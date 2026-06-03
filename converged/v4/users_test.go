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
	"fmt"
	"strings"
	"testing"
	"time"

	"k8s.io/utils/ptr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"

	iamModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
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

	t.Run("ListUserKeys", func(t *testing.T) {
		filter := "userType eq Iam.Authn.UserType'SERVICE_ACCOUNT'"
		users, err := client.Users.List(ctx, converged.WithLimit(1), converged.WithFilter(filter))
		require.NoError(t, err)
		if len(users) == 0 {
			t.Skip("No users available for testing user keys")
		}
		userUUID := *users[0].ExtId
		require.NotEmpty(t, userUUID)

		keys, err := client.Users.ListUserKeys(ctx, userUUID, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, keys)
		assert.GreaterOrEqual(t, len(keys), 0)
	})

	t.Run("GetUserKeyById", func(t *testing.T) {
		filter := "username eq 'prism-go-client-test-user'"
		users, err := client.Users.List(ctx, converged.WithLimit(1), converged.WithFilter(filter))
		require.NoError(t, err)
		userUUID := *users[0].ExtId
		require.NotEmpty(t, userUUID)

		keys, err := client.Users.ListUserKeys(ctx, userUUID, converged.WithLimit(1))
		require.NoError(t, err)
		if len(keys) == 0 || keys[0].ExtId == nil {
			t.Skip("No user keys available for testing")
		}
		keyID := *keys[0].ExtId
		require.NotEmpty(t, keyID)

		key, err := client.Users.GetUserKeyById(ctx, userUUID, "acf6c546-2e2a-56a0-83d5-c07733f9772f")
		assert.NoError(t, err)
		assert.NotNil(t, key)
		assert.Equal(t, keyID, *key.ExtId)
		fmt.Printf("test key: %+v\n", *key.ExtId)
	})

	// test for creating user with service account type
	t.Run("CreateUserWithServiceAccountType", func(t *testing.T) {
		user := iamModels.NewUser()
		user.Username = ptr.To("prism-go-client-test-user")
		user.UserType = ptr.To(iamModels.USERTYPE_SERVICE_ACCOUNT)
		createdUser, err := client.Users.Create(ctx, user)
		require.NoError(t, err)
		require.NotNil(t, createdUser)
		assert.Equal(t, user.Username, *createdUser.Username)
		assert.Equal(t, user.UserType, *createdUser.UserType)
		fmt.Printf("test user: %+v\n", createdUser.ExtId)
	})

	t.Run("CreateUserKey", func(t *testing.T) {
		key := iamModels.NewKey()
		userExtId := "ae7ed369-8642-53f2-ae76-a566de392c6c"
		key.Name = ptr.To("prism-go-client-test-key-1")
		key.KeyType = ptr.To(iamModels.KEYKIND_API_KEY)
		expiryTime := time.Now().Add(1 * time.Hour)
		key.ExpiryTime = &expiryTime
		key.Description = ptr.To("prism-go-client-test-key-description")

		createdKey, err := client.Users.CreateUserKey(ctx, userExtId, key)
		require.NoError(t, err)
		require.NotNil(t, createdKey)
		keyDetails, err := createdKey.KeyDetails.MarshalJSON()
		require.NoError(t, err)
		require.NotNil(t, keyDetails)

		fmt.Printf("test key details: %+v\n", string(keyDetails))
		fmt.Printf("test key expiry time: %+v\n", *createdKey.ExpiryTime)
		fmt.Printf("test key description: %+v\n", *createdKey.ExtId)
	})

	t.Run("DeleteUserKey", func(t *testing.T) {
		userExtId := "ae7ed369-8642-53f2-ae76-a566de392c6c"
		keyID := "fde67989-003a-59ab-8eb9-949e981621ce"
		err := client.Users.DeleteUserKeyById(ctx, userExtId, keyID)
		assert.NoError(t, err)
	})

	t.Run("RevokeUserKey", func(t *testing.T) {
		userExtId := "ae7ed369-8642-53f2-ae76-a566de392c6c"
		keyID := "fde67989-003a-59ab-8eb9-949e981621ce"
		err := client.Users.RevokeUserKey(ctx, userExtId, keyID)
		assert.NoError(t, err)
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

	t.Run("Create_ValidUser_NilClient", func(t *testing.T) {
		user := iamModels.NewUser()
		_, err := service.Create(ctx, user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("ListUserKeys_NilClient", func(t *testing.T) {
		_, err := service.ListUserKeys(ctx, "user-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("CreateUserKey_NilClient", func(t *testing.T) {
		key := iamModels.NewKey()
		_, err := service.CreateUserKey(ctx, "user-id", key)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("GetUserKeyById_NilClient", func(t *testing.T) {
		_, err := service.GetUserKeyById(ctx, "user-id", "key-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("DeleteUserKeyById_NilClient", func(t *testing.T) {
		err := service.DeleteUserKeyById(ctx, "user-id", "key-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("RevokeUserKey_NilClient", func(t *testing.T) {
		err := service.RevokeUserKey(ctx, "user-id", "key-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("NewIterator_NilClient", func(t *testing.T) {
		iterator := service.NewIterator(ctx)
		assert.Nil(t, iterator)
	})
}
