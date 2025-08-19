package facadetest

import (
	"context"
	"errors"
	"fmt"
	"testing"

	prismclientv4facade "github.com/nutanix-cloud-native/prism-go-client/facade/v4"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	"github.com/nutanix-cloud-native/prism-go-client/utils"
	prismclientv3 "github.com/nutanix-cloud-native/prism-go-client/v3"
	clustermgmtconfig "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	vmmconfig "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	vmmcontent "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
	"k8s.io/utils/ptr"
)

var (
	v4FacadeClient *prismclientv4facade.FacadeV4Client
	v3Client       *prismclientv3.Client
)

func initializeClients(t *testing.T) error {
	var err error
	credentials := testhelpers.CredentialsFromEnvironment(t)
	fmt.Printf("Using credentials for host: %s", credentials.Endpoint)

	v3Client, err = prismclientv3.NewV3Client(credentials)
	if err != nil {
		t.Errorf("Error creating v3 client: %v", err)
	}

	v4FacadeClient, err = prismclientv4facade.NewFacadeV4Client(credentials)
	if err != nil {
		t.Errorf("Error creating facade v4 client: %v", err)
	}

	return nil
}

func TestCreateVmCases(t *testing.T) {
	err := initializeClients(t)
	if err != nil {
		t.Errorf("failed to intialize prism clients, error: %v", err)
	}

	response, err := v3Client.V3.ListAllCluster(context.Background(), "")
	if err != nil {
		t.Errorf("unable to list all clusters: %v", err)
	}

	peUUID := ""
	peGpuUUID := ""
	for _, pe := range response.Entities {
		if pe.Spec.Resources.Config.GpuDriverVersion == "" {
			if peUUID == "" {
				peUUID = *pe.Metadata.UUID
			}
		} else {
			if peGpuUUID == "" {
				peGpuUUID = *pe.Metadata.UUID
			}
		}
	}
	if peUUID == "" || peGpuUUID == "" {
		t.Errorf("valid PEs not found")
	}

	baseVmConfig := vmmconfig.Vm{
		NumCoresPerSocket: utils.IntPtr(1),
		NumSockets:        utils.IntPtr(1),
		MemorySizeBytes:   utils.Int64Ptr(1e10),
	}

	tests := []struct {
		name            string
		vmName          string
		peUUID          string
		vmConfigBuilder func() (*vmmconfig.Vm, error)
	}{
		{
			name:   "create vm with data disks",
			vmName: "test-createvm-facade-v4-data-disks",
			peUUID: peUUID,
			vmConfigBuilder: func() (*vmmconfig.Vm, error) {
				vmConfig := ptr.To(baseVmConfig)

				images, err := v4FacadeClient.ListImages()
				if err != nil {
					return nil, err
				}

				imageExtId := ""
				for _, image := range images {
					if *image.Type != vmmcontent.IMAGETYPE_UNKNOWN && *image.Type != vmmcontent.IMAGETYPE_REDACTED {
						imageExtId = *image.ExtId
						break
					}
				}
				if imageExtId == "" {
					return nil, errors.New("no images found")
				}

				vmConfig.Disks = []vmmconfig.Disk{
					*newDisk(imageExtId, true, vmmconfig.DISKBUSTYPE_SCSI, 1),
				}
				vmConfig.CdRoms = []vmmconfig.CdRom{
					*newCdRom(imageExtId, true, vmmconfig.CDROMBUSTYPE_IDE, 0),
				}

				return vmConfig, nil
			},
		},
		{
			name:   "create vm with uefi boot type",
			vmName: "test-createvm-facade-v4-uefi",
			peUUID: peUUID,
			vmConfigBuilder: func() (*vmmconfig.Vm, error) {
				vmConfig := ptr.To(baseVmConfig)

				vmConfig.BootConfig = vmmconfig.NewOneOfVmBootConfig()
				vmConfig.BootConfig.SetValue(*vmmconfig.NewUefiBoot())

				return vmConfig, nil
			},
		},
		{
			name:   "create vm with attached gpu",
			vmName: "test-createvm-facade-v4-gpu",
			peUUID: peGpuUUID,
			vmConfigBuilder: func() (*vmmconfig.Vm, error) {
				vmConfig := ptr.To(baseVmConfig)

				gpus, err := v4FacadeClient.ListClusterVirtualGPUs(peGpuUUID)
				if err != nil {
					return nil, err
				}

				var foundGpu *clustermgmtconfig.VirtualGpuProfile
				for _, gpu := range gpus {
					if *gpu.VirtualGpuConfig.IsInUse {
						continue
					}

					foundGpu = &gpu
					break
				}
				if foundGpu == nil {
					return nil, errors.New("no gpus found")
				}

				deviceId := utils.IntPtr(int(*foundGpu.VirtualGpuConfig.DeviceId))
				mode := vmmconfig.GPUMODE_VIRTUAL.Ref()
				vendor := vendorStringToV4Model(foundGpu.VirtualGpuConfig.VendorName)

				vmConfig.Gpus = []vmmconfig.Gpu{
					*newGpu(deviceId, mode, vendor),
				}

				return vmConfig, nil
			},
		},
		{
			name:   "create vm attached to a project",
			vmName: "test-createvm-facade-v4-project",
			peUUID: peUUID,
			vmConfigBuilder: func() (*vmmconfig.Vm, error) {
				vmConfig := ptr.To(baseVmConfig)

				// TODO use facade v4
				projects, err := v3Client.V3.ListAllProject(context.Background(), "")
				if err != nil {
					return nil, err
				}
				if len(projects.Entities) == 0 {
					return nil, errors.New("no project found")
				}

				vmConfig.Project = newProjectRef(projects.Entities[0].Metadata.UUID)

				return vmConfig, nil
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vmConfig, err := test.vmConfigBuilder()
			if err != nil {
				t.Errorf("error: %v", err)
			}

			vmConfig.Name = &test.vmName
			vmConfig.Cluster = vmmconfig.NewClusterReference()
			vmConfig.Cluster.ExtId = utils.StringPtr(test.peUUID)

			// create VM
			newVmTaskWaiter, err := v4FacadeClient.CreateVM(vmConfig)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			taskUUID := newVmTaskWaiter.GetTaskUUID()
			if taskUUID == "" {
				t.Errorf("error: %v", err)
			}
			vmEntities, err := newVmTaskWaiter.WaitForTaskCompletion()
			if err != nil {
				t.Errorf("error: %v", err)
			}
			errs := newVmTaskWaiter.GetTaskErrors()
			if len(errs) > 0 {
				t.Errorf("errors: %v", err)
			}
			vmUUID := vmEntities[0].ExtId
			t.Logf("created - VM with name: %s, UUID: %s", *vmConfig.Name, *vmUUID)

			// power on the VM
			newVmTaskWaiter, err = v4FacadeClient.PowerOnVM(*vmUUID)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			taskUUID = newVmTaskWaiter.GetTaskUUID()
			if taskUUID == "" {
				t.Errorf("error: %v", err)
			}
			_, err = newVmTaskWaiter.WaitForTaskCompletion()
			if err != nil {
				t.Errorf("error: %v", err)
			}
			t.Logf("power on - VM with name: %s, UUID: %s", *vmConfig.Name, *vmUUID)

			// delete the VM
			waiter, err := v4FacadeClient.DeleteVM(*vmUUID)
			if err != nil {
				t.Fatal("failed to delete vm, errors:", err)
			}
			_, err = waiter.WaitForTaskCompletion()
			if err != nil {
				t.Fatal("failed to delete vm, errors:", err)
			}
			errs = waiter.GetTaskErrors()
			if len(errs) > 0 {
				t.Fatal("failed to delete vm, errors:", errs)
			}
			t.Logf("deleted - VM with name: %s, UUID: %s", *vmConfig.Name, *vmUUID)
		})
	}
}

func vendorStringToV4Model(vendor *string) *vmmconfig.GpuVendor {
	switch *vendor {
	case "kNvidia":
		return vmmconfig.GPUVENDOR_NVIDIA.Ref()
	case "kIntel":
		return vmmconfig.GPUVENDOR_INTEL.Ref()
	case "kAmd":
		return vmmconfig.GPUVENDOR_AMD.Ref()
	default:
		return vmmconfig.GPUVENDOR_UNKNOWN.Ref()
	}
}

func newDisk(imageExtId string, flashModeEnabled bool, busType vmmconfig.DiskBusType, deviceIndex int) *vmmconfig.Disk {
	vmDisk := vmmconfig.NewVmDisk()
	vmDisk.DiskSizeBytes = utils.Int64Ptr(1e10)

	imageRef := vmmconfig.NewImageReference()
	imageRef.ImageExtId = utils.StringPtr(imageExtId)
	vmDisk.DataSource = vmmconfig.NewDataSource()
	vmDisk.DataSource.SetReference(*imageRef)
	vmDisk.DataSource.ReferenceItemDiscriminator_ = nil

	vmDisk.StorageConfig = vmmconfig.NewVmDiskStorageConfig()
	vmDisk.StorageConfig.IsFlashModeEnabled = utils.BoolPtr(flashModeEnabled)

	disk := vmmconfig.NewDisk()
	disk.ExtId = vmDisk.DiskExtId
	disk.DiskAddress = vmmconfig.NewDiskAddress()
	disk.DiskAddress.BusType = &busType
	disk.DiskAddress.Index = &deviceIndex
	disk.SetBackingInfo(*vmDisk)

	return disk
}

func newCdRom(imageExtId string, flashModeEnabled bool, busType vmmconfig.CdRomBusType, deviceIndex int) *vmmconfig.CdRom {
	vmDisk := vmmconfig.NewVmDisk()
	vmDisk.DiskSizeBytes = utils.Int64Ptr(1e10)

	imageRef := vmmconfig.NewImageReference()
	imageRef.ImageExtId = utils.StringPtr(imageExtId)
	vmDisk.DataSource = vmmconfig.NewDataSource()
	vmDisk.DataSource.SetReference(*imageRef)
	vmDisk.DataSource.ReferenceItemDiscriminator_ = nil

	vmDisk.StorageConfig = vmmconfig.NewVmDiskStorageConfig()
	vmDisk.StorageConfig.IsFlashModeEnabled = utils.BoolPtr(flashModeEnabled)

	cdRom := vmmconfig.NewCdRom()
	cdRom.ExtId = vmDisk.DiskExtId
	cdRom.DiskAddress = vmmconfig.NewCdRomAddress()
	cdRom.DiskAddress.BusType = &busType
	cdRom.DiskAddress.Index = &deviceIndex
	cdRom.BackingInfo = vmDisk

	return cdRom
}

func newGpu(deviceId *int, mode *vmmconfig.GpuMode, vendor *vmmconfig.GpuVendor) *vmmconfig.Gpu {
	gpu := vmmconfig.NewGpu()
	gpu.DeviceId = deviceId
	gpu.Mode = mode
	gpu.Vendor = vendor

	return gpu
}

func newProjectRef(projectUUID *string) *vmmconfig.ProjectReference {
	projRef := vmmconfig.NewProjectReference()
	projRef.ExtId = projectUUID

	return projRef
}
