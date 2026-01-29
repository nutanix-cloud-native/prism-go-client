package converged

import "context"

// VolumeGroups defines the interface for volume group operations.
// Volume groups are collections of volumes that can be attached to or detached from VMs.
type VolumeGroups[VolumeGroup, VmAttachment any] interface {
	Getter[VolumeGroup]
	Lister[VolumeGroup]

	// AttachToVMAsync attaches a volume group to a VM asynchronously.
	AttachToVMAsync(ctx context.Context, volumeGroupUUID string, vmAttachment VmAttachment) (Operation[VolumeGroup], error)

	// AttachToVM attaches a volume group to a VM.
	AttachToVM(ctx context.Context, volumeGroupUUID string, vmAttachment VmAttachment) (*VolumeGroup, error)

	// DetachFromVMAsync detaches a volume group from a VM asynchronously.
	DetachFromVMAsync(ctx context.Context, volumeGroupUUID string, vmAttachment VmAttachment) (Operation[VolumeGroup], error)

	// DetachFromVM detaches a volume group from a VM.
	DetachFromVM(ctx context.Context, volumeGroupUUID string, vmAttachment VmAttachment) (*VolumeGroup, error)
}
