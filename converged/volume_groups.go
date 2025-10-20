package converged

import "context"

// Volumes is the interface for the Volumes service.
type VolumeGroups[VolumeGroup, VmAttachment any] interface {
	// Getter is the interface for Get operations.
	Getter[VolumeGroup]

	// Lister is the interface for List operations.
	Lister[VolumeGroup]

	// AttachToVMAsync attaches a volume group to a VM asynchronously.
	AttachToVMAsync(ctx context.Context, volumeGroupUUID string, vmAttachment VmAttachment) (Operation[VolumeGroup], error)

	// AttachToVM attaches a volume group to a VM.
	AttachToVM(ctx context.Context, volumeGroupUUID string, vmAttachment VmAttachment) (*VolumeGroup, error)

	// DetachFromVMAsync detaches a volume group from a VM asynchronously.
	DetachFromVMAsync(ctx context.Context, volumeGroupUUID string, vmAttachment VmAttachment) (Operation[VolumeGroup], error)

	// DetachFromVM detaches a volume group from a VM.
	DetachFromVM(ctx context.Context, volumeGroupUUID string, vmAttachment VmAttachment) (*VolumeGroup, error)

	// Additional methods can be added here as needed.
}
