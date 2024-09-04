package v3

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/keploy/go-sdk/integrations/khttpclient"
	"github.com/keploy/go-sdk/keploy"
	"github.com/keploy/go-sdk/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	"github.com/nutanix-cloud-native/prism-go-client/v3/models"
)

func setup(t *testing.T) (*http.ServeMux, *internal.Client, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	creds := &prismgoclient.Credentials{
		URL:      "",
		Username: "username",
		Password: "password",
		Port:     "",
		Endpoint: "0.0.0.0",
		Insecure: true,
	}
	c, err := internal.NewClient(
		internal.WithCredentials(creds),
		internal.WithUserAgent(userAgent),
		internal.WithAbsolutePath(absolutePath),
		internal.WithBaseURL(server.URL))
	require.NoError(t, err)

	return mux, c, server
}

func TestOperations_GetPrismCentral(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	// Changing insecure to true as that leads to modifying the transport underneath
	creds.Insecure = true

	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	v3Client, err := NewV3Client(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})
	pc, err := v3Client.V3.GetPrismCentral(kctx)
	require.NoError(t, err)
	assert.NotNil(t, pc)
	assert.Equal(t, "PC", *pc.Resources.Type)
	assert.Equal(t, "pc.2024.1.0.1", pc.Resources.Version)
}

func TestOperations_ListSubnet(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	v3Client, err := NewV3Client(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})

	subnets, err := v3Client.V3.ListSubnet(kctx, &models.SubnetListMetadata{})
	require.NoError(t, err)
	assert.Len(t, subnets.Entities, 3)
	for _, subnet := range subnets.Entities {
		assert.Equal(t, "subnet", subnet.Metadata.Kind)
	}
}

func TestOperations_GetSubnet(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	v3Client, err := NewV3Client(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})

	subnet, err := v3Client.V3.GetSubnet(kctx, "c7938dc6-7659-453e-a688-e26020c68e43")
	require.NoError(t, err)
	assert.NotNil(t, subnet)
	assert.Equal(t, "subnet", subnet.Metadata.Kind)
	assert.Equal(t, "sherlock_net", subnet.Spec.Name)
	assert.Nil(t, subnet.Spec.Resources.IsExternal)
}

func TestOperations_GetCluster(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	v3Client, err := NewV3Client(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})

	cluster, err := v3Client.V3.GetCluster(kctx, "0005b0f1-8f43-a0f2-02b7-3cecef193712")
	require.NoError(t, err)
	assert.NotNil(t, cluster)
	assert.Equal(t, "cluster", cluster.Metadata.Kind)
	assert.Equal(t, "ganon", *cluster.Status.Name)
}

func TestOperations_ListCluster(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	v3Client, err := NewV3Client(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})

	clusters, err := v3Client.V3.ListCluster(kctx, &models.ClusterListMetadata{})
	require.NoError(t, err)
	assert.Len(t, clusters.Entities, 2)
	assert.Equal(t, "cluster", clusters.Entities[0].Metadata.Kind)
	assert.Equal(t, "cluster", clusters.Entities[1].Metadata.Kind)
	assert.False(t, clusters.Entities[0].IsPrismCentral())
	assert.True(t, clusters.Entities[1].IsPrismCentral())

	prismElements := clusters.GetPrismElements()
	assert.Len(t, prismElements, 1)
	assert.Equal(t, "cluster", prismElements[0].Metadata.Kind)
	assert.Equal(t, "ganon", *prismElements[0].Status.Name)
}
