package facade

type FacadeClientV4 interface {
	AntiAffinityPolicyFacadeV4
	CategoriesFacadeV4
	VMsFacadeV4
	ClustersFacadeV4
	ImagesFacadeV4
	StorageContainersFacadeV4
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

type ODataListIterator[T any] interface {
	Next() bool
	GetCurrent() (T, error)
	Count() int
}
