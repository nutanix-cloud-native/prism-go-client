package v3

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
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

// Test GetOrCreate defensive programming - nil Address pointer validation
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

// Test GetOrCreate defensive programming - empty Address.Host validation
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

// Test GetOrCreate defensive programming - empty Username validation
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
	assert.Contains(t, err.Error(), "API credentials username is empty for cachedClientParams with key test-client")
}

// Test GetOrCreate defensive programming - empty Password validation
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
	assert.Contains(t, err.Error(), "API credentials password is empty for cachedClientParams with key test-client")
}

// Test GetOrCreate defensive programming - all validations pass
func TestGetOrCreateValidationsPassed(t *testing.T) {
	cache := NewClientCache()
	params := &cachedClientParams{
		name: "test-client",
		mgmtEndpoint: types.ManagementEndpoint{
			ApiCredentials: types.ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address:  &url.URL{Host: "test.com"},
			Insecure: true,
		},
	}

	// This will pass all defensive validations
	client, err := cache.GetOrCreate(params)

	// We don't assert on the error/client result since it depends on whether
	// the credentials work or not. The important thing is that we don't get
	// any of our defensive programming validation errors
	if err != nil {
		// If there is an error, ensure it's NOT from our defensive programming validations
		assert.NotContains(t, err.Error(), "management endpoint address is nil")
		assert.NotContains(t, err.Error(), "address host is empty")
		assert.NotContains(t, err.Error(), "username is empty")
		assert.NotContains(t, err.Error(), "password is empty")
	}

	// The test passes as long as we didn't get validation errors
	// Whether client creation succeeds or fails depends on the actual credentials
	_ = client // Suppress unused variable warning if client is returned
}
