// Package v4 provides integration tests for the converged client projects service.
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run TestProjectsIntegration
//	go test -v ./converged/v4 -run TestProjectsService_ErrorHandling
//	go test -v ./converged/v4 -run TestProjectsService_UnsupportedODataOptions
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

// TestProjectsIntegration tests the client.Projects service with real Nutanix API calls.
func TestProjectsIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ListProjects", func(t *testing.T) {
		projects, err := client.Projects.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, projects)
		assert.GreaterOrEqual(t, len(projects), 0)

		if len(projects) > 0 {
			assert.NotNil(t, projects[0].ExtId)
		}
	})

	t.Run("ListProjectsWithValidOptions", func(t *testing.T) {
		projects, err := client.Projects.List(ctx,
			converged.WithPage(0),
			converged.WithLimit(5),
			converged.WithSelect("extId,name"),
			converged.WithOrderBy("name asc"),
			converged.WithFilter("name ne null"),
		)
		assert.NoError(t, err)
		assert.NotNil(t, projects)
	})

	t.Run("GetProject", func(t *testing.T) {
		projects, err := client.Projects.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(projects) == 0 {
			t.Skip("No projects available for testing")
		}

		extID := *projects[0].ExtId
		require.NotEmpty(t, extID)

		project, err := client.Projects.Get(ctx, extID)
		assert.NoError(t, err)
		assert.NotNil(t, project)
		assert.Equal(t, extID, *project.ExtId)
	})

	t.Run("NewIterator", func(t *testing.T) {
		iterator := client.Projects.NewIterator(ctx, converged.WithLimit(5))
		assert.NotNil(t, iterator)

		var count int
		for _, err := range iterator {
			if err != nil {
				t.Logf("Iterator error: %v", err)
				break
			}
			count++
		}
		assert.GreaterOrEqual(t, count, 0)
	})

	t.Run("GetDefaultProject", func(t *testing.T) {
		defaultProject, err := client.Projects.GetDefaultProject(ctx)
		if err != nil {
			// Empty filtered list means the caller cannot see the system default
			// project (RBAC), not that it is missing.
			assert.True(t, converged.IsUnauthorized(err))
			assert.Contains(t, err.Error(), "not authorized to access the default project")
			return
		}
		assert.NotNil(t, defaultProject)
		assert.NotNil(t, defaultProject.ExtId)
		if defaultProject.IsDefault != nil {
			assert.True(t, *defaultProject.IsDefault)
		}
	})
}

// TestProjectsService_ErrorHandling tests error handling for nil client.
func TestProjectsService_ErrorHandling(t *testing.T) {
	service := NewProjectsService(nil)
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

	t.Run("GetDefaultProject_NilClient", func(t *testing.T) {
		_, err := service.GetDefaultProject(ctx)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})
}

// TestProjectsService_UnsupportedODataOptions tests that apply and expand options are rejected.
func TestProjectsService_UnsupportedODataOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("List_WithExpand", func(t *testing.T) {
		_, err := client.Projects.List(ctx, converged.WithExpand("config"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing Projects")
	})

	t.Run("List_WithApply", func(t *testing.T) {
		_, err := client.Projects.List(ctx, converged.WithApply("groupby((extId))"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing Projects")
	})

	t.Run("List_WithExpandAndLimit", func(t *testing.T) {
		_, err := client.Projects.List(ctx, converged.WithExpand("config"), converged.WithLimit(5))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing Projects")
	})

	t.Run("NewIterator_WithExpand", func(t *testing.T) {
		iter := client.Projects.NewIterator(ctx, converged.WithExpand("config"))
		require.NotNil(t, iter)

		var gotError error
		for _, err := range iter {
			gotError = err
			break
		}
		require.Error(t, gotError)
		assert.Contains(t, gotError.Error(), "apply and expand options are not supported for listing Projects")
	})

	t.Run("NewIterator_WithApply", func(t *testing.T) {
		iter := client.Projects.NewIterator(ctx, converged.WithApply("groupby((extId))"))
		require.NotNil(t, iter)

		var gotError error
		for _, err := range iter {
			gotError = err
			break
		}
		require.Error(t, gotError)
		assert.Contains(t, gotError.Error(), "apply and expand options are not supported for listing Projects")
	})
}
