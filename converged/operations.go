package converged

// Operations defines the interface for IAM Operation API.
type Operations[Operation any] interface {
	// Getter is the interface for Get operations.
	Getter[Operation]

	// Lister is the interface for List operations.
	Lister[Operation]
}
