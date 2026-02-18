// Package v4 provides tests for the converged client VMs service custom attributes methods.
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run TestVMsCustomAttributesIntegration
//	go test -v ./converged/v4 -run TestVMsCustomAttributes_NilClient
package v4

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
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
		// Test with empty attributes - should not panic
		assert.NotPanics(t, func() {
			_, _ = client.VMs.AddVmCustomAttributesAsync("invalid-uuid", []string{})
		})

		assert.NotPanics(t, func() {
			_, _ = client.VMs.RemoveVmCustomAttributesAsync("invalid-uuid", []string{})
		})
	})
}
