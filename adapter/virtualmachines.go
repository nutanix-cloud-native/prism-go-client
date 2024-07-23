package adapter

import (
	"fmt"

	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// VirtualMachinesClient is the interface for the virtual machines
type VirtualMachinesClient interface {
	GetVM(vmUUID string) (*config.Vm, error)
}

// VM is the struct for the virtual machine
type VM struct {
}

// VirtualMachines returns the VirtualMachinesClient
func (c *client) VirtualMachines() VirtualMachinesClient {
	return &vmClient{
		client: c,
	}
}

// vmClient implements the VirtualMachinesClient
type vmClient struct {
	client *client
}

// type assertion to ensure vmClient implements VirtualMachinesClient
var _ VirtualMachinesClient = &vmClient{}

// GetVM returns the virtual machine
// TODO: This function can also be handled via v3 in absence of v4 support. Ideally the return value should be
// a struct with the task status and completion details that are not specific to v4.
func (c *vmClient) GetVM(vmUUID string) (*config.Vm, error) {
	if !c.client.envIsv4Compatible {
		return nil, fmt.Errorf("%w:%w", V4APINotSupported, MethodNotSupported)
	}

	vmResp, err := c.client.v4Client.VmApiInstance.GetVmById(&vmUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get virtual machine: %w", err)
	}

	data := vmResp.GetData()
	vm, ok := data.(config.Vm)
	if !ok {
		return nil, fmt.Errorf("failed to cast response to VM")
	}

	return &vm, nil
}
