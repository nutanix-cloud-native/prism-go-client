package converged

import (
	"context"
)

// Subnets defines the interface for subnet operations.
// It supports get/list for subnets and IP reserve/unreserve operations.
type Subnets[Subnet, TaskReference any] interface {
	Getter[Subnet]
	Lister[Subnet]

	// ReserveIpsBySubnetId reserves IP addresses on a subnet.
	// Returns a TaskReference that can be used to track the async operation.
	// spec should be of type *subnetModels.IpReserveSpec (for v4 implementation).
	ReserveIpsBySubnetId(ctx context.Context, subnetExtId string, spec any) (*TaskReference, error)

	// UnreserveIpsBySubnetId unreserves IP addresses on a subnet.
	// Returns a TaskReference that can be used to track the async operation.
	// spec should be of type *subnetModels.IpUnreserveSpec (for v4 implementation).
	UnreserveIpsBySubnetId(ctx context.Context, subnetExtId string, spec any) (*TaskReference, error)

	// ListReservedIpsBySubnetId returns the list of reserved IP addresses on a subnet.
	ListReservedIpsBySubnetId(ctx context.Context, subnetExtId string, opts ...ODataOption) (any, error)
}
