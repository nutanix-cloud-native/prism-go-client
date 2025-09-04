package test

import (
	"context"
	"errors"
	"testing"

	convergedclient "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	"github.com/nutanix-cloud-native/prism-go-client/converged_client/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"k8s.io/utils/ptr"
)

// Test custom errors
func TestCustomErrors(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "ErrTaskNotComplete",
			err:      convergedclient.ErrTaskNotComplete,
			expected: "task not yet complete",
		},
		{
			name:     "ErrTaskResultsNotReady",
			err:      convergedclient.ErrTaskResultsNotReady,
			expected: "task results not ready",
		},
		{
			name:     "ErrTaskFailed",
			err:      convergedclient.ErrTaskFailed,
			expected: "task failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Error(t, tt.err)
			assert.Equal(t, tt.expected, tt.err.Error())
		})
	}
}

// Test TaskStatus constants
func TestTaskStatusConstants(t *testing.T) {
	tests := []struct {
		name     string
		status   convergedclient.TaskStatus
		expected string
	}{
		{
			name:     "TaskStatusFailed",
			status:   convergedclient.TaskStatusFailed,
			expected: "FAILED",
		},
		{
			name:     "TaskStatusQueued",
			status:   convergedclient.TaskStatusQueued,
			expected: "QUEUED",
		},
		{
			name:     "TaskStatusRunning",
			status:   convergedclient.TaskStatusRunning,
			expected: "RUNNING",
		},
		{
			name:     "TaskStatusSuspended",
			status:   convergedclient.TaskStatusSuspended,
			expected: "SUSPENDED",
		},
		{
			name:     "TaskStatusSucceeded",
			status:   convergedclient.TaskStatusSucceeded,
			expected: "SUCCEEDED",
		},
		{
			name:     "TaskStatusCanceled",
			status:   convergedclient.TaskStatusCanceled,
			expected: "CANCELED",
		},
		{
			name:     "TaskStatusCanceling",
			status:   convergedclient.TaskStatusCanceling,
			expected: "CANCELLING",
		},
		{
			name:     "TaskStatusUnknown",
			status:   convergedclient.TaskStatusUnknown,
			expected: "UNKNOWN",
		},
		{
			name:     "TaskStatusRedacted",
			status:   convergedclient.TaskStatusRedacted,
			expected: "REDACTED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, string(tt.status))
		})
	}
}

// Test TaskStatus type behavior
func TestTaskStatusComparison(t *testing.T) {
	// Test that TaskStatus constants can be compared
	assert.NotEqual(t, convergedclient.TaskStatusQueued, convergedclient.TaskStatusRunning)
	assert.NotEqual(t, convergedclient.TaskStatusRunning, convergedclient.TaskStatusSucceeded)
	assert.NotEqual(t, convergedclient.TaskStatusSucceeded, convergedclient.TaskStatusFailed)

	// Test that the same constants are equal
	assert.Equal(t, convergedclient.TaskStatusQueued, convergedclient.TaskStatusQueued)
	assert.Equal(t, convergedclient.TaskStatusSucceeded, convergedclient.TaskStatusSucceeded)

	// Test conversion to string
	status := convergedclient.TaskStatusRunning
	assert.Equal(t, "RUNNING", string(status))
}

// Test error behavior and identity
func TestErrorIdentity(t *testing.T) {
	// Test that errors are distinct
	assert.NotEqual(t, convergedclient.ErrTaskNotComplete, convergedclient.ErrTaskFailed)
	assert.NotEqual(t, convergedclient.ErrTaskFailed, convergedclient.ErrTaskResultsNotReady)

	// Test that the same error variables are identical
	assert.Equal(t, convergedclient.ErrTaskNotComplete, convergedclient.ErrTaskNotComplete)
	assert.Equal(t, convergedclient.ErrTaskFailed, convergedclient.ErrTaskFailed)

	// Test error wrapping behavior
	wrappedErr := errors.Join(convergedclient.ErrTaskFailed, errors.New("additional context"))
	assert.ErrorIs(t, wrappedErr, convergedclient.ErrTaskFailed)
}

// Test OData functional options with generated mocks
func TestWithPage(t *testing.T) {
	tests := []struct {
		name        string
		page        int
		shouldError bool
		expectError bool
	}{
		{
			name:        "successful page option",
			page:        5,
			shouldError: false,
			expectError: false,
		},
		{
			name:        "page option with error",
			page:        10,
			shouldError: true,
			expectError: true,
		},
		{
			name:        "zero page",
			page:        0,
			shouldError: false,
			expectError: false,
		},
		{
			name:        "negative page",
			page:        -1,
			shouldError: false,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockOData := mocks.NewMockODataOptions(ctrl)
			if tt.shouldError {
				mockOData.EXPECT().SetPageOption(tt.page).Return(errors.New("mock error"))
			} else {
				mockOData.EXPECT().SetPageOption(tt.page).Return(nil)
			}

			option := convergedclient.WithPage(tt.page)
			err := option(mockOData)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWithLimit(t *testing.T) {
	tests := []struct {
		name        string
		limit       int
		shouldError bool
		expectError bool
	}{
		{
			name:        "successful limit option",
			limit:       10,
			shouldError: false,
			expectError: false,
		},
		{
			name:        "limit option with error",
			limit:       20,
			shouldError: true,
			expectError: true,
		},
		{
			name:        "zero limit",
			limit:       0,
			shouldError: false,
			expectError: false,
		},
		{
			name:        "negative limit",
			limit:       -5,
			shouldError: false,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockOData := mocks.NewMockODataOptions(ctrl)
			if tt.shouldError {
				mockOData.EXPECT().SetLimitOption(tt.limit).Return(errors.New("mock error"))
			} else {
				mockOData.EXPECT().SetLimitOption(tt.limit).Return(nil)
			}

			option := convergedclient.WithLimit(tt.limit)
			err := option(mockOData)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWithFilter(t *testing.T) {
	tests := []struct {
		name        string
		filter      string
		shouldError bool
		expectError bool
	}{
		{
			name:        "successful filter option",
			filter:      "name eq 'test'",
			shouldError: false,
			expectError: false,
		},
		{
			name:        "filter option with error",
			filter:      "invalid filter",
			shouldError: true,
			expectError: true,
		},
		{
			name:        "empty filter",
			filter:      "",
			shouldError: false,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockOData := mocks.NewMockODataOptions(ctrl)
			if tt.shouldError {
				mockOData.EXPECT().SetFilterOption(tt.filter).Return(errors.New("mock error"))
			} else {
				mockOData.EXPECT().SetFilterOption(tt.filter).Return(nil)
			}

			option := convergedclient.WithFilter(tt.filter)
			err := option(mockOData)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWithOrderBy(t *testing.T) {
	tests := []struct {
		name        string
		orderBy     string
		shouldError bool
		expectError bool
	}{
		{
			name:        "successful orderBy option",
			orderBy:     "name asc",
			shouldError: false,
			expectError: false,
		},
		{
			name:        "orderBy option with error",
			orderBy:     "invalid order",
			shouldError: true,
			expectError: true,
		},
		{
			name:        "empty orderBy",
			orderBy:     "",
			shouldError: false,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockOData := mocks.NewMockODataOptions(ctrl)
			if tt.shouldError {
				mockOData.EXPECT().SetOrderByOption(tt.orderBy).Return(errors.New("mock error"))
			} else {
				mockOData.EXPECT().SetOrderByOption(tt.orderBy).Return(nil)
			}

			option := convergedclient.WithOrderBy(tt.orderBy)
			err := option(mockOData)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWithExpand(t *testing.T) {
	tests := []struct {
		name        string
		expand      string
		shouldError bool
		expectError bool
	}{
		{
			name:        "successful expand option",
			expand:      "metadata,spec",
			shouldError: false,
			expectError: false,
		},
		{
			name:        "expand option with error",
			expand:      "invalid expand",
			shouldError: true,
			expectError: true,
		},
		{
			name:        "empty expand",
			expand:      "",
			shouldError: false,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockOData := mocks.NewMockODataOptions(ctrl)
			if tt.shouldError {
				mockOData.EXPECT().SetExpandOption(tt.expand).Return(errors.New("mock error"))
			} else {
				mockOData.EXPECT().SetExpandOption(tt.expand).Return(nil)
			}

			option := convergedclient.WithExpand(tt.expand)
			err := option(mockOData)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWithSelect(t *testing.T) {
	tests := []struct {
		name         string
		selectFields string
		shouldError  bool
		expectError  bool
	}{
		{
			name:         "successful select option",
			selectFields: "name,id",
			shouldError:  false,
			expectError:  false,
		},
		{
			name:         "select option with error",
			selectFields: "invalid select",
			shouldError:  true,
			expectError:  true,
		},
		{
			name:         "empty select",
			selectFields: "",
			shouldError:  false,
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockOData := mocks.NewMockODataOptions(ctrl)
			if tt.shouldError {
				mockOData.EXPECT().SetSelectOption(tt.selectFields).Return(errors.New("mock error"))
			} else {
				mockOData.EXPECT().SetSelectOption(tt.selectFields).Return(nil)
			}

			option := convergedclient.WithSelect(tt.selectFields)
			err := option(mockOData)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWithApply(t *testing.T) {
	tests := []struct {
		name        string
		apply       string
		shouldError bool
		expectError bool
	}{
		{
			name:        "successful apply option",
			apply:       "groupby((name))",
			shouldError: false,
			expectError: false,
		},
		{
			name:        "apply option with error",
			apply:       "invalid apply",
			shouldError: true,
			expectError: true,
		},
		{
			name:        "empty apply",
			apply:       "",
			shouldError: false,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockOData := mocks.NewMockODataOptions(ctrl)
			if tt.shouldError {
				mockOData.EXPECT().SetApplyOption(tt.apply).Return(errors.New("mock error"))
			} else {
				mockOData.EXPECT().SetApplyOption(tt.apply).Return(nil)
			}

			option := convergedclient.WithApply(tt.apply)
			err := option(mockOData)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMultipleODataOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOData := mocks.NewMockODataOptions(ctrl)

	// Set up expectations for multiple options
	mockOData.EXPECT().SetPageOption(1).Return(nil)
	mockOData.EXPECT().SetLimitOption(10).Return(nil)
	mockOData.EXPECT().SetFilterOption("name eq 'test'").Return(nil)

	// Apply multiple options
	options := []convergedclient.ODataOption{
		convergedclient.WithPage(1),
		convergedclient.WithLimit(10),
		convergedclient.WithFilter("name eq 'test'"),
	}

	for _, option := range options {
		err := option(mockOData)
		assert.NoError(t, err)
	}
}

// Test Operation interface with generated mocks
func TestOperationInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOperation := mocks.NewMockOperation[TestVM](ctrl)
	ctx := context.Background()
	expectedVMs := []*TestVM{{ExtId: ptr.To("1"), Name: ptr.To("vm1")}, {ExtId: ptr.To("2"), Name: ptr.To("vm2")}}

	// Test Wait
	mockOperation.EXPECT().Wait(ctx).Return(expectedVMs, nil)
	result, err := mockOperation.Wait(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedVMs, result)

	// Test Results
	mockOperation.EXPECT().Results().Return(expectedVMs, nil)
	result, err = mockOperation.Results()
	assert.NoError(t, err)
	assert.Equal(t, expectedVMs, result)

	// Test IsDone
	mockOperation.EXPECT().IsDone().Return(true)
	assert.True(t, mockOperation.IsDone())

	// Test IsSuccess
	mockOperation.EXPECT().IsSuccess().Return(true)
	assert.True(t, mockOperation.IsSuccess())

	// Test IsFailed
	mockOperation.EXPECT().IsFailed().Return(false)
	assert.False(t, mockOperation.IsFailed())

	// Test UUID
	mockOperation.EXPECT().UUID().Return("test-uuid")
	assert.Equal(t, "test-uuid", mockOperation.UUID())

	// Test Status
	mockOperation.EXPECT().Status().Return(convergedclient.TaskStatusSucceeded)
	assert.Equal(t, convergedclient.TaskStatusSucceeded, mockOperation.Status())

	// Test Errors
	mockOperation.EXPECT().Errors().Return([]error{})
	assert.Empty(t, mockOperation.Errors())
}

// Test NoEntityGetter function
func TestNoEntityGetter(t *testing.T) {
	tests := []struct {
		name string
		uuid string
	}{
		{
			name: "valid uuid",
			uuid: "test-uuid-123",
		},
		{
			name: "empty uuid",
			uuid: "",
		},
		{
			name: "special characters in uuid",
			uuid: "test-uuid-!@#$%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := convergedclient.NoEntityGetter(tt.uuid)

			assert.NoError(t, err)
			assert.Nil(t, result)
		})
	}
}
