//go:build e2e

package e2e

import (
	"context"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	v4 "github.com/nutanix-cloud-native/prism-go-client/converged/v4"
	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	"github.com/nutanix-cloud-native/prism-go-client/tests/e2e/helpers"
	clusterModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

type cachedClientParams struct {
	name         string
	mgmtEndpoint types.ManagementEndpoint
}

func (c *cachedClientParams) Key() string                                  { return c.name }
func (c *cachedClientParams) ManagementEndpoint() types.ManagementEndpoint { return c.mgmtEndpoint }

func unexportedField(structPtr any, name string) reflect.Value {
	v := reflect.ValueOf(structPtr)
	if v.Kind() != reflect.Ptr || v.IsNil() || v.Elem().Kind() != reflect.Struct {
		return reflect.Value{}
	}
	f := v.Elem().FieldByName(name)
	if !f.IsValid() {
		return reflect.Value{}
	}
	return reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
}

func assertCacheBackref(t *testing.T, c *v4.Client) {
	t.Helper()
	cache := unexportedField(c, "cache")
	recreate := unexportedField(c, "recreate")
	require.True(t, cache.IsValid() && !cache.IsNil())
	require.True(t, recreate.IsValid() && !recreate.IsNil())
}

func newLiveCachedClient(t *testing.T, key string) (*v4.ClientCache, *cachedClientParams, *v4.Client) {
	t.Helper()
	endpoint := ManagementEndpointFromE2EEnvironment(t)
	cache := v4.NewClientCache()
	params := &cachedClientParams{name: key, mgmtEndpoint: *endpoint}
	client, err := cache.GetOrCreate(params)
	require.NoError(t, err)
	assertCacheBackref(t, client)
	_, err = client.Clusters.List(context.Background(), converged.WithLimit(1))
	require.NoError(t, err)
	return cache, params, client
}

func poisonAndAssertStale(t *testing.T, client *v4.Client) {
	t.Helper()
	helpers.StripSDKAuthorization(client)
	_, err := client.Clusters.List(context.Background(), converged.WithLimit(1))
	require.Error(t, err)
	require.Truef(t, converged.IsStaleClient(err), "got: %v", err)
}

func TestStaleClientRetryOnStaleE2E(t *testing.T) {
	cache, params, poisoned := newLiveCachedClient(t, "e2e-stale-retry-on-stale")
	poisonAndAssertStale(t, poisoned)
	calls := 0
	var recovered *v4.Client
	_, err := v4.RetryOnStale(poisoned, func(client *v4.Client) ([]clusterModels.Cluster, error) {
		calls++
		recovered = client
		if calls == 1 {
			helpers.StripSDKAuthorization(client)
		}
		return client.Clusters.List(context.Background(), converged.WithLimit(1))
	})
	require.NoError(t, err)
	assert.Equal(t, 2, calls)
	assert.NotSame(t, poisoned, recovered)
	assertCacheBackref(t, recovered)
	cached, err := cache.GetOrCreate(params)
	require.NoError(t, err)
	assert.Same(t, recovered, cached)
}

// TestStaleClientSelfHealsAutomaticallyE2E proves the no-caller-changes
// requirement: after poisoning the SDK client's auth state (simulating the
// "unsupported protocol scheme" stale-client bug), calling
// poisoned.Clusters.List directly -- with no RetryOnStale wrapper and no
// other caller code changes -- transparently self-heals and succeeds.
func TestStaleClientSelfHealsAutomaticallyE2E(t *testing.T) {
	cache, params, poisoned := newLiveCachedClient(t, "e2e-stale-self-heal")
	poisonAndAssertStale(t, poisoned)
	_, err := poisoned.Clusters.List(context.Background(), converged.WithLimit(1))
	require.NoError(t, err)
	// The same *Client pointer keeps working going forward too.
	_, err = poisoned.Clusters.List(context.Background(), converged.WithLimit(1))
	require.NoError(t, err)
	// The cache stays authoritative: SelfHeal goes through
	// Invalidate()+recreate() for cache-backed clients.
	cached, err := cache.GetOrCreate(params)
	require.NoError(t, err)
	_, err = cached.Clusters.List(context.Background(), converged.WithLimit(1))
	require.NoError(t, err)
}

func TestStaleClientRefreshRecoversE2E(t *testing.T) {
	cache, params, poisoned := newLiveCachedClient(t, "e2e-stale-refresh")
	poisonAndAssertStale(t, poisoned)
	fresh, err := poisoned.Refresh()
	require.NoError(t, err)
	assert.NotSame(t, poisoned, fresh)
	assertCacheBackref(t, fresh)
	_, err = fresh.Clusters.List(context.Background(), converged.WithLimit(1))
	require.NoError(t, err)
	cached, err := cache.GetOrCreate(params)
	require.NoError(t, err)
	assert.Same(t, fresh, cached)
}

func TestStaleClientInvalidateAllowsRecreateE2E(t *testing.T) {
	cache, params, poisoned := newLiveCachedClient(t, "e2e-stale-invalidate")
	poisonAndAssertStale(t, poisoned)
	poisoned.Invalidate()
	fresh, err := cache.GetOrCreate(params)
	require.NoError(t, err)
	assert.NotSame(t, poisoned, fresh)
	assertCacheBackref(t, fresh)
	_, err = fresh.Clusters.List(context.Background(), converged.WithLimit(1))
	require.NoError(t, err)
}

// TestStaleClientSelfHealsAutomaticallyStandaloneE2E proves the same
// no-caller-changes self-heal behavior for a Client built directly via
// v4.NewClient, without any ClientCache involved.
func TestStaleClientSelfHealsAutomaticallyStandaloneE2E(t *testing.T) {
	credentials := CredentialsFromE2EEnvironment(t)
	client, err := v4.NewClient(credentials)
	require.NoError(t, err)
	_, err = client.Clusters.List(context.Background(), converged.WithLimit(1))
	require.NoError(t, err)

	poisonAndAssertStale(t, client)

	_, err = client.Clusters.List(context.Background(), converged.WithLimit(1))
	require.NoError(t, err)
}
