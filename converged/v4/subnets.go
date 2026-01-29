package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	subnetModels "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	networkingprismapi "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/prism/v4/config"
)

// SubnetsService implements the converged Subnets interface for the Prism Central V4 API.
type SubnetsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewSubnetsService returns a new SubnetsService for the given V4 client.
func NewSubnetsService(client *v4prismGoClient.Client) *SubnetsService {
	return &SubnetsService{client: client, entitiesName: "subnet"}
}

// Get returns the subnet for the given UUID.
func (s *SubnetsService) Get(ctx context.Context, uuid string) (*subnetModels.Subnet, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*subnetModels.GetSubnetApiResponse, subnetModels.Subnet](
		func() (*subnetModels.GetSubnetApiResponse, error) {
			return s.client.SubnetsApiInstance.GetSubnetById(&uuid)
		},
		s.entitiesName,
	)
}

// List returns a list of subnets.
func (s *SubnetsService) List(ctx context.Context, opts ...converged.ODataOption) ([]subnetModels.Subnet, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams.Apply != nil {
		return nil, fmt.Errorf("apply is not supported for listing Subnets")
	}

	return GenericListEntities[*subnetModels.ListSubnetsApiResponse, subnetModels.Subnet](
		func(reqParams *V4ODataParams) (*subnetModels.ListSubnetsApiResponse, error) {
			return s.client.SubnetsApiInstance.ListSubnets(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing subnets.
func (s *SubnetsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[subnetModels.Subnet] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*subnetModels.ListSubnetsApiResponse, subnetModels.Subnet](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*subnetModels.ListSubnetsApiResponse, error) {
			return s.client.SubnetsApiInstance.ListSubnets(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// ReserveIpsBySubnetId reserves IP addresses on a subnet.
// It returns a TaskReference that can be used to track the async operation.
func (s *SubnetsService) ReserveIpsBySubnetId(
	ctx context.Context,
	subnetExtId string,
	spec any,
) (*networkingprismapi.TaskReference, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	if subnetExtId == "" {
		return nil, fmt.Errorf("subnetExtId is required")
	}

	ipReserveSpec, ok := spec.(*subnetModels.IpReserveSpec)
	if !ok {
		return nil, fmt.Errorf("spec must be of type *subnetModels.IpReserveSpec")
	}

	if ipReserveSpec == nil {
		return nil, fmt.Errorf("ipReserveSpec is required")
	}

	// Call the V4 API and extract TaskReference using CallAPI helper
	taskRef, err := CallAPI[*subnetModels.TaskReferenceApiResponse, networkingprismapi.TaskReference](
		s.client.SubnetIPReservationApi.ReserveIpsBySubnetId(&subnetExtId, ipReserveSpec),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to reserve IPs for subnet %s: %w", subnetExtId, err)
	}

	return &taskRef, nil
}

// UnreserveIpsBySubnetId unreserves IP addresses on a subnet.
// It returns a TaskReference that can be used to track the async operation.
func (s *SubnetsService) UnreserveIpsBySubnetId(
	ctx context.Context,
	subnetExtId string,
	spec any,
) (*networkingprismapi.TaskReference, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	if subnetExtId == "" {
		return nil, fmt.Errorf("subnetExtId is required")
	}

	ipUnreserveSpec, ok := spec.(*subnetModels.IpUnreserveSpec)
	if !ok {
		return nil, fmt.Errorf("spec must be of type *subnetModels.IpUnreserveSpec")
	}

	if ipUnreserveSpec == nil {
		return nil, fmt.Errorf("ipUnreserveSpec is required")
	}

	// Call the V4 API and extract TaskReference using CallAPI helper
	taskRef, err := CallAPI[*subnetModels.TaskReferenceApiResponse, networkingprismapi.TaskReference](
		s.client.SubnetIPReservationApi.UnreserveIpsBySubnetId(&subnetExtId, ipUnreserveSpec),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to unreserve IPs for subnet %s: %w", subnetExtId, err)
	}

	return &taskRef, nil
}

// ListReservedIpsBySubnetId fetches a list of reserved IP addresses on a particular managed subnet.
func (s *SubnetsService) ListReservedIpsBySubnetId(
	ctx context.Context,
	subnetExtId string,
	opts ...converged.ODataOption,
) (any, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	if subnetExtId == "" {
		return nil, fmt.Errorf("subnetExtId is required")
	}

	// Convert OData options to V4 OData parameters
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	// Check if unsupported OData options are provided
	if reqParams != nil {
		if reqParams.Apply != nil || reqParams.Expand != nil {
			return nil, errors.New("apply and expand are not supported for listing reserved IPs")
		}
	}

	// Call the V4 API
	resp, err := s.client.SubnetIPReservationApi.ListReservedIpsBySubnetId(
		&subnetExtId,
		reqParams.Page,
		reqParams.Limit,
		reqParams.Filter,
		reqParams.OrderBy,
		reqParams.Select,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list reserved IPs for subnet %s: %w", subnetExtId, err)
	}

	return resp, nil
}
