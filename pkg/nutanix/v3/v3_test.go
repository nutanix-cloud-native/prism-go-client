package v3

import (
	"testing"

	client "github.com/nutanix-cloud-native/prism-go-client/pkg/nutanix"
)

func TestNewV3Client(t *testing.T) {
	// verifies positive client creation
	cred := client.Credentials{
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
	_, err := NewV3Client(cred)
	if err != nil {
		t.Errorf(err.Error())
	}

	// verify missing client scenario
	cred2 := client.Credentials{
		URL:      "foo.com",
		Insecure: true,
		RequiredFields: map[string][]string{
			"prism_central": {"username", "password", "endpoint"},
		},
	}
	v3Client2, err2 := NewV3Client(cred2)
	if err2 != nil {
		t.Errorf(err2.Error())
	}

	if v3Client2.client.ErrorMsg == "" {
		t.Errorf("NewV3Client(%v) expected the base client in v3 client to have some error message", cred2)
	}
}
