package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	iamModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	"k8s.io/utils/ptr"
)

// UsersService provides implementation for IAM Users API operations.
type UsersService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewUsersService creates a new UsersService instance.
func NewUsersService(client *v4prismGoClient.Client) *UsersService {
	return &UsersService{
		client:       client,
		entitiesName: "user",
	}
}

// Get returns the user for the given UUID.
func (s *UsersService) Get(ctx context.Context, uuid string) (*iamModels.User, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamModels.GetUserApiResponse, iamModels.User](
		func() (*iamModels.GetUserApiResponse, error) {
			return s.client.UsersApiInstance.GetUserById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of users.
func (s *UsersService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamModels.User, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing Users")
	}

	return GenericListEntities[*iamModels.ListUsersApiResponse, iamModels.User](
		func(reqParams *V4ODataParams) (*iamModels.ListUsersApiResponse, error) {
			return s.client.UsersApiInstance.ListUsers(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing users.
func (s *UsersService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamModels.User] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamModels.ListUsersApiResponse, iamModels.User](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamModels.ListUsersApiResponse, error) {
			return s.client.UsersApiInstance.ListUsers(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// ValidateCredentials validates the provided credentials by making an authenticated API call.
// Returns an error if credentials are invalid or authentication fails.
func (s *UsersService) ValidateCredentials(ctx context.Context) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	// Use ListUsers API with limit 1 to validate credentials
	_, err := s.client.UsersApiInstance.ListUsers(nil, ptr.To(1), nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to validate credentials: %w", err)
	}

	return nil
}
