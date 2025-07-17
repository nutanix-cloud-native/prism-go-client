package v4

import (
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

func (f *FacadeV4Client) GetVM(uuid string) (*vmmModels.Vm, error) {
	return nil, fmt.Errorf("GetVM not implemented in FacadeV4Client")
}

func (f *FacadeV4Client) ListVMs(opts ...facade.ODataOption) ([]vmmModels.Vm, error) {
	return nil, fmt.Errorf("ListVMs not implemented in FacadeV4Client")
}

func (f *FacadeV4Client) ListAllVMs(filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]vmmModels.Vm, error) {
	return nil, fmt.Errorf("ListAllVMs not implemented in FacadeV4Client")
}

func (f *FacadeV4Client) GetListIteratorVMs(opts ...facade.ODataOption) (facade.ODataListIterator[*vmmModels.Vm], error) {
	return nil, fmt.Errorf("GetListIteratorVMs not implemented in FacadeV4Client")
}

func (f *FacadeV4Client) CreateVM(vm *vmmModels.Vm) (*vmmModels.Vm, error) {
	return nil, fmt.Errorf("CreateVM not implemented in FacadeV4Client")
}

func (f *FacadeV4Client) UpdateVM(uuid string, vm *vmmModels.Vm) (*vmmModels.Vm, error) {
	return nil, fmt.Errorf("UpdateVM not implemented in FacadeV4Client")
}

func (f *FacadeV4Client) DeleteVM(uuid string) error {
	return fmt.Errorf("DeleteVM not implemented in FacadeV4Client")
}
