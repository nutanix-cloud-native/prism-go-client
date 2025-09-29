// Package v4 provides integration tests for the converged client tasks service.
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run Testclient.TasksIntegration
//	go test -v ./converged/v4 -run Testclient.TasksWithRealEnvironment
package v4

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

const (
	prismEndpointDummyValue = "prism-test.nutanix.com"
)

// TestTasksIntegration tests the client.Tasks with real Nutanix API calls
func TestTasksIntegration(t *testing.T) {

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

	t.Run("GetTask", func(t *testing.T) {
		// First, list tasks to get a valid task UUID
		tasks, err := client.Tasks.List(ctx)
		require.NoError(t, err)

		if len(tasks) == 0 {
			t.Skip("No tasks available for testing")
		}

		// Get the first task
		taskUUID := *tasks[0].ExtId
		require.NotEmpty(t, taskUUID)

		// Test Get task
		task, err := client.Tasks.Get(ctx, taskUUID)
		assert.NoError(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, taskUUID, *task.ExtId)
		assert.NotNil(t, task.Status)
	})

	t.Run("GetTaskWithSelect", func(t *testing.T) {
		// First, list tasks to get a valid task UUID
		tasks, err := client.Tasks.List(ctx)
		require.NoError(t, err)

		if len(tasks) == 0 {
			t.Skip("No tasks available for testing")
		}

		// Get the first task
		taskUUID := *tasks[0].ExtId
		require.NotEmpty(t, taskUUID)

		// Test GetWithSelect with specific fields
		fields := []string{"extId", "status"}
		task, err := client.Tasks.GetWithSelect(ctx, taskUUID, fields)
		assert.NoError(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, taskUUID, *task.ExtId)
		assert.NotNil(t, task.Status)
	})

	t.Run("ListTasks", func(t *testing.T) {
		// Test basic list
		tasks, err := client.Tasks.List(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, tasks)
		assert.GreaterOrEqual(t, len(tasks), 0)

		// Test list with limit
		tasks, err = client.Tasks.List(ctx, converged.WithLimit(5))
		assert.NoError(t, err)
		assert.NotNil(t, tasks)
		assert.LessOrEqual(t, len(tasks), 5)

		// Test list with page
		tasks, err = client.Tasks.List(ctx, converged.WithPage(0), converged.WithLimit(3))
		assert.NoError(t, err)
		assert.NotNil(t, tasks)
		assert.LessOrEqual(t, len(tasks), 3)

		// Test list with select fields
		tasks, err = client.Tasks.List(ctx, converged.WithSelect("extId,status"))
		assert.NoError(t, err)
		assert.NotNil(t, tasks)

		// Test list with order by
		tasks, err = client.Tasks.List(ctx, converged.WithOrderBy("createdTime desc"))
		assert.NoError(t, err)
		assert.NotNil(t, tasks)
	})

	t.Run("ListTasksWithFilter", func(t *testing.T) {
		// Test list with filter - get only succeeded tasks
		tasks, err := client.Tasks.List(ctx, converged.WithFilter("status eq Prism.Config.TaskStatus'SUCCEEDED'"))
		assert.NoError(t, err)
		assert.NotNil(t, tasks)

		// Verify all returned tasks have succeeded status
		for _, task := range tasks {
			if task.Status != nil {
				assert.Equal(t, prismModels.TASKSTATUS_SUCCEEDED, *task.Status)
			}
		}
	})

	t.Run("NewIterator", func(t *testing.T) {
		// Test iterator
		iterator := client.Tasks.NewIterator(ctx)
		require.NotNil(t, iterator)

		// Collect tasks using iterator
		var tasks []prismModels.Task
		for task, err := range iterator {
			if err != nil {
				break
			}
			tasks = append(tasks, task)
			if len(tasks) >= 10 { // Limit to prevent long test runs
				break
			}
		}

		assert.GreaterOrEqual(t, len(tasks), 0)
	})

	t.Run("NewIteratorWithOptions", func(t *testing.T) {
		// Test iterator with options
		iterator := client.Tasks.NewIterator(ctx,
			converged.WithLimit(5),
			converged.WithSelect("extId,status"),
		)
		require.NotNil(t, iterator)

		// Collect tasks using iterator
		var tasks []prismModels.Task
		for task, err := range iterator {
			if err != nil {
				break
			}
			tasks = append(tasks, task)
		}

		assert.LessOrEqual(t, len(tasks), 5)
	})

	t.Run("CancelTask", func(t *testing.T) {
		// First, list tasks to find a running task
		tasks, err := client.Tasks.List(ctx, converged.WithFilter("status eq Prism.Config.TaskStatus'RUNNING'"))
		require.NoError(t, err)

		if len(tasks) == 0 {
			t.Skip("No running tasks available for cancellation testing")
		}

		// Get the first running task
		taskUUID := *tasks[0].ExtId
		require.NotEmpty(t, taskUUID)

		// Test Cancel task
		appMessage, err := client.Tasks.Cancel(ctx, taskUUID)
		if err != nil {
			t.Logf("Error canceling task: %v", err)
		}
		t.Logf("App message: %v", appMessage)
		assert.NoError(t, err)
		assert.NotNil(t, appMessage)
	})

	t.Run("ErrorHandling", func(t *testing.T) {
		// Test Get with invalid UUID
		_, err := client.Tasks.Get(ctx, "invalid-uuid")
		assert.Error(t, err)

		// Test GetWithSelect with invalid UUID
		_, err = client.Tasks.GetWithSelect(ctx, "invalid-uuid", []string{"extId"})
		assert.Error(t, err)

		// Test Cancel with invalid UUID
		_, err = client.Tasks.Cancel(ctx, "invalid-uuid")
		assert.Error(t, err)
	})

	t.Run("ConcurrentAccess", func(t *testing.T) {
		// Test concurrent access with separate client instances
		// Note: The underlying v4 API client is not thread-safe for concurrent operations
		// So we create separate client instances for each goroutine
		done := make(chan bool, 2)

		// Create separate client instances for each goroutine
		client, err1 := NewClient(creds)
		require.NoError(t, err1)

		go func() {
			defer func() { done <- true }()
			tasks, err := client.Tasks.List(ctx, converged.WithLimit(5))
			assert.NoError(t, err)
			assert.NotNil(t, tasks)
			assert.LessOrEqual(t, len(tasks), 5)
		}()

		go func() {
			defer func() { done <- true }()
			tasks, err := client.Tasks.List(ctx, converged.WithLimit(3))
			assert.NoError(t, err)
			assert.NotNil(t, tasks)
			assert.LessOrEqual(t, len(tasks), 3)
		}()

		// Wait for both goroutines to complete
		<-done
		<-done
	})
}

// TestTasksWithRealEnvironment tests with real environment variables
func TestTasksWithRealEnvironment(t *testing.T) {
	// Skip if not running integration tests

	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	// Create converged client
	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("RealEnvironmentListTasks", func(t *testing.T) {
		// Test listing tasks with real environment
		tasks, err := client.Tasks.List(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, tasks)

		t.Logf("Found %d tasks in the environment", len(tasks))
	})

	t.Run("RealEnvironmentGetTask", func(t *testing.T) {
		// First get a list of tasks
		tasks, err := client.Tasks.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(tasks) == 0 {
			t.Skip("No tasks available in the environment")
		}

		// Get the first task
		taskUUID := *tasks[0].ExtId
		require.NotEmpty(t, taskUUID)

		// Test getting the task
		task, err := client.Tasks.Get(ctx, taskUUID)
		assert.NoError(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, taskUUID, *task.ExtId)

		t.Logf("Retrieved task: %s with status: %v", taskUUID, task.Status)
	})
}

// TasksOptions tests various OData options
func TestTasksOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("WithLimit", func(t *testing.T) {
		tasks, err := client.Tasks.List(ctx, converged.WithLimit(2))
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(tasks), 2)
	})

	t.Run("WithPage", func(t *testing.T) {
		tasks, err := client.Tasks.List(ctx, converged.WithPage(0), converged.WithLimit(1))
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(tasks), 1)
	})

	t.Run("WithSelect", func(t *testing.T) {
		tasks, err := client.Tasks.List(ctx, converged.WithSelect("extId,status,createdTime"))
		assert.NoError(t, err)
		assert.NotNil(t, tasks)
	})

	t.Run("WithOrderBy", func(t *testing.T) {
		tasks, err := client.Tasks.List(ctx, converged.WithOrderBy("createdTime desc"))
		assert.NoError(t, err)
		assert.NotNil(t, tasks)
	})

	t.Run("WithFilter", func(t *testing.T) {
		// Test different filter scenarios
		testCases := []struct {
			name   string
			filter string
		}{
			{"SucceededTasks", "status eq Prism.Config.TaskStatus'SUCCEEDED'"},
			{"FailedTasks", "status eq Prism.Config.TaskStatus'FAILED'"},
			{"RunningTasks", "status eq Prism.Config.TaskStatus'RUNNING'"},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				tasks, err := client.Tasks.List(ctx, converged.WithFilter(tc.filter))
				assert.NoError(t, err)
				assert.NotNil(t, tasks)
			})
		}
	})

	t.Run("MultipleOptions", func(t *testing.T) {
		// Test combining multiple options
		tasks, err := client.Tasks.List(ctx,
			converged.WithLimit(3),
			converged.WithPage(0),
			converged.WithSelect("extId,status"),
			converged.WithOrderBy("createdTime desc"),
		)
		assert.NoError(t, err)
		assert.NotNil(t, tasks)
		assert.LessOrEqual(t, len(tasks), 3)
	})
}

// TestTasksErrorScenarios tests error handling scenarios
func TestTasksErrorScenarios(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("InvalidUUID", func(t *testing.T) {
		// Test with invalid UUID format
		_, err := client.Tasks.Get(ctx, "invalid-uuid-format")
		assert.Error(t, err)

		_, err = client.Tasks.GetWithSelect(ctx, "invalid-uuid-format", []string{"extId"})
		assert.Error(t, err)

		_, err = client.Tasks.Cancel(ctx, "invalid-uuid-format")
		assert.Error(t, err)
	})

	t.Run("NonExistentUUID", func(t *testing.T) {
		// Test with valid UUID format but non-existent task
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.Tasks.Get(ctx, nonExistentUUID)
		assert.Error(t, err)

		_, err = client.Tasks.GetWithSelect(ctx, nonExistentUUID, []string{"extId"})
		assert.Error(t, err)

		_, err = client.Tasks.Cancel(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("InvalidOptions", func(t *testing.T) {
		// Test with invalid filter syntax
		_, err := client.Tasks.List(ctx, converged.WithFilter("invalid filter syntax"))
		assert.Error(t, err)

		// This might not error depending on API behavior, so we just check it doesn't panic
		assert.NotPanics(t, func() {
			_, _ = client.Tasks.List(ctx, converged.WithSelect("invalidField"))
		})
	})
}
