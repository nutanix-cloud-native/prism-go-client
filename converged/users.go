package converged

import "context"

// Users defines the interface for IAM Users API operations.
type Users[User any] interface {
	// Getter is the interface for Get operations.
	Getter[User]

	// Lister is the interface for List operations.
	Lister[User]

	// Creator is the interface for Create operations.
	Creator[User]

	// Updater is the interface for Update operations.
	Updater[User]

	// UpdateUserState updates the state of a user (activate/deactivate).
	UpdateUserState(ctx context.Context, uuid string, stateUpdate any) (*User, error)

	// ChangeUserPassword changes the password for a user.
	ChangeUserPassword(ctx context.Context, passwordChange any) error

	// ResetUserPassword resets the password for a user.
	ResetUserPassword(ctx context.Context, uuid string, passwordReset any) error
}
