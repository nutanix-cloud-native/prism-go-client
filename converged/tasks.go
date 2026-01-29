package converged

import "context"

// Tasks defines the interface for Prism Central task operations.
// Tasks represent asynchronous operations (e.g. VM create, image import).
type Tasks[Task, AppMessage any] interface {
	Getter[Task]

	// GetWithSelect returns the task with the given UUID and selected fields.
	GetWithSelect(ctx context.Context, uuid string, fields []string) (*Task, error)

	Lister[Task]

	// Cancel cancels the task with the given UUID.
	Cancel(ctx context.Context, uuid string) (*AppMessage, error)
}
