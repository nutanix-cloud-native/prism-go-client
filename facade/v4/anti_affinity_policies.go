package v4

import (
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	v4VmmConfig "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	v4policies "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
)

// Implements the FacadeClientV4 interface
func (f *FacadeV4Client) GetAntiAffinityPolicy(uuid string) (*v4policies.VmAntiAffinityPolicy, error) {
	return CommonGetEntity[*v4policies.GetVmAntiAffinityPolicyApiResponse, v4policies.VmAntiAffinityPolicy](
		func() (*v4policies.GetVmAntiAffinityPolicyApiResponse, error) {
			return f.client.VmAntiAffinityPoliciesApiInstance.GetVmAntiAffinityPolicyById(&uuid)
		},
		"anti-affinity policy",
	)
}

func (f *FacadeV4Client) ListAntiAffinityPolicies(opts ...facade.ODataOption) ([]v4policies.VmAntiAffinityPolicy, error) {
	return CommonListEntities[*v4policies.ListVmAntiAffinityPoliciesApiResponse, v4policies.VmAntiAffinityPolicy](
		func(reqParams *V4ODataParams) (*v4policies.ListVmAntiAffinityPoliciesApiResponse, error) {
			return f.client.VmAntiAffinityPoliciesApiInstance.ListVmAntiAffinityPolicies(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
			)
		},
		opts,
		"anti-affinity policies",
	)
}

func (f *FacadeV4Client) ListAllAntiAffinityPolicies(filterParam *string, orderbyParam *string) ([]v4policies.VmAntiAffinityPolicy, error) {
	reqParams := &V4ODataParams{
		Filter:  filterParam,
		OrderBy: orderbyParam,
	}

	return CommonListAllEntities[*v4policies.ListVmAntiAffinityPoliciesApiResponse, v4policies.VmAntiAffinityPolicy](
		func(reqParams *V4ODataParams) (*v4policies.ListVmAntiAffinityPoliciesApiResponse, error) {
			return f.client.VmAntiAffinityPoliciesApiInstance.ListVmAntiAffinityPolicies(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
			)
		},
		reqParams,
		"anti-affinity policies",
	)
}

func (f *FacadeV4Client) GetListIteratorAntiAffinityPolicies(opts ...facade.ODataOption) facade.ODataListIterator[v4policies.VmAntiAffinityPolicy] {
	return CommonGetListIterator[*v4policies.ListVmAntiAffinityPoliciesApiResponse, v4policies.VmAntiAffinityPolicy](
		f,
		func(reqParams *V4ODataParams) (*v4policies.ListVmAntiAffinityPoliciesApiResponse, error) {
			return f.client.VmAntiAffinityPoliciesApiInstance.ListVmAntiAffinityPolicies(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
			)
		},
		opts,
		"anti-affinity policies",
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
