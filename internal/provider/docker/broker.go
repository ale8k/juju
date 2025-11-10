// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package docker

import (
	"context"

	"encoding/base64"

	"github.com/juju/errors"
	jujuversion "github.com/juju/juju/core/version"
	"github.com/juju/names/v6"

	"github.com/juju/juju/caas"
	"github.com/juju/juju/core/constraints"
	"github.com/juju/juju/core/instance"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/core/semversion"
	"github.com/juju/juju/environs"
	environscloudspec "github.com/juju/juju/environs/cloudspec"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/internal/cloudconfig/podcfg"
	"github.com/juju/juju/internal/docker"
	"github.com/juju/juju/internal/proxy"
	"github.com/juju/juju/internal/storage"
)

// dockerBroker minimal CAAS broker.
// NOTE: This is an early stub; many Broker interface methods intentionally
// return NotSupported.
type dockerBroker struct {
	prov           *environProvider
	controllerUUID string
	cloud          environscloudspec.CloudSpec
	cfg            *config.Config
	client         Client
	controllerCID  string // docker container id for controller
}

func newDockerBroker(p *environProvider, controllerUUID string, cloud environscloudspec.CloudSpec, cfg *config.Config) *dockerBroker {
	return &dockerBroker{prov: p, controllerUUID: controllerUUID, cloud: cloud, cfg: cfg}
}

// Provider returns the creating provider.
func (b *dockerBroker) Provider() caas.ContainerEnvironProvider { return b.prov }

// Configer
func (b *dockerBroker) Config() *config.Config { return b.cfg }
func (b *dockerBroker) SetConfig(_ context.Context, newCfg *config.Config) error {
	// Accept updated model configuration (eg agent-version) during bootstrap.
	if newCfg == nil {
		return errors.NotValidf("nil config")
	}
	b.cfg = newCfg
	return nil
}

// ConstraintsChecker
func (b *dockerBroker) ConstraintsValidator(context.Context) (constraints.Validator, error) {
	return constraints.NewValidator(), nil
}

// APIVersion returns a static API version string for the docker broker stub.
func (b *dockerBroker) APIVersion() (string, error) { return "docker-0", nil }

// Close no-op.
func (b *dockerBroker) Close(ctx context.Context) error {
	if b.client != nil {
		_ = b.client.Close(ctx)
	}
	return nil
}

// --- Minimal interface implementations ---
// The caas.Broker interface is large. We stub the essential embedded
// sub-interfaces with NotSupported errors where appropriate until
// actual Docker functionality is added.

// InstancePrechecker
func (b *dockerBroker) PrecheckInstance(context.Context, environs.PrecheckInstanceParams) error {
	return errors.NotSupportedf("precheck instance")
}

// BootstrapEnviron
func (b *dockerBroker) PrepareForBootstrap(ctx environs.BootstrapContext, controllerName string) error {
	// Ensure a per-model network exists; name derived from model UUID.
	if b.client == nil {
		return nil
	}
	return b.client.EnsureNetwork(ctx, "juju-"+b.cfg.UUID())
}
func (b *dockerBroker) Bootstrap(ctx environs.BootstrapContext, params environs.BootstrapParams) (*environs.BootstrapResult, error) {
	if b.client == nil {
		return nil, errors.NotSupportedf("bootstrap without docker client")
	}
	modelNet := "juju-" + b.cfg.UUID()
	if err := b.client.EnsureNetwork(ctx, modelNet); err != nil {
		return nil, errors.Annotate(err, "ensure model network")
	}
	// Ensure a persistent volume for controller state.
	ctrlVol := "juju-controller-" + b.cfg.UUID()
	if err := b.client.EnsureVolume(ctx, ctrlVol, 20480); err != nil { // 20GiB default
		return nil, errors.Annotate(err, "ensure controller volume")
	}

	// Defer container creation to CaasBootstrapFinalizer so we can use the fully populated pod config
	// (including Bootstrap.StateInitializationParams) to inject bootstrap-params at startup.
	res := &environs.BootstrapResult{
		Arch: "amd64", // TODO: detect arch
		CaasBootstrapFinalizer: func(bc environs.BootstrapContext, pc *podcfg.ControllerPodConfig, _ environs.BootstrapDialOpts) error {
			if pc == nil || pc.Bootstrap == nil {
				return errors.NotValidf("nil controller pod config")
			}

			bp, err := pc.Bootstrap.StateInitializationParams.Marshal()
			if err != nil {
				return errors.Annotate(err, "marshal bootstrap-params")
			}
			// Render controller agent config (state machine) to pre-create agent.conf like k8s path.
			agentCfg, err := pc.AgentConfig(names.NewControllerAgentTag(pc.ControllerId))
			if err != nil {
				return errors.Annotate(err, "render controller agent config")
			}
			agentBytes, err := agentCfg.Render()
			if err != nil {
				return errors.Annotate(err, "render controller agent config contents")
			}
			encodedAgent := base64.StdEncoding.EncodeToString(agentBytes)
			encoded := base64.StdEncoding.EncodeToString(bp)

			jujuVer := jujuversion.Current.String()
			image := "ghcr.io/juju/jujud-operator:" + jujuVer
			apiPort := pc.Bootstrap.ControllerAgentInfo.APIPort
			internalPort := 17022

			script := `echo "Starting Juju controller agent (docker provider)..." 1>&2
sleep 2
export JUJU_DATA_DIR=/var/lib/juju
arch=$(uname -m); case "$arch" in aarch64) arch=arm64;; x86_64) arch=amd64;; *) arch=$arch;; esac
VER="` + jujuVer + `-ubuntu-${arch}"
export JUJU_TOOLS_ROOT=$JUJU_DATA_DIR/tools
export JUJU_TOOLS_DIR=$JUJU_TOOLS_ROOT/$VER
mkdir -p $JUJU_TOOLS_DIR $JUJU_DATA_DIR/agents/controller-0
ln -snf $JUJU_TOOLS_DIR $JUJU_TOOLS_ROOT/controller-0

if [ ! -f $JUJU_DATA_DIR/bootstrap-params ]; then
  echo "$BOOTSTRAP_PARAMS_B64" | base64 -d > $JUJU_DATA_DIR/bootstrap-params || { echo "failed writing bootstrap-params" >&2; exit 1; }
  chmod 600 $JUJU_DATA_DIR/bootstrap-params || true
fi
if [ ! -f $JUJU_DATA_DIR/agents/controller-0/agent.conf ]; then
  echo "$AGENT_CONF_B64" | base64 -d > $JUJU_DATA_DIR/agents/controller-0/agent.conf || { echo "failed writing agent.conf" >&2; exit 1; }
  chmod 600 $JUJU_DATA_DIR/agents/controller-0/agent.conf || true
fi

install -m0755 /opt/jujud $JUJU_TOOLS_DIR/jujud

# Ensure required tooling exists
for bin in tar sha256sum stat; do
  if ! command -v $bin >/dev/null 2>&1; then
    echo "$bin not found in image; cannot build tools tarball" >&2
    exit 1
  fi
done

TOOLS_TAR=$JUJU_TOOLS_DIR/tools.tar.gz
if [ ! -f "$TOOLS_TAR" ]; then
  ( cd $JUJU_TOOLS_DIR && tar -czf tools.tar.gz jujud ) || { echo "failed to create tools.tar.gz" >&2; exit 1; }
fi
sha=$(sha256sum "$TOOLS_TAR" | awk '{print $1}')
size=$(stat -c%s "$TOOLS_TAR")
printf '{"version":"%s","url":"","sha256":"%s","size":%s}' "$VER" "$sha" "$size" > $JUJU_TOOLS_DIR/downloaded-tools.txt
sha256sum "$TOOLS_TAR" > $JUJU_TOOLS_DIR/juju${VER}.sha256

if ! grep -q 'old-password' $JUJU_DATA_DIR/agents/controller-0/agent.conf 2>/dev/null; then
  $JUJU_TOOLS_DIR/jujud bootstrap-state --data-dir $JUJU_DATA_DIR --show-log --timeout 20m0s
  rc=$?
  if [ $rc -ne 0 ]; then
    echo "bootstrap-state failed (rc=$rc)" >&2
    sleep 5
    exit $rc
  fi
fi

mkdir -p /var/lib/pebble/default/layers
cat > /var/lib/pebble/default/layers/001-jujud.yaml <<EOF
summary: jujud service
services:
    jujud:
        summary: Juju controller agent
        startup: enabled
        override: replace
        command: $JUJU_TOOLS_DIR/jujud machine --data-dir $JUJU_DATA_DIR --controller-id 0 --log-to-stderr --show-log
EOF

exec /opt/pebble run --http :38811 --verbose`
			spec := ContainerSpec{
				Name:  "juju-controller",
				Image: image,
				Env: map[string]string{
					"JUJU_MODEL_UUID":      b.cfg.UUID(),
					"JUJU_CONTROLLER_UUID": b.controllerUUID,
					"JUJU_CONTAINER_NAME":  "api-server",
					"JUJU_CONTAINER_NAMES": "api-server",
					"BOOTSTRAP_PARAMS_B64": encoded,
					"AGENT_CONF_B64":       encodedAgent,
				},
				Labels: map[string]string{
					"juju-model":         b.cfg.UUID(),
					"juju-controller":    b.controllerUUID,
					"juju.is-controller": "true",
				},
				NetworkName: modelNet,
				Ports: []PortMapping{
					{ContainerPort: apiPort, HostPort: apiPort, Protocol: "tcp"},
					{ContainerPort: internalPort, HostPort: internalPort, Protocol: "tcp"},
					{ContainerPort: 38811, HostPort: 38811, Protocol: "tcp"},
				},
				Mounts: []VolumeMount{{VolumeName: ctrlVol, MountPath: "/var/lib/juju", ReadOnly: false}},
				Cmd:    []string{script},
				User:   "root",
			}
			id, err := b.client.RunContainer(context.Background(), spec)
			if err != nil {
				return errors.Annotate(err, "run controller container")
			}
			b.controllerCID = id
			return nil
		},
	}
	return res, nil
}

func (b *dockerBroker) Destroy(_ context.Context) error {
	// Best effort model teardown.
	if b.client != nil {
		modelNet := "juju-" + b.cfg.UUID()
		_ = b.client.RemoveNetwork(context.Background(), modelNet)
		ctrlVol := "juju-controller-" + b.cfg.UUID()
		_ = b.client.RemoveVolume(context.Background(), ctrlVol)
	}
	return nil
}

func (b *dockerBroker) DestroyController(_ context.Context, _ string) error {
	if b.client == nil {
		return nil
	}
	if b.controllerCID != "" {
		_ = b.client.RemoveContainer(context.Background(), b.controllerCID)
	}
	modelNet := "juju-" + b.cfg.UUID()
	_ = b.client.RemoveNetwork(context.Background(), modelNet)
	ctrlVol := "juju-controller-" + b.cfg.UUID()
	_ = b.client.RemoveVolume(context.Background(), ctrlVol)
	return nil
}

// ServiceManager
func (b *dockerBroker) GetService(_ context.Context, appName string, _ bool) (*caas.Service, error) {
	// Minimal implementation to satisfy bootstrap controllerDataRefresher.
	if appName != "controller" { // k8s constant JujuControllerStackName
		return nil, errors.NotSupportedf("get service for %s", appName)
	}
	addrs := network.ProviderAddresses{network.NewMachineAddress("127.0.0.1").AsProviderAddress()}
	return &caas.Service{Id: appName, Addresses: addrs}, nil
}

// ResourceAdopter
func (b *dockerBroker) AdoptResources(context.Context, string, semversion.Number) error { return nil }

// Networking (return not supported for all methods)
func (b *dockerBroker) Subnets(context.Context, []network.Id) ([]network.SubnetInfo, error) {
	return nil, errors.NotSupportedf("subnets")
}
func (b *dockerBroker) NetworkInterfaces(context.Context, []instance.Id) ([]network.InterfaceInfos, error) {
	return nil, errors.NotSupportedf("network interfaces")
}
func (b *dockerBroker) SupportsSpaces() (bool, error) { return false, errors.NotSupportedf("spaces") }
func (b *dockerBroker) SupportsSpaceDiscovery() (bool, error) {
	return false, errors.NotSupportedf("space discovery")
}
func (b *dockerBroker) Spaces(context.Context) (network.SpaceInfos, error) {
	return nil, errors.NotSupportedf("spaces")
}
func (b *dockerBroker) ProviderSpaceInfo(context.Context, *network.SpaceInfo) (*environs.ProviderSpaceInfo, error) {
	return nil, errors.NotSupportedf("provider space info")
}
func (b *dockerBroker) SupportsContainerAddresses() bool { return false }
func (b *dockerBroker) AllocateContainerAddresses(context.Context, instance.Id, string, network.InterfaceInfos) (network.InterfaceInfos, error) {
	return nil, errors.NotSupportedf("allocate container addresses")
}
func (b *dockerBroker) ReleaseContainerAddresses(context.Context, []string) error {
	return errors.NotSupportedf("release container addresses")
}

// Storage ProviderRegistry
func (b *dockerBroker) RecommendedPoolForKind(storage.StorageKind) *storage.Config { return nil }
func (b *dockerBroker) StorageProviderTypes() ([]storage.ProviderType, error)      { return nil, nil }
func (b *dockerBroker) StorageProvider(storage.ProviderType) (storage.Provider, error) {
	return nil, errors.NotSupportedf("storage provider")
}

// CredentialChecker
func (b *dockerBroker) CheckCloudCredentials(context.Context) error { return nil }

// ApplicationBroker
func (b *dockerBroker) Application(string, caas.DeploymentType) caas.Application { return nil }
func (b *dockerBroker) Units(context.Context, string) ([]caas.Unit, error) {
	return nil, errors.NotSupportedf("units")
}
func (b *dockerBroker) AnnotateUnit(context.Context, string, string, names.UnitTag) error {
	return errors.NotSupportedf("annotate unit")
}

// ModelOperatorManager
func (b *dockerBroker) ModelOperatorExists(context.Context) (bool, error) {
	return false, errors.NotSupportedf("model operator exists")
}
func (b *dockerBroker) EnsureModelOperator(context.Context, string, string, *caas.ModelOperatorConfig) error {
	return errors.NotSupportedf("ensure model operator")
}
func (b *dockerBroker) ModelOperator(context.Context) (*caas.ModelOperatorConfig, error) {
	return nil, errors.NotSupportedf("model operator config")
}
func (b *dockerBroker) GetModelOperatorDeploymentImage(context.Context) (string, error) {
	return "", errors.NotSupportedf("model operator deployment image")
}

// EnsureImageRepoSecret
func (b *dockerBroker) EnsureImageRepoSecret(context.Context, docker.ImageRepoDetails) error {
	return errors.NotSupportedf("ensure image repo secret")
}

// ProxyManager
func (b *dockerBroker) ProxyToApplication(context.Context, string, string) (proxy.Proxier, error) {
	return nil, errors.NotSupportedf("proxy to application")
}

// StorageValidator
func (b *dockerBroker) ValidateStorageClass(context.Context, map[string]interface{}) error {
	return errors.NotSupportedf("validate storage class")
}

// Upgrader
func (b *dockerBroker) Upgrade(context.Context, string, semversion.Number) error {
	return errors.NotSupportedf("upgrade")
}

// GetSecretToken
func (b *dockerBroker) GetSecretToken(context.Context, string) (string, error) {
	return "", errors.NotSupportedf("secret token")
}

// ClusterVersionGetter
func (b *dockerBroker) Version() (*semversion.Number, error) {
	return nil, errors.NotSupportedf("cluster version")
}
