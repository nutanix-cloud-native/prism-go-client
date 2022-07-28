package local

/*
 Local environment provider which simply reads management endpoint,
 credentials and settings from OS environment variables.

 This environment is meant for local testing.
*/

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
)

const (
	endpointEnv   = "NUTANIX_ENDPOINT"
	userEnv       = "NUTANIX_USERNAME"
	passwordEnv   = "NUTANIX_PASSWORD"
	categoriesEnv = "NUTANIX_CATEGORIES"
)

type provider struct{}

func (prov *provider) GetManagementEndpoint(
	topology types.Topology,
) (*types.ManagementEndpoint, error) {
	endpoint := os.Getenv(endpointEnv)
	// No local environment defined
	if endpoint == "" {
		return nil, types.ErrNotFound
	}
	if !strings.HasPrefix(endpoint, "https://") {
		endpoint = fmt.Sprintf("https://%s", endpoint)
	}
	addr, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &types.ManagementEndpoint{
		Address: addr,
		ApiCredentials: types.ApiCredentials{
			Username: os.Getenv(userEnv),
			Password: os.Getenv(passwordEnv),
		},
	}, nil
}

func (prov *provider) Get(topology types.Topology, key string) (
	interface{}, error,
) {
	switch key {
	case types.CategoriesKey:
		return strings.Split(os.Getenv(categoriesEnv), ","), nil
	}
	return nil, types.ErrNotFound
}

func NewProvider() types.Provider {
	return &provider{}
}
