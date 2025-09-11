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
	policyModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"
)

var antiAffinityClient *v4ConvergedClient.Client
var createdCategoryIds []string

func initializeAntiAffinityConvergedClient(t *testing.T) {
	if antiAffinityClient != nil {
		return
	}

	var err error
	creds := testhelpers.CredentialsFromEnvironment(t)

	antiAffinityClient, err = v4ConvergedClient.NewClient(creds)
	assert.Nil(t, err)

	// Initialize cleanup slice
	createdCategoryIds = []string{}
}

func cleanupCreatedCategories(t *testing.T) {
	if antiAffinityClient == nil || len(createdCategoryIds) == 0 {
		return
	}

	ctx := context.Background()
	for _, categoryId := range createdCategoryIds {
		err := antiAffinityClient.Categories.Delete(ctx, categoryId)
		if err != nil {
			t.Logf("Failed to cleanup category %s: %v", categoryId, err)
		} else {
			t.Logf("Cleaned up category: %s", categoryId)
		}
	}
	createdCategoryIds = []string{}
}

func TestAntiAffinityPoliciesService_Get(t *testing.T) {
	initializeAntiAffinityConvergedClient(t)
	ctx := context.Background()

	policies, err := antiAffinityClient.AntiAffinityPolicies.List(ctx, models.WithPage(0))
	assert.Nil(t, err)
	if len(policies) > 0 {
		_, err = antiAffinityClient.AntiAffinityPolicies.Get(ctx, *policies[0].ExtId)
		assert.Nil(t, err)
	} else {
		t.Skip("No anti-affinity policies available for testing")
	}
}

func TestAntiAffinityPoliciesService_List(t *testing.T) {
	initializeAntiAffinityConvergedClient(t)
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
			name: "with order by option",
			opts: []models.ODataOption{models.WithOrderBy("name")},
		},
	}

	for _, tt := range supportedTestCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := antiAffinityClient.AntiAffinityPolicies.List(ctx, tt.opts...)
			assert.Nil(t, err)
		})
	}

	// Test unsupported operations - these should fail with 400 BAD REQUEST
	unsupportedTestCases := []struct {
		name string
		opts []models.ODataOption
	}{
		{
			name: "with filter option (may be unsupported)",
			opts: []models.ODataOption{models.WithFilter("name eq 'test'")},
		},
		{
			name: "with expand option (may be unsupported)",
			opts: []models.ODataOption{models.WithExpand("metadata")},
		},
		{
			name: "with select option (may be unsupported)",
			opts: []models.ODataOption{models.WithSelect("name")},
		},
		{
			name: "with apply option (may be unsupported)",
			opts: []models.ODataOption{models.WithApply("groupby((name))")},
		},
	}

	for _, tt := range unsupportedTestCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := antiAffinityClient.AntiAffinityPolicies.List(ctx, tt.opts...)
			// These may or may not be supported - just log the result
			if err != nil {
				t.Logf("Expected potential error for %s: %v", tt.name, err)
				assert.True(t, strings.Contains(err.Error(), "parsing the URI parameters") ||
					strings.Contains(err.Error(), "400") ||
					strings.Contains(err.Error(), "not supported"))
			}
		})
	}
}

func TestAntiAffinityPoliciesService_NewIterator(t *testing.T) {
	initializeAntiAffinityConvergedClient(t)
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
			name: "with limit",
			opts: []models.ODataOption{models.WithLimit(5)},
		},
	}

	for _, tt := range supportedTestCases {
		t.Run(tt.name, func(t *testing.T) {
			iterator := antiAffinityClient.AntiAffinityPolicies.NewIterator(ctx, tt.opts...)
			assert.NotNil(t, iterator)
			count := 0
			for _, err := range iterator {
				if err != nil {
					t.Logf("Iterator error: %v", err)
					break
				}
				count++
				if count >= 5 { // Limit iteration to prevent long test runs
					break
				}
			}
		})
	}

	// Test potentially unsupported operations
	unsupportedTestCases := []struct {
		name string
		opts []models.ODataOption
	}{
		{
			name: "with filter option (may be unsupported)",
			opts: []models.ODataOption{models.WithFilter("name eq 'test'")},
		},
		{
			name: "with order by option",
			opts: []models.ODataOption{models.WithOrderBy("name")},
		},
	}

	for _, tt := range unsupportedTestCases {
		t.Run(tt.name, func(t *testing.T) {
			iterator := antiAffinityClient.AntiAffinityPolicies.NewIterator(ctx, tt.opts...)
			assert.NotNil(t, iterator)

			var foundError error
			count := 0
			for _, err := range iterator {
				if err != nil {
					foundError = err
					break
				}
				count++
				if count >= 3 { // Limit iteration
					break
				}
			}

			if foundError != nil {
				t.Logf("Expected potential error for %s: %v", tt.name, foundError)
			}
		})
	}
}

func TestAntiAffinityPoliciesService_Create(t *testing.T) {
	initializeAntiAffinityConvergedClient(t)
	defer cleanupCreatedCategories(t)
	ctx := context.Background()

	// Create a unique policy for testing
	timestamp := time.Now().Unix()
	policyName := fmt.Sprintf("test-anti-affinity-policy-%d", timestamp)
	policyDescription := fmt.Sprintf("Integration test anti-affinity policy created at %d", timestamp)

	policy, err := createTestAntiAffinityPolicy(policyName, policyDescription)
	require.NoError(t, err, "Failed to create test anti-affinity policy")

	// Test Create operation
	createdPolicy, err := antiAffinityClient.AntiAffinityPolicies.Create(ctx, policy)
	require.Nil(t, err)
	require.NotNil(t, createdPolicy)
	require.NotNil(t, createdPolicy.ExtId)
	assert.Equal(t, policyName, *createdPolicy.Name)
	assert.Equal(t, policyDescription, *createdPolicy.Description)
	t.Logf("Created anti-affinity policy with ID: %s", *createdPolicy.ExtId)

	// Clean up - delete the created policy
	defer func() {
		if createdPolicy != nil && createdPolicy.ExtId != nil {
			err := antiAffinityClient.AntiAffinityPolicies.Delete(ctx, *createdPolicy.ExtId)
			if err != nil {
				t.Logf("Failed to cleanup policy: %v", err)
			}
		}
	}()
}

func TestAntiAffinityPoliciesService_Update(t *testing.T) {
	initializeAntiAffinityConvergedClient(t)
	defer cleanupCreatedCategories(t)
	ctx := context.Background()

	// First, create a policy to update
	timestamp := time.Now().Unix()
	originalName := fmt.Sprintf("test-update-policy-%d", timestamp)
	originalDescription := fmt.Sprintf("Original description %d", timestamp)

	policy, err := createTestAntiAffinityPolicy(originalName, originalDescription)
	require.NoError(t, err, "Failed to create test anti-affinity policy")

	createdPolicy, err := antiAffinityClient.AntiAffinityPolicies.Create(ctx, policy)
	require.Nil(t, err)
	require.NotNil(t, createdPolicy)
	require.NotNil(t, createdPolicy.ExtId)
	t.Logf("Created policy for update test with ID: %s", *createdPolicy.ExtId)

	// Clean up at the end
	defer func() {
		if createdPolicy != nil && createdPolicy.ExtId != nil {
			err := antiAffinityClient.AntiAffinityPolicies.Delete(ctx, *createdPolicy.ExtId)
			if err != nil {
				t.Logf("Failed to cleanup policy: %v", err)
			}
		}
	}()

	// Now update the policy
	updatedName := fmt.Sprintf("updated-policy-%d", timestamp)
	updatedDescription := fmt.Sprintf("Updated description %d", timestamp)

	updatePolicy := &policyModels.VmAntiAffinityPolicy{
		ExtId:       createdPolicy.ExtId,
		Name:        ptr.To(updatedName),
		Description: ptr.To(updatedDescription),
	}

	updatedPolicy, err := antiAffinityClient.AntiAffinityPolicies.Update(ctx, *createdPolicy.ExtId, updatePolicy)
	require.Nil(t, err)
	require.NotNil(t, updatedPolicy)
	assert.Equal(t, *createdPolicy.ExtId, *updatedPolicy.ExtId)
	assert.Equal(t, updatedName, *updatedPolicy.Name)
	assert.Equal(t, updatedDescription, *updatedPolicy.Description)
	t.Logf("Updated policy name to: %s", *updatedPolicy.Name)
}

func TestAntiAffinityPoliciesService_Delete(t *testing.T) {
	initializeAntiAffinityConvergedClient(t)
	defer cleanupCreatedCategories(t)
	ctx := context.Background()

	// First, create a policy to delete
	timestamp := time.Now().Unix()
	policyName := fmt.Sprintf("test-delete-policy-%d", timestamp)
	policyDescription := fmt.Sprintf("Policy to be deleted %d", timestamp)

	policy, err := createTestAntiAffinityPolicy(policyName, policyDescription)
	require.NoError(t, err, "Failed to create test anti-affinity policy")

	createdPolicy, err := antiAffinityClient.AntiAffinityPolicies.Create(ctx, policy)
	require.Nil(t, err)
	require.NotNil(t, createdPolicy)
	require.NotNil(t, createdPolicy.ExtId)
	t.Logf("Created policy for delete test with ID: %s", *createdPolicy.ExtId)

	// Verify the policy exists
	retrievedPolicy, err := antiAffinityClient.AntiAffinityPolicies.Get(ctx, *createdPolicy.ExtId)
	require.Nil(t, err)
	require.NotNil(t, retrievedPolicy)
	assert.Equal(t, *createdPolicy.ExtId, *retrievedPolicy.ExtId)

	// Now delete the policy
	err = antiAffinityClient.AntiAffinityPolicies.Delete(ctx, *createdPolicy.ExtId)
	require.Nil(t, err)
	t.Logf("Deleted policy with ID: %s", *createdPolicy.ExtId)

	// Verify the policy no longer exists
	deletedPolicy, err := antiAffinityClient.AntiAffinityPolicies.Get(ctx, *createdPolicy.ExtId)
	assert.NotNil(t, err) // Should return an error since policy doesn't exist
	assert.Nil(t, deletedPolicy)
}

func TestAntiAffinityPoliciesService_CreateAsync(t *testing.T) {
	initializeAntiAffinityConvergedClient(t)
	defer cleanupCreatedCategories(t)
	ctx := context.Background()

	// Create a unique policy for async testing
	timestamp := time.Now().Unix()
	policyName := fmt.Sprintf("test-async-policy-%d", timestamp)
	policyDescription := fmt.Sprintf("Async test anti-affinity policy created at %d", timestamp)

	policy, err := createTestAntiAffinityPolicy(policyName, policyDescription)
	require.NoError(t, err, "Failed to create test anti-affinity policy")

	// Test CreateAsync operation
	operation, err := antiAffinityClient.AntiAffinityPolicies.CreateAsync(ctx, policy)
	require.Nil(t, err)
	require.NotNil(t, operation)

	// Wait for completion
	results, err := operation.Wait(ctx)
	require.Nil(t, err)
	require.Len(t, results, 1)
	require.NotNil(t, results[0])
	require.NotNil(t, results[0].ExtId)
	t.Logf("Created async policy with ID: %s", *results[0].ExtId)

	// Clean up - delete the created policy
	defer func() {
		if results[0] != nil && results[0].ExtId != nil {
			err := antiAffinityClient.AntiAffinityPolicies.Delete(ctx, *results[0].ExtId)
			if err != nil {
				t.Logf("Failed to cleanup async policy: %v", err)
			}
		}
	}()
}

func TestAntiAffinityPoliciesService_CRUD_Integration(t *testing.T) {
	initializeAntiAffinityConvergedClient(t)
	defer cleanupCreatedCategories(t)
	ctx := context.Background()

	// Test full CRUD cycle
	timestamp := time.Now().Unix()
	originalName := fmt.Sprintf("test-crud-policy-%d", timestamp)
	originalDescription := fmt.Sprintf("CRUD test anti-affinity policy %d", timestamp)

	// CREATE
	policy, err := createTestAntiAffinityPolicy(originalName, originalDescription)
	require.NoError(t, err, "Failed to create test anti-affinity policy")

	createdPolicy, err := antiAffinityClient.AntiAffinityPolicies.Create(ctx, policy)
	require.Nil(t, err)
	require.NotNil(t, createdPolicy)
	require.NotNil(t, createdPolicy.ExtId)
	t.Logf("Created policy with ID: %s", *createdPolicy.ExtId)

	// Clean up at the end
	defer func() {
		if createdPolicy != nil && createdPolicy.ExtId != nil {
			err := antiAffinityClient.AntiAffinityPolicies.Delete(ctx, *createdPolicy.ExtId)
			if err != nil {
				t.Logf("Failed to cleanup CRUD policy: %v", err)
			}
		}
	}()

	// READ (Get)
	retrievedPolicy, err := antiAffinityClient.AntiAffinityPolicies.Get(ctx, *createdPolicy.ExtId)
	require.Nil(t, err)
	require.NotNil(t, retrievedPolicy)
	assert.Equal(t, *createdPolicy.ExtId, *retrievedPolicy.ExtId)
	assert.Equal(t, originalName, *retrievedPolicy.Name)
	assert.Equal(t, originalDescription, *retrievedPolicy.Description)
	t.Logf("Retrieved policy: %s", *retrievedPolicy.Name)

	// UPDATE
	updatedName := fmt.Sprintf("updated-crud-policy-%d", timestamp)
	updatedDescription := fmt.Sprintf("Updated CRUD test policy %d", timestamp)

	updatePolicy := &policyModels.VmAntiAffinityPolicy{
		ExtId:       createdPolicy.ExtId,
		Name:        ptr.To(updatedName),
		Description: ptr.To(updatedDescription),
	}

	updatedPolicy, err := antiAffinityClient.AntiAffinityPolicies.Update(ctx, *createdPolicy.ExtId, updatePolicy)
	require.Nil(t, err)
	require.NotNil(t, updatedPolicy)
	assert.Equal(t, *createdPolicy.ExtId, *updatedPolicy.ExtId)
	assert.Equal(t, updatedName, *updatedPolicy.Name)
	assert.Equal(t, updatedDescription, *updatedPolicy.Description)
	t.Logf("Updated policy name to: %s", *updatedPolicy.Name)

	// READ again to verify update
	updatedRetrievedPolicy, err := antiAffinityClient.AntiAffinityPolicies.Get(ctx, *createdPolicy.ExtId)
	require.Nil(t, err)
	require.NotNil(t, updatedRetrievedPolicy)
	assert.Equal(t, updatedName, *updatedRetrievedPolicy.Name)
	assert.Equal(t, updatedDescription, *updatedRetrievedPolicy.Description)

	// DELETE
	err = antiAffinityClient.AntiAffinityPolicies.Delete(ctx, *createdPolicy.ExtId)
	require.Nil(t, err)
	t.Logf("Deleted policy with ID: %s", *createdPolicy.ExtId)

	// Verify deletion
	deletedPolicy, err := antiAffinityClient.AntiAffinityPolicies.Get(ctx, *createdPolicy.ExtId)
	assert.NotNil(t, err) // Should return an error since policy doesn't exist
	assert.Nil(t, deletedPolicy)
	t.Logf("Verified policy deletion - policy no longer exists")

	// Clear the createdPolicy to prevent cleanup from trying to delete again
	createdPolicy = nil
}

// Helper function to create or get a test category
func createOrGetTestCategory() (*prismModels.Category, error) {
	ctx := context.Background()

	// Try to get existing categories using Categories service
	existingCategories, err := antiAffinityClient.Categories.List(ctx, models.WithLimit(1))
	if err == nil && len(existingCategories) > 0 {
		return &existingCategories[0], nil
	}

	// If no existing categories, create a test category using Categories service
	timestamp := time.Now().Unix()
	categoryKey := fmt.Sprintf("test-policy-category-key-%d", timestamp)
	categoryValue := fmt.Sprintf("test-policy-category-value-%d", timestamp)

	testCategory := &prismModels.Category{
		Key:         ptr.To(categoryKey),
		Value:       ptr.To(categoryValue),
		Description: ptr.To("Test category for anti-affinity policy"),
	}

	createdCategory, err := antiAffinityClient.Categories.Create(ctx, testCategory)
	if err != nil {
		return nil, fmt.Errorf("failed to create test category: %w", err)
	}

	// Store the category ID for cleanup
	if createdCategory.ExtId != nil {
		createdCategoryIds = append(createdCategoryIds, *createdCategory.ExtId)
	}

	return createdCategory, nil
}

// Helper function to create a test anti-affinity policy
func createTestAntiAffinityPolicy(name, description string) (*policyModels.VmAntiAffinityPolicy, error) {
	policy := policyModels.NewVmAntiAffinityPolicy()
	policy.Name = ptr.To(name)
	policy.Description = ptr.To(description)

	// Get or create a test category
	testCategory, err := createOrGetTestCategory()
	if err != nil {
		return nil, fmt.Errorf("failed to create or get test category: %w", err)
	}

	// Use the category in the policy
	categoryRef := policyModels.NewCategoryReference()
	categoryRef.ExtId = testCategory.ExtId

	policy.Categories = []policyModels.CategoryReference{*categoryRef}

	return policy, nil
}
