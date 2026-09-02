package v4

import "strings"

const unsupportedProtocolScheme = "unsupported protocol scheme"

// IsStaleClientError reports whether err is the upstream SDK transport failure
// that occurs when a relative OIDC redirect is followed after session expiry
// (or similar auth header loss), typically containing "unsupported protocol scheme".
func IsStaleClientError(err error) bool {
	return err != nil && strings.Contains(err.Error(), unsupportedProtocolScheme)
}
