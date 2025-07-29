package v4

import (
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	v4VmmConfig "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

func (f *FacadeV4Client) GetVM(uuid string) (*vmmModels.Vm, error) {
	result, err := CallAPI[*vmmModels.GetVmApiResponse, vmmModels.Vm](
		f.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (f *FacadeV4Client) ListVMs(opts ...facade.ODataOption) ([]vmmModels.Vm, error) {
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to OData params: %w", err)
	}

	result, err := CallAPI[*vmmModels.ListVmsApiResponse, []vmmModels.Vm](
		f.client.VmApiInstance.ListVms(reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy, reqParams.Select),
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (f *FacadeV4Client) ListAllVMs(filterParam *string, orderbyParam *string, selectParam *string) ([]vmmModels.Vm, error) {
	var result []vmmModels.Vm
	page := 0

	reqParams := &V4ODataParams{
		Page:    &page,
		Limit:   nil, // Let API use the default limit
		Filter:  filterParam,
		OrderBy: orderbyParam,
		Select:  selectParam,
	}

	vms, totalCount, err := CallListAPI[*vmmModels.ListVmsApiResponse, vmmModels.Vm](
		f.client.VmApiInstance.ListVms(
			reqParams.Page,
			reqParams.Limit,
			reqParams.Filter,
			reqParams.OrderBy,
			reqParams.Select,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list all VMs: %w", err)
	}
	result = append(result, vms...)

	for len(result) < totalCount {
		page++
		reqParams.Page = &page
		vms, _, err = CallListAPI[*vmmModels.ListVmsApiResponse, vmmModels.Vm](
			f.client.VmApiInstance.ListVms(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to list all VMs: %w", err)
		}
		result = append(result, vms...)
	}

	return result, nil
}

func (f *FacadeV4Client) GetListIteratorVMs(opts ...facade.ODataOption) (facade.ODataListIterator[vmmModels.Vm], error) {
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to OData params: %w", err)
	}

	vms, totalCount, err := CallListAPI[*vmmModels.ListVmsApiResponse, vmmModels.Vm](
		f.client.VmApiInstance.ListVms(
			reqParams.Page,
			reqParams.Limit,
			reqParams.Filter,
			reqParams.OrderBy,
			reqParams.Select,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list VMs: %w", err)
	}

	return NewFacadeV4ODataIterator(
		f.client,
		totalCount,
		vms,
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
	), nil
}

func (f *FacadeV4Client) CreateVM(vm *vmmModels.Vm) (facade.TaskWaiter[vmmModels.Vm], error) {
	taskRef, err := CallAPI[*vmmModels.CreateVmApiResponse, v4VmmConfig.TaskReference](
		f.client.VmApiInstance.CreateVm(vm),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create VM: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for created VM")
	}

	return NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, f.GetVM), nil
}

func (f *FacadeV4Client) UpdateVM(uuid string, vm *vmmModels.Vm) (facade.TaskWaiter[vmmModels.Vm], error) {
	currentVM, args, err := GetEntityAndEtag(
		f.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get VM for deletion: %w", err)
	}

	vm = CopyEtag(currentVM, vm).(*vmmModels.Vm)

	taskRef, err := CallAPI[*vmmModels.UpdateVmApiResponse, v4VmmConfig.TaskReference](
		f.client.VmApiInstance.UpdateVmById(&uuid, vm, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update VM: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for updated VM")
	}

	waiter := NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, f.GetVM)
	return waiter, nil
}

func (f *FacadeV4Client) DeleteVM(uuid string) (facade.TaskWaiter[facade.NoEntity], error) {
	_, args, err := GetEntityAndEtag(
		f.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get VM for deletion: %w", err)
	}

	taskRef, err := CallAPI[*vmmModels.DeleteVmApiResponse, v4VmmConfig.TaskReference](
		f.client.VmApiInstance.DeleteVmById(&uuid, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete VM: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for deleted VM")
	}

	waiter := NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, facade.NoEntityGetter)
	return waiter, nil
}

func (f *FacadeV4Client) PowerOnVM(uuid string) (facade.TaskWaiter[vmmModels.Vm], error) {
	_, args, err := GetEntityAndEtag(
		f.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get VM for deletion: %w", err)
	}

	taskRef, err := CallAPI[*vmmModels.PowerOnVmApiResponse, v4VmmConfig.TaskReference](
		f.client.VmApiInstance.PowerOnVm(&uuid, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to power on VM: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for powered on VM")
	}

	waiter := NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, f.GetVM)
	return waiter, nil
}

func (f *FacadeV4Client) PowerOffVM(uuid string) (facade.TaskWaiter[vmmModels.Vm], error) {
	_, args, err := GetEntityAndEtag(
		f.client.VmApiInstance.GetVmById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get VM for deletion: %w", err)
	}

	taskRef, err := CallAPI[*vmmModels.PowerOffVmApiResponse, v4VmmConfig.TaskReference](
		f.client.VmApiInstance.PowerOffVm(&uuid, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to power off VM: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for powered off VM")
	}

	waiter := NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, f.GetVM)
	return waiter, nil
}
