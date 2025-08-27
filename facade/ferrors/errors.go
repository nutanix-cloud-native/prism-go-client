package ferrors

import (
	"fmt"
	"runtime"
)

type ErrorType string

const (
	ErrorTypeUnknown       ErrorType = "UNKNOWN_ERROR"
	ErrorTypeApiError      ErrorType = "V4_API_ERROR"
	ErrorTypeTypeAssertion ErrorType = "TYPE_ASSERTION_ERROR"
)

type Err struct {
	Message    string                   `json:"-"`
	Type       ErrorType                `json:"type"`
	Err        interface{}              `json:"error"`
	StackTrace []string                 `json:"-"`
	Args       []map[string]interface{} `json:"args,omitempty"`
}

func (fe Err) Error() string {
	return fmt.Sprintf("ErrorType: %v, Message: %s, Detail: %v", fe.Type, fe.Message, fe.Err)
}

func (fe Err) GetErrorType() ErrorType {
	return fe.Type
}

func (fe Err) GetError() interface{} {
	return fe.Err
}

func New(errType ErrorType, msg string, err interface{}, args ...map[string]interface{}) error {
	pc := make([]uintptr, 10)
	count := runtime.Callers(1, pc)
	frames := runtime.CallersFrames(pc[:count])

	var stacktrace []string
	for frame, hasMore := frames.Next(); hasMore; frame, hasMore = frames.Next() {
		stacktrace = append(stacktrace, fmt.Sprintf("%s:%d", frame.File, frame.Line))
	}

	return Err{
		Type:       errType,
		Message:    msg,
		Err:        err,
		StackTrace: stacktrace,
		Args:       args,
	}
}

func NewErrUnknownError(msg string, err interface{}, args ...map[string]interface{}) error {
	return New(ErrorTypeUnknown, msg, err, args...)
}

func NewErrTypeAssertionError(msg string, err interface{}, args ...map[string]interface{}) error {
	return New(ErrorTypeTypeAssertion, msg, err, args...)
}

func NewErrApiError(msg string, err interface{}, args ...map[string]interface{}) error {
	return New(ErrorTypeApiError, msg, err, args...)
}
