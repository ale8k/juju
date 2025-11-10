// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"

	environscloudspec "github.com/juju/juju/environs/cloudspec"

	dockerapi "github.com/docker/docker/api/types"
	container "github.com/docker/docker/api/types/container"
	mounttypes "github.com/docker/docker/api/types/mount"
	networktypes "github.com/docker/docker/api/types/network"
	volumetypes "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// PortMapping represents a mapping of container ports to host ports.
type PortMapping struct {
	ContainerPort int
	HostPort      int
	Protocol      string // tcp/udp
}

// VolumeMount represents a mount point for a volume in the container.
type VolumeMount struct {
	VolumeName string
	MountPath  string
	ReadOnly   bool
}

// ContainerSpec holds the specification for creating a container.
type ContainerSpec struct {
	Name        string
	Image       string
	Env         map[string]string
	Labels      map[string]string
	NetworkName string
	Ports       []PortMapping
	Mounts      []VolumeMount
	Cmd         []string // optional command/entrypoint override
	User        string   // optional user to run as (e.g., "root" or "170:170")
}

// Client abstracts the minimal Docker operations the broker needs.
// Initial implementation is now backed by the Docker SDK.
//
// NOTE: This is intentionally small; we'll extend as features land.

type Client interface {
	Ping(ctx context.Context) error
	EnsureNetwork(ctx context.Context, name string) error
	EnsureVolume(ctx context.Context, name string, sizeMiB int) error
	RunContainer(ctx context.Context, spec ContainerSpec) (id string, err error)
	RemoveContainer(ctx context.Context, id string) error                                       // new: stop and remove
	RemoveNetwork(ctx context.Context, name string) error                                       // new
	RemoveVolume(ctx context.Context, name string) error                                        // new
	WriteFile(ctx context.Context, containerID, path string, content []byte, mode uint32) error // new: copy file into container
	Close(ctx context.Context) error
}

// NewClientFromSpec creates a Client from the cloud spec.
func NewClientFromSpec(_ context.Context, spec environscloudspec.CloudSpec) (Client, error) {
	endpoint := spec.Endpoint
	if endpoint == "" {
		endpoint = "unix:///var/run/docker.sock"
	}
	// Validate scheme.
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid docker endpoint: %w", err)
	}
	_ = u // endpoint passed as-is to docker client

	cli, err := client.NewClientWithOpts(
		client.WithHost(endpoint),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}
	return &dockerSDKClient{cli: cli}, nil
}

type dockerSDKClient struct {
	cli *client.Client
}

func (d *dockerSDKClient) Ping(ctx context.Context) error {
	_, err := d.cli.Ping(ctx)
	return err
}

func (d *dockerSDKClient) EnsureNetwork(ctx context.Context, name string) error {
	// Inspect
	_, err := d.cli.NetworkInspect(ctx, name, dockerapi.NetworkInspectOptions{})
	if err == nil {
		return nil
	}
	// Create bridge network
	_, err = d.cli.NetworkCreate(ctx, name, dockerapi.NetworkCreate{Driver: "bridge"})
	return err
}

func (d *dockerSDKClient) EnsureVolume(ctx context.Context, name string, _ int) error {
	_, err := d.cli.VolumeInspect(ctx, name)
	if err == nil {
		return nil
	}
	// Updated to use CreateOptions (VolumeCreateBody removed in newer SDK).
	_, err = d.cli.VolumeCreate(ctx, volumetypes.CreateOptions{
		Name: name,
	})

	return err
}

func (d *dockerSDKClient) RunContainer(ctx context.Context, spec ContainerSpec) (string, error) {
	// Pull image (best effort)
	rc, err := d.cli.ImagePull(ctx, spec.Image, dockerapi.ImagePullOptions{})
	if err == nil {
		io.Copy(io.Discard, rc)
		_ = rc.Close()
	}

	env := make([]string, 0, len(spec.Env))
	for k, v := range spec.Env {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	exposed := nat.PortSet{}
	bindings := nat.PortMap{}
	for _, p := range spec.Ports {
		proto := strings.ToLower(strings.TrimSpace(p.Protocol))
		if proto == "" {
			proto = "tcp"
		}
		portKey, err := nat.NewPort(proto, fmt.Sprintf("%d", p.ContainerPort))
		if err != nil {
			return "", err
		}
		exposed[portKey] = struct{}{}
		if p.HostPort > 0 {
			bindings[portKey] = []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: fmt.Sprintf("%d", p.HostPort)}}
		}
	}

	mounts := []mounttypes.Mount{}
	for _, m := range spec.Mounts {
		mounts = append(mounts, mounttypes.Mount{
			Type:     mounttypes.TypeVolume,
			Source:   m.VolumeName,
			Target:   m.MountPath,
			ReadOnly: m.ReadOnly,
		})
	}

	containerCfg := &container.Config{
		Image:        spec.Image,
		Env:          env,
		Labels:       spec.Labels,
		ExposedPorts: exposed,
		WorkingDir:   "/var/lib/juju",
	}
	if len(spec.Cmd) > 0 {
		containerCfg.Cmd = spec.Cmd
	}
	if spec.User != "" {
		containerCfg.User = spec.User
	}
	hostCfg := &container.HostConfig{
		PortBindings: bindings,
		Mounts:       mounts,
		RestartPolicy: container.RestartPolicy{
			Name: "always", // replaced container.RestartPolicyAlways constant
		},
	}
	netCfg := &networktypes.NetworkingConfig{}
	if spec.NetworkName != "" {
		netCfg.EndpointsConfig = map[string]*networktypes.EndpointSettings{
			spec.NetworkName: {},
		}
	}

	resp, err := d.cli.ContainerCreate(ctx, containerCfg, hostCfg, netCfg, nil, spec.Name)
	if err != nil {
		return "", err
	}
	if err := d.cli.ContainerStart(ctx, resp.ID, dockerapi.ContainerStartOptions{}); err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (d *dockerSDKClient) RemoveContainer(ctx context.Context, id string) error {
	// Attempt stop then remove.
	err := d.cli.ContainerStop(ctx, id, container.StopOptions{})
	if err != nil {
		return err
	}
	return d.cli.ContainerRemove(ctx, id, dockerapi.ContainerRemoveOptions{Force: true})
}

func (d *dockerSDKClient) RemoveNetwork(ctx context.Context, name string) error {
	return d.cli.NetworkRemove(ctx, name)
}

func (d *dockerSDKClient) RemoveVolume(ctx context.Context, name string) error {
	return d.cli.VolumeRemove(ctx, name, true)
}

func (d *dockerSDKClient) WriteFile(ctx context.Context, containerID, path string, content []byte, mode uint32) error {
	// Create tar archive in-memory containing the single file at the desired absolute path.
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	// Docker expects unix style path; ensure no leading duplicate slashes.
	hdr := &tar.Header{
		Name: strings.TrimLeft(path, "/"), // tar header uses relative path inside archive
		Mode: int64(mode),
		Size: int64(len(content)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}
	if _, err := tw.Write(content); err != nil {
		return err
	}
	if err := tw.Close(); err != nil {
		return err
	}
	// Copy archive to container root.
	// Using "/" so that header Name is treated as absolute path components.
	return d.cli.CopyToContainer(ctx, containerID, "/", &buf, dockerapi.CopyToContainerOptions{AllowOverwriteDirWithFile: true})
}

func (d *dockerSDKClient) Close(ctx context.Context) error { return d.cli.Close() }

// TODO: Add label constants and filtering helpers.
