package v4

import (
	"errors"
	"fmt"
	"sync"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
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

	v4sdkClient, err := c.v4sdkClientCache.GetOrCreate(cachedClientParams, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to get or create v4sdk client for cachedClientParams with key %s: %w", cachedClientParams.Key(), err)
	}

	client = NewClientFromV4SDKClient(v4sdkClient)
	client.cache = c
	client.cacheKey = cachedClientParams.Key()
	client.recreate = func() (*Client, error) {
		return c.GetOrCreate(cachedClientParams, opts...)
	}

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
	c.deleteByKey(params.Key())
}

func (c *ClientCache) deleteByKey(key string) {
	c.mtx.Lock()
	delete(c.cache, key)
	delete(c.validationHashes, key)
	c.mtx.Unlock()

	c.v4sdkClientCache.Delete(&cacheKeyParams{key: key})
}

type cacheKeyParams struct {
	key string
}

func (p *cacheKeyParams) Key() string {
	return p.key
}

func (p *cacheKeyParams) ManagementEndpoint() types.ManagementEndpoint {
	return types.ManagementEndpoint{}
}

// Invalidate removes this client from the cache that created it.
func (c *Client) Invalidate() {
	if c == nil || c.cache == nil || c.cacheKey == "" {
		return
	}
	c.cache.deleteByKey(c.cacheKey)
	c.cache, c.cacheKey, c.recreate = nil, "", nil
}

// Refresh invalidates this cached client and returns a newly created one.
func (c *Client) Refresh() (*Client, error) {
	if c == nil || c.recreate == nil {
		return nil, fmt.Errorf("client was not created via ClientCache")
	}
	recreate := c.recreate
	c.Invalidate()
	return recreate()
}

// SelfHeal rebuilds this client in place after a stale-client error, without
// requiring any caller code changes: callers keep using the same *Client
// pointer, since only its pointee's contents change.
//
// If this client was created via a ClientCache, SelfHeal reuses Refresh()
// (which calls Invalidate() to remove the stale cache entry, then recreate()
// to register a freshly-built client with a new validation hash) so the
// cache stays authoritative.
//
// If this client was created directly (no ClientCache backref, e.g. via
// NewClient outside of any cache), SelfHeal rebuilds it from the
// credentials/options originally passed to NewClient.
//
// SelfHeal is intended to be called rarely (roughly once per auth-session
// lifetime), so it takes a full mutex for the entire rebuild-and-copy
// operation; this adds no overhead to the normal request hot path, since the
// mutex is only acquired while a heal is in flight.
func (c *Client) SelfHeal() error {
	if c == nil {
		return fmt.Errorf("client is nil")
	}

	c.healMu.Lock()
	defer c.healMu.Unlock()

	var fresh *Client
	var err error
	if c.recreate != nil {
		fresh, err = c.Refresh()
	} else if c.hasCredentials {
		fresh, err = NewClient(c.credentials, c.opts...)
	} else {
		return fmt.Errorf("client cannot self-heal: not created via ClientCache or NewClient")
	}
	if err != nil {
		return fmt.Errorf("failed to self-heal client: %w", err)
	}

	*c = *fresh
	// The services created for fresh point back at fresh, not at c. Re-point
	// them at c so future calls keep using the same converged Client that
	// callers already hold a pointer to.
	c.rebindServices()

	return nil
}

// RetryOnStale runs fn with client and retries once with a refreshed client
// when the stale-client transport error is returned.
func RetryOnStale[T any](client *Client, fn func(*Client) (T, error)) (T, error) {
	result, err := fn(client)
	if err == nil || !converged.IsStaleClient(err) {
		return result, err
	}

	fresh, err := client.Refresh()
	if err != nil {
		var zero T
		return zero, err
	}
	return fn(fresh)
}
