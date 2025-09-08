package v4

import (
	"context"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// VMsService provides default "not implemented" implementation for all VMs interface methods.
type VMsService struct {
	client *v4prismGoClient.Client
}

// NewVMsService creates a new VMsService instance.
func NewVMsService(client *v4prismGoClient.Client) *VMsService {
	return &VMsService{client: client}
}

// Get returns the VM for the given UUID.
func (s *VMsService) Get(ctx context.Context, uuid string) (*vmmModels.Vm, error) {
	return nil, fmt.Errorf("not implemented")
}

// List returns a list of VMs.
func (s *VMsService) List(ctx context.Context, opts ...converged.ODataOption) ([]vmmModels.Vm, error) {
	return nil, fmt.Errorf("not implemented")
}

// ListAll returns all VMs without pagination.
func (s *VMsService) ListAll(ctx context.Context, opts ...converged.ODataOption) ([]vmmModels.Vm, error) {
	return nil, fmt.Errorf("not implemented")
}

// NewIterator returns an iterator for listing VMs.
func (s *VMsService) NewIterator(opts ...converged.ODataOption) converged.Iterator[vmmModels.Vm] {
	return nil
}

// Create creates a new VM.
func (s *VMsService) Create(ctx context.Context, entity *vmmModels.Vm) (*vmmModels.Vm, error) {
	return nil, fmt.Errorf("not implemented")
}

// Update updates an existing VM.
func (s *VMsService) Update(ctx context.Context, uuid string, entity *vmmModels.Vm) (*vmmModels.Vm, error) {
	return nil, fmt.Errorf("not implemented")
}

// Delete deletes an existing VM.
func (s *VMsService) Delete(ctx context.Context, uuid string) (*vmmModels.Vm, error) {
	return nil, fmt.Errorf("not implemented")
}

// CreateAsync creates a new VM asynchronously.
func (s *VMsService) CreateAsync(ctx context.Context, entity *vmmModels.Vm) (converged.Operation[vmmModels.Vm], error) {
	return nil, fmt.Errorf("not implemented")
}

// UpdateAsync updates an existing VM asynchronously.
func (s *VMsService) UpdateAsync(ctx context.Context, uuid string, entity *vmmModels.Vm) (converged.Operation[vmmModels.Vm], error) {
	return nil, fmt.Errorf("not implemented")
}

// DeleteAsync deletes an existing VM asynchronously.
func (s *VMsService) DeleteAsync(ctx context.Context, uuid string) (converged.Operation[vmmModels.Vm], error) {
	return nil, fmt.Errorf("not implemented")
}

// PowerOnVM powers on the VM with the given UUID.
func (s *VMsService) PowerOnVM(uuid string) (converged.Operation[vmmModels.Vm], error) {
	return nil, fmt.Errorf("PowerOnVM not implemented")
}

// PowerOffVM powers off the VM with the given UUID.
func (s *VMsService) PowerOffVM(uuid string) (converged.Operation[vmmModels.Vm], error) {
	return nil, fmt.Errorf("PowerOffVM not implemented")
}
