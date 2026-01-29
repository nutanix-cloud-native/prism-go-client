package converged

// VMs defines the interface for virtual machine operations.
// It supports get, list, create, update, delete (sync and async) and power actions.
type VMs[VM any] interface {
	Getter[VM]
	Lister[VM]
	Creator[VM]
	Updater[VM]
	Deleter[VM]
	AsyncCreator[VM]
	AsyncUpdater[VM]
	AsyncDeleter[VM]

	// PowerOnVM powers on the VM with the given UUID.
	PowerOnVM(uuid string) (Operation[VM], error)

	// PowerOffVM powers off the VM with the given UUID.
	PowerOffVM(uuid string) (Operation[VM], error)
}
