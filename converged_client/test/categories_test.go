package test

import (
	"context"
	"errors"
	"testing"

	convergedclient "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	"github.com/nutanix-cloud-native/prism-go-client/converged_client/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"k8s.io/utils/ptr"
)

// TestCategory represents a test category entity
type TestCategory struct {
	ExtId *string
	Key   *string
	Value *string
}

// TestCategoriesInterface tests the Categories interface using generated mocks
func TestCategoriesInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategories := mocks.NewMockCategories[TestCategory](ctrl)
	ctx := context.Background()

	// Test Get method
	t.Run("Get", func(t *testing.T) {
		expectedCategory := &TestCategory{
			ExtId: ptr.To("test-category-1"),
			Key:   ptr.To("Test Category"),
			Value: ptr.To("A test category"),
		}

		mockCategories.EXPECT().
			Get(ctx, "test-category-1").
			Return(expectedCategory, nil)

		result, err := mockCategories.Get(ctx, "test-category-1")
		assert.NoError(t, err)
		assert.Equal(t, expectedCategory, result)
	})

	// Test Get with error
	t.Run("Get with error", func(t *testing.T) {
		expectedError := errors.New("category not found")

		mockCategories.EXPECT().
			Get(ctx, "non-existent-category").
			Return(nil, expectedError)

		result, err := mockCategories.Get(ctx, "non-existent-category")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test List method
	t.Run("List", func(t *testing.T) {
		expectedCategories := []TestCategory{
			{ExtId: ptr.To("1"), Key: ptr.To("Category 1"), Value: ptr.To("First category")},
			{ExtId: ptr.To("2"), Key: ptr.To("Category 2"), Value: ptr.To("Second category")},
		}

		mockCategories.EXPECT().
			List(ctx, gomock.Any()).
			Return(expectedCategories, nil)

		result, err := mockCategories.List(ctx, convergedclient.WithPage(1), convergedclient.WithLimit(10))
		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	// Test List with error
	t.Run("List with error", func(t *testing.T) {
		expectedError := errors.New("failed to list categories")

		mockCategories.EXPECT().
			List(ctx, gomock.Any()).
			Return(nil, expectedError)

		result, err := mockCategories.List(ctx)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test ListAll method
	t.Run("ListAll", func(t *testing.T) {
		expectedCategories := []TestCategory{
			{ExtId: ptr.To("1"), Key: ptr.To("Category 1"), Value: ptr.To("First category")},
			{ExtId: ptr.To("2"), Key: ptr.To("Category 2"), Value: ptr.To("Second category")},
			{ExtId: ptr.To("3"), Key: ptr.To("Category 3"), Value: ptr.To("Third category")},
		}

		mockCategories.EXPECT().
			ListAll(ctx, gomock.Any()).
			Return(expectedCategories, nil)

		result, err := mockCategories.ListAll(ctx, convergedclient.WithFilter("name eq 'test'"))
		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	// Test ListAll with error
	t.Run("ListAll with error", func(t *testing.T) {
		expectedError := errors.New("failed to list all categories")

		mockCategories.EXPECT().
			ListAll(ctx, gomock.Any()).
			Return(nil, expectedError)

		result, err := mockCategories.ListAll(ctx)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test NewIterator method
	t.Run("NewIterator", func(t *testing.T) {
		// Create a mock iterator
		mockIterator := func(yield func(TestCategory, error) bool) {
			categories := []TestCategory{
				{ExtId: ptr.To("1"), Key: ptr.To("Category 1")},
				{ExtId: ptr.To("2"), Key: ptr.To("Category 2")},
			}
			for _, cat := range categories {
				if !yield(cat, nil) {
					return
				}
			}
		}

		mockCategories.EXPECT().
			NewIterator(gomock.Any()).
			Return(convergedclient.Iterator[TestCategory](mockIterator))

		iterator := mockCategories.NewIterator(convergedclient.WithPage(1))
		assert.NotNil(t, iterator)

		// Test iteration
		var results []TestCategory
		for category, err := range iterator {
			assert.NoError(t, err)
			results = append(results, category)
		}
		assert.Len(t, results, 2)
	})

	// Test Create method
	t.Run("Create", func(t *testing.T) {
		newCategory := &TestCategory{
			Key:   ptr.To("New Category"),
			Value: ptr.To("A newly created category"),
		}

		expectedCategory := &TestCategory{
			ExtId: ptr.To("new-category-id"),
			Key:   ptr.To("New Category"),
			Value: ptr.To("A newly created category"),
		}

		mockCategories.EXPECT().
			Create(ctx, newCategory).
			Return(expectedCategory, nil)

		result, err := mockCategories.Create(ctx, newCategory)
		assert.NoError(t, err)
		assert.Equal(t, expectedCategory, result)
	})

	// Test Create with error
	t.Run("Create with error", func(t *testing.T) {
		newCategory := &TestCategory{
			Key: ptr.To("Invalid Category"),
		}

		expectedError := errors.New("invalid category data")

		mockCategories.EXPECT().
			Create(ctx, newCategory).
			Return(nil, expectedError)

		result, err := mockCategories.Create(ctx, newCategory)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test Update method
	t.Run("Update", func(t *testing.T) {
		updatedCategory := &TestCategory{
			ExtId: ptr.To("existing-category-id"),
			Key:   ptr.To("Updated Category"),
			Value: ptr.To("An updated category"),
		}

		mockCategories.EXPECT().
			Update(ctx, "existing-category-id", updatedCategory).
			Return(updatedCategory, nil)

		result, err := mockCategories.Update(ctx, "existing-category-id", updatedCategory)
		assert.NoError(t, err)
		assert.Equal(t, updatedCategory, result)
	})

	// Test Update with error
	t.Run("Update with error", func(t *testing.T) {
		updatedCategory := &TestCategory{
			ExtId: ptr.To("non-existent-id"),
			Key:   ptr.To("Updated Category"),
		}

		expectedError := errors.New("category not found")

		mockCategories.EXPECT().
			Update(ctx, "non-existent-id", updatedCategory).
			Return(nil, expectedError)

		result, err := mockCategories.Update(ctx, "non-existent-id", updatedCategory)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test Delete method
	t.Run("Delete", func(t *testing.T) {
		deletedCategory := &TestCategory{
			ExtId: ptr.To("category-to-delete"),
			Key:   ptr.To("Category to Delete"),
			Value: ptr.To("This category will be deleted"),
		}

		mockCategories.EXPECT().
			Delete(ctx, "category-to-delete").
			Return(deletedCategory, nil)

		result, err := mockCategories.Delete(ctx, "category-to-delete")
		assert.NoError(t, err)
		assert.Equal(t, deletedCategory, result)
	})

	// Test Delete with error
	t.Run("Delete with error", func(t *testing.T) {
		expectedError := errors.New("category not found")

		mockCategories.EXPECT().
			Delete(ctx, "non-existent-category").
			Return(nil, expectedError)

		result, err := mockCategories.Delete(ctx, "non-existent-category")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})
}

// TestCategoriesInterfaceComposition tests that Categories interface properly composes other interfaces
func TestCategoriesInterfaceComposition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategories := mocks.NewMockCategories[TestCategory](ctrl)
	ctx := context.Background()

	// Test that Categories can be used as Getter
	t.Run("as Getter", func(t *testing.T) {
		var getter convergedclient.Getter[TestCategory] = mockCategories

		expectedCategory := &TestCategory{ExtId: ptr.To("1"), Key: ptr.To("Test")}
		mockCategories.EXPECT().Get(ctx, "1").Return(expectedCategory, nil)

		result, err := getter.Get(ctx, "1")
		assert.NoError(t, err)
		assert.Equal(t, expectedCategory, result)
	})

	// Test that Categories can be used as Lister
	t.Run("as Lister", func(t *testing.T) {
		var lister convergedclient.Lister[TestCategory] = mockCategories

		expectedCategories := []TestCategory{{ExtId: ptr.To("1"), Key: ptr.To("Test")}}
		mockCategories.EXPECT().List(ctx).Return(expectedCategories, nil)

		result, err := lister.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	// Test that Categories can be used as Creator
	t.Run("as Creator", func(t *testing.T) {
		var creator convergedclient.Creator[TestCategory] = mockCategories

		newCategory := &TestCategory{Key: ptr.To("New Category")}
		expectedCategory := &TestCategory{ExtId: ptr.To("1"), Key: ptr.To("New Category")}
		mockCategories.EXPECT().Create(ctx, newCategory).Return(expectedCategory, nil)

		result, err := creator.Create(ctx, newCategory)
		assert.NoError(t, err)
		assert.Equal(t, expectedCategory, result)
	})

	// Test that Categories can be used as Updater
	t.Run("as Updater", func(t *testing.T) {
		var updater convergedclient.Updater[TestCategory] = mockCategories

		updatedCategory := &TestCategory{ExtId: ptr.To("1"), Key: ptr.To("Updated")}
		mockCategories.EXPECT().Update(ctx, "1", updatedCategory).Return(updatedCategory, nil)

		result, err := updater.Update(ctx, "1", updatedCategory)
		assert.NoError(t, err)
		assert.Equal(t, updatedCategory, result)
	})

	// Test that Categories can be used as Deleter
	t.Run("as Deleter", func(t *testing.T) {
		var deleter convergedclient.Deleter[TestCategory] = mockCategories

		deletedCategory := &TestCategory{ExtId: ptr.To("1"), Key: ptr.To("Deleted")}
		mockCategories.EXPECT().Delete(ctx, "1").Return(deletedCategory, nil)

		result, err := deleter.Delete(ctx, "1")
		assert.NoError(t, err)
		assert.Equal(t, deletedCategory, result)
	})
}

// TestCategoriesErrorHandling tests various error scenarios
func TestCategoriesErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategories := mocks.NewMockCategories[TestCategory](ctrl)
	ctx := context.Background()

	// Test context cancellation
	t.Run("context cancellation", func(t *testing.T) {
		cancelledCtx, cancel := context.WithCancel(ctx)
		cancel()

		expectedError := context.Canceled
		mockCategories.EXPECT().Get(cancelledCtx, "test").Return(nil, expectedError)

		_, err := mockCategories.Get(cancelledCtx, "test")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	// Test context timeout
	t.Run("context timeout", func(t *testing.T) {
		timeoutCtx, cancel := context.WithTimeout(ctx, 0) // Immediate timeout
		defer cancel()

		expectedError := context.DeadlineExceeded
		mockCategories.EXPECT().List(timeoutCtx).Return(nil, expectedError)

		_, err := mockCategories.List(timeoutCtx)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	// Test invalid input
	t.Run("invalid input", func(t *testing.T) {
		expectedError := errors.New("invalid UUID format")
		mockCategories.EXPECT().Get(ctx, "").Return(nil, expectedError)

		_, err := mockCategories.Get(ctx, "")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

// TestCategoriesWithODataOptions tests Categories interface with various OData options
func TestCategoriesWithODataOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategories := mocks.NewMockCategories[TestCategory](ctrl)
	ctx := context.Background()

	// Test with page option
	t.Run("with page option", func(t *testing.T) {
		expectedCategories := []TestCategory{{ExtId: ptr.To("1"), Key: ptr.To("Page 1")}}
		mockCategories.EXPECT().List(ctx, gomock.Any()).Return(expectedCategories, nil)

		result, err := mockCategories.List(ctx, convergedclient.WithPage(1))
		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	// Test with limit option
	t.Run("with limit option", func(t *testing.T) {
		expectedCategories := []TestCategory{{ExtId: ptr.To("1"), Key: ptr.To("Limited")}}
		mockCategories.EXPECT().List(ctx, gomock.Any()).Return(expectedCategories, nil)

		result, err := mockCategories.List(ctx, convergedclient.WithLimit(5))
		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	// Test with filter option
	t.Run("with filter option", func(t *testing.T) {
		expectedCategories := []TestCategory{{ExtId: ptr.To("1"), Key: ptr.To("Filtered")}}
		mockCategories.EXPECT().List(ctx, gomock.Any()).Return(expectedCategories, nil)

		result, err := mockCategories.List(ctx, convergedclient.WithFilter("name eq 'test'"))
		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	// Test with order by option
	t.Run("with order by option", func(t *testing.T) {
		expectedCategories := []TestCategory{{ExtId: ptr.To("1"), Key: ptr.To("Ordered")}}
		mockCategories.EXPECT().List(ctx, gomock.Any()).Return(expectedCategories, nil)

		result, err := mockCategories.List(ctx, convergedclient.WithOrderBy("name asc"))
		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	// Test with expand option
	t.Run("with expand option", func(t *testing.T) {
		expectedCategories := []TestCategory{{ExtId: ptr.To("1"), Key: ptr.To("Expanded")}}
		mockCategories.EXPECT().List(ctx, gomock.Any()).Return(expectedCategories, nil)

		result, err := mockCategories.List(ctx, convergedclient.WithExpand("metadata"))
		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	// Test with select option
	t.Run("with select option", func(t *testing.T) {
		expectedCategories := []TestCategory{{ExtId: ptr.To("1"), Key: ptr.To("Selected")}}
		mockCategories.EXPECT().List(ctx, gomock.Any()).Return(expectedCategories, nil)

		result, err := mockCategories.List(ctx, convergedclient.WithSelect("id,name"))
		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	// Test with apply option
	t.Run("with apply option", func(t *testing.T) {
		expectedCategories := []TestCategory{{ExtId: ptr.To("1"), Key: ptr.To("Applied")}}
		mockCategories.EXPECT().List(ctx, gomock.Any()).Return(expectedCategories, nil)

		result, err := mockCategories.List(ctx, convergedclient.WithApply("groupby((name))"))
		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	// Test with multiple options
	t.Run("with multiple options", func(t *testing.T) {
		expectedCategories := []TestCategory{{ExtId: ptr.To("1"), Key: ptr.To("Multiple")}}
		mockCategories.EXPECT().List(ctx, gomock.Any()).Return(expectedCategories, nil)

		result, err := mockCategories.List(ctx, convergedclient.WithPage(1), convergedclient.WithLimit(10), convergedclient.WithFilter("name eq 'test'"), convergedclient.WithOrderBy("name asc"))
		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})
}
