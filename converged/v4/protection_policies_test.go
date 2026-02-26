// Package v4 provides integration tests for the converged client protection policies service.
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run TestProtectionPoliciesIntegration
//	go test -v ./converged/v4 -run TestProtectionPoliciesService_ErrorHandling
//	go test -v ./converged/v4 -run TestProtectionPoliciesService_UnsupportedODataOptions
package v4

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"

	dpModels "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
)

// TestProtectionPoliciesIntegration tests the client.DataPolicies.ProtectionPolicies with real Nutanix API calls.
func TestProtectionPoliciesIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ListProtectionPolicies", func(t *testing.T) {
		policies, err := client.DataPolicies.ProtectionPolicies.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, policies)
		assert.GreaterOrEqual(t, len(policies), 0)
	})

	t.Run("ListProtectionPoliciesWithValidOptions", func(t *testing.T) {
		policies, err := client.DataPolicies.ProtectionPolicies.List(ctx,
			converged.WithPage(0),
			converged.WithLimit(5),
			converged.WithSelect("extId,name"),
			converged.WithOrderBy("name asc"),
		)
		assert.NoError(t, err)
		assert.NotNil(t, policies)
	})

	t.Run("GetProtectionPolicy", func(t *testing.T) {
		policies, err := client.DataPolicies.ProtectionPolicies.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(policies) == 0 {
			t.Skip("No protection policies available for testing")
		}

		uuid := *policies[0].ExtId
		require.NotEmpty(t, uuid)

		policy, err := client.DataPolicies.ProtectionPolicies.Get(ctx, uuid)
		assert.NoError(t, err)
		assert.NotNil(t, policy)
		assert.Equal(t, uuid, *policy.ExtId)
	})

	t.Run("NewIterator", func(t *testing.T) {
		iter := client.DataPolicies.ProtectionPolicies.NewIterator(ctx, converged.WithLimit(5))
		require.NotNil(t, iter)

		var count int
		for _, err := range iter {
			if err != nil {
				t.Logf("Iterator error: %v", err)
				break
			}
			count++
		}
		assert.GreaterOrEqual(t, count, 0)
	})
}

// TestProtectionPoliciesService_ErrorHandling tests error handling for nil client.
func TestProtectionPoliciesService_ErrorHandling(t *testing.T) {
	service := NewProtectionPoliciesService(nil)
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
		iter := service.NewIterator(ctx)
		assert.Nil(t, iter)
	})

	t.Run("Create_NilClient", func(t *testing.T) {
		policy := &dpModels.ProtectionPolicy{}
		_, err := service.Create(ctx, policy)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("CreateAsync_NilClient", func(t *testing.T) {
		policy := &dpModels.ProtectionPolicy{}
		_, err := service.CreateAsync(ctx, policy)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("Update_NilClient", func(t *testing.T) {
		policy := &dpModels.ProtectionPolicy{}
		_, err := service.Update(ctx, "test-id", policy)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("UpdateAsync_NilClient", func(t *testing.T) {
		policy := &dpModels.ProtectionPolicy{}
		_, err := service.UpdateAsync(ctx, "test-id", policy)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("Delete_NilClient", func(t *testing.T) {
		err := service.Delete(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("DeleteAsync_NilClient", func(t *testing.T) {
		_, err := service.DeleteAsync(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})
}

// TestProtectionPoliciesService_UnsupportedODataOptions tests that apply and expand options are rejected.
func TestProtectionPoliciesService_UnsupportedODataOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("List_WithExpand", func(t *testing.T) {
		_, err := client.DataPolicies.ProtectionPolicies.List(ctx, converged.WithExpand("config"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing")
		assert.Contains(t, err.Error(), "protection policy")
	})

	t.Run("List_WithApply", func(t *testing.T) {
		_, err := client.DataPolicies.ProtectionPolicies.List(ctx, converged.WithApply("groupby((extId))"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing")
		assert.Contains(t, err.Error(), "protection policy")
	})

	t.Run("List_WithExpandAndLimit", func(t *testing.T) {
		_, err := client.DataPolicies.ProtectionPolicies.List(ctx, converged.WithExpand("config"), converged.WithLimit(5))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing")
	})

	t.Run("NewIterator_WithExpand", func(t *testing.T) {
		iter := client.DataPolicies.ProtectionPolicies.NewIterator(ctx, converged.WithExpand("config"))
		require.NotNil(t, iter)

		var gotError error
		for _, err := range iter {
			gotError = err
			break
		}
		require.Error(t, gotError)
		assert.Contains(t, gotError.Error(), "apply and expand options are not supported for listing")
	})

	t.Run("NewIterator_WithApply", func(t *testing.T) {
		iter := client.DataPolicies.ProtectionPolicies.NewIterator(ctx, converged.WithApply("groupby((extId))"))
		require.NotNil(t, iter)

		var gotError error
		for _, err := range iter {
			gotError = err
			break
		}
		require.Error(t, gotError)
		assert.Contains(t, gotError.Error(), "apply and expand options are not supported for listing")
	})
}
