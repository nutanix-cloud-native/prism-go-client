// Package v4 provides tests for the converged client resource groups service.
//
// Integration tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run TestResourceGroupsIntegration
//	go test -v ./converged/v4 -run TestResourceGroupsService_ErrorHandling
//	go test -v ./converged/v4 -run TestResourceGroupsService_UnsupportedODataOptions
//	go test -v ./converged/v4 -run TestExtractPrismElements
//	go test -v ./converged/v4 -run TestExtractStorageContainers
package v4

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	multidomainModels "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/config"
	commonConfig "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/common/v1/config"
)

// TestResourceGroupsIntegration tests the client.ResourceGroups service with real Nutanix API calls.
func TestResourceGroupsIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ListResourceGroups", func(t *testing.T) {
		resourceGroups, err := client.ResourceGroups.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, resourceGroups)
		assert.GreaterOrEqual(t, len(resourceGroups), 0)

		if len(resourceGroups) > 0 {
			assert.NotNil(t, resourceGroups[0].ExtId)
		}
	})

	t.Run("ListResourceGroupsWithValidOptions", func(t *testing.T) {
		resourceGroups, err := client.ResourceGroups.List(ctx,
			converged.WithPage(0),
			converged.WithLimit(5),
			converged.WithSelect("extId,name"),
			converged.WithOrderBy("name asc"),
		)
		assert.NoError(t, err)
		assert.NotNil(t, resourceGroups)
	})

	t.Run("GetResourceGroup", func(t *testing.T) {
		resourceGroups, err := client.ResourceGroups.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(resourceGroups) == 0 {
			t.Skip("No resource groups available for testing")
		}

		extID := *resourceGroups[0].ExtId
		require.NotEmpty(t, extID)

		resourceGroup, err := client.ResourceGroups.Get(ctx, extID)
		assert.NoError(t, err)
		assert.NotNil(t, resourceGroup)
		assert.Equal(t, extID, *resourceGroup.ExtId)
	})

	t.Run("NewIterator", func(t *testing.T) {
		iterator := client.ResourceGroups.NewIterator(ctx, converged.WithLimit(5))
		assert.NotNil(t, iterator)

		var count int
		for _, err := range iterator {
			if err != nil {
				t.Logf("Iterator error: %v", err)
				break
			}
			count++
		}
		assert.GreaterOrEqual(t, count, 0)
	})

	t.Run("ListPrismElements", func(t *testing.T) {
		resourceGroups, err := client.ResourceGroups.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(resourceGroups) == 0 {
			t.Skip("No resource groups available for testing")
		}

		extID := *resourceGroups[0].ExtId
		pes, err := client.ResourceGroups.ListPrismElements(ctx, extID)
		assert.NoError(t, err)
		assert.NotNil(t, pes)

		for _, pe := range pes {
			assert.NotEmpty(t, pe.ExtId)
			t.Logf("PE: %s (%s)", pe.Name, pe.ExtId)
		}
	})

	t.Run("ListStorageContainers", func(t *testing.T) {
		resourceGroups, err := client.ResourceGroups.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(resourceGroups) == 0 {
			t.Skip("No resource groups available for testing")
		}

		extID := *resourceGroups[0].ExtId
		scs, err := client.ResourceGroups.ListStorageContainers(ctx, extID)
		assert.NoError(t, err)
		assert.NotNil(t, scs)

		for _, sc := range scs {
			assert.NotEmpty(t, sc.ExtId)
			t.Logf("Storage Container: %s (%s) on PE: %s (%s)", sc.Name, sc.ExtId, sc.PrismElement.Name, sc.PrismElement.ExtId)
		}
	})
}

// TestResourceGroupsService_ErrorHandling tests error handling for nil client.
func TestResourceGroupsService_ErrorHandling(t *testing.T) {
	service := NewResourceGroupsService(nil)
	require.NotNil(t, service)

	ctx := context.Background()

	t.Run("Get_NilClient", func(t *testing.T) {
		_, err := service.Get(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("List_NilClient", func(t *testing.T) {
		_, err := service.List(ctx)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("NewIterator_NilClient", func(t *testing.T) {
		iterator := service.NewIterator(ctx)
		assert.Nil(t, iterator)
	})

	t.Run("ListPrismElements_NilClient", func(t *testing.T) {
		_, err := service.ListPrismElements(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get resource group")
	})

	t.Run("ListStorageContainers_NilClient", func(t *testing.T) {
		_, err := service.ListStorageContainers(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get resource group")
	})
}

// TestResourceGroupsService_UnsupportedODataOptions tests that apply and expand options are rejected.
func TestResourceGroupsService_UnsupportedODataOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("List_WithExpand", func(t *testing.T) {
		_, err := client.ResourceGroups.List(ctx, converged.WithExpand("config"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing ResourceGroups")
	})

	t.Run("List_WithApply", func(t *testing.T) {
		_, err := client.ResourceGroups.List(ctx, converged.WithApply("groupby((extId))"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "apply and expand options are not supported for listing ResourceGroups")
	})

	t.Run("NewIterator_WithExpand", func(t *testing.T) {
		iter := client.ResourceGroups.NewIterator(ctx, converged.WithExpand("config"))
		require.NotNil(t, iter)

		var gotError error
		for _, err := range iter {
			gotError = err
			break
		}
		require.Error(t, gotError)
		assert.Contains(t, gotError.Error(), "apply and expand options are not supported for listing ResourceGroups")
	})

	t.Run("NewIterator_WithApply", func(t *testing.T) {
		iter := client.ResourceGroups.NewIterator(ctx, converged.WithApply("groupby((extId))"))
		require.NotNil(t, iter)

		var gotError error
		for _, err := range iter {
			gotError = err
			break
		}
		require.Error(t, gotError)
		assert.Contains(t, gotError.Error(), "apply and expand options are not supported for listing ResourceGroups")
	})
}

// Helper function to create a KVPair with string value
func createKVPair(name, value string) commonConfig.KVPair {
	kv := commonConfig.KVPair{
		Name:  &name,
		Value: commonConfig.NewOneOfKVPairValue(),
	}
	_ = kv.Value.SetValue(value)
	return kv
}

// TestExtractPrismElements tests the extractPrismElements function.
func TestExtractPrismElements(t *testing.T) {
	t.Run("NilResourceGroup", func(t *testing.T) {
		result := extractPrismElements(nil)
		assert.Nil(t, result)
	})

	t.Run("NilPlacementTargets", func(t *testing.T) {
		rg := &multidomainModels.ResourceGroup{}
		result := extractPrismElements(rg)
		assert.Nil(t, result)
	})

	t.Run("EmptyPlacementTargets", func(t *testing.T) {
		rg := &multidomainModels.ResourceGroup{
			PlacementTargets: []multidomainModels.TargetDetails{},
		}
		result := extractPrismElements(rg)
		assert.Empty(t, result)
	})

	t.Run("SinglePEWithName", func(t *testing.T) {
		clusterExtId := "cluster-123"
		rg := &multidomainModels.ResourceGroup{
			PlacementTargets: []multidomainModels.TargetDetails{
				{
					ClusterExtId: &clusterExtId,
					Capabilities: []commonConfig.KVPair{
						createKVPair("cluster_name", "my-cluster"),
						createKVPair("cluster_type", "kHyperConverged"),
					},
				},
			},
		}

		result := extractPrismElements(rg)
		require.Len(t, result, 1)
		assert.Equal(t, "cluster-123", result[0].ExtId)
		assert.Equal(t, "my-cluster", result[0].Name)
	})

	t.Run("MultiplePEs", func(t *testing.T) {
		clusterExtId1 := "cluster-123"
		clusterExtId2 := "cluster-456"
		rg := &multidomainModels.ResourceGroup{
			PlacementTargets: []multidomainModels.TargetDetails{
				{
					ClusterExtId: &clusterExtId1,
					Capabilities: []commonConfig.KVPair{
						createKVPair("cluster_name", "cluster-one"),
					},
				},
				{
					ClusterExtId: &clusterExtId2,
					Capabilities: []commonConfig.KVPair{
						createKVPair("cluster_name", "cluster-two"),
					},
				},
			},
		}

		result := extractPrismElements(rg)
		require.Len(t, result, 2)
		assert.Equal(t, "cluster-123", result[0].ExtId)
		assert.Equal(t, "cluster-one", result[0].Name)
		assert.Equal(t, "cluster-456", result[1].ExtId)
		assert.Equal(t, "cluster-two", result[1].Name)
	})

	t.Run("PEWithNilClusterExtId", func(t *testing.T) {
		clusterExtId := "cluster-123"
		rg := &multidomainModels.ResourceGroup{
			PlacementTargets: []multidomainModels.TargetDetails{
				{
					ClusterExtId: nil,
				},
				{
					ClusterExtId: &clusterExtId,
					Capabilities: []commonConfig.KVPair{
						createKVPair("cluster_name", "valid-cluster"),
					},
				},
			},
		}

		result := extractPrismElements(rg)
		require.Len(t, result, 1)
		assert.Equal(t, "cluster-123", result[0].ExtId)
		assert.Equal(t, "valid-cluster", result[0].Name)
	})

	t.Run("PEWithoutClusterNameCapability", func(t *testing.T) {
		clusterExtId := "cluster-123"
		rg := &multidomainModels.ResourceGroup{
			PlacementTargets: []multidomainModels.TargetDetails{
				{
					ClusterExtId: &clusterExtId,
					Capabilities: []commonConfig.KVPair{
						createKVPair("other_capability", "some-value"),
					},
				},
			},
		}

		result := extractPrismElements(rg)
		require.Len(t, result, 1)
		assert.Equal(t, "cluster-123", result[0].ExtId)
		assert.Empty(t, result[0].Name)
	})
}

// TestExtractStorageContainers tests the extractStorageContainers function.
func TestExtractStorageContainers(t *testing.T) {
	t.Run("NilResourceGroup", func(t *testing.T) {
		result := extractStorageContainers(nil)
		assert.Nil(t, result)
	})

	t.Run("NilPlacementTargets", func(t *testing.T) {
		rg := &multidomainModels.ResourceGroup{}
		result := extractStorageContainers(rg)
		assert.Nil(t, result)
	})

	t.Run("EmptyPlacementTargets", func(t *testing.T) {
		rg := &multidomainModels.ResourceGroup{
			PlacementTargets: []multidomainModels.TargetDetails{},
		}
		result := extractStorageContainers(rg)
		assert.Empty(t, result)
	})

	t.Run("SingleStorageContainerWithPE", func(t *testing.T) {
		clusterExtId := "cluster-123"
		scExtId := "sc-456"
		rg := &multidomainModels.ResourceGroup{
			PlacementTargets: []multidomainModels.TargetDetails{
				{
					ClusterExtId: &clusterExtId,
					Capabilities: []commonConfig.KVPair{
						createKVPair("cluster_name", "my-cluster"),
					},
					StorageContainers: []multidomainModels.StorageContainerDetails{
						{
							ExtId: &scExtId,
							Capabilities: []commonConfig.KVPair{
								createKVPair("vstore_name_list", "my-container"),
							},
						},
					},
				},
			},
		}

		result := extractStorageContainers(rg)
		require.Len(t, result, 1)
		assert.Equal(t, "sc-456", result[0].ExtId)
		assert.Equal(t, "my-container", result[0].Name)
		assert.Equal(t, "cluster-123", result[0].PrismElement.ExtId)
		assert.Equal(t, "my-cluster", result[0].PrismElement.Name)
	})

	t.Run("MultipleStorageContainersOnSamePE", func(t *testing.T) {
		clusterExtId := "cluster-123"
		scExtId1 := "sc-456"
		scExtId2 := "sc-789"
		rg := &multidomainModels.ResourceGroup{
			PlacementTargets: []multidomainModels.TargetDetails{
				{
					ClusterExtId: &clusterExtId,
					Capabilities: []commonConfig.KVPair{
						createKVPair("cluster_name", "my-cluster"),
					},
					StorageContainers: []multidomainModels.StorageContainerDetails{
						{
							ExtId: &scExtId1,
							Capabilities: []commonConfig.KVPair{
								createKVPair("vstore_name_list", "container-one"),
							},
						},
						{
							ExtId: &scExtId2,
							Capabilities: []commonConfig.KVPair{
								createKVPair("vstore_name_list", "container-two"),
							},
						},
					},
				},
			},
		}

		result := extractStorageContainers(rg)
		require.Len(t, result, 2)
		assert.Equal(t, "sc-456", result[0].ExtId)
		assert.Equal(t, "container-one", result[0].Name)
		assert.Equal(t, "cluster-123", result[0].PrismElement.ExtId)
		assert.Equal(t, "my-cluster", result[0].PrismElement.Name)
		assert.Equal(t, "sc-789", result[1].ExtId)
		assert.Equal(t, "container-two", result[1].Name)
		assert.Equal(t, "cluster-123", result[1].PrismElement.ExtId)
		assert.Equal(t, "my-cluster", result[1].PrismElement.Name)
	})

	t.Run("StorageContainersOnMultiplePEs", func(t *testing.T) {
		clusterExtId1 := "cluster-123"
		clusterExtId2 := "cluster-456"
		scExtId1 := "sc-111"
		scExtId2 := "sc-222"
		rg := &multidomainModels.ResourceGroup{
			PlacementTargets: []multidomainModels.TargetDetails{
				{
					ClusterExtId: &clusterExtId1,
					Capabilities: []commonConfig.KVPair{
						createKVPair("cluster_name", "cluster-one"),
					},
					StorageContainers: []multidomainModels.StorageContainerDetails{
						{
							ExtId: &scExtId1,
							Capabilities: []commonConfig.KVPair{
								createKVPair("vstore_name_list", "container-on-pe1"),
							},
						},
					},
				},
				{
					ClusterExtId: &clusterExtId2,
					Capabilities: []commonConfig.KVPair{
						createKVPair("cluster_name", "cluster-two"),
					},
					StorageContainers: []multidomainModels.StorageContainerDetails{
						{
							ExtId: &scExtId2,
							Capabilities: []commonConfig.KVPair{
								createKVPair("vstore_name_list", "container-on-pe2"),
							},
						},
					},
				},
			},
		}

		result := extractStorageContainers(rg)
		require.Len(t, result, 2)
		assert.Equal(t, "sc-111", result[0].ExtId)
		assert.Equal(t, "container-on-pe1", result[0].Name)
		assert.Equal(t, "cluster-123", result[0].PrismElement.ExtId)
		assert.Equal(t, "cluster-one", result[0].PrismElement.Name)
		assert.Equal(t, "sc-222", result[1].ExtId)
		assert.Equal(t, "container-on-pe2", result[1].Name)
		assert.Equal(t, "cluster-456", result[1].PrismElement.ExtId)
		assert.Equal(t, "cluster-two", result[1].PrismElement.Name)
	})

	t.Run("StorageContainerWithNilExtId", func(t *testing.T) {
		clusterExtId := "cluster-123"
		scExtId := "sc-valid"
		rg := &multidomainModels.ResourceGroup{
			PlacementTargets: []multidomainModels.TargetDetails{
				{
					ClusterExtId: &clusterExtId,
					Capabilities: []commonConfig.KVPair{
						createKVPair("cluster_name", "my-cluster"),
					},
					StorageContainers: []multidomainModels.StorageContainerDetails{
						{
							ExtId: nil,
						},
						{
							ExtId: &scExtId,
							Capabilities: []commonConfig.KVPair{
								createKVPair("vstore_name_list", "valid-container"),
							},
						},
					},
				},
			},
		}

		result := extractStorageContainers(rg)
		require.Len(t, result, 1)
		assert.Equal(t, "sc-valid", result[0].ExtId)
		assert.Equal(t, "valid-container", result[0].Name)
	})

	t.Run("StorageContainerWithoutVstoreNameCapability", func(t *testing.T) {
		clusterExtId := "cluster-123"
		scExtId := "sc-456"
		rg := &multidomainModels.ResourceGroup{
			PlacementTargets: []multidomainModels.TargetDetails{
				{
					ClusterExtId: &clusterExtId,
					Capabilities: []commonConfig.KVPair{
						createKVPair("cluster_name", "my-cluster"),
					},
					StorageContainers: []multidomainModels.StorageContainerDetails{
						{
							ExtId: &scExtId,
							Capabilities: []commonConfig.KVPair{
								createKVPair("other_capability", "some-value"),
							},
						},
					},
				},
			},
		}

		result := extractStorageContainers(rg)
		require.Len(t, result, 1)
		assert.Equal(t, "sc-456", result[0].ExtId)
		assert.Empty(t, result[0].Name)
		assert.Equal(t, "cluster-123", result[0].PrismElement.ExtId)
		assert.Equal(t, "my-cluster", result[0].PrismElement.Name)
	})
}
