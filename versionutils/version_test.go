package versionutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNormalizes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"lowercases and trims", "  PC.7.3.0  ", "7.3.0"},
		{"handles missing prefix", "7.5.1", "7.5.1"},
		{"handles uppercase PC prefix", "Pc.2024.1.0.1", "2024.1.0.1"},
		{"strips only a leading prefix", "pc.pc.7.3", "pc.7.3"},
		{"leaves an empty version empty", "   ", ""},
		{"leaves an AOS version alone", "6.5.1", "6.5.1"},
		{"strips a v prefix", "v7.6", "7.6"},
		{"strips an uppercase V prefix", "V7.6", "7.6"},
		{"strips a v prefix after pc.", "pc.v7.6.0.5", "7.6.0.5"},
		{"leaves master alone", "master", "master"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, Parse(tt.input).String())
		})
	}
}

func TestCompare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		v1       string
		v2       string
		expected int
	}{
		{"equal versions", "3.5.2.1", "3.5.2.1", 0},
		{"greater patch", "3.5.2.2", "3.5.2.1", 1},
		{"lesser patch", "3.5.2.1", "3.5.2.2", -1},
		{"greater major", "10.0", "9.9.9", 1},
		{"missing components are zero", "7.6", "7.6.0", 0},
		{"missing components are zero, reversed", "7.6.0", "7.6", 0},
		{"trailing component breaks the tie", "7.6.1", "7.6", 1},
		{"master sorts newest", "master", "7.5.0", 1},
		{"master sorts newest, reversed", "7.5.0", "master", -1},
		{"master is case insensitive", "MASTER", "7.5.0", 1},
		{"unparsable sorts newest", "not-a-version", "7.5.0", 1},
		{"empty sorts newest", "", "7.5.0", 1},
		{"two unparsable versions are equal", "master", "", 0},
		// Prism Central's 7.x line succeeded its calendar-versioned releases.
		{"7.x beats 202x", "7.3.0", "2024.1.0.1", 1},
		{"202x loses to 7.x", "2024.1.0.1", "7.3.0", -1},
		{"pc prefix is ignored", "pc.7.6.0.5", "7.6.0.5", 0},
		{"pc prefix is ignored across lines", "PC.7.3", "pc.2024.3.1", 1},
		{"two calendar versions compare numerically", "2024.3.1", "2023.4", 1},
		// AOS has no calendar-versioned releases, so the rule never fires.
		{"AOS versions compare numerically", "6.8.1", "6.5.3.1", 1},
		{"AOS 7.x compares numerically", "7.0", "6.8", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, Parse(tt.v1).compare(Parse(tt.v2)))
		})
	}
}

func TestAtLeast(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		version    string
		minVersion string
		expected   bool
	}{
		{"equal", "7.6", "7.6", true},
		{"greater", "7.6.0.5", "7.6", true},
		{"lesser", "7.5.0.5", "7.6", false},
		{"pc prefix is ignored", "pc.7.6.0.5", "7.6", true},
		{"calendar version does not satisfy a 7.x minimum", "2024.3", "7.6", false},
		{"7.x satisfies a calendar minimum", "7.0", "2024.3", true},
		{"missing components are zero", "7.6", "7.6.0", true},
		{"AOS minimum", "6.8.1", "6.5.1", true},
		{"AOS below minimum", "6.5.0", "6.5.1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, Parse(tt.version).AtLeast(Parse(tt.minVersion)))
		})
	}
}

// TestVPrefixIsNotUnknown guards the distinction between a version we cannot
// read and one wearing a conventional prefix. "master" is genuinely unknown
// and gates open by design; "v7.6" is 7.6 and must gate like it, rather than
// collapsing to the unparsable sentinel and satisfying every minimum.
func TestVPrefixIsNotUnknown(t *testing.T) {
	t.Parallel()

	for _, version := range []string{"v7.6", "V7.6", "pc.v7.6", "  v7.6  "} {
		t.Run(version, func(t *testing.T) {
			t.Parallel()

			assert.True(t, Parse(version).AtLeast(Parse("7.6")))
			assert.True(t, Parse(version).AtLeast(Parse("6.5.1")))
			assert.False(t, Parse(version).AtLeast(Parse("9.9")),
				"a v-prefixed version must not satisfy every minimum")
			assert.False(t, Parse(version).AtLeast(Parse("7.7")))
			assert.Equal(t, 0, Parse(version).compare(Parse("7.6")))
		})
	}

	// Stripping the prefix off something still unreadable leaves it
	// unreadable, and so still the newest version.
	assert.True(t, Parse("v-some-branch").AtLeast(Parse("9.9")))
	assert.True(t, Parse("master").AtLeast(Parse("9.9")))
}

// TestMasterIsLatest pins the policy that a development build reporting
// "master" satisfies every minimum-version gate. Consumers rely on this to
// avoid blocking test and debug clusters.
func TestMasterIsLatest(t *testing.T) {
	t.Parallel()

	for _, version := range []string{"master", "MASTER", "Master", "pc.master", "  master  "} {
		t.Run(version, func(t *testing.T) {
			t.Parallel()

			for _, minVersion := range []string{"6.5.1", "7.6", "7.10.2", "2024.3", "2025.1.0.1"} {
				assert.True(t, Parse(version).AtLeast(Parse(minVersion)),
					"Parse(%q).AtLeast(%q) must treat master as latest", version, minVersion)
			}
		})
	}
}

// TestUnreportedVersionIsTheCallersQuestion documents the pattern for a
// caller that has no version at all. That is a question about whether the
// lookup succeeded, not about version ordering, so it is answered before
// parsing -- unlike "master", which must still gate as the newest version.
func TestUnreportedVersionIsTheCallersQuestion(t *testing.T) {
	t.Parallel()

	minPC := Parse("7.6")
	supported := func(version string) bool {
		if version == "" {
			return false
		}
		return Parse(version).AtLeast(minPC)
	}

	assert.True(t, supported("pc.7.6.0.5"))
	assert.False(t, supported("pc.7.5.0.5"))
	assert.False(t, supported("2024.3"))
	assert.False(t, supported(""))
	// A development build is newer than any release, so it gates open.
	assert.True(t, supported("master"))
}

// TestParsedVersionIsReusable guards the reason for parsing separately from
// comparing: consumers parse once and gate many times.
func TestParsedVersionIsReusable(t *testing.T) {
	t.Parallel()

	v := Parse("pc.7.6.0.5")
	for range 3 {
		assert.True(t, v.AtLeast(Parse("7.6")))
		assert.False(t, v.AtLeast(Parse("7.7")))
		assert.Equal(t, "7.6.0.5", v.String())
	}
}
