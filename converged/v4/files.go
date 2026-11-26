package v4

import (
	"context"
	"fmt"

	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
)

// FilesNFSVersion is the NFS protocol version configured on a Nutanix Files server.
type FilesNFSVersion = v4prismGoClient.FilesNFSVersion

const (
	FilesNFSVersionUnknown = v4prismGoClient.FilesNFSVersionUnknown
	FilesNFSVersionV3      = v4prismGoClient.FilesNFSVersionV3
	FilesNFSVersionV4      = v4prismGoClient.FilesNFSVersionV4
	FilesNFSVersionV3V4    = v4prismGoClient.FilesNFSVersionV3V4
)

// FilesNFSVersion reads NFS version directly from the file server configured in Credentials.
func (c *Client) FilesNFSVersion(ctx context.Context) (FilesNFSVersion, error) {
	if c == nil || c.client == nil {
		return FilesNFSVersionUnknown, fmt.Errorf("client is not initialized")
	}
	return c.client.FilesNFSVersion(ctx)
}

// FilesNFSVersionFromPC resolves fileServer through Prism Central and reads its NFS version.
// fileServer can be a file server name, external ID, FQDN, external IP, or virtual IP.
func (c *Client) FilesNFSVersionFromPC(ctx context.Context, fileServer string) (FilesNFSVersion, error) {
	if c == nil || c.client == nil {
		return FilesNFSVersionUnknown, fmt.Errorf("client is not initialized")
	}
	return c.client.FilesNFSVersionFromPC(ctx, fileServer)
}
