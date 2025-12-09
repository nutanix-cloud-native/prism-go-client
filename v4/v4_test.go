package v4

import (
	"testing"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/stretchr/testify/assert"
)

func TestNewV4ClientBasicAuth(t *testing.T) {
	// verifies positive client creation
	cred := prismgoclient.Credentials{
		URL:                "foo.com",
		Username:           "username",
		Password:           "password",
		Port:               "",
		Endpoint:           "0.0.0.0",
		Insecure:           true,
		FoundationEndpoint: "10.0.0.0",
		FoundationPort:     "8000",
		RequiredFields:     nil,
	}
	v4Client, err := NewV4Client(cred)
	assert.NoError(t, err)
	assert.NotNil(t, v4Client)
	assert.NotNil(t, v4Client.VmApiInstance)
	assert.NotNil(t, v4Client.ImagesApiInstance)
	assert.NotNil(t, v4Client.SubnetsApiInstance)
	assert.NotNil(t, v4Client.SubnetIPReservationApi)
	assert.NotNil(t, v4Client.ClustersApiInstance)
	assert.NotNil(t, v4Client.TasksApiInstance)
	assert.NotNil(t, v4Client.StorageContainerAPI)
	assert.NotNil(t, v4Client.CategoriesApiInstance)
	assert.NotNil(t, v4Client.VolumeGroupsApiInstance)

	// verify missing client scenario
	cred = prismgoclient.Credentials{
		URL:      "foo.com",
		Insecure: true,
		RequiredFields: map[string][]string{
			"prism_central": {"username", "password", "endpoint"},
		},
	}

	v4Client, err = NewV4Client(cred)
	assert.Nil(t, v4Client)
	assert.EqualError(t, err, "username, password and endpoint are required for basic auth")
}

func TestNewV4Client(t *testing.T) {
	// verifies positive client creation
	cred := prismgoclient.Credentials{
		URL:                "foo.com",
		APIKey:             "my-api-key",
		Port:               "",
		Endpoint:           "0.0.0.0",
		Insecure:           true,
		FoundationEndpoint: "10.0.0.0",
		FoundationPort:     "8000",
		RequiredFields:     nil,
	}
	v4Client, err := NewV4Client(cred)
	assert.NoError(t, err)
	assert.NotNil(t, v4Client)
	assert.NotNil(t, v4Client.VmApiInstance)
	assert.NotNil(t, v4Client.ImagesApiInstance)
	assert.NotNil(t, v4Client.SubnetsApiInstance)
	assert.NotNil(t, v4Client.SubnetIPReservationApi)
	assert.NotNil(t, v4Client.ClustersApiInstance)
	assert.NotNil(t, v4Client.TasksApiInstance)
	assert.NotNil(t, v4Client.StorageContainerAPI)
	assert.NotNil(t, v4Client.CategoriesApiInstance)
	assert.NotNil(t, v4Client.VolumeGroupsApiInstance)

	// verify missing client scenario
	cred = prismgoclient.Credentials{
		URL:      "foo.com",
		Insecure: true,
		RequiredFields: map[string][]string{
			"prism_central": {"username", "password", "endpoint"},
		},
		APIKey: "my-api-key",
	}

	v4Client, err = NewV4Client(cred)
	assert.Nil(t, v4Client)
	assert.EqualError(t, err, "endpoint is required for api key auth")
}

type fakeApiClientV4Test struct {
	headers map[string]string
}

func (f *fakeApiClientV4Test) AddDefaultHeader(key, value string) {
	if f.headers == nil {
		f.headers = make(map[string]string)
	}
	f.headers[key] = value
}

func Test_setAuthHeader_ServiceAccountModes_V4(t *testing.T) {
	tests := []struct {
		name               string
		creds              prismgoclient.Credentials
		expectedAPIKey     string
		expectAuth         bool
		expectedAuthPrefix string
	}{
		{
			name: "explicit API key",
			creds: prismgoclient.Credentials{
				APIKey: "explicit-api-key",
			},
			expectedAPIKey: "explicit-api-key",
			expectAuth:     false,
		},
		{
			name: "username is header key, password is API key",
			creds: prismgoclient.Credentials{
				Username: ntnxAPIKeyHeaderKey,
				Password: "password-as-api-key",
			},
			expectedAPIKey: "password-as-api-key",
			expectAuth:     false,
		},
		{
			name: "basic auth fallback",
			creds: prismgoclient.Credentials{
				Username: "user",
				Password: "pass",
			},
			expectedAPIKey:     "",
			expectAuth:         true,
			expectedAuthPrefix: "Basic ",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ac := &fakeApiClientV4Test{}
			setAuthHeader(ac, tt.creds)

			if tt.expectedAPIKey != "" {
				assert.Equal(t, tt.expectedAPIKey, ac.headers[ntnxAPIKeyHeaderKey])
			} else {
				_, hasAPIKey := ac.headers[ntnxAPIKeyHeaderKey]
				assert.False(t, hasAPIKey)
			}

			if tt.expectAuth {
				val, hasAuth := ac.headers[authorizationHeader]
				assert.True(t, hasAuth)
				if tt.expectedAuthPrefix != "" {
					assert.Contains(t, val, tt.expectedAuthPrefix)
				}
			} else {
				_, hasAuth := ac.headers[authorizationHeader]
				assert.False(t, hasAuth)
			}
		})
	}
}
