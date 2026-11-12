package v4

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientFilesNFSVersionDirect(t *testing.T) {
	var requestPath, authorization string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath = r.URL.Path
		authorization = r.Header.Get("Authorization")
		_, _ = w.Write([]byte(nameServicesResponse("NFSV4")))
	}))
	defer server.Close()

	client := newFilesTestClient(t, server, true)
	version, err := client.FilesNFSVersion(context.Background())

	require.NoError(t, err)
	assert.Equal(t, FilesNFSVersionV4, version)
	assert.Equal(t, "/api/files/v4.0.a2/config/file-server/name-services", requestPath)
	assert.NotEmpty(t, authorization)
}

func TestClientFilesNFSVersionFromPC(t *testing.T) {
	const extID = "d8b59823-3416-4149-62c2-c8e86e1990d1"

	tests := []struct {
		name       string
		identifier string
	}{
		{name: "name", identifier: "prod-fs"},
		{name: "external ID", identifier: extID},
		{name: "FQDN", identifier: "prod-fs.example.com"},
		{name: "external IP", identifier: "10.15.98.173"},
		{name: "virtual IP", identifier: "10.15.98.174"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nameServicesPath string
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/api/files/v4.0.a2/config/file-servers":
					_, _ = w.Write([]byte(`{"data":[{"name":"files-a","extId":"other"},{"name":"prod-fs","extId":"` + extID + `","dnsDomainName":"example.com","externalNetworks":[{"ipAddresses":[{"ipv4":{"value":"10.15.98.173"}}],"virtualIpAddress":{"ipv4":{"value":"10.15.98.174"}}}]}]}`))
				default:
					nameServicesPath = r.URL.Path
					_, _ = w.Write([]byte(nameServicesResponse("NFSV3V4")))
				}
			}))
			defer server.Close()

			client := newFilesTestClient(t, server, true)
			version, err := client.FilesNFSVersionFromPC(context.Background(), tt.identifier)

			require.NoError(t, err)
			assert.Equal(t, FilesNFSVersionV3V4, version)
			assert.Equal(t, "/api/files/v4.0.a2/config/file-servers/"+extID+"/name-services", nameServicesPath)
		})
	}
}

func TestClientFilesNFSVersionFromPCMalformedResponse(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/files/v4.0.a2/config/file-servers" {
			_, _ = w.Write([]byte(`{"data":[{"name":"prod-fs","extId":"id-1"}]}`))
			return
		}
		_, _ = w.Write([]byte(`{"data":{"nfsVersion":"NFSV3"}}`))
	}))
	defer server.Close()

	client := newFilesTestClient(t, server, true)
	assert.NotPanics(t, func() {
		version, err := client.FilesNFSVersionFromPC(context.Background(), "prod-fs")
		assert.Error(t, err)
		assert.Equal(t, FilesNFSVersionUnknown, version)
	})
}

func TestClientFilesNFSVersionFromPCFileServerNotFound(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"data":[]}`))
	}))
	defer server.Close()

	client := newFilesTestClient(t, server, true)
	version, err := client.FilesNFSVersionFromPC(context.Background(), "missing")

	assert.ErrorContains(t, err, "not found")
	assert.Equal(t, FilesNFSVersionUnknown, version)
}

func TestClientFilesNFSVersionDirectUsesAPIKey(t *testing.T) {
	var apiKey, authorization string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey = r.Header.Get("X-ntnx-api-key")
		authorization = r.Header.Get("Authorization")
		_, _ = w.Write([]byte(nameServicesResponse("NFSV4")))
	}))
	defer server.Close()

	client := newFilesTestClientWithCredentials(t, server, prismgoclient.Credentials{
		APIKey:   "api-key",
		Insecure: true,
	})
	_, err := client.FilesNFSVersion(context.Background())

	require.NoError(t, err)
	assert.Equal(t, "api-key", apiKey)
	assert.Empty(t, authorization)
}

func TestClientFilesNFSVersionDirectAllowsExplicitInsecureTLS(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(nameServicesResponse("NFSV4")))
	}))
	defer server.Close()

	client := newFilesTestClientWithCredentials(t, server, prismgoclient.Credentials{
		Username: "admin",
		Password: "password",
		Insecure: true,
	})
	version, err := client.FilesNFSVersion(context.Background())

	require.NoError(t, err)
	assert.Equal(t, FilesNFSVersionV4, version)
}

func newFilesTestClient(t *testing.T, server *httptest.Server, insecure bool) *Client {
	t.Helper()

	credentials := prismgoclient.Credentials{
		Username: "admin",
		Password: "password",
		Insecure: insecure,
	}
	return newFilesTestClientWithCredentials(t, server, credentials)
}

func newFilesTestClientWithCredentials(
	t *testing.T,
	server *httptest.Server,
	credentials prismgoclient.Credentials,
) *Client {
	t.Helper()

	credentials.Endpoint = strings.TrimPrefix(server.URL, "https://")
	client, err := NewClient(credentials)
	require.NoError(t, err)
	return client
}

func nameServicesResponse(version string) string {
	return `{"data":{"$objectType":"files.v4.config.NameServiceSpec","nfsVersion":"` + version + `"}}`
}
