package converged

import "context"

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

	// AddVmCustomAttributes adds custom attributes to the VM with the given UUID.
	AddVmCustomAttributes(ctx context.Context, uuid string, customAttributes []string) (*VM, error)

	// AddVmCustomAttributesAsync adds custom attributes to the VM asynchronously.
	AddVmCustomAttributesAsync(uuid string, customAttributes []string) (Operation[VM], error)

	// RemoveVmCustomAttributes removes custom attributes from the VM with the given UUID.
	RemoveVmCustomAttributes(ctx context.Context, uuid string, customAttributes []string) (*VM, error)

	// RemoveVmCustomAttributesAsync removes custom attributes from the VM asynchronously.
	RemoveVmCustomAttributesAsync(uuid string, customAttributes []string) (Operation[VM], error)

	// DeleteCdRom deletes a CD-ROM device from the VM identified by uuid
	// and waits for the asynchronous task to complete.
	DeleteCdRom(ctx context.Context, uuid string, cdRomUUID string) error

	// DeleteCdRomAsync starts an asynchronous CD-ROM deletion on the VM.
	DeleteCdRomAsync(uuid string, cdRomUUID string) (Operation[NoEntity], error)

	// GenerateConsoleToken obtains a JWT token and WebSocket URI for VNC
	// console access to the VM identified by uuid.
	GenerateConsoleToken(ctx context.Context, uuid string) (*VMConsoleToken, error)

	// GenerateConsoleTokenAsync starts the generate-console-token API call
	// asynchronously and returns an Operation to track the task.
	GenerateConsoleTokenAsync(uuid string) (Operation[NoEntity], error)
}
