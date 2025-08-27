package facade

import (
	"context"

	clusterModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

type ClustersFacadeV4 interface {
	// GetCluster returns the cluster for the given UUID.
	GetCluster(ctx context.Context, uuid string) (*clusterModels.Cluster, error)

	// ListClusters returns a list of clusters.
	ListClusters(ctx context.Context, opts ...ODataOption) ([]clusterModels.Cluster, error)

	// ListAllClusters returns all clusters without pagination.
	ListAllClusters(ctx context.Context, filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]clusterModels.Cluster, error)

	// GetListIteratorClusters returns an iterator for listing clusters.
	GetListIteratorClusters(ctx context.Context, opts ...ODataOption) ODataListIterator[clusterModels.Cluster]

	// ListClusterVirtualGPUs returns the virtual GPU configuration for the given cluster UUID.
	ListClusterVirtualGPUs(ctx context.Context, clusterUuid string, opts ...ODataOption) ([]clusterModels.VirtualGpuProfile, error)

	// ListClusterPhysicalGPUs returns the physical GPU configuration for the given cluster UUID.
	ListClusterPhysicalGPUs(ctx context.Context, clusterUuid string, opts ...ODataOption) ([]clusterModels.PhysicalGpuProfile, error)
}
