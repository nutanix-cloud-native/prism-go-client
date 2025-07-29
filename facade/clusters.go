package facade

import (
	clusterModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

type ClustersFacadeV4 interface {
	// GetCluster returns the cluster for the given UUID.
	GetCluster(uuid string) (*clusterModels.Cluster, error)

	// ListClusters returns a list of clusters.
	ListClusters(opts ...ODataOption) ([]clusterModels.Cluster, error)

	// ListAllClusters returns all clusters without pagination.
	ListAllClusters(filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]clusterModels.Cluster, error)

	// GetListIteratorClusters returns an iterator for listing clusters.
	GetListIteratorClusters(opts ...ODataOption) (ODataListIterator[clusterModels.Cluster], error)

	// ListClusterVirtualGPUs returns the virtual GPU configuration for the given cluster UUID.
	ListClusterVirtualGPUs(clusterUuid string, opts ...ODataOption) ([]clusterModels.VirtualGpuProfile, error)

	// ListClusterPhysicalGPUs returns the physical GPU configuration for the given cluster UUID.
	ListClusterPhysicalGPUs(clusterUuid string, opts ...ODataOption) ([]clusterModels.PhysicalGpuProfile, error)
}
