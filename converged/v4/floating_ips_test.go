// Package v4 provides integration tests for the converged client floating IPs service.
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// Optional fixtures (override discovery when set):
//   - NUTANIX_EXTERNAL_SUBNET_NAME: External subnet used to allocate floating IPs.
//     Creates use that subnet's projectExtId.
//   - NUTANIX_PRIVATE_SUBNET_NAME: Overlay subnet in a VPC for PrivateIp assign/deassign.
//     The floating IP is created in this subnet's project, which must own the VPC.
//     The external subnet must be in that project or shared with it.
//   - NUTANIX_PRIVATE_IP: Private IPv4 used with that subnet's VPC for PrivateIp association
//
// To run these tests:
//
//	go test -v ./converged/v4 -run TestFloatingIpsIntegration
//	go test -v ./converged/v4 -run TestFloatingIpsCRUDIntegration
//	go test -v ./converged/v4 -run TestFloatingIpsAssignDeassignIntegration
//	go test -v ./converged/v4 -run TestFloatingIpsErrorScenarios
//	go test -v ./converged/v4 -run TestFloatingIpsService_ErrorHandling
//	go test -v ./converged/v4 -run TestFloatingIpsService_UnsupportedODataOptions
//	go test -v ./converged/v4 -run TestFloatingIpsService_AssignAssociationValidation
package v4

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	commonConfig "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/config"
	subnetModels "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

const (
	envExternalSubnetName = "NUTANIX_EXTERNAL_SUBNET_NAME"
	envPrivateSubnetName  = "NUTANIX_PRIVATE_SUBNET_NAME"
	envPrivateIP          = "NUTANIX_PRIVATE_IP"
	nonExistentUUID       = "00000000-0000-0000-0000-000000000000"
)

func odataStringLiteral(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "''") + "'"
}

// subnetByName looks up a subnet by exact name. Skips if none is found.
func subnetByName(t *testing.T, client *Client, name string) subnetModels.Subnet {
	t.Helper()

	ctx := context.Background()
	subnets, err := client.Subnets.List(ctx,
		converged.WithFilter(fmt.Sprintf("name eq %s", odataStringLiteral(name))),
		converged.WithLimit(10),
	)
	require.NoError(t, err)

	for _, subnet := range subnets {
		if subnet.Name != nil && *subnet.Name == name && subnet.ExtId != nil && *subnet.ExtId != "" {
			return subnet
		}
	}

	t.Skipf("No subnet named %q found", name)
	return subnetModels.Subnet{}
}

// externalSubnet returns an external subnet for floating IP allocation.
// Prefers NUTANIX_EXTERNAL_SUBNET_NAME; otherwise discovers the first external subnet.
func externalSubnet(t *testing.T, client *Client) subnetModels.Subnet {
	t.Helper()

	if name := strings.TrimSpace(os.Getenv(envExternalSubnetName)); name != "" {
		subnet := subnetByName(t, client, name)
		if subnet.IsExternal == nil || !*subnet.IsExternal {
			t.Skipf("subnet %q is not external", name)
		}
		return subnet
	}

	ctx := context.Background()
	subnets, err := client.Subnets.List(ctx, converged.WithLimit(100))
	require.NoError(t, err)

	for _, subnet := range subnets {
		if subnet.ExtId == nil {
			continue
		}
		if subnet.IsExternal != nil && *subnet.IsExternal {
			return subnet
		}
	}

	t.Skip("No external subnet available for floating IP testing (set NUTANIX_EXTERNAL_SUBNET_NAME)")
	return subnetModels.Subnet{}
}

func subnetProjectExtID(subnet subnetModels.Subnet) string {
	if subnet.ProjectExtId == nil {
		return ""
	}
	return *subnet.ProjectExtId
}

// subnetVisibleToProject reports whether subnet is owned by projectExtID or shared with it.
func subnetVisibleToProject(subnet subnetModels.Subnet, projectExtID string) bool {
	if projectExtID == "" || (subnet.ProjectExtId != nil && *subnet.ProjectExtId == projectExtID) {
		return true
	}
	for _, shared := range subnet.SharedWithProjects {
		if shared == projectExtID {
			return true
		}
	}
	return false
}

type privateIPFixture struct {
	association  subnetModels.PrivateIpAssociation
	projectExtID string
	subnetName   string
}

// privateIPFixtureFromEnv returns a PrivateIpAssociation and the project that owns
// the overlay subnet. The floating IP must be created in that project, which is
// the VPC's project. The default project cannot reference a VPC that lives elsewhere.
func privateIPFixtureFromEnv(t *testing.T, client *Client) privateIPFixture {
	t.Helper()

	subnetName := strings.TrimSpace(os.Getenv(envPrivateSubnetName))
	privateIP := strings.TrimSpace(os.Getenv(envPrivateIP))
	if subnetName == "" || privateIP == "" {
		t.Skip("No private IP fixtures available (set NUTANIX_PRIVATE_SUBNET_NAME and NUTANIX_PRIVATE_IP)")
	}

	subnet := subnetByName(t, client, subnetName)
	if subnet.VpcReference == nil || *subnet.VpcReference == "" {
		t.Skipf("subnet %q is not in a VPC (VpcReference is empty)", subnetName)
	}
	projectExtID := subnetProjectExtID(subnet)
	if projectExtID == "" {
		t.Skipf("subnet %q has no projectExtId", subnetName)
	}

	ip := commonConfig.NewIPAddress()
	ip.Ipv4 = commonConfig.NewIPv4Address()
	ip.Ipv4.Value = ptr.To(privateIP)

	association := subnetModels.NewPrivateIpAssociation()
	association.PrivateIp = ip
	association.VpcReference = subnet.VpcReference
	return privateIPFixture{
		association:  *association,
		projectExtID: projectExtID,
		subnetName:   subnetName,
	}
}

func runAssignDeassign(t *testing.T, client *Client, subnetExtID, projectExtID string, association any, wantStatus subnetModels.AssociationStatus) {
	t.Helper()

	ctx := context.Background()
	created, err := client.FloatingIps.Create(ctx, newTestFloatingIP(uniqueFloatingIPName(), subnetExtID, projectExtID))
	require.NoError(t, err)
	require.NotNil(t, created)
	require.NotNil(t, created.ExtId)
	extID := *created.ExtId
	registerFloatingIPCleanup(t, client, extID)

	t.Run("Assign", func(t *testing.T) {
		assigned, err := client.FloatingIps.AssignFloatingIp(ctx, extID, association)
		if err != nil {
			t.Skipf("Skipping assign: environment rejected association (%v)", err)
		}
		require.NotNil(t, assigned)

		got, err := client.FloatingIps.Get(ctx, extID)
		require.NoError(t, err)
		require.NotNil(t, got)
		if got.AssociationStatus != nil {
			assert.Equal(t, wantStatus, *got.AssociationStatus)
		}
		if got.GetAssociation() == nil {
			t.Log("association fields not populated on GET; assign succeeded without error")
		}
	})

	t.Run("Deassign", func(t *testing.T) {
		gotBefore, err := client.FloatingIps.Get(ctx, extID)
		require.NoError(t, err)
		if gotBefore.AssociationStatus != nil &&
			*gotBefore.AssociationStatus == subnetModels.ASSOCIATIONSTATUS_NOT_ASSOCIATED {
			_, assignErr := client.FloatingIps.AssignFloatingIp(ctx, extID, association)
			if assignErr != nil {
				t.Skipf("Skipping deassign: could not associate first (%v)", assignErr)
			}
		}

		deassigned, err := client.FloatingIps.DeassignFloatingIp(ctx, extID)
		if err != nil {
			t.Skipf("Skipping deassign: environment rejected deassociation (%v)", err)
		}
		require.NotNil(t, deassigned)

		got, err := client.FloatingIps.Get(ctx, extID)
		require.NoError(t, err)
		require.NotNil(t, got)
		if got.AssociationStatus != nil {
			assert.Equal(t, subnetModels.ASSOCIATIONSTATUS_NOT_ASSOCIATED, *got.AssociationStatus)
		}
		assert.Nil(t, got.GetAssociation())
	})
}

func newTestFloatingIP(name, externalSubnetExtID, projectExtID string) *subnetModels.FloatingIp {
	floatingIP := subnetModels.NewFloatingIp()
	floatingIP.Name = ptr.To(name)
	floatingIP.ExternalSubnetReference = ptr.To(externalSubnetExtID)
	if projectExtID != "" {
		floatingIP.ProjectExtId = ptr.To(projectExtID)
	}
	return floatingIP
}

func uniqueFloatingIPName() string {
	return fmt.Sprintf("prism-go-client-fip-%s", uuid.NewString()[:8])
}

// registerFloatingIPCleanup deletes the floating IP on test cleanup unless cancelled.
func registerFloatingIPCleanup(t *testing.T, client *Client, extID string) (cancel func()) {
	t.Helper()
	var mu sync.Mutex
	id := extID
	t.Cleanup(func() {
		mu.Lock()
		toDelete := id
		mu.Unlock()
		if toDelete == "" {
			return
		}
		ctx := context.Background()
		if err := client.FloatingIps.Delete(ctx, toDelete); err != nil {
			t.Logf("cleanup: failed to delete floating IP %s: %v", toDelete, err)
		}
	})
	return func() {
		mu.Lock()
		defer mu.Unlock()
		id = ""
	}
}

// TestFloatingIpsIntegration tests client.FloatingIps with real Nutanix API calls.
func TestFloatingIpsIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ListFloatingIps", func(t *testing.T) {
		floatingIPs, err := client.FloatingIps.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, floatingIPs)
		assert.GreaterOrEqual(t, len(floatingIPs), 0)
	})

	t.Run("ListFloatingIpsWithValidOptions", func(t *testing.T) {
		floatingIPs, err := client.FloatingIps.List(ctx,
			converged.WithPage(0),
			converged.WithLimit(5),
			converged.WithOrderBy("name asc"),
		)
		assert.NoError(t, err)
		assert.NotNil(t, floatingIPs)
	})

	t.Run("GetFloatingIp", func(t *testing.T) {
		floatingIPs, err := client.FloatingIps.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(floatingIPs) == 0 {
			t.Skip("No floating IPs available for testing")
		}

		uuid := *floatingIPs[0].ExtId
		require.NotEmpty(t, uuid)

		floatingIP, err := client.FloatingIps.Get(ctx, uuid)
		assert.NoError(t, err)
		assert.NotNil(t, floatingIP)
		assert.Equal(t, uuid, *floatingIP.ExtId)
	})

	t.Run("NewIterator", func(t *testing.T) {
		iter := client.FloatingIps.NewIterator(ctx, converged.WithLimit(5))
		require.NotNil(t, iter)

		var count int
		for _, err := range iter {
			if err != nil {
				t.Logf("Iterator error: %v", err)
				break
			}
			count++
		}
		assert.GreaterOrEqual(t, count, 0)
	})
}

// TestFloatingIpsCRUDIntegration exercises create/get/update/delete against a real PC.
func TestFloatingIpsCRUDIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()
	subnet := externalSubnet(t, client)
	require.NotNil(t, subnet.ExtId)
	subnetExtID := *subnet.ExtId
	projectExtID := subnetProjectExtID(subnet)

	t.Run("CreateGetUpdateDelete", func(t *testing.T) {
		name := uniqueFloatingIPName()
		created, err := client.FloatingIps.Create(ctx, newTestFloatingIP(name, subnetExtID, projectExtID))
		require.NoError(t, err)
		require.NotNil(t, created)
		require.NotNil(t, created.ExtId)
		require.NotEmpty(t, *created.ExtId)

		extID := *created.ExtId
		cancelCleanup := registerFloatingIPCleanup(t, client, extID)

		got, err := client.FloatingIps.Get(ctx, extID)
		require.NoError(t, err)
		require.NotNil(t, got)
		assert.Equal(t, extID, *got.ExtId)
		if got.Name != nil {
			assert.Equal(t, name, *got.Name)
		}
		if got.ExternalSubnetReference != nil {
			assert.Equal(t, subnetExtID, *got.ExternalSubnetReference)
		}
		if projectExtID != "" {
			require.NotNil(t, got.ProjectExtId)
			assert.Equal(t, projectExtID, *got.ProjectExtId)
		}

		updatedDesc := "updated by prism-go-client floating IP e2e"
		updateBody, err := floatingIpUpdateBody(got)
		require.NoError(t, err)
		require.NotNil(t, updateBody)
		updateBody.Description = ptr.To(updatedDesc)

		updated, err := client.FloatingIps.Update(ctx, extID, updateBody)
		require.NoError(t, err)
		require.NotNil(t, updated)

		gotAfterUpdate, err := client.FloatingIps.Get(ctx, extID)
		require.NoError(t, err)
		require.NotNil(t, gotAfterUpdate)
		require.NotNil(t, gotAfterUpdate.Description)
		assert.Equal(t, updatedDesc, *gotAfterUpdate.Description)

		require.NoError(t, client.FloatingIps.Delete(ctx, extID))
		cancelCleanup()

		_, err = client.FloatingIps.Get(ctx, extID)
		assert.Error(t, err)
	})

	t.Run("CreateAsync", func(t *testing.T) {
		name := uniqueFloatingIPName()
		op, err := client.FloatingIps.CreateAsync(ctx, newTestFloatingIP(name, subnetExtID, projectExtID))
		require.NoError(t, err)
		require.NotNil(t, op)

		results, err := op.Wait(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, results)
		require.NotNil(t, results[0])
		require.NotNil(t, results[0].ExtId)

		extID := *results[0].ExtId
		cancelCleanup := registerFloatingIPCleanup(t, client, extID)

		got, err := client.FloatingIps.Get(ctx, extID)
		require.NoError(t, err)
		require.NotNil(t, got)
		assert.Equal(t, extID, *got.ExtId)

		require.NoError(t, client.FloatingIps.Delete(ctx, extID))
		cancelCleanup()
	})
}

// TestFloatingIpsAssignDeassignIntegration exercises PrivateIp assign/deassign.
// Skips when private-subnet or private-IP fixtures are missing.
func TestFloatingIpsAssignDeassignIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	extSubnet := externalSubnet(t, client)
	require.NotNil(t, extSubnet.ExtId)

	fixture := privateIPFixtureFromEnv(t, client)
	if !subnetVisibleToProject(extSubnet, fixture.projectExtID) {
		name := ""
		if extSubnet.Name != nil {
			name = *extSubnet.Name
		}
		t.Skipf("external subnet %q is not owned by or shared with project %s (private subnet %s)", name, fixture.projectExtID, fixture.subnetName)
	}

	runAssignDeassign(t, client, *extSubnet.ExtId, fixture.projectExtID, fixture.association, subnetModels.ASSOCIATIONSTATUS_ASSOCIATED_TO_PRIVATE_IP)
}

// TestFloatingIpsErrorScenarios tests negative API paths against a real PC.
func TestFloatingIpsErrorScenarios(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("GetInvalidUUID", func(t *testing.T) {
		_, err := client.FloatingIps.Get(ctx, "invalid-uuid-format")
		assert.Error(t, err)
	})

	t.Run("GetNonExistentUUID", func(t *testing.T) {
		_, err := client.FloatingIps.Get(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("DeleteNonExistentUUID", func(t *testing.T) {
		err := client.FloatingIps.Delete(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("UpdateNonExistentUUID", func(t *testing.T) {
		_, err := client.FloatingIps.Update(ctx, nonExistentUUID, subnetModels.NewFloatingIp())
		assert.Error(t, err)
	})

	t.Run("CreateMissingExternalSubnet", func(t *testing.T) {
		floatingIP := subnetModels.NewFloatingIp()
		floatingIP.Name = ptr.To(uniqueFloatingIPName())
		_, err := client.FloatingIps.Create(ctx, floatingIP)
		assert.Error(t, err)
	})

	t.Run("AssignInvalidAssociationType", func(t *testing.T) {
		_, err := client.FloatingIps.AssignFloatingIp(ctx, nonExistentUUID, "not-an-association")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "association must be VmNicAssociation, PrivateIpAssociation, or LoadBalancerSessionAssociation")
	})

	t.Run("AssignNonExistentFloatingIp", func(t *testing.T) {
		_, err := client.FloatingIps.AssignFloatingIp(ctx, nonExistentUUID, subnetModels.VmNicAssociation{
			VmNicReference: ptr.To(nonExistentUUID),
		})
		assert.Error(t, err)
	})
}

// TestFloatingIpsService_ErrorHandling tests error handling for nil client.
func TestFloatingIpsService_ErrorHandling(t *testing.T) {
	service := NewFloatingIpsService(nil)
	require.NotNil(t, service)
	ctx := context.Background()

	t.Run("Get_NilClient", func(t *testing.T) {
		_, err := service.Get(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("List_NilClient", func(t *testing.T) {
		_, err := service.List(ctx)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("NewIterator_NilClient", func(t *testing.T) {
		iter := service.NewIterator(ctx)
		assert.Nil(t, iter)
	})

	t.Run("Create_NilClient", func(t *testing.T) {
		floatingIP := &subnetModels.FloatingIp{}
		_, err := service.Create(ctx, floatingIP)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("CreateAsync_NilClient", func(t *testing.T) {
		floatingIP := &subnetModels.FloatingIp{}
		_, err := service.CreateAsync(ctx, floatingIP)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("Update_NilClient", func(t *testing.T) {
		floatingIP := &subnetModels.FloatingIp{}
		_, err := service.Update(ctx, "test-id", floatingIP)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("UpdateAsync_NilClient", func(t *testing.T) {
		floatingIP := &subnetModels.FloatingIp{}
		_, err := service.UpdateAsync(ctx, "test-id", floatingIP)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("Delete_NilClient", func(t *testing.T) {
		err := service.Delete(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("DeleteAsync_NilClient", func(t *testing.T) {
		_, err := service.DeleteAsync(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("AssignFloatingIp_NilClient", func(t *testing.T) {
		_, err := service.AssignFloatingIp(ctx, "test-id", subnetModels.VmNicAssociation{
			VmNicReference: ptr.To("nic-id"),
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("AssignFloatingIpAsync_NilClient", func(t *testing.T) {
		_, err := service.AssignFloatingIpAsync(ctx, "test-id", subnetModels.VmNicAssociation{
			VmNicReference: ptr.To("nic-id"),
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("DeassignFloatingIp_NilClient", func(t *testing.T) {
		_, err := service.DeassignFloatingIp(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("DeassignFloatingIpAsync_NilClient", func(t *testing.T) {
		_, err := service.DeassignFloatingIpAsync(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})
}

// TestFloatingIpsService_UnsupportedODataOptions tests that select and apply options are rejected.
func TestFloatingIpsService_UnsupportedODataOptions(t *testing.T) {
	service := NewFloatingIpsService(&v4prismGoClient.Client{})
	require.NotNil(t, service)
	ctx := context.Background()

	t.Run("List_WithSelect", func(t *testing.T) {
		_, err := service.List(ctx, converged.WithSelect("extId,name"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "select and apply options are not supported for listing")
		assert.Contains(t, err.Error(), "floating IP")
	})

	t.Run("List_WithApply", func(t *testing.T) {
		_, err := service.List(ctx, converged.WithApply("groupby((extId))"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "select and apply options are not supported for listing")
		assert.Contains(t, err.Error(), "floating IP")
	})

	t.Run("NewIterator_WithSelect", func(t *testing.T) {
		iter := service.NewIterator(ctx, converged.WithSelect("extId,name"))
		require.NotNil(t, iter)

		var gotError error
		for _, err := range iter {
			gotError = err
			break
		}
		require.Error(t, gotError)
		assert.Contains(t, gotError.Error(), "select and apply options are not supported for listing")
	})

	t.Run("NewIterator_WithApply", func(t *testing.T) {
		iter := service.NewIterator(ctx, converged.WithApply("groupby((extId))"))
		require.NotNil(t, iter)

		var gotError error
		for _, err := range iter {
			gotError = err
			break
		}
		require.Error(t, gotError)
		assert.Contains(t, gotError.Error(), "select and apply options are not supported for listing")
	})
}

func TestFloatingIpUpdateBody_PreservesAllocatedAddressAndProject(t *testing.T) {
	current := subnetModels.NewFloatingIp()
	current.Name = ptr.To("fip")
	current.Description = ptr.To("keep")
	current.ExternalSubnetReference = ptr.To("subnet-id")
	current.ProjectExtId = ptr.To("project-id")
	current.FloatingIp = subnetModels.NewFloatingIPAddress()
	current.FloatingIp.Ipv4 = subnetModels.NewFloatingIPv4Address()
	current.FloatingIp.Ipv4.Value = ptr.To("10.22.193.134")

	body, err := floatingIpUpdateBody(current)
	require.NoError(t, err)
	require.NotNil(t, body.FloatingIp)
	require.NotNil(t, body.FloatingIp.Ipv4)
	require.NotNil(t, body.FloatingIp.Ipv4.Value)
	assert.Equal(t, "10.22.193.134", *body.FloatingIp.Ipv4.Value)
	assert.Nil(t, body.FloatingIp.Ipv4.PrefixLength)
	require.NotNil(t, body.ProjectExtId)
	assert.Equal(t, "project-id", *body.ProjectExtId)
	assert.Equal(t, "keep", *body.Description)
}

// TestFloatingIpsService_AssignAssociationValidation tests association type checking.
func TestFloatingIpsService_AssignAssociationValidation(t *testing.T) {
	service := NewFloatingIpsService(&v4prismGoClient.Client{})
	require.NotNil(t, service)
	ctx := context.Background()

	t.Run("AssignFloatingIpAsync_RejectsInvalidAssociationType", func(t *testing.T) {
		_, err := service.AssignFloatingIpAsync(ctx, "test-id", "not-an-association")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "association must be VmNicAssociation, PrivateIpAssociation, or LoadBalancerSessionAssociation")
	})

	t.Run("AssignFloatingIpAsync_RejectsMissingVmNicReference", func(t *testing.T) {
		_, err := service.AssignFloatingIpAsync(ctx, "test-id", subnetModels.VmNicAssociation{})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "VmNicAssociation.VmNicReference is required")
	})

	t.Run("NormalizePrivateIpAssociation_StructLiteralDoesNotPanic", func(t *testing.T) {
		ip := commonConfig.NewIPAddress()
		ip.Ipv4 = commonConfig.NewIPv4Address()
		ip.Ipv4.Value = ptr.To("172.16.0.50")
		normalized, err := normalizeFloatingIpAssociation(subnetModels.PrivateIpAssociation{
			PrivateIp:    ip,
			VpcReference: ptr.To("vpc-id"),
		})
		require.NoError(t, err)

		body := subnetModels.NewFloatingIp()
		require.NoError(t, body.SetAssociation(normalized))
		require.NotNil(t, body.GetAssociation())
	})

	t.Run("AssignFloatingIpAsync_ValidAssociationFailsWithoutAPI", func(t *testing.T) {
		_, err := service.AssignFloatingIpAsync(ctx, "test-id", subnetModels.VmNicAssociation{
			VmNicReference: ptr.To("nic-id"),
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "floating IPs API is not initialized")
	})
}
