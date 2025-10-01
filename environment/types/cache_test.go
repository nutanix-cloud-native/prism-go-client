package types

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCachedClientParams(t *testing.T) {
	// Test that the interface can be implemented
	t.Run("interface implementation", func(t *testing.T) {
		testURL, err := url.Parse("https://test.example.com")
		require.NoError(t, err)

		params := &testCachedClientParams{
			endpoint: ManagementEndpoint{
				ApiCredentials: ApiCredentials{
					Username: "testuser",
					Password: "testpass",
				},
				Address: testURL,
			},
			key: "test-key",
		}

		assert.Equal(t, params.endpoint, params.ManagementEndpoint())
		assert.Equal(t, "test-key", params.Key())
	})
}

// testCachedClientParams is a test implementation of CachedClientParams
type testCachedClientParams struct {
	endpoint ManagementEndpoint
	key      string
}

func (t *testCachedClientParams) ManagementEndpoint() ManagementEndpoint {
	return t.endpoint
}

func (t *testCachedClientParams) Key() string {
	return t.key
}
