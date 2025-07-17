package v4

import (
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	subnetModels "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// GetSubnet returns the subnet for the given UUID.
func (f *FacadeV4Client) GetSubnet(uuid string) (*subnetModels.Subnet, error) {
	return nil, fmt.Errorf("GetSubnet not implemented in FacadeV4Client")
}

// ListSubnets returns a list of subnets.
func (f *FacadeV4Client) ListSubnets(opts ...facade.ODataOption) ([]subnetModels.Subnet, error) {
	return nil, fmt.Errorf("ListSubnets not implemented in FacadeV4Client")
}

// ListAllSubnets returns all subnets without pagination.
func (f *FacadeV4Client) ListAllSubnets(filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]subnetModels.Subnet, error) {
	return nil, fmt.Errorf("ListAllSubnets not implemented in FacadeV4Client")
}

// GetListIteratorSubnets returns an iterator for listing subnets.
func (f *FacadeV4Client) GetListIteratorSubnets(opts ...facade.ODataOption) (facade.ODataListIterator[*subnetModels.Subnet], error) {
	return nil, fmt.Errorf("GetListIteratorSubnets not implemented in FacadeV4Client")
}
