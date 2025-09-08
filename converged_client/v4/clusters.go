package v4

import (
	"context"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	clusterModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// ClustersService provides implementation for all Clusters interface methods.
type ClustersService struct {
	client *Client
}

// NewClustersService creates a new ClustersService instance.
func NewClustersService() *ClustersService {
	return &ClustersService{}
}

// NewClustersServiceWithClient creates a new ClustersService instance with a client.
func NewClustersServiceWithClient(client *Client) *ClustersService {
	return &ClustersService{client: client}
}

// Get returns the cluster for the given UUID.
func (s *ClustersService) Get(ctx context.Context, uuid string) (*clusterModels.Cluster, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client not initialized")
	}

	return GenericGetEntity[
		*clusterModels.GetClusterApiResponse,
		clusterModels.Cluster,
	](
		func() (*clusterModels.GetClusterApiResponse, error) {
			return s.client.client.ClustersApiInstance.GetClusterById(&uuid, nil)
		},
		"cluster",
	)
}

// List returns a list of clusters.
func (s *ClustersService) List(ctx context.Context, opts ...converged.ODataOption) ([]clusterModels.Cluster, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client not initialized")
	}

	return GenericListEntities[
		*clusterModels.ListClustersApiResponse,
		clusterModels.Cluster,
	](
		func(reqParams *V4ODataParams) (*clusterModels.ListClustersApiResponse, error) {
			return s.client.client.ClustersApiInstance.ListClusters(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Apply,
				reqParams.Expand,
				reqParams.Select,
				nil,
			)
		},
		opts,
		"clusters",
	)
}

// ListAll returns all clusters without pagination.
func (s *ClustersService) ListAll(ctx context.Context, opts ...converged.ODataOption) ([]clusterModels.Cluster, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client not initialized")
	}

	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	return GenericListAllEntities[
		*clusterModels.ListClustersApiResponse,
		clusterModels.Cluster,
	](
		func(reqParams *V4ODataParams) (*clusterModels.ListClustersApiResponse, error) {
			return s.client.client.ClustersApiInstance.ListClusters(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Apply,
				reqParams.Expand,
				reqParams.Select,
				nil,
			)
		},
		reqParams,
		"clusters",
	)
}

// NewIterator returns an iterator for listing clusters.
func (s *ClustersService) NewIterator(opts ...converged.ODataOption) converged.Iterator[clusterModels.Cluster] {
	if s.client == nil {
		return func(yield func(clusterModels.Cluster, error) bool) {}
	}

	ctx := context.Background()
	return GenericNewIterator[
		*clusterModels.ListClustersApiResponse,
		clusterModels.Cluster,
	](
		ctx,
		s.client,
		func(ctx context.Context, reqParams *V4ODataParams) (*clusterModels.ListClustersApiResponse, error) {
			return s.client.client.ClustersApiInstance.ListClusters(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Apply,
				reqParams.Expand,
				reqParams.Select,
				nil,
			)
		},
		opts,
		"clusters",
	)
}

// ListClusterVirtualGPUs returns the virtual GPU configuration for the given cluster UUID.
func (s *ClustersService) ListClusterVirtualGPUs(ctx context.Context, clusterUuid string, opts ...converged.ODataOption) ([]clusterModels.VirtualGpuProfile, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client not initialized")
	}

	return GenericListEntities[
		*clusterModels.ListVirtualGpuProfilesApiResponse,
		clusterModels.VirtualGpuProfile,
	](
		func(reqParams *V4ODataParams) (*clusterModels.ListVirtualGpuProfilesApiResponse, error) {
			return s.client.client.ClustersApiInstance.ListVirtualGpuProfiles(
				&clusterUuid,
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				nil,
			)
		},
		opts,
		"virtual GPU profiles",
	)
}

// ListClusterPhysicalGPUs returns the physical GPU configuration for the given cluster UUID.
func (s *ClustersService) ListClusterPhysicalGPUs(ctx context.Context, clusterUuid string, opts ...converged.ODataOption) ([]clusterModels.PhysicalGpuProfile, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client not initialized")
	}

	return GenericListEntities[
		*clusterModels.ListPhysicalGpuProfilesApiResponse,
		clusterModels.PhysicalGpuProfile,
	](
		func(reqParams *V4ODataParams) (*clusterModels.ListPhysicalGpuProfilesApiResponse, error) {
			return s.client.client.ClustersApiInstance.ListPhysicalGpuProfiles(
				&clusterUuid,
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				nil,
			)
		},
		opts,
		"physical GPU profiles",
	)
}
