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
func (s *UsersService) Create(ctx context.Context, user *iamModels.User) (*iamModels.User, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*iamModels.CreateUserApiResponse, iamModels.User](
		s.client.UsersApiInstance.CreateUser(user, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	return &result, nil
}

// Update updates an existing user.
func (s *UsersService) Update(ctx context.Context, uuid string, user *iamModels.User) (*iamModels.User, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*iamModels.UpdateUserApiResponse, iamModels.User](
		s.client.UsersApiInstance.UpdateUserById(&uuid, user, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	return &result, nil
}

// UpdateUserState updates the state of a user (activate/deactivate).
func (s *UsersService) UpdateUserState(ctx context.Context, uuid string, stateUpdate any) (*iamModels.User, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	userStateUpdate, ok := stateUpdate.(*iamModels.UserStateUpdate)
	if !ok {
		return nil, fmt.Errorf("invalid state update type: expected *iamModels.UserStateUpdate")
	}

	result, err := CallAPI[*iamModels.ActivateUserApiResponse, iamModels.User](
		s.client.UsersApiInstance.UpdateUserState(&uuid, userStateUpdate, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to update %s state: %w", s.entitiesName, err)
	}
	return &result, nil
}

// ChangeUserPassword changes the password for a user.
func (s *UsersService) ChangeUserPassword(ctx context.Context, passwordChange any) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	passwordChangeRequest, ok := passwordChange.(*iamModels.PasswordChangeRequest)
	if !ok {
		return fmt.Errorf("invalid password change type: expected *iamModels.PasswordChangeRequest")
	}

	_, err := s.client.UsersApiInstance.ChangeUserPassword(passwordChangeRequest, nil)
	if err != nil {
		return fmt.Errorf("failed to change %s password: %w", s.entitiesName, err)
	}
	return nil
}

// ResetUserPassword resets the password for a user.
func (s *UsersService) ResetUserPassword(ctx context.Context, uuid string, passwordReset any) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	passwordResetRequest, ok := passwordReset.(*iamModels.PasswordResetRequest)
	if !ok {
		return fmt.Errorf("invalid password reset type: expected *iamModels.PasswordResetRequest")
	}

	_, err := s.client.UsersApiInstance.ResetUserPassword(&uuid, passwordResetRequest, nil)
	if err != nil {
		return fmt.Errorf("failed to reset %s password: %w", s.entitiesName, err)
	}
	return nil
}
