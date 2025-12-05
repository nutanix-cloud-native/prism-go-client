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
	// 1) Explicit APIKey
	{
		ac := &fakeApiClientV4Test{}
		creds := prismgoclient.Credentials{
			APIKey: "explicit-api-key",
		}
		setAuthHeader(ac, creds)
		assert.Equal(t, "explicit-api-key", ac.headers[ntnxAPIKeyHeaderKey])
		_, hasAuth := ac.headers[authorizationHeader]
		assert.False(t, hasAuth)
	}

	// 2) Username equals API key header key, Password is API key
	{
		ac := &fakeApiClientV4Test{}
		creds := prismgoclient.Credentials{
			Username: ntnxAPIKeyHeaderKey,
			Password: "password-as-api-key",
		}
		setAuthHeader(ac, creds)
		assert.Equal(t, "password-as-api-key", ac.headers[ntnxAPIKeyHeaderKey])
		_, hasAuth := ac.headers[authorizationHeader]
		assert.False(t, hasAuth)
	}

	// 3) Normal basic auth fallback
	{
		ac := &fakeApiClientV4Test{}
		creds := prismgoclient.Credentials{
			Username: "user",
			Password: "pass",
		}
		setAuthHeader(ac, creds)
		_, hasApiKey := ac.headers[ntnxAPIKeyHeaderKey]
		assert.False(t, hasApiKey)
		val, hasAuth := ac.headers[authorizationHeader]
		assert.True(t, hasAuth)
		assert.Contains(t, val, "Basic ")
	}
}
