package foundationcentral

import (
	"testing"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
)

func TestNewFoundationCentralClient(t *testing.T) {
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
	_, err := NewFoundationCentralClient(cred)
	if err != nil {
		t.Error(err)
	}

	// verify missing client scenario
	cred2 := prismgoclient.Credentials{
		URL:      "foo.com",
		Insecure: true,
		RequiredFields: map[string][]string{
			"prism_central":      {"username", "password", "endpoint"},
			"foundation_central": {"username", "password", "endpoint"},
		},
	}
	FcClient2, err2 := NewFoundationCentralClient(cred2)
	if err2 != nil {
		t.Error(err2)
	}

	if FcClient2.client.ErrorMsg == "" {
		t.Errorf("NewFoundationCentralClient(%v) expected the base client in v3 client to have some error message", cred2)
	}
}
