package v4

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
)

// TestAlertsIntegration tests the client.Alerts with real Nutanix API calls.
func TestAlertsIntegration(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ListAlerts", func(t *testing.T) {
		alerts, err := client.Alerts.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(alerts), 0)
	})

	t.Run("GetAlert", func(t *testing.T) {
		// Discover a real alert via List, then Get it by external identifier.
		alerts, err := client.Alerts.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(alerts) == 0 || alerts[0].ExtId == nil {
			t.Skip("No alerts available for testing")
		}

		alertExtID := *alerts[0].ExtId
		alert, err := client.Alerts.Get(ctx, alertExtID)
		assert.NoError(t, err)
		assert.NotNil(t, alert)
		assert.Equal(t, alertExtID, *alert.ExtId)
	})

	t.Run("ListAlertsWithFilter", func(t *testing.T) {
		// Filter on a severity value; an empty result is acceptable, the call
		// must simply succeed and propagate the filter without error.
		alerts, err := client.Alerts.List(ctx, converged.WithFilter(
			"severity eq Monitoring.Alert.AlertSeverity'kCritical'"))
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(alerts), 0)
	})
}

// TestAlertsService_ErrorHandling tests error handling for nil client.
func TestAlertsService_ErrorHandling(t *testing.T) {
	service := NewAlertsService(nil)
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
		iterator := service.NewIterator(ctx)
		assert.Nil(t, iterator)
	})
}
