package v4

import (
	"context"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	policyModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
)

// AntiAffinityPoliciesService provides default "not implemented" implementation for all AntiAffinityPolicies interface methods.
type AntiAffinityPoliciesService struct {
	client *v4prismGoClient.Client
}

// NewAntiAffinityPoliciesService creates a new AntiAffinityPoliciesService instance.
func NewAntiAffinityPoliciesService(client *v4prismGoClient.Client) *AntiAffinityPoliciesService {
	return &AntiAffinityPoliciesService{client: client}
}

// Get returns the anti-affinity policy for the given UUID.
func (s *AntiAffinityPoliciesService) Get(ctx context.Context, uuid string) (*policyModels.VmAntiAffinityPolicy, error) {
	return nil, fmt.Errorf("not implemented")
}

// List returns a list of anti-affinity policies.
func (s *AntiAffinityPoliciesService) List(ctx context.Context, opts ...converged.ODataOption) ([]policyModels.VmAntiAffinityPolicy, error) {
	return nil, fmt.Errorf("not implemented")
}

// NewIterator returns an iterator for listing anti-affinity policies.
func (s *AntiAffinityPoliciesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[policyModels.VmAntiAffinityPolicy] {
	return nil
}

// Create creates a new anti-affinity policy.
func (s *AntiAffinityPoliciesService) Create(ctx context.Context, entity *policyModels.VmAntiAffinityPolicy) (*policyModels.VmAntiAffinityPolicy, error) {
	return nil, fmt.Errorf("not implemented")
}

// Update updates an existing anti-affinity policy.
func (s *AntiAffinityPoliciesService) Update(ctx context.Context, uuid string, entity *policyModels.VmAntiAffinityPolicy) (*policyModels.VmAntiAffinityPolicy, error) {
	return nil, fmt.Errorf("not implemented")
}

// Delete deletes an existing anti-affinity policy.
func (s *AntiAffinityPoliciesService) Delete(ctx context.Context, uuid string) error {
	return fmt.Errorf("not implemented")
}

// CreateAsync creates a new anti-affinity policy asynchronously.
func (s *AntiAffinityPoliciesService) CreateAsync(ctx context.Context, entity *policyModels.VmAntiAffinityPolicy) (converged.Operation[policyModels.VmAntiAffinityPolicy], error) {
	return nil, fmt.Errorf("not implemented")
}

// UpdateAsync updates an existing anti-affinity policy asynchronously.
func (s *AntiAffinityPoliciesService) UpdateAsync(ctx context.Context, uuid string, entity *policyModels.VmAntiAffinityPolicy) (converged.Operation[policyModels.VmAntiAffinityPolicy], error) {
	return nil, fmt.Errorf("not implemented")
}

// DeleteAsync deletes an existing anti-affinity policy asynchronously.
func (s *AntiAffinityPoliciesService) DeleteAsync(ctx context.Context, uuid string) (converged.Operation[converged.NoEntity], error) {
	return nil, fmt.Errorf("not implemented")
}
