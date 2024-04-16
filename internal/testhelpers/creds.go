package testhelpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/environment/providers/local"
	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
)

const (
	nutanixEndpointDummyValue = "prism-test.nutanix.com"
	nutanixPortDummyValue     = "9440"
	nutanixUsernameDummyValue = "test"
	nutanixPasswordDummyValue = "test"
	nutanixInsecureDummyValue = false
	prismDevConfigFileName    = ".prism-dev.yaml"
	prismDevConfigFileEnvVar  = "PRISM_DEV_CONFIG"
)

type prismCentral struct {
	Endpoint              string `yaml:"endpoint"`
	Port                  string `yaml:"port"`
	Username              string `yaml:"username"`
	Password              string `yaml:"password"`
	Insecure              bool   `yaml:"insecure"`
	AdditionalTrustBundle string `yaml:"additionalTrustBundle"`
}

type devConf struct {
	PrismCentral prismCentral `yaml:"prismCentral"`
}

func prismDevConfigFilePath() (string, error) {
	// Check if the developer config file path is set in the environment.
	// If not, use the default path.
	prismDevConfigFilePath, ok := os.LookupEnv(prismDevConfigFileEnvVar)
	if ok {
		if _, err := os.Stat(prismDevConfigFilePath); err != nil {
			return "", fmt.Errorf("unable to find developer config file %q mentioned in environment variable %q: %w", prismDevConfigFilePath, prismDevConfigFileEnvVar, err)
		}

		path, err := filepath.Abs(prismDevConfigFilePath)
		if err != nil {
			return "", fmt.Errorf("unable to get absolute path of developer config file %q: %w", prismDevConfigFilePath, err)
		}

		return path, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Abs(filepath.Join(homeDir, prismDevConfigFileName))
}

func dummyDeveloperConfig() devConf {
	return devConf{
		PrismCentral: prismCentral{
			Endpoint: nutanixEndpointDummyValue,
			Port:     nutanixPortDummyValue,
			Username: nutanixUsernameDummyValue,
			Password: nutanixPasswordDummyValue,
			Insecure: nutanixInsecureDummyValue,
		},
	}
}

func getDeveloperConfig(t *testing.T, filePath string) devConf {
	conf := dummyDeveloperConfig()
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		t.Log("unable to read developer config file, using dummy values")
		return conf
	}

	dConf := devConf{}
	if err := yaml.Unmarshal(bytes, &dConf); err != nil {
		t.Log("unable to unmarshal developer config file, using dummy values")
		return conf
	}

	return dConf
}

// CredentialsFromEnvironment returns a Credentials object from the developer environment
// or dummy values if the developer environment is not configured.
// The prism developer config is configured by creating a `.prism-dev.yaml` file in the user's home directory.
// Alternatively, the path to the developer config file can be set in the environment variable `PRISM_DEV_CONFIG`.
// More information about the developer config file can be found in the DEVELOPMENT.md in repo root.
func CredentialsFromEnvironment(t *testing.T) prismgoclient.Credentials {
	endpoint := ManagementEndpointFromEnvironment(t)

	return prismgoclient.Credentials{
		Endpoint: endpoint.Address.Host,
		Port:     endpoint.Address.Port(),
		URL:      endpoint.Address.String(),
		Username: endpoint.Username,
		Password: endpoint.Password,
		Insecure: endpoint.Insecure,
	}
}

// ManagementEndpointFromEnvironment returns a ManagementEndpoint object from the developer environment
func ManagementEndpointFromEnvironment(t *testing.T) *types.ManagementEndpoint {
	confFile, err := prismDevConfigFilePath()
	require.NoError(t, err)
	conf := getDeveloperConfig(t, confFile)
	// NOTE: These environment variables are set to dummy values for the purpose of unit testing.
	// These environment variables need to be replaced with a real Prism Central endpoint and credentials
	// when recording the mocks. Once done with recording, please reset them back to the original dummy values.
	// and ensure that the actual credentials are not checked in to the repository.
	t.Setenv("NUTANIX_ENDPOINT", conf.PrismCentral.Endpoint)
	t.Setenv("NUTANIX_PORT", conf.PrismCentral.Port)
	t.Setenv("NUTANIX_USERNAME", conf.PrismCentral.Username)
	t.Setenv("NUTANIX_PASSWORD", conf.PrismCentral.Password)
	t.Setenv("NUTANIX_INSECURE", strconv.FormatBool(conf.PrismCentral.Insecure))
	endpoint, err := local.NewProvider().GetManagementEndpoint(nil)
	require.NoError(t, err)

	return endpoint
}
