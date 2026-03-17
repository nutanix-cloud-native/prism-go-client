package v4

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakeOpenAPIError mirrors the SDK's GenericOpenAPIError struct.
type fakeOpenAPIError struct {
	Body   []byte
	Model  interface{}
	Status string
}

func (e fakeOpenAPIError) Error() string {
	return string(e.Body)
}

type noFieldsError struct {
	Msg string
}

func (e noFieldsError) Error() string { return e.Msg }

func makeV4Body(errorGroup, code, message string) []byte {
	return []byte(fmt.Sprintf(`{
		"$dataItemDiscriminator": "vmm.v4.error.ErrorResponse",
		"data": {
			"$errorItemDiscriminator": "List<vmm.v4.error.AppMessage>",
			"error": [{
				"$objectType": "vmm.v4.error.AppMessage",
				"errorGroup": %q,
				"code": %q,
				"message": %q
			}]
		}
	}`, errorGroup, code, message))
}

func makeSchemaValidationBody() []byte {
	return []byte(`{
		"$dataItemDiscriminator": "vmm.v4.error.ErrorResponse",
		"data": {
			"$errorItemDiscriminator": "prism.v4.error.SchemaValidationError",
			"error": {
				"$objectType": "prism.v4.error.SchemaValidationError",
				"error": "Schema validation failed",
				"statusCode": 400
			}
		}
	}`)
}

// ---------------------------------------------------------------------------
// parseStatusCode
// ---------------------------------------------------------------------------

func TestParseStatusCode(t *testing.T) {
	tests := []struct {
		status string
		want   int
	}{
		{"404 Not Found", 404},
		{"429 Too Many Requests", 429},
		{"500 Internal Server Error", 500},
		{"502 Bad Gateway", 502},
		{"503 Service Unavailable", 503},
		{"400 Bad Request", 400},
		{"200 OK", 200},
		{"", 0},
		{"not-a-status", 0},
	}
	for _, tt := range tests {
		t.Run(tt.status, func(t *testing.T) {
			assert.Equal(t, tt.want, parseStatusCode(tt.status))
		})
	}
}

// ---------------------------------------------------------------------------
// Categorise — status-based
// ---------------------------------------------------------------------------

func TestCategorise_NilError(t *testing.T) {
	assert.Nil(t, Categorise(nil, "404 Not Found", nil))
}

func TestCategorise_404_NotFound(t *testing.T) {
	cause := fmt.Errorf("sdk error")
	err := Categorise(cause, "404 Not Found", nil)
	require.Error(t, err)
	assert.True(t, errors.Is(err, converged.ErrNotFound))

	var apiErr *converged.APIError
	require.True(t, errors.As(err, &apiErr))
	assert.Equal(t, converged.ErrNotFound, apiErr.Kind)
	assert.Equal(t, cause, apiErr.Cause)
}

func TestCategorise_429_RateLimit(t *testing.T) {
	cause := fmt.Errorf("sdk error")
	err := Categorise(cause, "429 Too Many Requests", nil)
	require.Error(t, err)
	assert.True(t, errors.Is(err, converged.ErrRateLimit))
}

func TestCategorise_5xx_Internal(t *testing.T) {
	cause := fmt.Errorf("sdk error")

	for _, status := range []string{
		"500 Internal Server Error",
		"502 Bad Gateway",
		"503 Service Unavailable",
		"504 Gateway Timeout",
		"599 Unknown",
	} {
		t.Run(status, func(t *testing.T) {
			err := Categorise(cause, status, nil)
			require.Error(t, err)
			assert.True(t, errors.Is(err, converged.ErrInternal))
		})
	}
}

func TestCategorise_Other4xx_Unclassified(t *testing.T) {
	cause := fmt.Errorf("sdk error")

	for _, status := range []string{
		"400 Bad Request",
		"401 Unauthorized",
		"403 Forbidden",
		"409 Conflict",
		"422 Unprocessable Entity",
	} {
		t.Run(status, func(t *testing.T) {
			err := Categorise(cause, status, nil)
			require.Error(t, err)

			var apiErr *converged.APIError
			require.True(t, errors.As(err, &apiErr))
			assert.Nil(t, apiErr.Kind, "other 4xx without body → unclassified")
			assert.False(t, errors.Is(err, converged.ErrNotFound))
			assert.False(t, errors.Is(err, converged.ErrRateLimit))
			assert.False(t, errors.Is(err, converged.ErrInternal))
		})
	}
}

// ---------------------------------------------------------------------------
// Categorise — body-based fallback
// ---------------------------------------------------------------------------

func TestCategorise_BodyFallback_NotFound(t *testing.T) {
	cause := fmt.Errorf("sdk error")
	body := makeV4Body("RESOURCE_NOT_FOUND", "PLAT-10003", "VM not found")

	err := Categorise(cause, "400 Bad Request", body)
	require.Error(t, err)
	assert.True(t, errors.Is(err, converged.ErrNotFound))

	var apiErr *converged.APIError
	require.True(t, errors.As(err, &apiErr))
	assert.Equal(t, "VM not found", apiErr.Message)
}

func TestCategorise_BodyFallback_RateLimit(t *testing.T) {
	cause := fmt.Errorf("sdk error")
	body := makeV4Body("RATE_LIMIT_EXCEEDED", "SYS-001", "Too many requests")

	err := Categorise(cause, "400 Bad Request", body)
	require.Error(t, err)
	assert.True(t, errors.Is(err, converged.ErrRateLimit))
}

func TestCategorise_BodyFallback_Internal(t *testing.T) {
	cause := fmt.Errorf("sdk error")
	body := makeV4Body("CLUSTERMGMT_SERVICE_ERROR", "CL-001", "Service unavailable")

	err := Categorise(cause, "400 Bad Request", body)
	require.Error(t, err)
	assert.True(t, errors.Is(err, converged.ErrInternal))
}

func TestCategorise_BodyFallback_UnknownGroup(t *testing.T) {
	cause := fmt.Errorf("sdk error")
	body := makeV4Body("SOME_FUTURE_ERROR_GROUP", "X-001", "Something new")

	err := Categorise(cause, "400 Bad Request", body)
	require.Error(t, err)

	var apiErr *converged.APIError
	require.True(t, errors.As(err, &apiErr))
	assert.Nil(t, apiErr.Kind, "unknown errorGroup → unclassified")
}

func TestCategorise_SchemaValidationBody(t *testing.T) {
	cause := fmt.Errorf("sdk error")
	body := makeSchemaValidationBody()

	err := Categorise(cause, "400 Bad Request", body)
	require.Error(t, err)

	var apiErr *converged.APIError
	require.True(t, errors.As(err, &apiErr))
	assert.Nil(t, apiErr.Kind, "schema validation → unclassified")
}

func TestCategorise_StatusTakesPrecedenceOverBody(t *testing.T) {
	cause := fmt.Errorf("sdk error")
	body := makeV4Body("RATE_LIMIT_EXCEEDED", "SYS-001", "Rate limited")

	err := Categorise(cause, "404 Not Found", body)
	require.Error(t, err)
	assert.True(t, errors.Is(err, converged.ErrNotFound), "status 404 should win over body errorGroup")
}

func TestCategorise_BodyMessage_AttachedOnSpecificCode(t *testing.T) {
	cause := fmt.Errorf("sdk error")
	body := makeV4Body("RESOURCE_NOT_FOUND", "PLAT-10003", "VM 00000000 not found")

	err := Categorise(cause, "404 Not Found", body)
	require.Error(t, err)

	var apiErr *converged.APIError
	require.True(t, errors.As(err, &apiErr))
	assert.Equal(t, "VM 00000000 not found", apiErr.Message)
}

func TestCategorise_InvalidJSON(t *testing.T) {
	cause := fmt.Errorf("sdk error")
	err := Categorise(cause, "400 Bad Request", []byte("not json"))
	require.Error(t, err)

	var apiErr *converged.APIError
	require.True(t, errors.As(err, &apiErr))
	assert.Nil(t, apiErr.Kind, "invalid body → unclassified")
}

func TestCategorise_NoStatus_BodyOnly(t *testing.T) {
	cause := fmt.Errorf("sdk error")
	body := makeV4Body("RATE_LIMIT_EXCEEDED", "SYS-001", "Rate limited")

	err := Categorise(cause, "", body)
	require.Error(t, err)
	assert.True(t, errors.Is(err, converged.ErrRateLimit))
}

func TestCategorise_NoStatus_NoBody(t *testing.T) {
	cause := fmt.Errorf("sdk error")
	err := Categorise(cause, "", nil)
	require.Error(t, err)

	var apiErr *converged.APIError
	require.True(t, errors.As(err, &apiErr))
	assert.Nil(t, apiErr.Kind, "no status, no body → unclassified")
}

// ---------------------------------------------------------------------------
// classifyErrorGroup (map-based)
// ---------------------------------------------------------------------------

func TestClassifyErrorGroup(t *testing.T) {
	tests := []struct {
		group string
		want  error
	}{
		{"RESOURCE_NOT_FOUND", converged.ErrNotFound},
		{"ENTITY_NOT_FOUND", converged.ErrNotFound},
		{"NOT_FOUND", converged.ErrNotFound},
		{"REPORTING_RESOURCE_NOT_FOUND", converged.ErrNotFound},

		{"RATE_LIMIT_EXCEEDED", converged.ErrRateLimit},
		{"RATE_LIMIT_ERROR", converged.ErrRateLimit},
		{"OBJECTS_RATE_LIMIT_ERROR", converged.ErrRateLimit},

		{"CLUSTERMGMT_SERVICE_ERROR", converged.ErrInternal},
		{"DISK_INTERNAL_ERROR", converged.ErrInternal},
		{"VOLUME_INTERNAL_ERROR", converged.ErrInternal},
		{"OBJECTS_SERVICE_UNAVAILABLE_ERROR", converged.ErrInternal},

		// Case insensitive
		{"resource_not_found", converged.ErrNotFound},
		{"Rate_Limit_Exceeded", converged.ErrRateLimit},

		// Unknown
		{"SOME_UNKNOWN_GROUP", nil},
		{"", nil},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			assert.Equal(t, tt.want, classifyErrorGroup(tt.group))
		})
	}
}

// ---------------------------------------------------------------------------
// APIError behavior
// ---------------------------------------------------------------------------

func TestAPIError_ErrorString(t *testing.T) {
	t.Run("with Kind and Cause", func(t *testing.T) {
		e := &converged.APIError{Kind: converged.ErrNotFound, Cause: fmt.Errorf("original")}
		assert.Equal(t, "resource not found: original", e.Error())
	})

	t.Run("with Kind, Message, and Cause", func(t *testing.T) {
		e := &converged.APIError{Kind: converged.ErrNotFound, Cause: fmt.Errorf("original"), Message: "VM not found"}
		assert.Equal(t, "resource not found: VM not found: original", e.Error())
	})

	t.Run("with nil Kind", func(t *testing.T) {
		e := &converged.APIError{Kind: nil, Cause: fmt.Errorf("original")}
		assert.Equal(t, "original", e.Error())
	})

	t.Run("empty", func(t *testing.T) {
		e := &converged.APIError{}
		assert.Equal(t, "unknown api error", e.Error())
	})
}

func TestAPIError_Unwrap(t *testing.T) {
	cause := fmt.Errorf("original error")
	e := &converged.APIError{Kind: converged.ErrNotFound, Cause: cause}
	assert.Equal(t, cause, e.Unwrap())
}

func TestAPIError_Is(t *testing.T) {
	e := &converged.APIError{Kind: converged.ErrRateLimit, Cause: fmt.Errorf("x")}
	assert.True(t, errors.Is(e, converged.ErrRateLimit))
	assert.False(t, errors.Is(e, converged.ErrNotFound))
	assert.False(t, errors.Is(e, converged.ErrInternal))
}

func TestAPIError_Is_NilKind(t *testing.T) {
	e := &converged.APIError{Kind: nil, Cause: fmt.Errorf("x")}
	assert.False(t, errors.Is(e, converged.ErrNotFound))
	assert.False(t, errors.Is(e, converged.ErrRateLimit))
	assert.False(t, errors.Is(e, converged.ErrInternal))
}

func TestAPIError_As(t *testing.T) {
	cause := fmt.Errorf("inner")
	wrapped := fmt.Errorf("outer: %w", &converged.APIError{Kind: converged.ErrNotFound, Cause: cause})

	var apiErr *converged.APIError
	require.True(t, errors.As(wrapped, &apiErr))
	assert.Equal(t, converged.ErrNotFound, apiErr.Kind)
	assert.Equal(t, cause, apiErr.Cause)
}

// ---------------------------------------------------------------------------
// GetStatusAndBody
// ---------------------------------------------------------------------------

func TestGetStatusAndBody_ValidStruct(t *testing.T) {
	err := fakeOpenAPIError{
		Body:   []byte(`{"error":"test"}`),
		Status: "404 Not Found",
	}
	status, body := GetStatusAndBody(err)
	assert.Equal(t, "404 Not Found", status)
	assert.Equal(t, []byte(`{"error":"test"}`), body)
}

func TestGetStatusAndBody_Pointer(t *testing.T) {
	err := &fakeOpenAPIError{
		Body:   []byte(`{"error":"test"}`),
		Status: "500 Internal Server Error",
	}
	status, body := GetStatusAndBody(err)
	assert.Equal(t, "500 Internal Server Error", status)
	assert.Equal(t, []byte(`{"error":"test"}`), body)
}

func TestGetStatusAndBody_NotAStruct(t *testing.T) {
	status, body := GetStatusAndBody(fmt.Errorf("plain error"))
	assert.Empty(t, status)
	assert.Nil(t, body)
}

func TestGetStatusAndBody_NoFields(t *testing.T) {
	status, body := GetStatusAndBody(noFieldsError{Msg: "no fields"})
	assert.Empty(t, status)
	assert.Nil(t, body)
}

// ---------------------------------------------------------------------------
// CategoriseFromOpenAPI
// ---------------------------------------------------------------------------

func TestCategoriseFromOpenAPI_NilError(t *testing.T) {
	assert.Nil(t, CategoriseFromOpenAPI(nil))
}

func TestCategoriseFromOpenAPI_OpenAPIError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantKind error
	}{
		{"404", fakeOpenAPIError{Status: "404 Not Found", Body: []byte(`{}`)}, converged.ErrNotFound},
		{"429", fakeOpenAPIError{Status: "429 Too Many Requests", Body: []byte(`{}`)}, converged.ErrRateLimit},
		{"500", fakeOpenAPIError{Status: "500 Internal Server Error", Body: []byte(`{}`)}, converged.ErrInternal},
		{"503", fakeOpenAPIError{Status: "503 Service Unavailable", Body: []byte(`{}`)}, converged.ErrInternal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CategoriseFromOpenAPI(tt.err)
			require.Error(t, err)
			assert.True(t, errors.Is(err, tt.wantKind))
		})
	}
}

func TestCategoriseFromOpenAPI_Other4xx_Unclassified(t *testing.T) {
	for _, status := range []string{
		"400 Bad Request",
		"401 Unauthorized",
		"403 Forbidden",
		"409 Conflict",
	} {
		t.Run(status, func(t *testing.T) {
			err := CategoriseFromOpenAPI(fakeOpenAPIError{Status: status, Body: []byte(`{}`)})
			require.Error(t, err)

			var apiErr *converged.APIError
			require.True(t, errors.As(err, &apiErr))
			assert.Nil(t, apiErr.Kind)
		})
	}
}

func TestCategoriseFromOpenAPI_PlainError(t *testing.T) {
	err := CategoriseFromOpenAPI(fmt.Errorf("connection refused"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "api call failed")
	assert.Contains(t, err.Error(), "connection refused")
}

func TestCategoriseFromOpenAPI_BodyOverridesGeneric4xx(t *testing.T) {
	body := makeV4Body("RATE_LIMIT_EXCEEDED", "SYS-001", "Too many requests")
	err := CategoriseFromOpenAPI(fakeOpenAPIError{
		Status: "400 Bad Request",
		Body:   body,
	})
	require.Error(t, err)
	assert.True(t, errors.Is(err, converged.ErrRateLimit))
}

// ---------------------------------------------------------------------------
// Convenience helpers
// ---------------------------------------------------------------------------

func TestConvenienceHelpers(t *testing.T) {
	assert.True(t, converged.IsNotFound(&converged.APIError{Kind: converged.ErrNotFound}))
	assert.True(t, converged.IsRateLimit(&converged.APIError{Kind: converged.ErrRateLimit}))
	assert.True(t, converged.IsInternal(&converged.APIError{Kind: converged.ErrInternal}))

	assert.False(t, converged.IsNotFound(&converged.APIError{Kind: converged.ErrRateLimit}))
	assert.False(t, converged.IsRateLimit(&converged.APIError{Kind: converged.ErrNotFound}))
	assert.False(t, converged.IsInternal(&converged.APIError{Kind: converged.ErrNotFound}))
}

// ---------------------------------------------------------------------------
// End-to-end: simulating real SDK error flows
// ---------------------------------------------------------------------------

func TestEndToEnd_404Flow(t *testing.T) {
	body := makeV4Body("RESOURCE_NOT_FOUND", "PLAT-10003", "VM 00000000 not found")
	sdkErr := fakeOpenAPIError{
		Body:   body,
		Status: "404 Not Found",
	}

	err := CategoriseFromOpenAPI(sdkErr)
	require.Error(t, err)

	assert.True(t, errors.Is(err, converged.ErrNotFound))
	assert.False(t, errors.Is(err, converged.ErrRateLimit))
	assert.False(t, errors.Is(err, converged.ErrInternal))

	var apiErr *converged.APIError
	require.True(t, errors.As(err, &apiErr))
	assert.Equal(t, converged.ErrNotFound, apiErr.Kind)
	assert.Equal(t, sdkErr, apiErr.Cause)
	assert.Equal(t, "VM 00000000 not found", apiErr.Message)
	assert.Contains(t, err.Error(), "resource not found")
}

func TestEndToEnd_429Flow(t *testing.T) {
	body := makeV4Body("RATE_LIMIT_EXCEEDED", "SYS-001", "Request rate limit exceeded")
	sdkErr := fakeOpenAPIError{
		Body:   body,
		Status: "429 Too Many Requests",
	}

	err := CategoriseFromOpenAPI(sdkErr)
	require.Error(t, err)

	assert.True(t, errors.Is(err, converged.ErrRateLimit))
	assert.False(t, errors.Is(err, converged.ErrNotFound))
	assert.False(t, errors.Is(err, converged.ErrInternal))
	assert.Contains(t, err.Error(), "rate limit exceeded")
}

func TestEndToEnd_500Flow(t *testing.T) {
	body := makeV4Body("CLUSTERMGMT_SERVICE_ERROR", "CL-001", "Cluster service unavailable")
	sdkErr := fakeOpenAPIError{
		Body:   body,
		Status: "500 Internal Server Error",
	}

	err := CategoriseFromOpenAPI(sdkErr)
	require.Error(t, err)

	assert.True(t, errors.Is(err, converged.ErrInternal))
	assert.False(t, errors.Is(err, converged.ErrNotFound))
	assert.False(t, errors.Is(err, converged.ErrRateLimit))
	assert.Contains(t, err.Error(), "internal error")
}

func TestEndToEnd_BodyBasedRateLimit_On400(t *testing.T) {
	body := makeV4Body("RATE_LIMIT_EXCEEDED", "SYS-001", "Rate limited")
	sdkErr := fakeOpenAPIError{
		Body:   body,
		Status: "400 Bad Request",
	}

	err := CategoriseFromOpenAPI(sdkErr)
	require.Error(t, err)
	assert.True(t, errors.Is(err, converged.ErrRateLimit),
		"body errorGroup should classify even when status is 400")

	var apiErr *converged.APIError
	require.True(t, errors.As(err, &apiErr))
	assert.Equal(t, "Rate limited", apiErr.Message)
}

func TestEndToEnd_BodyBasedNotFound_On400(t *testing.T) {
	body := makeV4Body("ENTITY_NOT_FOUND", "PLAT-10004", "Entity does not exist")
	sdkErr := fakeOpenAPIError{
		Body:   body,
		Status: "400 Bad Request",
	}

	err := CategoriseFromOpenAPI(sdkErr)
	require.Error(t, err)
	assert.True(t, errors.Is(err, converged.ErrNotFound),
		"body errorGroup should classify even when status is 400")
}

func TestEndToEnd_BodyBasedInternal_On400(t *testing.T) {
	body := makeV4Body("DISK_INTERNAL_ERROR", "DSK-001", "Disk failure")
	sdkErr := fakeOpenAPIError{
		Body:   body,
		Status: "400 Bad Request",
	}

	err := CategoriseFromOpenAPI(sdkErr)
	require.Error(t, err)
	assert.True(t, errors.Is(err, converged.ErrInternal),
		"body errorGroup should classify even when status is 400")
}

func TestEndToEnd_WrappedError_ErrorsIs(t *testing.T) {
	sdkErr := fakeOpenAPIError{
		Body:   []byte(`{}`),
		Status: "404 Not Found",
	}
	classified := CategoriseFromOpenAPI(sdkErr)
	wrapped := fmt.Errorf("failed to get VM: %w", classified)

	assert.True(t, errors.Is(wrapped, converged.ErrNotFound),
		"errors.Is should work through fmt.Errorf wrapping")
}
