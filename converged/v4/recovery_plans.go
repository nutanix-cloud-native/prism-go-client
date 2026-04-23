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

// RecoveryPlansService provides methods to interact with Data Policies Recovery Plans API (v4.2).
type RecoveryPlansService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewRecoveryPlansService creates a new RecoveryPlansService instance.
func NewRecoveryPlansService(client *v4prismGoClient.Client) *RecoveryPlansService {
	return &RecoveryPlansService{
		client:       client,
		entitiesName: "recovery plan",
	}
}

// Get returns a single RecoveryPlan by its external ID.
func (s *RecoveryPlansService) Get(ctx context.Context, uuid string) (*dpModels.RecoveryPlan, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	return GenericGetEntity[*dpModels.GetRecoveryPlanApiResponse, dpModels.RecoveryPlan](
		func() (*dpModels.GetRecoveryPlanApiResponse, error) {
			return s.client.RecoveryPlansApiInstance.GetRecoveryPlanById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of RecoveryPlan entities.
func (s *RecoveryPlansService) List(ctx context.Context, opts ...converged.ODataOption) ([]dpModels.RecoveryPlan, error) {
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
	return GenericListEntities[*dpModels.ListRecoveryPlansApiResponse, dpModels.RecoveryPlan](
		func(reqParams *V4ODataParams) (*dpModels.ListRecoveryPlansApiResponse, error) {
			return s.client.RecoveryPlansApiInstance.ListRecoveryPlans(
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

// NewIterator returns an iterator for listing RecoveryPlan entities.
func (s *RecoveryPlansService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[dpModels.RecoveryPlan] {
	if s.client == nil {
		return nil
	}
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil || (reqParams != nil && (reqParams.Apply != nil || reqParams.Expand != nil)) {
		return func(yield func(dpModels.RecoveryPlan, error) bool) {
			var zero dpModels.RecoveryPlan
			if err != nil {
				yield(zero, fmt.Errorf("failed to convert options to V4ODataParams: %w", err))
				return
			}
			yield(zero, fmt.Errorf("apply and expand options are not supported for listing %s", s.entitiesName))
		}
	}
	return GenericNewIterator[*dpModels.ListRecoveryPlansApiResponse, dpModels.RecoveryPlan](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*dpModels.ListRecoveryPlansApiResponse, error) {
			return s.client.RecoveryPlansApiInstance.ListRecoveryPlans(
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

// Create creates a new recovery plan and waits for the task to complete.
func (s *RecoveryPlansService) Create(ctx context.Context, plan *dpModels.RecoveryPlan) (*dpModels.RecoveryPlan, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	op, err := s.CreateAsync(ctx, plan)
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

// CreateAsync creates a new recovery plan asynchronously.
func (s *RecoveryPlansService) CreateAsync(ctx context.Context, plan *dpModels.RecoveryPlan) (converged.Operation[dpModels.RecoveryPlan], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	resp, err := s.client.RecoveryPlansApiInstance.CreateRecoveryPlan(plan, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	taskRef, err := CallAPI[*dpModels.CreateRecoveryPlanApiResponse, dpPrism.TaskReference](resp, err)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for created %s", s.entitiesName)
	}
	return NewOperation(*taskRef.ExtId, s.client, s.Get), nil
}

// Update updates an existing recovery plan and waits for the task to complete.
func (s *RecoveryPlansService) Update(ctx context.Context, uuid string, plan *dpModels.RecoveryPlan) (*dpModels.RecoveryPlan, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	op, err := s.UpdateAsync(ctx, uuid, plan)
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

// UpdateAsync updates an existing recovery plan asynchronously.
func (s *RecoveryPlansService) UpdateAsync(ctx context.Context, uuid string, plan *dpModels.RecoveryPlan) (converged.Operation[dpModels.RecoveryPlan], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	resp, err := s.client.RecoveryPlansApiInstance.UpdateRecoveryPlanById(&uuid, plan, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	taskRef, err := CallAPI[*dpModels.UpdateRecoveryPlanApiResponse, dpPrism.TaskReference](resp, err)
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for updated %s", s.entitiesName)
	}
	return NewOperation(*taskRef.ExtId, s.client, s.Get), nil
}

// Delete deletes a recovery plan and waits for the task to complete.
func (s *RecoveryPlansService) Delete(ctx context.Context, uuid string) error {
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

// DeleteAsync deletes a recovery plan asynchronously.
func (s *RecoveryPlansService) DeleteAsync(ctx context.Context, uuid string) (converged.Operation[converged.NoEntity], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	_, args, err := GetEntityAndEtag(
		s.client.RecoveryPlansApiInstance.GetRecoveryPlanById(&uuid, nil),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s for deletion: %w", s.entitiesName, err)
	}
	resp, err := s.client.RecoveryPlansApiInstance.DeleteRecoveryPlanById(&uuid, args)
	if err != nil {
		return nil, fmt.Errorf("failed to delete %s: %w", s.entitiesName, err)
	}
	taskRef, err := CallAPI[*dpModels.DeleteRecoveryPlanApiResponse, dpPrism.TaskReference](resp, err)
	if err != nil {
		return nil, fmt.Errorf("failed to delete %s: %w", s.entitiesName, err)
	}
	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for deleted %s", s.entitiesName)
	}
	return NewOperation(*taskRef.ExtId, s.client, converged.NoEntityGetter), nil
}
