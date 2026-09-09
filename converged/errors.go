package converged

import (
	"errors"
	"strings"
)

var (
	ErrNotFound        = errors.New("resource not found")
	ErrRateLimit       = errors.New("rate limit exceeded")
	ErrInternal        = errors.New("internal error")
	ErrStaleClient     = errors.New("stale client")
	ErrUnauthenticated = errors.New("unauthenticated") // HTTP 401: missing/invalid credentials
	ErrUnauthorized    = errors.New("unauthorized")    // HTTP 403: authenticated but not allowed
)

// APIError wraps SDK errors with a classified Kind.
// Kind is one of the sentinel errors above, or nil for unclassified errors.
// Cause is the original error (e.g. GenericOpenAPIError from the SDK).
type APIError struct {
	Kind    error
	Cause   error
	Message string
}

func (e *APIError) Error() string {
	parts := make([]string, 0, 3)
	if e.Kind != nil {
		parts = append(parts, e.Kind.Error())
	}
	if e.Message != "" {
		parts = append(parts, e.Message)
	}
	if e.Cause != nil {
		parts = append(parts, e.Cause.Error())
	}
	if len(parts) == 0 {
		return "unknown api error"
	}
	return strings.Join(parts, ": ")
}

func (e *APIError) Unwrap() error { return e.Cause }

func (e *APIError) Is(target error) bool {
	return e.Kind != nil && errors.Is(e.Kind, target)
}

// IsNotFound returns true if the error is or wraps an ErrNotFound.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// IsRateLimit returns true if the error is or wraps an ErrRateLimit.
func IsRateLimit(err error) bool {
	return errors.Is(err, ErrRateLimit)
}

// IsInternal returns true if the error is or wraps an ErrInternal.
func IsInternal(err error) bool {
	return errors.Is(err, ErrInternal)
}

// IsStaleClient returns true if the error is or wraps an ErrStaleClient.
func IsStaleClient(err error) bool {
	return errors.Is(err, ErrStaleClient) || (err != nil && strings.Contains(err.Error(), "unsupported protocol scheme"))
}

// IsUnauthenticated returns true if the error is or wraps an ErrUnauthenticated
// (HTTP 401 — missing or invalid credentials), including stale-client auth
// failures that were observed after a 401/OIDC redirect.
func IsUnauthenticated(err error) bool {
	return errors.Is(err, ErrUnauthenticated) || IsStaleClient(err)
}

// IsUnauthorized returns true if the error is or wraps an ErrUnauthorized
// (HTTP 403 — authenticated but not permitted).
func IsUnauthorized(err error) bool {
	return errors.Is(err, ErrUnauthorized)
}
