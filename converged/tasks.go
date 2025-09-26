package converged

import "context"

type Tasks[Task, AppMessage any] interface {
	// Getter is the interface for Get operations.
	Getter[Task]

	// Get task with the given UUID and select fields.
	GetWithSelect(ctx context.Context, uuid string, fields []string) (*Task, error)

	// Lister is the interface for List operations.
	Lister[Task]

	// Cancel task with the given UUID.
	Cancel(ctx context.Context, uuid string) (*AppMessage, error)
}
