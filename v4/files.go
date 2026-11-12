package v4

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	filesConfig "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/config"
)

const (
	filesPCFileServersPath    = "/api/files/v4.0.a2/config/file-servers"
	filesPCNameServicesFormat = "/api/files/v4.0.a2/config/file-servers/%s/name-services"
	filesPCListPageSize       = 100
	filesPCListMaxPages       = 100
)

// FilesNFSVersion is the NFS protocol version configured on a Nutanix Files server.
type FilesNFSVersion string

const (
	FilesNFSVersionUnknown FilesNFSVersion = "UNKNOWN"
	FilesNFSVersionV3      FilesNFSVersion = "NFSV3"
	FilesNFSVersionV4      FilesNFSVersion = "NFSV4"
	FilesNFSVersionV3V4    FilesNFSVersion = "NFSV3V4"
)

type filesNameServicesResponse struct {
	Data struct {
		ObjectType string `json:"$objectType"`
		NFSVersion string `json:"nfsVersion"`
	} `json:"data"`
}

type filesPCFileServer struct {
	ExtID            string `json:"extId"`
	Name             string `json:"name"`
	DNSDomainName    string `json:"dnsDomainName"`
	ExternalNetworks []struct {
		IPAddresses      []filesIPAddress `json:"ipAddresses"`
		VirtualIPAddress *filesIPAddress  `json:"virtualIpAddress"`
	} `json:"externalNetworks"`
}

type filesIPAddress struct {
	IPv4 *struct {
		Value string `json:"value"`
	} `json:"ipv4"`
	IPv6 *struct {
		Value string `json:"value"`
	} `json:"ipv6"`
}

type filesPCFileServersResponse struct {
	Data     []filesPCFileServer `json:"data"`
	Metadata struct {
		TotalAvailableResults *int `json:"totalAvailableResults"`
	} `json:"metadata"`
}

// FilesNFSVersion reads NFS version directly from the file server configured in Credentials.
func (c *Client) FilesNFSVersion(ctx context.Context) (FilesNFSVersion, error) {
	if c == nil || c.filesNameServicesAPI == nil {
		return FilesNFSVersionUnknown, fmt.Errorf("Files client is not initialized")
	}
	if err := ctx.Err(); err != nil {
		return FilesNFSVersionUnknown, err
	}
	// Files SDK v4.0.1-alpha.2 does not accept a context. This check handles an already-cancelled
	// context; an in-flight call uses the SDK client's 30-second default timeout.
	response, err := c.filesNameServicesAPI.GetNameServicesByFileServer()
	if err != nil {
		return FilesNFSVersionUnknown, err
	}
	if response == nil {
		return FilesNFSVersionUnknown, fmt.Errorf("Files name-services response is empty")
	}
	nameServices, ok := response.GetData().(filesConfig.NameServiceSpec)
	if !ok || nameServices.NfsVersion == nil {
		return FilesNFSVersionUnknown, fmt.Errorf("Files name-services response has invalid data")
	}
	return filesNFSVersionFromSDK(*nameServices.NfsVersion)
}

// FilesNFSVersionFromPC resolves fileServer through Prism Central and reads its NFS version.
// fileServer can be a file server name, external ID, FQDN, external IP, or virtual IP.
func (c *Client) FilesNFSVersionFromPC(ctx context.Context, fileServer string) (FilesNFSVersion, error) {
	if c == nil || c.filesAPIClient == nil {
		return FilesNFSVersionUnknown, fmt.Errorf("Files client is not initialized")
	}
	if err := ctx.Err(); err != nil {
		return FilesNFSVersionUnknown, err
	}
	extID, err := c.fileServerExtID(fileServer)
	if err != nil {
		return FilesNFSVersionUnknown, err
	}
	return c.nameServicesFromPC(fmt.Sprintf(filesPCNameServicesFormat, url.PathEscape(extID)))
}

func filesNFSVersionFromSDK(version filesConfig.NfsVersion) (FilesNFSVersion, error) {
	switch version {
	case filesConfig.NFSVERSION_NFSV3:
		return FilesNFSVersionV3, nil
	case filesConfig.NFSVERSION_NFSV4:
		return FilesNFSVersionV4, nil
	case filesConfig.NFSVERSION_NFSV3V4:
		return FilesNFSVersionV3V4, nil
	default:
		return FilesNFSVersionUnknown, fmt.Errorf("Files name-services response has no usable NFS version")
	}
}

func (c *Client) nameServicesFromPC(path string) (FilesNFSVersion, error) {
	body, err := c.callFilesAPI(path, nil)
	if err != nil {
		return FilesNFSVersionUnknown, err
	}
	var response filesNameServicesResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return FilesNFSVersionUnknown, fmt.Errorf("failed to decode Files name-services response: %w", err)
	}
	if response.Data.ObjectType != "files.v4.config.NameServiceSpec" {
		return FilesNFSVersionUnknown, fmt.Errorf("Files name-services response has invalid data type %q", response.Data.ObjectType)
	}
	switch FilesNFSVersion(response.Data.NFSVersion) {
	case FilesNFSVersionV3, FilesNFSVersionV4, FilesNFSVersionV3V4:
		return FilesNFSVersion(response.Data.NFSVersion), nil
	default:
		return FilesNFSVersionUnknown, fmt.Errorf("Files name-services response has no usable NFS version")
	}
}

func (c *Client) fileServerExtID(identifier string) (string, error) {
	seen := 0
	for page := range filesPCListMaxPages {
		query := url.Values{
			"$page":  {strconv.Itoa(page)},
			"$limit": {strconv.Itoa(filesPCListPageSize)},
		}
		body, err := c.callFilesAPI(filesPCFileServersPath, query)
		if err != nil {
			return "", err
		}
		if len(body) == 0 {
			return "", fmt.Errorf("empty response listing Nutanix Files servers")
		}

		var response filesPCFileServersResponse
		if err := json.Unmarshal(body, &response); err != nil {
			return "", fmt.Errorf("failed to decode Nutanix Files servers list: %w", err)
		}
		for i := range response.Data {
			server := &response.Data[i]
			if server.matches(identifier) {
				if server.ExtID == "" {
					return "", fmt.Errorf("file server %q has no external ID", identifier)
				}
				return server.ExtID, nil
			}
		}

		seen += len(response.Data)
		switch {
		case len(response.Data) == 0:
			return "", fmt.Errorf("file server %q not found on Prism Central", identifier)
		case response.Metadata.TotalAvailableResults != nil && seen >= *response.Metadata.TotalAvailableResults:
			return "", fmt.Errorf("file server %q not found on Prism Central", identifier)
		case response.Metadata.TotalAvailableResults == nil && len(response.Data) < filesPCListPageSize:
			return "", fmt.Errorf("file server %q not found on Prism Central", identifier)
		}
	}
	return "", fmt.Errorf("file server %q not found within first %d Prism Central file servers", identifier, filesPCListMaxPages*filesPCListPageSize)
}

func (c *Client) callFilesAPI(path string, query url.Values) ([]byte, error) {
	body, err := c.filesAPIClient.CallApi(
		&path,
		http.MethodGet,
		nil,
		query,
		make(map[string]string),
		url.Values{},
		[]string{"application/json"},
		nil,
		[]string{"basicAuthScheme"},
	)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (s *filesPCFileServer) matches(identifier string) bool {
	id := normalizeFilesIdentifier(identifier)
	if id == "" {
		return false
	}
	if strings.EqualFold(s.Name, id) || strings.EqualFold(s.ExtID, id) {
		return true
	}
	name := strings.TrimSuffix(strings.ToLower(strings.TrimSpace(s.Name)), ".")
	domain := strings.TrimSuffix(strings.ToLower(strings.TrimSpace(s.DNSDomainName)), ".")
	if name != "" && domain != "" && id == name+"."+domain {
		return true
	}
	for i := range s.ExternalNetworks {
		if filesIPAddressMatches(s.ExternalNetworks[i].VirtualIPAddress, id) {
			return true
		}
		for j := range s.ExternalNetworks[i].IPAddresses {
			if filesIPAddressMatches(&s.ExternalNetworks[i].IPAddresses[j], id) {
				return true
			}
		}
	}
	return false
}

func normalizeFilesIdentifier(identifier string) string {
	identifier = strings.TrimSpace(strings.ToLower(identifier))
	if identifier == "" {
		return ""
	}
	parsed, err := url.Parse("https://" + identifier)
	if err == nil && parsed.Hostname() != "" {
		return strings.TrimSuffix(parsed.Hostname(), ".")
	}
	return strings.TrimSuffix(identifier, ".")
}

func filesIPAddressMatches(address *filesIPAddress, identifier string) bool {
	if address == nil {
		return false
	}
	return (address.IPv4 != nil && strings.EqualFold(strings.TrimSpace(address.IPv4.Value), identifier)) ||
		(address.IPv6 != nil && strings.EqualFold(strings.TrimSpace(address.IPv6.Value), identifier))
}
