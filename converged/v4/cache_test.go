package v4

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
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

func TestGetOrCreateWithInvalidEndpoint(t *testing.T) {
	cache := NewClientCache()
	params := &cachedClientParams{
		name: "test-client",
		mgmtEndpoint: types.ManagementEndpoint{
			ApiCredentials: types.ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address: nil, // This should trigger validation in the underlying v4 SDK
		},
	}

	client, err := cache.GetOrCreate(params)

	// The converged cache delegates to v4 SDK cache, which handles validation
	assert.Error(t, err)
	assert.Nil(t, client)
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

	// Check that the underlying v4 SDK client has VerifySSL set correctly
	// The converged client wraps the v4 SDK client, so we need to access it through the client field
	assert.True(t, c.client.CategoriesApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.client.ClustersApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.client.ImagesApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.client.StorageContainerAPI.ApiClient.VerifySSL)
	assert.True(t, c.client.SubnetsApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.client.SubnetIPReservationApi.ApiClient.VerifySSL)
	assert.True(t, c.client.TasksApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.client.VolumeGroupsApiInstance.ApiClient.VerifySSL)
	assert.True(t, c.client.VmApiInstance.ApiClient.VerifySSL)

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

	// With trust bundle, VerifySSL should be false (insecure)
	assert.False(t, c.client.CategoriesApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.client.ClustersApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.client.ImagesApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.client.StorageContainerAPI.ApiClient.VerifySSL)
	assert.False(t, c.client.SubnetsApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.client.SubnetIPReservationApi.ApiClient.VerifySSL)
	assert.False(t, c.client.TasksApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.client.VolumeGroupsApiInstance.ApiClient.VerifySSL)
	assert.False(t, c.client.VmApiInstance.ApiClient.VerifySSL)
}

func TestGetOrCreateWithSessionAuth(t *testing.T) {
	cache := NewClientCache()
	mgmtEndpoint := testhelpers.ManagementEndpointFromEnvironment(t)
	cp := &cachedClientParams{
		name:         "cluster1",
		mgmtEndpoint: *mgmtEndpoint,
	}
	c, err := cache.GetOrCreate(cp)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.NotNil(t, c.client)
}

func TestGetOrCreateCaching(t *testing.T) {
	cache := NewClientCache()
	mgmtEndpoint := testhelpers.ManagementEndpointFromEnvironment(t)
	cp := &cachedClientParams{
		name:         "cluster1",
		mgmtEndpoint: *mgmtEndpoint,
	}

	// First call should create the client
	c1, err := cache.GetOrCreate(cp)
	assert.NoError(t, err)
	assert.NotNil(t, c1)

	// Second call should return the same client from cache
	c2, err := cache.GetOrCreate(cp)
	assert.NoError(t, err)
	assert.NotNil(t, c2)
	assert.Equal(t, c1, c2)
}

func TestGetOrCreateWithDifferentEndpoints(t *testing.T) {
	cache := NewClientCache()
	mgmtEndpoint := testhelpers.ManagementEndpointFromEnvironment(t)

	cp1 := &cachedClientParams{
		name:         "cluster1",
		mgmtEndpoint: *mgmtEndpoint,
	}

	cp2 := &cachedClientParams{
		name:         "cluster2",
		mgmtEndpoint: *mgmtEndpoint,
	}

	c1, err := cache.GetOrCreate(cp1)
	assert.NoError(t, err)
	assert.NotNil(t, c1)

	c2, err := cache.GetOrCreate(cp2)
	assert.NoError(t, err)
	assert.NotNil(t, c2)

	// Different keys should result in different clients
	assert.NotEqual(t, c1, c2)
}

func TestCacheInvalidationDeletesV4SDKClient(t *testing.T) {
	cache := NewClientCache()
	mgmtEndpoint := testhelpers.ManagementEndpointFromEnvironment(t)
	cp := &cachedClientParams{
		name:         "cluster1",
		mgmtEndpoint: *mgmtEndpoint,
	}

	// Create a client and verify it exists in the converged cache
	c1, err := cache.GetOrCreate(cp)
	assert.NoError(t, err)
	assert.NotNil(t, c1)

	// Verify the client exists in the converged cache
	_, _, err = cache.get("cluster1")
	assert.NoError(t, err)

	// Store reference to the underlying v4 SDK client
	v4Client1 := c1.client

	// Delete from converged cache
	cache.Delete(cp)

	// Verify the client is removed from converged cache
	_, _, err = cache.get("cluster1")
	assert.ErrorIs(t, err, types.ErrorClientNotFound)

	// Create a new client with the same parameters
	// This should create a new v4 SDK client since the old one was deleted
	c2, err := cache.GetOrCreate(cp)
	assert.NoError(t, err)
	assert.NotNil(t, c2)

	// The new client should be different from the original
	assert.NotEqual(t, c1, c2)

	// The underlying v4 SDK client should also be different (recreated)
	assert.NotEqual(t, v4Client1, c2.client)
}

func TestCacheInvalidationWithDifferentEndpoints(t *testing.T) {
	cache := NewClientCache()
	mgmtEndpoint := testhelpers.ManagementEndpointFromEnvironment(t)

	cp1 := &cachedClientParams{
		name:         "cluster1",
		mgmtEndpoint: *mgmtEndpoint,
	}

	cp2 := &cachedClientParams{
		name:         "cluster2",
		mgmtEndpoint: *mgmtEndpoint,
	}

	// Create clients for both endpoints
	c1, err := cache.GetOrCreate(cp1)
	assert.NoError(t, err)
	assert.NotNil(t, c1)

	c2, err := cache.GetOrCreate(cp2)
	assert.NoError(t, err)
	assert.NotNil(t, c2)

	// Store references to the underlying v4 SDK clients
	v4Client1 := c1.client
	v4Client2 := c2.client

	// Verify both clients exist in the converged cache
	_, _, err = cache.get("cluster1")
	assert.NoError(t, err)
	_, _, err = cache.get("cluster2")
	assert.NoError(t, err)

	// Delete only cluster1
	cache.Delete(cp1)

	// Verify cluster1 is removed from converged cache
	_, _, err = cache.get("cluster1")
	assert.ErrorIs(t, err, types.ErrorClientNotFound)

	// Verify cluster2 still exists in converged cache
	_, _, err = cache.get("cluster2")
	assert.NoError(t, err)

	// Recreate cluster1 and verify it gets a new v4 SDK client
	c1New, err := cache.GetOrCreate(cp1)
	assert.NoError(t, err)
	assert.NotNil(t, c1New)

	// The new cluster1 client should be different from the original
	assert.NotEqual(t, c1, c1New)

	// The underlying v4 SDK client should also be different (recreated)
	assert.NotEqual(t, v4Client1, c1New.client)

	// Verify cluster2's v4 SDK client is unchanged
	c2Retrieved, err := cache.GetOrCreate(cp2)
	assert.NoError(t, err)
	assert.Equal(t, c2, c2Retrieved)
	assert.Equal(t, v4Client2, c2Retrieved.client)
}

func TestNewClient(t *testing.T) {
	t.Run("creates client with valid credentials", func(t *testing.T) {
		mgmtEndpoint := testhelpers.ManagementEndpointFromEnvironment(t)
		credentials := prismgoclient.Credentials{
			Endpoint: mgmtEndpoint.Address.Host,
			Port:     mgmtEndpoint.Address.Port(),
			URL:      mgmtEndpoint.Address.String(),
			Username: mgmtEndpoint.Username,
			Password: mgmtEndpoint.Password,
			Insecure: mgmtEndpoint.Insecure,
		}

		client, err := NewClient(credentials)

		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.NotNil(t, client.client)
		assert.NotNil(t, client.AntiAffinityPolicies)
		assert.NotNil(t, client.Clusters)
		assert.NotNil(t, client.Categories)
		assert.NotNil(t, client.Images)
		assert.NotNil(t, client.StorageContainers)
		assert.NotNil(t, client.Subnets)
		assert.NotNil(t, client.VMs)
		assert.NotNil(t, client.Tasks)
	})

	t.Run("returns error with invalid credentials", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			Endpoint: "", // Invalid endpoint
			Username: "testuser",
			Password: "testpass",
		}

		client, err := NewClient(credentials)

		assert.Error(t, err)
		assert.Nil(t, client)
	})

	t.Run("returns error with missing username for basic auth", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			Endpoint: "test.example.com",
			Username: "", // Missing username
			Password: "testpass",
		}

		client, err := NewClient(credentials)

		assert.Error(t, err)
		assert.Nil(t, client)
	})

	t.Run("returns error with missing password for basic auth", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			Endpoint: "test.example.com",
			Username: "testuser",
			Password: "", // Missing password
		}

		client, err := NewClient(credentials)

		assert.Error(t, err)
		assert.Nil(t, client)
	})

	t.Run("returns error with missing endpoint for basic auth", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			Endpoint: "", // Missing endpoint
			Username: "testuser",
			Password: "testpass",
		}

		client, err := NewClient(credentials)

		assert.Error(t, err)
		assert.Nil(t, client)
	})

	t.Run("returns error with missing endpoint for api key auth", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			Endpoint: "", // Missing endpoint
			APIKey:   "test-api-key",
		}

		client, err := NewClient(credentials)

		assert.Error(t, err)
		assert.Nil(t, client)
	})
}

func TestNewClientFromV4SDKClient(t *testing.T) {
	t.Run("creates client from valid v4 SDK client", func(t *testing.T) {
		mgmtEndpoint := testhelpers.ManagementEndpointFromEnvironment(t)
		credentials := prismgoclient.Credentials{
			Endpoint: mgmtEndpoint.Address.Host,
			Port:     mgmtEndpoint.Address.Port(),
			URL:      mgmtEndpoint.Address.String(),
			Username: mgmtEndpoint.Username,
			Password: mgmtEndpoint.Password,
			Insecure: mgmtEndpoint.Insecure,
		}

		v4sdkClient, err := v4prismGoClient.NewV4Client(credentials)
		require.NoError(t, err)

		client := NewClientFromV4SDKClient(v4sdkClient)

		assert.NotNil(t, client)
		assert.NotNil(t, client.client)
		assert.Equal(t, v4sdkClient, client.client)
		assert.NotNil(t, client.AntiAffinityPolicies)
		assert.NotNil(t, client.Clusters)
		assert.NotNil(t, client.Categories)
		assert.NotNil(t, client.Images)
		assert.NotNil(t, client.StorageContainers)
		assert.NotNil(t, client.Subnets)
		assert.NotNil(t, client.VMs)
		assert.NotNil(t, client.Tasks)
	})

	t.Run("creates client with nil v4 SDK client", func(t *testing.T) {
		client := NewClientFromV4SDKClient(nil)

		assert.NotNil(t, client)
		assert.Nil(t, client.client)
		assert.NotNil(t, client.AntiAffinityPolicies)
		assert.NotNil(t, client.Clusters)
		assert.NotNil(t, client.Categories)
		assert.NotNil(t, client.Images)
		assert.NotNil(t, client.StorageContainers)
		assert.NotNil(t, client.Subnets)
		assert.NotNil(t, client.VMs)
		assert.NotNil(t, client.Tasks)
	})
}

func TestGetOrCreateErrorClientNotFound(t *testing.T) {
	cache := NewClientCache()
	mgmtEndpoint := testhelpers.ManagementEndpointFromEnvironment(t)
	cp := &cachedClientParams{
		name:         "nonexistent-cluster",
		mgmtEndpoint: *mgmtEndpoint,
	}

	// Test GetOrCreate when client is not found in cache
	// This should trigger the ErrorClientNotFound path and create a new client
	client, err := cache.GetOrCreate(cp)

	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.NotNil(t, client.client)
}

func TestGetNoValidationHash(t *testing.T) {
	cache := NewClientCache()
	client := &Client{}

	// Manually set a client in cache without validation hash
	// This simulates the scenario where client exists but no validation hash
	cache.set("test-client", "", client)

	// Test get() when no validation hash exists in validationHashes map
	returnedClient, hash, err := cache.get("test-client")

	assert.NoError(t, err)
	assert.Equal(t, "", hash) // Should return empty hash when not found in validationHashes
	assert.Equal(t, client, returnedClient)
}
