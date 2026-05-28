package v3

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
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
	assert.Contains(t, err.Error(), "API credentials(either username and password or API key) is required for authentication")
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
	assert.Contains(t, err.Error(), "API credentials(either username and password or API key) is required for authentication")
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

	// Require that there is no error - all validations pass and client creation succeeds
	require.NoError(t, err)
	assert.NotNil(t, client)
}

// Test validateManagementEndpoint function directly
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
			name: "valid endpoint with api key",
			endpoint: types.ManagementEndpoint{
				ApiCredentials: types.ApiCredentials{
					APIKey: "test-api-key",
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
			key:     "key1",
			wantErr: "",
		},
		{
			name: "valid api key credentials",
			credentials: prismgoclient.Credentials{
				APIKey: "test-api-key",
			},
			key:     "key2",
			wantErr: "",
		},
		{
			name: "missing username for basic auth",
			credentials: prismgoclient.Credentials{
				Password: "testpass",
			},
			key:     "key3",
			wantErr: "API credentials(either username and password or API key) is required for authentication in cachedClientParams with key key3",
		},
		{
			name: "missing password for basic auth",
			credentials: prismgoclient.Credentials{
				Username: "testuser",
			},
			key:     "key4",
			wantErr: "API credentials(either username and password or API key) is required for authentication in cachedClientParams with key key4",
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
