package v4

import (
	"context"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	clusterModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// ClustersService provides default "not implemented" implementation for all Clusters interface methods.
type ClustersService struct {
	client *v4prismGoClient.Client
}

// NewClustersService creates a new ClustersService instance.
func NewClustersService(client *v4prismGoClient.Client) *ClustersService {
	return &ClustersService{client: client}
}

// Get returns the cluster for the given UUID.
func (s *ClustersService) Get(ctx context.Context, uuid string) (*clusterModels.Cluster, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	if uuid == "" {
		return nil, fmt.Errorf("uuid is required")
	}
	return GenericGetEntity[*clusterModels.GetClusterApiResponse, clusterModels.Cluster](
		func() (*clusterModels.GetClusterApiResponse, error) {
			return s.client.ClustersApiInstance.GetClusterById(&uuid, nil)
		},
		"cluster",
	)
}

// List returns a list of clusters.
func (s *ClustersService) List(ctx context.Context, opts ...converged.ODataOption) ([]clusterModels.Cluster, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	return GenericListEntities[*clusterModels.ListClustersApiResponse, clusterModels.Cluster](
		func(reqParams *V4ODataParams) (*clusterModels.ListClustersApiResponse, error) {
			return s.client.ClustersApiInstance.ListClusters(reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy, reqParams.Apply, reqParams.Expand, reqParams.Select)
		},
		opts,
		"clusters",
	)
}

// NewIterator returns an iterator for listing clusters.
func (s *ClustersService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[clusterModels.Cluster] {
	if s.client == nil {
		return nil
	}
	return GenericNewIterator[*clusterModels.ListClustersApiResponse, clusterModels.Cluster](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*clusterModels.ListClustersApiResponse, error) {
			return s.client.ClustersApiInstance.ListClusters(reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy, reqParams.Apply, reqParams.Expand, reqParams.Select)
		},
		opts,
		"clusters",
	)
}

// ListClusterVirtualGPUs returns the virtual GPU configuration for the given cluster UUID.
func (s *ClustersService) ListClusterVirtualGPUs(ctx context.Context, clusterUuid string, opts ...converged.ODataOption) ([]clusterModels.VirtualGpuProfile, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	if clusterUuid == "" {
		return nil, fmt.Errorf("clusterUuid is required")
	}
	return GenericListEntities[*clusterModels.ListVirtualGpuProfilesApiResponse, clusterModels.VirtualGpuProfile](
		func(reqParams *V4ODataParams) (*clusterModels.ListVirtualGpuProfilesApiResponse, error) {
			return s.client.ClustersApiInstance.ListVirtualGpuProfiles(&clusterUuid, reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy)
		},
		opts,
		"virtual GPU profiles",
	)
}

// ListClusterPhysicalGPUs returns the physical GPU configuration for the given cluster UUID.
func (s *ClustersService) ListClusterPhysicalGPUs(ctx context.Context, clusterUuid string, opts ...converged.ODataOption) ([]clusterModels.PhysicalGpuProfile, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	if clusterUuid == "" {
		return nil, fmt.Errorf("clusterUuid is required")
	}
	return GenericListEntities[*clusterModels.ListPhysicalGpuProfilesApiResponse, clusterModels.PhysicalGpuProfile](
		func(reqParams *V4ODataParams) (*clusterModels.ListPhysicalGpuProfilesApiResponse, error) {
			return s.client.ClustersApiInstance.ListPhysicalGpuProfiles(&clusterUuid, reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy)
		},
		opts,
		"physical GPU profiles",
	)
}
