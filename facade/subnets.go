package facade

import (
	"context"

	subnetModels "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

type SubnetsFacadeV4 interface {
	// GetSubnet returns the subnet for the given UUID.
	GetSubnet(ctx context.Context, uuid string) (*subnetModels.Subnet, error)

	// ListSubnets returns a list of subnets.
	ListSubnets(ctx context.Context, opts ...ODataOption) ([]subnetModels.Subnet, error)

	// ListAllSubnets returns all subnets without pagination.
	ListAllSubnets(ctx context.Context, filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]subnetModels.Subnet, error)

	// GetListIteratorSubnets returns an iterator for listing subnets.
	GetListIteratorSubnets(ctx context.Context, opts ...ODataOption) ODataListIterator[subnetModels.Subnet]
}
