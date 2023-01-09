package karbon

import (
	"testing"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/environment/providers/local"
	"github.com/stretchr/testify/require"
)

func testCredsFromEnv(t *testing.T) prismgoclient.Credentials {
	// NOTE: These environment variables are set to dummy values for the purpose of unit testing.
	// These environment variables need to be replaced with a real Prism Central endpoint and credentials
	// when recording the mocks. Once done with recording, please reset them back to the original dummy values.
	// and ensure that the actual credentials are not checked in to the repository.
	t.Setenv("NUTANIX_ENDPOINT", "prism-test.nutanix.com")
	t.Setenv("NUTANIX_PORT", "9440")
	t.Setenv("NUTANIX_USERNAME", "test")
	t.Setenv("NUTANIX_PASSWORD", "test")
	t.Setenv("NUTANIX_INSECURE", "false")
	provider, err := local.NewProvider().GetManagementEndpoint(nil)
	require.NoError(t, err)

	return prismgoclient.Credentials{
		URL:      provider.Address.String(),
		Username: provider.Username,
		Password: provider.Password,
		Insecure: provider.Insecure,
	}
}
