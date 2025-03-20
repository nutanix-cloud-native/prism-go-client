package version

import (
	"fmt"
	"strconv"
	"strings"
)

// ComponentType identifies whether a version is for AOS or PC
type ComponentType int

const (
	// TypeUnknown is for unspecified component types
	TypeUnknown ComponentType = iota
	// TypeAOS is for AOS components
	TypeAOS
	// TypePC is for PC components
	TypePC
)

// String returns the string representation of the component type
func (c ComponentType) String() string {
	switch c {
	case TypeAOS:
		return "AOS"
	case TypePC:
		return "PC"
	default:
		return "Unknown"
	}
}

// NutanixVersion represents a parsed Nutanix version
type NutanixVersion struct {
	major       int
	minor       int
	patch       *int
	build       *int
	hasPCPrefix bool
	compType    ComponentType
	original    string
}

// ParseOption allows customizing the parsing behavior
type ParseOption func(*NutanixVersion)

// WithComponentType explicitly sets the component type (AOS or PC)
func WithComponentType(compType ComponentType) ParseOption {
	return func(v *NutanixVersion) {
		v.compType = compType
	}
}

// ParseVersion parses a Nutanix version string into a structured format
// Handles:
// - Standard SemVer-like format (X.Y.Z) - e.g., 5.17.1, 6.5.0, 7.3.0
// - PC format with prefix (pc.YYYY.MM.X.Y) - e.g., pc.2024.1.0.2
func ParseVersion(version string, opts ...ParseOption) (*NutanixVersion, error) {
	if version == "" {
		return nil, fmt.Errorf("empty version string")
	}

	nutanixVersion := &NutanixVersion{
		original: version,
		compType: TypeUnknown,
	}

	// Apply options
	for _, opt := range opts {
		opt(nutanixVersion)
	}

	// Check for "pc." prefix
	hasPCPrefix := false
	versionStr := version

	if strings.HasPrefix(strings.ToLower(version), "pc.") {
		hasPCPrefix = true
		versionStr = version[3:] // Strip the "pc." prefix

		// If component type is explicitly set to AOS but has pc. prefix, that's an error
		if nutanixVersion.compType == TypeAOS {
			return nil, fmt.Errorf("version with 'pc.' prefix cannot be an AOS version: %s", version)
		}

		// If component type wasn't explicitly set, versions with "pc." prefix are PC
		if nutanixVersion.compType == TypeUnknown {
			nutanixVersion.compType = TypePC
		}
	}

	nutanixVersion.hasPCPrefix = hasPCPrefix

	// Parse the version components
	parts := strings.Split(versionStr, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid version format: %s", version)
	}

	// Parse major version
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid major version: %s", parts[0])
	}

	// Parse minor version
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid minor version: %s", parts[1])
	}

	// Parse patch version if present
	if len(parts) > 2 {
		patch, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("invalid patch version: %s", parts[2])
		}
		nutanixVersion.patch = &patch
	}

	// Parse build version if present
	if len(parts) > 3 {
		build, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, fmt.Errorf("invalid build version: %s", parts[3])
		}
		nutanixVersion.build = &build
	}

	if len(parts) > 4 {
		return nil, fmt.Errorf("invalid version format: %s", version)
	}

	nutanixVersion.major = major
	nutanixVersion.minor = minor

	// If component type wasn't explicitly set and doesn't have pc. prefix
	if nutanixVersion.compType == TypeUnknown {
		// Default to AOS
		nutanixVersion.compType = TypeAOS
	}

	if nutanixVersion.compType == TypePC {
		// PC versions can exist in three types:
		// Original PC format (pc.x.y where x ≤ 5 and y ≤ 17)
		// Legacy PC format (pc.YYYY.MM) where 2020 <= YYYY <= 2024
		// Future PC format (pc.x.y where x ≥ 7 and y >= 3)
		// here we check the bounds to ensure the version is valid
		if (nutanixVersion.major == 5 && nutanixVersion.minor > 17) ||
			(nutanixVersion.major == 6) ||
			(nutanixVersion.major == 7 && nutanixVersion.minor < 3) {
			return nil, fmt.Errorf("invalid PC version: %s", version)
		}
	}

	return nutanixVersion, nil
}

// Major returns the major version component
func (v *NutanixVersion) Major() int {
	return v.major
}

// Minor returns the minor version component
func (v *NutanixVersion) Minor() int {
	return v.minor
}

// Patch returns the patch version component, or nil if not present
func (v *NutanixVersion) Patch() *int {
	return v.patch
}

// Build returns the build version component, or nil if not present
func (v *NutanixVersion) Build() *int {
	return v.build
}

// ComponentType returns the component type (AOS or PC)
func (v *NutanixVersion) ComponentType() ComponentType {
	return v.compType
}

// Original returns the original version string
func (v *NutanixVersion) Original() string {
	return v.original
}

// String returns a string representation of the version
func (v *NutanixVersion) String() string {
	prefix := ""
	if v.hasPCPrefix {
		prefix = "pc."
	}

	if v.build != nil {
		return fmt.Sprintf("%s%d.%d.%d.%d", prefix, v.major, v.minor, *v.patch, *v.build)
	}
	if v.patch != nil {
		return fmt.Sprintf("%s%d.%d.%d", prefix, v.major, v.minor, *v.patch)
	}
	return fmt.Sprintf("%s%d.%d", prefix, v.major, v.minor)
}

// unexported helper function to check if version uses YYYY.MM format
func (v *NutanixVersion) isYearMonth() bool {
	return v.major >= 2000
}

// GreaterThanOrEqual checks if this version is greater than or equal to the other version
func (v *NutanixVersion) GreaterThanOrEqual(other *NutanixVersion) (bool, error) {
	// versions cannot be compared if they have different component types
	if v.compType != other.compType {
		return false, fmt.Errorf("cannot compare incompatible versions: %s (%s) vs %s (%s)",
			v.String(), v.compType, other.String(), other.compType)
	}

	// Same component types can be compared directly in most cases
	if v.compType == other.compType {
		// PC versions with different formats (semver vs year.month) need special handling
		if v.compType == TypePC {
			// One is year.month format and the other is semver format
			if v.isYearMonth() != other.isYearMonth() {
				// Rules:
				// 5.x (Legacy PC) < YYYY.MM (Modern PC) < 7.x (Future PC)

				// Legacy PC (5.x) vs Modern PC (YYYY.MM)
				if !v.isYearMonth() && v.major == 5 && other.isYearMonth() {
					return false, nil
				}
				if v.isYearMonth() && !other.isYearMonth() && other.major == 5 {
					return true, nil
				}

				// Modern PC (YYYY.MM) vs Future PC (7.x)
				if v.isYearMonth() && !other.isYearMonth() && other.major == 7 {
					return false, nil
				}
				if !v.isYearMonth() && v.major == 7 && other.isYearMonth() {
					return true, nil
				}
			}
		}

		// For all other same component type comparisons, compare components directly
		return isGreaterThanOrEqual(v, other), nil
	}

	// For different component types

	// If both are using the same format, we might be able to compare them
	if v.isYearMonth() == other.isYearMonth() {
		if !v.isYearMonth() { // Both using semver format
			// PC versions 5.18 through 7.2 don't exist, so we can't compare across components
			// in that range
			isPCInInvalidRange := func(version *NutanixVersion) bool {
				return version.compType == TypePC && ((version.major == 5 && version.minor > 17) ||
					(version.major == 6) ||
					(version.major == 7 && version.minor < 3))
			}

			if isPCInInvalidRange(v) || isPCInInvalidRange(other) {
				return false, fmt.Errorf("cannot compare PC and AOS versions in range 5.18-7.2: %s (%s) vs %s (%s)",
					v.String(), v.compType, other.String(), other.compType)
			}

			// Otherwise we can compare them directly
			return isGreaterThanOrEqual(v, other), nil
		}

		// Both are year.month format - direct comparison
		return isGreaterThanOrEqual(v, other), nil
	}

	// Otherwise, different formats across component types cannot be compared
	return false, fmt.Errorf("cannot compare incompatible versions: %s (%s) vs %s (%s)",
		v.String(), v.compType, other.String(), other.compType)
}

// compareVersionComponents directly compares the version components
func isGreaterThanOrEqual(v1, v2 *NutanixVersion) bool {
	// Compare major
	if v1.major > v2.major {
		return true
	}
	if v1.major < v2.major {
		return false
	}

	// Compare minor
	if v1.minor > v2.minor {
		return true
	}
	if v1.minor < v2.minor {
		return false
	}

	// Compare patch (nil patch < any patch value)
	switch {
	case v1.patch == nil && v2.patch == nil:
		// Equal patches (both nil)
	case v1.patch != nil && v2.patch == nil:
		return true // v1 has patch, v2 doesn't
	case v1.patch == nil && v2.patch != nil:
		return false // v1 doesn't have patch, v2 does
	case *v1.patch > *v2.patch:
		return true
	case *v1.patch < *v2.patch:
		return false
	}

	// Compare build (nil build < any build value)
	switch {
	case v1.build == nil && v2.build == nil:
		return true // Equal builds (both nil)
	case v1.build != nil && v2.build == nil:
		return true // v1 has build, v2 doesn't
	case v1.build == nil && v2.build != nil:
		return false // v1 doesn't have build, v2 does
	default:
		return *v1.build >= *v2.build
	}
}

// LessThan checks if this version is less than the other version
func (v *NutanixVersion) LessThan(other *NutanixVersion) (bool, error) {
	gte, err := v.GreaterThanOrEqual(other)
	if err != nil {
		return false, err
	}
	return !gte, nil
}

// MustParseVersion parses a version string and panics if it fails
// This is useful for static initialization and testing
func MustParseVersion(version string, opts ...ParseOption) *NutanixVersion {
	v, err := ParseVersion(version, opts...)
	if err != nil {
		panic(fmt.Sprintf("Invalid version string: %s - %v", version, err))
	}
	return v
}
