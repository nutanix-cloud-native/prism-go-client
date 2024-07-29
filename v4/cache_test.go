package v4

import (
	"testing"

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
	assert.ErrorIs(t, err, ErrorClientNotFound)
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
	assert.ErrorIs(t, err, ErrorClientNotFound)
}

func TestDeleteDoesNotErrorIfClientNotPresentInCache(t *testing.T) {
	cache := NewClientCache()

	cache.Delete(&cachedClientParams{name: "cluster1"}) // No error expected
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
}
