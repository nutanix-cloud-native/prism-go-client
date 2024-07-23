package adapter

import (
	"context"
	"fmt"

	vmconfig "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	prismconfig "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/prism/v4/config"
	volumesconfig "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
)

// VolumesClient is the interface for the volumes
type VolumesClient interface {
	DetachVolumeGroupsFromVM(vmUUID string) error
}

// volumes implements the Volumes interface
type volumes struct {
	client *client
}

// type assertion to ensure volumes implements Volumes
var _ VolumesClient = &volumes{}

func (c *client) Volumes() VolumesClient {
	return &volumes{
		client: c,
	}
}

func (vc *volumes) DetachVolumeGroupsFromVM(vmUUID string) error {
	// DetachVolume is only supported in v4. If the environment is not v4 compatible, return an error
	if !vc.client.envIsv4Compatible {
		return fmt.Errorf("%w:%w", V4APINotSupported, MethodNotSupported)
	}

	vm, err := vc.client.VirtualMachines().GetVM(vmUUID)
	if err != nil {
		return fmt.Errorf("failed to get virtual machine: %w", err)
	}

	volumeGroupsToDetach := make([]string, 0)
	for _, disk := range vm.Disks {
		backingInfo, ok := disk.GetBackingInfo().(vmconfig.ADSFVolumeGroupReference)
		if !ok {
			continue
		}

		volumeGroupsToDetach = append(volumeGroupsToDetach, *backingInfo.VolumeGroupExtId)
	}

	// Detach the volume groups from the virtual machine
	for _, volumeGroup := range volumeGroupsToDetach {
		body := &volumesconfig.VmAttachment{
			ExtId: &vmUUID,
		}
		resp, err := vc.client.v4Client.VolumeGroupsApiInstance.DetachVm(&volumeGroup, body)
		if err != nil {
			return fmt.Errorf("failed to detach volume group %s from virtual machine %s: %w", volumeGroup, vmUUID, err)
		}

		data := resp.GetData()
		task, ok := data.(prismconfig.TaskReference)
		if !ok {
			return fmt.Errorf("failed to cast response to TaskReference")
		}

		// Wait for the task to complete
		if _, err := vc.client.Prism().WaitForTaskCompletion(context.TODO(), *task.ExtId); err != nil {
			return fmt.Errorf("failed to wait for task %s to complete: %w", *task.ExtId, err)
		}
	}

	return nil
}
