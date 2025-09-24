package v4

import (
	"context"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	subnetModels "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// SubnetsService provides default "not implemented" implementation for all Subnets interface methods.
type SubnetsService struct {
	client *v4prismGoClient.Client
}

// NewSubnetsService creates a new SubnetsService instance.
func NewSubnetsService(client *v4prismGoClient.Client) *SubnetsService {
	return &SubnetsService{client: client}
}

// Get returns the subnet for the given UUID.
func (s *SubnetsService) Get(ctx context.Context, uuid string) (*subnetModels.Subnet, error) {
	return nil, fmt.Errorf("not implemented")
}

// List returns a list of subnets.
func (s *SubnetsService) List(ctx context.Context, opts ...converged.ODataOption) ([]subnetModels.Subnet, error) {
	return nil, fmt.Errorf("not implemented")
}

// NewIterator returns an iterator for listing subnets.
func (s *SubnetsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[subnetModels.Subnet] {
	return nil
}
