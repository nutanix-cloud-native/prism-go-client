package v4

import (
	"context"
	"testing"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	clusterModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	"github.com/stretchr/testify/assert"
)

func TestNewClustersService(t *testing.T) {
	service := NewClustersService(nil)
	assert.NotNil(t, service)
	assert.IsType(t, &ClustersService{}, service)
}

func TestClustersService_Get(t *testing.T) {
	service := NewClustersService(nil)
	ctx := context.Background()

	result, err := service.Get(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")
}

func TestClustersService_List(t *testing.T) {
	service := NewClustersService(nil)
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
			name: "with order by option",
			opts: []converged.ODataOption{converged.WithOrderBy("name")},
		},
		{
			name: "with expand option",
			opts: []converged.ODataOption{converged.WithExpand("resources")},
		},
		{
			name: "with select option",
			opts: []converged.ODataOption{converged.WithSelect("name,uuid")},
		},
		{
			name: "with apply option",
			opts: []converged.ODataOption{converged.WithApply("groupby((name))")},
		},
		{
			name: "with multiple options",
			opts: []converged.ODataOption{
				converged.WithPage(1),
				converged.WithLimit(10),
				converged.WithFilter("name eq 'test'"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.List(ctx, tt.opts...)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "not implemented")
		})
	}
}

func TestClustersService_NewIterator(t *testing.T) {
	service := NewClustersService(nil)

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
			name: "with order by option",
			opts: []converged.ODataOption{converged.WithOrderBy("name")},
		},
		{
			name: "with multiple options",
			opts: []converged.ODataOption{
				converged.WithFilter("name eq 'test'"),
				converged.WithOrderBy("name"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			iterator := service.NewIterator(ctx, tt.opts...)
			assert.Nil(t, iterator)
		})
	}
}

func TestClustersService_ListClusterVirtualGPUs(t *testing.T) {
	service := NewClustersService(nil)
	ctx := context.Background()
	clusterUuid := "test-cluster-uuid"

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
			name: "with order by option",
			opts: []converged.ODataOption{converged.WithOrderBy("name")},
		},
		{
			name: "with expand option",
			opts: []converged.ODataOption{converged.WithExpand("resources")},
		},
		{
			name: "with select option",
			opts: []converged.ODataOption{converged.WithSelect("name,uuid")},
		},
		{
			name: "with apply option",
			opts: []converged.ODataOption{converged.WithApply("groupby((name))")},
		},
		{
			name: "with multiple options",
			opts: []converged.ODataOption{
				converged.WithPage(1),
				converged.WithLimit(10),
				converged.WithFilter("name eq 'test'"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ListClusterVirtualGPUs(ctx, clusterUuid, tt.opts...)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "ListClusterVirtualGPUs not implemented")
		})
	}
}

func TestClustersService_ListClusterPhysicalGPUs(t *testing.T) {
	service := NewClustersService(nil)
	ctx := context.Background()
	clusterUuid := "test-cluster-uuid"

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
			name: "with order by option",
			opts: []converged.ODataOption{converged.WithOrderBy("name")},
		},
		{
			name: "with expand option",
			opts: []converged.ODataOption{converged.WithExpand("resources")},
		},
		{
			name: "with select option",
			opts: []converged.ODataOption{converged.WithSelect("name,uuid")},
		},
		{
			name: "with apply option",
			opts: []converged.ODataOption{converged.WithApply("groupby((name))")},
		},
		{
			name: "with multiple options",
			opts: []converged.ODataOption{
				converged.WithPage(1),
				converged.WithLimit(10),
				converged.WithFilter("name eq 'test'"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ListClusterPhysicalGPUs(ctx, clusterUuid, tt.opts...)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "ListClusterPhysicalGPUs not implemented")
		})
	}
}

// TestClustersInterface tests the Clusters interface implementation
func TestClustersInterface(t *testing.T) {
	service := NewClustersService(nil)

	// Test that the service implements the Clusters interface
	var _ converged.Clusters[clusterModels.Cluster, clusterModels.VirtualGpuProfile, clusterModels.PhysicalGpuProfile] = service

	ctx := context.Background()

	// Test Get method
	_, err := service.Get(ctx, "test-uuid")
	assert.Error(t, err)

	// Test List method
	_, err = service.List(ctx)
	assert.Error(t, err)

	// Test NewIterator method
	iterator := service.NewIterator(ctx)
	assert.Nil(t, iterator)

	// Test ListClusterVirtualGPUs method
	_, err = service.ListClusterVirtualGPUs(ctx, "test-cluster-uuid")
	assert.Error(t, err)

	// Test ListClusterPhysicalGPUs method
	_, err = service.ListClusterPhysicalGPUs(ctx, "test-cluster-uuid")
	assert.Error(t, err)
}

// TestClustersService_EdgeCases tests edge cases and error conditions
func TestClustersService_EdgeCases(t *testing.T) {
	service := NewClustersService(nil)
	ctx := context.Background()

	t.Run("empty uuid", func(t *testing.T) {
		result, err := service.Get(ctx, "")
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("empty cluster uuid for virtual GPUs", func(t *testing.T) {
		result, err := service.ListClusterVirtualGPUs(ctx, "")
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("empty cluster uuid for physical GPUs", func(t *testing.T) {
		result, err := service.ListClusterPhysicalGPUs(ctx, "")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
