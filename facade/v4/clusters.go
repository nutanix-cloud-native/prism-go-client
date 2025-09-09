package v4

import (
	"context"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	clustersModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// GetCluster returns the cluster for the given UUID.
func (f *FacadeV4Client) GetCluster(ctx context.Context, uuid string) (*clustersModels.Cluster, error) {
	return CommonGetEntity[*clustersModels.GetClusterApiResponse, clustersModels.Cluster](
		func() (*clustersModels.GetClusterApiResponse, error) {
			return f.client.ClustersApiInstance.GetClusterById(&uuid, nil)
		},
		"cluster",
	)
}

// ListClusters returns a list of clusters.
func (f *FacadeV4Client) ListClusters(ctx context.Context, opts ...facade.ODataOption) ([]clustersModels.Cluster, error) {
	return CommonListEntities[*clustersModels.ListClustersApiResponse, clustersModels.Cluster](
		func(reqParams *V4ODataParams) (*clustersModels.ListClustersApiResponse, error) {
			return f.client.ClustersApiInstance.ListClusters(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Apply,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		"clusters",
	)
}

// ListAllClusters returns all clusters without pagination.
func (f *FacadeV4Client) ListAllClusters(ctx context.Context, filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]clustersModels.Cluster, error) {
	reqParams := &V4ODataParams{
		Filter:  filterParam,
		OrderBy: orderbyParam,
		Expand:  expandParam,
		Select:  selectParam,
	}

	return CommonListAllEntities[*clustersModels.ListClustersApiResponse, clustersModels.Cluster](
		func(reqParams *V4ODataParams) (*clustersModels.ListClustersApiResponse, error) {
			return f.client.ClustersApiInstance.ListClusters(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Apply,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		reqParams,
		"clusters",
	)
}

// GetListIteratorClusters returns an iterator for listing clusters.
func (f *FacadeV4Client) GetListIteratorClusters(ctx context.Context, opts ...facade.ODataOption) facade.ODataListIterator[clustersModels.Cluster] {
	return CommonGetListIterator[*clustersModels.ListClustersApiResponse, clustersModels.Cluster](
		ctx,
		f,
		func(reqParams *V4ODataParams) (*clustersModels.ListClustersApiResponse, error) {
			return f.client.ClustersApiInstance.ListClusters(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Apply,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		"clusters",
	)
}

// ListClusterVirtualGPUs returns the virtual GPU configuration for the given cluster UUID.
func (f *FacadeV4Client) ListClusterVirtualGPUs(ctx context.Context, clusterUuid string, opts ...facade.ODataOption) ([]clustersModels.VirtualGpuProfile, error) {
	return CommonListEntities[*clustersModels.ListVirtualGpuProfilesApiResponse, clustersModels.VirtualGpuProfile](
		func(reqParams *V4ODataParams) (*clustersModels.ListVirtualGpuProfilesApiResponse, error) {
			return f.client.ClustersApiInstance.ListVirtualGpuProfiles(
				&clusterUuid,
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
			)
		},
		opts,
		"virtual GPU profiles",
	)
}

// ListClusterPhysicalGPUs returns the physical GPU configuration for the given cluster UUID.
func (f *FacadeV4Client) ListClusterPhysicalGPUs(ctx context.Context, clusterUuid string, opts ...facade.ODataOption) ([]clustersModels.PhysicalGpuProfile, error) {
	return CommonListEntities[*clustersModels.ListPhysicalGpuProfilesApiResponse, clustersModels.PhysicalGpuProfile](
		func(reqParams *V4ODataParams) (*clustersModels.ListPhysicalGpuProfilesApiResponse, error) {
			return f.client.ClustersApiInstance.ListPhysicalGpuProfiles(
				&clusterUuid,
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
			)
		},
		opts,
		"physical GPU profiles",
	)
}
