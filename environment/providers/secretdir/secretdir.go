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

// read parameters return "" for non-existing keys
func readParam(path, key string) (string, error) {
	if b, err := os.ReadFile(filepath.Join(path, key)); err == nil {
		return string(b), err
	} else if os.IsNotExist(err) {
		return "", nil
	} else {
		return "", err
	}
}

// mangementEndpointFromSecretParams parses management endpoint and credentials from
// parameters embedded as data in a k8s secret.
func getMgmtEndpointFromSecretDir(path string) (
	*types.ManagementEndpoint, error,
) {
	certString, err := readParam(path, secretKeyCertName)
	if err != nil {
		return nil, err
	}
	if certString != "" {
		// TLS key pair
		endpoint, err := readParam(path, secretKeyCertEndpoint)
		if err != nil {
			return nil, err
		}
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
		str, err := readParam(path, secretKeyName)
		if err != nil {
			return nil, err
		}
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
