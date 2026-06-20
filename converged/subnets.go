package converged

import (
	"context"
)

type Subnets[Subnet, TaskReference any] interface {
	// Getter is the interface for Get operations.
	Getter[Subnet]

	// Lister is the interface for List operations.
	Lister[Subnet]

	// ReserveIpsBySubnetId reserves IP addresses on a subnet and waits for completion.
	// Returns the reserved IP addresses from the completed task.
	// spec should be of type *subnetModels.IpReserveSpec (for v4 implementation).
	ReserveIpsBySubnetId(ctx context.Context, subnetExtId string, spec any) ([]string, error)

	// ReserveIpsBySubnetIdAsync reserves IP addresses on a subnet and returns an Operation that can wait for completion.
	// spec should be of type *subnetModels.IpReserveSpec (for v4 implementation).
	ReserveIpsBySubnetIdAsync(ctx context.Context, subnetExtId string, spec any) (Operation[NoEntity], error)

	// UnreserveIpsBySubnetId unreserves IP addresses on a subnet and waits for completion.
	// Returns the IP addresses that were unreserved, as reported by the completed task.
	// This is primarily useful when unreserving by client context, where the caller does
	// not know up front which IPs the server will release. The slice may be empty if the
	// task does not report the unreserved IPs.
	// spec should be of type *subnetModels.IpUnreserveSpec (for v4 implementation).
	UnreserveIpsBySubnetId(ctx context.Context, subnetExtId string, spec any) ([]string, error)

	// UnreserveIpsBySubnetIdAsync unreserves IP addresses on a subnet and returns an Operation that can wait for completion.
	// spec should be of type *subnetModels.IpUnreserveSpec (for v4 implementation).
	UnreserveIpsBySubnetIdAsync(ctx context.Context, subnetExtId string, spec any) (Operation[NoEntity], error)

	// ListReservedIpsBySubnetId fetches a list of reserved IP addresses on a subnet.
	ListReservedIpsBySubnetId(ctx context.Context, subnetExtId string, opts ...ODataOption) (any, error)

	// Additional methods can be added here as needed.
}
