package converged

import "context"

// VMs is the interface for the VMs service.
type VMs[VM, NIC, VMDisk any] interface {
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
	PowerOnVM(ctx context.Context, uuid string) (Operation[VM], error)

	// PowerOffVM powers off the VM with the given UUID.
	PowerOffVM(ctx context.Context, uuid string) (Operation[VM], error)

	// AddVmCustomAttributes adds custom attributes to the VM with the given UUID.
	AddVmCustomAttributes(ctx context.Context, uuid string, customAttributes []string) (*VM, error)

	// AddVmCustomAttributesAsync adds custom attributes to the VM asynchronously.
	AddVmCustomAttributesAsync(ctx context.Context, uuid string, customAttributes []string) (Operation[VM], error)

	// RemoveVmCustomAttributes removes custom attributes from the VM with the given UUID.
	RemoveVmCustomAttributes(ctx context.Context, uuid string, customAttributes []string) (*VM, error)

	// RemoveVmCustomAttributesAsync removes custom attributes from the VM asynchronously.
	RemoveVmCustomAttributesAsync(ctx context.Context, uuid string, customAttributes []string) (Operation[VM], error)

	// DeleteCdRom deletes a CD-ROM device from the VM identified by uuid
	// and waits for the asynchronous task to complete.
	DeleteCdRom(ctx context.Context, uuid string, cdRomUUID string) error

	// DeleteCdRomAsync starts an asynchronous CD-ROM deletion on the VM.
	DeleteCdRomAsync(ctx context.Context, uuid string, cdRomUUID string) (Operation[NoEntity], error)

	// GenerateConsoleToken obtains a JWT token and WebSocket URI for VNC
	// console access to the VM identified by uuid.
	GenerateConsoleToken(ctx context.Context, uuid string) (*VMConsoleToken, error)

	// GenerateConsoleTokenAsync starts the generate-console-token API call
	// asynchronously and returns an Operation to track the task.
	GenerateConsoleTokenAsync(ctx context.Context, uuid string) (Operation[NoEntity], error)

	// GetVMByBiosUUID returns the VM matching the given BIOS UUID. When multiple VMs
	// share the same BIOS UUID (e.g. due to cloning), it resolves the chain by returning
	// the leaf VM that is not referenced as a source by any other VM.
	GetVMByBiosUUID(ctx context.Context, biosUUID string) (*VM, error)

	// ListNicsByVmId lists the NICs attached to the VM with the given UUID.
	ListNicsByVmId(ctx context.Context, vmUUID string) ([]NIC, error)

	// AddDisk adds a disk to the VM and returns an async operation for the task.
	AddDisk(ctx context.Context, vmUUID string, disk *VMDisk) (Operation[NoEntity], error)

	// GrowDisk updates an existing VM disk and returns an async operation for the task.
	GrowDisk(ctx context.Context, vmUUID string, diskUUID string, disk *VMDisk) (Operation[NoEntity], error)

	// DeleteDisk deletes a disk from the VM and returns an async operation for the task.
	DeleteDisk(ctx context.Context, vmUUID string, diskUUID string) (Operation[NoEntity], error)

	// AddNIC adds a NIC to the VM and returns an async operation for the task.
	AddNIC(ctx context.Context, vmUUID string, nic *NIC) (Operation[NoEntity], error)

	// DeleteNIC deletes a NIC from the VM and returns an async operation for the task.
	DeleteNIC(ctx context.Context, vmUUID string, nicUUID string) (Operation[NoEntity], error)
}
