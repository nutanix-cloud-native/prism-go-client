// Package v4 provides integration tests for the converged client vm profiles service.
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run Testclient.TestVMProfilesIntegration
//	go test -v ./converged/v4 -run Testclient.TestDeployVmWithVmProfile

package v4

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"

	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// TestVMProfilesIntegration tests the client.VMProfiles with real Nutanix API calls
func TestVMProfilesIntegration(t *testing.T) {
	// Get credentials from environment
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	// Create converged client
	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ListVMProfiles", func(t *testing.T) {
		// Test basic list
		profiles, err := client.VMProfiles.List(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, profiles)
		assert.GreaterOrEqual(t, len(profiles), 0)

		t.Logf("Found %d VM profiles", len(profiles))

		// Print profile details
		for i, profile := range profiles {
			if i >= 5 { // Limit output to first 5
				break
			}
			if profile.ExtId != nil {
				if profile.Name != nil {
					t.Logf("  Profile %d: UUID=%s, Name=%s", i+1, *profile.ExtId, *profile.Name)
				}
			}
		}

		// Test list with limit
		profiles, err = client.VMProfiles.List(ctx, converged.WithLimit(5))
		assert.NoError(t, err)
		assert.NotNil(t, profiles)
		assert.LessOrEqual(t, len(profiles), 5)

		// Test list with page
		profiles, err = client.VMProfiles.List(ctx, converged.WithPage(0), converged.WithLimit(3))
		assert.NoError(t, err)
		assert.NotNil(t, profiles)
		assert.LessOrEqual(t, len(profiles), 3)

		// Test list with select fields
		profiles, err = client.VMProfiles.List(ctx, converged.WithSelect("extId,name"))
		assert.NoError(t, err)
		assert.NotNil(t, profiles)
	})

	t.Run("GetVMProfile", func(t *testing.T) {
		// First, list profiles to get a valid profile UUID
		profiles, err := client.VMProfiles.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(profiles) == 0 {
			t.Skip("No VM profiles available for testing")
		}

		// Get the first profile
		profileUUID := *profiles[0].ExtId
		require.NotEmpty(t, profileUUID)

		// Test Get profile
		profile, err := client.VMProfiles.Get(ctx, profileUUID)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, profileUUID, *profile.ExtId)

		if profile.Name != nil {
			t.Logf("Retrieved VM profile: %s with name: %s", profileUUID, *profile.Name)
		}
	})

	t.Run("NewIterator", func(t *testing.T) {
		// Test iterator
		iterator := client.VMProfiles.NewIterator(ctx)
		require.NotNil(t, iterator)

		// Collect profiles using iterator
		var profiles []vmmModels.VmProfile
		for profile, err := range iterator {
			if err != nil {
				break
			}
			profiles = append(profiles, profile)
			if len(profiles) >= 10 { // Limit to prevent long test runs
				break
			}
		}

		assert.GreaterOrEqual(t, len(profiles), 0)
		t.Logf("Iterated through %d VM profiles", len(profiles))
	})

	t.Run("ErrorHandling", func(t *testing.T) {
		// Test Get with invalid UUID
		_, err := client.VMProfiles.Get(ctx, "invalid-uuid")
		assert.Error(t, err)

		// Test Deploy with invalid profile UUID (using nil params - will default to empty)
		_, err = client.VMProfiles.DeployVmWithVmProfile(ctx, "invalid-uuid", nil)
		assert.Error(t, err)
	})
}

// TestVMProfilesWithRealEnvironment tests with real environment variables
func TestVMProfilesWithRealEnvironment(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("RealEnvironmentListProfiles", func(t *testing.T) {
		profiles, err := client.VMProfiles.List(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, profiles)

		t.Logf("Found %d VM profiles in the environment", len(profiles))
	})

	t.Run("RealEnvironmentGetProfile", func(t *testing.T) {
		profiles, err := client.VMProfiles.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(profiles) == 0 {
			t.Skip("No VM profiles available in the environment")
		}

		profileUUID := *profiles[0].ExtId
		require.NotEmpty(t, profileUUID)

		profile, err := client.VMProfiles.Get(ctx, profileUUID)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, profileUUID, *profile.ExtId)

		if profile.Name != nil {
			t.Logf("Retrieved VM profile: %s with name: %s", profileUUID, *profile.Name)
		} else {
			t.Logf("Retrieved VM profile: %s with name: <nil>", profileUUID)
		}
	})
}

func TestVMProfilesOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("WithLimit", func(t *testing.T) {
		profiles, err := client.VMProfiles.List(ctx, converged.WithLimit(2))
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(profiles), 2)
	})

	t.Run("WithPage", func(t *testing.T) {
		profiles, err := client.VMProfiles.List(ctx, converged.WithPage(0), converged.WithLimit(1))
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(profiles), 1)
	})

	t.Run("WithSelect", func(t *testing.T) {
		profiles, err := client.VMProfiles.List(ctx, converged.WithSelect("extId,name,description"))
		assert.NoError(t, err)
		assert.NotNil(t, profiles)
	})

	t.Run("WithOrderBy", func(t *testing.T) {
		profiles, err := client.VMProfiles.List(ctx, converged.WithOrderBy("name asc"))
		assert.NoError(t, err)
		assert.NotNil(t, profiles)
	})

	t.Run("WithFilter", func(t *testing.T) {
		// Test filtering by name (if you know a profile name)
		// This is just an example - adjust the filter based on your environment
		profiles, err := client.VMProfiles.List(ctx, converged.WithFilter("contains(name, '.medium')"))
		assert.NoError(t, err)
		assert.NotNil(t, profiles)
		t.Logf("Found %d profiles with '.medium' in name", len(profiles))
	})

	t.Run("MultipleOptions", func(t *testing.T) {
		profiles, err := client.VMProfiles.List(ctx,
			converged.WithLimit(3),
			converged.WithPage(0),
			converged.WithSelect("extId,name"),
			converged.WithOrderBy("name asc"),
		)
		assert.NoError(t, err)
		assert.NotNil(t, profiles)
		assert.LessOrEqual(t, len(profiles), 3)
	})
}

// TestVMProfilesErrorScenarios tests error handling scenarios
func TestVMProfilesErrorScenarios(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("InvalidUUID", func(t *testing.T) {
		_, err := client.VMProfiles.Get(ctx, "invalid-uuid-format")
		assert.Error(t, err)
	})

	t.Run("NonExistentUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.VMProfiles.Get(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("InvalidOptions", func(t *testing.T) {
		_, err := client.VMProfiles.List(ctx, converged.WithFilter("invalid filter syntax"))
		assert.Error(t, err)

		assert.NotPanics(t, func() {
			_, _ = client.VMProfiles.List(ctx, converged.WithSelect("invalidField"))
		})
	})
}

func TestDeployVmWithVmProfile(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	// Get first available VM profile
	profiles, err := client.VMProfiles.List(ctx, converged.WithLimit(1))
	require.NoError(t, err)
	if len(profiles) == 0 {
		t.Skip("No VM profiles available for testing")
	}
	profileUUID := *profiles[0].ExtId
	if profiles[0].Name != nil {
		t.Logf("Using VM profile: %s (UUID: %s)", *profiles[0].Name, profileUUID)
	}

	// Get first available cluster
	clusters, err := client.Clusters.List(ctx)
	require.NoError(t, err)
	if len(clusters) == 0 {
		t.Fatalf("No clusters available for deployment")
	}
	clusterUUID := *clusters[0].ExtId
	if clusters[0].Name != nil {
		t.Logf("Using cluster: %s (UUID: %s)", *clusters[0].Name, clusterUUID)
	}

	// Get first available subnet (optional)
	var subnetUUID *string
	subnets, err := client.Subnets.List(ctx)
	if err == nil && len(subnets) > 0 {
		subnetUUID = subnets[0].ExtId
		if subnets[0].Name != nil {
			t.Logf("Using subnet: %s (UUID: %s)", *subnets[0].Name, *subnetUUID)
		}
	}

	// Generate VM name
	vmName := "test-vm-" + time.Now().Format("150405") // HHMMSS format

	// Create deployment parameters
	clusterRef := vmmModels.ClusterReference{
		ExtId: &clusterUUID,
	}

	deployParams := &vmmModels.DeployVmFromVmProfileParams{
		VmName:  &vmName,
		Cluster: &clusterRef,
	}

	// Don't add Nics - the VM profile already has NIC configuration
	// If you need to override NICs, you must reference the profile's NIC ExtIds

	t.Run("DeployVm", func(t *testing.T) {
		t.Logf("Deploying VM '%s' from profile UUID: %s", vmName, profileUUID)

		operation, err := client.VMProfiles.DeployVmWithVmProfile(ctx, profileUUID, deployParams)
		require.NoError(t, err)
		require.NotNil(t, operation)

		taskUUID := operation.UUID()
		t.Logf("Deployment initiated. Task UUID: %s", taskUUID)
	})

	t.Run("DeployVmWithInvalidProfile", func(t *testing.T) {
		_, err := client.VMProfiles.DeployVmWithVmProfile(ctx, "invalid-profile-uuid", deployParams)
		assert.Error(t, err)
	})
}
