package foundation

import (
	"fmt"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

const (
	absolutePath = "foundation"
	userAgent    = "foundation"
	clientName   = "foundation"
)

// Foundation internal with its services
type Client struct {
	// base client
	client *internal.Client

	// Service for Imaging Nodes and Cluster Creation
	NodeImaging NodeImagingService

	// Service for File Management in foundation VM
	FileManagement FileManagementService

	// Service for Networking apis in foundation VM
	Networking NetworkingService
}

// NewFoundationAPIClient creates a new Foundation API Client
func NewFoundationAPIClient(credentials prismgoclient.Credentials) (*Client, error) {
	if credentials.FoundationEndpoint == "" {
		return nil, fmt.Errorf("foundation Client is missing: %s %s",
			"Please provide required detail - %s in provider configuration.",
			strings.Join(credentials.RequiredFields[clientName], ", "))
	}
	// for foundation internal, url should be based on foundation's endpoint and port
	baseURL := fmt.Sprintf("%s:%s", credentials.FoundationEndpoint, credentials.FoundationPort)
	if !strings.HasPrefix(baseURL, string(internal.SchemeHTTP)) {
		baseURL = fmt.Sprintf("%s://%s", internal.SchemeHTTP, baseURL)
	}
	c, err := internal.NewClient(
		internal.WithCredentials(&credentials),
		internal.WithAbsolutePath(absolutePath),
		internal.WithUserAgent(userAgent),
		internal.WithBaseURL(baseURL),
	)
	if err != nil {
		return nil, err
	}

	foundationClient := &Client{
		client: c,
		NodeImaging: NodeImagingOperations{
			client: c,
		},
		FileManagement: FileManagementOperations{
			client: c,
		},
		Networking: NetworkingOperations{
			client: c,
		},
	}
	return foundationClient, nil
}
