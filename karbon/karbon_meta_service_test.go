package karbon

import (
	"net/http"
	"testing"

	"github.com/keploy/go-sdk/integrations/khttpclient"
	"github.com/keploy/go-sdk/keploy"
	"github.com/keploy/go-sdk/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/environment/providers/local"
)

func testCredsFromEnv(t *testing.T) prismgoclient.Credentials {
	// NOTE: These environment variables are set to dummy values for the purpose of unit testing.
	// These environment variables need to be replaced with a real Prism Central endpoint and credentials
	// when recording the mocks. Once done with recording, please reset them back to the original dummy values.
	// and ensure that the actual credentials are not checked in to the repository.
	t.Setenv("NUTANIX_ENDPOINT", "prism-test.nutanix.com")
	t.Setenv("NUTANIX_PORT", "9440")
	t.Setenv("NUTANIX_USERNAME", "test")
	t.Setenv("NUTANIX_PASSWORD", "test")
	provider, err := local.NewProvider().GetManagementEndpoint(nil)
	require.NoError(t, err)

	return prismgoclient.Credentials{
		URL:      provider.Address.String(),
		Username: provider.Username,
		Password: provider.Password,
	}
}

func TestMetaOperations_GetVersion(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testCredsFromEnv(t)
	kc, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})
	v, err := kc.Meta.GetVersion(kctx)
	assert.NoError(t, err)

	assert.Equal(t, "2.6.0", *v.Version)
}

func TestMetaOperations_GetSemanticVersion(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testCredsFromEnv(t)
	kc, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})
	v, err := kc.Meta.GetSemanticVersion(kctx)
	assert.NoError(t, err)

	assert.Equal(t, int64(2), v.MajorVersion)
	assert.Equal(t, int64(6), v.MinorVersion)
	assert.Equal(t, int64(0), v.RevisionVersion)
}
