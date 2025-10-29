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
	AntiAffinityPolicy, AntiAffinityPolicyListParams, // AntiaffinityPolicies service dependencies
	Cluster, VirtualGpuProfile, PhysicalGpuProfile, ClusterListParams, GpuProfileListParams, // Clusters service dependencies
	Category, CategoryListParams, // Categories service dependencies
	Image, ImageListParams, // Images service dependencies
	StorageContainer, StorageContainerListParams, // StorageContainers service dependencies
	Subnet, SubnetListParams, // Subnets service dependencies
	VM, VmListParams, // VMs service dependencies
	Task, AppMessage, TaskListParams, // Tasks service dependencies
	VolumeGroup, VmAttachment, VolumeGroupListParams any, // VolumeGroups service dependencies
] struct {
	AntiAffinityPolicies AntiAffinityPolicies[AntiAffinityPolicy, AntiAffinityPolicyListParams]
	Clusters             Clusters[Cluster, VirtualGpuProfile, PhysicalGpuProfile, ClusterListParams, GpuProfileListParams]
	Categories           Categories[Category, CategoryListParams]
	Images               Images[Image, ImageListParams]
	StorageContainers    StorageContainers[StorageContainer, StorageContainerListParams]
	Subnets              Subnets[Subnet, SubnetListParams]
	VMs                  VMs[VM, VmListParams]
	Tasks                Tasks[Task, AppMessage, TaskListParams]
	VolumeGroups         VolumeGroups[VolumeGroup, VmAttachment, VolumeGroupListParams]
}

// Getter is the interface for Get operations.
type Getter[T any] interface {
	// Get returns the entity for the given UUID.
	Get(ctx context.Context, uuid string) (*T, error)
}

// Lister is the interface for List operations.
type Lister[T any, Constraint any] interface {
	// List returns a list of entities.
	// If no page and limit are provided, the API will return all entities.
	List(ctx context.Context, opts ...ODataOption[Constraint]) ([]T, error)

	// NewIterator returns an iterator for listing entities.
	NewIterator(ctx context.Context, opts ...ODataOption[Constraint]) Iterator[T]
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
type ODataOption[Constraint any] func(Constraint) error

type EmptySetter interface{}
type PageSetter interface{ SetPageOption(int) error }
type LimitSetter interface{ SetLimitOption(int) error }
type FilterSetter interface{ SetFilterOption(string) error }
type OrderBySetter interface{ SetOrderByOption(string) error }
type ExpandSetter interface{ SetExpandOption(string) error }
type SelectSetter interface{ SetSelectOption(string) error }
type ApplySetter interface{ SetApplyOption(string) error }

type OptionConvertibleTo[U any] interface {
	~func(U) error
}

type ConvertibleOption[P any] interface {
	OptionConvertibleTo[P]
}

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
func WithPage[T PageSetter](page int) ODataOption[T] {
	return func(setter T) error {
		return setter.SetPageOption(page)
	}
}

// WithLimit is a functional option for the ODataOptions.
func WithLimit[T LimitSetter](limit int) ODataOption[T] {
	return func(setter T) error {
		return setter.SetLimitOption(limit)
	}
}

// WithFilter is a functional option for the ODataOptions.
func WithFilter[T FilterSetter](filter string) ODataOption[T] {
	return func(setter T) error {
		return setter.SetFilterOption(filter)
	}
}

// WithOrderBy is a functional option for the ODataOptions.
func WithOrderBy[T OrderBySetter](orderBy string) ODataOption[T] {
	return func(setter T) error {
		return setter.SetOrderByOption(orderBy)
	}
}

// WithExpand is a functional option for the ODataOptions.
func WithExpand[T ExpandSetter](expand string) ODataOption[T] {
	return func(setter T) error {
		return setter.SetExpandOption(expand)
	}
}

// WithSelect is a functional option for the ODataOptions.
func WithSelect[T SelectSetter](selectFields string) ODataOption[T] {
	return func(setter T) error {
		return setter.SetSelectOption(selectFields)
	}
}

// WithApply is a functional option for the ODataOptions.
func WithApply[T ApplySetter](apply string) ODataOption[T] {
	return func(setter T) error {
		return setter.SetApplyOption(apply)
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
