package converged

import (
	"context"
)

type Subnets[Subnet, TaskReference any] interface {
	// Getter is the interface for Get operations.
	Getter[Subnet]

	// Lister is the interface for List operations.
	Lister[Subnet]

	// ReserveIpsBySubnetId reserves IP addresses on a subnet.
	// Returns a TaskReference that can be used to track the async operation.
	// spec should be of type *subnetModels.IpReserveSpec (for v4 implementation).
	ReserveIpsBySubnetId(ctx context.Context, subnetExtId string, spec any) (*TaskReference, error)

	// UnreserveIpsBySubnetId unreserves IP addresses on a subnet.
	// Returns a TaskReference that can be used to track the async operation.
	// spec should be of type *subnetModels.IpUnreserveSpec (for v4 implementation).
	UnreserveIpsBySubnetId(ctx context.Context, subnetExtId string, spec any) (*TaskReference, error)

	// ListReservedIpsBySubnetId fetches a list of reserved IP addresses on a subnet.
	ListReservedIpsBySubnetId(ctx context.Context, subnetExtId string, opts ...ODataOption) (any, error)

	// Additional methods can be added here as needed.
}
