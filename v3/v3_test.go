package v3

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nutanix-cloud-native/prism-go-client"
)

func TestNewV3Client(t *testing.T) {
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
	v3Client, err := NewV3Client(cred)
	assert.NoError(t, err)
	assert.NotNil(t, v3Client)

	// verify missing client scenario
	cred = prismgoclient.Credentials{
		URL:      "foo.com",
		Insecure: true,
		RequiredFields: map[string][]string{
			"prism_central": {"username", "password", "endpoint"},
		},
	}

	v3Client, err = NewV3Client(cred)
	assert.Nil(t, v3Client)
	assert.EqualError(t, err, "username, password and endpoint are required")
}
