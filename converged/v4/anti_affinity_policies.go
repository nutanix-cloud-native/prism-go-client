package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	vmmPrismModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	policyModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
)

// AntiAffinityPoliciesService provides default "not implemented" implementation for all AntiAffinityPolicies interface methods.
type AntiAffinityPoliciesService struct {
	client   *v4prismGoClient.Client
	entities string
}

// NewAntiAffinityPoliciesService creates a new AntiAffinityPoliciesService instance.
func NewAntiAffinityPoliciesService(client *v4prismGoClient.Client) *AntiAffinityPoliciesService {
	return &AntiAffinityPoliciesService{client: client, entities: "anti-affinity policy"}
}

// Get returns the anti-affinity policy for the given UUID.
func (s *AntiAffinityPoliciesService) Get(ctx context.Context, uuid string) (*policyModels.VmAntiAffinityPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	return GenericGetEntity[*policyModels.GetVmAntiAffinityPolicyApiResponse, policyModels.VmAntiAffinityPolicy](
		func() (*policyModels.GetVmAntiAffinityPolicyApiResponse, error) {
			return s.client.VmAntiAffinityPoliciesApiInstance.GetVmAntiAffinityPolicyById(&uuid)
		},
		s.entities,
	)
}

// List returns a list of anti-affinity policies.
func (s *AntiAffinityPoliciesService) List(ctx context.Context, opts ...converged.ODataOption) ([]policyModels.VmAntiAffinityPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	// Check if unsupported OData options are provided
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}
	if reqParams != nil {
		if reqParams.Apply != nil && reqParams.Expand != nil && reqParams.Select != nil {
			return nil, errors.New("apply, expand and select are not supported")
		}
	}

	return GenericListEntities[*policyModels.ListVmAntiAffinityPoliciesApiResponse, policyModels.VmAntiAffinityPolicy](
		func(reqParams *V4ODataParams) (*policyModels.ListVmAntiAffinityPoliciesApiResponse, error) {
			return s.client.VmAntiAffinityPoliciesApiInstance.ListVmAntiAffinityPolicies(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
			)
		},
		opts,
		s.entities,
	)
}

// NewIterator returns an iterator for listing anti-affinity policies.
func (s *AntiAffinityPoliciesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[policyModels.VmAntiAffinityPolicy] {
	if s.client == nil {
		return nil
	}
	return GenericNewIterator[*policyModels.ListVmAntiAffinityPoliciesApiResponse, policyModels.VmAntiAffinityPolicy](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*policyModels.ListVmAntiAffinityPoliciesApiResponse, error) {
			return s.client.VmAntiAffinityPoliciesApiInstance.ListVmAntiAffinityPolicies(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
			)
		},
		opts,
		s.entities,
	)
}

// Create creates a new anti-affinity policy.
func (s *AntiAffinityPoliciesService) Create(ctx context.Context, entity *policyModels.VmAntiAffinityPolicy) (*policyModels.VmAntiAffinityPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	taskRef, err := CallAPI[*policyModels.CreateVmAntiAffinityPolicyApiResponse, vmmPrismModels.TaskReference](
		s.client.VmAntiAffinityPoliciesApiInstance.CreateVmAntiAffinityPolicy(entity),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create anti-affinity policy: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for created anti-affinity policy")
	}

	waiter := NewOperation(*taskRef.ExtId, s.client, func(ctx context.Context, uuid string) (*policyModels.VmAntiAffinityPolicy, error) {
		return s.Get(ctx, uuid)
	})
	createdPolicy, err := waiter.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create anti-affinity policy: %w", err)
	}
	return createdPolicy[0], nil
}

// Update updates an existing anti-affinity policy.
func (s *AntiAffinityPoliciesService) Update(ctx context.Context, uuid string, entity *policyModels.VmAntiAffinityPolicy) (*policyModels.VmAntiAffinityPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	taskRef, err := CallAPI[*policyModels.UpdateVmAntiAffinityPolicyApiResponse, vmmPrismModels.TaskReference](
		s.client.VmAntiAffinityPoliciesApiInstance.UpdateVmAntiAffinityPolicyById(&uuid, entity),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update anti-affinity policy with UUID %s: %w", uuid, err)
	}

	waiter := NewOperation(*taskRef.ExtId, s.client, s.Get)
	updatedPolicy, err := waiter.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update anti-affinity policy: %w", err)
	}
	return updatedPolicy[0], nil
}

// Delete deletes an existing anti-affinity policy.
func (s *AntiAffinityPoliciesService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}
	policy, args, err := GetEntityAndEtag(
		s.Get(ctx, uuid),
	)
	if err != nil {
		return fmt.Errorf("failed to get anti-affinity policy with UUID %s: %w", uuid, err)
	}

	if policy == nil {
		return fmt.Errorf("no anti-affinity policy found with UUID %s", uuid)
	}

	taskRef, err := CallAPI[*policyModels.DeleteVmAntiAffinityPolicyApiResponse, vmmPrismModels.TaskReference](
		s.client.VmAntiAffinityPoliciesApiInstance.DeleteVmAntiAffinityPolicyById(&uuid, args),
	)
	if err != nil {
		return fmt.Errorf("failed to delete anti-affinity policy with UUID %s: %w", uuid, err)
	}

	if taskRef.ExtId == nil {
		return fmt.Errorf("task reference ExtId is nil for deleted anti-affinity policy with UUID %s", uuid)
	}

	waiter := NewOperation(*taskRef.ExtId, s.client, func(ctx context.Context, uuid string) (*converged.NoEntity, error) {
		return nil, nil
	})
	_, err = waiter.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete anti-affinity policy: %w", err)
	}
	return nil
}

// CreateAsync creates a new anti-affinity policy asynchronously.
// Note: Anti-affinity policy creation is actually synchronous, so this returns an already-completed operation.
func (s *AntiAffinityPoliciesService) CreateAsync(ctx context.Context, entity *policyModels.VmAntiAffinityPolicy) (converged.Operation[policyModels.VmAntiAffinityPolicy], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	taskRef, err := CallAPI[*policyModels.CreateVmAntiAffinityPolicyApiResponse, vmmPrismModels.TaskReference](
		s.client.VmAntiAffinityPoliciesApiInstance.CreateVmAntiAffinityPolicy(entity),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create anti-affinity policy: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for created anti-affinity policy")
	}

	waiter := NewOperation(*taskRef.ExtId, s.client, s.Get)
	return waiter, nil
}

// UpdateAsync updates an existing anti-affinity policy asynchronously.
// Note: Anti-affinity policy updates are actually synchronous, so this returns an already-completed operation.
func (s *AntiAffinityPoliciesService) UpdateAsync(ctx context.Context, uuid string, entity *policyModels.VmAntiAffinityPolicy) (converged.Operation[policyModels.VmAntiAffinityPolicy], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	taskRef, err := CallAPI[*policyModels.UpdateVmAntiAffinityPolicyApiResponse, vmmPrismModels.TaskReference](
		s.client.VmAntiAffinityPoliciesApiInstance.UpdateVmAntiAffinityPolicyById(&uuid, entity),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update anti-affinity policy with UUID %s: %w", uuid, err)
	}

	waiter := NewOperation(*taskRef.ExtId, s.client, s.Get)
	return waiter, nil
}

// DeleteAsync deletes an existing anti-affinity policy asynchronously.
// Note: Anti-affinity policy deletion is actually synchronous, so this returns an already-completed operation.
func (s *AntiAffinityPoliciesService) DeleteAsync(ctx context.Context, uuid string) (converged.Operation[converged.NoEntity], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	policy, args, err := GetEntityAndEtag(
		s.Get(ctx, uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get anti-affinity policy with UUID %s: %w", uuid, err)
	}

	if policy == nil {
		return nil, fmt.Errorf("no anti-affinity policy found with UUID %s", uuid)
	}

	taskRef, err := CallAPI[*policyModels.DeleteVmAntiAffinityPolicyApiResponse, vmmPrismModels.TaskReference](
		s.client.VmAntiAffinityPoliciesApiInstance.DeleteVmAntiAffinityPolicyById(&uuid, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete anti-affinity policy with UUID %s: %w", uuid, err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for deleted anti-affinity policy with UUID %s", uuid)
	}

	waiter := NewOperation(*taskRef.ExtId, s.client, func(ctx context.Context, uuid string) (*converged.NoEntity, error) {
		return nil, nil
	})
	return waiter, nil
}
