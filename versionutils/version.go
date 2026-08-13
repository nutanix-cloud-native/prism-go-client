// Package versionutils parses and compares Nutanix version strings.
//
// Nutanix products report versions as '.'-separated integers: a plain release
// version such as "3.5.2.1" from AOS, or a Prism Central version such as
// "2024.3.1" or "pc.7.6.0.5". Parse a version once, then compare it as often
// as needed:
//
//	minPC := versionutils.Parse("7.6")
//	if versionutils.Parse(pcVersion).AtLeast(minPC) {
//		// Prism Central is new enough.
//	}
//
// [Parse] handles every Nutanix version format, so callers do not classify a
// version before parsing it. It normalizes away an optional "pc." prefix,
// surrounding whitespace and casing, and it knows that Prism Central moved
// from calendar versioning (202x.xx.xx) to a 7.x line, so a 7.x version is
// newer than any 202x.xx version despite comparing lower numerically. That
// rule is inert for AOS, which has no calendar-versioned releases.
//
// # Unparsable versions
//
// A version that is not '.'-separated integers, such as a development build
// reporting "master" or an empty string from a failed lookup, sorts as the
// newest possible version. A development build therefore satisfies every
// minimum-version gate, which is deliberate: preflight-style checks should not
// block a test or debug cluster.
//
// A caller that has no version at all is asking a different question, and
// answers it before parsing:
//
//	if pcVersion == "" {
//		return false // nothing was reported, use the fallback path
//	}
//	return versionutils.Parse(pcVersion).AtLeast(minPC)
package versionutils

import (
	"regexp"
	"strconv"
	"strings"
)

// unparsableOrdinal is the value an unparsable version is treated as, making
// it sort above any realistic release.
const unparsableOrdinal = 9999

// masterVersion is the version reported by development builds that are not
// built from a release branch.
const masterVersion = "master"

// pcVersion202xRe matches the calendar-versioned Prism Central format,
// 202x.xx.xx.xx.
var pcVersion202xRe = regexp.MustCompile(`^202(\d(\.\d+)+)$`)

// pcVersion7xRe matches the Prism Central 7.xx.xx format.
var pcVersion7xRe = regexp.MustCompile(`^7((\.\d+)+)$`)

// Version is a parsed Nutanix version, ready to compare. The zero Version is
// not meaningful; construct one with [Parse].
type Version struct {
	// normalized is the version the parts were derived from: trimmed,
	// lower-cased and stripped of an optional "pc." prefix.
	normalized string

	// parts holds the '.'-separated components, or a single unparsableOrdinal
	// if the version could not be parsed.
	parts []int
}

// Parse parses any Nutanix version string, whether it came from AOS
// ("3.5.2.1") or from Prism Central ("2024.3.1", "7.6", "pc.7.6.0.5"). An
// optional "pc." prefix, surrounding whitespace, and casing are normalized
// away.
//
// Parse never fails. A version that is not '.'-separated integers sorts as the
// newest possible version.
func Parse(version string) Version {
	// Prism Central reports versions with inconsistent casing, padding, and an
	// optional "pc." prefix. AOS versions are unaffected by normalization.
	normalized := strings.TrimPrefix(strings.ToLower(strings.TrimSpace(version)), "pc.")

	// A "v" prefix is not a format Nutanix documents, but the version arrives
	// as a free-form string and an unrecognized one sorts as the newest. That
	// is the right answer for "master", where the version is genuinely
	// unknown, and the wrong one for "v7.6", where it is not: leaving the
	// prefix on would let a "v7.6" Prism Central satisfy a 9.9 minimum.
	normalized = strings.TrimPrefix(normalized, "v")

	v := Version{normalized: normalized}
	if normalized == masterVersion {
		v.parts = []int{unparsableOrdinal}
		return v
	}

	v.parts = toIntList(normalized)
	return v
}

// toIntList splits a '.'-separated string into its integer components. If any
// component is not an integer, or the input is empty, it returns
// [unparsableOrdinal] so the version sorts as the newest.
func toIntList(str string) []int {
	parts := strings.Split(str, ".")
	ints := make([]int, 0, len(parts))
	for _, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			return []int{unparsableOrdinal}
		}
		ints = append(ints, val)
	}
	return ints
}

// compare returns 0 if v and other are equal, 1 if v is newer, and -1 if v is
// older. Missing trailing components are treated as zero, so "7.6" and "7.6.0"
// are equal, and a 7.x version outranks a calendar-versioned one.
//
// compare is unexported because no consumer needs three-way ordering; they all
// ask "is this new enough", which is [Version.AtLeast]. Export it when
// something genuinely needs to sort or switch on the result.
func (v Version) compare(other Version) int {
	// Prism Central's 7.x line succeeded its calendar-versioned releases, so
	// it is newer despite the lower leading component.
	if is7X(v.normalized) && is202X(other.normalized) {
		return 1
	}
	if is7X(other.normalized) && is202X(v.normalized) {
		return -1
	}

	maxLen := max(len(v.parts), len(other.parts))

	vParts := make([]int, maxLen)
	otherParts := make([]int, maxLen)
	copy(vParts, v.parts)
	copy(otherParts, other.parts)

	for i, part := range vParts {
		switch {
		case part > otherParts[i]:
			return 1
		case part < otherParts[i]:
			return -1
		}
	}
	return 0
}

// AtLeast reports whether v is newer than or equal to minVersion. An
// unparsable version, including an empty one, reports true.
func (v Version) AtLeast(minVersion Version) bool {
	return v.compare(minVersion) >= 0
}

// String returns the normalized version: trimmed, lower-cased and without any
// "pc." prefix.
func (v Version) String() string {
	return v.normalized
}

// is202X reports whether a normalized version is calendar-versioned.
func is202X(normalized string) bool {
	return pcVersion202xRe.MatchString(normalized)
}

// is7X reports whether a normalized version is on the Prism Central 7.x line.
func is7X(normalized string) bool {
	return pcVersion7xRe.MatchString(normalized)
}
