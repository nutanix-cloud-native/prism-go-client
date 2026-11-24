package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	managementModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/management"
	domainManagerRequest "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/request/domainmanager"
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
			return s.client.DomainManagerApiInstance.ServiceClient.GetDomainManagerById(ctx,
				&domainManagerRequest.GetDomainManagerByIdRequest{
					ExtId: &extID,
				},
			)
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
			return s.client.DomainManagerApiInstance.ServiceClient.ListDomainManagers(ctx,
				&domainManagerRequest.ListDomainManagersRequest{
					Select_: reqParams.Select,
				},
			)
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
			return s.client.DomainManagerApiInstance.ServiceClient.ListDomainManagers(ctx,
				&domainManagerRequest.ListDomainManagersRequest{
					Select_: reqParams.Select,
				},
			)
		},
		opts,
		s.entitiesName,
	)
}

// ListProducts returns the products registered on the domain manager identified
// by domainManagerExtID. Additional OData options (e.g. converged.WithFilter) may
// be supplied; for example, to filter by product name:
//
//	converged.WithFilter("name eq Prism.Management.ProductName'NUTANIX_DISASTER_RECOVERY'")
func (s *DomainManagerService) ListProducts(ctx context.Context, domainManagerExtID string, opts ...converged.ODataOption) ([]managementModels.Product, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericListEntities[*managementModels.ListProductsApiResponse, managementModels.Product](
		func(reqParams *V4ODataParams) (*managementModels.ListProductsApiResponse, error) {
			return s.client.DomainManagerApiInstance.ServiceClient.ListProducts(ctx,
				&domainManagerRequest.ListProductsRequest{
					DomainManagerExtId: &domainManagerExtID,
					Page_:              reqParams.Page,
					Limit_:             reqParams.Limit,
					Filter_:            reqParams.Filter,
					Orderby_:           reqParams.OrderBy,
					Select_:            reqParams.Select,
				},
			)
		},
		opts,
		"domain manager product",
	)
}

// GetProduct returns a single product (identified by extID) registered on the
// domain manager identified by domainManagerExtID.
func (s *DomainManagerService) GetProduct(ctx context.Context, domainManagerExtID, extID string) (*managementModels.Product, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*managementModels.GetProductApiResponse, managementModels.Product](
		func() (*managementModels.GetProductApiResponse, error) {
			return s.client.DomainManagerApiInstance.ServiceClient.GetProductById(ctx,
				&domainManagerRequest.GetProductByIdRequest{
					DomainManagerExtId: &domainManagerExtID,
					ExtId:              &extID,
				},
			)
		},
		"domain manager product",
	)
}

// IsProductDisabled reports whether the product with the given productName is in
// the DISABLED enablement state on the domain manager identified by
// domainManagerExtID.
//
// It performs the two-step lookup documented by the API: it lists products
// filtered by name to resolve the product's extId, then fetches that product
// and inspects its enablementState.
//
// If no product matching productName is found, found is false and err is nil.
func (s *DomainManagerService) IsProductDisabled(ctx context.Context, domainManagerExtID string, productName managementModels.ProductName) (disabled bool, found bool, err error) {
	if s.client == nil {
		return false, false, errors.New("client is not initialized")
	}

	filter := fmt.Sprintf("name eq Prism.Management.ProductName'%s'", productName.GetName())
	products, err := s.ListProducts(ctx, domainManagerExtID, converged.WithFilter(filter))
	if err != nil {
		return false, false, fmt.Errorf("failed to list products: %w", err)
	}

	if len(products) == 0 {
		return false, false, nil
	}

	if products[0].ExtId == nil {
		return false, false, fmt.Errorf("product %q is missing an extId", productName.GetName())
	}

	product, err := s.GetProduct(ctx, domainManagerExtID, *products[0].ExtId)
	if err != nil {
		return false, false, fmt.Errorf("failed to get product %q: %w", productName.GetName(), err)
	}

	if product.EnablementState == nil {
		return false, true, fmt.Errorf("product %q has no enablementState", productName.GetName())
	}

	return *product.EnablementState == managementModels.ENABLEMENTSTATE_DISABLED, true, nil
}

// GetPrismCentralVersion gets the Prism Central version from Domain Manager API
// This is a lightweight alternative to V3 GetPrismCentral() API
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
