package v4

import (
	"net/http"
	"testing"

	"github.com/keploy/go-sdk/integrations/khttpclient"
	"github.com/keploy/go-sdk/keploy"
	"github.com/keploy/go-sdk/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/internal"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
)

func TestActuatorGetVersionRoutes(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testhelpers.CredentialsFromEnvironment(t)

	actuatorAPI, err := NewActuatorAPI(creds, internal.WithRoundTripper(interceptor))
	require.NoError(t, err)
	require.NotNil(t, actuatorAPI)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})

	versionRoutes, err := actuatorAPI.GetVersionRoutes(kctx)
	require.NoError(t, err)
	require.NotNil(t, versionRoutes)

	// Basic validation
	assert.Greater(t, len(versionRoutes), 0, "Expected at least one namespace")

	for _, namespace := range versionRoutes {
		assert.NotEmpty(t, namespace.Namespace, "Namespace should not be empty")
		assert.Greater(t, len(namespace.VersionRoutes), 0, "Expected at least one version route for namespace %s", namespace.Namespace)

		for _, versionRoute := range namespace.VersionRoutes {
			assert.NotEmpty(t, versionRoute.Version, "Version should not be empty")
			assert.Greater(t, len(versionRoute.Routes), 0, "Expected at least one route for version %s", versionRoute.Version)
		}
	}
}
