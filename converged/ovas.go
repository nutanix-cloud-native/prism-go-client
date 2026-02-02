package converged

import "context"

// Ovas defines the interface for Prism Central OVAs.
type Ovas[Ova, FileDetail any] interface {
	// Getter is the interface for Get operations.
	Getter[Ova]

	// Lister is the interface for List operations.
	Lister[Ova]

	// Creator is the interface for Create operations.
	Creator[Ova]

	// GetFile downloads the OVA file for the given UUID.
	GetFile(ctx context.Context, uuid string) (*FileDetail, error)
}
