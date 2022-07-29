package karbon

import (
	"fmt"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

const (
	absolutePath = "karbon"
	userAgent    = "nutanix"
	clientName   = "karbon"
)

// Client manages the V3 API
type Client struct {
	httpClient      *internal.Client
	Cluster         ClusterService
	PrivateRegistry PrivateRegistryService
	Meta            MetaService
}

// NewKarbonAPIClient return a internal to operate Karbon resources
func NewKarbonAPIClient(credentials prismgoclient.Credentials) (*Client, error) {
	if credentials.URL == "" || credentials.Username == "" || credentials.Password == "" {
		return nil, fmt.Errorf("karbon Client is missing: %s %s",
			"Please provide required details - %s in provider configuration",
			strings.Join(credentials.RequiredFields[clientName], ", "))
	}
	c, err := internal.NewClient(
		internal.WithCredentials(&credentials),
		internal.WithUserAgent(userAgent),
		internal.WithAbsolutePath(absolutePath))
	if err != nil {
		return nil, err
	}

	f := &Client{
		httpClient: c,
		Cluster: ClusterOperations{
			httpClient: c,
		},
		PrivateRegistry: PrivateRegistryOperations{
			httpClient: c,
		},
		Meta: MetaOperations{
			httpClient: c,
		},
	}

	return f, nil
}
