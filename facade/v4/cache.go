package v4

import (
	"errors"
	"fmt"
	"sync"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	"github.com/nutanix-cloud-native/prism-go-client/facade"
	v4 "github.com/nutanix-cloud-native/prism-go-client/v4"
)

type clientCacheMap map[string]*FacadeV4Client

type FacadeV4ClientCache struct {
	cache            map[string]*FacadeV4Client
	validationHashes map[string]string
	mtx              sync.RWMutex

	useSessionAuth bool
}

func NewFacadeV4ClientCache(opts ...types.CacheOpts[v4.ClientCache]) *FacadeV4ClientCache {
	cache := &FacadeV4ClientCache{
		cache:            make(clientCacheMap),
		validationHashes: make(map[string]string),
		mtx:              sync.RWMutex{},
	}

	return cache
}

func (c *FacadeV4ClientCache) GetOrCreate(cachedClientParams types.CachedClientParams, opts ...types.ClientOption[v4.Client]) (facade.FacadeClientV4, error) {
	client, validationHash, err := c.get(cachedClientParams.Key())
	if err != nil {
		if !errors.Is(err, types.ErrorClientNotFound) {
			return nil, fmt.Errorf("failed to get client with key %s from cache: %w", cachedClientParams.Key(), err)
		}
	}

	isValidationHashEqual, currentValidationHash, err := types.IsManagementEndpointHashEqual(cachedClientParams.ManagementEndpoint(), validationHash)
	if err != nil {
		return nil, fmt.Errorf("failed to compare validation hash for cachedClientParams with key %s: %w", cachedClientParams.Key(), err)
	}

	if isValidationHashEqual {
		// validation hash is the same, return the client
		return client, nil
	}

	c.Delete(cachedClientParams)

	credentials := prismgoclient.Credentials{
		URL:         cachedClientParams.ManagementEndpoint().Address.Host,
		Endpoint:    cachedClientParams.ManagementEndpoint().Address.Host,
		Insecure:    cachedClientParams.ManagementEndpoint().Insecure,
		Username:    cachedClientParams.ManagementEndpoint().ApiCredentials.Username,
		Password:    cachedClientParams.ManagementEndpoint().ApiCredentials.Password,
		SessionAuth: c.useSessionAuth,
	}

	// TODO(sid): v4 SDK doesn't have trust bundle as an input. Until we have a better solution, we will
	// set Insecure to true if trust bundle is provided to avoid breaking existing consumers of v3 SDK.
	if cachedClientParams.ManagementEndpoint().AdditionalTrustBundle != "" {
		credentials.Insecure = true
	}

	setDefaultsForCredentials(&credentials)
	if err := validateCredentials(credentials); err != nil {
		return nil, fmt.Errorf("failed to validate credentials for cachedClientParams with key %s: %w", cachedClientParams.Key(), err)
	}

	client, err = NewFacadeV4Client(credentials, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create client for cachedClientParams with key %s: %w", cachedClientParams.Key(), err)
	}

	c.set(cachedClientParams.Key(), currentValidationHash, client)
	return client, nil
}

func setDefaultsForCredentials(credentials *prismgoclient.Credentials) {
	if credentials.Port == "" {
		credentials.Port = "9440"
	}

	if credentials.URL == "" {
		credentials.URL = fmt.Sprintf("%s:%s", credentials.Endpoint, credentials.Port)
	}
}

func validateCredentials(credentials prismgoclient.Credentials) error {
	if credentials.Username == "" {
		return types.ErrorPrismUsernameNotSet
	}

	if credentials.Password == "" {
		return types.ErrorPrismPasswordNotSet
	}

	return nil
}

func (c *FacadeV4ClientCache) get(clientName string) (*FacadeV4Client, string, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	clnt, ok := c.cache[clientName]
	if !ok {
		return nil, "", types.ErrorClientNotFound
	}

	validationHash, ok := c.validationHashes[clientName]
	if !ok {
		return nil, "", types.ErrorClientNotFound
	}

	return clnt, validationHash, nil
}

func (c *FacadeV4ClientCache) set(clientName string, validationHash string, client *FacadeV4Client) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.cache[clientName] = client
	c.validationHashes[clientName] = validationHash
}

func (c *FacadeV4ClientCache) Delete(cachedClientParams types.CachedClientParams) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	delete(c.cache, cachedClientParams.Key())
	delete(c.validationHashes, cachedClientParams.Key())
}
