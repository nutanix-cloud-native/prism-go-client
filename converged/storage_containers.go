package converged

// StorageContainers defines the interface for managing storage containers.
type StorageContainers[StorageContainer, StorageContainerListParams any] interface {
	// Getter is the interface for Get operations.
	Getter[StorageContainer]

	// Lister is the interface for List operations.
	Lister[StorageContainer, StorageContainerListParams]
}
