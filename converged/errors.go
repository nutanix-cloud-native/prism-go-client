package converged

import (
	"errors"
	"strings"
)

var (
	ErrNotFound  = errors.New("resource not found")
	ErrRateLimit = errors.New("rate limit exceeded")
	ErrInternal  = errors.New("internal error")
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
