package karbon

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

// MetaOperations wrap the internal http client and provide the implementation for the MetaService interface
type MetaOperations struct {
	httpClient *internal.Client
}

// MetaService provides the interface for the karbon metadata e.g. Versions
// Karbon v2.1
type MetaService interface {
	GetVersion(ctx context.Context) (*MetaVersionResponse, error)
	GetSemanticVersion(ctx context.Context) (*MetaSemanticVersionResponse, error)
}

// GetVersion returns the karbon version
func (op MetaOperations) GetVersion(ctx context.Context) (*MetaVersionResponse, error) {
	path := "/v1-alpha.1/version"
	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	karbonMetaVersionResponse := new(MetaVersionResponse)
	if err != nil {
		return nil, err
	}
	return karbonMetaVersionResponse, op.httpClient.Do(ctx, req, karbonMetaVersionResponse)
}

// GetSemanticVersion is a wrapper on GetVersion and returns the karbon semantic version
func (op MetaOperations) GetSemanticVersion(ctx context.Context) (*MetaSemanticVersionResponse, error) {
	const expectedVersionLength int = 3
	metaSemanticVersionResponse := new(MetaSemanticVersionResponse)
	rawVersion, err := op.GetVersion(ctx)
	if err != nil {
		return nil, err
	}
	splitted := strings.Split(*rawVersion.Version, ".")
	if len(splitted) != expectedVersionLength {
		return nil, fmt.Errorf("expected karbon version to be consisting out of 3 elements but was %v", len(splitted))
	}

	major, err := strconv.Atoi(splitted[0])
	if err != nil {
		return nil, fmt.Errorf("could not convert major version to int64")
	}
	minor, err := strconv.Atoi(splitted[1])
	if err != nil {
		return nil, fmt.Errorf("could not convert minor version to int64")
	}
	rev, err := strconv.Atoi(splitted[2])
	if err != nil {
		return nil, fmt.Errorf("could not convert rev version to int64")
	}
	metaSemanticVersionResponse.MajorVersion = int64(major)
	metaSemanticVersionResponse.MinorVersion = int64(minor)
	metaSemanticVersionResponse.RevisionVersion = int64(rev)
	return metaSemanticVersionResponse, nil
}
