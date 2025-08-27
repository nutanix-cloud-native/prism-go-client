package facade

import "iter"

type FacadeClientV4 interface {
	AntiAffinityPolicyFacadeV4
	CategoriesFacadeV4
	VMsFacadeV4
	ClustersFacadeV4
	ImagesFacadeV4
	StorageContainersFacadeV4
	SubnetsFacadeV4
	// Additional facade interfaces can be added here as needed.
}

type ODataOption func(params ODataOptions) error

type ODataOptions interface {
	SetPageOption(page int) error
	SetLimitOption(limit int) error
	SetFilterOption(filter string) error
	SetOrderByOption(orderBy string) error
	SetExpandOption(expand string) error
	SetSelectOption(selectFields string) error
	SetApplyOption(apply string) error
}

func WithPage(page int) ODataOption {
	return func(params ODataOptions) error {
		return params.SetPageOption(page)
	}
}

func WithLimit(limit int) ODataOption {
	return func(params ODataOptions) error {
		return params.SetLimitOption(limit)
	}
}

func WithFilter(filter string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetFilterOption(filter)
	}
}

func WithOrderBy(orderBy string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetOrderByOption(orderBy)
	}
}

func WithExpand(expand string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetExpandOption(expand)
	}
}

func WithSelect(selectFields string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetSelectOption(selectFields)
	}
}

func WithApply(apply string) ODataOption {
	return func(params ODataOptions) error {
		return params.SetApplyOption(apply)
	}
}

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

type TaskWaiter[T any] interface {
	WaitForTaskCompletion() ([]*T, error)
	GetTaskUUID() string
	GetTaskStatus() TaskStatus
	GetTaskErrors() []error
}

// NoEntity is a placeholder for cases where no entity is returned (e.g. delete operations).
type NoEntity interface{}

func NoEntityGetter(uuid string) (*NoEntity, error) {
	return nil, nil
}

// make it more Go-ish style iterator
type ODataListIterator[T any] iter.Seq2[T, error]

//
// 1. Error handling
// 2. Go iterator
// 3. Naming
// 4. Context in every API call
//
