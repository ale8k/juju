package docker

import (
	"fmt"
	"os"
	"strings"

	"github.com/juju/errors"
	"github.com/juju/juju/agent"
	"github.com/juju/juju/mongo"
	jujuversion "github.com/juju/juju/version"
	"github.com/juju/loggo"
	"github.com/juju/names/v5"
	"github.com/juju/version/v2"
	"github.com/mitchellh/go-linereader"

	"github.com/juju/juju/internal/provider/kubernetes/constants"

	stdcontext "context"

	dockercontainer "github.com/docker/docker/api/types/container"
	dockerimage "github.com/docker/docker/api/types/image"
	dockermount "github.com/docker/docker/api/types/mount"
	dockernetwork "github.com/docker/docker/api/types/network"
	dockerclient "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/juju/juju/caas"
	"github.com/juju/juju/cloudconfig/podcfg"
	"github.com/juju/juju/core/arch"
	"github.com/juju/juju/core/config"
	"github.com/juju/juju/core/constraints"
	"github.com/juju/juju/core/watcher"
	"github.com/juju/juju/docker"
	"github.com/juju/juju/environs"
	environsbootstrap "github.com/juju/juju/environs/bootstrap"
	environconfig "github.com/juju/juju/environs/config"
	"github.com/juju/juju/environs/context"
	"github.com/juju/juju/proxy"
	"github.com/juju/juju/storage"
)

var logger = loggo.GetLogger("juju.provider.docker")

// Compile-time check that dockerBroker implements caas.Broker
var _ caas.Broker = (*dockerBroker)(nil)

type dockerBroker struct {
	provider *dockerProvider
	client   *dockerclient.Client
	config   *environconfig.Config
}

// Provider returns the ContainerEnvironProvider that created this Broker.
func (b *dockerBroker) Provider() caas.ContainerEnvironProvider {
	return b.provider
}

// These methods are in call order of the bootstrap:
// - PrepareForBootstrap
// - ConstraintsValidator
// - Next the client checks the controller name
// - Bootstrap (but the finaliser doesn't run until later)
// - Config

// PrepareForBootstrap will be called very early in the bootstrap
// procedure to give an Environ a chance to perform interactive
// operations that are required for bootstrapping.
func (b *dockerBroker) PrepareForBootstrap(ctx environs.BootstrapContext, controllerName string) error {
	// TODO: Check container names, networks, etc don't have conflicts.
	return nil
}

// ConstraintsValidator returns a Validator instance which
// is used to validate and merge constraints.
func (b *dockerBroker) ConstraintsValidator(ctx context.ProviderCallContext) (_ constraints.Validator, _ error) {
	// TODO: Implement constraints validation for Docker. Not 100% sure what constraints make sense here.
	return constraints.NewValidator(), nil
}

// Bootstrap creates a new environment, and an instance to host the
// controller for that environment. The instance will have the
// series and architecture of the Environ's choice, constrained to
// those of the available tools. Bootstrap will return the instance's
// architecture, series, and a function that must be called to finalize
// the bootstrap process by transferring the tools and installing the
// initial Juju controller.
//
// It is possible to direct Bootstrap to use a specific architecture
// (or fail if it cannot start an instance of that architecture) by
// using an architecture constraint; this will have the effect of
// limiting the available tools to just those matching the specified
// architecture.
func (b *dockerBroker) Bootstrap(ctx environs.BootstrapContext, callCtx context.ProviderCallContext, args environs.BootstrapParams) (*environs.BootstrapResult, error) {
	// 1. Container Config

	// K8S does:
	// - VerifyConfig
	// - Initial model stuff
	// - Controller namespace / checks it exists
	// - COntroller annotations
	// - Creates controller stack
	// - Deploys stack
	//
	// See: internal/provider/kubernetes/k8s.go
	// See: internal/provider/kubernetes/bootstrap.go

	if !args.BootstrapBase.Empty() {
		return nil, errors.NotSupportedf("set base for bootstrapping to docker")
	}

	// This runs after GetService.
	finaliser := func(ctx environs.BootstrapContext, pcfg *podcfg.ControllerPodConfig, opts environs.BootstrapDialOpts) (err error) {

		var agentConfig agent.ConfigSetterWriter
		agentConfig, err = pcfg.AgentConfig(names.NewControllerAgentTag(pcfg.ControllerId))
		if err != nil {
			return errors.Trace(err)
		}

		si, ok := agentConfig.StateServingInfo()
		if !ok {
			return errors.NewNotValid(nil, "agent config has no state serving info")
		}

		// ensures shared-secret content.
		if si.SharedSecret == "" {
			// Generate a shared secret for the Mongo replica set.
			sharedSecret, err := mongo.GenerateSharedSecret()
			if err != nil {
				return errors.Trace(err)
			}
			si.SharedSecret = sharedSecret
		}

		agentConfig.SetStateServingInfo(si)
		pcfg.Bootstrap.StateServingInfo = si // TODO: Probably doesn't make sense for us?

		unitAgentConfig, err := pcfg.UnitAgentConfig()
		if err != nil {
			return errors.Trace(err)
		}

		_ = unitAgentConfig

		// They run this between every step (below) to cancel it and run cleanup.
		if environsbootstrap.IsContextDone(ctx.Context()) {
			return environsbootstrap.Cancelled()
		}

		bootstrapParamsContent, err := pcfg.Bootstrap.StateInitializationParams.Marshal()
		if err != nil {
			panic(err)
		}

		agentConfigFileContent, err := agentConfig.Render()
		if err != nil {
			panic(err)
		}

		unitAgentConfigFileContent, err := unitAgentConfig.Render()
		if err != nil {
			panic(err)
		}

		// TODO: replace with copy from in memory.
		// Now create files for controller
		// shared-secret: MongoDB replica set authentication key
		// Mount: MongoDB container -> /var/lib/juju/shared-secret
		os.WriteFile("./shared-secret", []byte(si.SharedSecret), 0400)

		// server.pem: TLS certificate for MongoDB and controller
		// Mount: MongoDB container -> /var/lib/juju/server.pem
		// Mount: Controller container -> /var/lib/juju/template-server.pem
		os.WriteFile("./server.pem", []byte(mongo.GenerateSSLKey(si.Cert, si.PrivateKey)), 0600)

		// bootstrap-params: State initialization parameters
		// Mount: Controller container -> /var/lib/juju/bootstrap-params
		os.WriteFile("bootstrap-params", bootstrapParamsContent, 0400)

		// agent.conf: Controller agent configuration
		// Mount: Controller container -> /var/lib/juju/agents/controller-0/template-agent.conf
		os.WriteFile("agent.conf", agentConfigFileContent, 0400)

		// unit-agent.conf: Unit agent configuration template
		// Mount: Controller container -> /var/lib/juju/template-agent.conf
		os.WriteFile("unit-agent.conf", unitAgentConfigFileContent, 0400)

		// Environment for container:
		// pcfg.AgentEnvironment - set these
		controllerUnitPassword := unitAgentConfig.OldPassword()
		apiInfo, ok := unitAgentConfig.APIInfo()
		if ok {
			controllerUnitPassword = apiInfo.Password
		}

		containerEnv := map[string]string{
			constants.EnvJujuK8sUnitPassword: controllerUnitPassword,
			"JUJU_CONTAINER_NAME":            "api-server", // prob not needed?
			"JUJU_PROVIDER_TYPE":             "docker",
		}
		_ = containerEnv

		// Now da gooood stuff, we can makez our containerz
		// und gets uz a controlla rollen

		// TODO: Where to persist the network ID?
		// One network for all controllers on this local provider?
		_, err = b.client.NetworkCreate(
			callCtx,
			"juju",
			dockernetwork.CreateOptions{
				Driver: "bridge",
				Scope:  "local",
			},
		)
		if err != nil {
			if !strings.Contains(err.Error(), "already exists") {
				panic(err)
			}
		}

		b.makesAMongoPlz(ctx.Context(), "juju", []byte(si.SharedSecret), []byte(mongo.GenerateSSLKey(si.Cert, si.PrivateKey)))
		// TODO: Wait for MongoDB to be ready

		return nil
	}

	return &environs.BootstrapResult{
		CaasBootstrapFinalizer: finaliser,
		Arch:                   arch.ARM64,
		Base:                   jujuversion.DefaultSupportedLTSBase(),
	}, nil
}

func (b *dockerBroker) makesAControllerPlz(ctx stdcontext.Context) {

}

func (b *dockerBroker) makesAMongoPlz(ctx stdcontext.Context, networkName string, sharedSecret, serverPem []byte) {
	progress, err := b.client.ImagePull(ctx, "jujusolutions/juju-db:4.4", dockerimage.PullOptions{})
	if err != nil {
		panic(err)
	}
	defer progress.Close()

	lr := linereader.New(progress)
	for line := range lr.Ch {
		fmt.Println(line)
	}

	containerConfig := &dockercontainer.Config{
		Image:      "jujusolutions/juju-db:4.4",
		Entrypoint: []string{},
		Cmd: []string{
			"/bin/sh", "-c",
			"chmod 400 /var/lib/juju/shared-secret && " +
				"chmod 600 /var/lib/juju/server.pem && " +
				"mkdir -p /data/db && " +
				"mongod --replSet juju --port 37017 --sslMode requireSSL --sslPEMKeyFile /var/lib/juju/server.pem --auth --keyFile /var/lib/juju/shared-secret",
		},
		ExposedPorts: nat.PortSet{
			"37017/tcp": struct{}{},
		},
	}

	hostConfig := &dockercontainer.HostConfig{
		Mounts: []dockermount.Mount{
			{
				Type:   dockermount.TypeBind,
				Source: "/home/ubuntu/repos/juju/shared-secret",
				Target: "/var/lib/juju/shared-secret",
			},
			{
				Type:   dockermount.TypeBind,
				Source: "/home/ubuntu/repos/juju/server.pem",
				Target: "/var/lib/juju/server.pem",
			},
		},
	}

	// 3. Create and start the container
	resp, err := b.client.ContainerCreate(
		ctx,
		containerConfig,
		hostConfig,
		nil,
		nil,
		"juju-db",
	)
	if err != nil {
		panic(err)
	}

	err = b.client.ContainerStart(ctx, resp.ID, dockercontainer.StartOptions{})
	if err != nil {
		panic(err)
	}
}

// Config returns the configuration data with which the Environ was created.
// Note that this is not necessarily current; the canonical location
// for the configuration data is stored in the state.
func (b *dockerBroker) Config() *environconfig.Config {
	return b.config
}

// SetConfig updates the Environ's configuration.
//
// Calls to SetConfig do not affect the configuration of
// values previously obtained from Storage.
func (b *dockerBroker) SetConfig(cfg *environconfig.Config) (_ error) {
	b.config = cfg
	return nil
}

// GetService returns the service for the specified application.
func (b *dockerBroker) GetService(appName string, mode caas.DeploymentMode, includeClusterIP bool) (_ *caas.Service, _ error) {
	// Would we return the docker addr here? it isn't deployed yet so how will we go about this?
	// Create the network earlier?
	return nil, nil
}

// PrecheckInstance performs a preflight check on the specified
// series and constraints, ensuring that they are possibly valid for
// creating an instance in this model.
//
// PrecheckInstance is best effort, and not guaranteed to eliminate
// all invalid parameters. If PrecheckInstance returns nil, it is not
// guaranteed that the constraints are valid; if a non-nil error is
// returned, then the constraints are definitely invalid.
func (b *dockerBroker) PrecheckInstance(_ context.ProviderCallContext, _ environs.PrecheckInstanceParams) (_ error) {
	panic("not implemented") // TODO: Implement
}

// Destroy shuts down all known machines and destroys the
// rest of the environment. Note that on some providers,
// very recently started instances may not be destroyed
// because they are not yet visible.
//
// When Destroy has been called, any Environ referring to the
// same remote environment may become invalid.
func (b *dockerBroker) Destroy(ctx context.ProviderCallContext) (_ error) {
	return errors.New("not implemented")
}

func (b *dockerBroker) DestroyController(ctx context.ProviderCallContext, controllerUUID string) (_ error) {
	return errors.New("not implemented")
}

// StorageProviderTypes returns the storage provider types
// contained within this registry.
//
// Determining the supported storage providers may be dynamic.
// Multiple calls for the same registry must return consistent
// results.
func (b *dockerBroker) StorageProviderTypes() (_ []storage.ProviderType, _ error) {
	panic("not implemented") // TODO: Implement
}

// StorageProvider returns the storage provider with the given
// provider type. StorageProvider must return an errors satisfying
// errors.IsNotFound if the registry does not contain the
// specified provider type.
func (b *dockerBroker) StorageProvider(_ storage.ProviderType) (_ storage.Provider, _ error) {
	panic("not implemented") // TODO: Implement
}

// Create creates the environment for a new hosted model.
//
// This will be called before any workers begin operating on the
// Environ, to give an Environ a chance to perform operations that
// are required for further use.
//
// Create is not called for the initial controller model; it is
// the Bootstrap method's job to create the controller model.
func (b *dockerBroker) Create(_ context.ProviderCallContext, _ environs.CreateParams) (_ error) {
	panic("not implemented") // TODO: Implement
}

// AdoptResources is called when the model is moved from one
// controller to another using model migration. Some providers tag
// instances, disks, and cloud storage with the controller UUID to
// aid in clean destruction. This method will be called on the
// environ for the target controller so it can update the
// controller tags for all of those things. For providers that do
// not track the controller UUID, a simple method returning nil
// will suffice. The version number of the source controller is
// provided for backwards compatibility - if the technique used to
// tag items changes, the version number can be used to decide how
// to remove the old tags correctly.
func (b *dockerBroker) AdoptResources(ctx context.ProviderCallContext, controllerUUID string, fromVersion version.Number) (_ error) {
	panic("not implemented") // TODO: Implement
}

// ValidateStorageClass returns an error if the storage config is not valid.
func (b *dockerBroker) ValidateStorageClass(config map[string]interface{}) (_ error) {
	panic("not implemented") // TODO: Implement
}

// Upgrade sets the OCI image for the app to the specified version.
func (b *dockerBroker) Upgrade(appName string, vers version.Number) (_ error) {
	panic("not implemented") // TODO: Implement
}

// APIVersion returns the version of the container orchestration layer.
func (b *dockerBroker) APIVersion() (_ string, _ error) {
	panic("not implemented") // TODO: Implement
}

// GetSecretToken returns the token content for the specified secret name.
func (b *dockerBroker) GetSecretToken(name string) (_ string, _ error) {
	panic("not implemented") // TODO: Implement
}

// Version returns cluster version information.
func (b *dockerBroker) Version() (_ *version.Number, _ error) {
	panic("not implemented") // TODO: Implement
}

// CheckCloudCredentials verifies that the provided cloud credentials
// are still valid for the cloud.
func (b *dockerBroker) CheckCloudCredentials() (_ error) {
	panic("not implemented") // TODO: Implement
}

// Application returns the broker interface for an Application
func (b *dockerBroker) Application(_ string, _ caas.DeploymentType) (_ caas.Application) {
	panic("not implemented") // TODO: Implement
}

// WatchUnits returns a watcher which notifies when there
// are changes to units of the specified application.
func (b *dockerBroker) WatchUnits(appName string, mode caas.DeploymentMode) (_ watcher.NotifyWatcher, _ error) {
	panic("not implemented") // TODO: Implement
}

// Units returns all units and any associated filesystems
// of the specified application. Filesystems are mounted
// via volumes bound to the unit.
func (b *dockerBroker) Units(appName string, mode caas.DeploymentMode) (_ []caas.Unit, _ error) {
	panic("not implemented") // TODO: Implement
}

// AnnotateUnit annotates the specified pod (name or uid) with a unit tag.
func (b *dockerBroker) AnnotateUnit(appName string, mode caas.DeploymentMode, podName string, unit names.UnitTag) (_ error) {
	panic("not implemented") // TODO: Implement
}

// WatchContainerStart returns a watcher which is notified when the specified container
// for each unit in the application is starting/restarting. Each string represents
// the provider id for the unit. If containerName is empty, then the first workload container
// is used.
func (b *dockerBroker) WatchContainerStart(appName string, containerName string) (_ watcher.StringsWatcher, _ error) {
	panic("not implemented") // TODO: Implement
}

// EnsureService creates or updates a service for pods with the given params.
func (b *dockerBroker) EnsureService(appName string, statusCallback caas.StatusCallbackFunc, params *caas.ServiceParams, numUnits int, config config.ConfigAttributes) (_ error) {
	panic("not implemented") // TODO: Implement
}

// DeleteService deletes the specified service with all related resources.
func (b *dockerBroker) DeleteService(appName string) (_ error) {
	panic("not implemented") // TODO: Implement
}

// ExposeService sets up external access to the specified service.
func (b *dockerBroker) ExposeService(appName string, resourceTags map[string]string, config config.ConfigAttributes) (_ error) {
	panic("not implemented") // TODO: Implement
}

// UnexposeService removes external access to the specified service.
func (b *dockerBroker) UnexposeService(appName string) (_ error) {
	panic("not implemented") // TODO: Implement
}

// WatchService returns a watcher which notifies when there
// are changes to the deployment of the specified application.
func (b *dockerBroker) WatchService(appName string, mode caas.DeploymentMode) (_ watcher.NotifyWatcher, _ error) {
	panic("not implemented") // TODO: Implement
}

// ModelOperatorExists indicates if the model operator for the given broker
// exists
func (b *dockerBroker) ModelOperatorExists() (_ bool, _ error) {
	panic("not implemented") // TODO: Implement
}

// EnsureModelOperator creates or updates a model operator pod for running
// model operations in a CAAS namespace/model
func (b *dockerBroker) EnsureModelOperator(modelUUID string, agentPath string, config *caas.ModelOperatorConfig) (_ error) {
	panic("not implemented") // TODO: Implement
}

// ModelOperator return the model operator config used to create the current
// model operator for this broker
func (b *dockerBroker) ModelOperator() (_ *caas.ModelOperatorConfig, _ error) {
	panic("not implemented") // TODO: Implement
}

// GetModelOperatorDeploymentImage returns the image used for the model operator deployment.
func (b *dockerBroker) GetModelOperatorDeploymentImage() (_ string, _ error) {
	panic("not implemented") // TODO: Implement
}

// OperatorExists indicates if the operator for the specified
// application exists, and whether the operator is terminating.
func (b *dockerBroker) OperatorExists(appName string) (_ caas.DeploymentState, _ error) {
	panic("not implemented") // TODO: Implement
}

// EnsureOperator creates or updates an operator pod for running
// a charm for the specified application.
func (b *dockerBroker) EnsureOperator(appName string, agentPath string, config *caas.OperatorConfig) (_ error) {
	panic("not implemented") // TODO: Implement
}

// DeleteOperator deletes the specified operator.
func (b *dockerBroker) DeleteOperator(appName string) (_ error) {
	panic("not implemented") // TODO: Implement
}

// Operator returns an Operator with current status and life details.
func (b *dockerBroker) Operator(_ string) (_ *caas.Operator, _ error) {
	panic("not implemented") // TODO: Implement
}

// WatchOperator returns a watcher which notifies when there
// are changes to the operator of the specified application.
func (b *dockerBroker) WatchOperator(_ string) (_ watcher.NotifyWatcher, _ error) {
	panic("not implemented") // TODO: Implement
}

// EnsureImageRepoSecret ensures the image pull secret gets created.
func (b *dockerBroker) EnsureImageRepoSecret(_ docker.ImageRepoDetails) (_ error) {
	panic("not implemented") // TODO: Implement
}

func (b *dockerBroker) ProxyToApplication(appName string, remotePort string) (_ proxy.Proxier, _ error) {
	panic("not implemented") // TODO: Implement
}
