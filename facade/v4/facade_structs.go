package v4

import (
	"fmt"
	"sync"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/facade"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	v4prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

type FacadeV4Client struct {
	client *v4prismGoClient.Client
}

// NewFacadeV4Client creates a new FacadeV4Client with the provided credentials and options.
func NewFacadeV4Client(credentials prismgoclient.Credentials, opts ...v4prismGoClient.ClientOption) (facade.FacadeClientV4, error) {
	client, err := v4prismGoClient.NewV4Client(credentials, opts...)
	if err != nil {
		return nil, err
	}

	result := &FacadeV4Client{
		client: client,
	}

	return result, nil
}

type FacadeV4TaskWaiter[T any] struct {
	lock             *sync.Mutex
	entityUUIDs      []string
	entitiesAffected int
	taskUUID         string
	taskStatus       facade.TaskStatus
	taskErrors       []error
	client           *v4prismGoClient.Client
	entityGetter     func(uuid string) (*T, error)
}

// NewFacadeV4TaskWaiter creates a new FacadeV4TaskWaiter for the given task UUID and client.
func NewFacadeV4TaskWaiter[T any](taskUUID string, client *v4prismGoClient.Client, entityGetter func(uuid string) (*T, error)) *FacadeV4TaskWaiter[T] {
	return &FacadeV4TaskWaiter[T]{
		lock:         &sync.Mutex{},
		entityUUIDs:  make([]string, 0),
		taskUUID:     taskUUID,
		taskStatus:   facade.TaskStatusQueued,
		taskErrors:   make([]error, 0),
		client:       client,
		entityGetter: entityGetter,
	}
}

func (f *FacadeV4TaskWaiter[T]) WaitForTaskCompletion() ([]*T, error) {
	var result []*T

	task, err := CallAPI[*v4prismModels.GetTaskApiResponse, v4prismModels.Task](
		f.client.TasksApiInstance.GetTaskById(&f.taskUUID, nil),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get task %s: %w", f.taskUUID, err)
	}

	taskStatus := task.Status
	if taskStatus == nil {
		return nil, fmt.Errorf("task %s status is nil", f.taskUUID)
	}

	f.setTaskStatus(ConvertTaskStatus(*taskStatus))

	for *taskStatus != v4prismModels.TASKSTATUS_SUCCEEDED {
		// Wait for the task to complete
		task, err = CallAPI[*v4prismModels.GetTaskApiResponse, v4prismModels.Task](
			f.client.TasksApiInstance.GetTaskById(&f.taskUUID, nil),
		)

		if err != nil {
			return nil, fmt.Errorf("failed to get task %s: %w", f.taskUUID, err)
		}

		taskStatus = task.Status
		if taskStatus == nil {
			return nil, fmt.Errorf("task %s status is nil", f.taskUUID)
		}

		f.setTaskStatus(ConvertTaskStatus(*taskStatus))

		if *taskStatus == v4prismModels.TASKSTATUS_FAILED {
			errMessages := ""
			if task.ErrorMessages != nil {
				for _, msg := range task.ErrorMessages {
					errMessages += fmt.Sprintf("%s; ", *msg.Message)
				}

				for _, msg := range task.ErrorMessages {
					f.appendTaskError(fmt.Errorf("task %s error: %s", f.taskUUID, *msg.Message))
				}

				return nil, fmt.Errorf("task %s failed: %s", f.taskUUID, errMessages)
			}

			if *taskStatus == v4prismModels.TASKSTATUS_CANCELED || *taskStatus == v4prismModels.TASKSTATUS_CANCELING {
				return nil, fmt.Errorf("task %s was canceled", f.taskUUID)
			}
		}

		if task.EntitiesAffected != nil {
			f.setEntitiesAffected(len(task.EntitiesAffected))
		}

		if f.entitiesAffected == 0 {
			return nil, fmt.Errorf("task %s did not affect any entities", f.taskUUID)
		}

		for _, entityRef := range task.EntitiesAffected {
			if entityRef.ExtId == nil {
				return nil, fmt.Errorf("task %s affected entity reference is nil or has no UUID", f.taskUUID)
			}
			f.appendEntityUUID(*entityRef.ExtId)
		}

		for _, uuid := range f.entityUUIDs {
			entity, err := f.entityGetter(uuid)
			if err != nil {
				return nil, fmt.Errorf("failed to get entity %s: %w", uuid, err)
			}
			if entity == nil {
				return nil, fmt.Errorf("entity %s not found", uuid)
			}
			result = append(result, entity)
		}

	}
	return result, nil
}

func (f *FacadeV4TaskWaiter[T]) GetTaskUUID() string {
	return f.taskUUID
}

func (f *FacadeV4TaskWaiter[T]) GetTaskStatus() facade.TaskStatus {
	return f.taskStatus
}

func (f *FacadeV4TaskWaiter[T]) GetTaskErrors() []error {
	return f.taskErrors
}

func (f *FacadeV4TaskWaiter[T]) setTaskStatus(status facade.TaskStatus) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.taskStatus = status
}

func (f *FacadeV4TaskWaiter[T]) setEntitiesAffected(count int) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.entitiesAffected = count
}

func (f *FacadeV4TaskWaiter[T]) appendTaskError(err error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	if f.taskErrors == nil {
		f.taskErrors = make([]error, 0)
	}
	f.taskErrors = append(f.taskErrors, err)
}

func (f *FacadeV4TaskWaiter[T]) appendEntityUUID(uuid string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.entityUUIDs = append(f.entityUUIDs, uuid)
}
