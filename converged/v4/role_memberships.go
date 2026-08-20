package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	iamAuthzModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
	rolemembershipRequestModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/request/rolemembership"
)

// RoleMembershipsService provides implementation for IAM Role Membership API operations.
type RoleMembershipsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewRoleMembershipsService creates a new RoleMembershipsService instance.
func NewRoleMembershipsService(client *v4prismGoClient.Client) *RoleMembershipsService {
	return &RoleMembershipsService{
		client:       client,
		entitiesName: "role membership",
	}
}

// Get returns the role membership for the given UUID.
func (s *RoleMembershipsService) Get(ctx context.Context, uuid string) (*iamAuthzModels.RoleMembership, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthzModels.GetRoleMembershipApiResponse, iamAuthzModels.RoleMembership](
		func() (*iamAuthzModels.GetRoleMembershipApiResponse, error) {
			return s.client.RoleMembershipApiInstance.GetRoleMembershipById(
				ctx,
				&rolemembershipRequestModels.GetRoleMembershipByIdRequest{ExtId: &uuid},
			)
		},
		s.entitiesName,
	)
}

// List returns a list of role memberships.
func (s *RoleMembershipsService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthzModels.RoleMembershipProjection, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && myParams.Apply != nil {
		return nil, fmt.Errorf("apply option is not supported for listing Role Memberships")
	}

	return GenericListEntities[*iamAuthzModels.ListRoleMembershipsApiResponse, iamAuthzModels.RoleMembershipProjection](
		func(reqParams *V4ODataParams) (*iamAuthzModels.ListRoleMembershipsApiResponse, error) {
			return s.client.RoleMembershipApiInstance.ListRoleMemberships(ctx, &rolemembershipRequestModels.ListRoleMembershipsRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Expand_:  reqParams.Expand,
				Select_:  reqParams.Select,
			})
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing role memberships.
func (s *RoleMembershipsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthzModels.RoleMembershipProjection] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthzModels.ListRoleMembershipsApiResponse, iamAuthzModels.RoleMembershipProjection](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthzModels.ListRoleMembershipsApiResponse, error) {
			return s.client.RoleMembershipApiInstance.ListRoleMemberships(ctx, &rolemembershipRequestModels.ListRoleMembershipsRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Expand_:  reqParams.Expand,
				Select_:  reqParams.Select,
			})
		},
		opts,
		s.entitiesName,
	)
}

// Create creates a new role membership.
func (s *RoleMembershipsService) Create(ctx context.Context, entity *iamAuthzModels.RoleMembership) (*iamAuthzModels.RoleMembership, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	newRoleMembership, err := CallAPI[*iamAuthzModels.CreateRoleMembershipApiResponse, iamAuthzModels.RoleMembership](
		s.client.RoleMembershipApiInstance.CreateRoleMembership(
			ctx,
			&rolemembershipRequestModels.CreateRoleMembershipRequest{Body: entity},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create role membership: %w", err)
	}

	return &newRoleMembership, nil
}

// Delete deletes an existing role membership.
func (s *RoleMembershipsService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}
	_, args, err := GetEntityAndEtag(
		s.client.RoleMembershipApiInstance.GetRoleMembershipById(
			ctx,
			&rolemembershipRequestModels.GetRoleMembershipByIdRequest{ExtId: &uuid},
		),
	)
	if err != nil {
		return fmt.Errorf("failed to get role membership with UUID %s: %w", uuid, err)
	}

	_, err = s.client.RoleMembershipApiInstance.DeleteRoleMembershipById(
		ctx,
		&rolemembershipRequestModels.DeleteRoleMembershipByIdRequest{ExtId: &uuid},
		args,
	)
	return err
}
