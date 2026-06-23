package credentials

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseCredentials(t *testing.T) {
	tests := []struct {
		name        string
		credsData   []byte
		wantErr     bool
		errContains string
		wantBasic   *BasicAuthCredential
		wantAPIKey  *APIKeyCredential
	}{
		{
			name:        "invalid top-level JSON",
			credsData:   []byte(`not-json`),
			wantErr:     true,
			errContains: "failed to unmarshal the credentials data",
		},
		{
			name:        "empty credentials list",
			credsData:   []byte(`[]`),
			wantErr:     true,
			errContains: "no credentials found",
		},
		{
			name: "unsupported credential type",
			credsData: []byte(`[
				{"type":"token","data":{"foo":"bar"}}
			]`),
			wantErr:     true,
			errContains: "unsupported credentials type",
		},
		{
			name: "basic auth success",
			credsData: []byte(`[
				{
					"type":"basic_auth",
					"data":{
						"prismCentral":{"username":"test-user","password":"test-password"},
						"prismElements":null
					}
				}
			]`),
			wantBasic: &BasicAuthCredential{
				PrismCentral: PrismCentralBasicAuth{
					BasicAuth: BasicAuth{
						Username: "test-user",
						Password: "test-password",
					},
				},
			},
		},
		{
			name: "basic auth data unmarshal failure",
			credsData: []byte(`[
				{
					"type":"basic_auth",
					"data":"not-an-object"
				}
			]`),
			wantErr:     true,
			errContains: "failed to unmarshal the basic-auth data",
		},
		{
			name: "basic auth missing username",
			credsData: []byte(`[
				{
					"type":"basic_auth",
					"data":{
						"prismCentral":{"username":"","password":"test-password"},
						"prismElements":null
					}
				}
			]`),
			wantErr:     true,
			errContains: "either username and password, or an API key must be set in provider configuration",
		},
		{
			name: "basic auth missing password",
			credsData: []byte(`[
				{
					"type":"basic_auth",
					"data":{
						"prismCentral":{"username":"test-user","password":""},
						"prismElements":null
					}
				}
			]`),
			wantErr:     true,
			errContains: "either username and password, or an API key must be set in provider configuration",
		},
		{
			name: "api key success",
			credsData: []byte(`[
				{
					"type":"api_key",
					"data":{"prismCentral":{"apiKey":"test-api-key"}}
				}
			]`),
			wantAPIKey: &APIKeyCredential{
				PrismCentral: PrismCentralAPIKey{
					APIKey: APIKey("test-api-key"),
				},
			},
		},
		{
			name: "api key data unmarshal failure",
			credsData: []byte(`[
				{
					"type":"api_key",
					"data":"not-an-object"
				}
			]`),
			wantErr:     true,
			errContains: "failed to unmarshal the api key data",
		},
		{
			name: "api key empty value",
			credsData: []byte(`[
				{
					"type":"api_key",
					"data":{"prismCentral":{"apiKey":""}}
				}
			]`),
			wantErr:     true,
			errContains: "either username and password, or an API key must be set in provider configuration",
		},
		{
			name: "multiple credentials uses first only",
			credsData: []byte(`[
				{
					"type":"basic_auth",
					"data":{
						"prismCentral":{"username":"first-user","password":"first-pass"},
						"prismElements":null
					}
				},
				{
					"type":"api_key",
					"data":{"prismCentral":{"apiKey":"second-api-key"}}
				}
			]`),
			wantBasic: &BasicAuthCredential{
				PrismCentral: PrismCentralBasicAuth{
					BasicAuth: BasicAuth{
						Username: "first-user",
						Password: "first-pass",
					},
				},
			},
		},
		{
			name: "multiple credentials first invalid fails even if second valid",
			credsData: []byte(`[
				{
					"type":"basic_auth",
					"data":{
						"prismCentral":{"username":"first-user","password":""},
						"prismElements":null
					}
				},
				{
					"type":"api_key",
					"data":{"prismCentral":{"apiKey":"second-api-key"}}
				}
			]`),
			wantErr:     true,
			errContains: "either username and password, or an API key must be set in provider configuration",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds, err := ParseCredentials(tt.credsData)
			if tt.wantErr {
				require.Error(t, err)
				require.ErrorContains(t, err, tt.errContains)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, creds)
			if tt.wantBasic != nil {
				require.Equal(t, tt.wantBasic.PrismCentral.Username, creds.Username)
				require.Equal(t, tt.wantBasic.PrismCentral.Password, creds.Password)
				require.Empty(t, creds.APIKey)
				return
			}
			if tt.wantAPIKey != nil {
				require.Empty(t, creds.Username)
				require.Empty(t, creds.Password)
				require.Equal(t, string(tt.wantAPIKey.PrismCentral.APIKey), creds.APIKey)
				return
			}
			require.Fail(t, "either wantBasic or wantAPIKey must be set for success cases")
		})
	}
}
