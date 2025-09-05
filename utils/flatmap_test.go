package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	t.Run("expand simple string value", func(t *testing.T) {
		m := map[string]string{
			"key": "value",
		}

		result := Expand(m, "key")
		assert.Equal(t, "value", result)
	})

	t.Run("expand boolean true", func(t *testing.T) {
		m := map[string]string{
			"key": "true",
		}

		result := Expand(m, "key")
		assert.Equal(t, true, result)
	})

	t.Run("expand boolean false", func(t *testing.T) {
		m := map[string]string{
			"key": "false",
		}

		result := Expand(m, "key")
		assert.Equal(t, false, result)
	})

	t.Run("expand non-existent key", func(t *testing.T) {
		m := map[string]string{
			"other": "value",
		}

		result := Expand(m, "key")
		assert.Nil(t, result)
	})

	t.Run("expand array with count", func(t *testing.T) {
		m := map[string]string{
			"key.#": "2",
			"key.0": "value1",
			"key.1": "value2",
		}

		result := Expand(m, "key")
		expected := []interface{}{"value1", "value2"}
		assert.Equal(t, expected, result)
	})

	t.Run("expand array with unknown count", func(t *testing.T) {
		m := map[string]string{
			"key.#": UnknownVariableValue,
		}

		result := Expand(m, "key")
		assert.Equal(t, UnknownVariableValue, result)
	})

	t.Run("expand empty array", func(t *testing.T) {
		m := map[string]string{
			"key.#": "0",
		}

		result := Expand(m, "key")
		expected := []interface{}{}
		assert.Equal(t, expected, result)
	})

	t.Run("expand map", func(t *testing.T) {
		m := map[string]string{
			"key.subkey1": "value1",
			"key.subkey2": "value2",
		}

		result := Expand(m, "key")
		expected := map[string]interface{}{
			"subkey1": "value1",
			"subkey2": "value2",
		}
		assert.Equal(t, expected, result)
	})

	t.Run("expand nested map", func(t *testing.T) {
		m := map[string]string{
			"key.subkey.nested": "value",
		}

		result := Expand(m, "key")
		expected := map[string]interface{}{
			"subkey": map[string]interface{}{
				"nested": "value",
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("expand array with boolean values", func(t *testing.T) {
		m := map[string]string{
			"key.#": "2",
			"key.0": "true",
			"key.1": "false",
		}

		result := Expand(m, "key")
		expected := []interface{}{true, false}
		assert.Equal(t, expected, result)
	})

	t.Run("expand array with nested structures", func(t *testing.T) {
		m := map[string]string{
			"key.#":        "1",
			"key.0.subkey": "value",
		}

		result := Expand(m, "key")
		expected := []interface{}{
			map[string]interface{}{
				"subkey": "value",
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("expand map with array values", func(t *testing.T) {
		m := map[string]string{
			"key.array.#": "2",
			"key.array.0": "value1",
			"key.array.1": "value2",
		}

		result := Expand(m, "key")
		expected := map[string]interface{}{
			"array": []interface{}{"value1", "value2"},
		}
		assert.Equal(t, expected, result)
	})
}

func TestExpandArray(t *testing.T) {
	t.Run("expand simple array", func(t *testing.T) {
		m := map[string]string{
			"key.#": "3",
			"key.0": "value1",
			"key.1": "value2",
			"key.2": "value3",
		}

		result := expandArray(m, "key")
		expected := []interface{}{"value1", "value2", "value3"}
		assert.Equal(t, expected, result)
	})

	t.Run("expand array with gaps", func(t *testing.T) {
		m := map[string]string{
			"key.#": "3",
			"key.0": "value1",
			"key.2": "value3",
		}

		result := expandArray(m, "key")
		expected := []interface{}{"value1", "value3"}
		assert.Equal(t, expected, result)
	})

	t.Run("expand array with computed values", func(t *testing.T) {
		m := map[string]string{
			"key.#":  "2",
			"key.0":  "value1",
			"key.~1": "computed_value",
		}

		result := expandArray(m, "key")
		expected := []interface{}{"value1", "computed_value"}
		assert.Equal(t, expected, result)
	})

	t.Run("expand array with nested structures", func(t *testing.T) {
		m := map[string]string{
			"key.#":        "1",
			"key.0.subkey": "value",
		}

		result := expandArray(m, "key")
		expected := []interface{}{
			map[string]interface{}{
				"subkey": "value",
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("expand array with boolean values", func(t *testing.T) {
		m := map[string]string{
			"key.#": "2",
			"key.0": "true",
			"key.1": "false",
		}

		result := expandArray(m, "key")
		expected := []interface{}{true, false}
		assert.Equal(t, expected, result)
	})

	t.Run("expand array with non-numeric keys", func(t *testing.T) {
		m := map[string]string{
			"key.#":   "1",
			"key.abc": "value1",
			"key.0":   "value2",
		}

		// This should panic due to non-numeric key
		assert.Panics(t, func() {
			expandArray(m, "key")
		})
	})

	t.Run("expand array with invalid count", func(t *testing.T) {
		m := map[string]string{
			"key.#": "invalid",
		}

		// This should panic due to invalid count
		assert.Panics(t, func() {
			expandArray(m, "key")
		})
	})
}

func TestExpandMap(t *testing.T) {
	t.Run("expand simple map", func(t *testing.T) {
		m := map[string]string{
			"key.subkey1": "value1",
			"key.subkey2": "value2",
		}

		result := expandMap(m, "key.")
		expected := map[string]interface{}{
			"subkey1": "value1",
			"subkey2": "value2",
		}
		assert.Equal(t, expected, result)
	})

	t.Run("expand map with count zero", func(t *testing.T) {
		m := map[string]string{
			"key.%": "0",
		}

		result := expandMap(m, "key.")
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})

	t.Run("expand map with nested structures", func(t *testing.T) {
		m := map[string]string{
			"key.subkey.nested": "value",
		}

		result := expandMap(m, "key.")
		expected := map[string]interface{}{
			"subkey": map[string]interface{}{
				"nested": "value",
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("expand map with array values", func(t *testing.T) {
		m := map[string]string{
			"key.array.#": "2",
			"key.array.0": "value1",
			"key.array.1": "value2",
		}

		result := expandMap(m, "key.")
		expected := map[string]interface{}{
			"array": []interface{}{"value1", "value2"},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("expand map with boolean values", func(t *testing.T) {
		m := map[string]string{
			"key.bool1": "true",
			"key.bool2": "false",
		}

		result := expandMap(m, "key.")
		expected := map[string]interface{}{
			"bool1": true,
			"bool2": false,
		}
		assert.Equal(t, expected, result)
	})

	t.Run("expand map with no matching keys", func(t *testing.T) {
		m := map[string]string{
			"other.key": "value",
		}

		result := expandMap(m, "key.")
		expected := map[string]interface{}{}
		assert.Equal(t, expected, result)
	})

	t.Run("expand map with count key", func(t *testing.T) {
		m := map[string]string{
			"key.%":    "2",
			"key.key1": "value1",
			"key.key2": "value2",
		}

		result := expandMap(m, "key.")
		expected := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}
		assert.Equal(t, expected, result)
	})
}

func TestUnknownVariableValue(t *testing.T) {
	t.Run("unknown variable value constant", func(t *testing.T) {
		assert.Equal(t, "74D93920-ED26-11E3-AC10-0800200C9A66", UnknownVariableValue)
	})
}

func TestComplexExpansion(t *testing.T) {
	t.Run("complex nested structure", func(t *testing.T) {
		m := map[string]string{
			"config.database.host":     "localhost",
			"config.database.port":     "5432",
			"config.database.ssl":      "true",
			"config.servers.#":         "2",
			"config.servers.0.name":    "server1",
			"config.servers.0.enabled": "true",
			"config.servers.1.name":    "server2",
			"config.servers.1.enabled": "false",
		}

		result := Expand(m, "config")
		expected := map[string]interface{}{
			"database": map[string]interface{}{
				"host": "localhost",
				"port": "5432",
				"ssl":  true,
			},
			"servers": []interface{}{
				map[string]interface{}{
					"name":    "server1",
					"enabled": true,
				},
				map[string]interface{}{
					"name":    "server2",
					"enabled": false,
				},
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("mixed array and map structures", func(t *testing.T) {
		m := map[string]string{
			"data.items.#":        "2",
			"data.items.0.id":     "1",
			"data.items.0.tags.#": "2",
			"data.items.0.tags.0": "tag1",
			"data.items.0.tags.1": "tag2",
			"data.items.1.id":     "2",
			"data.items.1.tags.#": "1",
			"data.items.1.tags.0": "tag3",
			"data.metadata.type":  "collection",
		}

		result := Expand(m, "data")
		expected := map[string]interface{}{
			"items": []interface{}{
				map[string]interface{}{
					"id":   "1",
					"tags": []interface{}{"tag1", "tag2"},
				},
				map[string]interface{}{
					"id":   "2",
					"tags": []interface{}{"tag3"},
				},
			},
			"metadata": map[string]interface{}{
				"type": "collection",
			},
		}
		assert.Equal(t, expected, result)
	})
}
