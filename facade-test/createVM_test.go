package facadetest

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	v4 "github.com/nutanix-cloud-native/prism-go-client/facade/v4"
	vmmconfig "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

func GetCredentials() prismgoclient.Credentials {
	err := godotenv.Load("../../prism-client-dotenv")
	if err != nil {
		fmt.Println("Error loading .env file: ", err)
	}
	username := os.Getenv("NUTANIX_USER")
	password := os.Getenv("NUTANIX_PASSWORD")
	endpoint := os.Getenv("NUTANIX_ENDPOINT")
	return prismgoclient.Credentials{
		Endpoint:    endpoint,
		Username:    username,
		Password:    password,
		SessionAuth: true,
		Insecure:    true,
	}
}

var (
	v4FacadeClient *v4.FacadeV4Client
)

func init() {
	var err error
	credentials := GetCredentials()
	fmt.Printf("Using credentials for host: %s\n", credentials.Endpoint)

	v4FacadeClient, err = v4.NewFacadeV4Client(credentials)
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}
}

const (
	BootTypeUefiVMConfigFileName = "bootTypeUefiVmConfig.json"
	DataDisksVmConfigFileName    = "dataDisksVmConfig.json"
)

func loadFile[T any](filePath string) (out T, err error) {
	file, err := os.Open(fmt.Sprintf("./testdata/%s", filePath))
	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&out)
	return
}

func TestCreateVm(t *testing.T) {
	dataDisksVmConfig, err := loadFile[vmmconfig.Vm](DataDisksVmConfigFileName)
	if err != nil {
		t.Errorf("unable to load data disks test file, error: %v", err)
	}
	bootTypeUefiVmConfig, err := loadFile[vmmconfig.Vm](BootTypeUefiVMConfigFileName)
	if err != nil {
		t.Errorf("unable to load boot type test file, error: %v", err)
	}

	tests := []struct {
		name     string
		vmConfig *vmmconfig.Vm
	}{
		{
			name:     "VM with data disks",
			vmConfig: &dataDisksVmConfig,
		},
		{
			name:     "VM with UEFI boot type",
			vmConfig: &bootTypeUefiVmConfig,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vmName := *test.vmConfig.Name

			// create VM
			newVmTaskWaiter, err := v4FacadeClient.CreateVM(test.vmConfig)
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
			t.Logf("created - VM with name: %s, UUID: %s", vmName, *vmUUID)

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
			t.Logf("power on - VM with name: %s, UUID: %s", vmName, *vmUUID)

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
			t.Logf("deleted - VM with name: %s, UUID: %s", vmName, *vmUUID)
		})
	}
}
