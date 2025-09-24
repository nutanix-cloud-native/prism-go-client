package converged

// Images defines the interface for Prism Central images.
type Images[Image any] interface {
	// Getter is the interface for Get operations.
	Getter[Image]

	// Lister is the interface for List operations.
	Lister[Image]
}
