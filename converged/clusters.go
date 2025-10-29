package converged

import "context"

type Clusters[
	Cluster,
	VirtualGpuProfile,
	PhysicalGpuProfile,
	ClusterListParams,
	GpuProfileListParams any] interface {
	// Getter is the interface for Get operations.
	Getter[Cluster]

	// Lister is the interface for List operations.
	Lister[Cluster, ClusterListParams]

	// ListClusterVirtualGPUs returns the virtual GPU configuration for the given cluster UUID.
	ListClusterVirtualGPUs(ctx context.Context, clusterUuid string, opts ...ODataOption[GpuProfileListParams]) ([]VirtualGpuProfile, error)

	// ListClusterPhysicalGPUs returns the physical GPU configuration for the given cluster UUID.
	ListClusterPhysicalGPUs(ctx context.Context, clusterUuid string, opts ...ODataOption[GpuProfileListParams]) ([]PhysicalGpuProfile, error)
}
