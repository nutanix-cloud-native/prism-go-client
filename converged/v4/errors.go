package v4

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
)

// v4BodyEnvelope matches the outer V4 API error response.
type v4BodyEnvelope struct {
	Data json.RawMessage `json:"data"`
}

// v4ErrorData matches the inner data structure.
type v4ErrorData struct {
	Discriminator string          `json:"$errorItemDiscriminator"`
	Error         json.RawMessage `json:"error"`
}

// v4AppMessage matches a single AppMessage from the error array.
type v4AppMessage struct {
	ErrorGroup string `json:"errorGroup"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

// errorGroupToKind maps V4 API errorGroup strings to sentinel errors.
// To add a new mapping, add a single line here — no other changes needed.
var errorGroupToKind = map[string]error{
	// Not Found
	"RESOURCE_NOT_FOUND":           converged.ErrNotFound,
	"ENTITY_NOT_FOUND":             converged.ErrNotFound,
	"NOT_FOUND":                    converged.ErrNotFound,
	"REPORTING_RESOURCE_NOT_FOUND": converged.ErrNotFound,

	// Rate Limit
	"RATE_LIMIT_EXCEEDED":      converged.ErrRateLimit,
	"RATE_LIMIT_ERROR":         converged.ErrRateLimit,
	"OBJECTS_RATE_LIMIT_ERROR": converged.ErrRateLimit,

	// Internal / Service
	"CLUSTERMGMT_SERVICE_ERROR":         converged.ErrInternal,
	"DISK_INTERNAL_ERROR":               converged.ErrInternal,
	"VOLUME_INTERNAL_ERROR":             converged.ErrInternal,
	"OBJECTS_SERVICE_UNAVAILABLE_ERROR": converged.ErrInternal,
}

// parseStatusCode extracts the numeric HTTP status code from a status line.
// The SDK's GenericOpenAPIError.Status is the full HTTP status line
// (Go's http.Response.Status), e.g. "404 Not Found", "429 Too Many Requests".
func parseStatusCode(status string) int {
	if status == "" {
		return 0
	}
	codeStr := strings.SplitN(status, " ", 2)[0]
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		return 0
	}
	return code
}

// Categorise classifies an error based on HTTP status code and optional response body.
//
// Classification order:
//  1. HTTP status: 404 → ErrNotFound, 429 → ErrRateLimit, 500-599 → ErrInternal.
//  2. Body fallback: for any other status, the body's errorGroup can classify
//     (e.g. a 400 with errorGroup RATE_LIMIT_EXCEEDED → ErrRateLimit).
//  3. If nothing matches, Kind is nil (unclassified) — Cause is still preserved.
func Categorise(err error, status string, body []byte) error {
	if err == nil {
		return nil
	}

	code := parseStatusCode(status)

	var bodyKind error
	var bodyMsg string
	if len(body) > 0 {
		bodyMsg, bodyKind = kindFromBody(body)
	}

	switch {
	case code == 404:
		return &converged.APIError{Kind: converged.ErrNotFound, Cause: err, Message: bodyMsg}
	case code == 429:
		return &converged.APIError{Kind: converged.ErrRateLimit, Cause: err, Message: bodyMsg}
	case code >= 500 && code <= 599:
		return &converged.APIError{Kind: converged.ErrInternal, Cause: err, Message: bodyMsg}
	}

	// Body-based fallback for any other status (including other 4xx)
	if bodyKind != nil {
		return &converged.APIError{Kind: bodyKind, Cause: err, Message: bodyMsg}
	}

	// Unclassified — Cause is still preserved for debugging
	return &converged.APIError{Kind: nil, Cause: err}
}

// kindFromBody attempts to classify an error from the V4 API JSON body.
// Returns ("", nil) if the body can't be parsed or doesn't contain a known errorGroup.
func kindFromBody(body []byte) (string, error) {
	var envelope v4BodyEnvelope
	if json.Unmarshal(body, &envelope) != nil || len(envelope.Data) == 0 {
		return "", nil
	}

	var data v4ErrorData
	if json.Unmarshal(envelope.Data, &data) != nil {
		return "", nil
	}

	if strings.Contains(data.Discriminator, "SchemaValidationError") {
		return "", nil
	}

	var messages []v4AppMessage
	if json.Unmarshal(data.Error, &messages) != nil || len(messages) == 0 {
		return "", nil
	}

	first := messages[0]
	kind := classifyErrorGroup(first.ErrorGroup)
	return first.Message, kind
}

// classifyErrorGroup looks up the errorGroup string in errorGroupToKind.
// Returns nil for unknown groups (caller falls through to status-based or unclassified).
func classifyErrorGroup(errorGroup string) error {
	if kind, ok := errorGroupToKind[strings.ToUpper(errorGroup)]; ok {
		return kind
	}
	return nil
}

// GetStatusAndBody extracts the Status and Body fields from any error struct
// using reflection. Works across all SDK clients (vmm, clustermgmt, prism, etc.)
// which each define their own GenericOpenAPIError with the same field layout.
func GetStatusAndBody(err error) (status string, body []byte) {
	v := reflect.ValueOf(err)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return "", nil
	}
	if f := v.FieldByName("Status"); f.IsValid() && f.Kind() == reflect.String {
		status = f.String()
	}
	if f := v.FieldByName("Body"); f.IsValid() && f.Kind() == reflect.Slice && f.Type().Elem().Kind() == reflect.Uint8 {
		body = f.Bytes()
	}
	return status, body
}

// CategoriseFromOpenAPI is the single entry point for classifying SDK errors.
// It extracts Status and Body from the error via reflection, then delegates to Categorise.
func CategoriseFromOpenAPI(err error) error {
	if err == nil {
		return nil
	}
	status, body := GetStatusAndBody(err)
	if status == "" && len(body) == 0 {
		return fmt.Errorf("api call failed: %w", err)
	}
	return Categorise(err, status, body)
}
