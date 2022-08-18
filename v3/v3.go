package v3

import (
	"crypto/x509"
	"fmt"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

const (
	libraryVersion = "v3"
	absolutePath   = "api/nutanix/" + libraryVersion
	userAgent      = "nutanix/" + libraryVersion
)

// Client manages the V3 API
type Client struct {
	client *internal.Client
	V3     Service

	clientOpts []internal.ClientOption
}

// ClientOption is a functional option for the Client
type ClientOption func(*Client) error

// WithCertificate sets the certificate for the client
func WithCertificate(certificate *x509.Certificate) ClientOption {
	return func(c *Client) error {
		c.clientOpts = append(c.clientOpts, internal.WithCertificate(certificate))
		return nil
	}
}

// WithRoundTripper overrides the transport for the underlying http client
// Overriding transport is useful for testing against API Mocks
// This is not recommended for production use
func WithRoundTripper(transport http.RoundTripper) ClientOption {
	return func(c *Client) error {
		c.clientOpts = append(c.clientOpts, internal.WithRoundTripper(transport))
		return nil
	}
}

// WithLogger sets the logger for the client
func WithLogger(logger *logr.Logger) ClientOption {
	return func(c *Client) error {
		c.clientOpts = append(c.clientOpts, internal.WithLogger(logger))
		return nil
	}
}

// NewV3Client return a internal to operate V3 resources
func NewV3Client(credentials prismgoclient.Credentials, opts ...ClientOption) (*Client, error) {
	v3Client := &Client{}
	for _, opt := range opts {
		if err := opt(v3Client); err != nil {
			return nil, err
		}
	}

	if credentials.Username == "" || credentials.Password == "" || credentials.Endpoint == "" {
		return nil, fmt.Errorf("username, password and endpoint are required")
	}

	v3Client.clientOpts = append(v3Client.clientOpts,
		internal.WithCredentials(&credentials),
		internal.WithUserAgent(userAgent),
		internal.WithAbsolutePath(absolutePath),
	)

	httpClient, err := internal.NewClient(v3Client.clientOpts...)
	if err != nil {
		return nil, err
	}

	v3Client.client = httpClient
	v3Client.V3 = Operations{client: httpClient}

	return v3Client, nil
}
