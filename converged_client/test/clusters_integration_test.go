package test

import (
	"context"
	"strings"
	"testing"

	models "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	v4ConvergedClient "github.com/nutanix-cloud-native/prism-go-client/converged_client/v4"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var clustersClient *v4ConvergedClient.Client

func initializeClustersClient(t *testing.T) {
	if clustersClient != nil {
		return
	}

	var err error
	creds := testhelpers.CredentialsFromEnvironment(t)

	clustersClient, err = v4ConvergedClient.NewClient(creds)
	assert.Nil(t, err)
}

func TestClustersService_Get(t *testing.T) {
	initializeClustersClient(t)
	ctx := context.Background()

	// First, get a list of clusters to find one to retrieve
	clusters, err := clustersClient.Clusters.List(ctx, models.WithPage(0), models.WithLimit(1))
	require.Nil(t, err)

	if len(clusters) > 0 {
		// Test Get operation with the first cluster
		cluster, err := clustersClient.Clusters.Get(ctx, *clusters[0].ExtId)
		require.Nil(t, err)
		require.NotNil(t, cluster)
		assert.Equal(t, *clusters[0].ExtId, *cluster.ExtId)
		t.Logf("Retrieved cluster: %s (ID: %s)", *cluster.Name, *cluster.ExtId)
	} else {
		t.Skip("No clusters available for testing")
	}
}

func TestClustersService_List(t *testing.T) {
	initializeClustersClient(t)
	ctx := context.Background()

	// Test supported operations - these should succeed
	supportedTestCases := []struct {
		name string
		opts []models.ODataOption
	}{
		{
			name: "no options",
			opts: []models.ODataOption{},
		},
		{
			name: "with page option",
			opts: []models.ODataOption{models.WithPage(1)},
		},
		{
			name: "with limit option",
			opts: []models.ODataOption{models.WithLimit(10)},
		},
	}

	for _, tt := range supportedTestCases {
		t.Run(tt.name, func(t *testing.T) {
			clusters, err := clustersClient.Clusters.List(ctx, tt.opts...)
			assert.Nil(t, err)
			t.Logf("Listed %d clusters with options: %s", len(clusters), tt.name)
		})
	}

	// Test potentially unsupported operations - these may fail depending on API support
	potentiallyUnsupportedTestCases := []struct {
		name string
		opts []models.ODataOption
	}{
		{
			name: "with filter option",
			opts: []models.ODataOption{models.WithFilter("name eq 'test'")},
		},
		{
			name: "with order by option",
			opts: []models.ODataOption{models.WithOrderBy("name")},
		},
		{
			name: "with expand option",
			opts: []models.ODataOption{models.WithExpand("nodes")},
		},
		{
			name: "with select option",
			opts: []models.ODataOption{models.WithSelect("name,extId")},
		},
		{
			name: "with apply option",
			opts: []models.ODataOption{models.WithApply("groupby((name))")},
		},
	}

	for _, tt := range potentiallyUnsupportedTestCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := clustersClient.Clusters.List(ctx, tt.opts...)
			// Note: We don't assert on error here since support varies by API version
			if err != nil {
				t.Logf("Option %s is not supported: %v", tt.name, err)
			} else {
				t.Logf("Option %s is supported", tt.name)
			}
		})
	}
}

func TestClustersService_NewIterator(t *testing.T) {
	initializeClustersClient(t)
	ctx := context.Background()

	// Test supported operations - these should succeed
	supportedTestCases := []struct {
		name string
		opts []models.ODataOption
	}{
		{
			name: "no options",
			opts: []models.ODataOption{},
		},
		{
			name: "with limit option",
			opts: []models.ODataOption{models.WithLimit(5)},
		},
	}

	for _, tt := range supportedTestCases {
		t.Run(tt.name, func(t *testing.T) {
			iterator := clustersClient.Clusters.NewIterator(ctx, tt.opts...)
			require.NotNil(t, iterator)

			count := 0
			for cluster, err := range iterator {
				assert.Nil(t, err)
				assert.NotNil(t, cluster.ExtId)
				count++

				// Limit iteration to prevent long-running tests
				if count >= 3 {
					break
				}
			}
			t.Logf("Iterated over %d clusters with options: %s", count, tt.name)
		})
	}
}

func TestClustersService_ListClusterVirtualGPUs(t *testing.T) {
	initializeClustersClient(t)
	ctx := context.Background()

	// First, get a cluster to test GPU listing
	clusters, err := clustersClient.Clusters.List(ctx, models.WithLimit(1))
	require.Nil(t, err)

	if len(clusters) > 0 {
		clusterUuid := *clusters[0].ExtId

		// Test supported operations
		supportedTestCases := []struct {
			name string
			opts []models.ODataOption
		}{
			{
				name: "no options",
				opts: []models.ODataOption{},
			},
			{
				name: "with page option",
				opts: []models.ODataOption{models.WithPage(0)},
			},
			{
				name: "with limit option",
				opts: []models.ODataOption{models.WithLimit(10)},
			},
		}

		for _, tt := range supportedTestCases {
			t.Run(tt.name, func(t *testing.T) {
				virtualGPUs, err := clustersClient.Clusters.ListClusterVirtualGPUs(ctx, clusterUuid, tt.opts...)
				assert.Nil(t, err)
				t.Logf("Listed %d virtual GPUs for cluster %s with options: %s",
					len(virtualGPUs), clusterUuid, tt.name)
			})
		}

		// Test potentially unsupported operations
		potentiallyUnsupportedTestCases := []struct {
			name string
			opts []models.ODataOption
		}{
			{
				name: "with filter option",
				opts: []models.ODataOption{models.WithFilter("name eq 'test'")},
			},
			{
				name: "with order by option",
				opts: []models.ODataOption{models.WithOrderBy("name")},
			},
		}

		for _, tt := range potentiallyUnsupportedTestCases {
			t.Run(tt.name, func(t *testing.T) {
				_, err := clustersClient.Clusters.ListClusterVirtualGPUs(ctx, clusterUuid, tt.opts...)
				if err != nil {
					t.Logf("Virtual GPU listing with %s is not supported: %v", tt.name, err)
				} else {
					t.Logf("Virtual GPU listing with %s is supported", tt.name)
				}
			})
		}
	} else {
		t.Skip("No clusters available for virtual GPU testing")
	}
}

func TestClustersService_ListClusterPhysicalGPUs(t *testing.T) {
	initializeClustersClient(t)
	ctx := context.Background()

	// First, get a cluster to test GPU listing
	clusters, err := clustersClient.Clusters.List(ctx, models.WithLimit(1))
	require.Nil(t, err)

	if len(clusters) > 0 {
		clusterUuid := *clusters[0].ExtId

		// Test supported operations
		supportedTestCases := []struct {
			name string
			opts []models.ODataOption
		}{
			{
				name: "no options",
				opts: []models.ODataOption{},
			},
			{
				name: "with page option",
				opts: []models.ODataOption{models.WithPage(0)},
			},
			{
				name: "with limit option",
				opts: []models.ODataOption{models.WithLimit(10)},
			},
		}

		for _, tt := range supportedTestCases {
			t.Run(tt.name, func(t *testing.T) {
				physicalGPUs, err := clustersClient.Clusters.ListClusterPhysicalGPUs(ctx, clusterUuid, tt.opts...)
				assert.Nil(t, err)
				t.Logf("Listed %d physical GPUs for cluster %s with options: %s",
					len(physicalGPUs), clusterUuid, tt.name)
			})
		}

		// Test potentially unsupported operations
		potentiallyUnsupportedTestCases := []struct {
			name string
			opts []models.ODataOption
		}{
			{
				name: "with filter option",
				opts: []models.ODataOption{models.WithFilter("name eq 'test'")},
			},
			{
				name: "with order by option",
				opts: []models.ODataOption{models.WithOrderBy("name")},
			},
		}

		for _, tt := range potentiallyUnsupportedTestCases {
			t.Run(tt.name, func(t *testing.T) {
				_, err := clustersClient.Clusters.ListClusterPhysicalGPUs(ctx, clusterUuid, tt.opts...)
				if err != nil {
					t.Logf("Physical GPU listing with %s is not supported: %v", tt.name, err)
				} else {
					t.Logf("Physical GPU listing with %s is supported", tt.name)
				}
			})
		}
	} else {
		t.Skip("No clusters available for physical GPU testing")
	}
}

func TestClustersService_EdgeCases(t *testing.T) {
	initializeClustersClient(t)
	ctx := context.Background()

	t.Run("Get with invalid UUID", func(t *testing.T) {
		_, err := clustersClient.Clusters.Get(ctx, "invalid-uuid")
		assert.NotNil(t, err)
		assert.True(t, strings.Contains(err.Error(), "404") ||
			strings.Contains(err.Error(), "not found") ||
			strings.Contains(err.Error(), "invalid"))
	})

	t.Run("Get with empty UUID", func(t *testing.T) {
		_, err := clustersClient.Clusters.Get(ctx, "")
		assert.NotNil(t, err)
	})

	t.Run("ListClusterVirtualGPUs with invalid cluster UUID", func(t *testing.T) {
		_, err := clustersClient.Clusters.ListClusterVirtualGPUs(ctx, "invalid-uuid")
		assert.NotNil(t, err)
	})

	t.Run("ListClusterVirtualGPUs with empty cluster UUID", func(t *testing.T) {
		_, err := clustersClient.Clusters.ListClusterVirtualGPUs(ctx, "")
		assert.NotNil(t, err)
	})

	t.Run("ListClusterPhysicalGPUs with invalid cluster UUID", func(t *testing.T) {
		_, err := clustersClient.Clusters.ListClusterPhysicalGPUs(ctx, "invalid-uuid")
		assert.NotNil(t, err)
	})

	t.Run("ListClusterPhysicalGPUs with empty cluster UUID", func(t *testing.T) {
		_, err := clustersClient.Clusters.ListClusterPhysicalGPUs(ctx, "")
		assert.NotNil(t, err)
	})
}

func TestClustersService_Comprehensive(t *testing.T) {
	initializeClustersClient(t)
	ctx := context.Background()

	// Test comprehensive cluster information retrieval
	t.Run("comprehensive cluster info", func(t *testing.T) {
		// List clusters
		clusters, err := clustersClient.Clusters.List(ctx, models.WithLimit(3))
		require.Nil(t, err)
		t.Logf("Found %d clusters in the environment", len(clusters))

		if len(clusters) == 0 {
			t.Skip("No clusters available for comprehensive testing")
			return
		}

		// Test each cluster
		for i, cluster := range clusters {
			if i >= 2 { // Limit to 2 clusters to keep test time reasonable
				break
			}

			clusterUuid := *cluster.ExtId
			clusterName := "unknown"
			if cluster.Name != nil {
				clusterName = *cluster.Name
			}

			t.Logf("Testing cluster %d: %s (%s)", i+1, clusterName, clusterUuid)

			// Get detailed cluster info
			detailedCluster, err := clustersClient.Clusters.Get(ctx, clusterUuid)
			require.Nil(t, err)
			assert.Equal(t, clusterUuid, *detailedCluster.ExtId)

			// List virtual GPUs for this cluster (may not be supported on all cluster types)
			virtualGPUs, err := clustersClient.Clusters.ListClusterVirtualGPUs(ctx, clusterUuid)
			if err != nil {
				if strings.Contains(err.Error(), "not supported") || strings.Contains(err.Error(), "CLU-10008") {
					t.Logf("  - Virtual GPUs: Not supported on this cluster type")
				} else {
					t.Logf("  - Virtual GPUs: Error - %v", err)
				}
			} else {
				t.Logf("  - Virtual GPUs: %d", len(virtualGPUs))
			}

			// List physical GPUs for this cluster (may not be supported on all cluster types)
			physicalGPUs, err := clustersClient.Clusters.ListClusterPhysicalGPUs(ctx, clusterUuid)
			if err != nil {
				if strings.Contains(err.Error(), "not supported") || strings.Contains(err.Error(), "CLU-10008") {
					t.Logf("  - Physical GPUs: Not supported on this cluster type")
				} else {
					t.Logf("  - Physical GPUs: Error - %v", err)
				}
			} else {
				t.Logf("  - Physical GPUs: %d", len(physicalGPUs))
			}
		}
	})
}
