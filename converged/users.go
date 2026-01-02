package converged

import "context"

// Users defines the interface for IAM Users API operations.
type Users[User any] interface {
	// Getter is the interface for Get operations.
	Getter[User]

	// Lister is the interface for List operations.
	Lister[User]

	// GetCurrentLoggedInUser returns the currently logged-in user based on the authenticated credentials.
	// Since V4 IAM API doesn't support a direct "me" endpoint (unlike V3's /users/me),
	// this method internally uses List with limit 1 to retrieve the authenticated user.
	// Returns nil, nil if credentials are valid but no user is found (valid for credential validation).
	GetCurrentLoggedInUser(ctx context.Context) (*User, error)
}
