package converged

import "context"

// Images defines the interface for Prism Central images.
type Images[Image, FileDetail any] interface {
	// Getter is the interface for Get operations.
	Getter[Image]

	// Lister is the interface for List operations.
	Lister[Image]

	// Creator is the interface for Create operations.
	Creator[Image]

	// Deleter is the interface for Delete operations.
	Deleter[Image]

	// GetFile downloads the image file for the given UUID.
	GetFile(ctx context.Context, uuid string) (*FileDetail, error)

	// Upload uploads the image file to the given UUID.
	Upload(ctx context.Context, uuid, filepath string) error
}
