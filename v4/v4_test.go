package v4

import (
	"testing"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/stretchr/testify/assert"
)

func TestNewV4Client(t *testing.T) {
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
