package facade

import (
	"context"

	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

type VMsFacadeV4 interface {
	// GetVM returns the VM for the given UUID.
	GetVM(ctx context.Context, uuid string) (*vmmModels.Vm, error)

	// ListVMs returns a list of VMs.
	ListVMs(ctx context.Context, opts ...ODataOption) ([]vmmModels.Vm, error)

	// ListAllVMs returns all VMs without pagination.
	ListAllVMs(ctx context.Context, filterParam *string, orderbyParam *string, selectParam *string) ([]vmmModels.Vm, error)

	// GetListIteratorVMs returns an iterator for listing VMs.
	GetListIteratorVMs(ctx context.Context, opts ...ODataOption) ODataListIterator[vmmModels.Vm]

	// CreateVM creates a new VM.
	CreateVM(ctx context.Context, vm *vmmModels.Vm) (TaskWaiter[vmmModels.Vm], error)

	// UpdateVM updates an existing VM.
	UpdateVM(ctx context.Context, uuid string, vm *vmmModels.Vm) (TaskWaiter[vmmModels.Vm], error)

	// DeleteVM deletes the VM with the given UUID.
	DeleteVM(ctx context.Context, uuid string) (TaskWaiter[NoEntity], error)

	// PowerOnVM powers on the VM with the given UUID.
	PowerOnVM(ctx context.Context, uuid string) (TaskWaiter[vmmModels.Vm], error)

	// PowerOffVM powers off the VM with the given UUID.
	PowerOffVM(ctx context.Context, uuid string) (TaskWaiter[vmmModels.Vm], error)

	// Additional methods can be added here as needed.
}
