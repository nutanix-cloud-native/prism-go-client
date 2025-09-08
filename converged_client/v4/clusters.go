package v4

import (
	"context"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	clustermgmtModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// ClustersService provides implementation for all Clusters interface methods.
type ClustersService struct {
	client *v4prismGoClient.Client
}

// NewClustersService creates a new ClustersService instance.
func NewClustersService(client *v4prismGoClient.Client) *ClustersService {
	return &ClustersService{
		client: client,
	}
}

// Get returns the cluster for the given UUID.
func (s *ClustersService) Get(ctx context.Context, uuid string) (*clustermgmtModels.Cluster, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	return GenericGetEntity[
		*clustermgmtModels.GetClusterApiResponse,
		clustermgmtModels.Cluster,
	](
		func() (*clustermgmtModels.GetClusterApiResponse, error) {
			return s.client.ClustersApiInstance.GetClusterById(&uuid, nil, nil)
		},
		"cluster",
	)
}

// List returns a list of clusters.
func (s *ClustersService) List(ctx context.Context, opts ...converged.ODataOption) ([]clustermgmtModels.Cluster, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	return GenericListEntities[
		*clustermgmtModels.ListClustersApiResponse,
		clustermgmtModels.Cluster,
	](
		func(reqParams *V4ODataParams) (*clustermgmtModels.ListClustersApiResponse, error) {
			return s.client.ClustersApiInstance.ListClusters(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
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
func (s *ClustersService) ListAll(ctx context.Context, opts ...converged.ODataOption) ([]clustermgmtModels.Cluster, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	return GenericListAllEntities[
		*clustermgmtModels.ListClustersApiResponse,
		clustermgmtModels.Cluster,
	](
		func(reqParams *V4ODataParams) (*clustermgmtModels.ListClustersApiResponse, error) {
			return s.client.ClustersApiInstance.ListClusters(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
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
func (s *ClustersService) NewIterator(opts ...converged.ODataOption) converged.Iterator[clustermgmtModels.Cluster] {
	if s.client == nil {
		return nil
	}

	return NewIterator[
		*clustermgmtModels.ListClustersApiResponse,
		clustermgmtModels.Cluster,
	](
		context.Background(),
		s.client,
		func(ctx context.Context, reqParams *V4ODataParams) (*clustermgmtModels.ListClustersApiResponse, error) {
			return s.client.ClustersApiInstance.ListClusters(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
				nil,
			)
		},
		opts...,
	)
}
