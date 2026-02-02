package converged

// Templates defines the interface for Prism Central templates.
type Templates[Template any] interface {
	// Getter is the interface for Get operations.
	Getter[Template]

	// Lister is the interface for List operations.
	Lister[Template]

	// Creator is the interface for Create operations.
	Creator[Template]
}
