package v4

import (
	"context"
	"fmt"
	"iter"
	"sync"
	"time"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	clusterModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	vmmModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	imageModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
	"k8s.io/utils/ptr"
)

// Client struct for converged client
// It contains implementation for all required API operations grouped by service
type Client struct {
	converged.Client[
		clusterModels.Cluster,
		clusterModels.VirtualGpuProfile,
		clusterModels.PhysicalGpuProfile,
		prismModels.Category,
		imageModels.Image,
		clusterModels.StorageContainer,
		vmmModels.Vm,
	]

	client *v4prismGoClient.Client
}

// NewClient creates a new converged client
// It initializes the V4 client and creates the service implementations
func NewClient(credentials prismgoclient.Credentials, opts ...types.ClientOption[v4prismGoClient.Client]) (*Client, error) {
	v4Client, err := v4prismGoClient.NewV4Client(credentials, opts...)
	if err != nil {
		return nil, err
	}
	client := &Client{
		Client: converged.Client[
			clusterModels.Cluster,
			clusterModels.VirtualGpuProfile,
			clusterModels.PhysicalGpuProfile,
			prismModels.Category,
			imageModels.Image,
			clusterModels.StorageContainer,
			vmmModels.Vm,
		]{
			Images:            NewImagesService(v4Client),
			Clusters:          NewClustersService(v4Client),
			Categories:        NewCategoriesService(v4Client),
			StorageContainers: NewStorageContainersService(v4Client),
			VMs:               NewVMsService(v4Client),
		},
		client: v4Client,
	}
	return client, nil
}

// Operation struct async PC task operation
// It contains the async PC task details and the results
type Operation[T any] struct {
	lock             *sync.Mutex
	entityUUIDs      []string
	entitiesAffected int
	taskUUID         string
	taskStatus       converged.TaskStatus
	taskErrors       []error
	client           *v4prismGoClient.Client
	entityGetter     func(ctx context.Context, uuid string) (*T, error)
	result           []*T
}

// NewOperation creates a new Operation for the given PC async task UUID and client.
func NewOperation[T any](
	taskUUID string,
	client *v4prismGoClient.Client,
	entityGetter func(ctx context.Context, uuid string) (*T, error)) *Operation[T] {
	return &Operation[T]{
		lock:         &sync.Mutex{},
		entityUUIDs:  make([]string, 0),
		taskUUID:     taskUUID,
		taskStatus:   converged.TaskStatusQueued,
		taskErrors:   make([]error, 0),
		client:       client,
		entityGetter: entityGetter,
	}
}

// Wait waits for the PC async task to complete and returns the entities affected.
func (o *Operation[T]) Wait(ctx context.Context) ([]*T, error) {
	var result []*T

	var task prismModels.Task
	var err error

	taskStatusValue := prismModels.TASKSTATUS_UNKNOWN
	taskStatus := &taskStatusValue

	for *taskStatus != prismModels.TASKSTATUS_SUCCEEDED {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("task wait canceled: %w", ctx.Err())
		default:
			// Wait for the task to complete
			time.Sleep(1 * time.Second)

			task, err = CallAPI[*prismModels.GetTaskApiResponse, prismModels.Task](
				o.client.TasksApiInstance.GetTaskById(&o.taskUUID, nil),
			)

			if err != nil {
				return nil, fmt.Errorf("failed to get task %s: %w", o.taskUUID, err)
			}

			taskStatus = task.Status
			if taskStatus == nil {
				return nil, fmt.Errorf("task %s status is nil", o.taskUUID)
			}

			o.setTaskStatus(ConvertTaskStatus(*taskStatus))

			if *taskStatus == prismModels.TASKSTATUS_FAILED {
				errMessages := ""
				if task.ErrorMessages != nil {
					for _, msg := range task.ErrorMessages {
						errMessages += fmt.Sprintf("%s; ", *msg.Message)
					}

					for _, msg := range task.ErrorMessages {
						o.appendTaskError(fmt.Errorf("task %s error: %s", o.taskUUID, *msg.Message))
					}

					return nil, fmt.Errorf("task %s failed: %s", o.taskUUID, errMessages)
				}

				if *taskStatus == prismModels.TASKSTATUS_CANCELED || *taskStatus == prismModels.TASKSTATUS_CANCELING {
					return nil, fmt.Errorf("task %s was canceled", o.taskUUID)
				}
			}
		}

		if task.EntitiesAffected != nil {
			o.setEntitiesAffected(len(task.EntitiesAffected))
		}

		if o.entitiesAffected == 0 {
			return nil, fmt.Errorf("task %s did not affect any entities", o.taskUUID)
		}

		for _, entityRef := range task.EntitiesAffected {
			if entityRef.ExtId == nil {
				return nil, fmt.Errorf("task %s affected entity reference is nil or has no UUID", o.taskUUID)
			}
			o.appendEntityUUID(*entityRef.ExtId)
		}
	}

	for _, uuid := range o.entityUUIDs {
		entity, err := o.entityGetter(ctx, uuid)

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

	// Set results
	o.setResults(result)

	return result, nil
}

// Results returns the results of the PC async task operation.
func (o *Operation[T]) Results() ([]*T, error) {
	if o.taskStatus == converged.TaskStatusFailed || o.taskStatus == converged.TaskStatusCanceled || o.taskStatus == converged.TaskStatusCanceling {
		return nil, converged.ErrTaskFailed
	}

	if o.taskStatus != converged.TaskStatusSucceeded {
		return nil, converged.ErrTaskNotComplete
	}

	if o.result == nil {
		return nil, converged.ErrTaskResultsNotReady
	}

	return o.result, nil
}

// UUID returns the UUID of the PC async task.
func (o *Operation[T]) UUID() string {
	return o.taskUUID
}

// Status returns the status of the PC async task.
func (o *Operation[T]) Status() converged.TaskStatus {
	return o.taskStatus
}

// Errors returns the errors of the PC async task.
func (o *Operation[T]) Errors() []error {
	return o.taskErrors
}

func (o *Operation[T]) setTaskStatus(status converged.TaskStatus) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.taskStatus = status
}

func (o *Operation[T]) setEntitiesAffected(count int) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.entitiesAffected = count
}

func (o *Operation[T]) appendTaskError(err error) {
	o.lock.Lock()
	defer o.lock.Unlock()

	o.taskErrors = append(o.taskErrors, err)
}

func (o *Operation[T]) appendEntityUUID(uuid string) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.entityUUIDs = append(o.entityUUIDs, uuid)
}

func (o *Operation[T]) setResults(results []*T) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.result = results
}

// IsDone checks if the PC async task is finished.
func (o *Operation[T]) IsDone() bool {
	return o.taskStatus == converged.TaskStatusSucceeded ||
		o.taskStatus == converged.TaskStatusFailed ||
		o.taskStatus == converged.TaskStatusCanceled ||
		o.taskStatus == converged.TaskStatusCanceling
}

// IsSuccess checks if the PC async task is successful.
func (o *Operation[T]) IsSuccess() bool {
	return o.taskStatus == converged.TaskStatusSucceeded
}

// IsFailed checks if the PC async task is failed.
func (o *Operation[T]) IsFailed() bool {
	return o.taskStatus == converged.TaskStatusFailed
}

// NewIterator creates a new iterator for the given list function and options.
func NewIterator[R APIResponse, T any](
	ctx context.Context,
	listFunc func(context.Context, *V4ODataParams) (R, error),
	opts ...converged.ODataOption,
) converged.Iterator[T] {
	var zero T

	iterFunc := func(yield func(T, error) bool) {
		var items []T
		var err error

		iterateAllPages := false
		totalCount := 0
		batchIndex := 0
		totalIndex := 0

		reqParams, err := OptsToV4ODataParams(opts...)
		if err != nil {
			return
		}

		if reqParams.Page == nil {
			reqParams.Page = ptr.To(0) // Start from the first page
			reqParams.Limit = nil      // Let API use the default limit
			iterateAllPages = true     // Iterate through all pages
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				// continue with the loop
			}

			// Get next page
			items, totalCount, err = CallListAPI[R, T](listFunc(ctx, reqParams))
			if err != nil {
				yield(zero, err)
				return
			}

			for batchIndex < len(items) && totalIndex < totalCount {
				select {
				case <-ctx.Done():
					return
				default:
					// continue with the loop
				}

				if !yield(items[batchIndex], nil) {
					return
				}

				totalIndex += 1
				batchIndex += 1
			}

			if totalIndex >= totalCount || !iterateAllPages || len(items) == 0 {
				return // All items have been yielded or we are not iterating through all pages
			}

			reqParams.Page = ptr.To(*reqParams.Page + 1) // Move to the next page
			batchIndex = 0
		}
	}

	return converged.Iterator[T](iter.Seq2[T, error](iterFunc))
}
