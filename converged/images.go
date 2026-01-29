package converged

// Images defines the interface for Prism Central images.
type Images[Image any] interface {
	// Getter is the interface for Get operations.
	Getter[Image]

	// Lister is the interface for List operations.
	Lister[Image]

	// Creator is the interface for Create operations.
	Creator[Image]

	// Deleter is the interface for Delete operations.
	Deleter[Image]

	// AsyncCreator is the interface for Async Create operations.
	AsyncCreator[Image]

	// AsyncDeleter is the interface for Async Delete operations.
	AsyncDeleter[Image]
}
