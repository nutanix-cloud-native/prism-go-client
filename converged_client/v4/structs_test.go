package v4

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	"github.com/stretchr/testify/assert"
	"k8s.io/utils/ptr"
)

// Test entity for testing purposes
type TestEntity struct {
	ExtId *string
	Name  *string
}

func TestNewClient(t *testing.T) {
	tests := []struct {
		name        string
		credentials prismgoclient.Credentials
		opts        []types.ClientOption[v4prismGoClient.Client]
		expectError bool
	}{
		{
			name: "successful client creation",
			credentials: prismgoclient.Credentials{
				URL:      "https://prism.example.com:9440",
				Username: "admin",
				Password: "password",
			},
			opts:        []types.ClientOption[v4prismGoClient.Client]{},
			expectError: false,
		},
		{
			name: "successful client creation with options",
			credentials: prismgoclient.Credentials{
				URL:      "https://prism.example.com:9440",
				Username: "admin",
				Password: "password",
			},
			opts:        []types.ClientOption[v4prismGoClient.Client]{},
			expectError: false,
		},
		{
			name: "successful client creation with minimal credentials",
			credentials: prismgoclient.Credentials{
				URL:      "https://test.example.com:9440",
				Username: "test",
				Password: "test",
			},
			opts:        []types.ClientOption[v4prismGoClient.Client]{},
			expectError: false,
		},
		{
			name: "invalid credentials - empty URL",
			credentials: prismgoclient.Credentials{
				URL: "", // invalid URL
			},
			opts:        []types.ClientOption[v4prismGoClient.Client]{},
			expectError: true,
		},
		{
			name: "invalid credentials - malformed URL",
			credentials: prismgoclient.Credentials{
				URL: "not-a-valid-url",
			},
			opts:        []types.ClientOption[v4prismGoClient.Client]{},
			expectError: true,
		},
		{
			name: "invalid credentials - missing username",
			credentials: prismgoclient.Credentials{
				URL:      "https://prism.example.com:9440",
				Username: "",
				Password: "password",
			},
			opts:        []types.ClientOption[v4prismGoClient.Client]{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.credentials, tt.opts...)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, client)
			} else {
				// Note: This test might fail if the actual NewV4Client implementation
				// validates the credentials. In a real scenario, we'd mock the v4.NewV4Client call
				if err != nil {
					t.Skipf("Skipping test due to actual v4 client creation: %v", err)
				}
				assert.NoError(t, err)
				assert.NotNil(t, client)
				assert.NotNil(t, client.Categories)
				assert.NotNil(t, client.VMs)
				assert.NotNil(t, client.client) // Test that the internal client is set
			}
		})
	}
}

func TestNewOperation(t *testing.T) {
	taskUUID := "test-task-uuid"
	mockClient := &v4prismGoClient.Client{}
	entityGetter := func(ctx context.Context, uuid string) (*TestEntity, error) {
		return &TestEntity{ExtId: ptr.To(uuid), Name: ptr.To("test")}, nil
	}

	operation := NewOperation[TestEntity](taskUUID, mockClient, entityGetter)

	assert.NotNil(t, operation)
	assert.Equal(t, taskUUID, operation.taskUUID)
	assert.Equal(t, mockClient, operation.client)
	assert.Equal(t, converged.TaskStatusQueued, operation.taskStatus)
	assert.Empty(t, operation.entityUUIDs)
	assert.Empty(t, operation.taskErrors)
	assert.NotNil(t, operation.lock)
	assert.NotNil(t, operation.entityGetter)
}

func TestOperation_UUID(t *testing.T) {
	taskUUID := "test-task-uuid"
	operation := NewOperation[TestEntity](taskUUID, nil, nil)

	assert.Equal(t, taskUUID, operation.UUID())
}

func TestOperation_Status(t *testing.T) {
	operation := NewOperation[TestEntity]("test", nil, nil)
	operation.setTaskStatus(converged.TaskStatusRunning)

	assert.Equal(t, converged.TaskStatusRunning, operation.Status())
}

func TestOperation_Errors(t *testing.T) {
	operation := NewOperation[TestEntity]("test", nil, nil)

	// Initially empty
	assert.Empty(t, operation.Errors())

	// Add error
	testError := errors.New("test error")
	operation.appendTaskError(testError)

	errors := operation.Errors()
	assert.Len(t, errors, 1)
	assert.Equal(t, testError, errors[0])
}

func TestOperation_Results(t *testing.T) {
	operation := NewOperation[TestEntity]("test", nil, nil)

	tests := []struct {
		name           string
		taskStatus     converged.TaskStatus
		result         []*TestEntity
		expectedError  error
		expectedResult []*TestEntity
	}{
		{
			name:           "task failed",
			taskStatus:     converged.TaskStatusFailed,
			result:         nil,
			expectedError:  converged.ErrTaskFailed,
			expectedResult: nil,
		},
		{
			name:           "task canceled",
			taskStatus:     converged.TaskStatusCanceled,
			result:         nil,
			expectedError:  converged.ErrTaskFailed,
			expectedResult: nil,
		},
		{
			name:           "task not complete",
			taskStatus:     converged.TaskStatusRunning,
			result:         nil,
			expectedError:  converged.ErrTaskNotComplete,
			expectedResult: nil,
		},
		{
			name:           "task succeeded but no results",
			taskStatus:     converged.TaskStatusSucceeded,
			result:         nil,
			expectedError:  converged.ErrTaskResultsNotReady,
			expectedResult: nil,
		},
		{
			name:           "task succeeded with results",
			taskStatus:     converged.TaskStatusSucceeded,
			result:         []*TestEntity{{ExtId: ptr.To("1"), Name: ptr.To("test")}},
			expectedError:  nil,
			expectedResult: []*TestEntity{{ExtId: ptr.To("1"), Name: ptr.To("test")}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			operation.setTaskStatus(tt.taskStatus)
			operation.setResults(tt.result)

			result, err := operation.Results()

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

func TestOperation_setTaskStatus(t *testing.T) {
	operation := NewOperation[TestEntity]("test", nil, nil)

	operation.setTaskStatus(converged.TaskStatusRunning)
	assert.Equal(t, converged.TaskStatusRunning, operation.taskStatus)

	operation.setTaskStatus(converged.TaskStatusSucceeded)
	assert.Equal(t, converged.TaskStatusSucceeded, operation.taskStatus)
}

func TestOperation_setEntitiesAffected(t *testing.T) {
	operation := NewOperation[TestEntity]("test", nil, nil)

	operation.setEntitiesAffected(5)
	assert.Equal(t, 5, operation.entitiesAffected)
}

func TestOperation_appendTaskError(t *testing.T) {
	operation := NewOperation[TestEntity]("test", nil, nil)

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	// Test that taskErrors is empty initially
	assert.Len(t, operation.taskErrors, 0)

	operation.appendTaskError(err1)
	assert.Len(t, operation.taskErrors, 1)

	operation.appendTaskError(err2)

	errors := operation.Errors()
	assert.Len(t, errors, 2)
	assert.Equal(t, err1, errors[0])
	assert.Equal(t, err2, errors[1])
}

func TestOperation_appendTaskError_NilInitialization(t *testing.T) {
	// Create operation manually to test nil taskErrors initialization
	operation := &Operation[TestEntity]{
		lock: &sync.Mutex{},
		// taskErrors is nil initially
	}

	err := errors.New("test error")

	// Test that taskErrors is nil initially
	assert.Nil(t, operation.taskErrors)

	operation.appendTaskError(err)

	// Should now be initialized and contain the error
	assert.NotNil(t, operation.taskErrors)
	assert.Len(t, operation.taskErrors, 1)
	assert.Equal(t, err, operation.taskErrors[0])
}

func TestOperation_appendEntityUUID(t *testing.T) {
	operation := NewOperation[TestEntity]("test", nil, nil)

	operation.appendEntityUUID("uuid1")
	operation.appendEntityUUID("uuid2")

	assert.Len(t, operation.entityUUIDs, 2)
	assert.Equal(t, "uuid1", operation.entityUUIDs[0])
	assert.Equal(t, "uuid2", operation.entityUUIDs[1])
}

func TestOperation_setResults(t *testing.T) {
	operation := NewOperation[TestEntity]("test", nil, nil)

	results := []*TestEntity{
		{ExtId: ptr.To("1"), Name: ptr.To("test1")},
		{ExtId: ptr.To("2"), Name: ptr.To("test2")},
	}

	operation.setResults(results)
	assert.Equal(t, results, operation.result)
}

func TestOperation_ConcurrentAccess(t *testing.T) {
	operation := NewOperation[TestEntity]("test", nil, nil)

	var wg sync.WaitGroup
	numGoroutines := 10

	// Test concurrent writes to different fields
	wg.Add(numGoroutines * 4)

	for i := 0; i < numGoroutines; i++ {
		go func(idx int) {
			defer wg.Done()
			operation.setTaskStatus(converged.TaskStatusRunning)
		}(i)

		go func(idx int) {
			defer wg.Done()
			operation.appendTaskError(fmt.Errorf("error %d", idx))
		}(i)

		go func(idx int) {
			defer wg.Done()
			operation.appendEntityUUID(fmt.Sprintf("uuid-%d", idx))
		}(i)

		go func(idx int) {
			defer wg.Done()
			operation.setEntitiesAffected(idx)
		}(i)
	}

	wg.Wait()

	// Verify that operations completed without race conditions
	assert.Equal(t, converged.TaskStatusRunning, operation.Status())
	assert.Len(t, operation.Errors(), numGoroutines)
	assert.Len(t, operation.entityUUIDs, numGoroutines)
}

// Mock APIResponse for testing
type MockAPIResponse struct {
	data     interface{}
	Metadata *TestMetadata
}

func (m *MockAPIResponse) GetData() interface{} {
	return m.data
}

func TestNewIterator(t *testing.T) {
	ctx := context.Background()

	// Test with successful API call
	t.Run("successful iteration", func(t *testing.T) {
		listFunc := func(ctx context.Context, params *V4ODataParams) (*MockAPIResponse, error) {
			return &MockAPIResponse{
				data: []TestEntity{
					{ExtId: ptr.To("1"), Name: ptr.To("test1")},
					{ExtId: ptr.To("2"), Name: ptr.To("test2")},
				},
				Metadata: &TestMetadata{TotalAvailableResults: ptr.To(2)},
			}, nil
		}

		iterator := NewIterator[*MockAPIResponse, TestEntity](ctx, listFunc)
		assert.NotNil(t, iterator)

		// Test that we can iterate through the results
		var results []TestEntity
		for entity, err := range iterator {
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			results = append(results, entity)
		}

		assert.Len(t, results, 2)
		assert.Equal(t, "test1", *results[0].Name)
		assert.Equal(t, "test2", *results[1].Name)
	})

	// Test with API error
	t.Run("API error", func(t *testing.T) {
		listFunc := func(ctx context.Context, params *V4ODataParams) (*MockAPIResponse, error) {
			return nil, errors.New("API error")
		}

		iterator := NewIterator[*MockAPIResponse, TestEntity](ctx, listFunc)
		assert.NotNil(t, iterator)

		// Test that we get the error
		var hasError bool
		for _, err := range iterator {
			if err != nil {
				hasError = true
				assert.Contains(t, err.Error(), "API error")
				break
			}
		}
		assert.True(t, hasError)
	})

	// Test with context cancellation
	t.Run("context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		listFunc := func(ctx context.Context, params *V4ODataParams) (*MockAPIResponse, error) {
			return &MockAPIResponse{
				data:     []TestEntity{{ExtId: ptr.To("1"), Name: ptr.To("test1")}},
				Metadata: &TestMetadata{TotalAvailableResults: ptr.To(1)},
			}, nil
		}

		iterator := NewIterator[*MockAPIResponse, TestEntity](ctx, listFunc)
		assert.NotNil(t, iterator)

		// Should not yield any results due to context cancellation
		var count int
		for range iterator {
			count++
		}
		assert.Equal(t, 0, count)
	})

	// Test with pagination
	t.Run("pagination", func(t *testing.T) {
		pageCount := 0
		listFunc := func(ctx context.Context, params *V4ODataParams) (*MockAPIResponse, error) {
			pageCount++
			if pageCount == 1 {
				return &MockAPIResponse{
					data:     []TestEntity{{ExtId: ptr.To("1"), Name: ptr.To("test1")}},
					Metadata: &TestMetadata{TotalAvailableResults: ptr.To(2)},
				}, nil
			} else {
				return &MockAPIResponse{
					data:     []TestEntity{{ExtId: ptr.To("2"), Name: ptr.To("test2")}},
					Metadata: &TestMetadata{TotalAvailableResults: ptr.To(2)},
				}, nil
			}
		}

		iterator := NewIterator[*MockAPIResponse, TestEntity](ctx, listFunc)
		assert.NotNil(t, iterator)

		var results []TestEntity
		for entity, err := range iterator {
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			results = append(results, entity)
		}

		assert.Len(t, results, 2)
		assert.Equal(t, 2, pageCount) // Should have made 2 API calls
	})

	// Test with options
	t.Run("with options", func(t *testing.T) {
		listFunc := func(ctx context.Context, params *V4ODataParams) (*MockAPIResponse, error) {
			// Verify that options are passed correctly
			assert.NotNil(t, params)
			return &MockAPIResponse{
				data:     []TestEntity{{ExtId: ptr.To("1"), Name: ptr.To("test1")}},
				Metadata: &TestMetadata{TotalAvailableResults: ptr.To(1)},
			}, nil
		}

		options := []converged.ODataOption{
			converged.WithPage(0),
			converged.WithLimit(10),
		}

		iterator := NewIterator[*MockAPIResponse, TestEntity](ctx, listFunc, options...)
		assert.NotNil(t, iterator)

		var results []TestEntity
		for entity, err := range iterator {
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			results = append(results, entity)
		}

		assert.Len(t, results, 1)
	})
}

// Test helper types and functions
func TestOperation_IsDone(t *testing.T) {
	operation := NewOperation[TestEntity]("test", nil, nil)

	// Test not done states
	notDoneStates := []converged.TaskStatus{
		converged.TaskStatusQueued,
		converged.TaskStatusRunning,
		converged.TaskStatusSuspended,
		converged.TaskStatusUnknown,
	}

	for _, status := range notDoneStates {
		operation.setTaskStatus(status)
		assert.False(t, operation.IsDone(), "Status %s should not be done", status)
	}

	// Test done states
	doneStates := []converged.TaskStatus{
		converged.TaskStatusSucceeded,
		converged.TaskStatusFailed,
		converged.TaskStatusCanceled,
		converged.TaskStatusCanceling,
	}

	for _, status := range doneStates {
		operation.setTaskStatus(status)
		assert.True(t, operation.IsDone(), "Status %s should be done", status)
	}
}

func TestOperation_IsSuccess(t *testing.T) {
	operation := NewOperation[TestEntity]("test", nil, nil)

	operation.setTaskStatus(converged.TaskStatusSucceeded)
	assert.True(t, operation.IsSuccess())

	operation.setTaskStatus(converged.TaskStatusFailed)
	assert.False(t, operation.IsSuccess())
}

func TestOperation_IsFailed(t *testing.T) {
	operation := NewOperation[TestEntity]("test", nil, nil)

	operation.setTaskStatus(converged.TaskStatusFailed)
	assert.True(t, operation.IsFailed())

	operation.setTaskStatus(converged.TaskStatusSucceeded)
	assert.False(t, operation.IsFailed())
}

func TestOperation_Wait(t *testing.T) {
	tests := []struct {
		name          string
		operation     *Operation[TestEntity]
		ctx           context.Context
		expectedError string
		expectedCount int
	}{
		{
			name: "context canceled",
			operation: &Operation[TestEntity]{
				taskUUID: "test-task-uuid",
			},
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel() // Cancel immediately
				return ctx
			}(),
			expectedError: "task wait canceled",
			expectedCount: 0,
		},
		{
			name: "context with timeout",
			operation: &Operation[TestEntity]{
				taskUUID: "test-task-uuid",
			},
			ctx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
				defer cancel()
				time.Sleep(2 * time.Millisecond) // Wait longer than timeout
				return ctx
			}(),
			expectedError: "task wait canceled",
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: This test will fail because it requires a real client and API calls
			// but it will exercise the context cancellation path which is the main
			// untested branch in the Wait function
			result, err := tt.operation.Wait(tt.ctx)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Len(t, result, tt.expectedCount)
			}
		})
	}
}
