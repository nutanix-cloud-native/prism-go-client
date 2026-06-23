package converged

// Roles defines the interface for IAM Roles API operations.
type Roles[Role any] interface {
	// Getter is the interface for Get operations.
	Getter[Role]

	// Lister is the interface for List operations.
	Lister[Role]

	// Creator is the interface for Create operations.
	Creator[Role]

	// Deleter is the interface for Delete operations.
	Deleter[Role]
}
