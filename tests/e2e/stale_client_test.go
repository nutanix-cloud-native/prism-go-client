//go:build e2e

// E2E tests for stale-client cache recovery against a real Prism Central.
//
// Setup:
//
//	cp .env.e2e.example .env.e2e
//	# or: cp .env.e2e.example .env.e2e.pc-7.5
//
// Run:
//
//	make test-e2e
//	E2E_PROFILE=pc-7.5 make test-e2e
//
// Required env vars (see .env.e2e.example):
//
//	NUTANIX_E2E_ENDPOINT, NUTANIX_E2E_PORT, NUTANIX_E2E_USERNAME,
//	NUTANIX_E2E_PASSWORD, NUTANIX_E2E_INSECURE
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

func (c *cachedClientParams) Key() string { return c.name }

func (c *cachedClientParams) ManagementEndpoint() types.ManagementEndpoint {
	return c.mgmtEndpoint
}

func unexportedField(structPtr any, name string) reflect.Value {
	v := reflect.ValueOf(structPtr)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return reflect.Value{}
	}
	elem := v.Elem()
	if elem.Kind() != reflect.Struct {
		return reflect.Value{}
	}
	f := elem.FieldByName(name)
	if !f.IsValid() {
		return reflect.Value{}
	}
	return reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
}

func assertCacheBackref(t *testing.T, c *v4.Client) {
	t.Helper()
	cacheField := unexportedField(c, "cache")
	recreateField := unexportedField(c, "recreate")
	require.True(t, cacheField.IsValid() && !cacheField.IsNil(), "client.cache backref must be set")
	require.True(t, recreateField.IsValid() && !recreateField.IsNil(), "client.recreate backref must be set")
}

func newLiveCachedClient(t *testing.T, cacheKey string) (*v4.ClientCache, *cachedClientParams, *v4.Client) {
	t.Helper()
	mgmtEndpoint := ManagementEndpointFromE2EEnvironment(t)
	cache := v4.NewClientCache()
	cp := &cachedClientParams{
		name:         cacheKey,
		mgmtEndpoint: *mgmtEndpoint,
	}
	client, err := cache.GetOrCreate(cp)
	require.NoError(t, err)
	require.NotNil(t, client)
	assertCacheBackref(t, client)

	ctx := context.Background()
	_, err = client.Clusters.List(ctx, converged.WithLimit(1))
	require.NoError(t, err, "sanity Clusters.List must succeed with e2e credentials")
	return cache, cp, client
}

func poisonAndAssertStale(t *testing.T, client *v4.Client) {
	t.Helper()
	helpers.StripSDKAuthorization(client)
	_, err := client.Clusters.List(context.Background(), converged.WithLimit(1))
	require.Error(t, err, "expected error after stripping Authorization")
	require.Truef(t, converged.IsStaleClient(err),
		"expected ErrStaleClient (unsupported protocol scheme from relative OIDC redirect); got: %v", err)
}

// TestStaleClientRetryOnStaleE2E forces the unsupported-protocol-scheme failure mode by
// stripping Authorization, then verifies RetryOnStale refreshes the cached client and retries.
func TestStaleClientRetryOnStaleE2E(t *testing.T) {
	cache, cp, poisoned := newLiveCachedClient(t, "e2e-stale-retry-on-stale")
	poisonAndAssertStale(t, poisoned)

	ctx := context.Background()
	calls := 0
	var recovered *v4.Client
	_, err := v4.RetryOnStale(poisoned, func(c *v4.Client) ([]clusterModels.Cluster, error) {
		calls++
		recovered = c
		if calls == 1 {
			helpers.StripSDKAuthorization(c)
		}
		return c.Clusters.List(ctx, converged.WithLimit(1))
	})
	require.NoError(t, err)
	assert.Equal(t, 2, calls)
	require.NotNil(t, recovered)
	assert.NotSame(t, poisoned, recovered)
	assertCacheBackref(t, recovered)

	cached, err := cache.GetOrCreate(cp)
	require.NoError(t, err)
	assert.NotSame(t, poisoned, cached)
	assert.Same(t, recovered, cached)
}

// TestStaleClientWithoutRetryOnStaleRemainsPoisonedE2E is the NCN-116866 failure mode:
// a direct List after auth loss stays ErrStaleClient, and GetOrCreate keeps returning
// the poisoned client because nothing invalidated the cache.
func TestStaleClientWithoutRetryOnStaleRemainsPoisonedE2E(t *testing.T) {
	cache, cp, poisoned := newLiveCachedClient(t, "e2e-stale-no-retry")
	poisonAndAssertStale(t, poisoned)

	ctx := context.Background()

	// Without RetryOnStale, subsequent calls keep failing the same way.
	_, err := poisoned.Clusters.List(ctx, converged.WithLimit(1))
	require.Error(t, err)
	assert.True(t, converged.IsStaleClient(err), "got: %v", err)

	cached, err := cache.GetOrCreate(cp)
	require.NoError(t, err)
	assert.Same(t, poisoned, cached, "poisoned client must remain cached without RetryOnStale/Invalidate")

	_, err = cached.Clusters.List(ctx, converged.WithLimit(1))
	require.Error(t, err)
	assert.True(t, converged.IsStaleClient(err), "cached client must still be stale; got: %v", err)
}

// TestStaleClientRefreshRecoversE2E verifies Refresh alone recreates an authenticated client.
func TestStaleClientRefreshRecoversE2E(t *testing.T) {
	cache, cp, poisoned := newLiveCachedClient(t, "e2e-stale-refresh")
	poisonAndAssertStale(t, poisoned)

	fresh, err := poisoned.Refresh()
	require.NoError(t, err)
	require.NotNil(t, fresh)
	assert.NotSame(t, poisoned, fresh)
	assertCacheBackref(t, fresh)

	_, err = fresh.Clusters.List(context.Background(), converged.WithLimit(1))
	require.NoError(t, err)

	cached, err := cache.GetOrCreate(cp)
	require.NoError(t, err)
	assert.Same(t, fresh, cached)
	assert.NotSame(t, poisoned, cached)
}

// TestStaleClientInvalidateAllowsRecreateE2E verifies Invalidate drops the poisoned entry so
// a later GetOrCreate returns a working client (CCM restart-equivalent without RetryOnStale).
func TestStaleClientInvalidateAllowsRecreateE2E(t *testing.T) {
	cache, cp, poisoned := newLiveCachedClient(t, "e2e-stale-invalidate")
	poisonAndAssertStale(t, poisoned)

	poisoned.Invalidate()

	fresh, err := cache.GetOrCreate(cp)
	require.NoError(t, err)
	require.NotNil(t, fresh)
	assert.NotSame(t, poisoned, fresh)
	assertCacheBackref(t, fresh)

	_, err = fresh.Clusters.List(context.Background(), converged.WithLimit(1))
	require.NoError(t, err)
}
