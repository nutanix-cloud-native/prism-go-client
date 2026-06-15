package credentials

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseCredentialsBasicAuth(t *testing.T) {
	credsData := []byte(`[
		{
			"type": "basic_auth",
			"data": {
				"prismCentral": {
					"username": "test-user",
					"password": "test-password"
				},
				"prismElements": null
			}
		}
	]`)

	creds, err := ParseCredentials(credsData)
	require.NoError(t, err)
	require.Equal(t, "test-user", creds.Username)
	require.Equal(t, "test-password", creds.Password)
	require.Empty(t, creds.APIKey)
}

func TestParseCredentialsAPIKey(t *testing.T) {
	credsData := []byte(`[
		{
			"type": "api_key",
			"data": {
				"prismCentral": {
					"apiKey": "test-api-key"
				}
			}
		}
	]`)

	creds, err := ParseCredentials(credsData)
	require.NoError(t, err)
	require.Empty(t, creds.Username)
	require.Empty(t, creds.Password)
	require.Equal(t, "test-api-key", creds.APIKey)
}

func TestParseCredentialsEmptyAPIKey(t *testing.T) {
	credsData := []byte(`[
		{
			"type": "api_key",
			"data": {
				"prismCentral": {
					"apiKey": ""
				}
			}
		}
	]`)

	_, err := ParseCredentials(credsData)
	require.ErrorContains(t, err, "the PrismCentral api key data is not set")
}
