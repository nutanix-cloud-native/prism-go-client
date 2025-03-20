package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseVersion(t *testing.T) {
	tests := []struct {
		name          string
		versionStr    string
		options       []ParseOption
		expectedMajor int
		expectedMinor int
		expectedPatch *int
		expectedBuild *int
		expectPC      bool
		expectError   bool
	}{
		// Basic valid version formats
		{
			name:          "AOS format - major.minor",
			versionStr:    "6.5",
			expectedMajor: 6,
			expectedMinor: 5,
			expectedPatch: nil,
			expectedBuild: nil,
			expectPC:      false,
			expectError:   false,
		},
		{
			name:          "AOS format - major.minor.patch",
			versionStr:    "6.5.1",
			expectedMajor: 6,
			expectedMinor: 5,
			expectedPatch: intPtr(1),
			expectedBuild: nil,
			expectPC:      false,
			expectError:   false,
		},
		{
			name:          "AOS format - major.minor.patch.build",
			versionStr:    "6.5.1.2",
			expectedMajor: 6,
			expectedMinor: 5,
			expectedPatch: intPtr(1),
			expectedBuild: intPtr(2),
			expectPC:      false,
			expectError:   false,
		},
		{
			name:          "PC format with pc prefix",
			versionStr:    "pc.2024.1",
			expectedMajor: 2024,
			expectedMinor: 1,
			expectedPatch: nil,
			expectedBuild: nil,
			expectPC:      true,
			expectError:   false,
		},
		{
			name:          "PC format with pc prefix and patch",
			versionStr:    "pc.2024.1.0",
			expectedMajor: 2024,
			expectedMinor: 1,
			expectedPatch: intPtr(0),
			expectedBuild: nil,
			expectPC:      true,
			expectError:   false,
		},
		{
			name:          "PC format with pc prefix, patch and build",
			versionStr:    "pc.2024.1.0.2",
			expectedMajor: 2024,
			expectedMinor: 1,
			expectedPatch: intPtr(0),
			expectedBuild: intPtr(2),
			expectPC:      true,
			expectError:   false,
		},
		{
			name:          "Legacy PC format",
			versionStr:    "5.17",
			options:       []ParseOption{WithComponentType(TypePC)},
			expectedMajor: 5,
			expectedMinor: 17,
			expectedPatch: nil,
			expectedBuild: nil,
			expectPC:      true,
			expectError:   false,
		},
		{
			name:          "Future PC format",
			versionStr:    "7.3",
			options:       []ParseOption{WithComponentType(TypePC)},
			expectedMajor: 7,
			expectedMinor: 3,
			expectedPatch: nil,
			expectedBuild: nil,
			expectPC:      true,
			expectError:   false,
		},

		// PC prefix tests
		{
			name:        "AOS can't have pc prefix",
			versionStr:  "pc.6.5",
			options:     []ParseOption{WithComponentType(TypeAOS)},
			expectError: true,
		},
		{
			name:          "PC with pc prefix for legacy format (pc.5.17)",
			versionStr:    "pc.5.17",
			options:       []ParseOption{WithComponentType(TypePC)},
			expectedMajor: 5,
			expectedMinor: 17,
			expectedPatch: nil,
			expectedBuild: nil,
			expectPC:      true,
			expectError:   false,
		},
		{
			name:          "PC with pc prefix for future format (pc.7.3)",
			versionStr:    "pc.7.3",
			options:       []ParseOption{WithComponentType(TypePC)},
			expectedMajor: 7,
			expectedMinor: 3,
			expectedPatch: nil,
			expectedBuild: nil,
			expectPC:      true,
			expectError:   false,
		},
		{
			name:          "PC with pc prefix for legacy format (pc.5.17) without explicit component type",
			versionStr:    "pc.5.17",
			expectedMajor: 5,
			expectedMinor: 17,
			expectedPatch: nil,
			expectedBuild: nil,
			expectPC:      true,
			expectError:   false,
		},
		{
			name:          "PC with pc prefix for future format (pc.7.3) without explicit component type",
			versionStr:    "pc.7.3",
			expectedMajor: 7,
			expectedMinor: 3,
			expectedPatch: nil,
			expectedBuild: nil,
			expectPC:      true,
			expectError:   false,
		},
		{
			name:        "PC with pc prefix for invalid range (pc.6.5)",
			versionStr:  "pc.6.5",
			options:     []ParseOption{WithComponentType(TypePC)},
			expectError: true,
		},
		{
			name:        "PC with pc prefix for invalid range (pc.6.5) without explicit component type",
			versionStr:  "pc.6.5",
			expectError: true,
		},

		// Default component types
		{
			name:          "No pc prefix and no component type defaults to AOS",
			versionStr:    "6.5",
			expectedMajor: 6,
			expectedMinor: 5,
			expectedPatch: nil,
			expectedBuild: nil,
			expectPC:      false,
			expectError:   false,
		},

		// Invalid formats
		{
			name:        "Empty string",
			versionStr:  "",
			expectError: true,
		},
		{
			name:        "Only major",
			versionStr:  "6",
			expectError: true,
		},
		{
			name:        "Non-numeric major",
			versionStr:  "a.5",
			expectError: true,
		},
		{
			name:        "Non-numeric minor",
			versionStr:  "6.b",
			expectError: true,
		},
		{
			name:        "Non-numeric patch",
			versionStr:  "6.5.c",
			expectError: true,
		},
		{
			name:        "Non-numeric build",
			versionStr:  "6.5.1.d",
			expectError: true,
		},
		{
			name:        "Too many components",
			versionStr:  "6.5.1.2.3",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := ParseVersion(tt.versionStr, tt.options...)

			// Check for expected errors
			if tt.expectError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedMajor, v.Major(), "Major version mismatch")
			assert.Equal(t, tt.expectedMinor, v.Minor(), "Minor version mismatch")
			assert.Equal(t, tt.expectPC, v.ComponentType() == TypePC, "Component type mismatch")

			// Check patch
			if tt.expectedPatch == nil {
				assert.Nil(t, v.Patch(), "Patch should be nil")
			} else {
				assert.NotNil(t, v.Patch(), "Patch should not be nil")
				assert.Equal(t, *tt.expectedPatch, *v.Patch(), "Patch value mismatch")
			}

			// Check build
			if tt.expectedBuild == nil {
				assert.Nil(t, v.Build(), "Build should be nil")
			} else {
				assert.NotNil(t, v.Build(), "Build should not be nil")
				assert.Equal(t, *tt.expectedBuild, *v.Build(), "Build value mismatch")
			}

			// Check that Original() always returns the original string
			assert.Equal(t, tt.versionStr, v.Original(), "Original string mismatch")
		})
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	tests := []struct {
		name        string
		version1    string
		options1    []ParseOption
		version2    string
		options2    []ParseOption
		expected    bool
		expectError bool
	}{
		// AOS vs AOS comparisons
		{
			name:        "AOS 6.5 >= 5.15",
			version1:    "6.5",
			version2:    "5.15",
			expected:    true,
			expectError: false,
		},
		{
			name:        "AOS 5.15 >= 6.5",
			version1:    "5.15",
			version2:    "6.5",
			expected:    false,
			expectError: false,
		},
		{
			name:        "AOS 6.5.1 >= 6.5.0",
			version1:    "6.5.1",
			version2:    "6.5.0",
			expected:    true,
			expectError: false,
		},
		{
			name:        "AOS 6.5.0 >= 6.5.1",
			version1:    "6.5.0",
			version2:    "6.5.1",
			expected:    false,
			expectError: false,
		},
		{
			name:        "AOS 6.5.0 >= 6.5.0",
			version1:    "6.5.0",
			version2:    "6.5.0",
			expected:    true,
			expectError: false,
		},
		{
			name:        "AOS 6.5.0.1 >= 6.5.0.0",
			version1:    "6.5.0.1",
			version2:    "6.5.0.0",
			expected:    true,
			expectError: false,
		},
		{
			name:        "AOS 6.5.0.0 >= 6.5.0.1",
			version1:    "6.5.0.0",
			version2:    "6.5.0.1",
			expected:    false,
			expectError: false,
		},
		{
			name:        "AOS 6.5 >= 6.5.0",
			version1:    "6.5",
			version2:    "6.5.0",
			expected:    false,
			expectError: false,
		},
		{
			name:        "AOS 6.5.0 >= 6.5",
			version1:    "6.5.0",
			version2:    "6.5",
			expected:    true,
			expectError: false,
		},

		// PC vs PC comparisons
		{
			name:        "PC pc.2024.1 >= pc.2020.7",
			version1:    "pc.2024.1",
			version2:    "pc.2020.7",
			expected:    true,
			expectError: false,
		},
		{
			name:        "PC pc.2020.7 >= pc.2024.1",
			version1:    "pc.2020.7",
			version2:    "pc.2024.1",
			expected:    false,
			expectError: false,
		},
		{
			name:        "PC pc.2024.1.0.2 >= pc.2024.1.0.1",
			version1:    "pc.2024.1.0.2",
			version2:    "pc.2024.1.0.1",
			expected:    true,
			expectError: false,
		},
		{
			name:        "PC pc.2024.1.0.1 >= pc.2024.1.0.2",
			version1:    "pc.2024.1.0.1",
			version2:    "pc.2024.1.0.2",
			expected:    false,
			expectError: false,
		},

		// Original PC vs Legacy PC
		{
			name:        "Original PC 5.17 >= Legacy PC pc.2020.7",
			version1:    "5.17",
			options1:    []ParseOption{WithComponentType(TypePC)},
			version2:    "pc.2020.7",
			expected:    false,
			expectError: false,
		},
		{
			name:        "Legacy PC pc.2020.7 >= Original PC 5.17",
			version1:    "pc.2020.7",
			version2:    "5.17",
			options2:    []ParseOption{WithComponentType(TypePC)},
			expected:    true,
			expectError: false,
		},

		// Legacy PC vs Future PC
		{
			name:        "Legacy PC pc.2024.3 >= Future PC 7.3",
			version1:    "pc.2024.3",
			version2:    "7.3",
			options2:    []ParseOption{WithComponentType(TypePC)},
			expected:    false,
			expectError: false,
		},
		{
			name:        "Future PC 7.3 >= Legacy PC pc.2024.3",
			version1:    "7.3",
			options1:    []ParseOption{WithComponentType(TypePC)},
			version2:    "pc.2024.3",
			expected:    true,
			expectError: false,
		},

		// Original PC vs Future PC
		{
			name:        "Original PC 5.17 >= Future PC 7.3",
			version1:    "5.17",
			options1:    []ParseOption{WithComponentType(TypePC)},
			version2:    "7.3",
			options2:    []ParseOption{WithComponentType(TypePC)},
			expected:    false,
			expectError: false,
		},
		{
			name:        "Future PC 7.3 >= Original PC 5.17",
			version1:    "7.3",
			options1:    []ParseOption{WithComponentType(TypePC)},
			version2:    "5.17",
			options2:    []ParseOption{WithComponentType(TypePC)},
			expected:    true,
			expectError: false,
		},

		// Cross-component comparisons
		{
			name:        "AOS 5.17 >= PC 5.17",
			version1:    "5.17",
			version2:    "5.17",
			options2:    []ParseOption{WithComponentType(TypePC)},
			expectError: true,
		},
		{
			name:        "PC 5.17 >= AOS 5.17",
			version1:    "5.17",
			options1:    []ParseOption{WithComponentType(TypePC)},
			version2:    "5.17",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, err := ParseVersion(tt.version1, tt.options1...)
			if err != nil {
				t.Fatalf("Failed to parse version1 %s: %v", tt.version1, err)
			}

			v2, err := ParseVersion(tt.version2, tt.options2...)
			if err != nil {
				t.Fatalf("Failed to parse version2 %s: %v", tt.version2, err)
			}

			result, err := v1.GreaterThanOrEqual(v2)

			// Check error expectations
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("GreaterThanOrEqual: expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name           string
		versionStr     string
		options        []ParseOption
		expectedOutput string
	}{
		{
			name:           "AOS major.minor",
			versionStr:     "6.5",
			expectedOutput: "6.5",
		},
		{
			name:           "AOS major.minor.patch",
			versionStr:     "6.5.0",
			expectedOutput: "6.5.0",
		},
		{
			name:           "AOS major.minor.patch.build",
			versionStr:     "6.5.0.1",
			expectedOutput: "6.5.0.1",
		},
		{
			name:           "PC with prefix",
			versionStr:     "pc.2024.1",
			expectedOutput: "pc.2024.1",
		},
		{
			name:           "PC with patch",
			versionStr:     "pc.2024.1.0",
			expectedOutput: "pc.2024.1.0",
		},
		{
			name:           "PC with build",
			versionStr:     "pc.2024.1.0.2",
			expectedOutput: "pc.2024.1.0.2",
		},
		// Test that String() generates the same output as Original()
		{
			name:           "PC with different order of components",
			versionStr:     "pc.2024.1.0",
			options:        []ParseOption{WithComponentType(TypePC)},
			expectedOutput: "pc.2024.1.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := ParseVersion(tt.versionStr, tt.options...)
			if err != nil {
				t.Fatalf("Failed to parse version %s: %v", tt.versionStr, err)
			}

			// Check that String() returns the expected format
			result := v.String()
			if result != tt.expectedOutput {
				t.Errorf("String(): expected %q, got %q", tt.expectedOutput, result)
			}

			// Check that Original() returns the original string
			original := v.Original()
			if original != tt.versionStr {
				t.Errorf("Original(): expected %q, got %q", tt.versionStr, original)
			}
		})
	}
}

// Helper function to create int pointers
func intPtr(i int) *int {
	return &i
}
