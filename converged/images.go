package converged

import "context"

// UploadOptions contains optional configuration for image upload.
type UploadOptions struct {
	ProjectUUID string
}

// UploadOption mutates UploadOptions.
type UploadOption func(*UploadOptions)

// WithUploadProjectUUID scopes image upload to the given Prism Central project UUID.
func WithUploadProjectUUID(projectUUID string) UploadOption {
	return func(opts *UploadOptions) {
		if opts == nil {
			return
		}
		opts.ProjectUUID = projectUUID
	}
}

// NewUploadOptions materializes upload options from functional options.
func NewUploadOptions(opts ...UploadOption) UploadOptions {
	resolved := UploadOptions{}
	for _, opt := range opts {
		if opt != nil {
			opt(&resolved)
		}
	}
	return resolved
}

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
	Upload(ctx context.Context, uuid, filepath string, opts ...UploadOption) error
}
