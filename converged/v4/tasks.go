package v4

import (
	"context"
	"fmt"
	"strings"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	prismErrors "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/error"
)

// TasksService is the service for the tasks
type TasksService struct {
	client *v4prismGoClient.Client

	entityName string
}

// NewTasksService creates a new TasksService for the converged client
func NewTasksService(client *v4prismGoClient.Client) *TasksService {
	return &TasksService{
		client:     client,
		entityName: "task",
	}
}

// Get returns the task for the given UUID
func (s *TasksService) Get(ctx context.Context, uuid string) (*prismModels.Task, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	return GenericGetEntity[*prismModels.GetTaskApiResponse, prismModels.Task](
		func() (*prismModels.GetTaskApiResponse, error) {
			return s.client.TasksApiInstance.GetTaskById(&uuid, nil, nil)
		},
		s.entityName,
	)
}

// GetWithSelect returns the task for the given UUID and select fields
func (s *TasksService) GetWithSelect(ctx context.Context, uuid string, fields []string) (*prismModels.Task, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	selectOption := strings.Join(fields, ",")

	return GenericGetEntity[*prismModels.GetTaskApiResponse, prismModels.Task](
		func() (*prismModels.GetTaskApiResponse, error) {
			return s.client.TasksApiInstance.GetTaskById(&uuid, &selectOption, nil)
		},
		s.entityName,
	)
}

// List returns the list of tasks
func (s *TasksService) List(ctx context.Context, opts ...converged.ODataOption) ([]prismModels.Task, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for tasks")
	}

	return GenericListEntities[*prismModels.ListTasksApiResponse, prismModels.Task](
		func(reqParams *V4ODataParams) (*prismModels.ListTasksApiResponse, error) {
			return s.client.TasksApiInstance.ListTasks(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		s.entityName,
	)
}

// NewIterator returns a new iterator for the tasks
func (s *TasksService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[prismModels.Task] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*prismModels.ListTasksApiResponse, prismModels.Task](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*prismModels.ListTasksApiResponse, error) {
			return s.client.TasksApiInstance.ListTasks(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		s.entityName,
	)
}

// Cancel cancels the task for the given UUID
func (s *TasksService) Cancel(ctx context.Context, uuid string) (*prismErrors.AppMessage, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	return GenericGetEntity[*prismModels.CancelTaskApiResponse, prismErrors.AppMessage](
		func() (*prismModels.CancelTaskApiResponse, error) {
			return s.client.TasksApiInstance.CancelTask(&uuid)
		},
		s.entityName,
	)
}
