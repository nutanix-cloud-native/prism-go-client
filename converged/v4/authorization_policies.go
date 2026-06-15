package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	iamAuthzModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
)

// AuthorizationPoliciesService provides implementation for IAM Authorization Policies API operations.
type AuthorizationPoliciesService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewAuthorizationPoliciesService creates a new AuthorizationPoliciesService instance.
func NewAuthorizationPoliciesService(client *v4prismGoClient.Client) *AuthorizationPoliciesService {
	return &AuthorizationPoliciesService{
		client:       client,
		entitiesName: "authorization policy",
	}
}

// Get returns the authorization policy for the given UUID.
func (s *AuthorizationPoliciesService) Get(ctx context.Context, uuid string) (*iamAuthzModels.AuthorizationPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthzModels.GetAuthorizationPolicyApiResponse, iamAuthzModels.AuthorizationPolicy](
		func() (*iamAuthzModels.GetAuthorizationPolicyApiResponse, error) {
			return s.client.AuthorizationPoliciesApiInstance.GetAuthorizationPolicyById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of authorization policies.
func (s *AuthorizationPoliciesService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthzModels.AuthorizationPolicyProjection, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && myParams.Apply != nil {
		return nil, fmt.Errorf("apply option is not supported for listing Authorization Policies")
	}

	return GenericListEntities[*iamAuthzModels.ListAuthorizationPoliciesApiResponse, iamAuthzModels.AuthorizationPolicyProjection](
		func(reqParams *V4ODataParams) (*iamAuthzModels.ListAuthorizationPoliciesApiResponse, error) {
			return s.client.AuthorizationPoliciesApiInstance.ListAuthorizationPolicies(
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

// NewIterator returns an iterator for listing authorization policies.
func (s *AuthorizationPoliciesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthzModels.AuthorizationPolicyProjection] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthzModels.ListAuthorizationPoliciesApiResponse, iamAuthzModels.AuthorizationPolicyProjection](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthzModels.ListAuthorizationPoliciesApiResponse, error) {
			return s.client.AuthorizationPoliciesApiInstance.ListAuthorizationPolicies(
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

// Create creates a new authorization policy.
func (s *AuthorizationPoliciesService) Create(ctx context.Context, entity *iamAuthzModels.AuthorizationPolicy) (*iamAuthzModels.AuthorizationPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	newPolicy, err := CallAPI[*iamAuthzModels.CreateAuthorizationPolicyApiResponse, iamAuthzModels.AuthorizationPolicy](
		s.client.AuthorizationPoliciesApiInstance.CreateAuthorizationPolicy(entity),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create authorization policy: %w", err)
	}

	return &newPolicy, nil
}

// Delete deletes an existing authorization policy.
func (s *AuthorizationPoliciesService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	_, args, err := GetEntityAndEtag(
		s.client.AuthorizationPoliciesApiInstance.GetAuthorizationPolicyById(&uuid),
	)
	if err != nil {
		return fmt.Errorf("failed to get authorization policy with UUID %s: %w", uuid, err)
	}

	_, err = s.client.AuthorizationPoliciesApiInstance.DeleteAuthorizationPolicyById(&uuid, args)
	if err != nil {
		return fmt.Errorf("failed to delete authorization policy with UUID %s: %w", uuid, err)
	}

	return nil
}
