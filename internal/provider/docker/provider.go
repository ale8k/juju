package docker

import (
	"github.com/juju/jsonschema"

	dockerclient "github.com/docker/docker/client"
	"github.com/juju/juju/caas"
	"github.com/juju/juju/cloud"
	"github.com/juju/juju/environs"
	environconfig "github.com/juju/juju/environs/config"
	environcontext "github.com/juju/juju/environs/context"
)

var _ caas.ContainerEnvironProvider = (*dockerProvider)(nil)

func init() {
	caas.RegisterContainerProvider("docker", &dockerProvider{})
}

type dockerProvider struct{}

// Open opens the environment and returns it. The configuration must
// have passed through PrepareConfig at some point in its lifecycle.
//
// Open should not perform any expensive operations, such as querying
// the cloud API, as it will be called frequently.
func (p *dockerProvider) Open(params environs.OpenParams) (caas.Broker, error) {
	// Get the Docker endpoint from the cloud configuration
	endpoint := params.Cloud.Endpoint

	// TODO: If it's a http protocol, create client with HTTP client
	c, err := dockerclient.NewClientWithOpts(
		dockerclient.WithHost(endpoint),
		dockerclient.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}

	broker := &dockerBroker{
		provider: p,
		client:   c,
		config:   params.Config,
	}

	return broker, nil
}

// Version returns the version of the provider. This is recorded as the
// environ version for each model, and used to identify which upgrade
// operations to run when upgrading a model's environ. Providers should
// start out at version 0.
func (p *dockerProvider) Version() int {
	return 0
}

// CloudSchema returns the schema used to validate input for add-cloud.  If
// a provider does not support custom clouds, CloudSchema should return
// nil.
func (p *dockerProvider) CloudSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:     []jsonschema.Type{jsonschema.ObjectType},
		Required: []string{cloud.EndpointKey},
		Properties: map[string]*jsonschema.Schema{
			cloud.EndpointKey: {
				Type:        []jsonschema.Type{jsonschema.StringType},
				Description: "Docker daemon endpoint (unix:///var/run/docker.sock or tcp://host:port)",
				Singular:    "Docker daemon endpoint (unix:///var/run/docker.sock or tcp://host:port)",
				Format:      jsonschema.FormatURI,
			},
		},
	}
}

// Ping tests the connection to the cloud, to verify the endpoint is valid.
func (p *dockerProvider) Ping(ctx environcontext.ProviderCallContext, endpoint string) error {
	// TODO: If it's a http protocol, create client with HTTP client
	// Create a temporary Docker client to test the endpoint
	client, err := dockerclient.NewClientWithOpts(
		dockerclient.WithHost(endpoint),
		dockerclient.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Ping(ctx)
	return err
}

// PrepareConfig prepares the configuration for a new model, based on
// the provided arguments. PrepareConfig is expected to produce a
// deterministic output. Any unique values should be based on the
// "uuid" attribute of the base configuration. This is called for the
// controller model during bootstrap, and also for new hosted models.
func (p *dockerProvider) PrepareConfig(prepareParams environs.PrepareConfigParams) (*environconfig.Config, error) {
	// Return the config as-is for now. Docker provider doesn't need special config preparation.
	return prepareParams.Config, nil
}

// Validate ensures that cfg is a valid configuration.
// If old is not nil, Validate should use it to determine
// whether a configuration change is valid.
func (p *dockerProvider) Validate(cfg, old *environconfig.Config) (valid *environconfig.Config, _ error) {
	// For now, just return the config as valid. Docker provider doesn't have special validation needs.
	return cfg, nil
}

// TODO: Credential schema, detection and finalisation need handling later for registries. For now we're just gonna
// use whatever default registry is in the daemon we connect to.

// CredentialSchemas returns credential schemas, keyed on
// authentication type. These may be used to validate existing
// credentials, or to generate new ones (e.g. to create an
// interactive form.)
func (p *dockerProvider) CredentialSchemas() map[cloud.AuthType]cloud.CredentialSchema {
	return map[cloud.AuthType]cloud.CredentialSchema{cloud.EmptyAuthType: {}}
}

// DetectCredentials automatically detects one or more credentials
// from the environment. This may involve, for example, inspecting
// environment variables, or reading configuration files in
// well-defined locations.
//
// If no credentials can be detected, DetectCredentials should
// return an error satisfying errors.IsNotFound.
//
// If cloud name is not passed (empty-string), all credentials are
// returned, otherwise only the credential for that cloud.
func (p *dockerProvider) DetectCredentials(cloudName string) (*cloud.CloudCredential, error) {
	// errors.NotFoundf("credentials")
	// like lxd / manual, we don't auto detect.
	return cloud.NewEmptyCloudCredential(), nil
}

// FinalizeCredential finalizes a credential, updating any attributes
// as necessary. This is always done client-side, when adding the
// credential to credentials.yaml and before uploading credentials to
// the controller. The provider may completely alter a credential, even
// going as far as changing the auth-type, but the output must be a
// fully formed credential.
//
// Docker does not need edits, so send registry credential as-is.
func (p *dockerProvider) FinalizeCredential(
	_ environs.FinalizeCredentialContext,
	args environs.FinalizeCredentialParams,
) (*cloud.Credential, error) {
	return &args.Credential, nil
}
