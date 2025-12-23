package converged

import "context"

type Clusters[
	Cluster,
	VirtualGpuProfile,
	PhysicalGpuProfile,
	Host any] interface {
	// Getter is the interface for Get operations.
	Getter[Cluster]

	// Lister is the interface for List operations.
	Lister[Cluster]

	// ListClusterVirtualGPUs returns the virtual GPU configuration for the given cluster UUID.
	ListClusterVirtualGPUs(ctx context.Context, clusterUuid string, opts ...ODataOption) ([]VirtualGpuProfile, error)

	// ListClusterPhysicalGPUs returns the physical GPU configuration for the given cluster UUID.
	ListClusterPhysicalGPUs(ctx context.Context, clusterUuid string, opts ...ODataOption) ([]PhysicalGpuProfile, error)

	// GetClusterHost returns the host for the given cluster UUID and host id.
	GetClusterHost(ctx context.Context, clusterUuid string, hostId string) (*Host, error)

	// ListClusterHosts returns the hosts for the given cluster UUID.
	ListClusterHosts(ctx context.Context, clusterUuid string, opts ...ODataOption) ([]Host, error)

	// NewClusterHostsIterator returns an iterator for listing cluster hosts.
	NewClusterHostsIterator(ctx context.Context, clusterUuid string, opts ...ODataOption) Iterator[Host]

	// ListAllHosts returns all hosts.
	ListAllHosts(ctx context.Context, opts ...ODataOption) ([]Host, error)

	// NewAllHostsIterator returns an iterator for listing all hosts.
	NewAllHostsIterator(ctx context.Context, opts ...ODataOption) Iterator[Host]
}
