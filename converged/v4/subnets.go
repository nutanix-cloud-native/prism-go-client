package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	subnetModels "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// SubnetsService provides implementation for all Subnets interface methods.
type SubnetsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewSubnetsService creates a new SubnetsService instance.
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
	if s.client != nil {
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
