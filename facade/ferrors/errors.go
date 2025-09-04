package ferrors

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrSchemaValidationError = errors.New("schema validation error")
	ErrAuthorizationError    = errors.New("authorization error")
	ErrResourceNotFound      = errors.New("resource not found")
	ErrRateLimitExceeded     = errors.New("rate limit exceeded")
	ErrInternalServiceError  = errors.New("internal service error")
	ErrInvalidInput          = errors.New("invalid input error")
)

type ApiError struct {
	Type          error `json:"type"`
	OriginalError error `json:"original_error"`
}

type UnexpectedTypeError struct {
	ExpectedType reflect.Type `json:"expected_type"`
	ActualType   reflect.Type `json:"actual_type"`
}

func (ae *ApiError) Error() string {
	errStr := fmt.Sprintf("api error [%v]", ae.Type.Error())
	if ae.OriginalError != nil {
		errStr = fmt.Sprintf("%s: %v", errStr, ae.OriginalError.Error())
	}

	return errStr
}

func NewApiError(errorType error, originalError error) error {
	return &ApiError{
		Type:          errorType,
		OriginalError: originalError,
	}
}

func NewUnexpectedTypeError(expectedValue any, actualValue any) error {
	return &UnexpectedTypeError{
		ExpectedType: reflect.TypeOf(expectedValue),
		ActualType:   reflect.TypeOf(actualValue),
	}
}

func (ute *UnexpectedTypeError) Error() string {
	return fmt.Sprintf("unexpected type: expected %v, got %v", ute.ExpectedType, ute.ActualType)
}
