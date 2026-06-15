package v4

import (
	"net/url"
	"testing"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/stretchr/testify/assert"

	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
)

func TestNewClientCacheReturnsNewCache(t *testing.T) {
	cache := NewClientCache()

	assert.NotNil(t, cache)
}

type cachedClientParams struct {
	name         string
	mgmtEndpoint types.ManagementEndpoint
}

func (c *cachedClientParams) Key() string {
	return c.name
}

func (c *cachedClientParams) ManagementEndpoint() types.ManagementEndpoint {
	return c.mgmtEndpoint
}

func TestGetReturnsClientIfPresentInCache(t *testing.T) {
	cache := NewClientCache()
	client := &Client{}
	cache.set("cluster1", "bm90X2FfcmVhbF9oYXNo", client)

	returnedClient, hash, err := cache.get("cluster1")

	assert.NoError(t, err)
	assert.Equal(t, "bm90X2FfcmVhbF9oYXNo", hash)
	assert.Equal(t, client, returnedClient)
}

func TestGetReturnsErrorIfClientNotPresentInCache(t *testing.T) {
	cache := NewClientCache()

	_, hash, err := cache.get("cluster1")

	assert.Equal(t, "", hash)
	assert.ErrorIs(t, err, types.ErrorClientNotFound)
}

func TestAddAddsClientToCache(t *testing.T) {
	cache := NewClientCache()
	client := &Client{}

	cache.set("cluster1", "bm90X2FfcmVhbF9oYXNo", client)

	returnedClient, hash, err := cache.get("cluster1")

	assert.NoError(t, err)
	assert.Equal(t, "bm90X2FfcmVhbF9oYXNo", hash)
	assert.Equal(t, client, returnedClient)
}

func TestAddOverwritesExistingClientInCache(t *testing.T) {
	cache := NewClientCache()
	client1 := &Client{}
	client2 := &Client{}

	cache.set("cluster1", "bm90X2FfcmVhbF9oYXNo", client1)
	cache.set("cluster1", "ZGVmaW5pdGVseV9ub3RfYV9yZWFsX2hhc2gK", client2)

	returnedClient, hash, err := cache.get("cluster1")

	assert.NoError(t, err)
	assert.Equal(t, "ZGVmaW5pdGVseV9ub3RfYV9yZWFsX2hhc2gK", hash)
	assert.Equal(t, client2, returnedClient)
}

func TestDeleteRemovesClientFromCache(t *testing.T) {
	cache := NewClientCache()
	client := &Client{}
	cache.set("cluster1", "bm90X2FfcmVhbF9oYXNo", client)

	cache.Delete(&cachedClientParams{name: "cluster1"})

	_, hash, err := cache.get("cluster1")
	assert.Equal(t, "", hash)
	assert.ErrorIs(t, err, types.ErrorClientNotFound)
}

func TestDeleteDoesNotErrorIfClientNotPresentInCache(t *testing.T) {
	cache := NewClientCache()

	cache.Delete(&cachedClientParams{name: "cluster1"}) // No error expected
}

func TestGetOrCreateNilAddressValidation(t *testing.T) {
	cache := NewClientCache()
	params := &cachedClientParams{
		name: "test-client",
		mgmtEndpoint: types.ManagementEndpoint{
			ApiCredentials: types.ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address: nil, // This should trigger the nil check
		},
	}

	client, err := cache.GetOrCreate(params)

	assert.Error(t, err)
	assert.Nil(t, client)
	assert.Contains(t, err.Error(), "management endpoint address is nil for cachedClientParams with key test-client")
}

func TestGetOrCreateEmptyHostValidation(t *testing.T) {
	cache := NewClientCache()
	params := &cachedClientParams{
		name: "test-client",
		mgmtEndpoint: types.ManagementEndpoint{
			ApiCredentials: types.ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address: &url.URL{Host: ""}, // Empty host should trigger validation
		},
	}

	client, err := cache.GetOrCreate(params)

	assert.Error(t, err)
	assert.Nil(t, client)
	assert.Contains(t, err.Error(), "management endpoint address host is empty for cachedClientParams with key test-client")
}

func TestGetOrCreateEmptyUsernameValidation(t *testing.T) {
	cache := NewClientCache()
	params := &cachedClientParams{
		name: "test-client",
		mgmtEndpoint: types.ManagementEndpoint{
			ApiCredentials: types.ApiCredentials{
				Username: "", // Empty username should trigger validation
				Password: "testpass",
			},
			Address: &url.URL{Host: "test.com"},
		},
	}

	client, err := cache.GetOrCreate(params)

	assert.Error(t, err)
	assert.Nil(t, client)
	assert.Contains(t, err.Error(), "API credentials(either a username and password, or an API key) is required for authentication")
}

func TestGetOrCreateEmptyPasswordValidation(t *testing.T) {
	cache := NewClientCache()
	params := &cachedClientParams{
		name: "test-client",
		mgmtEndpoint: types.ManagementEndpoint{
			ApiCredentials: types.ApiCredentials{
				Username: "testuser",
				Password: "", // Empty password should trigger validation
			},
			Address: &url.URL{Host: "test.com"},
		},
	}

	client, err := cache.GetOrCreate(params)

	assert.Error(t, err)
	assert.Nil(t, client)
	assert.Contains(t, err.Error(), "API credentials(either a username and password, or an API key) is required for authentication")
}

func TestValidateManagementEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		endpoint types.ManagementEndpoint
		key      string
		wantErr  bool
		errMsg   string
	}{
		{
			name: "valid endpoint",
			endpoint: types.ManagementEndpoint{
				ApiCredentials: types.ApiCredentials{
					Username: "testuser",
					Password: "testpass",
				},
				Address:  &url.URL{Host: "test.com"},
				Insecure: true,
			},
			key:     "test-key",
			wantErr: false,
		},
		{
			name: "nil address",
			endpoint: types.ManagementEndpoint{
				ApiCredentials: types.ApiCredentials{
					Username: "testuser",
					Password: "testpass",
				},
				Address: nil,
			},
			key:     "test-key",
			wantErr: true,
			errMsg:  "management endpoint address is nil for cachedClientParams with key test-key",
		},
		{
			name: "empty host",
			endpoint: types.ManagementEndpoint{
				ApiCredentials: types.ApiCredentials{
					Username: "testuser",
					Password: "testpass",
				},
				Address: &url.URL{Host: ""},
			},
			key:     "test-key",
			wantErr: true,
			errMsg:  "management endpoint address host is empty for cachedClientParams with key test-key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateManagementEndpoint(tt.endpoint, tt.key)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetOrCreateSetsVerifySSL(t *testing.T) {
	cache := NewClientCache()
	mgmtEndpoint := testhelpers.ManagementEndpointFromEnvironment(t)
	cp := &cachedClientParams{
		name:         "cluster1",
		mgmtEndpoint: *mgmtEndpoint,
	}
	c, err := cache.GetOrCreate(cp)
	assert.NoError(t, err)
	assert.True(t, c.CategoriesApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.ClustersApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.ImagesApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.StorageContainerAPI.ApiClient.VerifySSL)
	assert.True(t, c.SubnetsApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.SubnetIPReservationApi.ApiClient.VerifySSL)
	assert.True(t, c.TasksApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.VolumeGroupsApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.VmApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.UsersApiInstance.ApiClient.VerifySSL)

	cache.Delete(cp)

	certBundle := `-----BEGIN CERTIFICATE-----
MIIEYDCCA0igAwIBAgILBAAAAAABL07hRQwwDQYJKoZIhvcNAQEFBQAwVzELMAkG
A1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv
b3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw0xMTA0MTMxMDAw
MDBaFw0yMjA0MTMxMDAwMDBaMF0xCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i
YWxTaWduIG52LXNhMTMwMQYDVQQDEypHbG9iYWxTaWduIE9yZ2FuaXphdGlvbiBW
YWxpZGF0aW9uIENBIC0gRzIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB
AQDdNR3yIFQmGtDvpW+Bdllw3Of01AMkHyQOnSKf1Ccyeit87ovjYWI4F6+0S3qf
ZyEcLZVUunm6tsTyDSF0F2d04rFkCJlgePtnwkv3J41vNnbPMYzl8QbX3FcOW6zu
zi2rqqlwLwKGyLHQCAeV6irs0Z7kNlw7pja1Q4ur944+ABv/hVlrYgGNguhKujiz
4MP0bRmn6gXdhGfCZsckAnNate6kGdn8AM62pI3ffr1fsjqdhDFPyGMM5NgNUqN+
ARvUZ6UYKOsBp4I82Y4d5UcNuotZFKMfH0vq4idGhs6dOcRmQafiFSNrVkfB7cVT
5NSAH2v6gEaYsgmmD5W+ZoiTAgMBAAGjggElMIIBITAOBgNVHQ8BAf8EBAMCAQYw
EgYDVR0TAQH/BAgwBgEB/wIBADAdBgNVHQ4EFgQUXUayjcRLdBy77fVztjq3OI91
nn4wRwYDVR0gBEAwPjA8BgRVHSAAMDQwMgYIKwYBBQUHAgEWJmh0dHBzOi8vd3d3
Lmdsb2JhbHNpZ24uY29tL3JlcG9zaXRvcnkvMDMGA1UdHwQsMCowKKAmoCSGImh0
dHA6Ly9jcmwuZ2xvYmFsc2lnbi5uZXQvcm9vdC5jcmwwPQYIKwYBBQUHAQEEMTAv
MC0GCCsGAQUFBzABhiFodHRwOi8vb2NzcC5nbG9iYWxzaWduLmNvbS9yb290cjEw
HwYDVR0jBBgwFoAUYHtmGkUNl8qJUC99BM00qP/8/UswDQYJKoZIhvcNAQEFBQAD
ggEBABvgiADHBREc/6stSEJSzSBo53xBjcEnxSxZZ6CaNduzUKcbYumlO/q2IQen
fPMOK25+Lk2TnLryhj5jiBDYW2FQEtuHrhm70t8ylgCoXtwtI7yw07VKoI5lkS/Z
9oL2dLLffCbvGSuXL+Ch7rkXIkg/pfcNYNUNUUflWP63n41edTzGQfDPgVRJEcYX
pOBWYdw9P91nbHZF2krqrhqkYE/Ho9aqp9nNgSvBZnWygI/1h01fwlr1kMbawb30
hag8IyrhFHvBN91i0ZJsumB9iOQct+R2UTjEqUdOqCsukNK1OFHrwZyKarXMsh3o
wFZUTKiL8IkyhtyTMr5NGvo1dbU=
-----END CERTIFICATE-----`
	cp.mgmtEndpoint.AdditionalTrustBundle = certBundle
	c, err = cache.GetOrCreate(cp)
	assert.NoError(t, err)
	assert.False(t, c.CategoriesApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.ClustersApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.ImagesApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.StorageContainerAPI.ApiClient.VerifySSL)
	assert.False(t, c.SubnetsApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.SubnetIPReservationApi.ApiClient.VerifySSL)
	assert.False(t, c.TasksApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.VolumeGroupsApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.VmApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.UsersApiInstance.ApiClient.VerifySSL)
}

func TestValidateCredentials(t *testing.T) {
	tests := []struct {
		name        string
		credentials prismgoclient.Credentials
		key         string
		wantErr     string
	}{
		{
			name: "valid basic auth credentials",
			credentials: prismgoclient.Credentials{
				Username: "testuser",
				Password: "testpass",
			},
			key:     "test-key-basic",
			wantErr: "",
		},
		{
			name: "valid api key credentials",
			credentials: prismgoclient.Credentials{
				APIKey: "test-api-key",
			},
			key:     "test-key-api",
			wantErr: "",
		},
		{
			name: "missing username for basic auth",
			credentials: prismgoclient.Credentials{
				Password: "testpass",
			},
			key:     "test-key-missing-username",
			wantErr: "API credentials(either a username and password, or an API key) is required for authentication in cachedClientParams with key test-key-missing-username",
		},
		{
			name: "missing password for basic auth",
			credentials: prismgoclient.Credentials{
				Username: "testuser",
			},
			key:     "test-key-missing-password",
			wantErr: "API credentials(either a username and password, or an API key) is required for authentication in cachedClientParams with key test-key-missing-password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCredentials(tt.credentials, tt.key)
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
