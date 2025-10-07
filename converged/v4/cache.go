package v4

import (
	"errors"
	"fmt"
	"sync"

	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
)

type ClientCache struct {
	cache            map[string]*Client
	validationHashes map[string]string
	mtx              sync.RWMutex

	v4sdkClientCache *v4prismGoClient.ClientCache
}

func NewClientCache(opts ...v4prismGoClient.CacheOpts) *ClientCache {
	v4sdkClientCache := v4prismGoClient.NewClientCache(opts...)

	return &ClientCache{
		cache:            make(map[string]*Client),
		validationHashes: make(map[string]string),
		mtx:              sync.RWMutex{},
		v4sdkClientCache: v4sdkClientCache,
	}
}

func (c *ClientCache) GetOrCreate(cachedClientParams types.CachedClientParams, opts ...types.ClientOption[v4prismGoClient.Client]) (*Client, error) {
	currentValidationHash, err := cachedClientParams.ManagementEndpoint().GetHash()
	if err != nil {
		return nil, fmt.Errorf("failed to calculate validation hash for cachedClientParams with key %s: %w", cachedClientParams.Key(), err)
	}

	client, validationHash, err := c.get(cachedClientParams.Key())
	if err != nil {
		if !errors.Is(err, types.ErrorClientNotFound) {
			return nil, fmt.Errorf("failed to get client with key %s from cache: %w", cachedClientParams.Key(), err)
		}
	}

	if validationHash == currentValidationHash {
		// validation hash is the same, return the client
		return client, nil
	}

	// validation hash is different, regenerate the client
	c.Delete(cachedClientParams)
	c.v4sdkClientCache.Delete(cachedClientParams)

	v4sdkClient, err := c.v4sdkClientCache.GetOrCreate(cachedClientParams, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to get or create v4sdk client for cachedClientParams with key %s: %w", cachedClientParams.Key(), err)
	}

	client = NewClientFromV4SDKClient(v4sdkClient)

	c.set(cachedClientParams.Key(), currentValidationHash, client)

	return client, nil
}

func (c *ClientCache) get(clientName string) (*Client, string, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	client, ok := c.cache[clientName]
	if !ok {
		return nil, "", types.ErrorClientNotFound
	}

	validationHash, ok := c.validationHashes[clientName]
	if !ok {
		return client, "", nil
	}

	return client, validationHash, nil
}

func (c *ClientCache) set(clientName string, validationHash string, client *Client) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.cache[clientName] = client
	c.validationHashes[clientName] = validationHash
}

func (c *ClientCache) Delete(params types.CachedClientParams) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	delete(c.cache, params.Key())
	delete(c.validationHashes, params.Key())

	c.v4sdkClientCache.Delete(params)
}
