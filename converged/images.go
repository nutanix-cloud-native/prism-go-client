package converged

// Images defines the interface for Prism Central images.
type Images[Image any] interface {
	// Getter is the interface for Get operations.
	Getter[Image]

	// Lister is the interface for List operations.
	Lister[Image]

	// Creator is the interface for Create operations.
	Creator[Image]

	// Updater is the interface for Update operations.
	Updater[Image]

	// Deleter is the interface for Delete operations.
	Deleter[Image]

	// AsyncCreator is the interface for Async Create operations.
	AsyncCreator[Image]

	// AsyncUpdater is the interface for Async Update operations.
	AsyncUpdater[Image]

	// AsyncDeleter is the interface for Async Delete operations.
	AsyncDeleter[Image]
}
