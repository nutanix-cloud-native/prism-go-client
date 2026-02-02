// Package v4 provides integration tests for the converged client OVAs service.
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run TestOvasIntegration
//	go test -v ./converged/v4 -run TestOvasService_ErrorHandling
package v4

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"

	ovaModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// TestOvasService_ErrorHandling tests error handling for nil client
func TestOvasService_ErrorHandling(t *testing.T) {
	service := NewOvasService(nil)
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

	t.Run("Create_NilClient", func(t *testing.T) {
		ova := ovaModels.NewOva()
		_, err := service.Create(ctx, ova)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("CreateAsync_NilClient", func(t *testing.T) {
		ova := ovaModels.NewOva()
		_, err := service.CreateAsync(ctx, ova)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("GetFile_NilClient", func(t *testing.T) {
		_, err := service.GetFile(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})
}

// TestOvasIntegration tests the client.Ovas with real Nutanix API calls
func TestOvasIntegration(t *testing.T) {
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

	t.Run("ListOvas", func(t *testing.T) {
		// Test List OVAs
		ovas, err := client.Ovas.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, ovas)
		assert.GreaterOrEqual(t, len(ovas), 0)
	})

	t.Run("GetOva", func(t *testing.T) {
		// First, list OVAs to get a valid OVA UUID
		ovas, err := client.Ovas.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(ovas) == 0 {
			t.Skip("No OVAs available for testing")
		}

		// Get the first OVA's UUID
		ovaUUID := *ovas[0].ExtId
		require.NotEmpty(t, ovaUUID)

		// Test Get OVA
		ova, err := client.Ovas.Get(ctx, ovaUUID)
		assert.NoError(t, err)
		assert.NotNil(t, ova)
		assert.Equal(t, ovaUUID, *ova.ExtId)
	})

	t.Run("NewIterator", func(t *testing.T) {
		// Test iterator
		iterator := client.Ovas.NewIterator(ctx)
		require.NotNil(t, iterator)

		// Collect OVAs using iterator
		var ovas []ovaModels.Ova
		for ova, err := range iterator {
			if err != nil {
				break
			}
			ovas = append(ovas, ova)
			if len(ovas) >= 5 { // Limit to prevent long test runs
				break
			}
		}

		assert.GreaterOrEqual(t, len(ovas), 0)
	})
}

// TestOvasListOptions tests various OData options for listing OVAs
func TestOvasListOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("WithLimit", func(t *testing.T) {
		ovas, err := client.Ovas.List(ctx, converged.WithLimit(2))
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(ovas), 2)
	})

	t.Run("WithPage", func(t *testing.T) {
		ovas, err := client.Ovas.List(ctx, converged.WithPage(0), converged.WithLimit(1))
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(ovas), 1)
	})

	t.Run("WithSelect", func(t *testing.T) {
		ovas, err := client.Ovas.List(ctx, converged.WithSelect("extId,name"))
		assert.NoError(t, err)
		assert.NotNil(t, ovas)
	})

	t.Run("WithOrderBy", func(t *testing.T) {
		ovas, err := client.Ovas.List(ctx, converged.WithOrderBy("name asc"))
		assert.NoError(t, err)
		assert.NotNil(t, ovas)
	})

	t.Run("MultipleOptions", func(t *testing.T) {
		ovas, err := client.Ovas.List(ctx,
			converged.WithLimit(5),
			converged.WithPage(0),
			converged.WithSelect("extId,name"),
		)
		assert.NoError(t, err)
		assert.NotNil(t, ovas)
		assert.LessOrEqual(t, len(ovas), 5)
	})

	t.Run("UnsupportedExpandOption", func(t *testing.T) {
		// Expand is not supported for OVAs
		_, err := client.Ovas.List(ctx, converged.WithExpand("someField"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "expand and apply are not supported")
	})

	t.Run("UnsupportedApplyOption", func(t *testing.T) {
		// Apply is not supported for OVAs
		_, err := client.Ovas.List(ctx, converged.WithApply("someAggregation"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "expand and apply are not supported")
	})
}

// TestOvasErrorScenarios tests error handling scenarios
func TestOvasErrorScenarios(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("GetInvalidUUID", func(t *testing.T) {
		_, err := client.Ovas.Get(ctx, "invalid-uuid-format")
		assert.Error(t, err)
	})

	t.Run("GetNonExistentUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.Ovas.Get(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("GetFileNonExistentUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.Ovas.GetFile(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("InvalidFilterSyntax", func(t *testing.T) {
		_, err := client.Ovas.List(ctx, converged.WithFilter("invalid filter syntax"))
		assert.Error(t, err)
	})
}
