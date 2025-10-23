package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	prismModels "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/prism/v4/config"
	volumeModels "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
)

// VolumeGroupsService provides default "not implemented" implementation for all VolumeGroups interface methods.
type VolumeGroupsService struct {
	client   *v4prismGoClient.Client
	entities string
}

// NewVolumeGroupsService creates a new VolumeGroupsService instance.
func NewVolumeGroupsService(client *v4prismGoClient.Client) *VolumeGroupsService {
	return &VolumeGroupsService{
		client:   client,
		entities: "volume group",
	}
}

// Get returns the volume group for the given UUID.
func (s *VolumeGroupsService) Get(ctx context.Context, uuid string) (*volumeModels.VolumeGroup, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	return GenericGetEntity[*volumeModels.GetVolumeGroupApiResponse, volumeModels.VolumeGroup](
		func() (*volumeModels.GetVolumeGroupApiResponse, error) {
			return s.client.VolumeGroupsApiInstance.GetVolumeGroupById(&uuid)
		},
		s.entities,
	)
}

// List returns a list of volume groups.
func (s *VolumeGroupsService) List(ctx context.Context, opts ...converged.ODataOption) ([]volumeModels.VolumeGroup, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	return GenericListEntities[*volumeModels.ListVolumeGroupsApiResponse, volumeModels.VolumeGroup](
		func(reqParams *V4ODataParams) (*volumeModels.ListVolumeGroupsApiResponse, error) {
			return s.client.VolumeGroupsApiInstance.ListVolumeGroups(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		s.entities,
	)
}

// NewIterator returns an iterator for listing volume groups.
func (s *VolumeGroupsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[volumeModels.VolumeGroup] {
	return GenericNewIterator[*volumeModels.ListVolumeGroupsApiResponse, volumeModels.VolumeGroup](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*volumeModels.ListVolumeGroupsApiResponse, error) {
			return s.client.VolumeGroupsApiInstance.ListVolumeGroups(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		s.entities,
	)
}

// AttachToVMAsync attaches a volume group to a VM asynchronously.
func (s *VolumeGroupsService) AttachToVMAsync(ctx context.Context, volumeGroupUUID string, vmAttachment volumeModels.VmAttachment) (converged.Operation[volumeModels.VolumeGroup], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	taskRef, err := CallAPI[*volumeModels.AttachVmApiResponse, prismModels.TaskReference](
		s.client.VolumeGroupsApiInstance.AttachVm(&volumeGroupUUID, &vmAttachment),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to attach volume group to VM asynchronously: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for attached volume group")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		s.Get,
	), nil
}

// DetachFromVMAsync detaches a volume group from a VM asynchronously.
func (s *VolumeGroupsService) DetachFromVMAsync(ctx context.Context, volumeGroupUUID string, vmAttachment volumeModels.VmAttachment) (converged.Operation[volumeModels.VolumeGroup], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	taskRef, err := CallAPI[*volumeModels.DetachVmApiResponse, prismModels.TaskReference](
		s.client.VolumeGroupsApiInstance.DetachVm(&volumeGroupUUID, &vmAttachment),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to detach volume group from VM asynchronously: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for detached volume group")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		s.Get,
	), nil
}

// DetachFromVM detaches a volume group from a VM.
func (s *VolumeGroupsService) DetachFromVM(ctx context.Context, volumeGroupUUID string, vmAttachment volumeModels.VmAttachment) (*volumeModels.VolumeGroup, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	operation, err := s.DetachFromVMAsync(ctx, volumeGroupUUID, vmAttachment)
	if err != nil {
		return nil, fmt.Errorf("failed to detach volume group from VM: %w", err)
	}

	result, err := operation.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to detach volume group from VM: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("failed to detach volume group from VM: operation completed but no volume group returned")
	}

	return result[0], nil
}

// AttachToVM attaches a volume group to a VM.
func (s *VolumeGroupsService) AttachToVM(ctx context.Context, volumeGroupUUID string, vmAttachment volumeModels.VmAttachment) (*volumeModels.VolumeGroup, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	operation, err := s.AttachToVMAsync(ctx, volumeGroupUUID, vmAttachment)
	if err != nil {
		return nil, fmt.Errorf("failed to attach volume group to VM: %w", err)
	}

	result, err := operation.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to attach volume group to VM: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("failed to attach volume group to VM: operation completed but no volume group returned")
	}

	return result[0], nil
}
