package v4

import (
	"context"
	"testing"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	clusterModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	"github.com/stretchr/testify/assert"
)

func TestNewClustersService(t *testing.T) {
	service := NewClustersService()
	assert.NotNil(t, service)
	assert.IsType(t, &ClustersService{}, service)
	assert.Nil(t, service.client)
}

func TestNewClustersServiceWithClient(t *testing.T) {
	mockClient := &Client{}
	service := NewClustersServiceWithClient(mockClient)
	assert.NotNil(t, service)
	assert.IsType(t, &ClustersService{}, service)
	assert.Equal(t, mockClient, service.client)
}

func TestClustersService_Get_NilClient(t *testing.T) {
	service := NewClustersService()
	ctx := context.Background()

	result, err := service.Get(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "client not initialized")
}

func TestClustersService_List_NilClient(t *testing.T) {
	service := NewClustersService()
	ctx := context.Background()

	tests := []struct {
		name string
		opts []converged.ODataOption
	}{
		{
			name: "no options",
			opts: []converged.ODataOption{},
		},
		{
			name: "with page option",
			opts: []converged.ODataOption{converged.WithPage(1)},
		},
		{
			name: "with limit option",
			opts: []converged.ODataOption{converged.WithLimit(10)},
		},
		{
			name: "with filter option",
			opts: []converged.ODataOption{converged.WithFilter("name eq 'test'")},
		},
		{
			name: "with orderby option",
			opts: []converged.ODataOption{converged.WithOrderBy("name asc")},
		},
		{
			name: "with expand option",
			opts: []converged.ODataOption{converged.WithExpand("nodes")},
		},
		{
			name: "with select option",
			opts: []converged.ODataOption{converged.WithSelect("name,extId")},
		},
		{
			name: "with apply option",
			opts: []converged.ODataOption{converged.WithApply("groupby((name))")},
		},
		{
			name: "with multiple options",
			opts: []converged.ODataOption{
				converged.WithPage(2),
				converged.WithLimit(5),
				converged.WithFilter("name eq 'cluster1'"),
				converged.WithOrderBy("name desc"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.List(ctx, tt.opts...)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "client not initialized")
		})
	}
}

func TestClustersService_ListAll_NilClient(t *testing.T) {
	service := NewClustersService()
	ctx := context.Background()

	tests := []struct {
		name string
		opts []converged.ODataOption
	}{
		{
			name: "no options",
			opts: []converged.ODataOption{},
		},
		{
			name: "with filter option",
			opts: []converged.ODataOption{converged.WithFilter("name eq 'test'")},
		},
		{
			name: "with orderby option",
			opts: []converged.ODataOption{converged.WithOrderBy("name asc")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ListAll(ctx, tt.opts...)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "client not initialized")
		})
	}
}

func TestClustersService_NewIterator_NilClient(t *testing.T) {
	service := NewClustersService()

	tests := []struct {
		name string
		opts []converged.ODataOption
	}{
		{
			name: "no options",
			opts: []converged.ODataOption{},
		},
		{
			name: "with page option",
			opts: []converged.ODataOption{converged.WithPage(1)},
		},
		{
			name: "with limit option",
			opts: []converged.ODataOption{converged.WithLimit(10)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator := service.NewIterator(tt.opts...)
			assert.NotNil(t, iterator)

			// Test that the iterator doesn't yield any results (should be a no-op function)
			count := 0
			for cluster, err := range iterator {
				count++
				// Should not execute since it's a no-op function for nil client
				assert.Fail(t, "Iterator should not yield results for nil client", "cluster: %v, err: %v", cluster, err)
			}
			assert.Equal(t, 0, count)
		})
	}
}

func TestClustersService_ListClusterVirtualGPUs_NilClient(t *testing.T) {
	service := NewClustersService()
	ctx := context.Background()

	tests := []struct {
		name        string
		clusterUuid string
		opts        []converged.ODataOption
	}{
		{
			name:        "basic test",
			clusterUuid: "cluster-uuid-123",
			opts:        []converged.ODataOption{},
		},
		{
			name:        "with filter",
			clusterUuid: "cluster-uuid-456",
			opts:        []converged.ODataOption{converged.WithFilter("name eq 'gpu1'")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ListClusterVirtualGPUs(ctx, tt.clusterUuid, tt.opts...)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "client not initialized")
		})
	}
}

func TestClustersService_ListClusterPhysicalGPUs_NilClient(t *testing.T) {
	service := NewClustersService()
	ctx := context.Background()

	tests := []struct {
		name        string
		clusterUuid string
		opts        []converged.ODataOption
	}{
		{
			name:        "basic test",
			clusterUuid: "cluster-uuid-123",
			opts:        []converged.ODataOption{},
		},
		{
			name:        "with filter",
			clusterUuid: "cluster-uuid-456",
			opts:        []converged.ODataOption{converged.WithFilter("name eq 'gpu1'")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ListClusterPhysicalGPUs(ctx, tt.clusterUuid, tt.opts...)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "client not initialized")
		})
	}
}

// Test struct validation
func TestClustersService_Interface_Compliance(t *testing.T) {
	var service converged.Clusters[clusterModels.Cluster, clusterModels.VirtualGpuProfile, clusterModels.PhysicalGpuProfile]
	service = NewClustersService()
	assert.NotNil(t, service)
}

// Benchmarks for performance testing
func BenchmarkNewClustersService(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewClustersService()
	}
}

func BenchmarkNewClustersServiceWithClient(b *testing.B) {
	mockClient := &Client{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewClustersServiceWithClient(mockClient)
	}
}
