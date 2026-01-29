// Package converged provides a generic client for Nutanix Prism Central APIs.
// It defines resource interfaces (Clusters, VMs, Images, etc.) and a parameterized
// Client that can be backed by different API versions (e.g. v4) with concrete types.
//
// # Overview
//
// The converged client is a single entry point that exposes sub-clients for each
// resource type: AntiAffinityPolicies, Clusters, Categories, Images, StorageContainers,
// Subnets, VMs, Tasks, VolumeGroups, DomainManager, and Users. Each sub-client
// implements get, list (and often create/update/delete) using a common set of
// interfaces, so callers can work against the converged API without depending on
// a specific Prism Central API version.
//
// # Resource operations
//
// Most resources support:
//   - Get(ctx, uuid) – fetch a single entity by UUID
//   - List(ctx, opts...) – list entities with optional OData parameters (pagination, filter, orderBy, etc.)
//   - NewIterator(ctx, opts...) – stream entities page-by-page without loading all into memory
//
// Some resources also support Create, Update, Delete (sync and/or async). Async operations
// return an Operation[T] that can be waited on with Wait(ctx) or polled via IsDone(), Results(), etc.
//
// # List options (OData)
//
// List and iterator methods accept optional converged.ODataOption values:
//
//	converged.WithPage(0)       // 0-based page index
//	converged.WithLimit(100)   // page size
//	converged.WithFilter("name eq 'my-vm'")
//	converged.WithOrderBy("name asc")
//	converged.WithExpand("resource_domain")
//	converged.WithSelect("name,extId")
//	converged.WithApply("aggregate(...)")
//
// Example – list first page of clusters with a limit:
//
//	clusters, err := client.Clusters.List(ctx, converged.WithPage(0), converged.WithLimit(20))
//
// Example – list all VMs using an iterator (no full in-memory list):
//
//	it := client.VMs.NewIterator(ctx)
//	for vm, err := range it {
//	    if err != nil { return err }
//	    process(vm)
//	}
//
// Example – get a single entity and run an async create:
//
//	cluster, err := client.Clusters.Get(ctx, clusterUUID)
//	op, err := client.VMs.CreateAsync(ctx, vmSpec)
//	if err != nil { return err }
//	created, err := op.Wait(ctx)
//
// To use the converged client with Prism Central V4, get a client via the
// converged/v4 package's ClientCache (see package v4 doc for creating a client
// with cache and usage examples).
package converged

import (
	"context"
	"errors"
	"iter"
)

// Errors for Prism Central asynchronous task states.
var (
	ErrTaskNotComplete     = errors.New("task not yet complete")
	ErrTaskResultsNotReady = errors.New("task results not ready")
	ErrTaskFailed          = errors.New("task failed")
)

// Client is the main entry point for the converged Prism Central client.
// It is parameterized by entity types for each resource (Cluster, VM, Image, etc.)
// and exposes sub-clients for AntiAffinityPolicies, Clusters, Categories, Images,
// StorageContainers, Subnets, VMs, Tasks, VolumeGroups, DomainManager, and Users.
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
	SubnetTaskReference,
	VM,
	Task,
	AppMessage,
	VolumeGroup,
	VmAttachment,
	DomainManagerEntity,
	User any] struct {
	AntiAffinityPolicies AntiAffinityPolicies[AntiAffinityPolicy]
	Clusters             Clusters[Cluster, VirtualGpuProfile, PhysicalGpuProfile, Host]
	Categories           Categories[Category]
	Images               Images[Image]
	StorageContainers    StorageContainers[StorageContainer]
	Subnets              Subnets[Subnet, SubnetTaskReference]
	VMs                  VMs[VM]
	Tasks                Tasks[Task, AppMessage]
	VolumeGroups         VolumeGroups[VolumeGroup, VmAttachment]
	DomainManager        DomainManager[DomainManagerEntity]
	Users                Users[User]
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

// ODataOption is a functional option applied to list/query parameters (e.g. pagination, filter).
type ODataOption func(params ODataOptions) error

// ODataOptions holds OData query parameters for list operations (page, limit, filter, orderBy, etc.).
type ODataOptions interface {
	SetPageOption(page int) error
	SetLimitOption(limit int) error
	SetFilterOption(filter string) error
	SetOrderByOption(orderBy string) error
	SetExpandOption(expand string) error
	SetSelectOption(selectFields string) error
	SetApplyOption(apply string) error
}

// WithPage sets the page number for paginated list results.
func WithPage(page int) ODataOption {
	return func(params ODataOptions) error {
		return params.SetPageOption(page)
	}
}

// WithLimit sets the maximum number of items to return per page.
func WithLimit(limit int) ODataOption {
	return func(params ODataOptions) error {
		return params.SetLimitOption(limit)
	}
}

// WithFilter sets the OData filter expression for list results.
func WithFilter(filter string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetFilterOption(filter)
	}
}

// WithOrderBy sets the OData order-by expression for list results.
func WithOrderBy(orderBy string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetOrderByOption(orderBy)
	}
}

// WithExpand sets the OData expand expression for related entities.
func WithExpand(expand string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetExpandOption(expand)
	}
}

// WithSelect sets the OData select expression for returned fields.
func WithSelect(selectFields string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetSelectOption(selectFields)
	}
}

// WithApply sets the OData apply expression for aggregations or transformations.
func WithApply(apply string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetApplyOption(apply)
	}
}

// TaskStatus represents the state of a Prism Central asynchronous task.
type TaskStatus string

const (
	TaskStatusFailed    TaskStatus = "FAILED"    // Task execution failed.
	TaskStatusQueued    TaskStatus = "QUEUED"    // Task is queued for execution.
	TaskStatusRunning   TaskStatus = "RUNNING"  // Task is currently running.
	TaskStatusSuspended TaskStatus = "SUSPENDED" // Task execution is suspended.
	TaskStatusSucceeded TaskStatus = "SUCCEEDED" // Task completed successfully.
	TaskStatusCanceled  TaskStatus = "CANCELED"  // Task was canceled.
	TaskStatusCanceling TaskStatus = "CANCELLING" // Task cancelation is in progress.
	TaskStatusUnknown   TaskStatus = "UNKNOWN"   // Task state is unknown.
	TaskStatusRedacted  TaskStatus = "REDACTED"  // Task details are redacted.
)

// Operation represents an asynchronous Prism Central task (e.g. create VM, delete image).
// Call Wait to block until completion, or use Results, IsDone, IsSuccess, IsFailed for polling.
type Operation[T any] interface {
	// Wait blocks until the task completes and returns the result entities or an error.
	Wait(ctx context.Context) ([]*T, error)

	// Results returns the result entities if the task succeeded; error if not ready or failed.
	Results() ([]*T, error)
	// GetAffectedEntityRefs returns references to entities affected by the operation.
	GetAffectedEntityRefs() ([]any, error)

	// IsDone reports whether the task has finished (success, failure, or cancel).
	IsDone() bool
	// IsSuccess reports whether the task completed successfully.
	IsSuccess() bool
	// IsFailed reports whether the task failed.
	IsFailed() bool

	// UUID returns the task UUID.
	UUID() string
	// Status returns the current task status.
	Status() TaskStatus
	// Errors returns any errors associated with the task.
	Errors() []error
}

// NoEntity is a type placeholder for operations that return no entity (e.g. delete).
type NoEntity any

// NoEntityGetter is a placeholder getter used when an Operation returns no entity.
func NoEntityGetter(ctx context.Context, uuid string) (*NoEntity, error) {
	return nil, nil
}

// Iterator is a sequence that yields entities of type T or an error (e.g. from List).
type Iterator[T any] iter.Seq2[T, error]
