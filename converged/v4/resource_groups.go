package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	multidomainModels "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/config"
	resourceGroupRequest "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/request/resourcegroups"
)

// Capability key constants for extracting information from placement targets.
const (
	capabilityClusterName = "cluster_name"
	capabilityVstoreName  = "vstore_name_list"
)

// ResourceGroupsService provides implementation for ResourceGroups API operations.
type ResourceGroupsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewResourceGroupsService creates a new ResourceGroupsService instance.
func NewResourceGroupsService(client *v4prismGoClient.Client) *ResourceGroupsService {
	return &ResourceGroupsService{client: client, entitiesName: "resource group"}
}

// Get returns the resource group for the given UUID.
func (s *ResourceGroupsService) Get(ctx context.Context, uuid string) (*multidomainModels.ResourceGroup, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	return GenericGetEntity[*multidomainModels.GetResourceGroupApiResponse, multidomainModels.ResourceGroup](
		func() (*multidomainModels.GetResourceGroupApiResponse, error) {
			return s.client.ResourceGroupsApiInstance.ServiceClient.GetResourceGroupById(ctx, &resourceGroupRequest.GetResourceGroupByIdRequest{
				ExtId: &uuid,
			})
		},
		s.entitiesName,
	)
}

// List returns a list of resource groups.
func (s *ResourceGroupsService) List(ctx context.Context, opts ...converged.ODataOption) ([]multidomainModels.ResourceGroup, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing ResourceGroups")
	}

	return GenericListEntities[*multidomainModels.ListResourceGroupsApiResponse, multidomainModels.ResourceGroup](
		func(reqParams *V4ODataParams) (*multidomainModels.ListResourceGroupsApiResponse, error) {
			return s.client.ResourceGroupsApiInstance.ServiceClient.ListResourceGroups(ctx, &resourceGroupRequest.ListResourceGroupsRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Select_:  reqParams.Select,
			})
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing resource groups.
func (s *ResourceGroupsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[multidomainModels.ResourceGroup] {
	if s.client == nil {
		return nil
	}
	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return ErrorOnlyIterator[multidomainModels.ResourceGroup](fmt.Errorf("failed to convert options to V4ODataParams: %w", err))
	}
	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return ErrorOnlyIterator[multidomainModels.ResourceGroup](fmt.Errorf("apply and expand options are not supported for listing ResourceGroups"))
	}
	return GenericNewIterator[*multidomainModels.ListResourceGroupsApiResponse, multidomainModels.ResourceGroup](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*multidomainModels.ListResourceGroupsApiResponse, error) {
			return s.client.ResourceGroupsApiInstance.ServiceClient.ListResourceGroups(ctx, &resourceGroupRequest.ListResourceGroupsRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Select_:  reqParams.Select,
			})
		},
		opts,
		s.entitiesName,
	)
}

// ListPrismElements fetches a ResourceGroup and extracts a list of Prism Elements (clusters) from its placement targets.
// The PE name is extracted from the "cluster_name" capability.
func (s *ResourceGroupsService) ListPrismElements(ctx context.Context, resourceGroupUUID string) ([]converged.PrismElementInfo, error) {
	rg, err := s.Get(ctx, resourceGroupUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get resource group: %w", err)
	}

	return extractPrismElements(rg), nil
}

// ListStorageContainers fetches a ResourceGroup and extracts a list of Storage Containers from its placement targets.
// The storage container name is extracted from the "vstore_name_list" capability.
// Each storage container includes its associated PE name.
func (s *ResourceGroupsService) ListStorageContainers(ctx context.Context, resourceGroupUUID string) ([]converged.StorageContainerInfo, error) {
	rg, err := s.Get(ctx, resourceGroupUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get resource group: %w", err)
	}

	return extractStorageContainers(rg), nil
}

// extractPrismElements extracts PE info from a ResourceGroup's placement targets.
func extractPrismElements(rg *multidomainModels.ResourceGroup) []converged.PrismElementInfo {
	if rg == nil || rg.PlacementTargets == nil {
		return nil
	}

	result := make([]converged.PrismElementInfo, 0, len(rg.PlacementTargets))
	for _, target := range rg.PlacementTargets {
		if target.ClusterExtId == nil {
			continue
		}

		pe := converged.PrismElementInfo{
			ExtId: *target.ClusterExtId,
		}

		// Extract cluster name from capabilities
		for _, cap := range target.Capabilities {
			if cap.Name != nil && *cap.Name == capabilityClusterName && cap.Value != nil {
				if strVal, ok := cap.Value.GetValue().(string); ok {
					pe.Name = strVal
				}
				break
			}
		}

		result = append(result, pe)
	}

	return result
}

// extractStorageContainers extracts storage container info from a ResourceGroup's placement targets.
func extractStorageContainers(rg *multidomainModels.ResourceGroup) []converged.StorageContainerInfo {
	if rg == nil || rg.PlacementTargets == nil {
		return nil
	}

	var result []converged.StorageContainerInfo
	for _, target := range rg.PlacementTargets {
		peInfo := converged.PrismElementInfo{}

		// Extract PE ExtId
		if target.ClusterExtId != nil {
			peInfo.ExtId = *target.ClusterExtId
		}

		// Extract PE name from capabilities
		for _, cap := range target.Capabilities {
			if cap.Name != nil && *cap.Name == capabilityClusterName && cap.Value != nil {
				if strVal, ok := cap.Value.GetValue().(string); ok {
					peInfo.Name = strVal
				}
				break
			}
		}

		// Process storage containers for this placement target
		for _, sc := range target.StorageContainers {
			if sc.ExtId == nil {
				continue
			}

			scInfo := converged.StorageContainerInfo{
				ExtId:        *sc.ExtId,
				PrismElement: peInfo,
			}

			// Extract storage container name from capabilities (vstore_name_list)
			for _, cap := range sc.Capabilities {
				if cap.Name != nil && *cap.Name == capabilityVstoreName && cap.Value != nil {
					if strVal, ok := cap.Value.GetValue().(string); ok {
						scInfo.Name = strVal
					}
					break
				}
			}

			result = append(result, scInfo)
		}
	}

	return result
}
