package foundationcentral

import (
	"fmt"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

const (
	libraryVersion = "v1"
	absolutePath   = "api/fc/" + libraryVersion
	userAgent      = "nutanix/" + libraryVersion
	clientName     = "foundation_central"
)

// Client manages the foundation central API
type Client struct {
	client  *internal.Client
	Service Service
}

// NewFoundationCentralClient return a internal to operate foundation central resources
func NewFoundationCentralClient(credentials prismgoclient.Credentials) (*Client, error) {
	var baseClient *internal.Client

	// check if all required fields are present. Else create an empty internal
	if credentials.Username != "" && credentials.Password != "" && credentials.Endpoint != "" {
		c, err := internal.NewClient(&credentials, userAgent, absolutePath, false)
		if err != nil {
			return nil, err
		}
		baseClient = c
	} else {
		errorMsg := fmt.Sprintf("Foundation Central Client is missing. "+
			"Please provide required details - %s in provider configuration.", strings.Join(credentials.RequiredFields[clientName], ", "))

		baseClient = &internal.Client{UserAgent: userAgent, ErrorMsg: errorMsg}
	}

	fc := &Client{
		client: baseClient,
		Service: Operations{
			client: baseClient,
		},
	}
	return fc, nil
}
