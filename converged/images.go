package converged

import (
	"context"
	"net/http"
)

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

	// Upload uploads the image file to the given UUID. Optional UploadOptions
	// configure the upload — see WithExtraHeaders. Calling Upload with no
	// options preserves the pre-options behaviour exactly.
	Upload(ctx context.Context, uuid, filepath string, opts ...UploadOption) error
}

// UploadOptions holds optional configuration for an Upload call.
type UploadOptions struct {
	// ExtraHeaders are merged onto every HTTP request the Objects Lite
	// upload client makes. Nil or empty means no extra headers.
	ExtraHeaders http.Header
}

// UploadOption configures an Upload call.
type UploadOption func(*UploadOptions)

// WithExtraHeaders adds HTTP headers to every request the Objects Lite S3
// client makes during an Upload. This is intended for deployments where Prism
// Central sits behind a service-token gateway (e.g. Cloudflare Access) that
// requires extra credential headers on every request, including the S3
// PutObject — which the AWS SDK would otherwise not carry.
//
// Headers are applied by an http.RoundTripper wrapper after the AWS SDK
// signs the request, so they do not invalidate the AWS V4 signature. If a
// header is set both here and by the AWS SDK, the value supplied here wins.
//
// If called multiple times, values are merged; duplicate keys accumulate
// across calls.
//
// Calling WithExtraHeaders with a nil or empty map is a no-op.
func WithExtraHeaders(h http.Header) UploadOption {
	return func(o *UploadOptions) {
		if len(h) == 0 {
			return
		}
		if o.ExtraHeaders == nil {
			o.ExtraHeaders = make(http.Header, len(h))
		}
		for k, vs := range h {
			for _, v := range vs {
				o.ExtraHeaders.Add(k, v)
			}
		}
	}
}
