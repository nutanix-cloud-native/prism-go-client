package v4

import (
	"context"
	"testing"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	policyModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
	"github.com/stretchr/testify/assert"
)

func TestNewAntiAffinityPoliciesService(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	assert.NotNil(t, service)
	assert.IsType(t, &AntiAffinityPoliciesService{}, service)
}

func TestAntiAffinityPoliciesService_Get(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	result, err := service.Get(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")
}

func TestAntiAffinityPoliciesService_List(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	tests := []struct {
		name string
		opts []converged.ODataOption
	}{
		{
			name: "no options",
			opts: []converged.ODataOption{},
		},
		{
			name: "with page option",
			opts: []converged.ODataOption{converged.WithPage(1)},
		},
		{
			name: "with limit option",
			opts: []converged.ODataOption{converged.WithLimit(10)},
		},
		{
			name: "with filter option",
			opts: []converged.ODataOption{converged.WithFilter("name eq 'test'")},
		},
		{
			name: "with order by option",
			opts: []converged.ODataOption{converged.WithOrderBy("name")},
		},
		{
			name: "with expand option",
			opts: []converged.ODataOption{converged.WithExpand("resources")},
		},
		{
			name: "with select option",
			opts: []converged.ODataOption{converged.WithSelect("name,uuid")},
		},
		{
			name: "with apply option",
			opts: []converged.ODataOption{converged.WithApply("groupby((name))")},
		},
		{
			name: "with multiple options",
			opts: []converged.ODataOption{
				converged.WithPage(1),
				converged.WithLimit(10),
				converged.WithFilter("name eq 'test'"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.List(ctx, tt.opts...)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "not implemented")
		})
	}
}

func TestAntiAffinityPoliciesService_NewIterator(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)

	tests := []struct {
		name string
		opts []converged.ODataOption
	}{
		{
			name: "no options",
			opts: []converged.ODataOption{},
		},
		{
			name: "with filter option",
			opts: []converged.ODataOption{converged.WithFilter("name eq 'test'")},
		},
		{
			name: "with order by option",
			opts: []converged.ODataOption{converged.WithOrderBy("name")},
		},
		{
			name: "with expand option",
			opts: []converged.ODataOption{converged.WithExpand("resources")},
		},
		{
			name: "with select option",
			opts: []converged.ODataOption{converged.WithSelect("name,uuid")},
		},
		{
			name: "with multiple options",
			opts: []converged.ODataOption{
				converged.WithFilter("name eq 'test'"),
				converged.WithOrderBy("name"),
				converged.WithExpand("resources"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			iterator := service.NewIterator(ctx, tt.opts...)
			assert.Nil(t, iterator)
		})
	}
}

func TestAntiAffinityPoliciesService_Create(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	// Test with nil entity
	result, err := service.Create(ctx, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")

	// Test with valid entity
	entity := &policyModels.VmAntiAffinityPolicy{}
	result, err = service.Create(ctx, entity)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")
}

func TestAntiAffinityPoliciesService_Update(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	// Test with empty UUID
	result, err := service.Update(ctx, "", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")

	// Test with valid UUID and nil entity
	result, err = service.Update(ctx, "test-uuid", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")

	// Test with valid UUID and entity
	entity := &policyModels.VmAntiAffinityPolicy{}
	result, err = service.Update(ctx, "test-uuid", entity)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")
}

func TestAntiAffinityPoliciesService_Delete(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	// Test with empty UUID
	err := service.Delete(ctx, "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	// Test with valid UUID
	err = service.Delete(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
}

func TestAntiAffinityPoliciesService_CreateAsync(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	// Test with nil entity
	result, err := service.CreateAsync(ctx, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")

	// Test with valid entity
	entity := &policyModels.VmAntiAffinityPolicy{}
	result, err = service.CreateAsync(ctx, entity)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")
}

func TestAntiAffinityPoliciesService_UpdateAsync(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	// Test with empty UUID
	result, err := service.UpdateAsync(ctx, "", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")

	// Test with valid UUID and nil entity
	result, err = service.UpdateAsync(ctx, "test-uuid", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")

	// Test with valid UUID and entity
	entity := &policyModels.VmAntiAffinityPolicy{}
	result, err = service.UpdateAsync(ctx, "test-uuid", entity)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")
}

func TestAntiAffinityPoliciesService_DeleteAsync(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	// Test with empty UUID
	result, err := service.DeleteAsync(ctx, "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")

	// Test with valid UUID
	result, err = service.DeleteAsync(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")
}

// TestAntiAffinityPoliciesInterface tests the AntiAffinityPolicies interface implementation
func TestAntiAffinityPoliciesInterface(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)

	// Test that the service implements the AntiAffinityPolicies interface
	var _ converged.AntiAffinityPolicies[policyModels.VmAntiAffinityPolicy] = service

	ctx := context.Background()

	// Test Get method
	_, err := service.Get(ctx, "test-uuid")
	assert.Error(t, err)

	// Test List method
	_, err = service.List(ctx)
	assert.Error(t, err)

	// Test NewIterator method
	iterator := service.NewIterator(ctx)
	assert.Nil(t, iterator)

	// Test Create method
	_, err = service.Create(ctx, &policyModels.VmAntiAffinityPolicy{})
	assert.Error(t, err)

	// Test Update method
	_, err = service.Update(ctx, "test-uuid", &policyModels.VmAntiAffinityPolicy{})
	assert.Error(t, err)

	// Test Delete method
	err = service.Delete(ctx, "test-uuid")
	assert.Error(t, err)

	// Test CreateAsync method
	_, err = service.CreateAsync(ctx, &policyModels.VmAntiAffinityPolicy{})
	assert.Error(t, err)

	// Test UpdateAsync method
	_, err = service.UpdateAsync(ctx, "test-uuid", &policyModels.VmAntiAffinityPolicy{})
	assert.Error(t, err)

	// Test DeleteAsync method
	_, err = service.DeleteAsync(ctx, "test-uuid")
	assert.Error(t, err)
}

// TestAntiAffinityPoliciesService_EdgeCases tests edge cases and error conditions
func TestAntiAffinityPoliciesService_EdgeCases(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	t.Run("empty uuid for Get", func(t *testing.T) {
		result, err := service.Get(ctx, "")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

// TestAntiAffinityPoliciesService_ClientDependency tests that the service properly handles client dependency
func TestAntiAffinityPoliciesService_ClientDependency(t *testing.T) {
	// Test with nil client
	service := NewAntiAffinityPoliciesService(nil)
	assert.NotNil(t, service)
	assert.Nil(t, service.client)

	// Test that all methods still return "not implemented" regardless of client state
	ctx := context.Background()

	_, err := service.Get(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, err = service.List(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	iterator := service.NewIterator(ctx)
	assert.Nil(t, iterator)

	_, err = service.Create(ctx, &policyModels.VmAntiAffinityPolicy{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, err = service.Update(ctx, "test-uuid", &policyModels.VmAntiAffinityPolicy{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	err = service.Delete(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, err = service.CreateAsync(ctx, &policyModels.VmAntiAffinityPolicy{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, err = service.UpdateAsync(ctx, "test-uuid", &policyModels.VmAntiAffinityPolicy{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, err = service.DeleteAsync(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
}

// TestAntiAffinityPoliciesService_Consistency tests that all methods consistently return "not implemented"
func TestAntiAffinityPoliciesService_Consistency(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	// Test that all methods return consistent "not implemented" errors
	methods := []struct {
		name string
		fn   func() error
	}{
		{
			name: "Get",
			fn: func() error {
				_, err := service.Get(ctx, "test-uuid")
				return err
			},
		},
		{
			name: "List",
			fn: func() error {
				_, err := service.List(ctx)
				return err
			},
		},
		{
			name: "Create",
			fn: func() error {
				_, err := service.Create(ctx, &policyModels.VmAntiAffinityPolicy{})
				return err
			},
		},
		{
			name: "Update",
			fn: func() error {
				_, err := service.Update(ctx, "test-uuid", &policyModels.VmAntiAffinityPolicy{})
				return err
			},
		},
		{
			name: "Delete",
			fn: func() error {
				err := service.Delete(ctx, "test-uuid")
				return err
			},
		},
		{
			name: "CreateAsync",
			fn: func() error {
				_, err := service.CreateAsync(ctx, &policyModels.VmAntiAffinityPolicy{})
				return err
			},
		},
		{
			name: "UpdateAsync",
			fn: func() error {
				_, err := service.UpdateAsync(ctx, "test-uuid", &policyModels.VmAntiAffinityPolicy{})
				return err
			},
		},
		{
			name: "DeleteAsync",
			fn: func() error {
				_, err := service.DeleteAsync(ctx, "test-uuid")
				return err
			},
		},
	}

	for _, method := range methods {
		t.Run(method.name, func(t *testing.T) {
			err := method.fn()
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "not implemented")
		})
	}
}

// TestAntiAffinityPoliciesService_TypeSafety tests type safety of the service
func TestAntiAffinityPoliciesService_TypeSafety(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)

	// Test that the service is properly typed
	assert.IsType(t, &AntiAffinityPoliciesService{}, service)

	// Test that the service implements the correct interface
	var antiAffinityPoliciesInterface converged.AntiAffinityPolicies[policyModels.VmAntiAffinityPolicy] = service
	assert.NotNil(t, antiAffinityPoliciesInterface)

	// Test that the service has the expected client field
	assert.Nil(t, service.client) // Should be nil since we passed nil
}

// TestAntiAffinityPoliciesService_ODataOptions tests various OData option combinations
func TestAntiAffinityPoliciesService_ODataOptions(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	// Test complex OData option combinations
	complexOptions := []converged.ODataOption{
		converged.WithPage(2),
		converged.WithLimit(50),
		converged.WithFilter("name eq 'production' and status eq 'active'"),
		converged.WithOrderBy("created_time desc"),
		converged.WithExpand("resources,metadata"),
		converged.WithSelect("name,uuid,status,created_time"),
		converged.WithApply("groupby((status))"),
	}

	t.Run("complex options for List", func(t *testing.T) {
		result, err := service.List(ctx, complexOptions...)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "not implemented")
	})

	t.Run("complex options for NewIterator", func(t *testing.T) {
		iterator := service.NewIterator(ctx, complexOptions...)
		assert.Nil(t, iterator)
	})
}

// TestAntiAffinityPoliciesService_ErrorMessages tests that error messages are consistent
func TestAntiAffinityPoliciesService_ErrorMessages(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	// Test that all methods return the same "not implemented" error message
	expectedError := "not implemented"

	_, err := service.Get(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedError)

	_, err = service.List(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedError)

	_, err = service.Create(ctx, &policyModels.VmAntiAffinityPolicy{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedError)

	_, err = service.Update(ctx, "test-uuid", &policyModels.VmAntiAffinityPolicy{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedError)

	err = service.Delete(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedError)

	_, err = service.CreateAsync(ctx, &policyModels.VmAntiAffinityPolicy{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedError)

	_, err = service.UpdateAsync(ctx, "test-uuid", &policyModels.VmAntiAffinityPolicy{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedError)

	_, err = service.DeleteAsync(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedError)
}

// TestAntiAffinityPoliciesService_PolicySpecificOptions tests policy-specific OData options
func TestAntiAffinityPoliciesService_PolicySpecificOptions(t *testing.T) {
	service := NewAntiAffinityPoliciesService(nil)
	ctx := context.Background()

	// Test policy-specific filter options
	policyOptions := []converged.ODataOption{
		converged.WithFilter("policy_type eq 'ANTI_AFFINITY'"),
		converged.WithFilter("enabled eq true"),
		converged.WithFilter("vm_count gt 1"),
		converged.WithExpand("vm_references,policy_rules"),
		converged.WithSelect("name,uuid,policy_type,enabled,vm_count"),
	}

	t.Run("policy-specific options for List", func(t *testing.T) {
		result, err := service.List(ctx, policyOptions...)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "not implemented")
	})

	t.Run("policy-specific options for NewIterator", func(t *testing.T) {
		iterator := service.NewIterator(ctx, policyOptions...)
		assert.Nil(t, iterator)
	})
}
