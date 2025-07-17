package facade

import (
	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

type VMsFacadeV4 interface {
	// GetVM returns the VM for the given UUID.
	GetVM(uuid string) (*vmmModels.Vm, error)

	// ListVMs returns a list of VMs.
	ListVMs(opts ...ODataOption) ([]vmmModels.Vm, error)

	// ListAllVMs returns all VMs without pagination.
	ListAllVMs(filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]vmmModels.Vm, error)

	// GetListIteratorVMs returns an iterator for listing VMs.
	GetListIteratorVMs(opts ...ODataOption) (ODataListIterator[*vmmModels.Vm], error)

	// CreateVM creates a new VM.
	CreateVM(vm *vmmModels.Vm) (*vmmModels.Vm, error)

	// UpdateVM updates an existing VM.
	UpdateVM(uuid string, vm *vmmModels.Vm) (*vmmModels.Vm, error)

	// DeleteVM deletes the VM with the given UUID.
	DeleteVM(uuid string) error

	// Additional methods can be added here as needed.
}
