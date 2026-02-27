// Package v4 provides integration tests for the converged client recovery plans service.
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run TestRecoveryPlansIntegration
//	go test -v ./converged/v4 -run TestRecoveryPlansService_ErrorHandling
//	go test -v ./converged/v4 -run TestRecoveryPlansService_UnsupportedODataOptions
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

// TestRecoveryPlansIntegration tests the client.DataPolicies.RecoveryPlans with real Nutanix API calls.
func TestRecoveryPlansIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ListRecoveryPlans", func(t *testing.T) {
		plans, err := client.DataPolicies.RecoveryPlans.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, plans)
		assert.GreaterOrEqual(t, len(plans), 0)
	})

	t.Run("ListRecoveryPlansWithValidOptions", func(t *testing.T) {
		plans, err := client.DataPolicies.RecoveryPlans.List(ctx,
			converged.WithPage(0),
			converged.WithLimit(5),
			converged.WithSelect("extId,name"),
			converged.WithOrderBy("name asc"),
		)
		assert.NoError(t, err)
		assert.NotNil(t, plans)
	})

	t.Run("GetRecoveryPlan", func(t *testing.T) {
		plans, err := client.DataPolicies.RecoveryPlans.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(plans) == 0 {
			t.Skip("No recovery plans available for testing")
		}

		uuid := *plans[0].ExtId
		require.NotEmpty(t, uuid)

		plan, err := client.DataPolicies.RecoveryPlans.Get(ctx, uuid)
		assert.NoError(t, err)
		assert.NotNil(t, plan)
		assert.Equal(t, uuid, *plan.ExtId)
	})

	t.Run("NewIterator", func(t *testing.T) {
		iter := client.DataPolicies.RecoveryPlans.NewIterator(ctx, converged.WithLimit(5))
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

// TestRecoveryPlansService_ErrorHandling tests error handling for nil client.
func TestRecoveryPlansService_ErrorHandling(t *testing.T) {
	service := NewRecoveryPlansService(nil)
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
		plan := &dpModels.RecoveryPlan{}
		_, err := service.Create(ctx, plan)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("CreateAsync_NilClient", func(t *testing.T) {
		plan := &dpModels.RecoveryPlan{}
		_, err := service.CreateAsync(ctx, plan)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("Update_NilClient", func(t *testing.T) {
		plan := &dpModels.RecoveryPlan{}
		_, err := service.Update(ctx, "test-id", plan)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("UpdateAsync_NilClient", func(t *testing.T) {
		plan := &dpModels.RecoveryPlan{}
		_, err := service.UpdateAsync(ctx, "test-id", plan)
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

// TestRecoveryPlansService_UnsupportedODataOptions tests that apply and expand options are rejected.
func TestRecoveryPlansService_UnsupportedODataOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("List_WithExpand", func(t *testing.T) {
		_, err := client.DataPolicies.RecoveryPlans.List(ctx, converged.WithExpand("config"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing")
		assert.Contains(t, err.Error(), "recovery plan")
	})

	t.Run("List_WithApply", func(t *testing.T) {
		_, err := client.DataPolicies.RecoveryPlans.List(ctx, converged.WithApply("groupby((extId))"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing")
		assert.Contains(t, err.Error(), "recovery plan")
	})

	t.Run("List_WithExpandAndLimit", func(t *testing.T) {
		_, err := client.DataPolicies.RecoveryPlans.List(ctx, converged.WithExpand("config"), converged.WithLimit(5))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing")
	})

	t.Run("NewIterator_WithExpand", func(t *testing.T) {
		iter := client.DataPolicies.RecoveryPlans.NewIterator(ctx, converged.WithExpand("config"))
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
		iter := client.DataPolicies.RecoveryPlans.NewIterator(ctx, converged.WithApply("groupby((extId))"))
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
