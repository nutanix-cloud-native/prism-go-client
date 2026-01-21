package konnector

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

// MetaService provides the interface for the konnector metadata e.g. Versions
// Konnector v2.1
type MetaService interface {
	GetVersion(ctx context.Context) (*MetaVersionResponse, error)
	GetSemanticVersion(ctx context.Context) (*MetaSemanticVersionResponse, error)
}

// GetVersion returns the konnector version
func (op MetaOperations) GetVersion(ctx context.Context) (*MetaVersionResponse, error) {
	path := "/v1-alpha.1/version"
	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	konnectorMetaVersionResponse := new(MetaVersionResponse)
	if err != nil {
		return nil, err
	}
	return konnectorMetaVersionResponse, op.httpClient.Do(ctx, req, konnectorMetaVersionResponse)
}

// GetSemanticVersion is a wrapper on GetVersion and returns the konnector semantic version
func (op MetaOperations) GetSemanticVersion(ctx context.Context) (*MetaSemanticVersionResponse, error) {
	const expectedVersionLength int = 3
	metaSemanticVersionResponse := new(MetaSemanticVersionResponse)
	rawVersion, err := op.GetVersion(ctx)
	if err != nil {
		return nil, err
	}
	splitted := strings.Split(*rawVersion.Version, ".")
	if len(splitted) != expectedVersionLength {
		return nil, fmt.Errorf("expected konnector version to be consisting out of 3 elements but was %v", len(splitted))
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
