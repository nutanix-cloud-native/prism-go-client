package v4

import (
	"context"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
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
	return nil, fmt.Errorf("not implemented")
}

// List returns a list of clusters.
func (s *ClustersService) List(ctx context.Context, opts ...converged.ODataOption) ([]clusterModels.Cluster, error) {
	return nil, fmt.Errorf("not implemented")
}

// NewIterator returns an iterator for listing clusters.
func (s *ClustersService) NewIterator(opts ...converged.ODataOption) converged.Iterator[clusterModels.Cluster] {
	return nil
}

// ListClusterVirtualGPUs returns the virtual GPU configuration for the given cluster UUID.
func (s *ClustersService) ListClusterVirtualGPUs(ctx context.Context, clusterUuid string, opts ...converged.ODataOption) ([]clusterModels.VirtualGpuProfile, error) {
	return nil, fmt.Errorf("ListClusterVirtualGPUs not implemented")
}

// ListClusterPhysicalGPUs returns the physical GPU configuration for the given cluster UUID.
func (s *ClustersService) ListClusterPhysicalGPUs(ctx context.Context, clusterUuid string, opts ...converged.ODataOption) ([]clusterModels.PhysicalGpuProfile, error) {
	return nil, fmt.Errorf("ListClusterPhysicalGPUs not implemented")
}
