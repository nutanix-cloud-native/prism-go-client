package karbon

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nutanix-cloud-native/prism-go-client"
)

func TestNewKarbonAPIClient(t *testing.T) {
	// verifies positive httpClient creation
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
	_, err := NewKarbonAPIClient(cred)
	assert.NoError(t, err)

	// verify missing httpClient scenario
	cred = prismgoclient.Credentials{
		URL:      "foo.com",
		Insecure: true,
		RequiredFields: map[string][]string{
			"karbon": {"username", "password", "endpoint"},
		},
	}
	_, err = NewKarbonAPIClient(cred)
	assert.Error(t, err)
}
