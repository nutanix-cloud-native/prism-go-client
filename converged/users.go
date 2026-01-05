package converged

import "context"

// Users defines the interface for IAM Users API operations.
type Users[User any] interface {
	// Getter is the interface for Get operations.
	Getter[User]

	// Lister is the interface for List operations.
	Lister[User]

	// ValidateCredentials validates the provided credentials by making an authenticated API call.
	// Returns an error if credentials are invalid or authentication fails.
	ValidateCredentials(ctx context.Context) error
}
