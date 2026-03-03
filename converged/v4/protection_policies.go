package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	dpModels "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
	dpPrism "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/prism/v4/config"
)

// ProtectionPoliciesService provides methods to interact with Data Policies Protection Policies API (v4.2).
type ProtectionPoliciesService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewProtectionPoliciesService creates a new ProtectionPoliciesService instance.
func NewProtectionPoliciesService(client *v4prismGoClient.Client) *ProtectionPoliciesService {
	return &ProtectionPoliciesService{
		client:       client,
		entitiesName: "protection policy",
	}
}

// Get returns a single ProtectionPolicy by its external ID.
func (s *ProtectionPoliciesService) Get(ctx context.Context, uuid string) (*dpModels.ProtectionPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	return GenericGetEntity[*dpModels.GetProtectionPolicyApiResponse, dpModels.ProtectionPolicy](
		func() (*dpModels.GetProtectionPolicyApiResponse, error) {
			return s.client.ProtectionPoliciesApiInstance.GetProtectionPolicyById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of ProtectionPolicy entities.
func (s *ProtectionPoliciesService) List(ctx context.Context, opts ...converged.ODataOption) ([]dpModels.ProtectionPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}
	if reqParams != nil && (reqParams.Apply != nil || reqParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing %s", s.entitiesName)
	}
	return GenericListEntities[*dpModels.ListProtectionPoliciesApiResponse, dpModels.ProtectionPolicy](
		func(reqParams *V4ODataParams) (*dpModels.ListProtectionPoliciesApiResponse, error) {
			return s.client.ProtectionPoliciesApiInstance.ListProtectionPolicies(
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

// NewIterator returns an iterator for listing ProtectionPolicy entities.
func (s *ProtectionPoliciesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[dpModels.ProtectionPolicy] {
	if s.client == nil {
		return nil
	}
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil || (reqParams != nil && (reqParams.Apply != nil || reqParams.Expand != nil)) {
		return func(yield func(dpModels.ProtectionPolicy, error) bool) {
			var zero dpModels.ProtectionPolicy
			if err != nil {
				yield(zero, fmt.Errorf("failed to convert options to V4ODataParams: %w", err))
				return
			}
			yield(zero, fmt.Errorf("apply and expand options are not supported for listing %s", s.entitiesName))
		}
	}
	return GenericNewIterator[*dpModels.ListProtectionPoliciesApiResponse, dpModels.ProtectionPolicy](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*dpModels.ListProtectionPoliciesApiResponse, error) {
			return s.client.ProtectionPoliciesApiInstance.ListProtectionPolicies(
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

// Create creates a new protection policy and waits for the task to complete.
func (s *ProtectionPoliciesService) Create(ctx context.Context, policy *dpModels.ProtectionPolicy) (*dpModels.ProtectionPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	op, err := s.CreateAsync(ctx, policy)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	result, err := op.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("failed to create %s: operation completed but no entity returned", s.entitiesName)
	}
	return result[0], nil
}

// CreateAsync creates a new protection policy asynchronously.
func (s *ProtectionPoliciesService) CreateAsync(ctx context.Context, policy *dpModels.ProtectionPolicy) (converged.Operation[dpModels.ProtectionPolicy], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	resp, err := s.client.ProtectionPoliciesApiInstance.CreateProtectionPolicy(policy, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	taskRef, err := CallAPI[*dpModels.CreateProtectionPolicyApiResponse, *dpPrism.TaskReference](resp, err)
	if err != nil || taskRef == nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for created %s", s.entitiesName)
	}
	return NewOperation(*taskRef.ExtId, s.client, s.Get), nil
}

// Update updates an existing protection policy and waits for the task to complete.
func (s *ProtectionPoliciesService) Update(ctx context.Context, uuid string, policy *dpModels.ProtectionPolicy) (*dpModels.ProtectionPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	op, err := s.UpdateAsync(ctx, uuid, policy)
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	result, err := op.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("failed to update %s: operation completed but no entity returned", s.entitiesName)
	}
	return result[0], nil
}

// UpdateAsync updates an existing protection policy asynchronously.
func (s *ProtectionPoliciesService) UpdateAsync(ctx context.Context, uuid string, policy *dpModels.ProtectionPolicy) (converged.Operation[dpModels.ProtectionPolicy], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	resp, err := s.client.ProtectionPoliciesApiInstance.UpdateProtectionPolicyById(&uuid, policy, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	taskRef, err := CallAPI[*dpModels.UpdateProtectionPolicyApiResponse, *dpPrism.TaskReference](resp, err)
	if err != nil || taskRef == nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for updated %s", s.entitiesName)
	}
	return NewOperation(*taskRef.ExtId, s.client, s.Get), nil
}

// Delete deletes a protection policy and waits for the task to complete.
func (s *ProtectionPoliciesService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}
	op, err := s.DeleteAsync(ctx, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %w", s.entitiesName, err)
	}
	_, err = op.Wait(ctx)
	return err
}

// DeleteAsync deletes a protection policy asynchronously.
func (s *ProtectionPoliciesService) DeleteAsync(ctx context.Context, uuid string) (converged.Operation[converged.NoEntity], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	resp, err := s.client.ProtectionPoliciesApiInstance.DeleteProtectionPolicyById(&uuid, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete %s: %w", s.entitiesName, err)
	}
	taskRef, err := CallAPI[*dpModels.DeleteProtectionPolicyApiResponse, *dpPrism.TaskReference](resp, err)
	if err != nil || taskRef == nil {
		return nil, fmt.Errorf("failed to delete %s: %w", s.entitiesName, err)
	}
	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for deleted %s", s.entitiesName)
	}
	return NewOperation(*taskRef.ExtId, s.client, converged.NoEntityGetter), nil
}
