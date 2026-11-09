package converged

import "context"

// FloatingIps defines the interface for Floating IP API operations.
type FloatingIps[FloatingIp any] interface {
	Getter[FloatingIp]
	Lister[FloatingIp]
	Creator[FloatingIp]
	Updater[FloatingIp]
	Deleter[FloatingIp]
	AsyncCreator[FloatingIp]
	AsyncUpdater[FloatingIp]
	AsyncDeleter[FloatingIp]

	// AssignFloatingIp associates a Floating IP with a VM NIC, private IP, or
	// load balancer session and waits for the task to complete.
	// association should be a VmNicAssociation, PrivateIpAssociation, or
	// LoadBalancerSessionAssociation (value or pointer) for the v4 implementation.
	AssignFloatingIp(ctx context.Context, uuid string, association any) (*FloatingIp, error)

	// AssignFloatingIpAsync associates a Floating IP asynchronously.
	AssignFloatingIpAsync(ctx context.Context, uuid string, association any) (Operation[FloatingIp], error)

	// DeassignFloatingIp clears the Floating IP association and waits for the
	// task to complete.
	DeassignFloatingIp(ctx context.Context, uuid string) (*FloatingIp, error)

	// DeassignFloatingIpAsync clears the Floating IP association asynchronously.
	DeassignFloatingIpAsync(ctx context.Context, uuid string) (Operation[FloatingIp], error)
}
