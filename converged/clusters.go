package converged

import "context"

// Clusters defines the interface for cluster operations.
// It supports get/list for clusters and cluster-scoped resources such as
// virtual/physical GPU profiles and hosts.
type Clusters[
	Cluster,
	VirtualGpuProfile,
	PhysicalGpuProfile,
	Host any] interface {
	Getter[Cluster]
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

	// ListAllHosts returns all hosts across clusters.
	ListAllHosts(ctx context.Context, opts ...ODataOption) ([]Host, error)

	// NewAllHostsIterator returns an iterator for listing all hosts.
	NewAllHostsIterator(ctx context.Context, opts ...ODataOption) Iterator[Host]
}
