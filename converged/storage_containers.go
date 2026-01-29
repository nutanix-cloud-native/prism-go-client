package converged

// StorageContainers defines the interface for storage container operations.
// Storage containers are logical groupings of storage for VMs and images.
type StorageContainers[StorageContainer any] interface {
	Getter[StorageContainer]
	Lister[StorageContainer]
}
