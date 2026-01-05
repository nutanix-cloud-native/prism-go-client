package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

// DomainManagerService provides methods to interact with Domain Manager API
type DomainManagerService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewDomainManagerService creates a new DomainManagerService instance
func NewDomainManagerService(client *v4prismGoClient.Client) *DomainManagerService {
	return &DomainManagerService{client: client, entitiesName: "domain manager"}
}

// Get returns a single DomainManager entity by its external ID.
func (s *DomainManagerService) Get(ctx context.Context, extID string) (*prismModels.DomainManager, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*prismModels.GetDomainManagerApiResponse, prismModels.DomainManager](
		func() (*prismModels.GetDomainManagerApiResponse, error) {
			return s.client.DomainManagerApiInstance.GetDomainManagerById(&extID)
		},
		s.entitiesName,
	)
}

// List returns a list of DomainManager entities.
func (s *DomainManagerService) List(ctx context.Context, opts ...converged.ODataOption) ([]prismModels.DomainManager, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	// Domain Manager API only supports select parameter, not page, limit, filter, orderby, expand, apply
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	// Check for unsupported options
	if reqParams.Page != nil || reqParams.Limit != nil || reqParams.Filter != nil || reqParams.OrderBy != nil || reqParams.Expand != nil || reqParams.Apply != nil {
		return nil, fmt.Errorf("domain manager API only supports select parameter, not page, limit, filter, orderby, expand, or apply")
	}

	return GenericListEntities[*prismModels.ListDomainManagerApiResponse, prismModels.DomainManager](
		func(reqParams *V4ODataParams) (*prismModels.ListDomainManagerApiResponse, error) {
			return s.client.DomainManagerApiInstance.ListDomainManagers(reqParams.Select)
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing DomainManager entities.
func (s *DomainManagerService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[prismModels.DomainManager] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*prismModels.ListDomainManagerApiResponse, prismModels.DomainManager](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*prismModels.ListDomainManagerApiResponse, error) {
			return s.client.DomainManagerApiInstance.ListDomainManagers(reqParams.Select)
		},
		opts,
		s.entitiesName,
	)
}

// GetPrismCentralVersion gets the Prism Central version from Domain Manager API
// This is a lightweight alternative to V3 GetPrismCentral() API
// Endpoint: GET /prism/v4.2/domain-managers
func (s *DomainManagerService) GetPrismCentralVersion(ctx context.Context) (string, error) {
	if s.client == nil {
		return "", errors.New("client is not initialized")
	}

	// Use List to get domain managers, then extract version from first one
	domainManagers, err := s.List(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to list domain managers: %w", err)
	}

	if len(domainManagers) == 0 {
		return "", fmt.Errorf("no domain managers found")
	}

	// Extract version from BuildInfo
	if domainManagers[0].Config == nil || domainManagers[0].Config.BuildInfo == nil {
		return "", fmt.Errorf("version not found in domain manager response")
	}

	buildInfo := domainManagers[0].Config.BuildInfo
	if buildInfo.Version == nil || *buildInfo.Version == "" {
		return "", fmt.Errorf("version not found in domain manager response")
	}

	return *buildInfo.Version, nil
}
