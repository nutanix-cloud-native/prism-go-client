package v4

import (
	"github.com/nutanix-cloud-native/prism-go-client/facade"
	subnetModels "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	subnetModelsError "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/error"
)

// GetSubnet returns the subnet for the given UUID.
func (f *FacadeV4Client) GetSubnet(uuid string) (*subnetModels.Subnet, error) {
	return CommonGetEntity[*subnetModels.GetSubnetApiResponse, subnetModels.Subnet, *subnetModelsError.OneOfErrorResponseError](
		func() (*subnetModels.GetSubnetApiResponse, error) {
			return f.client.SubnetsApiInstance.GetSubnetById(&uuid)
		},
		"subnet",
	)
}

// ListSubnets returns a list of subnets.
func (f *FacadeV4Client) ListSubnets(opts ...facade.ODataOption) ([]subnetModels.Subnet, error) {
	return CommonListEntities[*subnetModels.ListSubnetsApiResponse, subnetModels.Subnet, *subnetModelsError.OneOfErrorResponseError](
		func(reqParams *V4ODataParams) (*subnetModels.ListSubnetsApiResponse, error) {
			return f.client.SubnetsApiInstance.ListSubnets(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		"subnets",
	)
}

// ListAllSubnets returns all subnets without pagination.
func (f *FacadeV4Client) ListAllSubnets(filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]subnetModels.Subnet, error) {
	myParams := &V4ODataParams{
		Filter:  filterParam,
		OrderBy: orderbyParam,
		Expand:  expandParam,
		Select:  selectParam,
	}

	return CommonListAllEntities[*subnetModels.ListSubnetsApiResponse, subnetModels.Subnet, *subnetModelsError.OneOfErrorResponseError](
		func(reqParams *V4ODataParams) (*subnetModels.ListSubnetsApiResponse, error) {
			return f.client.SubnetsApiInstance.ListSubnets(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		myParams,
		"subnets",
	)
}

// GetListIteratorSubnets returns an iterator for listing subnets.
func (f *FacadeV4Client) GetListIteratorSubnets(opts ...facade.ODataOption) facade.ODataListIterator[subnetModels.Subnet] {
	return CommonGetListIterator[*subnetModels.ListSubnetsApiResponse, subnetModels.Subnet, *subnetModelsError.OneOfErrorResponseError](
		f,
		func(reqParams *V4ODataParams) (*subnetModels.ListSubnetsApiResponse, error) {
			return f.client.SubnetsApiInstance.ListSubnets(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		"subnets",
	)
}
