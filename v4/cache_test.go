package v4

import (
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
