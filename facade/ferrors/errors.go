package ferrors

import (
	"fmt"
)

type ErrorType string

const (
	ErrorTypeV4ApiError    ErrorType = "V4_API_ERROR"
	ErrorTypeTypeAssertion ErrorType = "TYPE_ASSERTION_ERROR"
	ErrorTypeNetworkError  ErrorType = "NETWORK_ERROR"
	ErrorTypeInternalError ErrorType = "INTERNAL_ERROR"
)

type ErrorSubType string

type ErrorSubTypeV4Api ErrorSubType

const (
	ErrorSubTypeV4ApiUncategorisedError    ErrorSubTypeV4Api = "UNCATEGORISED_ERROR"
	ErrorSubTypeV4ApiSchemaValidationError ErrorSubTypeV4Api = "SCHEMA_VALIDATION_ERROR"
	ErrorSubTypeV4ApiAuthorizationError    ErrorSubTypeV4Api = "AUTHORIZATION_ERROR"
	ErrorSubTypeV4ApiResourceNotFoundError ErrorSubTypeV4Api = "RESOURCE_NOT_FOUND_ERROR"
	ErrorSubTypeV4ApiRateLimitError        ErrorSubTypeV4Api = "RATE_LIMIT_ERROR"
	ErrorSubTypeV4ApiInternalServiceError  ErrorSubTypeV4Api = "INTERNAL_SERVICE_ERROR"
	ErrorSubTypeV4ApiInvalidInputError     ErrorSubTypeV4Api = "INVALID_INPUT_ERROR"
)

var (
	V4ApiErrorSubTypeToErrConstructor = map[ErrorSubTypeV4Api]func(string, any) error{
		ErrorSubTypeV4ApiUncategorisedError:    NewErrV4ApiUncategorisedError,
		ErrorSubTypeV4ApiSchemaValidationError: NewErrV4ApiSchemaValidationError,
		ErrorSubTypeV4ApiAuthorizationError:    NewErrV4ApiAuthorizationError,
		ErrorSubTypeV4ApiResourceNotFoundError: NewErrV4ApiResourceNotFoundError,
		ErrorSubTypeV4ApiRateLimitError:        NewErrV4ApiRateLimitError,
		ErrorSubTypeV4ApiInternalServiceError:  NewErrV4ApiInternalError,
		ErrorSubTypeV4ApiInvalidInputError:     NewErrV4ApiInvalidInputError,
	}
)

type FacadeError struct {
	Message string       `json:"-"`
	Type    ErrorType    `json:"type"`
	SubType ErrorSubType `json:"sub_type,omitempty"`
	Data    any          `json:"data"`
}

func (fe *FacadeError) Error() string {
	errStr := fmt.Sprintf("ErrorType: %v", fe.Type)
	if fe.SubType != "" {
		errStr += fmt.Sprintf(", ErrorSubType: %v", fe.SubType)
	}
	if fe.Message != "" {
		errStr += fmt.Sprintf(", Message: %s", fe.Message)
	}
	errStr += fmt.Sprintf(", Detail: %v", fe.Data)

	return errStr
}

func new(errType ErrorType, errSubType ErrorSubType, msg string, err any) error {
	return &FacadeError{
		Type:    errType,
		SubType: errSubType,
		Message: msg,
		Data:    err,
	}
}

func NewErrTypeAssertionError(msg string, err any) error {
	return new(ErrorTypeTypeAssertion, "", msg, err)
}

func NewErrNetworkError(msg string, err any) error {
	return new(ErrorTypeNetworkError, "", msg, err)
}

func NewErrInternalError(msg string, err any) error {
	return new(ErrorTypeInternalError, "", msg, err)
}

func NewErrV4ApiUncategorisedError(msg string, err any) error {
	return new(ErrorTypeV4ApiError, ErrorSubType(ErrorSubTypeV4ApiUncategorisedError), msg, err)
}

func NewErrV4ApiSchemaValidationError(msg string, err any) error {
	return new(ErrorTypeV4ApiError, ErrorSubType(ErrorSubTypeV4ApiSchemaValidationError), msg, err)
}

func NewErrV4ApiAuthorizationError(msg string, err any) error {
	return new(ErrorTypeV4ApiError, ErrorSubType(ErrorSubTypeV4ApiAuthorizationError), msg, err)
}

func NewErrV4ApiResourceNotFoundError(msg string, err any) error {
	return new(ErrorTypeV4ApiError, ErrorSubType(ErrorSubTypeV4ApiResourceNotFoundError), msg, err)
}

func NewErrV4ApiRateLimitError(msg string, err any) error {
	return new(ErrorTypeV4ApiError, ErrorSubType(ErrorSubTypeV4ApiRateLimitError), msg, err)
}

func NewErrV4ApiInternalError(msg string, err any) error {
	return new(ErrorTypeV4ApiError, ErrorSubType(ErrorSubTypeV4ApiInternalServiceError), msg, err)
}

func NewErrV4ApiInvalidInputError(msg string, err any) error {
	return new(ErrorTypeV4ApiError, ErrorSubType(ErrorSubTypeV4ApiInvalidInputError), msg, err)
}

func (apiSubType ErrorSubTypeV4Api) ToError(apiError any) error {
	if constructor, found := V4ApiErrorSubTypeToErrConstructor[apiSubType]; found {
		return constructor("", apiError)
	}

	return NewErrInternalError("Invalid v4 API error sub-type", apiError)
}
