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

func TestAuthorizationPoliciesService_ErrorHandling(t *testing.T) {
	service := NewAuthorizationPoliciesService(nil)
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
		policy := iamAuthzModels.NewAuthorizationPolicy()
		_, err := service.Create(ctx, policy)
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

// TestAuthorizationPoliciesIntegration tests client.AuthorizationPolicies with real Nutanix API calls.
func TestAuthorizationPoliciesIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ListAuthorizationPolicies", func(t *testing.T) {
		policies, err := client.AuthorizationPolicies.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, policies)
		assert.GreaterOrEqual(t, len(policies), 0)
	})

	t.Run("GetAuthorizationPolicy", func(t *testing.T) {
		policies, err := client.AuthorizationPolicies.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)
		if len(policies) == 0 {
			t.Skip("No authorization policies available for testing")
		}

		policyUUID := *policies[0].ExtId
		require.NotEmpty(t, policyUUID)

		policy, err := client.AuthorizationPolicies.Get(ctx, policyUUID)
		assert.NoError(t, err)
		assert.NotNil(t, policy)
		assert.Equal(t, policyUUID, *policy.ExtId)
	})
}

func TestAuthorizationPoliciesListOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("WithLimit", func(t *testing.T) {
		policies, err := client.AuthorizationPolicies.List(ctx, converged.WithLimit(2))
		assert.NoError(t, err)
		assert.NotNil(t, policies)
		assert.LessOrEqual(t, len(policies), 2)
	})

	t.Run("WithSelect", func(t *testing.T) {
		policies, err := client.AuthorizationPolicies.List(ctx, converged.WithSelect("extId,displayName"))
		assert.NoError(t, err)
		assert.NotNil(t, policies)
	})

	t.Run("WithExpand", func(t *testing.T) {
		policies, err := client.AuthorizationPolicies.List(ctx, converged.WithExpand("role"))
		assert.NoError(t, err)
		assert.NotNil(t, policies)
	})

	t.Run("WithApply_NotSupported", func(t *testing.T) {
		_, err := client.AuthorizationPolicies.List(ctx, converged.WithApply("groupby((name))"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply option is not supported for listing Authorization Policies")
	})
}
