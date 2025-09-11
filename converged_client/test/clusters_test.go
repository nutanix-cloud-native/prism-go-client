package test

import (
	"context"
	"errors"
	"testing"

	convergedclient "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	"github.com/nutanix-cloud-native/prism-go-client/converged_client/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"k8s.io/utils/ptr"
)

// TestCluster represents a test cluster entity
type TestCluster struct {
	ExtId *string
	Name  *string
	State *string
}

// TestVirtualGpuProfile represents a test virtual GPU profile entity
type TestVirtualGpuProfile struct {
	ExtId       *string
	Name        *string
	Type        *string
	DeviceId    *string
	VendorId    *string
	SubsystemId *string
}

// TestPhysicalGpuProfile represents a test physical GPU profile entity
type TestPhysicalGpuProfile struct {
	ExtId    *string
	Name     *string
	Type     *string
	DeviceId *string
	VendorId *string
}

// TestClustersInterface tests the Clusters interface using generated mocks
func TestClustersInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClusters := mocks.NewMockClusters[TestCluster, TestVirtualGpuProfile, TestPhysicalGpuProfile](ctrl)
	ctx := context.Background()

	// Test Get method
	t.Run("Get", func(t *testing.T) {
		expectedCluster := &TestCluster{
			ExtId: ptr.To("test-cluster-1"),
			Name:  ptr.To("Test Cluster"),
			State: ptr.To("RUNNING"),
		}

		mockClusters.EXPECT().
			Get(ctx, "test-cluster-1").
			Return(expectedCluster, nil)

		result, err := mockClusters.Get(ctx, "test-cluster-1")
		assert.NoError(t, err)
		assert.Equal(t, expectedCluster, result)
	})

	// Test Get with error
	t.Run("Get with error", func(t *testing.T) {
		expectedError := errors.New("cluster not found")

		mockClusters.EXPECT().
			Get(ctx, "non-existent-cluster").
			Return(nil, expectedError)

		result, err := mockClusters.Get(ctx, "non-existent-cluster")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test List method
	t.Run("List", func(t *testing.T) {
		expectedClusters := []TestCluster{
			{ExtId: ptr.To("1"), Name: ptr.To("Cluster 1"), State: ptr.To("RUNNING")},
			{ExtId: ptr.To("2"), Name: ptr.To("Cluster 2"), State: ptr.To("STOPPED")},
		}

		mockClusters.EXPECT().
			List(ctx, gomock.Any()).
			Return(expectedClusters, nil)

		result, err := mockClusters.List(ctx, convergedclient.WithPage(1), convergedclient.WithLimit(10))
		assert.NoError(t, err)
		assert.Equal(t, expectedClusters, result)
	})

	// Test List with error
	t.Run("List with error", func(t *testing.T) {
		expectedError := errors.New("failed to list clusters")

		mockClusters.EXPECT().
			List(ctx, gomock.Any()).
			Return(nil, expectedError)

		result, err := mockClusters.List(ctx)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test NewIterator method
	t.Run("NewIterator", func(t *testing.T) {
		// Create a mock iterator
		mockIterator := func(yield func(TestCluster, error) bool) {
			clusters := []TestCluster{
				{ExtId: ptr.To("1"), Name: ptr.To("Cluster 1")},
				{ExtId: ptr.To("2"), Name: ptr.To("Cluster 2")},
			}
			for _, cluster := range clusters {
				if !yield(cluster, nil) {
					return
				}
			}
		}

		mockClusters.EXPECT().
			NewIterator(ctx, gomock.Any()).
			Return(convergedclient.Iterator[TestCluster](mockIterator))

		iterator := mockClusters.NewIterator(ctx, convergedclient.WithPage(1))
		assert.NotNil(t, iterator)

		// Test iteration
		var results []TestCluster
		for cluster, err := range iterator {
			assert.NoError(t, err)
			results = append(results, cluster)
		}
		assert.Len(t, results, 2)
	})

	// Test ListClusterVirtualGPUs method
	t.Run("ListClusterVirtualGPUs", func(t *testing.T) {
		expectedGpuProfiles := []TestVirtualGpuProfile{
			{ExtId: ptr.To("vgpu-1"), Name: ptr.To("NVIDIA Grid K2"), Type: ptr.To("passthrough")},
			{ExtId: ptr.To("vgpu-2"), Name: ptr.To("NVIDIA Grid K1"), Type: ptr.To("shared")},
		}

		mockClusters.EXPECT().
			ListClusterVirtualGPUs(ctx, "test-cluster-1", gomock.Any()).
			Return(expectedGpuProfiles, nil)

		result, err := mockClusters.ListClusterVirtualGPUs(ctx, "test-cluster-1", convergedclient.WithLimit(10))
		assert.NoError(t, err)
		assert.Equal(t, expectedGpuProfiles, result)
	})

	// Test ListClusterVirtualGPUs with error
	t.Run("ListClusterVirtualGPUs with error", func(t *testing.T) {
		expectedError := errors.New("failed to list virtual GPU profiles")

		mockClusters.EXPECT().
			ListClusterVirtualGPUs(ctx, "non-existent-cluster", gomock.Any()).
			Return(nil, expectedError)

		result, err := mockClusters.ListClusterVirtualGPUs(ctx, "non-existent-cluster")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	// Test ListClusterPhysicalGPUs method
	t.Run("ListClusterPhysicalGPUs", func(t *testing.T) {
		expectedGpuProfiles := []TestPhysicalGpuProfile{
			{ExtId: ptr.To("pgpu-1"), Name: ptr.To("NVIDIA Tesla V100"), Type: ptr.To("compute")},
			{ExtId: ptr.To("pgpu-2"), Name: ptr.To("NVIDIA Tesla P40"), Type: ptr.To("graphics")},
		}

		mockClusters.EXPECT().
			ListClusterPhysicalGPUs(ctx, "test-cluster-1", gomock.Any()).
			Return(expectedGpuProfiles, nil)

		result, err := mockClusters.ListClusterPhysicalGPUs(ctx, "test-cluster-1", convergedclient.WithFilter("type eq 'compute'"))
		assert.NoError(t, err)
		assert.Equal(t, expectedGpuProfiles, result)
	})

	// Test ListClusterPhysicalGPUs with error
	t.Run("ListClusterPhysicalGPUs with error", func(t *testing.T) {
		expectedError := errors.New("failed to list physical GPU profiles")

		mockClusters.EXPECT().
			ListClusterPhysicalGPUs(ctx, "non-existent-cluster", gomock.Any()).
			Return(nil, expectedError)

		result, err := mockClusters.ListClusterPhysicalGPUs(ctx, "non-existent-cluster")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})
}

// TestClustersInterfaceComposition tests that Clusters interface properly composes other interfaces
func TestClustersInterfaceComposition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClusters := mocks.NewMockClusters[TestCluster, TestVirtualGpuProfile, TestPhysicalGpuProfile](ctrl)
	ctx := context.Background()

	// Test that Clusters can be used as Getter
	t.Run("as Getter", func(t *testing.T) {
		var getter convergedclient.Getter[TestCluster] = mockClusters

		expectedCluster := &TestCluster{ExtId: ptr.To("1"), Name: ptr.To("Test")}
		mockClusters.EXPECT().Get(ctx, "1").Return(expectedCluster, nil)

		result, err := getter.Get(ctx, "1")
		assert.NoError(t, err)
		assert.Equal(t, expectedCluster, result)
	})

	// Test that Clusters can be used as Lister (through embedded interface)
	t.Run("as Lister", func(t *testing.T) {
		expectedClusters := []TestCluster{{ExtId: ptr.To("1"), Name: ptr.To("Test")}}
		mockClusters.EXPECT().List(ctx).Return(expectedClusters, nil)

		result, err := mockClusters.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, expectedClusters, result)
	})

	// Test that Clusters provides iterator functionality
	t.Run("as Iterator provider", func(t *testing.T) {
		mockIterator := func(yield func(TestCluster, error) bool) {
			clusters := []TestCluster{{ExtId: ptr.To("1"), Name: ptr.To("Test")}}
			for _, cluster := range clusters {
				if !yield(cluster, nil) {
					return
				}
			}
		}

		mockClusters.EXPECT().NewIterator(ctx).Return(convergedclient.Iterator[TestCluster](mockIterator))

		iterator := mockClusters.NewIterator(ctx)
		assert.NotNil(t, iterator)

		// Test iteration
		var results []TestCluster
		for cluster, err := range iterator {
			assert.NoError(t, err)
			results = append(results, cluster)
		}
		assert.Len(t, results, 1)
	})
}

// TestClustersErrorHandling tests various error scenarios
func TestClustersErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClusters := mocks.NewMockClusters[TestCluster, TestVirtualGpuProfile, TestPhysicalGpuProfile](ctrl)
	ctx := context.Background()

	// Test context cancellation
	t.Run("context cancellation", func(t *testing.T) {
		cancelledCtx, cancel := context.WithCancel(ctx)
		cancel()

		expectedError := context.Canceled
		mockClusters.EXPECT().Get(cancelledCtx, "test").Return(nil, expectedError)

		_, err := mockClusters.Get(cancelledCtx, "test")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	// Test context timeout
	t.Run("context timeout", func(t *testing.T) {
		timeoutCtx, cancel := context.WithTimeout(ctx, 0) // Immediate timeout
		defer cancel()

		expectedError := context.DeadlineExceeded
		mockClusters.EXPECT().List(timeoutCtx).Return(nil, expectedError)

		_, err := mockClusters.List(timeoutCtx)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	// Test invalid input
	t.Run("invalid input", func(t *testing.T) {
		expectedError := errors.New("invalid UUID format")
		mockClusters.EXPECT().Get(ctx, "").Return(nil, expectedError)

		_, err := mockClusters.Get(ctx, "")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	// Test invalid cluster UUID for GPU profiles
	t.Run("invalid cluster UUID for virtual GPUs", func(t *testing.T) {
		expectedError := errors.New("invalid cluster UUID format")
		mockClusters.EXPECT().ListClusterVirtualGPUs(ctx, "", gomock.Any()).Return(nil, expectedError)

		_, err := mockClusters.ListClusterVirtualGPUs(ctx, "")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	// Test invalid cluster UUID for physical GPUs
	t.Run("invalid cluster UUID for physical GPUs", func(t *testing.T) {
		expectedError := errors.New("invalid cluster UUID format")
		mockClusters.EXPECT().ListClusterPhysicalGPUs(ctx, "", gomock.Any()).Return(nil, expectedError)

		_, err := mockClusters.ListClusterPhysicalGPUs(ctx, "")
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

// TestClustersWithODataOptions tests Clusters interface with various OData options
func TestClustersWithODataOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClusters := mocks.NewMockClusters[TestCluster, TestVirtualGpuProfile, TestPhysicalGpuProfile](ctrl)
	ctx := context.Background()

	// Test with page option
	t.Run("with page option", func(t *testing.T) {
		expectedClusters := []TestCluster{{ExtId: ptr.To("1"), Name: ptr.To("Page 1")}}
		mockClusters.EXPECT().List(ctx, gomock.Any()).Return(expectedClusters, nil)

		result, err := mockClusters.List(ctx, convergedclient.WithPage(1))
		assert.NoError(t, err)
		assert.Equal(t, expectedClusters, result)
	})

	// Test with limit option
	t.Run("with limit option", func(t *testing.T) {
		expectedClusters := []TestCluster{{ExtId: ptr.To("1"), Name: ptr.To("Limited")}}
		mockClusters.EXPECT().List(ctx, gomock.Any()).Return(expectedClusters, nil)

		result, err := mockClusters.List(ctx, convergedclient.WithLimit(5))
		assert.NoError(t, err)
		assert.Equal(t, expectedClusters, result)
	})

	// Test with filter option
	t.Run("with filter option", func(t *testing.T) {
		expectedClusters := []TestCluster{{ExtId: ptr.To("1"), Name: ptr.To("Filtered")}}
		mockClusters.EXPECT().List(ctx, gomock.Any()).Return(expectedClusters, nil)

		result, err := mockClusters.List(ctx, convergedclient.WithFilter("name eq 'test'"))
		assert.NoError(t, err)
		assert.Equal(t, expectedClusters, result)
	})

	// Test with order by option
	t.Run("with order by option", func(t *testing.T) {
		expectedClusters := []TestCluster{{ExtId: ptr.To("1"), Name: ptr.To("Ordered")}}
		mockClusters.EXPECT().List(ctx, gomock.Any()).Return(expectedClusters, nil)

		result, err := mockClusters.List(ctx, convergedclient.WithOrderBy("name asc"))
		assert.NoError(t, err)
		assert.Equal(t, expectedClusters, result)
	})

	// Test with expand option
	t.Run("with expand option", func(t *testing.T) {
		expectedClusters := []TestCluster{{ExtId: ptr.To("1"), Name: ptr.To("Expanded")}}
		mockClusters.EXPECT().List(ctx, gomock.Any()).Return(expectedClusters, nil)

		result, err := mockClusters.List(ctx, convergedclient.WithExpand("metadata"))
		assert.NoError(t, err)
		assert.Equal(t, expectedClusters, result)
	})

	// Test with select option
	t.Run("with select option", func(t *testing.T) {
		expectedClusters := []TestCluster{{ExtId: ptr.To("1"), Name: ptr.To("Selected")}}
		mockClusters.EXPECT().List(ctx, gomock.Any()).Return(expectedClusters, nil)

		result, err := mockClusters.List(ctx, convergedclient.WithSelect("id,name"))
		assert.NoError(t, err)
		assert.Equal(t, expectedClusters, result)
	})

	// Test with apply option
	t.Run("with apply option", func(t *testing.T) {
		expectedClusters := []TestCluster{{ExtId: ptr.To("1"), Name: ptr.To("Applied")}}
		mockClusters.EXPECT().List(ctx, gomock.Any()).Return(expectedClusters, nil)

		result, err := mockClusters.List(ctx, convergedclient.WithApply("groupby((name))"))
		assert.NoError(t, err)
		assert.Equal(t, expectedClusters, result)
	})

	// Test with multiple options
	t.Run("with multiple options", func(t *testing.T) {
		expectedClusters := []TestCluster{{ExtId: ptr.To("1"), Name: ptr.To("Multiple")}}
		mockClusters.EXPECT().List(ctx, gomock.Any()).Return(expectedClusters, nil)

		result, err := mockClusters.List(ctx, convergedclient.WithPage(1), convergedclient.WithLimit(10), convergedclient.WithFilter("name eq 'test'"), convergedclient.WithOrderBy("name asc"))
		assert.NoError(t, err)
		assert.Equal(t, expectedClusters, result)
	})
}

// TestClustersGpuProfilesWithODataOptions tests GPU profile methods with various OData options
func TestClustersGpuProfilesWithODataOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClusters := mocks.NewMockClusters[TestCluster, TestVirtualGpuProfile, TestPhysicalGpuProfile](ctrl)
	ctx := context.Background()
	clusterUuid := "test-cluster-1"

	// Test Virtual GPU profiles with various options
	t.Run("Virtual GPU profiles", func(t *testing.T) {
		// Test with page option
		t.Run("with page option", func(t *testing.T) {
			expectedProfiles := []TestVirtualGpuProfile{{ExtId: ptr.To("vgpu-1"), Name: ptr.To("Page 1")}}
			mockClusters.EXPECT().ListClusterVirtualGPUs(ctx, clusterUuid, gomock.Any()).Return(expectedProfiles, nil)

			result, err := mockClusters.ListClusterVirtualGPUs(ctx, clusterUuid, convergedclient.WithPage(1))
			assert.NoError(t, err)
			assert.Equal(t, expectedProfiles, result)
		})

		// Test with limit option
		t.Run("with limit option", func(t *testing.T) {
			expectedProfiles := []TestVirtualGpuProfile{{ExtId: ptr.To("vgpu-1"), Name: ptr.To("Limited")}}
			mockClusters.EXPECT().ListClusterVirtualGPUs(ctx, clusterUuid, gomock.Any()).Return(expectedProfiles, nil)

			result, err := mockClusters.ListClusterVirtualGPUs(ctx, clusterUuid, convergedclient.WithLimit(5))
			assert.NoError(t, err)
			assert.Equal(t, expectedProfiles, result)
		})

		// Test with filter option
		t.Run("with filter option", func(t *testing.T) {
			expectedProfiles := []TestVirtualGpuProfile{{ExtId: ptr.To("vgpu-1"), Name: ptr.To("Filtered")}}
			mockClusters.EXPECT().ListClusterVirtualGPUs(ctx, clusterUuid, gomock.Any()).Return(expectedProfiles, nil)

			result, err := mockClusters.ListClusterVirtualGPUs(ctx, clusterUuid, convergedclient.WithFilter("type eq 'shared'"))
			assert.NoError(t, err)
			assert.Equal(t, expectedProfiles, result)
		})

		// Test with multiple options
		t.Run("with multiple options", func(t *testing.T) {
			expectedProfiles := []TestVirtualGpuProfile{{ExtId: ptr.To("vgpu-1"), Name: ptr.To("Multiple")}}
			mockClusters.EXPECT().ListClusterVirtualGPUs(ctx, clusterUuid, gomock.Any()).Return(expectedProfiles, nil)

			result, err := mockClusters.ListClusterVirtualGPUs(ctx, clusterUuid, convergedclient.WithPage(1), convergedclient.WithLimit(10), convergedclient.WithOrderBy("name asc"))
			assert.NoError(t, err)
			assert.Equal(t, expectedProfiles, result)
		})
	})

	// Test Physical GPU profiles with various options
	t.Run("Physical GPU profiles", func(t *testing.T) {
		// Test with page option
		t.Run("with page option", func(t *testing.T) {
			expectedProfiles := []TestPhysicalGpuProfile{{ExtId: ptr.To("pgpu-1"), Name: ptr.To("Page 1")}}
			mockClusters.EXPECT().ListClusterPhysicalGPUs(ctx, clusterUuid, gomock.Any()).Return(expectedProfiles, nil)

			result, err := mockClusters.ListClusterPhysicalGPUs(ctx, clusterUuid, convergedclient.WithPage(1))
			assert.NoError(t, err)
			assert.Equal(t, expectedProfiles, result)
		})

		// Test with limit option
		t.Run("with limit option", func(t *testing.T) {
			expectedProfiles := []TestPhysicalGpuProfile{{ExtId: ptr.To("pgpu-1"), Name: ptr.To("Limited")}}
			mockClusters.EXPECT().ListClusterPhysicalGPUs(ctx, clusterUuid, gomock.Any()).Return(expectedProfiles, nil)

			result, err := mockClusters.ListClusterPhysicalGPUs(ctx, clusterUuid, convergedclient.WithLimit(5))
			assert.NoError(t, err)
			assert.Equal(t, expectedProfiles, result)
		})

		// Test with filter option
		t.Run("with filter option", func(t *testing.T) {
			expectedProfiles := []TestPhysicalGpuProfile{{ExtId: ptr.To("pgpu-1"), Name: ptr.To("Filtered")}}
			mockClusters.EXPECT().ListClusterPhysicalGPUs(ctx, clusterUuid, gomock.Any()).Return(expectedProfiles, nil)

			result, err := mockClusters.ListClusterPhysicalGPUs(ctx, clusterUuid, convergedclient.WithFilter("type eq 'compute'"))
			assert.NoError(t, err)
			assert.Equal(t, expectedProfiles, result)
		})

		// Test with multiple options
		t.Run("with multiple options", func(t *testing.T) {
			expectedProfiles := []TestPhysicalGpuProfile{{ExtId: ptr.To("pgpu-1"), Name: ptr.To("Multiple")}}
			mockClusters.EXPECT().ListClusterPhysicalGPUs(ctx, clusterUuid, gomock.Any()).Return(expectedProfiles, nil)

			result, err := mockClusters.ListClusterPhysicalGPUs(ctx, clusterUuid, convergedclient.WithPage(1), convergedclient.WithLimit(10), convergedclient.WithOrderBy("name asc"))
			assert.NoError(t, err)
			assert.Equal(t, expectedProfiles, result)
		})
	})
}
