package converged

import (
	"context"
	"errors"
	"iter"
)

// Custom errors for different states Prism Central asynchronous task.
var (
	ErrTaskNotComplete     = errors.New("task not yet complete")
	ErrTaskResultsNotReady = errors.New("task results not ready")
	ErrTaskFailed          = errors.New("task failed")
)

// Client is the main struct for the converged client.
type Client[
	AntiAffinityPolicy,
	Cluster,
	VirtualGpuProfile,
	PhysicalGpuProfile,
	Host,
	Category,
	Image,
	StorageContainer,
	Subnet,
	VM,
	Task,
	AppMessage,
	VolumeGroup,
	VmAttachment any] struct {
	AntiAffinityPolicies AntiAffinityPolicies[AntiAffinityPolicy]
	Clusters             Clusters[Cluster, VirtualGpuProfile, PhysicalGpuProfile, Host]
	Categories           Categories[Category]
	Images               Images[Image]
	StorageContainers    StorageContainers[StorageContainer]
	Subnets              Subnets[Subnet]
	VMs                  VMs[VM]
	Tasks                Tasks[Task, AppMessage]
	VolumeGroups         VolumeGroups[VolumeGroup, VmAttachment]
	// Additional service interfaces can be added here as needed.
}

// Getter is the interface for Get operations.
type Getter[T any] interface {
	// Get returns the entity for the given UUID.
	Get(ctx context.Context, uuid string) (*T, error)
}

// Lister is the interface for List operations.
type Lister[T any] interface {
	// List returns a list of entities.
	// If no page and limit are provided, the API will return all entities.
	List(ctx context.Context, opts ...ODataOption) ([]T, error)

	// NewIterator returns an iterator for listing entities.
	NewIterator(ctx context.Context, opts ...ODataOption) Iterator[T]
}

// Creator is the interface for Create operations.
type Creator[T any] interface {
	// Create creates a new entity.
	Create(ctx context.Context, entity *T) (*T, error)
}

// AsyncCreator is the interface for Async Create operations.
type AsyncCreator[T any] interface {
	// Create creates a new entity.
	CreateAsync(ctx context.Context, entity *T) (Operation[T], error)
}

// Updater is the interface for Update operations.
type Updater[T any] interface {
	// Update updates an existing entity.
	Update(ctx context.Context, uuid string, entity *T) (*T, error)
}

// AsyncUpdater is the interface for Async Update operations.
type AsyncUpdater[T any] interface {
	// Update updates an existing entity.
	UpdateAsync(ctx context.Context, uuid string, entity *T) (Operation[T], error)
}

// Deleter is the interface for Delete operations.
type Deleter[T any] interface {
	// Delete deletes an existing entity.
	Delete(ctx context.Context, uuid string) error
}

// AsyncDeleter is the interface for Async Delete operations.
type AsyncDeleter[T any] interface {
	// Delete deletes an existing entity.
	DeleteAsync(ctx context.Context, uuid string) (Operation[NoEntity], error)
}

// ODataOption is a functional option for the ODataOptions.
type ODataOption func(params ODataOptions) error

// ODataOptions is the interface for the ODataOptions.
type ODataOptions interface {
	SetPageOption(page int) error
	SetLimitOption(limit int) error
	SetFilterOption(filter string) error
	SetOrderByOption(orderBy string) error
	SetExpandOption(expand string) error
	SetSelectOption(selectFields string) error
	SetApplyOption(apply string) error
}

// WithPage is a functional option for the ODataOptions.
func WithPage(page int) ODataOption {
	return func(params ODataOptions) error {
		return params.SetPageOption(page)
	}
}

// WithLimit is a functional option for the ODataOptions.
func WithLimit(limit int) ODataOption {
	return func(params ODataOptions) error {
		return params.SetLimitOption(limit)
	}
}

// WithFilter is a functional option for the ODataOptions.
func WithFilter(filter string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetFilterOption(filter)
	}
}

// WithOrderBy is a functional option for the ODataOptions.
func WithOrderBy(orderBy string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetOrderByOption(orderBy)
	}
}

// WithExpand is a functional option for the ODataOptions.
func WithExpand(expand string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetExpandOption(expand)
	}
}

// WithSelect is a functional option for the ODataOptions.
func WithSelect(selectFields string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetSelectOption(selectFields)
	}
}

// WithApply is a functional option for the ODataOptions.
func WithApply(apply string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetApplyOption(apply)
	}
}

// TaskStatus is the type for the task status.
type TaskStatus string

const (
	TaskStatusFailed    TaskStatus = "FAILED"
	TaskStatusQueued    TaskStatus = "QUEUED"
	TaskStatusRunning   TaskStatus = "RUNNING"
	TaskStatusSuspended TaskStatus = "SUSPENDED"
	TaskStatusSucceeded TaskStatus = "SUCCEEDED"
	TaskStatusCanceled  TaskStatus = "CANCELED"
	TaskStatusCanceling TaskStatus = "CANCELLING"
	TaskStatusUnknown   TaskStatus = "UNKNOWN"
	TaskStatusRedacted  TaskStatus = "REDACTED"
)

// Operation is the interface for the Prism Central asynchronous task.
type Operation[T any] interface {
	// Blocking wait for task completion
	Wait(ctx context.Context) ([]*T, error)

	// Non-blocking results
	Results() ([]*T, error)                // Returns results if complete, error if not ready
	GetAffectedEntityRefs() ([]any, error) // Returns the affected entity references

	// Non-blocking task states
	IsDone() bool
	IsSuccess() bool
	IsFailed() bool

	// Metadata
	UUID() string
	Status() TaskStatus
	Errors() []error
}

// NoEntity is a placeholder for cases where no entity is returned (e.g. delete operations).
type NoEntity any

// NoEntityGetter is a placeholder for cases where no entity is returned (e.g. delete operations).
func NoEntityGetter(ctx context.Context, uuid string) (*NoEntity, error) {
	return nil, nil
}

// ListIterator is the interface for the list iterator.
type Iterator[T any] iter.Seq2[T, error]
