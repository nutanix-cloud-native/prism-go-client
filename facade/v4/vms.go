package v4

import (
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	"github.com/nutanix-cloud-native/prism-go-client/facade/ferrors"
	v4VmmConfig "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	vmmModelsError "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/error"
)

func (f *FacadeV4Client) GetVM(uuid string) (*vmmModels.Vm, error) {
	return CommonGetEntity[*vmmModels.GetVmApiResponse, vmmModels.Vm, *vmmModelsError.OneOfErrorResponseError](
		func() (*vmmModels.GetVmApiResponse, error) {
			return f.client.VmApiInstance.GetVmById(&uuid)
		},
		"VM",
	)
}

func (f *FacadeV4Client) ListVMs(opts ...facade.ODataOption) ([]vmmModels.Vm, error) {
	return CommonListEntities[*vmmModels.ListVmsApiResponse, vmmModels.Vm, *vmmModelsError.OneOfErrorResponseError](
		func(reqParams *V4ODataParams) (*vmmModels.ListVmsApiResponse, error) {
			return f.client.VmApiInstance.ListVms(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		"VMs",
	)
}

func (f *FacadeV4Client) ListAllVMs(filterParam *string, orderbyParam *string, selectParam *string) ([]vmmModels.Vm, error) {
	reqParams := &V4ODataParams{
		Filter:  filterParam,
		OrderBy: orderbyParam,
		Select:  selectParam,
	}

	return CommonListAllEntities[*vmmModels.ListVmsApiResponse, vmmModels.Vm, *vmmModelsError.OneOfErrorResponseError](
		func(reqParams *V4ODataParams) (*vmmModels.ListVmsApiResponse, error) {
			return f.client.VmApiInstance.ListVms(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		reqParams,
		"VMs",
	)
}

func (f *FacadeV4Client) GetListIteratorVMs(opts ...facade.ODataOption) facade.ODataListIterator[vmmModels.Vm] {
	return NewFacadeV4ODataIterator[*vmmModels.ListVmsApiResponse, vmmModels.Vm, *vmmModelsError.OneOfErrorResponseError](
		f.client,
		func(params *V4ODataParams) (*vmmModels.ListVmsApiResponse, error) {
			return f.client.VmApiInstance.ListVms(
				params.Page,
				params.Limit,
				params.Filter,
				params.OrderBy,
				params.Select,
			)
		},
		opts...,
	)
}

func (f *FacadeV4Client) CreateVM(vm *vmmModels.Vm) (facade.TaskWaiter[vmmModels.Vm], error) {
	taskRef, err := CallAPI[*vmmModels.CreateVmApiResponse, v4VmmConfig.TaskReference, *vmmModelsError.OneOfErrorResponseError](
		f.client.VmApiInstance.CreateVm(vm),
	)
	if err != nil {
		return nil, err
	}

	if taskRef.ExtId == nil {
		return nil, ferrors.NewErrUncategorisedError("", fmt.Errorf("task reference ExtId is nil for created VM"))
	}

	return NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, f.GetVM), nil
}

func (f *FacadeV4Client) UpdateVM(uuid string, vm *vmmModels.Vm) (facade.TaskWaiter[vmmModels.Vm], error) {
	currentVM, args, err := GetEntityAndEtag(
		f.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, err
	}

	vm = CopyEtag(currentVM, vm).(*vmmModels.Vm)

	taskRef, err := CallAPI[*vmmModels.UpdateVmApiResponse, v4VmmConfig.TaskReference, *vmmModelsError.OneOfErrorResponseError](
		f.client.VmApiInstance.UpdateVmById(&uuid, vm, args),
	)
	if err != nil {
		return nil, err
	}

	if taskRef.ExtId == nil {
		return nil, ferrors.NewErrUncategorisedError("task reference ExtId is nil for updated VM", nil)
	}

	waiter := NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, f.GetVM)
	return waiter, nil
}

func (f *FacadeV4Client) DeleteVM(uuid string) (facade.TaskWaiter[facade.NoEntity], error) {
	_, args, err := GetEntityAndEtag(
		f.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, err
	}

	taskRef, err := CallAPI[*vmmModels.DeleteVmApiResponse, v4VmmConfig.TaskReference, *vmmModelsError.OneOfErrorResponseError](
		f.client.VmApiInstance.DeleteVmById(&uuid, args),
	)
	if err != nil {
		return nil, err
	}

	if taskRef.ExtId == nil {
		return nil, ferrors.NewErrUncategorisedError("task reference ExtId is nil for deleted VM", nil)
	}

	waiter := NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, facade.NoEntityGetter)
	return waiter, nil
}

func (f *FacadeV4Client) PowerOnVM(uuid string) (facade.TaskWaiter[vmmModels.Vm], error) {
	_, args, err := GetEntityAndEtag(
		f.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, err
	}

	taskRef, err := CallAPI[*vmmModels.PowerOnVmApiResponse, v4VmmConfig.TaskReference, *vmmModelsError.OneOfErrorResponseError](
		f.client.VmApiInstance.PowerOnVm(&uuid, args),
	)
	if err != nil {
		return nil, err
	}

	if taskRef.ExtId == nil {
		return nil, ferrors.NewErrUncategorisedError("task reference ExtId is nil for powered on VM", nil)
	}

	waiter := NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, f.GetVM)
	return waiter, nil
}

func (f *FacadeV4Client) PowerOffVM(uuid string) (facade.TaskWaiter[vmmModels.Vm], error) {
	_, args, err := GetEntityAndEtag(
		f.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, err
	}

	taskRef, err := CallAPI[*vmmModels.PowerOffVmApiResponse, v4VmmConfig.TaskReference, *vmmModelsError.OneOfErrorResponseError](
		f.client.VmApiInstance.PowerOffVm(&uuid, args),
	)
	if err != nil {
		return nil, err
	}

	if taskRef.ExtId == nil {
		return nil, ferrors.NewErrUncategorisedError("task reference ExtId is nil for powered off VM", nil)
	}

	waiter := NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, f.GetVM)
	return waiter, nil
}
