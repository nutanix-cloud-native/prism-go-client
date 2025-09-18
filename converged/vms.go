package converged

// VMs is the interface for the VMs service.
type VMs[VM any] interface {
	// Getter is the interface for Get operations.
	Getter[VM]

	// Lister is the interface for List operations.
	Lister[VM]

	// Creator is the interface for Create operations.
	Creator[VM]

	// Updater is the interface for Update operations.
	Updater[VM]

	// Deleter is the interface for Delete operations.
	Deleter[VM]

	// AsyncCreator is the interface for Async Create operations.
	AsyncCreator[VM]

	// AsyncUpdater is the interface for Async Update operations.
	AsyncUpdater[VM]

	// AsyncDeleter is the interface for Async Delete operations.
	AsyncDeleter[VM]

	// PowerOnVM powers on the VM with the given UUID.
	PowerOnVM(uuid string) (Operation[VM], error)

	// PowerOffVM powers off the VM with the given UUID.
	PowerOffVM(uuid string) (Operation[VM], error)

	// Additional methods can be added here as needed.
}
