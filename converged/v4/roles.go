package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	iamAuthzModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
)

// RolesService provides implementation for IAM Roles API operations.
type RolesService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewRolesService creates a new RolesService instance.
func NewRolesService(client *v4prismGoClient.Client) *RolesService {
	return &RolesService{
		client:       client,
		entitiesName: "role",
	}
}

// Get returns the role for the given UUID.
func (s *RolesService) Get(ctx context.Context, uuid string) (*iamAuthzModels.Role, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthzModels.GetRoleApiResponse, iamAuthzModels.Role](
		func() (*iamAuthzModels.GetRoleApiResponse, error) {
			return s.client.RolesApiInstance.GetRoleById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of roles.
func (s *RolesService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthzModels.Role, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing Roles")
	}

	return GenericListEntities[*iamAuthzModels.ListRolesApiResponse, iamAuthzModels.Role](
		func(reqParams *V4ODataParams) (*iamAuthzModels.ListRolesApiResponse, error) {
			return s.client.RolesApiInstance.ListRoles(
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

// NewIterator returns an iterator for listing roles.
func (s *RolesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthzModels.Role] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthzModels.ListRolesApiResponse, iamAuthzModels.Role](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthzModels.ListRolesApiResponse, error) {
			return s.client.RolesApiInstance.ListRoles(
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

// Create creates a new role.
func (s *RolesService) Create(ctx context.Context, entity *iamAuthzModels.Role) (*iamAuthzModels.Role, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	newRole, err := CallAPI[*iamAuthzModels.CreateRoleApiResponse, iamAuthzModels.Role](
		s.client.RolesApiInstance.CreateRole(entity),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %w", err)
	}

	return &newRole, nil
}

// Delete deletes an existing role.
func (s *RolesService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}
	_, args, err := GetEntityAndEtag(
		s.client.RolesApiInstance.GetRoleById(&uuid),
	)
	if err != nil {
		return fmt.Errorf("failed to get role with UUID %s: %w", uuid, err)
	}

	_, err = s.client.RolesApiInstance.DeleteRoleById(&uuid, args)
	return err
}
