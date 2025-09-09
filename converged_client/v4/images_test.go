package v4

import (
	"context"
	"testing"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	imageModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
	"github.com/stretchr/testify/assert"
)

func TestNewImagesService(t *testing.T) {
	service := NewImagesService(nil)
	assert.NotNil(t, service)
	assert.IsType(t, &ImagesService{}, service)
}

func TestImagesService_Get(t *testing.T) {
	service := NewImagesService(nil)
	ctx := context.Background()

	result, err := service.Get(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not implemented")
}

func TestImagesService_List(t *testing.T) {
	service := NewImagesService(nil)
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

func TestImagesService_NewIterator(t *testing.T) {
	service := NewImagesService(nil)

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

// TestImagesInterface tests the Images interface implementation
func TestImagesInterface(t *testing.T) {
	service := NewImagesService(nil)

	// Test that the service implements the Images interface
	var _ converged.Images[imageModels.Image] = service

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
}

// TestImagesService_EdgeCases tests edge cases and error conditions
func TestImagesService_EdgeCases(t *testing.T) {
	service := NewImagesService(nil)
	ctx := context.Background()

	t.Run("empty uuid", func(t *testing.T) {
		result, err := service.Get(ctx, "")
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("empty uuid for list operations", func(t *testing.T) {
		// Test that list operations don't require UUID but still return not implemented
		result, err := service.List(ctx)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "not implemented")
	})
}

// TestImagesService_ClientDependency tests that the service properly handles client dependency
func TestImagesService_ClientDependency(t *testing.T) {
	// Test with nil client
	service := NewImagesService(nil)
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
}

// TestImagesService_Consistency tests that all methods consistently return "not implemented"
func TestImagesService_Consistency(t *testing.T) {
	service := NewImagesService(nil)
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
	}

	for _, method := range methods {
		t.Run(method.name, func(t *testing.T) {
			err := method.fn()
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "not implemented")
		})
	}
}

// TestImagesService_TypeSafety tests type safety of the service
func TestImagesService_TypeSafety(t *testing.T) {
	service := NewImagesService(nil)

	// Test that the service is properly typed
	assert.IsType(t, &ImagesService{}, service)

	// Test that the service implements the correct interface
	var imagesInterface converged.Images[imageModels.Image] = service
	assert.NotNil(t, imagesInterface)

	// Test that the service has the expected client field
	assert.Nil(t, service.client) // Should be nil since we passed nil
}
