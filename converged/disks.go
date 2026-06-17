package converged

// Disks defines the interface for managing disks.
type Disks[Disk any] interface {
	// Getter is the interface for Get operations.
	Getter[Disk]

	// Lister is the interface for List operations.
	Lister[Disk]

	// Deleter is the interface for blocking Delete operations.
	Deleter[Disk]

	// AsyncDeleter is the interface for async Delete operations.
	AsyncDeleter[Disk]
}
