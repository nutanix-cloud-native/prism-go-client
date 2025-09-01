package cache

import (
	"context"

	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	"github.com/nutanix-cloud-native/prism-go-client/facade"
	facadeV4 "github.com/nutanix-cloud-native/prism-go-client/facade/v4"
	v3 "github.com/nutanix-cloud-native/prism-go-client/v3"
	v4 "github.com/nutanix-cloud-native/prism-go-client/v4"
)

type ClientCacheRegistry struct{}

func NewClientCacheRegistry() ClientCacheFactory {
	return &ClientCacheRegistry{}
}

type ClientCacheStruct[T any] interface {
	GetOrCreate(cachedClientParams types.CachedClientParams, opts ...types.ClientOption[T]) (*T, error)
	Delete(cachedClientParams types.CachedClientParams)
}

type ClientCacheIfc[T any, U any] interface {
	GetOrCreate(ctx context.Context, cachedClientParams types.CachedClientParams, opts ...types.ClientOption[T]) (U, error)
	Delete(cachedClientParams types.CachedClientParams)
}

type ClientCacheFactory interface {
	NewClientCacheV3(opts ...types.CacheOpts[v3.ClientCache]) ClientCacheStruct[v3.Client]
	NewClientCacheV4(opts ...types.CacheOpts[v4.ClientCache]) ClientCacheStruct[v4.Client]
	NewFacadeV4ClientCache(opts ...types.CacheOpts[v4.ClientCache]) ClientCacheIfc[v4.Client, facade.FacadeClientV4]
}

func (r *ClientCacheRegistry) NewClientCacheV3(opts ...types.CacheOpts[v3.ClientCache]) ClientCacheStruct[v3.Client] {
	return v3.NewClientCache(opts...)
}

func (r *ClientCacheRegistry) NewClientCacheV4(opts ...types.CacheOpts[v4.ClientCache]) ClientCacheStruct[v4.Client] {
	return v4.NewClientCache(opts...)
}

func (r *ClientCacheRegistry) NewFacadeV4ClientCache(opts ...types.CacheOpts[v4.ClientCache]) ClientCacheIfc[v4.Client, facade.FacadeClientV4] {
	return facadeV4.NewFacadeV4ClientCache(opts...)
}
