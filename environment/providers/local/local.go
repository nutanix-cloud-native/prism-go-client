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
	endpointEnv    = "NUTANIX_ENDPOINT"
	portEnv        = "NUTANIX_PORT"
	userEnv        = "NUTANIX_USERNAME"
	passwordEnv    = "NUTANIX_PASSWORD"
	apiKeyEnv      = "NUTANIX_API_KEY"
	insecureEnv    = "NUTANIX_INSECURE"
	trustBundleEnv = "NUTANIX_ADDITIONAL_TRUST_BUNDLE"
	categoriesEnv  = "NUTANIX_CATEGORIES"
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
	port := os.Getenv(portEnv)
	if port == "" {
		port = "9440"
	}
	address := fmt.Sprintf("%s:%s", endpoint, port)
	if !strings.HasPrefix(address, "https://") {
		address = fmt.Sprintf("https://%s", address)
	}
	addr, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	insecureTLS := os.Getenv(insecureEnv) == "true"
	trustBundle := os.Getenv(trustBundleEnv)
	
	credentials := &types.ApiCredentials{
		Username: os.Getenv(userEnv),
		Password: os.Getenv(passwordEnv),
		APIKey:   os.Getenv(apiKeyEnv),
	}
	if err := credentials.Validate(); err != nil {
		return nil, err
	}

	return &types.ManagementEndpoint{
		Address: addr,
		ApiCredentials: *credentials,
		Insecure:              insecureTLS,
		AdditionalTrustBundle: trustBundle,
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
