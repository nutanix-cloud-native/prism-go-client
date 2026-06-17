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

// TestDisksIntegration tests the client.Disks with real Nutanix API calls.
func TestDisksIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ListDisks", func(t *testing.T) {
		disks, err := client.Disks.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(disks), 0)
	})

	t.Run("GetNonExistentDisk", func(t *testing.T) {
		// A well-formed but non-existent disk UUID should surface an error
		// rather than panic or return a disk.
		_, err := client.Disks.Get(ctx, "00000000-0000-0000-0000-000000000000")
		assert.Error(t, err)
	})
}

// TestDisksService_ErrorHandling tests error handling for nil client.
func TestDisksService_ErrorHandling(t *testing.T) {
	service := NewDisksService(nil)
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
