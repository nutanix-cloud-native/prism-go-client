package test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	models "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	v4ConvergedClient "github.com/nutanix-cloud-native/prism-go-client/converged_client/v4"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"
)

var client *v4ConvergedClient.Client

func initializeConvergedClient(t *testing.T) {
	if client != nil {
		return
	}

	var err error
	creds := testhelpers.CredentialsFromEnvironment(t)

	client, err = v4ConvergedClient.NewClient(creds)
	assert.Nil(t, err)
}

func TestCategoriesService_Get(t *testing.T) {
	initializeConvergedClient(t)
	ctx := context.Background()

	categories, err := client.Categories.List(ctx, models.WithPage(0))
	assert.Nil(t, err)
	assert.True(t, len(categories) > 0)

	_, err = client.Categories.Get(ctx, *categories[0].ExtId)
	assert.Nil(t, err)
}

func TestCategoriesService_List(t *testing.T) {
	initializeConvergedClient(t)
	ctx := context.Background()

	// Test supported operations - these should succeed
	supportedTestCases := []struct {
		name string
		opts []models.ODataOption
	}{
		{
			name: "no options",
			opts: []models.ODataOption{},
		},
		{
			name: "with page option",
			opts: []models.ODataOption{models.WithPage(1)},
		},
		{
			name: "with limit option",
			opts: []models.ODataOption{models.WithLimit(10)},
		},
		{
			name: "with apply option",
			opts: []models.ODataOption{models.WithApply("groupby((name))")},
		},
	}

	for _, tt := range supportedTestCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.Categories.List(ctx, tt.opts...)
			assert.Nil(t, err)
		})
	}

	// Test unsupported operations - these should fail with 400 BAD REQUEST
	unsupportedTestCases := []struct {
		name string
		opts []models.ODataOption
	}{
		{
			name: "with filter option (unsupported)",
			opts: []models.ODataOption{models.WithFilter("name eq 'test'")},
		},
		{
			name: "with order by option (unsupported)",
			opts: []models.ODataOption{models.WithOrderBy("name")},
		},
		{
			name: "with expand option (unsupported)",
			opts: []models.ODataOption{models.WithExpand("metadata")},
		},
		{
			name: "with select option (unsupported)",
			opts: []models.ODataOption{models.WithSelect("name")},
		},
		{
			name: "with multiple unsupported options",
			opts: []models.ODataOption{
				models.WithPage(1),
				models.WithLimit(10),
				models.WithFilter("name eq 'test'"),
			},
		},
	}

	for _, tt := range unsupportedTestCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.Categories.List(ctx, tt.opts...)
			// Expect error for unsupported OData parameters
			assert.NotNil(t, err)
			assert.True(t, strings.Contains(err.Error(), "parsing the URI parameters") ||
				strings.Contains(err.Error(), "400"))
		})
	}
}

func TestCategoriesService_NewIterator(t *testing.T) {
	initializeConvergedClient(t)
	ctx := context.Background()

	// Test supported operations - these should succeed
	supportedTestCases := []struct {
		name string
		opts []models.ODataOption
	}{
		{
			name: "no options",
			opts: []models.ODataOption{},
		},
	}

	for _, tt := range supportedTestCases {
		t.Run(tt.name, func(t *testing.T) {
			iterator := client.Categories.NewIterator(ctx, tt.opts...)
			assert.NotNil(t, iterator)
			for _, err := range iterator {
				assert.Nil(t, err)
			}
		})
	}

	// Test unsupported operations - these should fail when iterating
	unsupportedTestCases := []struct {
		name string
		opts []models.ODataOption
	}{
		{
			name: "with filter option (unsupported)",
			opts: []models.ODataOption{models.WithFilter("name eq 'test'")},
		},
		{
			name: "with order by option (unsupported)",
			opts: []models.ODataOption{models.WithOrderBy("name")},
		},
		{
			name: "with expand option (unsupported)",
			opts: []models.ODataOption{models.WithExpand("metadata")},
		},
		{
			name: "with select option (unsupported)",
			opts: []models.ODataOption{models.WithSelect("name")},
		},
		{
			name: "with multiple unsupported options",
			opts: []models.ODataOption{
				models.WithFilter("name eq 'test'"),
				models.WithOrderBy("name"),
				models.WithExpand("metadata"),
			},
		},
	}

	for _, tt := range unsupportedTestCases {
		t.Run(tt.name, func(t *testing.T) {
			iterator := client.Categories.NewIterator(ctx, tt.opts...)
			assert.NotNil(t, iterator)

			// Expect error when iterating with unsupported parameters
			var foundError error
			for _, err := range iterator {
				if err != nil {
					foundError = err
					break
				}
			}

			assert.NotNil(t, foundError)
			assert.True(t, strings.Contains(foundError.Error(), "parsing the URI parameters") ||
				strings.Contains(foundError.Error(), "400"))
		})
	}
}

func TestCategoriesService_Create(t *testing.T) {
	initializeConvergedClient(t)
	ctx := context.Background()

	// Create a unique category for testing
	timestamp := time.Now().Unix()
	categoryKey := fmt.Sprintf("test-category-key-%d", timestamp)
	categoryValue := fmt.Sprintf("test-category-value-%d", timestamp)
	categoryDescription := fmt.Sprintf("Integration test category created at %d", timestamp)

	category := &prismModels.Category{
		Key:         ptr.To(categoryKey),
		Value:       ptr.To(categoryValue),
		Description: ptr.To(categoryDescription),
	}

	// Test Create operation
	createdCategory, err := client.Categories.Create(ctx, category)
	require.Nil(t, err)
	require.NotNil(t, createdCategory)
	require.NotNil(t, createdCategory.ExtId)
	assert.Equal(t, categoryKey, *createdCategory.Key)
	assert.Equal(t, categoryValue, *createdCategory.Value)
	assert.Equal(t, categoryDescription, *createdCategory.Description)

	// Clean up - delete the created category
	defer func() {
		if createdCategory != nil && createdCategory.ExtId != nil {
			_ = client.Categories.Delete(ctx, *createdCategory.ExtId)
		}
	}()
}

func TestCategoriesService_Update(t *testing.T) {
	initializeConvergedClient(t)
	ctx := context.Background()

	// First, create a category to update
	timestamp := time.Now().Unix()
	originalKey := fmt.Sprintf("test-update-key-%d", timestamp)
	originalValue := fmt.Sprintf("original-value-%d", timestamp)
	originalDescription := fmt.Sprintf("Original description %d", timestamp)

	category := &prismModels.Category{
		Key:         ptr.To(originalKey),
		Value:       ptr.To(originalValue),
		Description: ptr.To(originalDescription),
	}

	createdCategory, err := client.Categories.Create(ctx, category)
	require.Nil(t, err)
	require.NotNil(t, createdCategory)
	require.NotNil(t, createdCategory.ExtId)

	// Clean up at the end
	defer func() {
		if createdCategory != nil && createdCategory.ExtId != nil {
			_ = client.Categories.Delete(ctx, *createdCategory.ExtId)
		}
	}()

	// Now update the category (Key cannot be updated, but Value and Description can be)
	updatedValue := fmt.Sprintf("updated-value-%d", timestamp)
	updatedDescription := fmt.Sprintf("Updated description %d", timestamp)

	updateCategory := &prismModels.Category{
		ExtId:       createdCategory.ExtId,
		Key:         createdCategory.Key, // Keep original key (cannot be changed)
		Value:       ptr.To(updatedValue),
		Description: ptr.To(updatedDescription),
		OwnerUuid:   createdCategory.OwnerUuid, // Preserve owner (required)
	}

	updatedCategory, err := client.Categories.Update(ctx, *createdCategory.ExtId, updateCategory)
	require.Nil(t, err)
	require.NotNil(t, updatedCategory)
	assert.Equal(t, *createdCategory.ExtId, *updatedCategory.ExtId)
	assert.Equal(t, originalKey, *updatedCategory.Key) // Key should remain unchanged
	assert.Equal(t, updatedValue, *updatedCategory.Value)
	assert.Equal(t, updatedDescription, *updatedCategory.Description)
}

func TestCategoriesService_Delete(t *testing.T) {
	initializeConvergedClient(t)
	ctx := context.Background()

	// First, create a category to delete
	timestamp := time.Now().Unix()
	categoryKey := fmt.Sprintf("test-delete-key-%d", timestamp)
	categoryValue := fmt.Sprintf("test-delete-value-%d", timestamp)
	categoryDescription := fmt.Sprintf("Category to be deleted %d", timestamp)

	category := &prismModels.Category{
		Key:         ptr.To(categoryKey),
		Value:       ptr.To(categoryValue),
		Description: ptr.To(categoryDescription),
	}

	createdCategory, err := client.Categories.Create(ctx, category)
	require.Nil(t, err)
	require.NotNil(t, createdCategory)
	require.NotNil(t, createdCategory.ExtId)

	// Verify the category exists
	retrievedCategory, err := client.Categories.Get(ctx, *createdCategory.ExtId)
	require.Nil(t, err)
	require.NotNil(t, retrievedCategory)
	assert.Equal(t, *createdCategory.ExtId, *retrievedCategory.ExtId)

	// Now delete the category
	err = client.Categories.Delete(ctx, *createdCategory.ExtId)
	require.Nil(t, err)

	// Verify the category no longer exists
	deletedCategory, err := client.Categories.Get(ctx, *createdCategory.ExtId)
	assert.NotNil(t, err) // Should return an error since category doesn't exist
	assert.Nil(t, deletedCategory)
}

func TestCategoriesService_CRUD_Integration(t *testing.T) {
	initializeConvergedClient(t)
	ctx := context.Background()

	// Test full CRUD cycle
	timestamp := time.Now().Unix()
	originalKey := fmt.Sprintf("test-crud-key-%d", timestamp)
	originalValue := fmt.Sprintf("original-crud-value-%d", timestamp)
	originalDescription := fmt.Sprintf("CRUD test category %d", timestamp)

	// CREATE
	category := &prismModels.Category{
		Key:         ptr.To(originalKey),
		Value:       ptr.To(originalValue),
		Description: ptr.To(originalDescription),
	}

	createdCategory, err := client.Categories.Create(ctx, category)
	require.Nil(t, err)
	require.NotNil(t, createdCategory)
	require.NotNil(t, createdCategory.ExtId)
	t.Logf("Created category with ID: %s", *createdCategory.ExtId)

	// Clean up at the end
	defer func() {
		if createdCategory != nil && createdCategory.ExtId != nil {
			_ = client.Categories.Delete(ctx, *createdCategory.ExtId)
		}
	}()

	// READ (Get)
	retrievedCategory, err := client.Categories.Get(ctx, *createdCategory.ExtId)
	require.Nil(t, err)
	require.NotNil(t, retrievedCategory)
	assert.Equal(t, *createdCategory.ExtId, *retrievedCategory.ExtId)
	assert.Equal(t, originalKey, *retrievedCategory.Key)
	assert.Equal(t, originalValue, *retrievedCategory.Value)
	t.Logf("Retrieved category: %s=%s", *retrievedCategory.Key, *retrievedCategory.Value)

	// UPDATE
	updatedValue := fmt.Sprintf("updated-crud-value-%d", timestamp)
	updatedDescription := fmt.Sprintf("Updated CRUD test category %d", timestamp)

	updateCategory := &prismModels.Category{
		ExtId:       createdCategory.ExtId,
		Key:         createdCategory.Key, // Key cannot be changed
		Value:       ptr.To(updatedValue),
		Description: ptr.To(updatedDescription),
		OwnerUuid:   createdCategory.OwnerUuid, // Preserve owner (required)
	}

	updatedCategory, err := client.Categories.Update(ctx, *createdCategory.ExtId, updateCategory)
	require.Nil(t, err)
	require.NotNil(t, updatedCategory)
	assert.Equal(t, *createdCategory.ExtId, *updatedCategory.ExtId)
	assert.Equal(t, originalKey, *updatedCategory.Key) // Key should remain unchanged
	assert.Equal(t, updatedValue, *updatedCategory.Value)
	assert.Equal(t, updatedDescription, *updatedCategory.Description)
	t.Logf("Updated category value to: %s=%s", *updatedCategory.Key, *updatedCategory.Value)

	// READ again to verify update
	updatedRetrievedCategory, err := client.Categories.Get(ctx, *createdCategory.ExtId)
	require.Nil(t, err)
	require.NotNil(t, updatedRetrievedCategory)
	assert.Equal(t, originalKey, *updatedRetrievedCategory.Key)
	assert.Equal(t, updatedValue, *updatedRetrievedCategory.Value)
	assert.Equal(t, updatedDescription, *updatedRetrievedCategory.Description)

	// DELETE
	err = client.Categories.Delete(ctx, *createdCategory.ExtId)
	require.Nil(t, err)
	t.Logf("Deleted category with ID: %s", *createdCategory.ExtId)

	// Verify deletion
	deletedCategory, err := client.Categories.Get(ctx, *createdCategory.ExtId)
	assert.NotNil(t, err) // Should return an error since category doesn't exist
	assert.Nil(t, deletedCategory)
	t.Logf("Verified category deletion - category no longer exists")

	// Clear the createdCategory to prevent cleanup from trying to delete again
	createdCategory = nil
}
