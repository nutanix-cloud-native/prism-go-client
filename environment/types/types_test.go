package types

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApiCredentials(t *testing.T) {
	t.Run("JSON serialization", func(t *testing.T) {
		creds := ApiCredentials{
			Username: "testuser",
			Password: "testpass",
			KeyPair:  "testkeypair",
		}

		data, err := json.Marshal(creds)
		require.NoError(t, err)

		var decoded ApiCredentials
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, creds.Username, decoded.Username)
		assert.Equal(t, creds.Password, decoded.Password)
		assert.Equal(t, creds.KeyPair, decoded.KeyPair)
	})

	t.Run("JSON serialization with empty fields", func(t *testing.T) {
		creds := ApiCredentials{}

		data, err := json.Marshal(creds)
		require.NoError(t, err)

		var decoded ApiCredentials
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Empty(t, decoded.Username)
		assert.Empty(t, decoded.Password)
		assert.Empty(t, decoded.KeyPair)
	})
}

func TestManagementEndpoint(t *testing.T) {
	t.Run("JSON serialization", func(t *testing.T) {
		testURL, err := url.Parse("https://test.example.com")
		require.NoError(t, err)

		endpoint := ManagementEndpoint{
			ApiCredentials: ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address:               testURL,
			Insecure:              true,
			AdditionalTrustBundle: "test-bundle",
		}

		data, err := json.Marshal(endpoint)
		require.NoError(t, err)

		var decoded ManagementEndpoint
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, endpoint.Username, decoded.Username)
		assert.Equal(t, endpoint.Password, decoded.Password)
		assert.Equal(t, endpoint.Address.String(), decoded.Address.String())
		assert.Equal(t, endpoint.Insecure, decoded.Insecure)
		assert.Equal(t, endpoint.AdditionalTrustBundle, decoded.AdditionalTrustBundle)
	})

	t.Run("JSON serialization with nil address", func(t *testing.T) {
		endpoint := ManagementEndpoint{
			ApiCredentials: ApiCredentials{
				Username: "testuser",
			},
			Address:  nil,
			Insecure: false,
		}

		data, err := json.Marshal(endpoint)
		require.NoError(t, err)

		var decoded ManagementEndpoint
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, endpoint.Username, decoded.Username)
		assert.Nil(t, decoded.Address)
		assert.Equal(t, endpoint.Insecure, decoded.Insecure)
	})
}

func TestTopology(t *testing.T) {
	t.Run("topology map operations", func(t *testing.T) {
		topology := Topology{
			"region": "us-west-1",
			"zone":   "us-west-1a",
			"rack":   "rack-1",
		}

		assert.Equal(t, "us-west-1", topology["region"])
		assert.Equal(t, "us-west-1a", topology["zone"])
		assert.Equal(t, "rack-1", topology["rack"])

		// Test modification
		topology["rack"] = "rack-2"
		assert.Equal(t, "rack-2", topology["rack"])

		// Test addition
		topology["datacenter"] = "dc-1"
		assert.Equal(t, "dc-1", topology["datacenter"])
	})

	t.Run("empty topology", func(t *testing.T) {
		topology := Topology{}
		assert.Empty(t, topology)
		value, exists := topology["nonexistent"]
		assert.False(t, exists)
		assert.Empty(t, value)
	})
}

func TestGetManagementEndpointHash(t *testing.T) {
	t.Run("successful hash generation", func(t *testing.T) {
		testURL, err := url.Parse("https://test.example.com")
		require.NoError(t, err)

		endpoint := ManagementEndpoint{
			ApiCredentials: ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address:  testURL,
			Insecure: true,
		}

		hash, err := endpoint.GetHash()
		require.NoError(t, err)
		assert.NotEmpty(t, hash)
		assert.Len(t, hash, 64) // SHA256 hex string length
	})

	t.Run("hash consistency", func(t *testing.T) {
		testURL, err := url.Parse("https://test.example.com")
		require.NoError(t, err)

		endpoint := ManagementEndpoint{
			ApiCredentials: ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address:  testURL,
			Insecure: true,
		}

		hash1, err := endpoint.GetHash()
		require.NoError(t, err)

		hash2, err := endpoint.GetHash()
		require.NoError(t, err)

		assert.Equal(t, hash1, hash2)
	})

	t.Run("different endpoints produce different hashes", func(t *testing.T) {
		testURL1, err := url.Parse("https://test1.example.com")
		require.NoError(t, err)

		testURL2, err := url.Parse("https://test2.example.com")
		require.NoError(t, err)

		endpoint1 := ManagementEndpoint{
			ApiCredentials: ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address: testURL1,
		}

		endpoint2 := ManagementEndpoint{
			ApiCredentials: ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address: testURL2,
		}

		hash1, err := endpoint1.GetHash()
		require.NoError(t, err)

		hash2, err := endpoint2.GetHash()
		require.NoError(t, err)

		assert.NotEqual(t, hash1, hash2)
	})

	t.Run("hash changes with credential changes", func(t *testing.T) {
		testURL, err := url.Parse("https://test.example.com")
		require.NoError(t, err)

		endpoint1 := ManagementEndpoint{
			ApiCredentials: ApiCredentials{
				Username: "user1",
				Password: "pass1",
			},
			Address: testURL,
		}

		endpoint2 := ManagementEndpoint{
			ApiCredentials: ApiCredentials{
				Username: "user2",
				Password: "pass2",
			},
			Address: testURL,
		}

		hash1, err := endpoint1.GetHash()
		require.NoError(t, err)

		hash2, err := endpoint2.GetHash()
		require.NoError(t, err)

		assert.NotEqual(t, hash1, hash2)
	})

	t.Run("hash changes with insecure flag", func(t *testing.T) {
		testURL, err := url.Parse("https://test.example.com")
		require.NoError(t, err)

		endpoint1 := ManagementEndpoint{
			ApiCredentials: ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address:  testURL,
			Insecure: false,
		}

		endpoint2 := ManagementEndpoint{
			ApiCredentials: ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address:  testURL,
			Insecure: true,
		}

		hash1, err := endpoint1.GetHash()
		require.NoError(t, err)

		hash2, err := endpoint2.GetHash()
		require.NoError(t, err)

		assert.NotEqual(t, hash1, hash2)
	})

	t.Run("hash changes with trust bundle", func(t *testing.T) {
		testURL, err := url.Parse("https://test.example.com")
		require.NoError(t, err)

		endpoint1 := ManagementEndpoint{
			ApiCredentials: ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address:               testURL,
			AdditionalTrustBundle: "bundle1",
		}

		endpoint2 := ManagementEndpoint{
			ApiCredentials: ApiCredentials{
				Username: "testuser",
				Password: "testpass",
			},
			Address:               testURL,
			AdditionalTrustBundle: "bundle2",
		}

		hash1, err := endpoint1.GetHash()
		require.NoError(t, err)

		hash2, err := endpoint2.GetHash()
		require.NoError(t, err)

		assert.NotEqual(t, hash1, hash2)
	})
}

func TestConstants(t *testing.T) {
	t.Run("CategoriesKey constant", func(t *testing.T) {
		assert.Equal(t, "categories", CategoriesKey)
	})

	t.Run("ErrNotFound error", func(t *testing.T) {
		assert.Equal(t, "environment key not found", ErrNotFound.Error())
	})
}

func TestEnvironmentInterface(t *testing.T) {
	// Test that the interface can be implemented
	t.Run("interface implementation", func(t *testing.T) {
		env := &testEnvironment{}

		topology := Topology{
			"region": "us-west-1",
			"zone":   "us-west-1a",
		}

		endpoint, err := env.GetManagementEndpoint(topology)
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)

		value, err := env.Get(topology, "test-key")
		assert.NoError(t, err)
		assert.Equal(t, "test-value", value)
	})
}

// testEnvironment is a test implementation of Environment
type testEnvironment struct{}

func (t *testEnvironment) GetManagementEndpoint(topology Topology) (*ManagementEndpoint, error) {
	testURL, _ := url.Parse("https://test.example.com")
	return &ManagementEndpoint{
		ApiCredentials: ApiCredentials{
			Username: "testuser",
			Password: "testpass",
		},
		Address: testURL,
	}, nil
}

func (t *testEnvironment) Get(topology Topology, key string) (interface{}, error) {
	return "test-value", nil
}

func TestProviderInterface(t *testing.T) {
	// Test that Provider interface extends Environment
	t.Run("interface composition", func(t *testing.T) {
		provider := &testProvider{}

		topology := Topology{
			"region": "us-west-1",
		}

		// Test Environment methods
		endpoint, err := provider.GetManagementEndpoint(topology)
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)

		value, err := provider.Get(topology, "test-key")
		assert.NoError(t, err)
		assert.Equal(t, "test-value", value)
	})
}

// testProvider is a test implementation of Provider
type testProvider struct{}

func (t *testProvider) GetManagementEndpoint(topology Topology) (*ManagementEndpoint, error) {
	testURL, _ := url.Parse("https://test.example.com")
	return &ManagementEndpoint{
		ApiCredentials: ApiCredentials{
			Username: "testuser",
			Password: "testpass",
		},
		Address: testURL,
	}, nil
}

func (t *testProvider) Get(topology Topology, key string) (interface{}, error) {
	return "test-value", nil
}
