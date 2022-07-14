package karbon

import (
	"testing"

	"github.com/nutanix-cloud-native/prism-go-client"
)

func TestNewKarbonAPIClient(t *testing.T) {
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
	_, err := NewKarbonAPIClient(cred)
	if err != nil {
		t.Errorf(err.Error())
	}

	// verify missing client scenario
	cred2 := prismgoclient.Credentials{
		URL:      "foo.com",
		Insecure: true,
		RequiredFields: map[string][]string{
			"karbon": {"username", "password", "endpoint"},
		},
	}
	v3Client2, err2 := NewKarbonAPIClient(cred2)
	if err2 != nil {
		t.Errorf(err2.Error())
	}

	if v3Client2.client.ErrorMsg == "" {
		t.Errorf("NewKarbonAPIClient(%v) expected the base client in karbon client to have some error message", cred2)
	}
}
