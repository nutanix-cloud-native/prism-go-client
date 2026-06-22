package converged

import "context"

// Alerts defines the interface for managing alerts.
type Alerts[Alert any] interface {
	// Getter is the interface for Get operations.
	Getter[Alert]

	// Lister is the interface for List operations.
	Lister[Alert]

	// Acknowledge acknowledges the alert identified by the given external
	// identifier and returns the asynchronous task operation.
	Acknowledge(ctx context.Context, uuid string) (Operation[Alert], error)

	// Resolve resolves the alert identified by the given external identifier
	// and returns the asynchronous task operation.
	Resolve(ctx context.Context, uuid string) (Operation[Alert], error)
}
