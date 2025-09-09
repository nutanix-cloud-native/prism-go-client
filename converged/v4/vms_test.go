package v4

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	v4 "github.com/nutanix-cloud-native/prism-go-client/v4"
	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	"k8s.io/utils/ptr"
)

func TestVMsServiceClientNotInitialized(t *testing.T) {
	ctx := context.Background()
	service := NewVMsService(nil)
	assert.NotNil(t, service)

	// Test that all methods return "client is not initialized" errors
	vm, err := service.Get(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")
	assert.Nil(t, vm)

	vms, err := service.List(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")
	assert.Nil(t, vms)

	iterator := service.NewIterator(ctx, converged.WithPage(1), converged.WithLimit(10))
	assert.Nil(t, iterator) // Should return nil

	created, err := service.Create(ctx, &vmmModels.Vm{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")
	assert.Nil(t, created)

	updated, err := service.Update(ctx, "test-uuid", &vmmModels.Vm{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")
	assert.Nil(t, updated)

	err = service.Delete(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")

	// Test async operations
	operation, err := service.CreateAsync(ctx, &vmmModels.Vm{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")
	assert.Nil(t, operation)

	operation, err = service.UpdateAsync(ctx, "test-uuid", &vmmModels.Vm{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")
	assert.Nil(t, operation)

	operationDelete, err := service.DeleteAsync(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")
	assert.Nil(t, operationDelete)

	// Test power operations
	operation, err = service.PowerOnVM("test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")
	assert.Nil(t, operation)

	operation, err = service.PowerOffVM("test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")
	assert.Nil(t, operation)
}

func TestVMsServiceInterfaceCompliance(t *testing.T) {
	service := NewVMsService(nil)

	// Test that VMs service implements all required interfaces
	var getter converged.Getter[vmmModels.Vm] = service
	assert.NotNil(t, getter)

	var lister converged.Lister[vmmModels.Vm] = service
	assert.NotNil(t, lister)

	var creator converged.Creator[vmmModels.Vm] = service
	assert.NotNil(t, creator)

	var updater converged.Updater[vmmModels.Vm] = service
	assert.NotNil(t, updater)

	var deleter converged.Deleter[vmmModels.Vm] = service
	assert.NotNil(t, deleter)

	var asyncCreator converged.AsyncCreator[vmmModels.Vm] = service
	assert.NotNil(t, asyncCreator)

	var asyncUpdater converged.AsyncUpdater[vmmModels.Vm] = service
	assert.NotNil(t, asyncUpdater)

	var asyncDeleter converged.AsyncDeleter[vmmModels.Vm] = service
	assert.NotNil(t, asyncDeleter)
}

// TestVMsOperations tests VM operations using Keploy for API recording/replay
func TestVMsOperations(t *testing.T) {
	// Get credentials from environment (or use dummy ones for CI)
	creds := testhelpers.CredentialsFromEnvironment(t)

	// Create v4 client
	v4Client, err := v4.NewV4Client(creds)
	require.NoError(t, err)

	// Create VMs service
	vmService := NewVMsService(v4Client)

	ctx := context.Background()

	t.Run("ListVMs", func(t *testing.T) {
		// Test listing VMs
		vms, err := vmService.List(ctx)
		require.NoError(t, err)
		assert.NotNil(t, vms)
		t.Logf("Found %d VMs", len(vms))
	})

	t.Run("ListVMsWithPagination", func(t *testing.T) {
		// Test listing VMs with pagination
		vms, err := vmService.List(ctx, converged.WithPage(0), converged.WithLimit(5))
		require.NoError(t, err)
		assert.NotNil(t, vms)
		assert.LessOrEqual(t, len(vms), 5)
		t.Logf("Found %d VMs with pagination", len(vms))
	})

	t.Run("ListVMsWithOrderBy", func(t *testing.T) {
		// Test listing VMs with ordering
		vms, err := vmService.List(ctx, converged.WithOrderBy("name"))
		require.NoError(t, err)
		assert.NotNil(t, vms)
		t.Logf("Found %d VMs ordered by name", len(vms))
	})

	t.Run("ListVMsWithFilter", func(t *testing.T) {
		// Test listing VMs with filter
		vms, err := vmService.List(ctx, converged.WithFilter("name==test-vm"))
		require.NoError(t, err)
		assert.NotNil(t, vms)
		t.Logf("Found %d VMs with filter", len(vms))
	})

	t.Run("NewIterator", func(t *testing.T) {
		// Test VM iterator
		iterator := vmService.NewIterator(ctx, converged.WithLimit(3))
		require.NotNil(t, iterator)

		count := 0
		for vm, err := range iterator {
			require.NoError(t, err)
			assert.NotNil(t, vm)
			count++
			if count >= 3 {
				break
			}
		}
		t.Logf("Iterated through %d VMs", count)
	})
}

// TestVMsCRUDOperations tests VM CRUD operations
func TestVMsCRUDOperations(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	v4Client, err := v4.NewV4Client(creds)
	require.NoError(t, err)

	vmService := NewVMsService(v4Client)

	ctx := context.Background()

	// Create a test VM
	testVM := &vmmModels.Vm{
		Name: ptr.To("test-vm-crud"),
	}

	t.Run("CreateVM", func(t *testing.T) {
		// Test VM creation
		createdVM, err := vmService.Create(ctx, testVM)
		require.NoError(t, err)
		assert.NotNil(t, createdVM)
		assert.Equal(t, "test-vm-crud", *createdVM.Name)
		t.Logf("Created VM with UUID: %s", *createdVM.ExtId)

		// Store the created VM for cleanup
		testVM = createdVM
	})

	t.Run("GetVM", func(t *testing.T) {
		// Test getting a specific VM
		vm, err := vmService.Get(ctx, *testVM.ExtId)
		require.NoError(t, err)
		assert.NotNil(t, vm)
		assert.Equal(t, "test-vm-crud", *vm.Name)
		assert.Equal(t, *testVM.ExtId, *vm.ExtId)
	})

	t.Run("UpdateVM", func(t *testing.T) {
		// Test updating VM
		updatedVM := *testVM
		updatedVM.Name = ptr.To("updated-vm-name")

		result, err := vmService.Update(ctx, *testVM.ExtId, &updatedVM)
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "updated-vm-name", *result.Name)
	})

	t.Run("DeleteVM", func(t *testing.T) {
		// Test VM deletion
		err := vmService.Delete(ctx, *testVM.ExtId)
		require.NoError(t, err)
	})
}

// TestVMsAsyncOperations tests VM async operations
func TestVMsAsyncOperations(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	v4Client, err := v4.NewV4Client(creds)
	require.NoError(t, err)

	vmService := NewVMsService(v4Client)

	ctx := context.Background()

	testVM := &vmmModels.Vm{
		Name: ptr.To("test-vm-async"),
	}

	t.Run("CreateVMAsync", func(t *testing.T) {
		// Test async VM creation
		operation, err := vmService.CreateAsync(ctx, testVM)
		require.NoError(t, err)
		assert.NotNil(t, operation)

		// Wait for operation to complete
		result, err := operation.Wait(ctx)
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 1)
		assert.Equal(t, "test-vm-async", *result[0].Name)

		// Store for cleanup
		testVM = result[0]
	})

	t.Run("UpdateVMAsync", func(t *testing.T) {
		// Test async VM update
		updatedVM := *testVM
		updatedVM.Name = ptr.To("updated-async-vm")

		operation, err := vmService.UpdateAsync(ctx, *testVM.ExtId, &updatedVM)
		require.NoError(t, err)
		assert.NotNil(t, operation)

		// Wait for operation to complete
		result, err := operation.Wait(ctx)
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 1)
		assert.Equal(t, "updated-async-vm", *result[0].Name)
	})

	t.Run("DeleteVMAsync", func(t *testing.T) {
		// Test async VM deletion
		operation, err := vmService.DeleteAsync(ctx, *testVM.ExtId)
		require.NoError(t, err)
		assert.NotNil(t, operation)

		// Wait for operation to complete
		_, err = operation.Wait(ctx)
		require.NoError(t, err)
	})
}

// TestVMsPowerOperations tests VM power operations
func TestVMsPowerOperations(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	v4Client, err := v4.NewV4Client(creds)
	require.NoError(t, err)

	vmService := NewVMsService(v4Client)

	ctx := context.Background()

	// Create a test VM for power operations
	testVM := &vmmModels.Vm{
		Name: ptr.To("test-vm-power"),
	}

	// Create VM first
	createdVM, err := vmService.Create(ctx, testVM)
	require.NoError(t, err)
	testVM = createdVM

	t.Run("PowerOnVM", func(t *testing.T) {
		// Test powering on VM
		operation, err := vmService.PowerOnVM(*testVM.ExtId)
		require.NoError(t, err)
		assert.NotNil(t, operation)

		// Wait for operation to complete
		result, err := operation.Wait(ctx)
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 1)
		t.Logf("Powered on VM: %s", *result[0].Name)
	})

	t.Run("PowerOffVM", func(t *testing.T) {
		// Test powering off VM
		operation, err := vmService.PowerOffVM(*testVM.ExtId)
		require.NoError(t, err)
		assert.NotNil(t, operation)

		// Wait for operation to complete
		result, err := operation.Wait(ctx)
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 1)
		t.Logf("Powered off VM: %s", *result[0].Name)
	})

	// Cleanup
	t.Run("Cleanup", func(t *testing.T) {
		err := vmService.Delete(ctx, *testVM.ExtId)
		require.NoError(t, err)
	})
}

// TestVMsErrorHandling tests VM error handling scenarios
func TestVMsErrorHandling(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	v4Client, err := v4.NewV4Client(creds)
	require.NoError(t, err)

	vmService := NewVMsService(v4Client)
	ctx := context.Background()

	t.Run("GetNonExistentVM", func(t *testing.T) {
		// Test getting a non-existent VM
		vm, err := vmService.Get(ctx, "non-existent-uuid")
		assert.Error(t, err)
		assert.Nil(t, vm)
	})

	t.Run("UpdateNonExistentVM", func(t *testing.T) {
		// Test updating a non-existent VM
		testVM := &vmmModels.Vm{
			Name: ptr.To("non-existent"),
		}
		vm, err := vmService.Update(ctx, "non-existent-uuid", testVM)
		assert.Error(t, err)
		assert.Nil(t, vm)
	})

	t.Run("DeleteNonExistentVM", func(t *testing.T) {
		// Test deleting a non-existent VM
		err := vmService.Delete(ctx, "non-existent-uuid")
		assert.Error(t, err)
	})

	t.Run("PowerOnNonExistentVM", func(t *testing.T) {
		// Test powering on a non-existent VM
		operation, err := vmService.PowerOnVM("non-existent-uuid")
		assert.Error(t, err)
		assert.Nil(t, operation)
	})

	t.Run("PowerOffNonExistentVM", func(t *testing.T) {
		// Test powering off a non-existent VM
		operation, err := vmService.PowerOffVM("non-existent-uuid")
		assert.Error(t, err)
		assert.Nil(t, operation)
	})
}

// TestVMsWithODataOptions tests VM operations with various OData options
func TestVMsWithODataOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	v4Client, err := v4.NewV4Client(creds)
	require.NoError(t, err)

	vmService := NewVMsService(v4Client)
	ctx := context.Background()

	t.Run("ListWithMultipleOptions", func(t *testing.T) {
		// Test listing with multiple OData options
		vms, err := vmService.List(ctx,
			converged.WithPage(0),
			converged.WithLimit(10),
			converged.WithOrderBy("name"),
			converged.WithSelect("name,extId"),
		)
		require.NoError(t, err)
		assert.NotNil(t, vms)
		t.Logf("Found %d VMs with multiple OData options", len(vms))
	})

	t.Run("IteratorWithOptions", func(t *testing.T) {
		// Test iterator with OData options
		iterator := vmService.NewIterator(ctx,
			converged.WithPage(0),
			converged.WithLimit(5),
			converged.WithOrderBy("name"),
		)
		require.NotNil(t, iterator)

		count := 0
		for vm, err := range iterator {
			require.NoError(t, err)
			assert.NotNil(t, vm)
			count++
			if count >= 5 {
				break
			}
		}
		t.Logf("Iterated through %d VMs with options", count)
	})
}

// TestVMsServiceIntegration tests the VMs service integration
func TestVMsServiceIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	v4Client, err := v4.NewV4Client(creds)
	require.NoError(t, err)

	vmService := NewVMsService(v4Client)
	ctx := context.Background()

	t.Run("ServiceImplementsInterfaces", func(t *testing.T) {
		// Test that VMs service implements all required interfaces
		var getter converged.Getter[vmmModels.Vm] = vmService
		assert.NotNil(t, getter)

		var lister converged.Lister[vmmModels.Vm] = vmService
		assert.NotNil(t, lister)

		var creator converged.Creator[vmmModels.Vm] = vmService
		assert.NotNil(t, creator)

		var updater converged.Updater[vmmModels.Vm] = vmService
		assert.NotNil(t, updater)

		var deleter converged.Deleter[vmmModels.Vm] = vmService
		assert.NotNil(t, deleter)

		var asyncCreator converged.AsyncCreator[vmmModels.Vm] = vmService
		assert.NotNil(t, asyncCreator)

		var asyncUpdater converged.AsyncUpdater[vmmModels.Vm] = vmService
		assert.NotNil(t, asyncUpdater)

		var asyncDeleter converged.AsyncDeleter[vmmModels.Vm] = vmService
		assert.NotNil(t, asyncDeleter)
	})

	t.Run("ServiceConsistency", func(t *testing.T) {
		// Test that service behaves consistently
		service1 := NewVMsService(v4Client)
		service2 := NewVMsService(v4Client)

		vms1, err1 := service1.List(ctx)
		vms2, err2 := service2.List(ctx)

		if err1 != nil && err2 != nil {
			assert.Equal(t, err1.Error(), err2.Error())
		} else {
			assert.NoError(t, err1)
			assert.NoError(t, err2)
			assert.Equal(t, len(vms1), len(vms2))
		}
	})
}
