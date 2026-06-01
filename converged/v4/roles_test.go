package v4

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"

	iamAuthzModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
)

func TestRolesService_ErrorHandling(t *testing.T) {
	service := NewRolesService(nil)
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
		role := iamAuthzModels.NewRole()
		_, err := service.Create(ctx, role)
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

// TestRolesIntegration tests the client.Roles with real Nutanix API calls.
func TestRolesIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ListRoles", func(t *testing.T) {
		roles, err := client.Roles.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, roles)
		assert.GreaterOrEqual(t, len(roles), 0)
	})

	t.Run("GetRole", func(t *testing.T) {
		roles, err := client.Roles.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)
		if len(roles) == 0 {
			t.Skip("No roles available for testing")
		}

		roleUUID := *roles[0].ExtId
		require.NotEmpty(t, roleUUID)

		role, err := client.Roles.Get(ctx, roleUUID)
		assert.NoError(t, err)
		assert.NotNil(t, role)
		assert.Equal(t, roleUUID, *role.ExtId)
	})
	t.Run("CreateRole", func(t *testing.T) {
		role := iamAuthzModels.NewRole()
		role.DisplayName = ptr.To("prism-go-client-test-role")
		role.Description = ptr.To("prism-go-client-test-role-description")
		role.Operations = []string{"625c5143-b08d-5631-a39d-f1a89c63f9df"} // View_Role operation

		createdRole, err := client.Roles.Create(ctx, role)
		assert.NoError(t, err)
		assert.NotNil(t, createdRole)
		fmt.Printf("createdRole: %+v\n", *createdRole.ExtId)
	})

	t.Run("DeleteRole", func(t *testing.T) {
		err := client.Roles.Delete(ctx, "436b4bbc-1111-44a5-741e-da04329bc3bb")
		assert.NoError(t, err)
	})
}

func TestRolesListOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("WithLimit", func(t *testing.T) {
		roles, err := client.Roles.List(ctx, converged.WithLimit(2))
		assert.NoError(t, err)
		assert.NotNil(t, roles)
		assert.LessOrEqual(t, len(roles), 2)
	})

	t.Run("WithSelect", func(t *testing.T) {
		roles, err := client.Roles.List(ctx, converged.WithSelect("extId,displayName"))
		assert.NoError(t, err)
		assert.NotNil(t, roles)
	})

	t.Run("WithApply_NotSupported", func(t *testing.T) {
		_, err := client.Roles.List(ctx, converged.WithApply("groupby((displayName))"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing Roles")
	})
}
