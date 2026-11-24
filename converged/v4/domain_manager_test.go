// Package v4 provides integration tests for the converged client domain manager service.
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run TestDomainManagerIntegration
package v4

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	managementModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/management"
)

// TestDomainManagerIntegration tests the client.DomainManager with real Nutanix API calls
func TestDomainManagerIntegration(t *testing.T) {
	// Get credentials from environment
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	// Create converged client
	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("GetDomainManager", func(t *testing.T) {
		// First, list domain managers to get a valid external ID
		domainManagers, err := client.DomainManager.List(ctx)
		require.NoError(t, err)

		if len(domainManagers) == 0 {
			t.Skip("No domain managers available for testing")
		}

		// Get the first domain manager's external ID
		extID := *domainManagers[0].ExtId
		require.NotEmpty(t, extID)

		// Test Get domain manager
		dm, err := client.DomainManager.Get(ctx, extID)
		assert.NoError(t, err)
		assert.NotNil(t, dm)
		assert.Equal(t, extID, *dm.ExtId)
	})

	t.Run("ListDomainManagers", func(t *testing.T) {
		// Test basic list
		domainManagers, err := client.DomainManager.List(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, domainManagers)
		assert.GreaterOrEqual(t, len(domainManagers), 0)

		// Verify domain managers have required fields
		if len(domainManagers) > 0 {
			assert.NotNil(t, domainManagers[0].ExtId)
		}
	})

	t.Run("ListDomainManagersWithSelect", func(t *testing.T) {
		// Test list with select parameter (only supported OData option)
		domainManagers, err := client.DomainManager.List(ctx, converged.WithSelect("extId,config"))
		assert.NoError(t, err)
		assert.NotNil(t, domainManagers)

		if len(domainManagers) > 0 {
			assert.NotNil(t, domainManagers[0].ExtId)
		}
	})

	t.Run("ListDomainManagersWithUnsupportedOptions", func(t *testing.T) {
		// Test that unsupported OData options return errors
		testCases := []struct {
			name string
			opts []converged.ODataOption
		}{
			{
				name: "WithPage",
				opts: []converged.ODataOption{converged.WithPage(0)},
			},
			{
				name: "WithLimit",
				opts: []converged.ODataOption{converged.WithLimit(10)},
			},
			{
				name: "WithFilter",
				opts: []converged.ODataOption{converged.WithFilter("extId eq 'test'")},
			},
			{
				name: "WithOrderBy",
				opts: []converged.ODataOption{converged.WithOrderBy("extId")},
			},
			{
				name: "WithExpand",
				opts: []converged.ODataOption{converged.WithExpand("config")},
			},
			{
				name: "WithApply",
				opts: []converged.ODataOption{converged.WithApply("groupby((extId))")},
			},
			{
				name: "MultipleUnsupported",
				opts: []converged.ODataOption{
					converged.WithPage(0),
					converged.WithLimit(10),
				},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				_, err := client.DomainManager.List(ctx, tc.opts...)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "domain manager API only supports select parameter")
			})
		}
	})

	t.Run("NewIterator", func(t *testing.T) {
		// Test iterator creation
		iterator := client.DomainManager.NewIterator(ctx)
		assert.NotNil(t, iterator)

		// Test iterator with select option
		iteratorWithSelect := client.DomainManager.NewIterator(ctx, converged.WithSelect("extId"))
		assert.NotNil(t, iteratorWithSelect)

		// Iterate and collect results
		var count int
		for dm, err := range iterator {
			if err != nil {
				t.Logf("Iterator error: %v", err)
				break
			}
			if dm.ExtId != nil {
				count++
			}
		}
		assert.GreaterOrEqual(t, count, 0)
	})

	t.Run("Products", func(t *testing.T) {
		domainManagers, err := client.DomainManager.List(ctx)
		require.NoError(t, err)
		if len(domainManagers) == 0 {
			t.Skip("No domain managers available for testing")
		}
		dmExtID := *domainManagers[0].ExtId

		// Product methods are v4-SDK specific, exposed on the concrete service.
		dmService := NewDomainManagerService(client.client)

		// List all products on the domain manager.
		products, err := dmService.ListProducts(ctx, dmExtID)
		require.NoError(t, err)

		// Filtered list by product name, then Get by extId and inspect state.
		disabled, found, err := dmService.IsProductDisabled(
			ctx, dmExtID, managementModels.PRODUCTNAME_NUTANIX_DISASTER_RECOVERY)
		require.NoError(t, err)
		t.Logf("NUTANIX_DISASTER_RECOVERY found=%v disabled=%v (total products=%d)",
			found, disabled, len(products))

		if found {
			// Cross-check: GetProduct for each listed product returns an entity.
			for _, p := range products {
				if p.ExtId == nil {
					continue
				}
				got, gerr := dmService.GetProduct(ctx, dmExtID, *p.ExtId)
				assert.NoError(t, gerr)
				assert.NotNil(t, got)
			}
		}
	})

	t.Run("GetPrismCentralVersion", func(t *testing.T) {
		// Test GetPrismCentralVersion
		version, err := client.DomainManager.GetPrismCentralVersion(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, version)
		assert.NotEqual(t, "", version)
	})

	t.Run("GetPrismCentralVersionWithSelect", func(t *testing.T) {
		// Test GetPrismCentralVersion still works even if we can't use select
		// (since GetPrismCentralVersion calls List internally without options)
		version, err := client.DomainManager.GetPrismCentralVersion(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, version)
	})
}

// TestDomainManagerService_ErrorHandling tests error handling for various error scenarios
func TestDomainManagerService_ErrorHandling(t *testing.T) {
	service := NewDomainManagerService(nil)
	require.NotNil(t, service)

	ctx := context.Background()

	t.Run("Get_NilClient", func(t *testing.T) {
		_, err := service.Get(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("List_NilClient", func(t *testing.T) {
		_, err := service.List(ctx)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("NewIterator_NilClient", func(t *testing.T) {
		iterator := service.NewIterator(ctx)
		assert.Nil(t, iterator)
	})

	t.Run("GetPrismCentralVersion_NilClient", func(t *testing.T) {
		_, err := service.GetPrismCentralVersion(ctx)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("ListProducts_NilClient", func(t *testing.T) {
		_, err := service.ListProducts(ctx, "dm-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("GetProduct_NilClient", func(t *testing.T) {
		_, err := service.GetProduct(ctx, "dm-id", "product-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("IsProductDisabled_NilClient", func(t *testing.T) {
		_, _, err := service.IsProductDisabled(ctx, "dm-id", managementModels.PRODUCTNAME_NUTANIX_DISASTER_RECOVERY)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})
}

// TestProductNameFilterString documents the exact $filter string used by
// IsProductDisabled to look up a product by name, matching the API contract:
// $filter=name eq Prism.Management.ProductName'NUTANIX_DISASTER_RECOVERY'
func TestProductNameFilterString(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "NUTANIX_DISASTER_RECOVERY",
		managementModels.PRODUCTNAME_NUTANIX_DISASTER_RECOVERY.GetName())

	filter := fmt.Sprintf("name eq Prism.Management.ProductName'%s'",
		managementModels.PRODUCTNAME_NUTANIX_DISASTER_RECOVERY.GetName())
	assert.Equal(t, "name eq Prism.Management.ProductName'NUTANIX_DISASTER_RECOVERY'", filter)
}
