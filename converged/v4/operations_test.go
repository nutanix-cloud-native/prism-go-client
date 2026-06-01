package v4

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
)

func TestOperationsService_ErrorHandling(t *testing.T) {
	service := NewOperationsService(nil)
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
}

// TestOperationsIntegration tests client.Operations with real Nutanix API calls.
func TestOperationsIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)
	require.NotNil(t, client.Operations)

	ctx := context.Background()

	t.Run("ListOperations", func(t *testing.T) {
		operations, err := client.Operations.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, operations)
		assert.GreaterOrEqual(t, len(operations), 0)
	})

	t.Run("GetOperation", func(t *testing.T) {
		operations, err := client.Operations.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)
		if len(operations) == 0 {
			t.Skip("No operations available for testing")
		}

		operationUUID := *operations[0].ExtId
		require.NotEmpty(t, operationUUID)

		operation, err := client.Operations.Get(ctx, operationUUID)
		assert.NoError(t, err)
		assert.NotNil(t, operation)
		assert.Equal(t, operationUUID, *operation.ExtId)
	})
}

func TestOperationsListOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)
	require.NotNil(t, client.Operations)

	ctx := context.Background()

	t.Run("WithLimit", func(t *testing.T) {
		operations, err := client.Operations.List(ctx, converged.WithLimit(2))
		assert.NoError(t, err)
		assert.NotNil(t, operations)
		assert.LessOrEqual(t, len(operations), 2)
	})
	t.Run("WithFilter", func(t *testing.T) {
		filter := "displayName eq 'View_Role'"
		operations, err := client.Operations.List(ctx, converged.WithFilter(filter))
		assert.NoError(t, err)
		assert.NotNil(t, operations)
		assert.Equal(t, 1, len(operations))
		assert.Equal(t, "View_Role", *operations[0].DisplayName)
	})
	
	t.Run("WithSelect", func(t *testing.T) {
		operations, err := client.Operations.List(ctx, converged.WithSelect("extId,displayName"))
		assert.NoError(t, err)
		assert.NotNil(t, operations)
	})

	t.Run("WithApply_NotSupported", func(t *testing.T) {
		_, err := client.Operations.List(ctx, converged.WithApply("groupby((displayName))"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing Operations")
	})
}

