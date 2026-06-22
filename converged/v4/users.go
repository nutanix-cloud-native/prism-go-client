package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	iamModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
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

// Create creates a new user.
func (s *UsersService) Create(ctx context.Context, entity *iamModels.User) (*iamModels.User, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	newUser, err := CallAPI[*iamModels.CreateUserApiResponse, iamModels.User](
		s.client.UsersApiInstance.CreateUser(entity),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &newUser, nil
}

// Update deactivates an existing user.
func (s *UsersService) UpdateUserState(uuid string, status *iamModels.UserStateUpdate) (*iamModels.UserStateUpdateResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	_, args, err := GetEntityAndEtag(
		s.client.UsersApiInstance.GetUserById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user with UUID %s: %w", uuid, err)
	}
	
	userStateUpdateResponse, err := CallAPI[*iamModels.ActivateUserApiResponse, iamModels.UserStateUpdateResponse](
		s.client.UsersApiInstance.UpdateUserState(&uuid, status, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update user state for user %s: %w", uuid, err)
	}

	return &userStateUpdateResponse, nil
}

// ListUserKeys returns a list of user keys.
func (s *UsersService) ListUserKeys(ctx context.Context, userExtId string, opts ...converged.ODataOption) ([]iamModels.Key, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing User Keys")
	}

	return GenericListEntities[*iamModels.ListUserKeysApiResponse, iamModels.Key](
		func(reqParams *V4ODataParams) (*iamModels.ListUserKeysApiResponse, error) {
			return s.client.UsersApiInstance.ListUserKeys(&userExtId, reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy, reqParams.Select)
		},
		opts,
		"user_key",
	)
}

// CreateUserKey creates a new user key.
func (s *UsersService) CreateUserKey(ctx context.Context, userExtId string, key *iamModels.Key) (*iamModels.Key, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	apiKeyDetails, err := CallAPI[*iamModels.CreateKeyApiResponse, iamModels.Key](
		s.client.UsersApiInstance.CreateUserKey(&userExtId, key, nil),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user key for user %s: %w", userExtId, err)
	}

	return &apiKeyDetails, nil
}

// GetUserKeyById returns a user key by key ID.
func (s *UsersService) GetUserKeyById(ctx context.Context, userExtId string, keyID string) (*iamModels.Key, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamModels.GetUserKeyApiResponse, iamModels.Key](
		func() (*iamModels.GetUserKeyApiResponse, error) {
			return s.client.UsersApiInstance.GetUserKeyById(&userExtId, &keyID, nil)
		},
		"user_key",
	)
}

// DeleteUserKeyById deletes a user key by key ID.
func (s *UsersService) DeleteUserKeyById(ctx context.Context, userExtId string, keyID string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}
	
	_, args, err := GetEntityAndEtag(
		s.client.UsersApiInstance.GetUserKeyById(&userExtId, &keyID),
	)
	if err != nil {
		return fmt.Errorf("failed to get user key %s for user %s: %w", keyID, userExtId, err)
	}

	_, err = s.client.UsersApiInstance.DeleteUserKeyById(&userExtId, &keyID, args)
	if err != nil {
		return fmt.Errorf("failed to delete user key %s for user %s: %w", keyID, userExtId, err)
	}

	return nil
}

// RevokeUserKey revokes a user key by key ID.
func (s *UsersService) RevokeUserKey(ctx context.Context, userExtId string, keyID string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	_, err := s.client.UsersApiInstance.RevokeUserKey(&userExtId, &keyID, nil)
	if err != nil {
		return fmt.Errorf("failed to revoke user key %s for user %s: %w", keyID, userExtId, err)
	}

	return nil
}
