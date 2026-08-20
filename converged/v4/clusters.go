package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	gpuModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/ahv/config"
	clusterModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	clusterRequest "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/clusters"
	gpuProfilesRequest "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/gpuprofiles"
)

// ClustersService provides default "not implemented" implementation for all Clusters interface methods.
type ClustersService struct {
	client   *v4prismGoClient.Client
	entities string
}

// NewClustersService creates a new ClustersService instance.
func NewClustersService(client *v4prismGoClient.Client) *ClustersService {
	return &ClustersService{client: client, entities: "cluster"}
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
			return s.client.ClustersApiInstance.ServiceClient.GetClusterById(ctx, &clusterRequest.GetClusterByIdRequest{
				ExtId: &uuid,
			})
		},
		s.entities,
	)
}

// List returns a list of clusters.
func (s *ClustersService) List(ctx context.Context, opts ...converged.ODataOption) ([]clusterModels.Cluster, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	return GenericListEntities[*clusterModels.ListClustersApiResponse, clusterModels.Cluster](
		func(reqParams *V4ODataParams) (*clusterModels.ListClustersApiResponse, error) {
			return s.client.ClustersApiInstance.ServiceClient.ListClusters(ctx, &clusterRequest.ListClustersRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Apply_:   reqParams.Apply,
				Expand_:  reqParams.Expand,
				Select_:  reqParams.Select,
			})
		},
		opts,
		s.entities,
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
			return s.client.ClustersApiInstance.ServiceClient.ListClusters(ctx, &clusterRequest.ListClustersRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Apply_:   reqParams.Apply,
				Expand_:  reqParams.Expand,
				Select_:  reqParams.Select,
			})
		},
		opts,
		s.entities,
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

	// Check if unsupported OData options are provided
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}
	if reqParams != nil {
		if reqParams.Apply != nil || reqParams.Expand != nil {
			return nil, errors.New("apply and expand are not supported")
		}
	}

	//nolint:staticcheck // Intentionally using deprecated cluster-scoped API for device-level GPU details
	return GenericListEntities[*clusterModels.ListVirtualGpuProfilesApiResponse, clusterModels.VirtualGpuProfile](
		func(reqParams *V4ODataParams) (*clusterModels.ListVirtualGpuProfilesApiResponse, error) {
			return s.client.ClustersApiInstance.ServiceClient.ListVirtualGpuProfiles(ctx, &clusterRequest.ListVirtualGpuProfilesRequest{
				ClusterExtId: &clusterUuid,
				Page_:        reqParams.Page,
				Limit_:       reqParams.Limit,
				Filter_:      reqParams.Filter,
				Orderby_:     reqParams.OrderBy,
			})
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

	// Check if unsupported OData options are provided
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}
	if reqParams != nil {
		if reqParams.Apply != nil || reqParams.Expand != nil {
			return nil, errors.New("apply and expand are not supported")
		}
	}

	//nolint:staticcheck // Intentionally using deprecated cluster-scoped API for device-level GPU details
	return GenericListEntities[*clusterModels.ListPhysicalGpuProfilesApiResponse, clusterModels.PhysicalGpuProfile](
		func(reqParams *V4ODataParams) (*clusterModels.ListPhysicalGpuProfilesApiResponse, error) {
			return s.client.ClustersApiInstance.ServiceClient.ListPhysicalGpuProfiles(ctx, &clusterRequest.ListPhysicalGpuProfilesRequest{
				ClusterExtId: &clusterUuid,
				Page_:        reqParams.Page,
				Limit_:       reqParams.Limit,
				Filter_:      reqParams.Filter,
				Orderby_:     reqParams.OrderBy,
			})
		},
		opts,
		"physical GPU profiles",
	)
}

// ListAHVVirtualGPUProfiles returns AHV virtual GPU profile templates for the given cluster.
// This calls the AHV global endpoint which returns profile templates with configuration name.
func (s *ClustersService) ListAHVVirtualGPUProfiles(ctx context.Context, clusterUuid string, opts ...converged.ODataOption) ([]gpuModels.VirtualGpuProfile, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	if clusterUuid == "" {
		return nil, fmt.Errorf("clusterUuid is required")
	}

	// Check if unsupported OData options are provided
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}
	if reqParams != nil {
		if reqParams.Apply != nil || reqParams.Expand != nil {
			return nil, errors.New("apply and expand are not supported")
		}
	}

	// Build filter for cluster UUID
	clusterFilter := fmt.Sprintf("clusterExtIds/any(c: c eq '%s')", clusterUuid)
	var filter *string
	if reqParams.Filter != nil && *reqParams.Filter != "" {
		combined := fmt.Sprintf("(%s) and (%s)", *reqParams.Filter, clusterFilter)
		filter = &combined
	} else {
		filter = &clusterFilter
	}

	return GenericListEntities[*gpuModels.ListAhvVirtualGpuProfilesApiResponse, gpuModels.VirtualGpuProfile](
		func(reqParams *V4ODataParams) (*gpuModels.ListAhvVirtualGpuProfilesApiResponse, error) {
			return s.client.GpuProfilesApiInstance.ServiceClient.ListAhvVirtualGpuProfiles(ctx, &gpuProfilesRequest.ListAhvVirtualGpuProfilesRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  filter,
				Orderby_: reqParams.OrderBy,
				Select_:  reqParams.Select,
			})
		},
		opts,
		"AHV virtual GPU profiles",
	)
}

// ListAHVPhysicalGPUProfiles returns AHV physical GPU profile templates for the given cluster.
// This calls the AHV global endpoint which returns profile templates with configuration name.
func (s *ClustersService) ListAHVPhysicalGPUProfiles(ctx context.Context, clusterUuid string, opts ...converged.ODataOption) ([]gpuModels.PhysicalGpuProfile, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	if clusterUuid == "" {
		return nil, fmt.Errorf("clusterUuid is required")
	}

	// Check if unsupported OData options are provided
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}
	if reqParams != nil {
		if reqParams.Apply != nil || reqParams.Expand != nil {
			return nil, errors.New("apply and expand are not supported")
		}
	}

	// Build filter for cluster UUID
	clusterFilter := fmt.Sprintf("clusterExtIds/any(c: c eq '%s')", clusterUuid)
	var filter *string
	if reqParams.Filter != nil && *reqParams.Filter != "" {
		combined := fmt.Sprintf("(%s) and (%s)", *reqParams.Filter, clusterFilter)
		filter = &combined
	} else {
		filter = &clusterFilter
	}

	return GenericListEntities[*gpuModels.ListAhvPhysicalGpuProfilesApiResponse, gpuModels.PhysicalGpuProfile](
		func(reqParams *V4ODataParams) (*gpuModels.ListAhvPhysicalGpuProfilesApiResponse, error) {
			return s.client.GpuProfilesApiInstance.ServiceClient.ListAhvPhysicalGpuProfiles(ctx, &gpuProfilesRequest.ListAhvPhysicalGpuProfilesRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  filter,
				Orderby_: reqParams.OrderBy,
				Select_:  reqParams.Select,
			})
		},
		opts,
		"AHV physical GPU profiles",
	)
}

// GetClusterHost returns the host for the given cluster UUID and host id.
func (s *ClustersService) GetClusterHost(ctx context.Context, clusterUuid string, hostId string) (*clusterModels.Host, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	if clusterUuid == "" {
		return nil, fmt.Errorf("clusterUuid is required")
	}
	if hostId == "" {
		return nil, fmt.Errorf("hostId is required")
	}
	return GenericGetEntity[*clusterModels.GetHostApiResponse, clusterModels.Host](
		func() (*clusterModels.GetHostApiResponse, error) {
			return s.client.ClustersApiInstance.ServiceClient.GetHostById(ctx, &clusterRequest.GetHostByIdRequest{
				ClusterExtId: &clusterUuid,
				ExtId:        &hostId,
			})
		},
		"cluster host",
	)
}

// ListClusterHosts returns the hosts for the given cluster UUID.
func (s *ClustersService) ListClusterHosts(ctx context.Context, clusterUuid string, opts ...converged.ODataOption) ([]clusterModels.Host, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	if clusterUuid == "" {
		return nil, fmt.Errorf("clusterUuid is required")
	}

	// Check if unsupported OData options are provided
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}
	if reqParams != nil {
		if reqParams.Expand != nil {
			return nil, errors.New("expand is not supported")
		}
	}

	return GenericListEntities[*clusterModels.ListHostsByClusterIdApiResponse, clusterModels.Host](
		func(reqParams *V4ODataParams) (*clusterModels.ListHostsByClusterIdApiResponse, error) {
			return s.client.ClustersApiInstance.ServiceClient.ListHostsByClusterId(ctx, &clusterRequest.ListHostsByClusterIdRequest{
				ClusterExtId: &clusterUuid,
				Page_:        reqParams.Page,
				Limit_:       reqParams.Limit,
				Filter_:      reqParams.Filter,
				Orderby_:     reqParams.OrderBy,
				Apply_:       reqParams.Apply,
				Select_:      reqParams.Select,
			})
		},
		opts,
		"cluster hosts",
	)
}

// NewClusterHostsIterator returns an iterator for listing cluster hosts.
func (s *ClustersService) NewClusterHostsIterator(ctx context.Context, clusterUuid string, opts ...converged.ODataOption) converged.Iterator[clusterModels.Host] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*clusterModels.ListHostsByClusterIdApiResponse, clusterModels.Host](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*clusterModels.ListHostsByClusterIdApiResponse, error) {
			return s.client.ClustersApiInstance.ServiceClient.ListHostsByClusterId(ctx, &clusterRequest.ListHostsByClusterIdRequest{
				ClusterExtId: &clusterUuid,
				Page_:        reqParams.Page,
				Limit_:       reqParams.Limit,
				Filter_:      reqParams.Filter,
				Orderby_:     reqParams.OrderBy,
				Apply_:       reqParams.Apply,
				Select_:      reqParams.Select,
			})
		},
		opts,
		"cluster hosts",
	)
}

// ListAllHosts returns all hosts.
func (s *ClustersService) ListAllHosts(ctx context.Context, opts ...converged.ODataOption) ([]clusterModels.Host, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is nil")
	}

	// Check if unsupported OData options are provided
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}
	if reqParams != nil {
		if reqParams.Expand != nil {
			return nil, errors.New("expand is not supported")
		}
	}

	return GenericListEntities[*clusterModels.ListHostsApiResponse, clusterModels.Host](
		func(reqParams *V4ODataParams) (*clusterModels.ListHostsApiResponse, error) {
			return s.client.ClustersApiInstance.ServiceClient.ListHosts(ctx, &clusterRequest.ListHostsRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Apply_:   reqParams.Apply,
				Select_:  reqParams.Select,
			})
		},
		nil,
		"hosts",
	)
}

// NewAllHostsIterator returns an iterator for listing all hosts.
func (s *ClustersService) NewAllHostsIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[clusterModels.Host] {
	if s.client == nil {
		return nil
	}
	return GenericNewIterator[*clusterModels.ListHostsApiResponse, clusterModels.Host](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*clusterModels.ListHostsApiResponse, error) {
			return s.client.ClustersApiInstance.ServiceClient.ListHosts(ctx, &clusterRequest.ListHostsRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Apply_:   reqParams.Apply,
				Select_:  reqParams.Select,
			})
		},
		opts,
		"hosts",
	)
}
