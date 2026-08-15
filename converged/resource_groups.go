package converged

import "context"

// PrismElementInfo represents basic information about a Prism Element (cluster).
type PrismElementInfo struct {
	ExtId string
	Name  string
}

// StorageContainerInfo represents basic information about a storage container and its associated PE.
type StorageContainerInfo struct {
	ExtId       string
	Name        string
	PrismElement PrismElementInfo
}

// ResourceGroups is the interface for Resource Groups operations.
type ResourceGroups[ResourceGroup any] interface {
	// Getter is the interface for Get operations.
	Getter[ResourceGroup]

	// Lister is the interface for List operations.
	Lister[ResourceGroup]

	// ListPrismElements fetches a ResourceGroup and extracts PE info from its placement targets.
	ListPrismElements(ctx context.Context, resourceGroupUUID string) ([]PrismElementInfo, error)

	// ListStorageContainers fetches a ResourceGroup and extracts storage container info from its placement targets.
	ListStorageContainers(ctx context.Context, resourceGroupUUID string) ([]StorageContainerInfo, error)
}
