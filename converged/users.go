package converged

// Users defines the interface for IAM Users API operations.
type Users[User any] interface {
	// Getter is the interface for Get operations.
	Getter[User]

	// Lister is the interface for List operations.
	Lister[User]
}
