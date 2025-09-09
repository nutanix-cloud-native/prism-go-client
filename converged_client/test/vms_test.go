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

// Test types for VMs interface testing
type TestVM struct {
	ExtId *string `json:"id"`
	Name  *string `json:"name"`
}

// Test VMs interface implementation with generated mocks
func TestVMsInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVMs := mocks.NewMockVMs[TestVM](ctrl)
	mockOperation := mocks.NewMockOperation[TestVM](ctrl)
	mockOperationDelete := mocks.NewMockOperation[convergedclient.NoEntity](ctrl)
	ctx := context.Background()
	uuid := "test-uuid"
	expectedVM := &TestVM{ExtId: ptr.To(uuid), Name: ptr.To("test-vm")}

	// Test Get
	mockVMs.EXPECT().Get(ctx, uuid).Return(expectedVM, nil)
	result, err := mockVMs.Get(ctx, uuid)
	assert.NoError(t, err)
	assert.Equal(t, expectedVM, result)

	// Test List
	expectedVMs := []TestVM{*expectedVM}
	mockVMs.EXPECT().List(ctx, gomock.Any()).Return(expectedVMs, nil)
	resultList, err := mockVMs.List(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedVMs, resultList)

	// Test NewIterator
	iterator := func(yield func(TestVM, error) bool) {
		yield(*expectedVM, nil)
	}
	mockVMs.EXPECT().NewIterator(gomock.Any(), gomock.Any()).Return(iterator)
	resultIterator := mockVMs.NewIterator(ctx)
	count := 0
	for entity, err := range resultIterator {
		assert.NoError(t, err)
		assert.NotNil(t, entity)
		count++
	}
	assert.Equal(t, 1, count)

	// Test Create
	mockVMs.EXPECT().Create(ctx, expectedVM).Return(expectedVM, nil)
	resultCreate, err := mockVMs.Create(ctx, expectedVM)
	assert.NoError(t, err)
	assert.Equal(t, expectedVM, resultCreate)

	// Test Update
	mockVMs.EXPECT().Update(ctx, uuid, expectedVM).Return(expectedVM, nil)
	resultUpdate, err := mockVMs.Update(ctx, uuid, expectedVM)
	assert.NoError(t, err)
	assert.Equal(t, expectedVM, resultUpdate)

	// Test Delete
	mockVMs.EXPECT().Delete(ctx, uuid).Return(nil)
	err = mockVMs.Delete(ctx, uuid)
	assert.NoError(t, err)

	// Test CreateAsync
	mockVMs.EXPECT().CreateAsync(ctx, expectedVM).Return(mockOperation, nil)
	operation, err := mockVMs.CreateAsync(ctx, expectedVM)
	assert.NoError(t, err)
	assert.Equal(t, mockOperation, operation)

	// Test UpdateAsync
	mockVMs.EXPECT().UpdateAsync(ctx, uuid, expectedVM).Return(mockOperation, nil)
	operation, err = mockVMs.UpdateAsync(ctx, uuid, expectedVM)
	assert.NoError(t, err)
	assert.Equal(t, mockOperation, operation)

	// Test DeleteAsync
	mockVMs.EXPECT().DeleteAsync(ctx, uuid).Return(mockOperationDelete, nil)
	operationDelete, err := mockVMs.DeleteAsync(ctx, uuid)
	assert.NoError(t, err)
	assert.Equal(t, mockOperationDelete, operationDelete)

	// Test PowerOnVM
	mockVMs.EXPECT().PowerOnVM(uuid).Return(mockOperation, nil)
	operation, err = mockVMs.PowerOnVM(uuid)
	assert.NoError(t, err)
	assert.Equal(t, mockOperation, operation)

	// Test PowerOffVM
	mockVMs.EXPECT().PowerOffVM(uuid).Return(mockOperation, nil)
	operation, err = mockVMs.PowerOffVM(uuid)
	assert.NoError(t, err)
	assert.Equal(t, mockOperation, operation)
}

// Test VMs interface composition with generated mocks
func TestVMsInterfaceComposition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVMs := mocks.NewMockVMs[TestVM](ctrl)
	mockOperation := mocks.NewMockOperation[TestVM](ctrl)
	mockOperationDelete := mocks.NewMockOperation[convergedclient.NoEntity](ctrl)
	ctx := context.Background()
	expectedVM := &TestVM{ExtId: ptr.To("test-1"), Name: ptr.To("Test VM 1")}

	// Test that VMs implements all the required interfaces
	var getter convergedclient.Getter[TestVM] = mockVMs
	var lister convergedclient.Lister[TestVM] = mockVMs
	var creator convergedclient.Creator[TestVM] = mockVMs
	var updater convergedclient.Updater[TestVM] = mockVMs
	var deleter convergedclient.Deleter[TestVM] = mockVMs
	var asyncCreator convergedclient.AsyncCreator[TestVM] = mockVMs
	var asyncUpdater convergedclient.AsyncUpdater[TestVM] = mockVMs
	var asyncDeleter convergedclient.AsyncDeleter[TestVM] = mockVMs

	// Test each interface independently
	mockVMs.EXPECT().Get(ctx, "test-1").Return(expectedVM, nil)
	entity, err := getter.Get(ctx, "test-1")
	assert.NoError(t, err)
	assert.NotNil(t, entity)

	expectedVMs := []TestVM{*expectedVM}
	mockVMs.EXPECT().List(ctx, gomock.Any()).Return(expectedVMs, nil)
	entities, err := lister.List(ctx)
	assert.NoError(t, err)
	assert.Len(t, entities, 1)

	newVM := &TestVM{ExtId: ptr.To("test-2"), Name: ptr.To("Test VM 2")}
	mockVMs.EXPECT().Create(ctx, newVM).Return(newVM, nil)
	created, err := creator.Create(ctx, newVM)
	assert.NoError(t, err)
	assert.Equal(t, newVM, created)

	updatedVM := &TestVM{ExtId: ptr.To("test-1"), Name: ptr.To("Updated VM")}
	mockVMs.EXPECT().Update(ctx, "test-1", updatedVM).Return(updatedVM, nil)
	updated, err := updater.Update(ctx, "test-1", updatedVM)
	assert.NoError(t, err)
	assert.Equal(t, updatedVM, updated)

	mockVMs.EXPECT().Delete(ctx, "test-1").Return(nil)
	err = deleter.Delete(ctx, "test-1")
	assert.NoError(t, err)

	// Test async operations
	mockVMs.EXPECT().CreateAsync(ctx, gomock.Any()).Return(mockOperation, nil)
	_, err = asyncCreator.CreateAsync(ctx, &TestVM{})
	assert.NoError(t, err)

	mockVMs.EXPECT().UpdateAsync(ctx, "test-id", gomock.Any()).Return(mockOperation, nil)
	_, err = asyncUpdater.UpdateAsync(ctx, "test-id", &TestVM{})
	assert.NoError(t, err)

	mockVMs.EXPECT().DeleteAsync(ctx, "test-id").Return(mockOperationDelete, nil)
	_, err = asyncDeleter.DeleteAsync(ctx, "test-id")
	assert.NoError(t, err)
}

// Test VMs error handling with generated mocks
func TestVMsErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVMs := mocks.NewMockVMs[TestVM](ctrl)
	ctx := context.Background()
	testError := errors.New("test error")

	// Test error handling in Get
	mockVMs.EXPECT().Get(ctx, "test").Return(nil, testError)
	_, err := mockVMs.Get(ctx, "test")
	assert.Error(t, err)
	assert.Equal(t, testError, err)

	// Test error handling in List
	mockVMs.EXPECT().List(ctx, gomock.Any()).Return(nil, testError)
	_, err = mockVMs.List(ctx)
	assert.Error(t, err)
	assert.Equal(t, testError, err)

	// Test error handling in Create
	testVM := &TestVM{Name: ptr.To("test")}
	mockVMs.EXPECT().Create(ctx, testVM).Return(nil, testError)
	_, err = mockVMs.Create(ctx, testVM)
	assert.Error(t, err)
	assert.Equal(t, testError, err)

	// Test error handling in CreateAsync
	mockVMs.EXPECT().CreateAsync(ctx, testVM).Return(nil, testError)
	_, err = mockVMs.CreateAsync(ctx, testVM)
	assert.Error(t, err)
	assert.Equal(t, testError, err)

	// Test error handling in Update
	mockVMs.EXPECT().Update(ctx, "test", testVM).Return(nil, testError)
	_, err = mockVMs.Update(ctx, "test", testVM)
	assert.Error(t, err)
	assert.Equal(t, testError, err)

	// Test error handling in UpdateAsync
	mockVMs.EXPECT().UpdateAsync(ctx, "test", testVM).Return(nil, testError)
	_, err = mockVMs.UpdateAsync(ctx, "test", testVM)
	assert.Error(t, err)
	assert.Equal(t, testError, err)

	// Test error handling in Delete
	mockVMs.EXPECT().Delete(ctx, "test").Return(testError)
	err = mockVMs.Delete(ctx, "test")
	assert.Error(t, err)
	assert.Equal(t, testError, err)

	// Test error handling in DeleteAsync
	mockVMs.EXPECT().DeleteAsync(ctx, "test").Return(nil, testError)
	_, err = mockVMs.DeleteAsync(ctx, "test")
	assert.Error(t, err)
	assert.Equal(t, testError, err)

	// Test error handling in PowerOnVM
	mockVMs.EXPECT().PowerOnVM("test").Return(nil, testError)
	_, err = mockVMs.PowerOnVM("test")
	assert.Error(t, err)
	assert.Equal(t, testError, err)

	// Test error handling in PowerOffVM
	mockVMs.EXPECT().PowerOffVM("test").Return(nil, testError)
	_, err = mockVMs.PowerOffVM("test")
	assert.Error(t, err)
	assert.Equal(t, testError, err)
}

// Test VMs with OData options using generated mocks
func TestVMsWithODataOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVMs := mocks.NewMockVMs[TestVM](ctrl)
	ctx := context.Background()
	expectedVMs := []TestVM{
		{ExtId: ptr.To("test-1"), Name: ptr.To("Test VM 1")},
		{ExtId: ptr.To("test-2"), Name: ptr.To("Test VM 2")},
	}

	// Test List with OData options
	mockVMs.EXPECT().List(ctx, gomock.Any()).Return(expectedVMs, nil)
	entities, err := mockVMs.List(ctx, convergedclient.WithPage(1), convergedclient.WithLimit(10))
	assert.NoError(t, err)
	assert.Len(t, entities, 2)

	// Test NewIterator with OData options
	iterator := func(yield func(TestVM, error) bool) {
		for _, vm := range expectedVMs {
			yield(vm, nil)
		}
	}
	mockVMs.EXPECT().NewIterator(gomock.Any(), gomock.Any()).Return(iterator)
	resultIterator := mockVMs.NewIterator(ctx, convergedclient.WithOrderBy("name"))
	count := 0
	for _, err := range resultIterator {
		assert.NoError(t, err)
		count++
	}
	assert.Equal(t, 2, count)
}
