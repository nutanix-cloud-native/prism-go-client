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

	v4FacadeClient, err = prismclientv4facade.NewFacadeV4Client(context.Background(), credentials)
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

	peExtId, peGpuExtId, err := getPeExtIds(t)
	if err != nil {
		t.Errorf("failed to get PEs, error: %v", err)
	}

	baseVmConfig := vmmconfig.Vm{
		NumCoresPerSocket: utils.IntPtr(1),
		NumSockets:        utils.IntPtr(1),
		MemorySizeBytes:   utils.Int64Ptr(1e10),
	}
	systemDisk, err := newSystemDisk()
	if err != nil {
		t.Errorf("failed to load system disk, error: %v", err)
	}
	baseVmConfig.Disks = []vmmconfig.Disk{*systemDisk}

	tests := []struct {
		name            string
		vmName          string
		peExtId         *string
		vmConfigBuilder func() (*vmmconfig.Vm, error)
	}{
		{
			name:    "create vm with data disks",
			vmName:  "test-createvm-facade-v4-data-disks",
			peExtId: peExtId,
			vmConfigBuilder: func() (*vmmconfig.Vm, error) {
				vmConfig := ptr.To(baseVmConfig)

				images, err := v4FacadeClient.ListAllImages(context.Background(), nil, nil, nil)
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

				vmConfig.Disks = append(vmConfig.Disks,
					*newDisk(imageExtId, true, vmmconfig.DISKBUSTYPE_SCSI, 1),
				)
				vmConfig.CdRoms = []vmmconfig.CdRom{
					*newCdRom(imageExtId, true, vmmconfig.CDROMBUSTYPE_IDE, 0),
				}

				return vmConfig, nil
			},
		},
		{
			name:    "create vm with uefi boot type",
			vmName:  "test-createvm-facade-v4-uefi",
			peExtId: peExtId,
			vmConfigBuilder: func() (*vmmconfig.Vm, error) {
				vmConfig := ptr.To(baseVmConfig)

				vmConfig.BootConfig = vmmconfig.NewOneOfVmBootConfig()
				vmConfig.BootConfig.SetValue(*vmmconfig.NewUefiBoot())

				return vmConfig, nil
			},
		},
		{
			name:    "create vm with attached gpu",
			vmName:  "test-createvm-facade-v4-gpu",
			peExtId: peGpuExtId,
			vmConfigBuilder: func() (*vmmconfig.Vm, error) {
				vmConfig := ptr.To(baseVmConfig)

				gpus, err := v4FacadeClient.ListClusterVirtualGPUs(context.Background(), *peGpuExtId)
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
			name:    "create vm attached to a project",
			vmName:  "test-createvm-facade-v4-project",
			peExtId: peExtId,
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
			vmConfig.Cluster.ExtId = test.peExtId

			// create VM
			newVmTaskWaiter, err := v4FacadeClient.CreateVM(context.Background(), vmConfig)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			taskUUID := newVmTaskWaiter.GetTaskUUID()
			if taskUUID == "" {
				t.Errorf("error: %v", err)
			}
			vmEntities, err := newVmTaskWaiter.WaitForTaskCompletion(context.Background())
			if err != nil {
				t.Errorf("error: %v", err)
			}
			errs := newVmTaskWaiter.GetTaskErrors()
			if len(errs) > 0 {
				t.Errorf("errors: %v", err)
			}
			vmExtId := vmEntities[0].ExtId
			t.Logf("created - VM with name: %s, extId: %s", *vmConfig.Name, *vmExtId)

			// power on the VM
			newVmTaskWaiter, err = v4FacadeClient.PowerOnVM(context.Background(), *vmExtId)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			taskUUID = newVmTaskWaiter.GetTaskUUID()
			if taskUUID == "" {
				t.Errorf("error: %v", err)
			}
			_, err = newVmTaskWaiter.WaitForTaskCompletion(context.Background())
			if err != nil {
				t.Errorf("error: %v", err)
			}
			t.Logf("power on - VM with name: %s, extId: %s", *vmConfig.Name, *vmExtId)

			// delete the VM
			waiter, err := v4FacadeClient.DeleteVM(context.Background(), *vmExtId)
			if err != nil {
				t.Fatal("failed to delete vm, errors:", err)
			}
			_, err = waiter.WaitForTaskCompletion(context.Background())
			if err != nil {
				t.Fatal("failed to delete vm, errors:", err)
			}
			errs = waiter.GetTaskErrors()
			if len(errs) > 0 {
				t.Fatal("failed to delete vm, errors:", errs)
			}
			t.Logf("deleted - VM with name: %s, extId: %s", *vmConfig.Name, *vmExtId)
		})
	}
}

func getPeExtIds(t *testing.T) (*string, *string, error) {
	PEs, err := v4FacadeClient.ListAllClusters(context.Background(), nil, nil, nil, nil)
	if err != nil {
		t.Errorf("unable to list all clusters: %v", err)
	}

	var peExtId *string
	var peGpuExtId *string
	for _, pe := range PEs {
		gpus, _ := v4FacadeClient.ListClusterVirtualGPUs(context.Background(), *pe.ExtId)

		if len(gpus) == 0 {
			if peExtId == nil {
				peExtId = pe.ExtId
			}
		} else {
			if peGpuExtId == nil {
				peGpuExtId = pe.ExtId
			}
		}
		if peExtId != nil && peGpuExtId != nil {
			break
		}
	}
	if peExtId == nil || peGpuExtId == nil {
		return nil, nil, errors.New("valid PEs not found")
	}

	return peExtId, peGpuExtId, nil
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

func newSystemDisk() (*vmmconfig.Disk, error) {
	images, err := v4FacadeClient.ListAllImages(context.Background(), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var osImageExtId *string
	for _, image := range images {
		if *image.Type != vmmcontent.IMAGETYPE_UNKNOWN && *image.Type != vmmcontent.IMAGETYPE_REDACTED {
			osImageExtId = image.ExtId
			break
		}
	}
	if osImageExtId == nil {
		return nil, errors.New("no image found")
	}

	vmDisk := vmmconfig.NewVmDisk()
	vmDisk.DiskSizeBytes = utils.Int64Ptr(1e10)

	imageRef := vmmconfig.NewImageReference()
	imageRef.ImageExtId = osImageExtId
	vmDisk.DataSource = vmmconfig.NewDataSource()
	vmDisk.DataSource.SetReference(*imageRef)
	vmDisk.DataSource.ReferenceItemDiscriminator_ = nil

	systemDisk := vmmconfig.NewDisk()
	systemDisk.SetBackingInfo(*vmDisk)

	return systemDisk, nil
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

func newProjectRef(projectExtId *string) *vmmconfig.ProjectReference {
	projRef := vmmconfig.NewProjectReference()
	projRef.ExtId = projectExtId

	return projRef
}
