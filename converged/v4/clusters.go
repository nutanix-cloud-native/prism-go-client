package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	clusterModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
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
			return s.client.ClustersApiInstance.GetClusterById(&uuid, nil)
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
			return s.client.ClustersApiInstance.ListClusters(reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy, reqParams.Apply, reqParams.Expand, reqParams.Select)
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
			return s.client.ClustersApiInstance.ListClusters(reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy, reqParams.Apply, reqParams.Expand, reqParams.Select)
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
		if reqParams.Apply != nil || reqParams.Expand != nil || reqParams.Select != nil {
			return nil, errors.New("apply, expand and select are not supported")
		}
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

	// Check if unsupported OData options are provided
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}
	if reqParams != nil {
		if reqParams.Apply != nil || reqParams.Expand != nil || reqParams.Select != nil {
			return nil, errors.New("apply, expand and select are not supported")
		}
	}

	return GenericListEntities[*clusterModels.ListPhysicalGpuProfilesApiResponse, clusterModels.PhysicalGpuProfile](
		func(reqParams *V4ODataParams) (*clusterModels.ListPhysicalGpuProfilesApiResponse, error) {
			return s.client.ClustersApiInstance.ListPhysicalGpuProfiles(&clusterUuid, reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy)
		},
		opts,
		"physical GPU profiles",
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
			return s.client.ClustersApiInstance.GetHostById(&clusterUuid, &hostId)
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
			return nil, errors.New("apply, expand and select are not supported")
		}
	}

	return GenericListEntities[*clusterModels.ListHostsByClusterIdApiResponse, clusterModels.Host](
		func(reqParams *V4ODataParams) (*clusterModels.ListHostsByClusterIdApiResponse, error) {
			return s.client.ClustersApiInstance.ListHostsByClusterId(
				&clusterUuid,
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Apply,
				reqParams.Select,
			)
		},
		opts,
		"cluster hosts",
	)
}

// NewIterator returns an iterator for listing cluster hosts.
func (s *ClustersService) NewClusterHostsIterator(ctx context.Context, clusterUuid string, opts ...converged.ODataOption) converged.Iterator[clusterModels.Host] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*clusterModels.ListHostsByClusterIdApiResponse, clusterModels.Host](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*clusterModels.ListHostsByClusterIdApiResponse, error) {
			return s.client.ClustersApiInstance.ListHostsByClusterId(
				&clusterUuid,
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Apply,
				reqParams.Select,
			)
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
			return nil, errors.New("apply, expand and select are not supported")
		}
	}

	return GenericListEntities[*clusterModels.ListHostsApiResponse, clusterModels.Host](
		func(reqParams *V4ODataParams) (*clusterModels.ListHostsApiResponse, error) {
			return s.client.ClustersApiInstance.ListHosts(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Apply,
				reqParams.Select,
			)
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
			return s.client.ClustersApiInstance.ListHosts(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Apply,
				reqParams.Select,
			)
		},
		opts,
		"hosts",
	)
}
