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
		{
			"extracts AOS from Prism fullVersion",
			"el8.5-release-ganges-7.5.1-stable-b2a0591e15d312a5888e2eee14e1c4bffc3e3b3d",
			"7.5.1",
		},
		{
			"extracts AOS from a non-el OS prefix",
			"ol8.5-release-ganges-7.5.1-stable-b2a0591e15d312a5888e2eee14e1c4bffc3e3b3d",
			"7.5.1",
		},
		{
			"extracts AOS 7.6 from Prism fullVersion",
			"el9-release-ganges-7.6-stable-abc123",
			"7.6",
		},
		{
			"extracts AOS when there is no OS prefix",
			"release-ganges-7.5.1-stable-abc123",
			"7.5.1",
		},
		{
			"extracts a four-component AOS token",
			"el8.5-release-ganges-7.5.1.2-stable-abc123",
			"7.5.1.2",
		},
		{
			"extracts AOS 6.x",
			"el7-release-fraser-6.8.1-stable-abc123",
			"6.8.1",
		},
		{
			"fullVersion 7.5 stays 7.5",
			"el8.5-release-ganges-7.5-stable-abc123",
			"7.5",
		},
		{
			"trims whitespace around fullVersion",
			"  el8.5-release-ganges-7.5.1-stable-abc  ",
			"7.5.1",
		},
		{
			"internal opt-master fullVersion stays unknown",
			"el9-opt-master-56c4af6e0fb9c6a2aed0e517cbfb3768cde74311",
			"el9-opt-master-56c4af6e0fb9c6a2aed0e517cbfb3768cde74311",
		},
		{
			"dotted OS prefix without AOS stays unknown",
			"el8.5-opt-master-56c4af6e0fb9c6a2aed0e517cbfb3768cde74311",
			"el8.5-opt-master-56c4af6e0fb9c6a2aed0e517cbfb3768cde74311",
		},
		{
			"uppercased fullVersion still extracts AOS",
			"EL8.5-RELEASE-GANGES-7.5.1-STABLE-ABC123",
			"7.5.1",
		},
		{
			"generic OS letters are not limited to el/ol",
			"centos8.5-release-ganges-7.5.1-stable-abc123",
			"7.5.1",
		},
		{
			"an extra kernel-style token is not preferred over AOS",
			"el8.5-5.10.0-release-ganges-7.5.1-stable-abc123",
			"7.5.1",
		},
		{
			"a dotted token without OS prefix or -stable- is left unknown",
			"build-ganges-7.5.1-abc123",
			"build-ganges-7.5.1-abc123",
		},
		{
			"a single-component AOS token is not extracted",
			"el9-release-ganges-7-stable-abc123",
			"el9-release-ganges-7-stable-abc123",
		},
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
		{
			"truncated-looking fullVersion satisfies the patch floor",
			"el8.5-release-ganges-7.5.1-stable-b2a0591e15d312a5888e2eee14e1c4bffc3e3b3d",
			"7.5.1",
			true,
		},
		{
			"fullVersion 7.5 does not satisfy 7.5.1",
			"el8.5-release-ganges-7.5-stable-abc123",
			"7.5.1",
			false,
		},
		{
			"fullVersion 7.5.1 does not satisfy 7.6",
			"el8.5-release-ganges-7.5.1-stable-b2a0591e15d312a5888e2eee14e1c4bffc3e3b3d",
			"7.6",
			false,
		},
		{
			"fullVersion 7.5.1 satisfies a 7.5 floor",
			"el8.5-release-ganges-7.5.1-stable-abc123",
			"7.5",
			true,
		},
		{
			"fullVersion 7.6 satisfies the 7.5.1 floor",
			"el9-release-ganges-7.6-stable-abc123",
			"7.5.1",
			true,
		},
		{
			"four-component fullVersion satisfies the patch floor",
			"el8.5-release-ganges-7.5.1.2-stable-abc123",
			"7.5.1",
			true,
		},
		{
			"ol-prefixed fullVersion satisfies the patch floor",
			"ol8.5-release-ganges-7.5.1-stable-abc123",
			"7.5.1",
			true,
		},
		{
			"fullVersion without an OS prefix still satisfies the patch floor",
			"release-ganges-7.5.1-stable-abc123",
			"7.5.1",
			true,
		},
		{
			"AOS 6.x fullVersion does not satisfy 7.5.1",
			"el7-release-fraser-6.8.1-stable-abc123",
			"7.5.1",
			false,
		},
		{
			"AOS 6.x fullVersion satisfies a 6.5.1 floor",
			"el7-release-fraser-6.8.1-stable-abc123",
			"6.5.1",
			true,
		},
		{
			"an extra kernel-style token still satisfies 7.5.1",
			"el8.5-5.10.0-release-ganges-7.5.1-stable-abc123",
			"7.5.1",
			true,
		},
		{
			"internal opt-master fullVersion satisfies every gate",
			"el9-opt-master-56c4af6e0fb9c6a2aed0e517cbfb3768cde74311",
			"7.5.1",
			true,
		},
		{
			"dotted OS-only fullVersion is unknown and satisfies every gate",
			"ol8.5-opt-master-56c4af6e0fb9c6a2aed0e517cbfb3768cde74311",
			"9.9",
			true,
		},
		{
			"a non-fullVersion dotted token stays unknown and satisfies every gate",
			"build-ganges-7.5.1-abc123",
			"9.9",
			true,
		},
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

func TestParseReported(t *testing.T) {
	t.Parallel()

	full751 := "el8.5-release-ganges-7.5.1-stable-b2a0591e15d312a5888e2eee14e1c4bffc3e3b3d"
	full75 := "el8.5-release-ganges-7.5-stable-abc123"
	full76 := "el9-release-ganges-7.6-stable-abc123"
	fullMaster := "el9-opt-master-56c4af6e0fb9c6a2aed0e517cbfb3768cde74311"
	osOnlyFull := "ol8.5-opt-master-56c4af6e0fb9c6a2aed0e517cbfb3768cde74311"

	tests := []struct {
		name     string
		short    string
		full     string
		want     string
		atLeast  string
		wantGate bool
	}{
		{
			name:     "truncated short is refined by fullVersion patch",
			short:    "7.5",
			full:     full751,
			want:     "7.5.1",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "truncated short with matching 7.5 fullVersion stays below patch floor",
			short:    "7.5",
			full:     full75,
			want:     "7.5",
			atLeast:  "7.5.1",
			wantGate: false,
		},
		{
			name:     "short 7.6 with 7.6 fullVersion passes metro floor",
			short:    "7.6",
			full:     full76,
			want:     "7.6",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "short 7.6.1 with no fullVersion passes",
			short:    "7.6.1",
			full:     "",
			want:     "7.6.1",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "short 7.5 with no fullVersion stays 7.5",
			short:    "7.5",
			full:     "",
			want:     "7.5",
			atLeast:  "7.5.1",
			wantGate: false,
		},
		{
			name:     "empty short uses AOS token from fullVersion",
			short:    "",
			full:     full751,
			want:     "7.5.1",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "empty short and OS-only fullVersion stays unknown",
			short:    "",
			full:     osOnlyFull,
			want:     "",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "master short is not replaced by a GA fullVersion",
			short:    "master",
			full:     full751,
			want:     "master",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "master short with opt-master fullVersion stays master",
			short:    "master",
			full:     fullMaster,
			want:     "master",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "OS token 8.5 does not refine short 7.5",
			short:    "7.5",
			full:     osOnlyFull,
			want:     "7.5",
			atLeast:  "7.5.1",
			wantGate: false,
		},
		{
			name:     "OS token 8.5 does not refine short 8.6",
			short:    "8.6",
			full:     osOnlyFull,
			want:     "8.6",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "OS token 8.5 does not refine short 9.5",
			short:    "9.5",
			full:     osOnlyFull,
			want:     "9.5",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "short already at patch is not downgraded by a 7.5 fullVersion",
			short:    "7.5.1",
			full:     full75,
			want:     "7.5.1",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "whitespace around short and fullVersion is trimmed",
			short:    "  7.5  ",
			full:     "  " + full751 + "  ",
			want:     "7.5.1",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "uppercase MASTER is left as master",
			short:    "MASTER",
			full:     full751,
			want:     "master",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "ol prefix still yields AOS patch from fullVersion",
			short:    "7.5",
			full:     "ol8.5-release-ganges-7.5.1-stable-abc",
			want:     "7.5.1",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "both fields empty stays empty",
			short:    "",
			full:     "",
			want:     "",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "short 7.5 is not replaced by a mismatched 7.6 fullVersion",
			short:    "7.5",
			full:     full76,
			want:     "7.5",
			atLeast:  "7.5.1",
			wantGate: false,
		},
		{
			name:     "short 7.5 is refined by a four-component fullVersion",
			short:    "7.5",
			full:     "el8.5-release-ganges-7.5.1.2-stable-abc123",
			want:     "7.5.1.2",
			atLeast:  "7.5.1",
			wantGate: true,
		},
		{
			name:     "short 7.5.1 is refined by a more specific patch",
			short:    "7.5.1",
			full:     "el8.5-release-ganges-7.5.1.2-stable-abc123",
			want:     "7.5.1.2",
			atLeast:  "7.5.1",
			wantGate: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ParseReported(tt.short, tt.full)
			assert.Equal(t, tt.want, got.String())
			assert.Equal(t, tt.wantGate, got.AtLeast(Parse(tt.atLeast)))
		})
	}
}

func TestParseFullVersionDoesNotTreatOSPrefixAsAOS(t *testing.T) {
	t.Parallel()

	// An internal build's fullVersion has el8.5/el9 and no AOS token. It must
	// stay unknown (newest), not parse as 8.5 and fail a 7.5.1 gate.
	internal := Parse("el9-opt-master-56c4af6e0fb9c6a2aed0e517cbfb3768cde74311")
	assert.True(t, internal.AtLeast(Parse("7.5.1")))
	assert.True(t, internal.AtLeast(Parse("9.9")))

	olInternal := Parse("ol8.5-opt-master-56c4af6e0fb9c6a2aed0e517cbfb3768cde74311")
	assert.True(t, olInternal.AtLeast(Parse("7.5.1")))
	assert.True(t, olInternal.AtLeast(Parse("9.9")))
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
