package v4

import (
	"fmt"
	"sync"
	"time"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	"github.com/nutanix-cloud-native/prism-go-client/facade"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	v4prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

type FacadeV4Client struct {
	client *v4prismGoClient.Client
}

// NewFacadeV4Client creates a new FacadeV4Client with the provided credentials and options.
func NewFacadeV4Client(credentials prismgoclient.Credentials, opts ...types.ClientOption[v4prismGoClient.Client]) (*FacadeV4Client, error) {
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

	var task v4prismModels.Task
	var err error

	taskStatusValue := v4prismModels.TASKSTATUS_UNKNOWN
	taskStatus := &taskStatusValue

	for *taskStatus != v4prismModels.TASKSTATUS_SUCCEEDED {
		time.Sleep(1 * time.Second)

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
	}

	for _, uuid := range f.entityUUIDs {
		entity, err := f.entityGetter(uuid)

		// TODO: Get rid of this check when we will be sure that the correct amount of entities is returned (see comment below)
		if entity != nil {
			if err != nil {
				return nil, fmt.Errorf("failed to get entity %s: %w", uuid, err)
			}
			// TODO: Uncomment this when we will be sure that the correct amount of entities is returned
			// STATE: 2025-07-29 - Ilya Alekseyev - On VM creation sometimes returns 2 entities one if which can not be found.
			// if entity == nil {
			// 	return nil, fmt.Errorf("entity %s not found", uuid)
			// }

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

type FacadeV4ODataIterator[R APIResponse, T any] struct {
	client             *v4prismGoClient.Client
	opts               []facade.ODataOption
	totalCount         int
	currentPage        int
	currentIndex       int
	currentBufferIndex int
	itemsBuffer        []T
	iteratorError      error
	listFunc           func(*V4ODataParams) (R, error)
	mutex              sync.Mutex
}

func NewFacadeV4ODataIterator[R APIResponse, T any](
	client *v4prismGoClient.Client,
	totalCount int,
	items []T,
	listFunc func(*V4ODataParams) (R, error),
	opts ...facade.ODataOption,
) *FacadeV4ODataIterator[R, T] {
	return &FacadeV4ODataIterator[R, T]{
		client:       client,
		opts:         opts,
		currentPage:  0,
		itemsBuffer:  items,
		currentIndex: 0,
		listFunc:     listFunc,
		totalCount:   totalCount,
		mutex:        sync.Mutex{},
	}
}

func (it *FacadeV4ODataIterator[R, T]) Next() bool {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	if it.currentIndex >= it.totalCount {
		return false
	}

	it.currentIndex++
	it.currentBufferIndex++

	if it.currentBufferIndex < len(it.itemsBuffer) {
		return true
	}

	reqParams, err := OptsToV4ODataParams(it.opts...)
	if err != nil {
		it.iteratorError = fmt.Errorf("failed to convert options to V4 OData params: %w", err)
		return false
	}
	reqParams.Limit = nil // Let API use the default limit
	reqParams.Page = &it.currentPage

	items, totalCount, err := CallListAPI[R, T](
		it.listFunc(reqParams),
	)
	if err != nil {
		it.currentBufferIndex-- // Reset buffer index since we failed to fetch new items
		it.currentIndex--       // Reset current index since we failed to fetch new items
		it.iteratorError = fmt.Errorf("failed to list items: %w", err)
		return false
	}

	it.currentBufferIndex = 0
	it.currentPage++

	it.totalCount = totalCount
	it.itemsBuffer = items

	return len(it.itemsBuffer) > 0 && it.currentIndex < it.totalCount
}

func (it *FacadeV4ODataIterator[R, T]) GetCurrent() (T, error) {
	if it.iteratorError != nil {
		var zero T
		return zero, it.iteratorError
	}

	if it.currentBufferIndex < 0 || it.currentBufferIndex >= len(it.itemsBuffer) {
		var zero T
		return zero, fmt.Errorf("current buffer index %d is out of bounds for items buffer of length %d", it.currentBufferIndex, len(it.itemsBuffer))
	}

	return it.itemsBuffer[it.currentBufferIndex], nil
}

func (it *FacadeV4ODataIterator[R, T]) Count() int {
	return it.totalCount
}
