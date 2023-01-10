package karbon

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"strings"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

const (
	absolutePath = "karbon"
	userAgent    = "nutanix"
	clientName   = "karbon"
)

// Client manages the V3 API
type Client struct {
	httpClient                    *internal.Client
	clientOpts                    []internal.ClientOption
	Cluster                       ClusterService
	PrivateRegistry               PrivateRegistryService
	Meta                          MetaService
	ClusterRegistrationOperations ClusterRegistrationOperations
}

// ClientOption is a type alias for functional options for Client
type ClientOption func(*Client) error

// WithRoundTripper overrides the transport for the underlying http client
// Overriding transport is useful for testing against API Mocks
// This is not recommended for production use
func WithRoundTripper(transport http.RoundTripper) ClientOption {
	return func(c *Client) error {
		c.clientOpts = append(c.clientOpts, internal.WithRoundTripper(transport))
		return nil
	}
}

// WithCertificate sets the certificate for the client
func WithCertificate(certificate *x509.Certificate) ClientOption {
	return func(c *Client) error {
		c.clientOpts = append(c.clientOpts, internal.WithCertificate(certificate))
		return nil
	}
}

// WithPEMEncodedCertBundle sets the certificates for the client
func WithPEMEncodedCertBundle(certBundle []byte) ClientOption {
	return func(c *Client) error {
		for block, rest := pem.Decode(certBundle); block != nil; block, rest = pem.Decode(rest) {
			if block.Type != "CERTIFICATE" {
				return fmt.Errorf("unexpected PEM block type %q: was expecting CERTIFICATE", block.Type)
			}
			certs, err := x509.ParseCertificates(block.Bytes)
			if err != nil {
				return err
			}
			for _, cert := range certs {
				c.clientOpts = append(c.clientOpts, internal.WithCertificate(cert))
			}
		}
		return nil
	}
}

// NewKarbonAPIClient return a internal to operate Karbon resources
func NewKarbonAPIClient(credentials prismgoclient.Credentials, opts ...ClientOption) (*Client, error) {
	if credentials.URL == "" || credentials.Username == "" || credentials.Password == "" {
		return nil, fmt.Errorf("karbon Client is missing: %s %s",
			"Please provide required details - %s in provider configuration",
			strings.Join(credentials.RequiredFields[clientName], ", "))
	}
	kc := &Client{
		clientOpts: []internal.ClientOption{
			internal.WithCredentials(&credentials),
			internal.WithUserAgent(userAgent),
			internal.WithAbsolutePath(absolutePath),
		},
	}
	for _, opt := range opts {
		if err := opt(kc); err != nil {
			return nil, err
		}
	}

	c, err := internal.NewClient(kc.clientOpts...)
	if err != nil {
		return nil, err
	}

	kc.httpClient = c
	kc.Cluster = ClusterOperations{httpClient: c}
	kc.PrivateRegistry = PrivateRegistryOperations{httpClient: c}
	kc.Meta = MetaOperations{httpClient: c}
	kc.ClusterRegistrationOperations = ClusterRegistrationOperations{httpClient: c}

	return kc, nil
}
