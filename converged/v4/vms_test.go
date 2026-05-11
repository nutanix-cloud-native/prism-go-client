// Package v4 provides tests for the converged client VMs service methods
// (custom attributes, CD-ROM deletion, console token generation).
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run TestVMs                         # all VM tests
//	go test -v ./converged/v4 -run TestVMsCustomAttributes_NilClient
//	go test -v ./converged/v4 -run TestVMsDeleteCdRom_NilClient
//	go test -v ./converged/v4 -run TestVMsGenerateConsoleToken_NilClient
//	go test -v ./converged/v4 -run TestVMsDeleteCdRomIntegration
//	go test -v ./converged/v4 -run TestVMsGenerateConsoleTokenIntegration
package v4

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// TestVMsCustomAttributes_NilClient tests nil client error handling for custom attributes methods
func TestVMsCustomAttributes_NilClient(t *testing.T) {
	service := NewVMsService(nil)
	require.NotNil(t, service)
	ctx := context.Background()

	t.Run("AddVmCustomAttributes", func(t *testing.T) {
		_, err := service.AddVmCustomAttributes(ctx, "test-id", []string{"attr1"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("AddVmCustomAttributesAsync", func(t *testing.T) {
		_, err := service.AddVmCustomAttributesAsync("test-id", []string{"attr1"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("RemoveVmCustomAttributes", func(t *testing.T) {
		_, err := service.RemoveVmCustomAttributes(ctx, "test-id", []string{"attr1"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("RemoveVmCustomAttributesAsync", func(t *testing.T) {
		_, err := service.RemoveVmCustomAttributesAsync("test-id", []string{"attr1"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})
}

// TestVMsCustomAttributesIntegration tests custom attributes with real Nutanix API calls
func TestVMsCustomAttributesIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	// Get a VM to test with
	vms, err := client.VMs.List(ctx)
	require.NoError(t, err)

	if len(vms) == 0 {
		t.Skip("No VMs available for testing")
	}

	vmUUID := *vms[0].ExtId
	require.NotEmpty(t, vmUUID)

	t.Run("AddVmCustomAttributes", func(t *testing.T) {
		vm, err := client.VMs.AddVmCustomAttributes(ctx, vmUUID, []string{"test-attr-1", "test-attr-2"})
		assert.NoError(t, err)
		assert.NotNil(t, vm)
		assert.Equal(t, vmUUID, *vm.ExtId)
	})

	t.Run("AddVmCustomAttributesAsync", func(t *testing.T) {
		operation, err := client.VMs.AddVmCustomAttributesAsync(vmUUID, []string{"test-attr-async"})
		assert.NoError(t, err)
		assert.NotNil(t, operation)

		result, err := operation.Wait(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("RemoveVmCustomAttributes", func(t *testing.T) {
		vm, err := client.VMs.RemoveVmCustomAttributes(ctx, vmUUID, []string{"test-attr-1", "test-attr-2"})
		assert.NoError(t, err)
		assert.NotNil(t, vm)
		assert.Equal(t, vmUUID, *vm.ExtId)
	})

	t.Run("RemoveVmCustomAttributesAsync", func(t *testing.T) {
		operation, err := client.VMs.RemoveVmCustomAttributesAsync(vmUUID, []string{"test-attr-async"})
		assert.NoError(t, err)
		assert.NotNil(t, operation)

		result, err := operation.Wait(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})
}

// TestVMsCustomAttributesErrorScenarios tests error handling scenarios
func TestVMsCustomAttributesErrorScenarios(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("InvalidUUID", func(t *testing.T) {
		_, err := client.VMs.AddVmCustomAttributes(ctx, "invalid-uuid-format", []string{"attr1"})
		assert.Error(t, err)

		_, err = client.VMs.AddVmCustomAttributesAsync("invalid-uuid-format", []string{"attr1"})
		assert.Error(t, err)

		_, err = client.VMs.RemoveVmCustomAttributes(ctx, "invalid-uuid-format", []string{"attr1"})
		assert.Error(t, err)

		_, err = client.VMs.RemoveVmCustomAttributesAsync("invalid-uuid-format", []string{"attr1"})
		assert.Error(t, err)
	})

	t.Run("NonExistentUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"

		_, err := client.VMs.AddVmCustomAttributes(ctx, nonExistentUUID, []string{"attr1"})
		assert.Error(t, err)

		_, err = client.VMs.AddVmCustomAttributesAsync(nonExistentUUID, []string{"attr1"})
		assert.Error(t, err)

		_, err = client.VMs.RemoveVmCustomAttributes(ctx, nonExistentUUID, []string{"attr1"})
		assert.Error(t, err)

		_, err = client.VMs.RemoveVmCustomAttributesAsync(nonExistentUUID, []string{"attr1"})
		assert.Error(t, err)
	})

	t.Run("EmptyAttributes", func(t *testing.T) {
		assert.NotPanics(t, func() {
			_, _ = client.VMs.AddVmCustomAttributesAsync("invalid-uuid", []string{})
		})

		assert.NotPanics(t, func() {
			_, _ = client.VMs.RemoveVmCustomAttributesAsync("invalid-uuid", []string{})
		})
	})
}

// TestVMsDeleteCdRom_NilClient tests nil client error handling for CD-ROM deletion methods
func TestVMsDeleteCdRom_NilClient(t *testing.T) {
	service := NewVMsService(nil)
	require.NotNil(t, service)
	ctx := context.Background()

	t.Run("DeleteCdRom", func(t *testing.T) {
		err := service.DeleteCdRom(ctx, "test-vm-id", "test-cdrom-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("DeleteCdRomAsync", func(t *testing.T) {
		_, err := service.DeleteCdRomAsync("test-vm-id", "test-cdrom-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})
}

// TestVMsDeleteCdRomErrorScenarios tests CD-ROM deletion with invalid and non-existent UUIDs
func TestVMsDeleteCdRomErrorScenarios(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("InvalidVMUUID", func(t *testing.T) {
		err := client.VMs.DeleteCdRom(ctx, "invalid-uuid", "cdrom-id")
		assert.Error(t, err)
	})

	t.Run("InvalidVMUUIDAsync", func(t *testing.T) {
		_, err := client.VMs.DeleteCdRomAsync("invalid-uuid", "cdrom-id")
		assert.Error(t, err)
	})

	t.Run("NonExistentVMUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		err := client.VMs.DeleteCdRom(ctx, nonExistentUUID, "cdrom-id")
		assert.Error(t, err)
	})

	t.Run("NonExistentVMUUIDAsync", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.VMs.DeleteCdRomAsync(nonExistentUUID, "cdrom-id")
		assert.Error(t, err)
	})

	t.Run("NoPanic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			_ = client.VMs.DeleteCdRom(ctx, "", "")
		})
		assert.NotPanics(t, func() {
			_, _ = client.VMs.DeleteCdRomAsync("", "")
		})
	})
}

// TestVMsDeleteCdRomIntegration tests CD-ROM deletion with real Nutanix API calls
func TestVMsDeleteCdRomIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	vms, err := client.VMs.List(ctx)
	require.NoError(t, err)

	var vmUUID, cdRomUUID string
	for _, vm := range vms {
		if vm.ExtId == nil || len(vm.CdRoms) == 0 {
			continue
		}
		for _, cd := range vm.CdRoms {
			if cd.ExtId != nil {
				vmUUID = *vm.ExtId
				cdRomUUID = *cd.ExtId
				break
			}
		}
		if vmUUID != "" {
			break
		}
	}

	if vmUUID == "" {
		t.Skip("No VM with a CD-ROM found for testing")
	}

	t.Logf("Testing DeleteCdRom on VM %s, CdRom %s", vmUUID, cdRomUUID)

	t.Run("DeleteCdRom", func(t *testing.T) {
		err := client.VMs.DeleteCdRom(ctx, vmUUID, cdRomUUID)
		assert.NoError(t, err)
	})
}

// TestVMsGenerateConsoleToken_NilClient tests nil client error handling for console token methods
func TestVMsGenerateConsoleToken_NilClient(t *testing.T) {
	service := NewVMsService(nil)
	require.NotNil(t, service)
	ctx := context.Background()

	t.Run("GenerateConsoleToken", func(t *testing.T) {
		_, err := service.GenerateConsoleToken(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("GenerateConsoleTokenAsync", func(t *testing.T) {
		_, err := service.GenerateConsoleTokenAsync("test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})
}

// TestVMsGenerateConsoleTokenErrorScenarios tests console token generation with invalid and non-existent UUIDs
func TestVMsGenerateConsoleTokenErrorScenarios(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("InvalidUUID", func(t *testing.T) {
		_, err := client.VMs.GenerateConsoleToken(ctx, "invalid-uuid")
		assert.Error(t, err)
	})

	t.Run("InvalidUUIDAsync", func(t *testing.T) {
		_, err := client.VMs.GenerateConsoleTokenAsync("invalid-uuid")
		assert.Error(t, err)
	})

	t.Run("NonExistentUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.VMs.GenerateConsoleToken(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("NonExistentUUIDAsync", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.VMs.GenerateConsoleTokenAsync(nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("NoPanic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			_, _ = client.VMs.GenerateConsoleToken(ctx, "")
		})
		assert.NotPanics(t, func() {
			_, _ = client.VMs.GenerateConsoleTokenAsync("")
		})
	})
}

// TestVMsGenerateConsoleTokenIntegration tests console token generation with real Nutanix API calls
func TestVMsGenerateConsoleTokenIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	vms, err := client.VMs.List(ctx)
	require.NoError(t, err)

	var vmUUID string
	for _, vm := range vms {
		if vm.ExtId != nil && vm.PowerState != nil && *vm.PowerState == vmmModels.POWERSTATE_ON {
			vmUUID = *vm.ExtId
			break
		}
	}

	if vmUUID == "" {
		t.Skip("No powered-on VM found for testing")
	}

	t.Logf("Testing GenerateConsoleToken on VM %s", vmUUID)

	t.Run("GenerateConsoleToken", func(t *testing.T) {
		token, err := client.VMs.GenerateConsoleToken(ctx, vmUUID)
		assert.NoError(t, err)
		assert.NotNil(t, token)
		assert.NotEmpty(t, token.Token, "Token should not be empty")
		assert.NotEmpty(t, token.WsUri, "WsUri should not be empty")
		t.Logf("Console token generated, WsUri=%s", token.WsUri)
	})

	t.Run("GenerateConsoleTokenAsync", func(t *testing.T) {
		operation, err := client.VMs.GenerateConsoleTokenAsync(vmUUID)
		assert.NoError(t, err)
		assert.NotNil(t, operation)
		assert.NotEmpty(t, operation.UUID(), "operation UUID should not be empty")

		_, err = operation.Wait(ctx)
		assert.NoError(t, err)
		assert.True(t, operation.IsDone(), "operation should be done after Wait")
		assert.True(t, operation.IsSuccess(), "operation should be successful")
	})
}

// TestVMsListNicsByVmId_NilClient tests nil client error handling for ListNicsByVmId
func TestVMsListNicsByVmId_NilClient(t *testing.T) {
	service := NewVMsService(nil)
	require.NotNil(t, service)
	ctx := context.Background()

	t.Run("ListNicsByVmId", func(t *testing.T) {
		_, err := service.ListNicsByVmId(ctx, "test-vm-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})
}

// TestVMsListNicsByVmId_ClientConvenienceMethod tests the Client.ListNicsByVmId convenience method
func TestVMsListNicsByVmId_ClientConvenienceMethod(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("DelegatesToVMsService", func(t *testing.T) {
		_, err := client.ListNicsByVmId(ctx, "00000000-0000-0000-0000-000000000000")
		assert.Error(t, err)
	})
}

// TestVMsListNicsByVmIdErrorScenarios tests ListNicsByVmId with invalid and non-existent UUIDs
func TestVMsListNicsByVmIdErrorScenarios(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("InvalidVMUUID", func(t *testing.T) {
		_, err := client.ListNicsByVmId(ctx, "invalid-uuid")
		assert.Error(t, err)
	})

	t.Run("NonExistentVMUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.ListNicsByVmId(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("NoPanic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			_, _ = client.ListNicsByVmId(ctx, "")
		})
	})
}

// TestVMsListNicsByVmIdIntegration tests ListNicsByVmId with real Nutanix API calls
func TestVMsListNicsByVmIdIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	vms, err := client.VMs.List(ctx)
	require.NoError(t, err)

	var vmUUID string
	for _, vm := range vms {
		if vm.ExtId != nil && len(vm.Nics) > 0 {
			vmUUID = *vm.ExtId
			break
		}
	}

	if vmUUID == "" {
		t.Skip("No VM with NICs found for testing")
	}

	t.Logf("Testing ListNicsByVmId on VM %s", vmUUID)

	t.Run("ListNicsByVmId_VMsService", func(t *testing.T) {
		vmsService, ok := client.VMs.(*VMsService)
		require.True(t, ok)

		nics, err := vmsService.ListNicsByVmId(ctx, vmUUID)
		if err != nil && strings.Contains(err.Error(), "not authorized") {
			t.Skipf("Skipping: user not authorized for ListNicsByVmId on VM %s", vmUUID)
		}
		assert.NoError(t, err)
		assert.NotEmpty(t, nics, "VM should have at least one NIC")
	})

	t.Run("ListNicsByVmId_Client", func(t *testing.T) {
		nics, err := client.ListNicsByVmId(ctx, vmUUID)
		if err != nil && strings.Contains(err.Error(), "not authorized") {
			t.Skipf("Skipping: user not authorized for ListNicsByVmId on VM %s", vmUUID)
		}
		assert.NoError(t, err)
		assert.NotEmpty(t, nics, "VM should have at least one NIC")
	})
}
