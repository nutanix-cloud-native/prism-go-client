package secretdir

/*
 Read credentials in CSI format from k8s secret mounted at $CSI_SECRET_DIR.
*/

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
)

const (
	secretKeyName         = "key"
	secretKeyCertName     = "cert"
	secretKeyCertEndpoint = "endpoint"
	envSecretDir          = "CSI_SECRET_DIR"
)

var ErrMissingValue = fmt.Errorf("missing information in secret value " +
	"'<prism-ip>:<prism-port>:<user>:<password>'")

type provider struct{}

// GetManagementEndpoint returns PC credentials independent of HCI cluster
func (prov *provider) GetManagementEndpoint(
	topology types.Topology,
) (*types.ManagementEndpoint, error) {
	path := os.Getenv(envSecretDir)
	if path == "" {
		return nil, fmt.Errorf("environment variable %s not set", envSecretDir)
	}
	return getMgmtEndpointFromSecretDir(path)
}

// Get doesn't return any settings from CSI secret
func (prov *provider) Get(topology types.Topology, key string) (
	interface{}, error,
) {
	return nil, types.ErrNotFound
}

func NewProvider() types.Provider {
	return &provider{}
}

func getMgmtEndpointFromSecretDir(path string) (*types.ManagementEndpoint, error) {
	secretParams := map[string]string{}
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.Type().IsRegular() {
			b, err := os.ReadFile(filepath.Join(path, file.Name()))
			if err != nil {
				return nil, err
			}
			secretParams[file.Name()] = string(b)
		}
	}
	return mangementEndpointFromSecretParams(secretParams)
}

// mangementEndpointFromSecretParams parses management endpoint and credentials from
// parameters embedded as data in a k8s secret.
func mangementEndpointFromSecretParams(secretParams map[string]string) (
	*types.ManagementEndpoint, error,
) {
	certString := secretParams[secretKeyCertName]
	if certString != "" {
		// TLS key pair
		endpoint := secretParams[secretKeyCertEndpoint]
		if endpoint == "" {
			return nil, fmt.Errorf("no endpoint is provided in secret value")
		}
		addr, err := url.Parse(fmt.Sprintf("https://%s", endpoint))
		if err != nil {
			return nil, err
		}
		return &types.ManagementEndpoint{
			ApiCredentials: types.ApiCredentials{
				KeyPair: certString,
			},
			Address: addr,
		}, nil
	} else {
		// basic auth
		str := secretParams[secretKeyName]
		creds := strings.SplitN(str, ":", 4)
		if len(creds) != 4 {
			return nil, ErrMissingValue
		}
		addr, err := url.Parse(fmt.Sprintf("https://%s:%s", creds[0], creds[1]))
		if err != nil {
			return nil, err
		}
		return &types.ManagementEndpoint{
			ApiCredentials: types.ApiCredentials{
				Username: creds[2],
				Password: creds[3],
			},
			Address: addr,
		}, nil
	}
}
