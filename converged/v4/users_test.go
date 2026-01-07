// Package v4 provides integration tests for the converged client users service.
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run TestUsersIntegration
//	go test -v ./converged/v4 -run TestUsersService_ErrorHandling
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

// TestUsersIntegration tests the client.Users with real Nutanix API calls
func TestUsersIntegration(t *testing.T) {
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

	t.Run("ListUsers", func(t *testing.T) {
		// Test List users
		users, err := client.Users.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.GreaterOrEqual(t, len(users), 0)
	})

	t.Run("GetUser", func(t *testing.T) {
		// First, list users to get a valid user UUID
		users, err := client.Users.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(users) == 0 {
			t.Skip("No users available for testing")
		}

		// Get the first user's UUID
		userUUID := *users[0].ExtId
		require.NotEmpty(t, userUUID)

		// Test Get user
		user, err := client.Users.Get(ctx, userUUID)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userUUID, *user.ExtId)
	})
}

// TestUsersService_ErrorHandling tests error handling for nil client
func TestUsersService_ErrorHandling(t *testing.T) {
	service := NewUsersService(nil)
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
