package facade

import (
	v4policies "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
)

type AntiAffinityPolicyFacadeV4 interface {
	// GetAntiAffinityPolicy returns the anti-affinity policy for the given UUID.
	GetAntiAffinityPolicy(uuid string) (*v4policies.VmAntiAffinityPolicy, error)

	// ListAntiAffinityPolicies returns a list of anti-affinity policies.
	ListAntiAffinityPolicies(opts ...ODataOption) ([]v4policies.VmAntiAffinityPolicy, error)

	// ListAllAntiAffinityPolicies returns all anti-affinity policies without pagination.
	ListAllAntiAffinityPolicies(filterParam *string, orderbyParam *string) ([]v4policies.VmAntiAffinityPolicy, error)

	// GetListIteratorAntiAffinityPolicies returns an iterator for listing anti-affinity policies.
	GetListIteratorAntiAffinityPolicies(opts ...ODataOption) (ODataListIterator[v4policies.VmAntiAffinityPolicy], error)

	// CreateAntiAffinityPolicy creates a new anti-affinity policy.
	CreateAntiAffinityPolicy(policy v4policies.VmAntiAffinityPolicy) (TaskWaiter[v4policies.VmAntiAffinityPolicy], error)

	// UpdateAntiAffinityPolicy updates an existing anti-affinity policy.
	UpdateAntiAffinityPolicy(uuid string, policy v4policies.VmAntiAffinityPolicy) (TaskWaiter[v4policies.VmAntiAffinityPolicy], error)

	// DeleteAntiAffinityPolicy deletes the anti-affinity policy with the given UUID.
	DeleteAntiAffinityPolicy(uuid string) (TaskWaiter[NoEntity], error)
}
