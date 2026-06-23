package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	iamAuthzModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
)

// OperationsService provides implementation for IAM Operations API operations.
type OperationsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewOperationsService creates a new OperationsService instance.
func NewOperationsService(client *v4prismGoClient.Client) *OperationsService {
	return &OperationsService{
		client:       client,
		entitiesName: "operation",
	}
}

// Get returns the operation for the given UUID.
func (s *OperationsService) Get(ctx context.Context, uuid string) (*iamAuthzModels.Operation, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthzModels.GetOperationApiResponse, iamAuthzModels.Operation](
		func() (*iamAuthzModels.GetOperationApiResponse, error) {
			return s.client.OperationsApiInstance.GetOperationById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of operations.
func (s *OperationsService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthzModels.Operation, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing Operations")
	}

	return GenericListEntities[*iamAuthzModels.ListOperationsApiResponse, iamAuthzModels.Operation](
		func(reqParams *V4ODataParams) (*iamAuthzModels.ListOperationsApiResponse, error) {
			return s.client.OperationsApiInstance.ListOperations(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
				nil,
			)
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing operations.
func (s *OperationsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthzModels.Operation] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthzModels.ListOperationsApiResponse, iamAuthzModels.Operation](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthzModels.ListOperationsApiResponse, error) {
			return s.client.OperationsApiInstance.ListOperations(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
				nil,
			)
		},
		opts,
		s.entitiesName,
	)
}
