//go:build e2e

package e2e

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/converged"
	v4 "github.com/nutanix-cloud-native/prism-go-client/converged/v4"
)

func TestIncorrectUsernameE2E(t *testing.T) {
	credentials := CredentialsFromE2EEnvironment(t)
	credentials.Username += "-incorrect"

	assertUnauthenticatedCredentials(t, credentials)
}

func TestIncorrectPasswordE2E(t *testing.T) {
	credentials := CredentialsFromE2EEnvironment(t)
	credentials.Password += "-incorrect"

	assertUnauthenticatedCredentials(t, credentials)
}

func assertUnauthenticatedCredentials(t *testing.T, credentials prismgoclient.Credentials) {
	t.Helper()

	for _, operation := range []string{"get", "list", "iterator"} {
		t.Run(operation, func(t *testing.T) {
			client, err := v4.NewClient(credentials)
			require.NoError(t, err)

			ctx := context.Background()
			switch operation {
			case "get":
				_, err = client.Clusters.Get(ctx, "00000000-0000-0000-0000-000000000000")
			case "list":
				_, err = client.Clusters.List(ctx, converged.WithLimit(1))
			case "iterator":
				iterator := client.Clusters.NewIterator(ctx, converged.WithLimit(1))
				for _, iteratorErr := range iterator {
					err = iteratorErr
					break
				}
			}

			require.Error(t, err)
			require.Truef(t, converged.IsUnauthenticated(err), "expected unauthenticated error, got: %v", err)
		})
	}
}
