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
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	vmmConfig "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// newVMServiceAgainstServer builds a VMsService whose VM API client is pointed at
// the given test HTTP server, so CreateVm requests are captured locally instead of
// hitting a real Prism Central.
func newVMServiceAgainstServer(t *testing.T, server *httptest.Server) *VMsService {
	t.Helper()

	u, err := url.Parse(server.URL)
	require.NoError(t, err)
	host, portStr, err := net.SplitHostPort(u.Host)
	require.NoError(t, err)
	port, err := strconv.Atoi(portStr)
	require.NoError(t, err)

	v4c, err := v4prismGoClient.NewV4Client(prismgoclient.Credentials{
		Endpoint: u.Host,
		Username: "admin",
		Password: "password",
		Insecure: true,
	})
	require.NoError(t, err)

	api := v4c.VmApiInstance.ApiClient
	api.Scheme = "http"
	api.Host = host
	api.Port = port
	// Avoid a pre-flight version negotiation round-trip against the test server.
	api.AllowVersionNegotiation = false

	return NewVMsService(v4c)
}

// createVmResponseBody returns a marshalled CreateVmApiResponse carrying a
// TaskReference, matching what Prism Central returns for an accepted create.
func createVmResponseBody(t *testing.T, taskExtID string) []byte {
	t.Helper()
	resp := &vmmModels.CreateVmApiResponse{}
	require.NoError(t, resp.SetData(vmmConfig.TaskReference{
		ObjectType_: ptr.To("prism.v4.config.TaskReference"),
		ExtId:       ptr.To(taskExtID),
	}))
	body, err := json.Marshal(resp)
	require.NoError(t, err)
	return body
}

func getVmResponseBodyWithEtag(t *testing.T, vmExtID string, etag string) []byte {
	t.Helper()
	resp := &vmmModels.GetVmApiResponse{}
	vm := vmmModels.Vm{
		ObjectType_: ptr.To("vmm.v4.ahv.config.Vm"),
		ExtId:       ptr.To(vmExtID),
	}
	resp.Reserved_ = map[string]any{"Etag": etag}
	require.NoError(t, resp.SetData(vm))
	body, err := json.Marshal(resp)
	require.NoError(t, err)
	return body
}

func taskRefBody(taskExtID string) []byte {
	return fmt.Appendf(nil, `{"data":{"$objectType":"prism.v4.config.TaskReference","extId":"%s"}}`, taskExtID)
}

func taskRefBodyWithoutExtID() []byte {
	return []byte(`{"data":{"$objectType":"prism.v4.config.TaskReference"}}`)
}

// TestCreateAsync_HonoursRequestIDHeader verifies that a request id set on the
// context via WithRequestID is sent as the Ntnx-Request-Id header on the CreateVm
// request, and that exactly one such header value reaches the wire (i.e. the SDK
// does not also inject its own random one).
func TestCreateAsync_HonoursRequestIDHeader(t *testing.T) {
	const requestID = "my-idempotency-key"
	body := createVmResponseBody(t, "task-ext-id")

	var captured http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.Header.Clone()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body)
	}))
	defer server.Close()

	service := newVMServiceAgainstServer(t, server)

	ctx := WithRequestID(context.Background(), requestID)
	op, err := service.CreateAsync(ctx, &vmmModels.Vm{Name: ptr.To("test-vm")})
	require.NoError(t, err)
	require.NotNil(t, op)

	// HTTP header names are case-insensitive; the server canonicalises them, so
	// both "NTNX-Request-Id" (what we send) and the SDK's auto-injected variant
	// would collapse onto the same canonical key. Assert exactly one value, equal
	// to ours.
	values := captured.Values("Ntnx-Request-Id")
	require.Len(t, values, 1, "expected exactly one request-id header, got %v", values)
	assert.Equal(t, requestID, values[0])
}

// TestCreateAsync_NoRequestID_NoDuplicateHeader verifies current behaviour.
func TestCreateAsync_NoRequestID_NoDuplicateHeader(t *testing.T) {
	body := createVmResponseBody(t, "task-ext-id")

	var captured http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.Header.Clone()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body)
	}))
	defer server.Close()

	service := newVMServiceAgainstServer(t, server)

	op, err := service.CreateAsync(context.Background(), &vmmModels.Vm{Name: ptr.To("test-vm")})
	require.NoError(t, err)
	require.NotNil(t, op)

	values := captured.Values("Ntnx-Request-Id")
	assert.Len(t, values, 1, "expected exactly one request-id header, got %v", values)
	assert.NotEmpty(t, values[0])
}

// TestVMSubresourceAsyncOps_HonourRequestIDAndETag verifies all new day-2 VM
// subresource operations fetch VM ETag first and pass both If-Match + request-id
// on the mutation request.
func TestVMSubresourceAsyncOps_HonourRequestIDAndETag(t *testing.T) {
	tests := []struct {
		name       string
		expectedOp string
		call       func(t *testing.T, service *VMsService, ctx context.Context) error
	}{
		{
			name:       "AddDisk",
			expectedOp: "/api/vmm/v4.3/ahv/config/vms/vm-1/disks",
			call: func(t *testing.T, service *VMsService, ctx context.Context) error {
				_, err := service.AddDisk(ctx, "vm-1", &vmmModels.Disk{})
				return err
			},
		},
		{
			name:       "GrowDisk",
			expectedOp: "/api/vmm/v4.3/ahv/config/vms/vm-1/disks/disk-1",
			call: func(t *testing.T, service *VMsService, ctx context.Context) error {
				_, err := service.GrowDisk(ctx, "vm-1", "disk-1", &vmmModels.Disk{})
				return err
			},
		},
		{
			name:       "DeleteDisk",
			expectedOp: "/api/vmm/v4.3/ahv/config/vms/vm-1/disks/disk-1",
			call: func(t *testing.T, service *VMsService, ctx context.Context) error {
				_, err := service.DeleteDisk(ctx, "vm-1", "disk-1")
				return err
			},
		},
		{
			name:       "AddNIC",
			expectedOp: "/api/vmm/v4.3/ahv/config/vms/vm-1/nics",
			call: func(t *testing.T, service *VMsService, ctx context.Context) error {
				_, err := service.AddNIC(ctx, "vm-1", &vmmModels.Nic{})
				return err
			},
		},
		{
			name:       "DeleteNIC",
			expectedOp: "/api/vmm/v4.3/ahv/config/vms/vm-1/nics/nic-1",
			call: func(t *testing.T, service *VMsService, ctx context.Context) error {
				_, err := service.DeleteNIC(ctx, "vm-1", "nic-1")
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				mu       sync.Mutex
				requests []*http.Request
			)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mu.Lock()
				requests = append(requests, r.Clone(r.Context()))
				callNum := len(requests)
				mu.Unlock()

				w.Header().Set("Content-Type", "application/json")
				if callNum == 1 {
					_, _ = w.Write(getVmResponseBodyWithEtag(t, "vm-1", "etag-1"))
					return
				}
				_, _ = w.Write(taskRefBody("task-ext-id"))
			}))
			defer server.Close()

			service := newVMServiceAgainstServer(t, server)
			ctx := WithRequestID(context.Background(), "req-id-1")
			err := tt.call(t, service, ctx)
			require.NoError(t, err)

			mu.Lock()
			require.Len(t, requests, 2, "expected GET vm + one mutation call")
			getReq := requests[0]
			opReq := requests[1]
			mu.Unlock()

			assert.Equal(t, http.MethodGet, getReq.Method)
			assert.Contains(t, getReq.URL.Path, "/api/vmm/v4.3/ahv/config/vms/vm-1")

			assert.Equal(t, tt.expectedOp, opReq.URL.Path)
			assert.Equal(t, "etag-1", opReq.Header.Get("If-Match"))
			assert.Equal(t, "req-id-1", opReq.Header.Get("Ntnx-Request-Id"))
		})
	}
}

func TestVMSubresourceAsyncOps_NilClient(t *testing.T) {
	service := NewVMsService(nil)
	ctx := context.Background()

	_, err := service.AddDisk(ctx, "vm", &vmmModels.Disk{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")

	_, err = service.GrowDisk(ctx, "vm", "disk", &vmmModels.Disk{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")

	_, err = service.DeleteDisk(ctx, "vm", "disk")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")

	_, err = service.AddNIC(ctx, "vm", &vmmModels.Nic{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")

	_, err = service.DeleteNIC(ctx, "vm", "nic")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")
}

func TestVMSubresourceAsyncOps_NilPayload(t *testing.T) {
	service := NewVMsService(&v4prismGoClient.Client{})
	ctx := context.Background()

	_, err := service.AddDisk(ctx, "vm", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk payload must be *vmmModels.Disk")

	_, err = service.GrowDisk(ctx, "vm", "disk", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk payload must be *vmmModels.Disk")

	_, err = service.AddNIC(ctx, "vm", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "nic payload must be *vmmModels.Nic")
}

func TestVMSubresourceAsyncOps_ETagFetchFailure(t *testing.T) {
	tests := []struct {
		name         string
		call         func(service *VMsService) error
		expectedText string
	}{
		{
			name: "AddDisk",
			call: func(service *VMsService) error {
				_, err := service.AddDisk(context.Background(), "vm-1", &vmmModels.Disk{})
				return err
			},
			expectedText: "failed to get VM details",
		},
		{
			name: "GrowDisk",
			call: func(service *VMsService) error {
				_, err := service.GrowDisk(context.Background(), "vm-1", "disk-1", &vmmModels.Disk{})
				return err
			},
			expectedText: "failed to get VM details",
		},
		{
			name: "AddNIC",
			call: func(service *VMsService) error {
				_, err := service.AddNIC(context.Background(), "vm-1", &vmmModels.Nic{})
				return err
			},
			expectedText: "failed to get VM details",
		},
		{
			name: "DeleteNIC",
			call: func(service *VMsService) error {
				_, err := service.DeleteNIC(context.Background(), "vm-1", "nic-1")
				return err
			},
			expectedText: "failed to get VM for NIC delete",
		},
		{
			name: "DeleteDisk",
			call: func(service *VMsService) error {
				_, err := service.DeleteDisk(context.Background(), "vm-1", "disk-1")
				return err
			},
			expectedText: "failed to get VM details",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Return VM without ETag so GetEntityAndEtag fails.
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write(getVmResponseBodyWithEtag(t, "vm-1", ""))
			}))
			defer server.Close()

			service := newVMServiceAgainstServer(t, server)
			err := tt.call(service)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedText)
			assert.Contains(t, err.Error(), "no ETag found")
		})
	}
}

func TestVMSubresourceAsyncOps_MutationFailure(t *testing.T) {
	tests := []struct {
		name         string
		call         func(service *VMsService) error
		expectedText string
	}{
		{
			name: "AddDisk",
			call: func(service *VMsService) error {
				_, err := service.AddDisk(context.Background(), "vm-1", &vmmModels.Disk{})
				return err
			},
			expectedText: "failed to add disk to VM",
		},
		{
			name: "GrowDisk",
			call: func(service *VMsService) error {
				_, err := service.GrowDisk(context.Background(), "vm-1", "disk-1", &vmmModels.Disk{})
				return err
			},
			expectedText: "failed to grow disk size for VM",
		},
		{
			name: "AddNIC",
			call: func(service *VMsService) error {
				_, err := service.AddNIC(context.Background(), "vm-1", &vmmModels.Nic{})
				return err
			},
			expectedText: "failed to add NIC to VM",
		},
		{
			name: "DeleteNIC",
			call: func(service *VMsService) error {
				_, err := service.DeleteNIC(context.Background(), "vm-1", "nic-1")
				return err
			},
			expectedText: "failed to delete NIC from VM",
		},
		{
			name: "DeleteDisk",
			call: func(service *VMsService) error {
				_, err := service.DeleteDisk(context.Background(), "vm-1", "disk-1")
				return err
			},
			expectedText: "failed to delete disk from VM",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callNum := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				callNum++
				w.Header().Set("Content-Type", "application/json")
				if callNum == 1 {
					_, _ = w.Write(getVmResponseBodyWithEtag(t, "vm-1", "etag-1"))
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"errors":[{"message":"boom"}]}`))
			}))
			defer server.Close()

			service := newVMServiceAgainstServer(t, server)
			err := tt.call(service)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedText)
		})
	}
}

func TestVMSubresourceAsyncOps_TaskReferenceExtIDNil(t *testing.T) {
	tests := []struct {
		name         string
		call         func(service *VMsService) error
		expectedText string
	}{
		{
			name: "AddDisk",
			call: func(service *VMsService) error {
				_, err := service.AddDisk(context.Background(), "vm-1", &vmmModels.Disk{})
				return err
			},
			expectedText: "task reference ExtId is nil for VM disk add",
		},
		{
			name: "GrowDisk",
			call: func(service *VMsService) error {
				_, err := service.GrowDisk(context.Background(), "vm-1", "disk-1", &vmmModels.Disk{})
				return err
			},
			expectedText: "task reference ExtId is nil for VM disk grow",
		},
		{
			name: "AddNIC",
			call: func(service *VMsService) error {
				_, err := service.AddNIC(context.Background(), "vm-1", &vmmModels.Nic{})
				return err
			},
			expectedText: "task reference ExtId is nil for VM NIC add",
		},
		{
			name: "DeleteNIC",
			call: func(service *VMsService) error {
				_, err := service.DeleteNIC(context.Background(), "vm-1", "nic-1")
				return err
			},
			expectedText: "task reference ExtId is nil for VM NIC delete",
		},
		{
			name: "DeleteDisk",
			call: func(service *VMsService) error {
				_, err := service.DeleteDisk(context.Background(), "vm-1", "disk-1")
				return err
			},
			expectedText: "task reference ExtId is nil for VM disk delete",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callNum := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				callNum++
				w.Header().Set("Content-Type", "application/json")
				if callNum == 1 {
					_, _ = w.Write(getVmResponseBodyWithEtag(t, "vm-1", "etag-1"))
					return
				}
				_, _ = w.Write(taskRefBodyWithoutExtID())
			}))
			defer server.Close()

			service := newVMServiceAgainstServer(t, server)
			err := tt.call(service)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedText)
		})
	}
}

func TestVMSubresourceAsyncOps_NoRequestIDHeaderFallback(t *testing.T) {
	tests := []struct {
		name string
		call func(service *VMsService) error
	}{
		{
			name: "AddDisk",
			call: func(service *VMsService) error {
				_, err := service.AddDisk(context.Background(), "vm-1", &vmmModels.Disk{})
				return err
			},
		},
		{
			name: "GrowDisk",
			call: func(service *VMsService) error {
				_, err := service.GrowDisk(context.Background(), "vm-1", "disk-1", &vmmModels.Disk{})
				return err
			},
		},
		{
			name: "DeleteDisk",
			call: func(service *VMsService) error {
				_, err := service.DeleteDisk(context.Background(), "vm-1", "disk-1")
				return err
			},
		},
		{
			name: "AddNIC",
			call: func(service *VMsService) error {
				_, err := service.AddNIC(context.Background(), "vm-1", &vmmModels.Nic{})
				return err
			},
		},
		{
			name: "DeleteNIC",
			call: func(service *VMsService) error {
				_, err := service.DeleteNIC(context.Background(), "vm-1", "nic-1")
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				mu       sync.Mutex
				requests []*http.Request
			)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mu.Lock()
				requests = append(requests, r.Clone(r.Context()))
				callNum := len(requests)
				mu.Unlock()

				w.Header().Set("Content-Type", "application/json")
				if callNum == 1 {
					_, _ = w.Write(getVmResponseBodyWithEtag(t, "vm-1", "etag-1"))
					return
				}
				_, _ = w.Write(taskRefBody("task-ext-id"))
			}))
			defer server.Close()

			service := newVMServiceAgainstServer(t, server)
			err := tt.call(service)
			require.NoError(t, err)

			mu.Lock()
			require.Len(t, requests, 2)
			opReq := requests[1]
			mu.Unlock()

			values := opReq.Header.Values("Ntnx-Request-Id")
			require.Len(t, values, 1, "expected exactly one request-id header when caller does not set one")
			assert.NotEmpty(t, values[0], "sdk-generated request-id should be non-empty")
			assert.Equal(t, "etag-1", opReq.Header.Get("If-Match"))
		})
	}
}

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
		_, err := service.AddVmCustomAttributesAsync(ctx, "test-id", []string{"attr1"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("RemoveVmCustomAttributes", func(t *testing.T) {
		_, err := service.RemoveVmCustomAttributes(ctx, "test-id", []string{"attr1"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("RemoveVmCustomAttributesAsync", func(t *testing.T) {
		_, err := service.RemoveVmCustomAttributesAsync(ctx, "test-id", []string{"attr1"})
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
		operation, err := client.VMs.AddVmCustomAttributesAsync(ctx, vmUUID, []string{"test-attr-async"})
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
		operation, err := client.VMs.RemoveVmCustomAttributesAsync(ctx, vmUUID, []string{"test-attr-async"})
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

		_, err = client.VMs.AddVmCustomAttributesAsync(ctx, "invalid-uuid-format", []string{"attr1"})
		assert.Error(t, err)

		_, err = client.VMs.RemoveVmCustomAttributes(ctx, "invalid-uuid-format", []string{"attr1"})
		assert.Error(t, err)

		_, err = client.VMs.RemoveVmCustomAttributesAsync(ctx, "invalid-uuid-format", []string{"attr1"})
		assert.Error(t, err)
	})

	t.Run("NonExistentUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"

		_, err := client.VMs.AddVmCustomAttributes(ctx, nonExistentUUID, []string{"attr1"})
		assert.Error(t, err)

		_, err = client.VMs.AddVmCustomAttributesAsync(ctx, nonExistentUUID, []string{"attr1"})
		assert.Error(t, err)

		_, err = client.VMs.RemoveVmCustomAttributes(ctx, nonExistentUUID, []string{"attr1"})
		assert.Error(t, err)

		_, err = client.VMs.RemoveVmCustomAttributesAsync(ctx, nonExistentUUID, []string{"attr1"})
		assert.Error(t, err)
	})

	t.Run("EmptyAttributes", func(t *testing.T) {
		assert.NotPanics(t, func() {
			_, _ = client.VMs.AddVmCustomAttributesAsync(ctx, "invalid-uuid", []string{})
		})

		assert.NotPanics(t, func() {
			_, _ = client.VMs.RemoveVmCustomAttributesAsync(ctx, "invalid-uuid", []string{})
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
		_, err := service.DeleteCdRomAsync(ctx, "test-vm-id", "test-cdrom-id")
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
		_, err := client.VMs.DeleteCdRomAsync(ctx, "invalid-uuid", "cdrom-id")
		assert.Error(t, err)
	})

	t.Run("NonExistentVMUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		err := client.VMs.DeleteCdRom(ctx, nonExistentUUID, "cdrom-id")
		assert.Error(t, err)
	})

	t.Run("NonExistentVMUUIDAsync", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.VMs.DeleteCdRomAsync(ctx, nonExistentUUID, "cdrom-id")
		assert.Error(t, err)
	})

	t.Run("NoPanic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			_ = client.VMs.DeleteCdRom(ctx, "", "")
		})
		assert.NotPanics(t, func() {
			_, _ = client.VMs.DeleteCdRomAsync(ctx, "", "")
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
		_, err := service.GenerateConsoleTokenAsync(ctx, "test-id")
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
		_, err := client.VMs.GenerateConsoleTokenAsync(ctx, "invalid-uuid")
		assert.Error(t, err)
	})

	t.Run("NonExistentUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.VMs.GenerateConsoleToken(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("NonExistentUUIDAsync", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.VMs.GenerateConsoleTokenAsync(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("NoPanic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			_, _ = client.VMs.GenerateConsoleToken(ctx, "")
		})
		assert.NotPanics(t, func() {
			_, _ = client.VMs.GenerateConsoleTokenAsync(ctx, "")
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
		operation, err := client.VMs.GenerateConsoleTokenAsync(ctx, vmUUID)
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

// TestVMsListNicsByVmId_ClientVMsField tests accessing ListNicsByVmId through client.VMs
func TestVMsListNicsByVmId_ClientVMsField(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ViaVMsInterface", func(t *testing.T) {
		_, err := client.VMs.ListNicsByVmId(ctx, "00000000-0000-0000-0000-000000000000")
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
		_, err := client.VMs.ListNicsByVmId(ctx, "invalid-uuid")
		assert.Error(t, err)
	})

	t.Run("NonExistentVMUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.VMs.ListNicsByVmId(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("NoPanic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			_, _ = client.VMs.ListNicsByVmId(ctx, "")
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

	t.Run("ListNicsByVmId_ViaVMsInterface", func(t *testing.T) {
		nics, err := client.VMs.ListNicsByVmId(ctx, vmUUID)
		if err != nil && strings.Contains(err.Error(), "not authorized") {
			t.Skipf("Skipping: user not authorized for ListNicsByVmId on VM %s", vmUUID)
		}
		assert.NoError(t, err)
		assert.NotEmpty(t, nics, "VM should have at least one NIC")
	})
}

// TestVMsGetVMByBiosUUID_NilClient tests nil client error handling for GetVMByBiosUUID
func TestVMsGetVMByBiosUUID_NilClient(t *testing.T) {
	service := NewVMsService(nil)
	require.NotNil(t, service)
	ctx := context.Background()

	_, err := service.GetVMByBiosUUID(ctx, "00000000-0000-0000-0000-000000000000")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client is not initialized")
}

// TestResolveVMByBiosUUID exercises the pure leaf-resolution logic that
// GetVMByBiosUUID applies to the VMs returned by the biosUuid filter.
func TestResolveVMByBiosUUID(t *testing.T) {
	strPtr := func(s string) *string { return &s }
	biosUUID := "bios-uuid"

	vmWithSource := func(extID, sourceExtID string) vmmModels.Vm {
		vm := vmmModels.Vm{ExtId: strPtr(extID)}
		if sourceExtID != "" {
			vm.Source = &vmmModels.VmSourceReference{ExtId: strPtr(sourceExtID)}
		}
		return vm
	}

	tests := []struct {
		name        string
		vms         []vmmModels.Vm
		wantExtID   string
		wantErr     bool
		errContains string
	}{
		{
			name:        "no VMs returns error",
			vms:         nil,
			wantErr:     true,
			errContains: "no VM found with bios UUID",
		},
		{
			name:      "single VM returned directly",
			vms:       []vmmModels.Vm{vmWithSource("vm-a", "")},
			wantExtID: "vm-a",
		},
		{
			name: "clone chain resolves to leaf not referenced as a source",
			// vm-b was cloned from vm-a; vm-a is referenced as vm-b's source, so
			// the leaf vm-b is returned.
			vms:       []vmmModels.Vm{vmWithSource("vm-a", ""), vmWithSource("vm-b", "vm-a")},
			wantExtID: "vm-b",
		},
		{
			name: "multiple leaves cannot be resolved",
			// Two VMs share the bios UUID and neither is referenced as a source,
			// so the chain is ambiguous.
			vms:         []vmmModels.Vm{vmWithSource("vm-a", ""), vmWithSource("vm-b", "")},
			wantErr:     true,
			errContains: "could not resolve to a single VM",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm, err := resolveVMByBiosUUID(tt.vms, biosUUID)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, vm)
			require.NotNil(t, vm.ExtId)
			assert.Equal(t, tt.wantExtID, *vm.ExtId)
		})
	}
}
