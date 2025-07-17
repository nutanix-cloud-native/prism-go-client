package v4

import (
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	clustersModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// GetCluster returns the cluster for the given UUID.
func (f *FacadeV4Client) GetCluster(uuid string) (*clustersModels.Cluster, error) {
	return nil, fmt.Errorf("GetCluster not implemented in FacadeV4Client")
}

// ListClusters returns a list of clusters.
func (f *FacadeV4Client) ListClusters(opts ...facade.ODataOption) ([]clustersModels.Cluster, error) {
	return nil, fmt.Errorf("ListClusters not implemented in FacadeV4Client")
}

// ListAllClusters returns all clusters without pagination.
func (f *FacadeV4Client) ListAllClusters(filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]clustersModels.Cluster, error) {
	return nil, fmt.Errorf("ListAllClusters not implemented in FacadeV4Client")
}

// GetListIteratorClusters returns an iterator for listing clusters.
func (f *FacadeV4Client) GetListIteratorClusters(opts ...facade.ODataOption) (facade.ODataListIterator[*clustersModels.Cluster], error) {
	return nil, fmt.Errorf("GetListIteratorClusters not implemented in FacadeV4Client")
}

// ListClusterVirtualGPUs returns the virtual GPU configuration for the given cluster UUID.
func (f *FacadeV4Client) ListClusterVirtualGPUs(clusterUuid string, opts ...facade.ODataOption) ([]*clustersModels.VirtualGpuProfile, error) {
	return nil, fmt.Errorf("ListClusterVirtualGPUs not implemented in FacadeV4Client")
}

// ListClusterPhysicalGPUs returns the physical GPU configuration for the given cluster UUID.
func (f *FacadeV4Client) ListClusterPhysicalGPUs(clusterUuid string, opts ...facade.ODataOption) ([]*clustersModels.PhysicalGpuProfile, error) {
	return nil, fmt.Errorf("ListClusterPhysicalGPUs not implemented in FacadeV4Client")
}
