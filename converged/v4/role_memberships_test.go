package v4

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"

	iamAuthzModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
)

func TestRoleMembershipsService_ErrorHandling(t *testing.T) {
	service := NewRoleMembershipsService(nil)
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

	t.Run("Create_NilClient", func(t *testing.T) {
		roleMembership := iamAuthzModels.NewRoleMembership()
		_, err := service.Create(ctx, roleMembership)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("Delete_NilClient", func(t *testing.T) {
		err := service.Delete(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("NewIterator_NilClient", func(t *testing.T) {
		iterator := service.NewIterator(ctx)
		assert.Nil(t, iterator)
	})
}

// TestRoleMembershipsIntegration tests client.RoleMemberships with real Nutanix API calls.
func TestRoleMembershipsIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ListRoleMemberships", func(t *testing.T) {
		memberships, err := client.RoleMemberships.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, memberships)
		assert.GreaterOrEqual(t, len(memberships), 0)
	})

	t.Run("GetRoleMembership", func(t *testing.T) {
		memberships, err := client.RoleMemberships.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)
		if len(memberships) == 0 || memberships[0].ExtId == nil {
			t.Skip("No role memberships available for testing")
		}

		membershipUUID := *memberships[0].ExtId
		require.NotEmpty(t, membershipUUID)

		membership, err := client.RoleMemberships.Get(ctx, membershipUUID)
		assert.NoError(t, err)
		assert.NotNil(t, membership)
		assert.Equal(t, membershipUUID, *membership.ExtId)
	})

	t.Run("NewIterator", func(t *testing.T) {
		iterator := client.RoleMemberships.NewIterator(ctx, converged.WithLimit(1))
		require.NotNil(t, iterator)

		count := 0
		for membership, err := range iterator {
			require.NoError(t, err)
			count++
			assert.NotNil(t, membership.ExtId)
		}
		assert.GreaterOrEqual(t, count, 0)
	})
}
