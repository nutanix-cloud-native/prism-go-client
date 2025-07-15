package v4

import (
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	v4VmmConfig "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	v4policies "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
)

// Implements the FacadeClientV4 interface
func (f *FacadeV4Client) GetAntiAffinityPolicy(uuid string) (*v4policies.VmAntiAffinityPolicy, error) {
	policy, err := CallAPI[*v4policies.GetVmAntiAffinityPolicyApiResponse, v4policies.VmAntiAffinityPolicy](
		f.client.VmAntiAffinityPoliciesApiInstance.GetVmAntiAffinityPolicyById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get anti-affinity policy with UUID %s: %w", uuid, err)
	}
	return &policy, nil
}

func (f *FacadeV4Client) ListAntiAffinityPolicies(opts ...facade.ODataOption) ([]v4policies.VmAntiAffinityPolicy, error) {
	var policies []v4policies.VmAntiAffinityPolicy
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4 OData params: %w", err)
	}

	if reqParams.Page == nil && reqParams.Limit == nil {
		// Get all policies
		page := 0
		policiesChunk, totalCount, err := CallListAPI[*v4policies.ListVmAntiAffinityPoliciesApiResponse, v4policies.VmAntiAffinityPolicy](
			f.client.VmAntiAffinityPoliciesApiInstance.ListVmAntiAffinityPolicies(&page, nil, reqParams.Filter, reqParams.OrderBy),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to list anti-affinity policies: %w", err)
		}
		policies = append(policies, policiesChunk...)

		for len(policies) < totalCount {
			page++
			policiesChunk, totalCount, err = CallListAPI[*v4policies.ListVmAntiAffinityPoliciesApiResponse, v4policies.VmAntiAffinityPolicy](
				f.client.VmAntiAffinityPoliciesApiInstance.ListVmAntiAffinityPolicies(&page, nil, reqParams.Filter, reqParams.OrderBy),
			)
			if err != nil {
				return nil, fmt.Errorf("failed to list anti-affinity policies: %w", err)
			}
			policies = append(policies, policiesChunk...)
		}
	} else {
		// Get exactly requested page and limit with filter and order by
		policies, err = CallAPI[*v4policies.ListVmAntiAffinityPoliciesApiResponse, []v4policies.VmAntiAffinityPolicy](
			f.client.VmAntiAffinityPoliciesApiInstance.ListVmAntiAffinityPolicies(reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to list anti-affinity policies: %w", err)
		}
	}

	return policies, nil
}

func (f *FacadeV4Client) GetListIteratorAntiAffinityPolicies(opts ...facade.ODataOption) facade.ODataListIterator[*v4policies.VmAntiAffinityPolicy] {
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return &FacadeV4ODataIterator[*v4policies.ListVmAntiAffinityPoliciesApiResponse, *v4policies.VmAntiAffinityPolicy]{
			iteratorError: fmt.Errorf("failed to convert options to V4 OData params: %w", err),
		}
	}
	page := 0
	reqParams.Limit = nil  // Let API use the default limit
	reqParams.Page = &page // Start from the page 0
	itemsBuffer, totalCount, err := CallListAPI[*v4policies.ListVmAntiAffinityPoliciesApiResponse, *v4policies.VmAntiAffinityPolicy](
		f.client.VmAntiAffinityPoliciesApiInstance.ListVmAntiAffinityPolicies(&page, nil, reqParams.Filter, reqParams.OrderBy),
	)
	if err != nil {
		return &FacadeV4ODataIterator[*v4policies.ListVmAntiAffinityPoliciesApiResponse, *v4policies.VmAntiAffinityPolicy]{
			iteratorError: fmt.Errorf("failed to list anti-affinity policies: %w", err),
		}
	}

	return NewFacadeV4ODataIterator[*v4policies.ListVmAntiAffinityPoliciesApiResponse, *v4policies.VmAntiAffinityPolicy](
		f.client,
		totalCount,
		itemsBuffer,
		func(params *V4ODataParams) (*v4policies.ListVmAntiAffinityPoliciesApiResponse, error) {
			return f.client.VmAntiAffinityPoliciesApiInstance.ListVmAntiAffinityPolicies(
				params.Page,
				params.Limit,
				params.Filter,
				params.OrderBy,
			)
		},
		opts...,
	)
}

// CreateAntiAffinityPolicy creates a new anti-affinity policy.
func (f *FacadeV4Client) CreateAntiAffinityPolicy(policy v4policies.VmAntiAffinityPolicy) (facade.TaskWaiter[v4policies.VmAntiAffinityPolicy], error) {
	taskRef, err := CallAPI[*v4policies.CreateVmAntiAffinityPolicyApiResponse, v4VmmConfig.TaskReference](
		f.client.VmAntiAffinityPoliciesApiInstance.CreateVmAntiAffinityPolicy(&policy),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create anti-affinity policy: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for created anti-affinity policy")
	}

	waiter := NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, f.GetAntiAffinityPolicy)
	return waiter, nil
}

// UpdateAntiAffinityPolicy updates an existing anti-affinity policy.
func (f *FacadeV4Client) UpdateAntiAffinityPolicy(uuid string, policy v4policies.VmAntiAffinityPolicy) (facade.TaskWaiter[v4policies.VmAntiAffinityPolicy], error) {
	taskRef, err := CallAPI[*v4policies.UpdateVmAntiAffinityPolicyApiResponse, v4VmmConfig.TaskReference](
		f.client.VmAntiAffinityPoliciesApiInstance.UpdateVmAntiAffinityPolicyById(&uuid, &policy),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update anti-affinity policy with UUID %s: %w", uuid, err)
	}

	waiter := NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, f.GetAntiAffinityPolicy)
	return waiter, nil
}

// DeleteAntiAffinityPolicy deletes the anti-affinity policy with the given UUID.
func (f *FacadeV4Client) DeleteAntiAffinityPolicy(uuid string) (facade.TaskWaiter[facade.NoEntity], error) {
	policy, args, err := GetEntityAndEtag(
		f.GetAntiAffinityPolicy(uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get anti-affinity policy with UUID %s: %w", uuid, err)
	}

	if policy == nil {
		return nil, fmt.Errorf("no anti-affinity policy found with UUID %s", uuid)
	}

	taskRef, err := CallAPI[*v4policies.DeleteVmAntiAffinityPolicyApiResponse, v4VmmConfig.TaskReference](
		f.client.VmAntiAffinityPoliciesApiInstance.DeleteVmAntiAffinityPolicyById(&uuid, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete anti-affinity policy with UUID %s: %w", uuid, err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for deleted anti-affinity policy with UUID %s", uuid)
	}

	waiter := NewFacadeV4TaskWaiter(*taskRef.ExtId, f.client, facade.NoEntityGetter)
	return waiter, nil
}
