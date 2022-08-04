package foundation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nutanix-cloud-native/prism-go-client"
)

func TestNewFoundationAPIClient(t *testing.T) {
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
	foundationClient, err := NewFoundationAPIClient(cred)
	assert.NoError(t, err)

	outURL := fmt.Sprintf("http://%s:%s/", cred.FoundationEndpoint, cred.FoundationPort)
	assert.Equal(t, outURL, foundationClient.client.BaseURL.String())

	// verify missing client scenario
	cred = prismgoclient.Credentials{
		URL:      "foo.com",
		Username: "username",
		Password: "password",
		Port:     "",
		Endpoint: "0.0.0.0",
		Insecure: true,
		RequiredFields: map[string][]string{
			"foundation": {"foundation_endpoint"},
		},
	}
	foundationClient, err = NewFoundationAPIClient(cred)
	assert.Error(t, err)
	assert.Nil(t, foundationClient)
}
