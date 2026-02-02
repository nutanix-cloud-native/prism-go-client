// Package v4 provides integration tests for the converged client templates service.
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run TestTemplatesIntegration
//	go test -v ./converged/v4 -run TestTemplatesService_ErrorHandling
package v4

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"

	templateModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// TestTemplatesService_ErrorHandling tests error handling for nil client
func TestTemplatesService_ErrorHandling(t *testing.T) {
	service := NewTemplatesService(nil)
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

	t.Run("Create_NilClient", func(t *testing.T) {
		template := templateModels.NewTemplate()
		_, err := service.Create(ctx, template)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("CreateAsync_NilClient", func(t *testing.T) {
		template := templateModels.NewTemplate()
		_, err := service.CreateAsync(ctx, template)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})
}

// TestTemplatesIntegration tests the client.Templates with real Nutanix API calls
func TestTemplatesIntegration(t *testing.T) {
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

	t.Run("ListTemplates", func(t *testing.T) {
		// Test List templates
		templates, err := client.Templates.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, templates)
		assert.GreaterOrEqual(t, len(templates), 0)
	})

	t.Run("GetTemplate", func(t *testing.T) {
		// First, list templates to get a valid template UUID
		templates, err := client.Templates.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(templates) == 0 {
			t.Skip("No templates available for testing")
		}

		// Get the first template's UUID
		templateUUID := *templates[0].ExtId
		require.NotEmpty(t, templateUUID)

		// Test Get template
		template, err := client.Templates.Get(ctx, templateUUID)
		assert.NoError(t, err)
		assert.NotNil(t, template)
		assert.Equal(t, templateUUID, *template.ExtId)
	})

	t.Run("NewIterator", func(t *testing.T) {
		// Test iterator
		iterator := client.Templates.NewIterator(ctx)
		require.NotNil(t, iterator)

		// Collect templates using iterator
		var templates []templateModels.Template
		for template, err := range iterator {
			if err != nil {
				break
			}
			templates = append(templates, template)
			if len(templates) >= 5 { // Limit to prevent long test runs
				break
			}
		}

		assert.GreaterOrEqual(t, len(templates), 0)
	})
}

// TestTemplatesListOptions tests various OData options for listing templates
func TestTemplatesListOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("WithLimit", func(t *testing.T) {
		templates, err := client.Templates.List(ctx, converged.WithLimit(2))
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(templates), 2)
	})

	t.Run("WithPage", func(t *testing.T) {
		templates, err := client.Templates.List(ctx, converged.WithPage(0), converged.WithLimit(1))
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(templates), 1)
	})

	t.Run("WithSelect", func(t *testing.T) {
		templates, err := client.Templates.List(ctx, converged.WithSelect("extId,templateName"))
		assert.NoError(t, err)
		assert.NotNil(t, templates)
	})

	t.Run("WithOrderBy", func(t *testing.T) {
		templates, err := client.Templates.List(ctx, converged.WithOrderBy("templateName asc"))
		assert.NoError(t, err)
		assert.NotNil(t, templates)
	})

	t.Run("MultipleOptions", func(t *testing.T) {
		templates, err := client.Templates.List(ctx,
			converged.WithLimit(5),
			converged.WithPage(0),
			converged.WithSelect("extId,templateName"),
		)
		assert.NoError(t, err)
		assert.NotNil(t, templates)
		assert.LessOrEqual(t, len(templates), 5)
	})

	t.Run("UnsupportedExpandOption", func(t *testing.T) {
		// Expand is not supported for Templates
		_, err := client.Templates.List(ctx, converged.WithExpand("someField"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "expand and apply are not supported")
	})

	t.Run("UnsupportedApplyOption", func(t *testing.T) {
		// Apply is not supported for Templates
		_, err := client.Templates.List(ctx, converged.WithApply("someAggregation"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "expand and apply are not supported")
	})
}

// TestTemplatesErrorScenarios tests error handling scenarios
func TestTemplatesErrorScenarios(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("GetInvalidUUID", func(t *testing.T) {
		_, err := client.Templates.Get(ctx, "invalid-uuid-format")
		assert.Error(t, err)
	})

	t.Run("GetNonExistentUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.Templates.Get(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("InvalidFilterSyntax", func(t *testing.T) {
		_, err := client.Templates.List(ctx, converged.WithFilter("invalid filter syntax"))
		assert.Error(t, err)
	})
}
