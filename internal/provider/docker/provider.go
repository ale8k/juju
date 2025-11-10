// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package docker

import (
	context "context"
	"net/url"

	"github.com/juju/errors"
	"github.com/juju/jsonschema"
	"github.com/juju/schema"

	"github.com/juju/juju/caas"
	"github.com/juju/juju/cloud"
	"github.com/juju/juju/environs"
	environscloudspec "github.com/juju/juju/environs/cloudspec"
	"github.com/juju/juju/environs/config"
	internallogger "github.com/juju/juju/internal/logger"
)

var logger = internallogger.GetLogger("juju.provider.docker")

// environProvider implements a minimal CAAS-oriented provider backed by Docker.
type environProvider struct{}

func init() {
	// Register as CAAS container provider (not a traditional IaaS environ).
	caas.RegisterContainerProvider("docker", &environProvider{})
	// Note: not registering an IaaS environs provider; docker is CAAS-only.
}

// Version implements environs.EnvironProvider (unused for CAAS broker but required).
func (p *environProvider) Version() int { return 1 }

// Open returns a caas.Broker (DockerBroker) placeholder implementing minimal CAAS operations.
func (p *environProvider) Open(ctx context.Context, args environs.OpenParams, _ environs.CredentialInvalidator) (caas.Broker, error) {
	if err := p.ValidateCloud(ctx, args.Cloud); err != nil {
		return nil, errors.Annotate(err, "validating cloud spec")
	}
	if _, err := p.Validate(ctx, args.Config, nil); err != nil {
		return nil, errors.Annotate(err, "validating model config")
	}
	cl, err := NewClientFromSpec(ctx, args.Cloud)
	if err != nil {
		return nil, errors.Annotate(err, "creating docker client")
	}
	// Verify connectivity to the docker daemon.
	if err := cl.Ping(ctx); err != nil {
		return nil, errors.Annotate(err, "ping docker daemon")
	}
	broker := newDockerBroker(p, args.ControllerUUID, args.Cloud, args.Config)
	broker.client = cl
	return broker, nil
}

// CloudSchema returns the schema for adding new clouds of this type.
func (p *environProvider) CloudSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:     []jsonschema.Type{jsonschema.ObjectType},
		Required: []string{cloud.EndpointKey, cloud.AuthTypesKey, cloud.RegionsKey},
		Order:    []string{cloud.EndpointKey, cloud.AuthTypesKey, cloud.RegionsKey},
		Properties: map[string]*jsonschema.Schema{
			cloud.EndpointKey: {
				Singular: "the Docker daemon address or URL (unix:///var/run/docker.sock or tcp://host:port)",
				Type:     []jsonschema.Type{jsonschema.StringType},
				Format:   jsonschema.FormatURI,
			},
			cloud.AuthTypesKey: {
				Type: []jsonschema.Type{jsonschema.ArrayType},
				Enum: []interface{}{[]string{string(cloud.EmptyAuthType)}},
			},
			cloud.RegionsKey: {
				Type:                 []jsonschema.Type{jsonschema.ObjectType},
				Singular:             "region",
				Plural:               "regions",
				AdditionalProperties: &jsonschema.Schema{Type: []jsonschema.Type{jsonschema.ObjectType}, MaxProperties: jsonschema.Int(0)},
			},
		},
	}
}

// CredentialSchemas declares supported auth.
func (*environProvider) CredentialSchemas() map[cloud.AuthType]cloud.CredentialSchema {
	return map[cloud.AuthType]cloud.CredentialSchema{cloud.EmptyAuthType: {}}
}

// DetectCredentials returns empty credentials.
func (*environProvider) DetectCredentials(string) (*cloud.CloudCredential, error) {
	return cloud.NewEmptyCloudCredential(), nil
}

// FinalizeCredential passthrough.
func (*environProvider) FinalizeCredential(_ environs.FinalizeCredentialContext, args environs.FinalizeCredentialParams) (*cloud.Credential, error) {
	return &args.Credential, nil
}

// DetectRegions returns a single default region.
func (*environProvider) DetectRegions() ([]cloud.Region, error) {
	return []cloud.Region{{Name: "default"}}, nil
}

// Ping tests the connection to the Docker endpoint.
func (*environProvider) Ping(ctx context.Context, endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return errors.New("invalid endpoint format; use unix:///path or tcp://host:port")
	}
	switch u.Scheme {
	case "unix", "tcp", "http", "https":
		return nil
	default:
		return errors.Errorf("unsupported scheme %q", u.Scheme)
	}
}

// ValidateCloud validates the cloud spec.
func (*environProvider) ValidateCloud(ctx context.Context, spec environscloudspec.CloudSpec) error {
	if err := spec.Validate(); err != nil {
		return errors.Trace(err)
	}
	if spec.Credential == nil {
		return errors.NotValidf("missing credential")
	}
	return nil
}

// Validate implements environs.EnvironProvider.
func (*environProvider) Validate(ctx context.Context, cfg, old *config.Config) (*config.Config, error) {
	// Reuse base config validation; no extra attrs initially.
	if err := config.Validate(ctx, cfg, old); err != nil {
		return nil, err
	}
	return cfg, nil
}

// --- environs.ModelConfigProvider (no-op) ---
// These allow the bootstrap "model defaults" path to query provider defaults
// without reporting "not supported" for the docker CAAS provider.
func (*environProvider) ConfigDefaults() schema.Defaults { return schema.Defaults{} }
func (*environProvider) ConfigSchema() schema.Fields     { return schema.Fields{} }
func (*environProvider) ModelConfigDefaults(context.Context) (map[string]any, error) {
	return map[string]any{}, nil
}
