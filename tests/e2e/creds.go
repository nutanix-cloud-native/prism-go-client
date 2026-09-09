package e2e

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
)

func requireE2EEnv(t testing.TB, name string) string {
	t.Helper()
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		t.Fatalf("missing required e2e environment variable %q; required: %v", name, RequiredNutanixE2EEnvVars())
	}
	return value
}

func ManagementEndpointFromE2EEnvironment(t testing.TB) *types.ManagementEndpoint {
	t.Helper()
	endpoint := requireE2EEnv(t, EnvNutanixE2EEndpoint)
	port := requireE2EEnv(t, EnvNutanixE2EPort)
	username := requireE2EEnv(t, EnvNutanixE2EUsername)
	password := requireE2EEnv(t, EnvNutanixE2EPassword)
	insecureRaw := requireE2EEnv(t, EnvNutanixE2EInsecure)
	insecure, err := strconv.ParseBool(insecureRaw)
	if err != nil {
		t.Fatalf("invalid %s=%q: %v", EnvNutanixE2EInsecure, insecureRaw, err)
	}
	address, err := url.Parse(fmt.Sprintf("https://%s:%s", endpoint, port))
	if err != nil {
		t.Fatalf("invalid e2e endpoint URL: %v", err)
	}
	return &types.ManagementEndpoint{
		Address:        address,
		ApiCredentials: types.ApiCredentials{Username: username, Password: password},
		Insecure:       insecure,
	}
}

func CredentialsFromE2EEnvironment(t testing.TB) prismgoclient.Credentials {
	t.Helper()
	endpoint := ManagementEndpointFromE2EEnvironment(t)
	return prismgoclient.Credentials{
		Endpoint: endpoint.Address.Hostname(),
		Port:     endpoint.Address.Port(),
		URL:      endpoint.Address.String(),
		Username: endpoint.Username,
		Password: endpoint.Password,
		Insecure: endpoint.Insecure,
	}
}
