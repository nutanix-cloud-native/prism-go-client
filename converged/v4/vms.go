package v4

import (
	"context"
	"errors"
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	vmmConfig "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// VMsService provides default "not implemented" implementation for all VMs interface methods.
type VMsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewVMsService creates a new VMsService instance.
func NewVMsService(client *v4prismGoClient.Client) *VMsService {
	return &VMsService{
		client:       client,
		entitiesName: "VM",
	}
}

// Get returns the VM for the given UUID.
func (s *VMsService) Get(ctx context.Context, uuid string) (*vmmModels.Vm, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	return GenericGetEntity[*vmmModels.GetVmApiResponse, vmmModels.Vm](
		func() (*vmmModels.GetVmApiResponse, error) {
			return s.client.VmApiInstance.GetVmById(&uuid)
		},
		s.entitiesName,
	)
}

// List returns a list of VMs.
// If no page and limit are provided, the API will return all VMs.
func (s *VMsService) List(ctx context.Context, opts ...converged.ODataOption) ([]vmmModels.Vm, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams.Expand != nil || myParams.Apply != nil {
		return nil, fmt.Errorf("expand and apply are not supported for listing VMs")
	}

	return GenericListEntities[*vmmModels.ListVmsApiResponse, vmmModels.Vm](
		func(reqParams *V4ODataParams) (*vmmModels.ListVmsApiResponse, error) {
			return s.client.VmApiInstance.ListVms(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing VMs.
func (s *VMsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[vmmModels.Vm] {
	if s.client == nil {
		return nil
	}
	return GenericNewIterator[*vmmModels.ListVmsApiResponse, vmmModels.Vm](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*vmmModels.ListVmsApiResponse, error) {
			return s.client.VmApiInstance.ListVms(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// Create creates a new VM.
func (s *VMsService) Create(ctx context.Context, vm *vmmModels.Vm) (*vmmModels.Vm, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	operation, err := s.CreateAsync(ctx, vm)
	if err != nil {
		return nil, fmt.Errorf("failed to create VM: %w", err)
	}

	result, err := operation.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create VM: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("failed to create VM: operation completed but no VM returned")
	}

	return result[0], nil
}

// Update updates an existing VM.
func (s *VMsService) Update(ctx context.Context, uuid string, entity *vmmModels.Vm) (*vmmModels.Vm, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	operation, err := s.UpdateAsync(ctx, uuid, entity)
	if err != nil {
		return nil, fmt.Errorf("failed to update VM: %w", err)
	}

	result, err := operation.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update VM: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("failed to update VM: operation completed but no VM returned")
	}

	return result[0], nil
}

// Delete deletes an existing VM.
func (s *VMsService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}
	operation, err := s.DeleteAsync(ctx, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete VM: %w", err)
	}

	_, err = operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete VM: %w", err)
	}

	return nil
}

// CreateAsync creates a new VM asynchronously.
func (s *VMsService) CreateAsync(ctx context.Context, vm *vmmModels.Vm) (converged.Operation[vmmModels.Vm], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	taskRef, err := CallAPI[*vmmModels.CreateVmApiResponse, vmmConfig.TaskReference](
		s.client.VmApiInstance.CreateVm(vm),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create VM: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for created VM")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		s.Get,
	), nil
}

// UpdateAsync updates an existing VM asynchronously.
func (s *VMsService) UpdateAsync(ctx context.Context, uuid string, vm *vmmModels.Vm) (converged.Operation[vmmModels.Vm], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	currentVM, args, err := GetEntityAndEtag(
		s.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get VM for update: %w", err)
	}

	vm = CopyEtag(currentVM, vm).(*vmmModels.Vm)

	taskRef, err := CallAPI[*vmmModels.UpdateVmApiResponse, vmmConfig.TaskReference](
		s.client.VmApiInstance.UpdateVmById(&uuid, vm, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update VM: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for updated VM")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		s.Get,
	), nil
}

// DeleteAsync deletes an existing VM asynchronously.
func (s *VMsService) DeleteAsync(ctx context.Context, uuid string) (converged.Operation[converged.NoEntity], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	// Get the VM first to retrieve its eTag
	_, args, err := GetEntityAndEtag(
		s.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get VM for deletion: %w", err)
	}

	taskRef, err := CallAPI[*vmmModels.DeleteVmApiResponse, vmmConfig.TaskReference](
		s.client.VmApiInstance.DeleteVmById(&uuid, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete VM: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for deleted VM")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		func(ctx context.Context, uuid string) (*converged.NoEntity, error) {
			return converged.NoEntityGetter(ctx, uuid)
		},
	), nil
}

// PowerOnVM powers on the VM with the given UUID.
func (s *VMsService) PowerOnVM(uuid string) (converged.Operation[vmmModels.Vm], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	_, args, err := GetEntityAndEtag(
		s.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get VM for power on: %w", err)
	}

	taskRef, err := CallAPI[*vmmModels.PowerOnVmApiResponse, vmmConfig.TaskReference](
		s.client.VmApiInstance.PowerOnVm(&uuid, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to power on VM: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for powered on VM")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		s.Get,
	), nil
}

// PowerOffVM powers off the VM with the given UUID.
func (s *VMsService) PowerOffVM(uuid string) (converged.Operation[vmmModels.Vm], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	_, args, err := GetEntityAndEtag(
		s.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get VM for power off: %w", err)
	}

	taskRef, err := CallAPI[*vmmModels.PowerOffVmApiResponse, vmmConfig.TaskReference](
		s.client.VmApiInstance.PowerOffVm(&uuid, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to power off VM: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for powered off VM")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		s.Get,
	), nil
}
