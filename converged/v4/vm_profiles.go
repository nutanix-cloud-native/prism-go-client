package v4

import (
	"context"
	"errors"
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	vmmConfig "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	vmProfileRequest "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/vmprofile"
)

// VMProfilesService provides implementation for all VMProfiles interface methods.
type VMProfilesService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewVMProfilesService creates a new VMProfilesService instance.
func NewVMProfilesService(client *v4prismGoClient.Client) *VMProfilesService {
	return &VMProfilesService{
		client:       client,
		entitiesName: "VM Profile",
	}
}

// Get returns the VM Profile for the given UUID.
func (s *VMProfilesService) Get(ctx context.Context, uuid string) (*vmmModels.VmProfile, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	return GenericGetEntity[*vmmModels.GetVmProfileApiResponse, vmmModels.VmProfile](
		func() (*vmmModels.GetVmProfileApiResponse, error) {
			return s.client.VmProfilesApiInstance.ServiceClient.GetVmProfileById(ctx, &vmProfileRequest.GetVmProfileByIdRequest{
				ExtId: &uuid,
			})
		},
		s.entitiesName,
	)
}

// List returns a list of VM Profiles.
// If no page and limit are provided, the API will return all VM Profiles.
func (s *VMProfilesService) List(ctx context.Context, opts ...converged.ODataOption) ([]vmmModels.VmProfile, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("expand and apply are not supported for listing VM Profiles")
	}

	return GenericListEntities[*vmmModels.ListVmProfilesApiResponse, vmmModels.VmProfile](
		func(reqParams *V4ODataParams) (*vmmModels.ListVmProfilesApiResponse, error) {
			return s.client.VmProfilesApiInstance.ServiceClient.ListVmProfiles(ctx, &vmProfileRequest.ListVmProfilesRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Select_:  reqParams.Select,
			})
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing VM Profiles.
func (s *VMProfilesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[vmmModels.VmProfile] {
	if s.client == nil {
		return nil
	}
	return GenericNewIterator[*vmmModels.ListVmProfilesApiResponse, vmmModels.VmProfile](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*vmmModels.ListVmProfilesApiResponse, error) {
			return s.client.VmProfilesApiInstance.ServiceClient.ListVmProfiles(ctx, &vmProfileRequest.ListVmProfilesRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Select_:  reqParams.Select,
			})
		},
		opts,
		s.entitiesName,
	)
}

// DeployVmWithVmProfile deploys a VM from a VM Profile.
func (s *VMProfilesService) DeployVmWithVmProfile(ctx context.Context, vmProfileUUID string, params *vmmModels.DeployVmFromVmProfileParams) (converged.Operation[vmmModels.Vm], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	// Get VMs service to use its Get method for the operation
	vmsService := NewVMsService(s.client)

	taskRef, err := CallAPI[*vmmModels.DeployVmFromVmProfileApiResponse, vmmConfig.TaskReference](
		s.client.VmProfilesApiInstance.ServiceClient.DeployVmFromVmProfile(ctx, &vmProfileRequest.DeployVmFromVmProfileRequest{
			ExtId: &vmProfileUUID,
			Body:  params,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy VM from VM Profile: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for deployed VM")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		vmsService.Get,
	), nil
}
