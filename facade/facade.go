package facade

type FacadeClientV4 interface {
	AntiAffinityPolicyFacadeV4
	CategoriesFacadeV4
}

type ODataOption func(params ODataOptions) error

type ODataOptions interface {
	SetPageOption(page int) ODataOption
	SetLimitOption(limit int) ODataOption
	SetFilterOption(filter string) ODataOption
	SetOrderByOption(orderBy string) ODataOption
	SetExpandOption(expand string) ODataOption
	SetSelectOption(selectFields string) ODataOption
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
