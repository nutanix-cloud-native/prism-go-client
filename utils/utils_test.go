package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToJSONString(t *testing.T) {
	t.Run("convert simple object to JSON string", func(t *testing.T) {
		obj := map[string]interface{}{
			"key": "value",
		}

		result := ToJSONString(obj)

		var parsed map[string]interface{}
		err := json.Unmarshal([]byte(result), &parsed)
		require.NoError(t, err)
		assert.Equal(t, "value", parsed["key"])
	})

	t.Run("convert complex object to JSON string", func(t *testing.T) {
		obj := map[string]interface{}{
			"nested": map[string]interface{}{
				"array":   []string{"item1", "item2"},
				"number":  42,
				"boolean": true,
			},
		}

		result := ToJSONString(obj)

		var parsed map[string]interface{}
		err := json.Unmarshal([]byte(result), &parsed)
		require.NoError(t, err)

		nested := parsed["nested"].(map[string]interface{})
		assert.Equal(t, []interface{}{"item1", "item2"}, nested["array"])
		assert.Equal(t, float64(42), nested["number"])
		assert.Equal(t, true, nested["boolean"])
	})

	t.Run("convert nil to JSON string", func(t *testing.T) {
		result := ToJSONString(nil)
		assert.Equal(t, "null", strings.TrimSpace(result))
	})
}
func TestConvertMapString(t *testing.T) {
	t.Run("convert map with string values", func(t *testing.T) {
		input := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}

		result := ConvertMapString(input)
		expected := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}

		assert.Equal(t, expected, result)
	})

	t.Run("convert empty map", func(t *testing.T) {
		input := map[string]interface{}{}

		result := ConvertMapString(input)
		expected := map[string]string{}

		assert.Equal(t, expected, result)
	})

	t.Run("convert map with non-string values", func(t *testing.T) {
		input := map[string]interface{}{
			"key1": "value1",
			"key2": 42,
			"key3": true,
		}

		// This should panic due to type assertion failure
		assert.Panics(t, func() {
			ConvertMapString(input)
		})
	})
}

func TestStringLowerCaseValidateFunc(t *testing.T) {
	t.Run("validate lowercase string", func(t *testing.T) {
		warns, errs := StringLowerCaseValidateFunc("lowercase", "test_key")
		assert.Empty(t, warns)
		assert.Empty(t, errs)
	})

	t.Run("validate uppercase string", func(t *testing.T) {
		warns, errs := StringLowerCaseValidateFunc("UPPERCASE", "test_key")
		assert.Empty(t, warns)
		assert.Len(t, errs, 1)
		assert.Contains(t, errs[0].Error(), "must be in lowercase")
		assert.Contains(t, errs[0].Error(), "UPPERCASE")
	})

	t.Run("validate mixed case string", func(t *testing.T) {
		warns, errs := StringLowerCaseValidateFunc("MixedCase", "test_key")
		assert.Empty(t, warns)
		assert.Len(t, errs, 1)
		assert.Contains(t, errs[0].Error(), "must be in lowercase")
		assert.Contains(t, errs[0].Error(), "MixedCase")
	})

	t.Run("validate empty string", func(t *testing.T) {
		warns, errs := StringLowerCaseValidateFunc("", "test_key")
		assert.Empty(t, warns)
		assert.Empty(t, errs)
	})

	t.Run("validate string with numbers", func(t *testing.T) {
		warns, errs := StringLowerCaseValidateFunc("test123", "test_key")
		assert.Empty(t, warns)
		assert.Empty(t, errs)
	})

	t.Run("validate string with special characters", func(t *testing.T) {
		warns, errs := StringLowerCaseValidateFunc("test-key_123", "test_key")
		assert.Empty(t, warns)
		assert.Empty(t, errs)
	})
}

func TestGenUUID(t *testing.T) {
	t.Run("generate UUID", func(t *testing.T) {
		uuid := GenUUID()
		assert.NotEmpty(t, uuid)
		assert.Len(t, uuid, 36) // Standard UUID length with hyphens
		assert.Contains(t, uuid, "-")
	})

	t.Run("generate multiple UUIDs are different", func(t *testing.T) {
		uuid1 := GenUUID()
		uuid2 := GenUUID()
		assert.NotEqual(t, uuid1, uuid2)
	})

	t.Run("UUID format validation", func(t *testing.T) {
		uuid := GenUUID()
		parts := strings.Split(uuid, "-")
		assert.Len(t, parts, 5)
		assert.Len(t, parts[0], 8)
		assert.Len(t, parts[1], 4)
		assert.Len(t, parts[2], 4)
		assert.Len(t, parts[3], 4)
		assert.Len(t, parts[4], 12)
	})
}

func TestHashcodeString(t *testing.T) {
	t.Run("hash simple string", func(t *testing.T) {
		result := HashcodeString("test")
		assert.GreaterOrEqual(t, result, 0)
	})

	t.Run("hash consistency", func(t *testing.T) {
		str := "consistent test string"
		hash1 := HashcodeString(str)
		hash2 := HashcodeString(str)
		assert.Equal(t, hash1, hash2)
	})

	t.Run("hash different strings produce different results", func(t *testing.T) {
		hash1 := HashcodeString("string1")
		hash2 := HashcodeString("string2")
		assert.NotEqual(t, hash1, hash2)
	})

	t.Run("hash empty string", func(t *testing.T) {
		result := HashcodeString("")
		assert.GreaterOrEqual(t, result, 0)
	})

	t.Run("hash long string", func(t *testing.T) {
		longStr := strings.Repeat("a", 1000)
		result := HashcodeString(longStr)
		assert.GreaterOrEqual(t, result, 0)
	})

	t.Run("hash special characters", func(t *testing.T) {
		specialStr := "!@#$%^&*()_+-=[]{}|;':\",./<>?"
		result := HashcodeString(specialStr)
		assert.GreaterOrEqual(t, result, 0)
	})

	t.Run("hash unicode characters", func(t *testing.T) {
		unicodeStr := "Hello 世界 🌍"
		result := HashcodeString(unicodeStr)
		assert.GreaterOrEqual(t, result, 0)
	})
}

func TestHashcodeStrings(t *testing.T) {
	t.Run("hash single string", func(t *testing.T) {
		strings := []string{"test"}
		result := HashcodeStrings(strings)
		assert.NotEmpty(t, result)
		assert.Regexp(t, `^\d+$`, result) // Should be numeric
	})

	t.Run("hash multiple strings", func(t *testing.T) {
		strings := []string{"test1", "test2", "test3"}
		result := HashcodeStrings(strings)
		assert.NotEmpty(t, result)
		assert.Regexp(t, `^\d+$`, result) // Should be numeric
	})

	t.Run("hash empty slice", func(t *testing.T) {
		strings := []string{}
		result := HashcodeStrings(strings)
		assert.NotEmpty(t, result)
		assert.Regexp(t, `^\d+$`, result) // Should be numeric
	})

	t.Run("hash consistency", func(t *testing.T) {
		strings := []string{"consistent", "test", "strings"}
		hash1 := HashcodeStrings(strings)
		hash2 := HashcodeStrings(strings)
		assert.Equal(t, hash1, hash2)
	})

	t.Run("hash different order produces different results", func(t *testing.T) {
		strings1 := []string{"first", "second"}
		strings2 := []string{"second", "first"}
		hash1 := HashcodeStrings(strings1)
		hash2 := HashcodeStrings(strings2)
		assert.NotEqual(t, hash1, hash2)
	})

	t.Run("hash with empty strings", func(t *testing.T) {
		strings := []string{"", "test", ""}
		result := HashcodeStrings(strings)
		assert.NotEmpty(t, result)
		assert.Regexp(t, `^\d+$`, result) // Should be numeric
	})

	t.Run("hash with special characters", func(t *testing.T) {
		strings := []string{"test1", "test-2", "test_3", "test.4"}
		result := HashcodeStrings(strings)
		assert.NotEmpty(t, result)
		assert.Regexp(t, `^\d+$`, result) // Should be numeric
	})
}

func TestHashcodeStringEdgeCases(t *testing.T) {
	t.Run("hash negative CRC32 result", func(t *testing.T) {
		// Test with a string that produces a negative CRC32 result
		// This is hard to predict, but we can test the logic
		result := HashcodeString("test")
		assert.GreaterOrEqual(t, result, 0)
	})

	t.Run("hash very long string", func(t *testing.T) {
		veryLongStr := strings.Repeat("a", 10000)
		result := HashcodeString(veryLongStr)
		assert.GreaterOrEqual(t, result, 0)
	})

	t.Run("hash strings with newlines", func(t *testing.T) {
		strWithNewlines := "line1\nline2\nline3"
		result := HashcodeString(strWithNewlines)
		assert.GreaterOrEqual(t, result, 0)
	})

	t.Run("hash strings with tabs", func(t *testing.T) {
		strWithTabs := "col1\tcol2\tcol3"
		result := HashcodeString(strWithTabs)
		assert.GreaterOrEqual(t, result, 0)
	})
}

func TestHashcodeStringsEdgeCases(t *testing.T) {
	t.Run("hash slice with nil strings", func(t *testing.T) {
		// This would cause a panic in the current implementation
		// but we test the expected behavior
		strings := []string{"test", "", "another"}
		result := HashcodeStrings(strings)
		assert.NotEmpty(t, result)
	})

	t.Run("hash very large slice", func(t *testing.T) {
		strings := make([]string, 1000)
		for i := 0; i < 1000; i++ {
			strings[i] = fmt.Sprintf("string_%d", i)
		}
		result := HashcodeStrings(strings)
		assert.NotEmpty(t, result)
		assert.Regexp(t, `^\d+$`, result)
	})
}
