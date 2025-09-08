package v4

import (
	"context"
	"testing"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	"github.com/stretchr/testify/assert"
	"k8s.io/utils/ptr"
)

func TestNewCategoriesService(t *testing.T) {
	service := NewCategoriesService(nil)
	assert.NotNil(t, service)
	assert.IsType(t, &CategoriesService{}, service)
}

func TestCategoriesService_Get(t *testing.T) {
	service := NewCategoriesService(nil)
	ctx := context.Background()

	result, err := service.Get(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")
}

func TestCategoriesService_List(t *testing.T) {
	service := NewCategoriesService(nil)
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
			opts: []converged.ODataOption{converged.WithOrderBy("name asc")},
		},
		{
			name: "with expand option",
			opts: []converged.ODataOption{converged.WithExpand("metadata")},
		},
		{
			name: "with select option",
			opts: []converged.ODataOption{converged.WithSelect("id,name")},
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
				converged.WithOrderBy("name asc"),
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

func TestCategoriesService_ListAll(t *testing.T) {
	service := NewCategoriesService(nil)
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
			opts: []converged.ODataOption{converged.WithOrderBy("name asc")},
		},
		{
			name: "with expand option",
			opts: []converged.ODataOption{converged.WithExpand("metadata")},
		},
		{
			name: "with select option",
			opts: []converged.ODataOption{converged.WithSelect("id,name")},
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
				converged.WithOrderBy("name asc"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ListAll(ctx, tt.opts...)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "not implemented")
		})
	}
}

func TestCategoriesService_NewIterator(t *testing.T) {
	service := NewCategoriesService(nil)

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
			opts: []converged.ODataOption{converged.WithOrderBy("name asc")},
		},
		{
			name: "with expand option",
			opts: []converged.ODataOption{converged.WithExpand("metadata")},
		},
		{
			name: "with select option",
			opts: []converged.ODataOption{converged.WithSelect("id,name")},
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
				converged.WithOrderBy("name asc"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator := service.NewIterator(tt.opts...)
			assert.Nil(t, iterator)
		})
	}
}

func TestCategoriesService_Create(t *testing.T) {
	service := NewCategoriesService(nil)
	ctx := context.Background()

	tests := []struct {
		name     string
		category *prismModels.Category
	}{
		{
			name:     "nil category",
			category: nil,
		},
		{
			name:     "empty category",
			category: &prismModels.Category{},
		},
		{
			name: "category with name",
			category: &prismModels.Category{
				Key: ptr.To("Test Category"),
			},
		},
		{
			name: "category with description",
			category: &prismModels.Category{
				Key:   ptr.To("Test Category"),
				Value: ptr.To("A test category"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Create(ctx, tt.category)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "not implemented")
		})
	}
}

func TestCategoriesService_Update(t *testing.T) {
	service := NewCategoriesService(nil)
	ctx := context.Background()

	tests := []struct {
		name     string
		uuid     string
		category *prismModels.Category
	}{
		{
			name:     "empty uuid",
			uuid:     "",
			category: &prismModels.Category{Key: ptr.To("Test")},
		},
		{
			name:     "valid uuid",
			uuid:     "test-uuid-123",
			category: &prismModels.Category{Key: ptr.To("Test")},
		},
		{
			name:     "nil category",
			uuid:     "test-uuid-123",
			category: nil,
		},
		{
			name: "category with name",
			uuid: "test-uuid-123",
			category: &prismModels.Category{
				Key: ptr.To("Updated Category"),
			},
		},
		{
			name: "category with description",
			uuid: "test-uuid-123",
			category: &prismModels.Category{
				Key:   ptr.To("Updated Category"),
				Value: ptr.To("An updated test category"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Update(ctx, tt.uuid, tt.category)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "not implemented")
		})
	}
}

func TestCategoriesService_Delete(t *testing.T) {
	service := NewCategoriesService(nil)
	ctx := context.Background()

	tests := []struct {
		name string
		uuid string
	}{
		{
			name: "empty uuid",
			uuid: "",
		},
		{
			name: "valid uuid",
			uuid: "test-uuid-123",
		},
		{
			name: "uuid with special characters",
			uuid: "test-uuid-!@#$%",
		},
		{
			name: "long uuid",
			uuid: "very-long-uuid-that-exceeds-normal-length-limits-and-should-still-be-handled-properly",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Delete(ctx, tt.uuid)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "not implemented")
		})
	}
}

func TestCategoriesService_ContextHandling(t *testing.T) {
	service := NewCategoriesService(nil)

	t.Run("context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		result, err := service.Get(ctx, "test-uuid")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "not implemented")
	})

	t.Run("context timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 0) // Immediate timeout
		defer cancel()

		result, err := service.Get(ctx, "test-uuid")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "not implemented")
	})
}

func TestCategoriesService_InterfaceCompliance(t *testing.T) {
	service := NewCategoriesService(nil)

	// Test that CategoriesService implements the Categories interface
	var _ converged.Categories[prismModels.Category] = service

	// Test that CategoriesService implements the Getter interface
	var _ converged.Getter[prismModels.Category] = service

	// Test that CategoriesService implements the Lister interface
	var _ converged.Lister[prismModels.Category] = service

	// Test that CategoriesService implements the Creator interface
	var _ converged.Creator[prismModels.Category] = service

	// Test that CategoriesService implements the Updater interface
	var _ converged.Updater[prismModels.Category] = service

	// Test that CategoriesService implements the Deleter interface
	var _ converged.Deleter[prismModels.Category] = service
}

func TestCategoriesService_ErrorConsistency(t *testing.T) {
	service := NewCategoriesService(nil)
	ctx := context.Background()

	// Test that all methods return consistent error messages
	methods := []struct {
		name string
		fn   func() error
	}{
		{
			name: "Get",
			fn: func() error {
				_, err := service.Get(ctx, "test")
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
			name: "ListAll",
			fn: func() error {
				_, err := service.ListAll(ctx)
				return err
			},
		},
		{
			name: "Create",
			fn: func() error {
				_, err := service.Create(ctx, &prismModels.Category{})
				return err
			},
		},
		{
			name: "Update",
			fn: func() error {
				_, err := service.Update(ctx, "test", &prismModels.Category{})
				return err
			},
		},
		{
			name: "Delete",
			fn: func() error {
				_, err := service.Delete(ctx, "test")
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
