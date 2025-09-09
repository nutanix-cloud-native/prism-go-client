package v4

import (
	"context"
	"testing"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	"github.com/stretchr/testify/assert"
)

func TestVMsService(t *testing.T) {
	ctx := context.Background()
	service := NewVMsService(nil)
	assert.NotNil(t, service)

	// Test that all methods return "not implemented" errors
	vm, err := service.Get(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
	assert.Nil(t, vm)

	vms, err := service.List(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
	assert.Nil(t, vms)

	allVMs, err := service.ListAll(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
	assert.Nil(t, allVMs)

	iterator := service.NewIterator()
	assert.Nil(t, iterator) // Should return nil

	created, err := service.Create(ctx, &vmmModels.Vm{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
	assert.Nil(t, created)

	updated, err := service.Update(ctx, "test-uuid", &vmmModels.Vm{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
	assert.Nil(t, updated)

	deleted, err := service.Delete(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
	assert.Nil(t, deleted)

	// Test async operations
	operation, err := service.CreateAsync(ctx, &vmmModels.Vm{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
	assert.Nil(t, operation)

	operation, err = service.UpdateAsync(ctx, "test-uuid", &vmmModels.Vm{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
	assert.Nil(t, operation)

	operation, err = service.DeleteAsync(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
	assert.Nil(t, operation)

	// Test power operations
	operation, err = service.PowerOnVM("test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "PowerOnVM not implemented")
	assert.Nil(t, operation)

	operation, err = service.PowerOffVM("test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "PowerOffVM not implemented")
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

func TestVMsServiceWithODataOptions(t *testing.T) {
	ctx := context.Background()
	service := NewVMsService(nil)

	// Test List with OData options
	vms, err := service.List(ctx, converged.WithPage(1), converged.WithLimit(10))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
	assert.Nil(t, vms)

	// Test ListAll with OData options
	allVMs, err := service.ListAll(ctx, converged.WithFilter("name eq 'test'"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
	assert.Nil(t, allVMs)

	// Test NewIterator with OData options
	iterator := service.NewIterator(converged.WithOrderBy("name"))
	assert.Nil(t, iterator)
}

func TestVMsServiceErrorMessages(t *testing.T) {
	ctx := context.Background()
	service := NewVMsService(nil)

	// Test that error messages are consistent
	_, err := service.Get(ctx, "test-uuid")
	assert.Equal(t, "not implemented", err.Error())

	_, err = service.List(ctx)
	assert.Equal(t, "not implemented", err.Error())

	_, err = service.ListAll(ctx)
	assert.Equal(t, "not implemented", err.Error())

	_, err = service.Create(ctx, &vmmModels.Vm{})
	assert.Equal(t, "not implemented", err.Error())

	_, err = service.Update(ctx, "test-uuid", &vmmModels.Vm{})
	assert.Equal(t, "not implemented", err.Error())

	_, err = service.Delete(ctx, "test-uuid")
	assert.Equal(t, "not implemented", err.Error())

	_, err = service.CreateAsync(ctx, &vmmModels.Vm{})
	assert.Equal(t, "not implemented", err.Error())

	_, err = service.UpdateAsync(ctx, "test-uuid", &vmmModels.Vm{})
	assert.Equal(t, "not implemented", err.Error())

	_, err = service.DeleteAsync(ctx, "test-uuid")
	assert.Equal(t, "not implemented", err.Error())

	_, err = service.PowerOnVM("test-uuid")
	assert.Equal(t, "PowerOnVM not implemented", err.Error())

	_, err = service.PowerOffVM("test-uuid")
	assert.Equal(t, "PowerOffVM not implemented", err.Error())
}

func TestVMsServiceWithNilInputs(t *testing.T) {
	ctx := context.Background()
	service := NewVMsService(nil)

	// Test with empty UUID
	_, err := service.Get(ctx, "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	// Test with nil VM entity
	_, err = service.Create(ctx, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, err = service.Update(ctx, "test-uuid", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, err = service.CreateAsync(ctx, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, err = service.UpdateAsync(ctx, "test-uuid", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
}

func TestVMsServicePowerOperations(t *testing.T) {
	service := NewVMsService(nil)

	// Test PowerOnVM with different UUIDs
	_, err := service.PowerOnVM("")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "PowerOnVM not implemented")

	_, err = service.PowerOnVM("valid-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "PowerOnVM not implemented")

	// Test PowerOffVM with different UUIDs
	_, err = service.PowerOffVM("")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "PowerOffVM not implemented")

	_, err = service.PowerOffVM("valid-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "PowerOffVM not implemented")
}

func TestVMsServiceConsistency(t *testing.T) {
	ctx := context.Background()
	service1 := NewVMsService(nil)
	service2 := NewVMsService(nil)

	// Test that multiple instances behave consistently
	_, err1 := service1.Get(ctx, "test-uuid")
	_, err2 := service2.Get(ctx, "test-uuid")
	assert.Equal(t, err1.Error(), err2.Error())

	_, err1 = service1.List(ctx)
	_, err2 = service2.List(ctx)
	assert.Equal(t, err1.Error(), err2.Error())

	_, err1 = service1.PowerOnVM("test-uuid")
	_, err2 = service2.PowerOnVM("test-uuid")
	assert.Equal(t, err1.Error(), err2.Error())
}
