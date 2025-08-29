package ferrors

import (
	"fmt"
	"runtime"
)

type ErrorType string

const (
	ErrorTypeUncategorisedError ErrorType = "UNCATEGORISED_ERROR"
	ErrorTypeV4ApiError         ErrorType = "V4_API_ERROR"
	ErrorTypeTypeAssertion      ErrorType = "TYPE_ASSERTION_ERROR"
	ErrorTypeNetworkError       ErrorType = "NETWORK_ERROR"
	ErrorTypeTimeoutError       ErrorType = "TIMEOUT_ERROR"
)

type ErrorSubType string

type ErrorSubTypeV4Api ErrorSubType

const (
	ErrorSubTypeV4ApiUncategorisedError    ErrorSubTypeV4Api = "UNCATEGORISED_ERROR"
	ErrorSubTypeV4ApiAuthorizationError    ErrorSubTypeV4Api = "AUTHORIZATION_ERROR"
	ErrorSubTypeV4ApiResourceNotFoundError ErrorSubTypeV4Api = "RESOURCE_NOT_FOUND"
	ErrorSubTypeV4ApiRateLimitError        ErrorSubTypeV4Api = "RATE_LIMIT_ERROR"
	ErrorSubTypeV4ApiSchemaValidationError ErrorSubTypeV4Api = "SCHEMA_VALIDATION_ERROR"
)

type Err struct {
	Message    string                   `json:"-"`
	Type       ErrorType                `json:"type"`
	SubType    ErrorSubType             `json:"sub_type,omitempty"`
	Err        interface{}              `json:"error"`
	StackTrace []string                 `json:"-"`
	Args       []map[string]interface{} `json:"-"`
}

func (fe *Err) Error() string {
	return fmt.Sprintf("ErrorType: %v, Message: %s, Detail: %v", fe.Type, fe.Message, fe.Err)
}

func (fe *Err) GetErrorType() ErrorType {
	return fe.Type
}

func (fe *Err) GetError() interface{} {
	return fe.Err
}

func new(errType ErrorType, errSubType ErrorSubType, msg string, err interface{}, args ...map[string]interface{}) error {
	pc := make([]uintptr, 10)
	count := runtime.Callers(1, pc)
	frames := runtime.CallersFrames(pc[:count])

	var stacktrace []string
	for frame, hasMore := frames.Next(); hasMore; frame, hasMore = frames.Next() {
		stacktrace = append(stacktrace, fmt.Sprintf("%s:%d", frame.File, frame.Line))
	}

	return &Err{
		Type:       errType,
		SubType:    errSubType,
		Message:    msg,
		Err:        err,
		StackTrace: stacktrace,
		Args:       args,
	}
}

func NewErrUncategorisedError(msg string, err interface{}, args ...map[string]interface{}) error {
	return new(ErrorTypeUncategorisedError, "", msg, err, args...)
}

func NewErrTypeAssertionError(msg string, err interface{}, args ...map[string]interface{}) error {
	return new(ErrorTypeTypeAssertion, "", msg, err, args...)
}

func NewErrNetworkError(msg string, err interface{}, args ...map[string]interface{}) error {
	return new(ErrorTypeNetworkError, "", msg, err, args...)
}

func NewErrTimeoutError(msg string, err interface{}, args ...map[string]interface{}) error {
	return new(ErrorTypeNetworkError, "", msg, err, args...)
}

func NewErrV4ApiUncategorisedError(msg string, err interface{}, args ...map[string]interface{}) error {
	return new(ErrorTypeV4ApiError, ErrorSubType(ErrorSubTypeV4ApiUncategorisedError), msg, err, args...)
}

func NewErrV4ApiAuthorizationError(msg string, err interface{}, args ...map[string]interface{}) error {
	return new(ErrorTypeV4ApiError, ErrorSubType(ErrorSubTypeV4ApiAuthorizationError), msg, err, args...)
}

func NewErrV4ApiResourceNotFoundError(msg string, err interface{}, args ...map[string]interface{}) error {
	return new(ErrorTypeV4ApiError, ErrorSubType(ErrorSubTypeV4ApiResourceNotFoundError), msg, err, args...)
}

func NewErrV4ApiRateLimitError(msg string, err interface{}, args ...map[string]interface{}) error {
	return new(ErrorTypeV4ApiError, ErrorSubType(ErrorSubTypeV4ApiRateLimitError), msg, err, args...)
}

func NewErrV4ApiSchemaValidationError(msg string, err interface{}, args ...map[string]interface{}) error {
	return new(ErrorTypeV4ApiError, ErrorSubType(ErrorSubTypeV4ApiSchemaValidationError), msg, err, args...)
}
