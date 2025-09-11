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

// TestAntiAffinityPolicy represents a test anti-affinity policy entity
type TestAntiAffinityPolicy struct {
	ExtId       *string
	Name        *string
	Description *string
}

// TestAntiAffinityPoliciesInterface tests the AntiAffinityPolicies interface using generated mocks
func TestAntiAffinityPoliciesInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAntiAffinityPolicies := mocks.NewMockAntiAffinityPolicies[TestAntiAffinityPolicy](ctrl)
	ctx := context.Background()

	// Test Get method
	t.Run("Get", func(t *testing.T) {
		expectedPolicy := &TestAntiAffinityPolicy{
			ExtId:       ptr.To("test-policy-1"),
			Name:        ptr.To("Test Policy"),
			Description: ptr.To("A test anti-affinity policy"),
		}

		mockAntiAffinityPolicies.EXPECT().
			Get(ctx, "test-policy-1").
			Return(expectedPolicy, nil)

		result, err := mockAntiAffinityPolicies.Get(ctx, "test-policy-1")
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicy, result)
	})

	// Test Get with error
	t.Run("Get with error", func(t *testing.T) {
		expectedError := errors.New("policy not found")

		mockAntiAffinityPolicies.EXPECT().
			Get(ctx, "non-existent-policy").
			Return(nil, expectedError)

		result, err := mockAntiAffinityPolicies.Get(ctx, "non-existent-policy")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test List method
	t.Run("List", func(t *testing.T) {
		expectedPolicies := []TestAntiAffinityPolicy{
			{ExtId: ptr.To("1"), Name: ptr.To("Policy 1"), Description: ptr.To("First policy")},
			{ExtId: ptr.To("2"), Name: ptr.To("Policy 2"), Description: ptr.To("Second policy")},
		}

		mockAntiAffinityPolicies.EXPECT().
			List(ctx, gomock.Any()).
			Return(expectedPolicies, nil)

		result, err := mockAntiAffinityPolicies.List(ctx, convergedclient.WithPage(1), convergedclient.WithLimit(10))
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicies, result)
	})

	// Test List with error
	t.Run("List with error", func(t *testing.T) {
		expectedError := errors.New("failed to list policies")

		mockAntiAffinityPolicies.EXPECT().
			List(ctx, gomock.Any()).
			Return(nil, expectedError)

		result, err := mockAntiAffinityPolicies.List(ctx)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test ListAll method
	t.Run("ListAll", func(t *testing.T) {
		expectedPolicies := []TestAntiAffinityPolicy{
			{ExtId: ptr.To("1"), Name: ptr.To("Policy 1"), Description: ptr.To("First policy")},
			{ExtId: ptr.To("2"), Name: ptr.To("Policy 2"), Description: ptr.To("Second policy")},
			{ExtId: ptr.To("3"), Name: ptr.To("Policy 3"), Description: ptr.To("Third policy")},
		}

		mockAntiAffinityPolicies.EXPECT().
			List(ctx, gomock.Any()).
			Return(expectedPolicies, nil)

		result, err := mockAntiAffinityPolicies.List(ctx, convergedclient.WithFilter("name eq 'test'"))
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicies, result)
	})

	// Test ListAll with error
	t.Run("ListAll with error", func(t *testing.T) {
		expectedError := errors.New("failed to list all policies")

		mockAntiAffinityPolicies.EXPECT().
			List(ctx, gomock.Any()).
			Return(nil, expectedError)

		result, err := mockAntiAffinityPolicies.List(ctx)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test NewIterator method
	t.Run("NewIterator", func(t *testing.T) {
		// Create a mock iterator
		mockIterator := func(yield func(TestAntiAffinityPolicy, error) bool) {
			policies := []TestAntiAffinityPolicy{
				{ExtId: ptr.To("1"), Name: ptr.To("Policy 1")},
				{ExtId: ptr.To("2"), Name: ptr.To("Policy 2")},
			}
			for _, policy := range policies {
				if !yield(policy, nil) {
					return
				}
			}
		}

		mockAntiAffinityPolicies.EXPECT().
			NewIterator(gomock.Any(), gomock.Any()).
			Return(convergedclient.Iterator[TestAntiAffinityPolicy](mockIterator))

		iterator := mockAntiAffinityPolicies.NewIterator(ctx, convergedclient.WithPage(1))
		assert.NotNil(t, iterator)

		// Test iteration
		var results []TestAntiAffinityPolicy
		for policy, err := range iterator {
			assert.NoError(t, err)
			results = append(results, policy)
		}
		assert.Len(t, results, 2)
	})

	// Test Create method
	t.Run("Create", func(t *testing.T) {
		newPolicy := &TestAntiAffinityPolicy{
			Name:        ptr.To("New Policy"),
			Description: ptr.To("A newly created policy"),
		}

		expectedPolicy := &TestAntiAffinityPolicy{
			ExtId:       ptr.To("new-policy-id"),
			Name:        ptr.To("New Policy"),
			Description: ptr.To("A newly created policy"),
		}

		mockAntiAffinityPolicies.EXPECT().
			Create(ctx, newPolicy).
			Return(expectedPolicy, nil)

		result, err := mockAntiAffinityPolicies.Create(ctx, newPolicy)
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicy, result)
	})

	// Test Create with error
	t.Run("Create with error", func(t *testing.T) {
		newPolicy := &TestAntiAffinityPolicy{
			Name: ptr.To("Invalid Policy"),
		}

		expectedError := errors.New("invalid policy data")

		mockAntiAffinityPolicies.EXPECT().
			Create(ctx, newPolicy).
			Return(nil, expectedError)

		result, err := mockAntiAffinityPolicies.Create(ctx, newPolicy)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test Update method
	t.Run("Update", func(t *testing.T) {
		updatedPolicy := &TestAntiAffinityPolicy{
			ExtId:       ptr.To("existing-policy-id"),
			Name:        ptr.To("Updated Policy"),
			Description: ptr.To("An updated policy"),
		}

		mockAntiAffinityPolicies.EXPECT().
			Update(ctx, "existing-policy-id", updatedPolicy).
			Return(updatedPolicy, nil)

		result, err := mockAntiAffinityPolicies.Update(ctx, "existing-policy-id", updatedPolicy)
		assert.NoError(t, err)
		assert.Equal(t, updatedPolicy, result)
	})

	// Test Update with error
	t.Run("Update with error", func(t *testing.T) {
		updatedPolicy := &TestAntiAffinityPolicy{
			ExtId: ptr.To("non-existent-id"),
			Name:  ptr.To("Updated Policy"),
		}

		expectedError := errors.New("policy not found")

		mockAntiAffinityPolicies.EXPECT().
			Update(ctx, "non-existent-id", updatedPolicy).
			Return(nil, expectedError)

		result, err := mockAntiAffinityPolicies.Update(ctx, "non-existent-id", updatedPolicy)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test Delete method
	t.Run("Delete", func(t *testing.T) {
		mockAntiAffinityPolicies.EXPECT().
			Delete(ctx, "policy-to-delete").
			Return(nil)

		err := mockAntiAffinityPolicies.Delete(ctx, "policy-to-delete")
		assert.NoError(t, err)
	})

	// Test Delete with error
	t.Run("Delete with error", func(t *testing.T) {
		expectedError := errors.New("policy not found")

		mockAntiAffinityPolicies.EXPECT().
			Delete(ctx, "non-existent-policy").
			Return(expectedError)

		err := mockAntiAffinityPolicies.Delete(ctx, "non-existent-policy")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	// Test CreateAsync method
	t.Run("CreateAsync", func(t *testing.T) {
		newPolicy := &TestAntiAffinityPolicy{
			Name:        ptr.To("Async Policy"),
			Description: ptr.To("An async created policy"),
		}

		// Create a mock operation
		mockOperation := mocks.NewMockOperation[TestAntiAffinityPolicy](ctrl)
		expectedResult := []*TestAntiAffinityPolicy{
			{ExtId: ptr.To("async-policy-id"), Name: ptr.To("Async Policy"), Description: ptr.To("An async created policy")},
		}
		mockOperation.EXPECT().Wait(ctx).Return(expectedResult, nil)

		mockAntiAffinityPolicies.EXPECT().
			CreateAsync(ctx, newPolicy).
			Return(mockOperation, nil)

		operation, err := mockAntiAffinityPolicies.CreateAsync(ctx, newPolicy)
		assert.NoError(t, err)
		assert.NotNil(t, operation)

		// Test operation execution
		results, err := operation.Wait(ctx)
		assert.NoError(t, err)
		assert.Len(t, results, 1)
		assert.Equal(t, "async-policy-id", *results[0].ExtId)
	})

	// Test UpdateAsync method
	t.Run("UpdateAsync", func(t *testing.T) {
		updatedPolicy := &TestAntiAffinityPolicy{
			ExtId:       ptr.To("existing-policy-id"),
			Name:        ptr.To("Async Updated Policy"),
			Description: ptr.To("An async updated policy"),
		}

		// Create a mock operation
		mockOperation := mocks.NewMockOperation[TestAntiAffinityPolicy](ctrl)
		expectedResult := []*TestAntiAffinityPolicy{updatedPolicy}
		mockOperation.EXPECT().Wait(ctx).Return(expectedResult, nil)

		mockAntiAffinityPolicies.EXPECT().
			UpdateAsync(ctx, "existing-policy-id", updatedPolicy).
			Return(mockOperation, nil)

		operation, err := mockAntiAffinityPolicies.UpdateAsync(ctx, "existing-policy-id", updatedPolicy)
		assert.NoError(t, err)
		assert.NotNil(t, operation)

		// Test operation execution
		results, err := operation.Wait(ctx)
		assert.NoError(t, err)
		assert.Len(t, results, 1)
		assert.Equal(t, updatedPolicy, results[0])
	})

	// Test DeleteAsync method
	t.Run("DeleteAsync", func(t *testing.T) {
		// Create a mock operation
		mockOperation := mocks.NewMockOperation[convergedclient.NoEntity](ctrl)
		var noEntity convergedclient.NoEntity
		expectedResult := []*convergedclient.NoEntity{&noEntity}
		mockOperation.EXPECT().Wait(ctx).Return(expectedResult, nil)

		mockAntiAffinityPolicies.EXPECT().
			DeleteAsync(ctx, "policy-to-delete").
			Return(mockOperation, nil)

		operation, err := mockAntiAffinityPolicies.DeleteAsync(ctx, "policy-to-delete")
		assert.NoError(t, err)
		assert.NotNil(t, operation)

		// Test operation execution
		results, err := operation.Wait(ctx)
		assert.NoError(t, err)
		assert.Len(t, results, 1)
	})
}

// TestAntiAffinityPoliciesInterfaceComposition tests that AntiAffinityPolicies interface properly composes other interfaces
func TestAntiAffinityPoliciesInterfaceComposition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAntiAffinityPolicies := mocks.NewMockAntiAffinityPolicies[TestAntiAffinityPolicy](ctrl)
	ctx := context.Background()

	// Test that AntiAffinityPolicies can be used as Getter
	t.Run("as Getter", func(t *testing.T) {
		var getter convergedclient.Getter[TestAntiAffinityPolicy] = mockAntiAffinityPolicies

		expectedPolicy := &TestAntiAffinityPolicy{ExtId: ptr.To("1"), Name: ptr.To("Test")}
		mockAntiAffinityPolicies.EXPECT().Get(ctx, "1").Return(expectedPolicy, nil)

		result, err := getter.Get(ctx, "1")
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicy, result)
	})

	// Test that AntiAffinityPolicies can be used as Lister
	t.Run("as Lister", func(t *testing.T) {
		var lister convergedclient.Lister[TestAntiAffinityPolicy] = mockAntiAffinityPolicies

		expectedPolicies := []TestAntiAffinityPolicy{{ExtId: ptr.To("1"), Name: ptr.To("Test")}}
		mockAntiAffinityPolicies.EXPECT().List(ctx).Return(expectedPolicies, nil)

		result, err := lister.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicies, result)
	})

	// Test that AntiAffinityPolicies can be used as Creator
	t.Run("as Creator", func(t *testing.T) {
		var creator convergedclient.Creator[TestAntiAffinityPolicy] = mockAntiAffinityPolicies

		newPolicy := &TestAntiAffinityPolicy{Name: ptr.To("New Policy")}
		expectedPolicy := &TestAntiAffinityPolicy{ExtId: ptr.To("1"), Name: ptr.To("New Policy")}
		mockAntiAffinityPolicies.EXPECT().Create(ctx, newPolicy).Return(expectedPolicy, nil)

		result, err := creator.Create(ctx, newPolicy)
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicy, result)
	})

	// Test that AntiAffinityPolicies can be used as Updater
	t.Run("as Updater", func(t *testing.T) {
		var updater convergedclient.Updater[TestAntiAffinityPolicy] = mockAntiAffinityPolicies

		updatedPolicy := &TestAntiAffinityPolicy{ExtId: ptr.To("1"), Name: ptr.To("Updated")}
		mockAntiAffinityPolicies.EXPECT().Update(ctx, "1", updatedPolicy).Return(updatedPolicy, nil)

		result, err := updater.Update(ctx, "1", updatedPolicy)
		assert.NoError(t, err)
		assert.Equal(t, updatedPolicy, result)
	})

	// Test that AntiAffinityPolicies can be used as Deleter
	t.Run("as Deleter", func(t *testing.T) {
		var deleter convergedclient.Deleter[TestAntiAffinityPolicy] = mockAntiAffinityPolicies

		mockAntiAffinityPolicies.EXPECT().Delete(ctx, "1").Return(nil)

		err := deleter.Delete(ctx, "1")
		assert.NoError(t, err)
	})

	// Test that AntiAffinityPolicies can be used as AsyncCreator
	t.Run("as AsyncCreator", func(t *testing.T) {
		var asyncCreator convergedclient.AsyncCreator[TestAntiAffinityPolicy] = mockAntiAffinityPolicies

		newPolicy := &TestAntiAffinityPolicy{Name: ptr.To("Async Policy")}
		mockOperation := mocks.NewMockOperation[TestAntiAffinityPolicy](ctrl)
		mockAntiAffinityPolicies.EXPECT().CreateAsync(ctx, newPolicy).Return(mockOperation, nil)

		operation, err := asyncCreator.CreateAsync(ctx, newPolicy)
		assert.NoError(t, err)
		assert.NotNil(t, operation)
	})

	// Test that AntiAffinityPolicies can be used as AsyncUpdater
	t.Run("as AsyncUpdater", func(t *testing.T) {
		var asyncUpdater convergedclient.AsyncUpdater[TestAntiAffinityPolicy] = mockAntiAffinityPolicies

		updatedPolicy := &TestAntiAffinityPolicy{ExtId: ptr.To("1"), Name: ptr.To("Async Updated")}
		mockOperation := mocks.NewMockOperation[TestAntiAffinityPolicy](ctrl)
		mockAntiAffinityPolicies.EXPECT().UpdateAsync(ctx, "1", updatedPolicy).Return(mockOperation, nil)

		operation, err := asyncUpdater.UpdateAsync(ctx, "1", updatedPolicy)
		assert.NoError(t, err)
		assert.NotNil(t, operation)
	})

	// Test that AntiAffinityPolicies can be used as AsyncDeleter
	t.Run("as AsyncDeleter", func(t *testing.T) {
		var asyncDeleter convergedclient.AsyncDeleter[convergedclient.NoEntity] = mockAntiAffinityPolicies

		mockOperation := mocks.NewMockOperation[convergedclient.NoEntity](ctrl)
		mockAntiAffinityPolicies.EXPECT().DeleteAsync(ctx, "1").Return(mockOperation, nil)

		operation, err := asyncDeleter.DeleteAsync(ctx, "1")
		assert.NoError(t, err)
		assert.NotNil(t, operation)
	})
}

// TestAntiAffinityPoliciesErrorHandling tests various error scenarios
func TestAntiAffinityPoliciesErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAntiAffinityPolicies := mocks.NewMockAntiAffinityPolicies[TestAntiAffinityPolicy](ctrl)
	ctx := context.Background()

	// Test context cancellation
	t.Run("context cancellation", func(t *testing.T) {
		cancelledCtx, cancel := context.WithCancel(ctx)
		cancel()

		expectedError := context.Canceled
		mockAntiAffinityPolicies.EXPECT().Get(cancelledCtx, "test").Return(nil, expectedError)

		_, err := mockAntiAffinityPolicies.Get(cancelledCtx, "test")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	// Test context timeout
	t.Run("context timeout", func(t *testing.T) {
		timeoutCtx, cancel := context.WithTimeout(ctx, 0) // Immediate timeout
		defer cancel()

		expectedError := context.DeadlineExceeded
		mockAntiAffinityPolicies.EXPECT().List(timeoutCtx).Return(nil, expectedError)

		_, err := mockAntiAffinityPolicies.List(timeoutCtx)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	// Test invalid input
	t.Run("invalid input", func(t *testing.T) {
		expectedError := errors.New("invalid UUID format")
		mockAntiAffinityPolicies.EXPECT().Get(ctx, "").Return(nil, expectedError)

		_, err := mockAntiAffinityPolicies.Get(ctx, "")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	// Test async operation error
	t.Run("async operation error", func(t *testing.T) {
		newPolicy := &TestAntiAffinityPolicy{Name: ptr.To("Failed Policy")}
		expectedError := errors.New("async operation failed")

		mockAntiAffinityPolicies.EXPECT().CreateAsync(ctx, newPolicy).Return(nil, expectedError)

		operation, err := mockAntiAffinityPolicies.CreateAsync(ctx, newPolicy)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, operation)
	})
}

// TestAntiAffinityPoliciesWithODataOptions tests AntiAffinityPolicies interface with various OData options
func TestAntiAffinityPoliciesWithODataOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAntiAffinityPolicies := mocks.NewMockAntiAffinityPolicies[TestAntiAffinityPolicy](ctrl)
	ctx := context.Background()

	// Test with page option
	t.Run("with page option", func(t *testing.T) {
		expectedPolicies := []TestAntiAffinityPolicy{{ExtId: ptr.To("1"), Name: ptr.To("Page 1")}}
		mockAntiAffinityPolicies.EXPECT().List(ctx, gomock.Any()).Return(expectedPolicies, nil)

		result, err := mockAntiAffinityPolicies.List(ctx, convergedclient.WithPage(1))
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicies, result)
	})

	// Test with limit option
	t.Run("with limit option", func(t *testing.T) {
		expectedPolicies := []TestAntiAffinityPolicy{{ExtId: ptr.To("1"), Name: ptr.To("Limited")}}
		mockAntiAffinityPolicies.EXPECT().List(ctx, gomock.Any()).Return(expectedPolicies, nil)

		result, err := mockAntiAffinityPolicies.List(ctx, convergedclient.WithLimit(5))
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicies, result)
	})

	// Test with filter option
	t.Run("with filter option", func(t *testing.T) {
		expectedPolicies := []TestAntiAffinityPolicy{{ExtId: ptr.To("1"), Name: ptr.To("Filtered")}}
		mockAntiAffinityPolicies.EXPECT().List(ctx, gomock.Any()).Return(expectedPolicies, nil)

		result, err := mockAntiAffinityPolicies.List(ctx, convergedclient.WithFilter("name eq 'test'"))
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicies, result)
	})

	// Test with order by option
	t.Run("with order by option", func(t *testing.T) {
		expectedPolicies := []TestAntiAffinityPolicy{{ExtId: ptr.To("1"), Name: ptr.To("Ordered")}}
		mockAntiAffinityPolicies.EXPECT().List(ctx, gomock.Any()).Return(expectedPolicies, nil)

		result, err := mockAntiAffinityPolicies.List(ctx, convergedclient.WithOrderBy("name asc"))
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicies, result)
	})

	// Test with expand option
	t.Run("with expand option", func(t *testing.T) {
		expectedPolicies := []TestAntiAffinityPolicy{{ExtId: ptr.To("1"), Name: ptr.To("Expanded")}}
		mockAntiAffinityPolicies.EXPECT().List(ctx, gomock.Any()).Return(expectedPolicies, nil)

		result, err := mockAntiAffinityPolicies.List(ctx, convergedclient.WithExpand("metadata"))
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicies, result)
	})

	// Test with select option
	t.Run("with select option", func(t *testing.T) {
		expectedPolicies := []TestAntiAffinityPolicy{{ExtId: ptr.To("1"), Name: ptr.To("Selected")}}
		mockAntiAffinityPolicies.EXPECT().List(ctx, gomock.Any()).Return(expectedPolicies, nil)

		result, err := mockAntiAffinityPolicies.List(ctx, convergedclient.WithSelect("id,name"))
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicies, result)
	})

	// Test with apply option
	t.Run("with apply option", func(t *testing.T) {
		expectedPolicies := []TestAntiAffinityPolicy{{ExtId: ptr.To("1"), Name: ptr.To("Applied")}}
		mockAntiAffinityPolicies.EXPECT().List(ctx, gomock.Any()).Return(expectedPolicies, nil)

		result, err := mockAntiAffinityPolicies.List(ctx, convergedclient.WithApply("groupby((name))"))
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicies, result)
	})

	// Test with multiple options
	t.Run("with multiple options", func(t *testing.T) {
		expectedPolicies := []TestAntiAffinityPolicy{{ExtId: ptr.To("1"), Name: ptr.To("Multiple")}}
		mockAntiAffinityPolicies.EXPECT().List(ctx, gomock.Any()).Return(expectedPolicies, nil)

		result, err := mockAntiAffinityPolicies.List(ctx, convergedclient.WithPage(1), convergedclient.WithLimit(10), convergedclient.WithFilter("name eq 'test'"), convergedclient.WithOrderBy("name asc"))
		assert.NoError(t, err)
		assert.Equal(t, expectedPolicies, result)
	})

	// Test NewIterator with options
	t.Run("NewIterator with options", func(t *testing.T) {
		// Create a mock iterator
		mockIterator := func(yield func(TestAntiAffinityPolicy, error) bool) {
			policies := []TestAntiAffinityPolicy{
				{ExtId: ptr.To("1"), Name: ptr.To("Iterator Policy 1")},
				{ExtId: ptr.To("2"), Name: ptr.To("Iterator Policy 2")},
			}
			for _, policy := range policies {
				if !yield(policy, nil) {
					return
				}
			}
		}

		mockAntiAffinityPolicies.EXPECT().
			NewIterator(ctx, gomock.Any()).
			Return(convergedclient.Iterator[TestAntiAffinityPolicy](mockIterator))

		iterator := mockAntiAffinityPolicies.NewIterator(ctx, convergedclient.WithLimit(5), convergedclient.WithOrderBy("name"))
		assert.NotNil(t, iterator)

		// Test iteration
		var results []TestAntiAffinityPolicy
		for policy, err := range iterator {
			assert.NoError(t, err)
			results = append(results, policy)
		}
		assert.Len(t, results, 2)
	})
}
