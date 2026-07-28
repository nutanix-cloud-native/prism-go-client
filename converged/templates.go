package converged

import "context"

// Templates defines the interface for Prism Central templates.
type Templates[Template any] interface {
	// Getter is the interface for Get operations.
	Getter[Template]

	// Lister is the interface for List operations.
	Lister[Template]

	// Creator is the interface for Create operations.
	Creator[Template]

	// InitiateGuestUpdate starts a temporary VM for guest OS update.
	// for a given template UUID.
	InitiateGuestUpdate(ctx context.Context, templateExtID string, versionID string) (*Template, error)

	// InitiateGuestUpdateAsync is the non-blocking variant of InitiateGuestUpdate.
	InitiateGuestUpdateAsync(templateExtID string, versionID string) (Operation[Template], error)

	// CompleteGuestUpdate finalizes an in-progress guest OS update.
	// A new template version is created from the temporary VM state, and
	// the temporary VM is deleted.
	CompleteGuestUpdate(ctx context.Context, templateExtID string, versionName string, versionDescription string, isActiveVersion bool) (*Template, error)

	// CompleteGuestUpdateAsync is the non-blocking variant of CompleteGuestUpdate.
	CompleteGuestUpdateAsync(templateExtID string, versionName string, versionDescription string, isActiveVersion bool) (Operation[Template], error)

	// CancelGuestUpdate aborts an in-progress guest OS update.
	// The temporary VM is deleted and the pending update status is cleared.
	CancelGuestUpdate(ctx context.Context, templateExtID string) error

	// CancelGuestUpdateAsync is the non-blocking variant of CancelGuestUpdate.
	CancelGuestUpdateAsync(templateExtID string) (Operation[NoEntity], error)
}
