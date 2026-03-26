package client

import (
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-retryablehttp"

	"nusights-api/client/config"
)

const (
	defaultRetryMax     = 4
	defaultRetryWaitMin = 1 * time.Second
	defaultRetryWaitMax = 30 * time.Second
)

// RetryConfig holds retry and exponential backoff settings.
type RetryConfig struct {
	RetryMax     int
	RetryWaitMin time.Duration
	RetryWaitMax time.Duration

	// CheckRetry specifies the policy for handling retries, and is called after
	// each request. If nil, the default policy (retry on 5xx except 501 and
	// connection errors) is used.
	CheckRetry retryablehttp.CheckRetry

	// Backoff specifies the policy for how long to wait between retries.
	// If nil, exponential backoff with jitter is used.
	Backoff retryablehttp.Backoff
}

// DefaultRetryConfig returns a RetryConfig with sensible defaults:
// 4 retries, 1s min wait, 30s max wait, exponential backoff with jitter.
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		RetryMax:     defaultRetryMax,
		RetryWaitMin: defaultRetryWaitMin,
		RetryWaitMax: defaultRetryWaitMax,
	}
}

// ApplyRetryToTransport wraps the given go-openapi transport's HTTP
// RoundTripper with hashicorp/go-retryablehttp to add automatic retry with
// exponential backoff. The transport must be the concrete *client.Runtime type
// returned by httptransport.New().
func ApplyRetryToTransport(transport *httptransport.Runtime, cfg *RetryConfig) {
	if cfg == nil {
		cfg = DefaultRetryConfig()
	}

	rc := retryablehttp.NewClient()
	rc.RetryMax = cfg.RetryMax
	rc.RetryWaitMin = cfg.RetryWaitMin
	rc.RetryWaitMax = cfg.RetryWaitMax
	rc.Logger = nil // silence default logger; callers can set their own

	if cfg.CheckRetry != nil {
		rc.CheckRetry = cfg.CheckRetry
	}
	if cfg.Backoff != nil {
		rc.Backoff = cfg.Backoff
	}

	// Preserve the existing base RoundTripper (e.g. TLS config) if one is set.
	if transport.Transport != nil {
		rc.HTTPClient.Transport = transport.Transport
	}

	transport.Transport = &retryablehttp.RoundTripper{Client: rc}
}

// NewHTTPClientWithRetry creates a NutanixInsightsRESTAPIs client with
// automatic retry and exponential backoff on the underlying HTTP transport.
// If retryCfg is nil, DefaultRetryConfig() is used.
func NewHTTPClientWithRetry(formats strfmt.Registry, cfg *TransportConfig, retryCfg *RetryConfig) *NutanixInsightsRESTAPIs {
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	ApplyRetryToTransport(transport, retryCfg)
	return New(transport, formats)
}

// NewConfigClientWithBasicAuthAndRetry creates a config API client with basic
// auth credentials and automatic retry with exponential backoff.
// If retryCfg is nil, DefaultRetryConfig() is used.
func NewConfigClientWithBasicAuthAndRetry(host, basePath, scheme, user, password string, retryCfg *RetryConfig) config.ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	ApplyRetryToTransport(transport, retryCfg)
	return config.New(transport, strfmt.Default)
}

// NewConfigClientWithBearerTokenAndRetry creates a config API client with a
// bearer token and automatic retry with exponential backoff.
// If retryCfg is nil, DefaultRetryConfig() is used.
func NewConfigClientWithBearerTokenAndRetry(host, basePath, scheme, bearerToken string, retryCfg *RetryConfig) config.ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	ApplyRetryToTransport(transport, retryCfg)
	return config.New(transport, strfmt.Default)
}
