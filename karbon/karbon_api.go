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
	client          *internal.Client
	Cluster         ClusterService
	PrivateRegistry PrivateRegistryService
	Meta            MetaService
}

// NewKarbonAPIClient return a internal to operate Karbon resources
func NewKarbonAPIClient(credentials prismgoclient.Credentials) (*Client, error) {
	var baseClient *internal.Client

	// check if all required fields are present. Else create an empty internal
	if credentials.Username != "" && credentials.Password != "" && credentials.Endpoint != "" {
		c, err := internal.NewClient(&credentials, userAgent, absolutePath, false)
		if err != nil {
			return nil, err
		}
		baseClient = c
	} else {
		errorMsg := fmt.Sprintf("karbon Client is missing. "+
			"Please provide required details - %s in provider configuration.", strings.Join(credentials.RequiredFields[clientName], ", "))
		baseClient = &internal.Client{UserAgent: userAgent, ErrorMsg: errorMsg}
	}

	f := &Client{
		client: baseClient,
		Cluster: ClusterOperations{
			client: baseClient,
		},
		PrivateRegistry: PrivateRegistryOperations{
			client: baseClient,
		},
		Meta: MetaOperations{
			client: baseClient,
		},
	}

	return f, nil
}
