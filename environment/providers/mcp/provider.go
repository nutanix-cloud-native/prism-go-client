package mcp

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
)

// ModelContextClient defines an interface to interact with the model context protocol.
type ModelContextClient interface {
	// GetValue returns the value associated with the given key from the model context.
	GetValue(key string) (string, error)
}

// provider implements the types.Provider interface using modelcontextprotocol as a data source.
type provider struct {
	client ModelContextClient
}

// NewProvider creates a new modelcontextprotocol environment provider using the given client.
func NewProvider(client ModelContextClient) types.Provider {
	return &provider{client: client}
}

// GetManagementEndpoint fetches the management endpoint configuration from model context.
func (p *provider) GetManagementEndpoint(topology types.Topology) (*types.ManagementEndpoint, error) {
	endpoint, err := p.client.GetValue("endpoint")
	if err != nil || endpoint == "" {
		return nil, types.ErrNotFound
	}
	port, err := p.client.GetValue("port")
	if err != nil || port == "" {
		port = "9440" // Default port if not provided.
	}
	address := fmt.Sprintf("%s:%s", endpoint, port)
	if !strings.HasPrefix(address, "https://") {
		address = fmt.Sprintf("https://%s", address)
	}
	addr, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	insecureStr, _ := p.client.GetValue("insecure")
	insecureTLS := strings.ToLower(insecureStr) == "true"
	trustBundle, _ := p.client.GetValue("trustBundle")
	username, _ := p.client.GetValue("username")
	password, _ := p.client.GetValue("password")

	return &types.ManagementEndpoint{
		Address: addr,
		ApiCredentials: types.ApiCredentials{
			Username: username,
			Password: password,
		},
		Insecure:              insecureTLS,
		AdditionalTrustBundle: trustBundle,
	}, nil
}

// Get retrieves additional configuration values from model context.
// For example, if key is "categories", it splits the value on commas.
func (p *provider) Get(topology types.Topology, key string) (interface{}, error) {
	value, err := p.client.GetValue(key)
	if err != nil || value == "" {
		return nil, types.ErrNotFound
	}
	if key == types.CategoriesKey {
		return strings.Split(value, ","), nil
	}
	return value, nil
}
