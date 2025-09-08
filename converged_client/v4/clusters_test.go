package v4

import (
	"context"
	"testing"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	clustermgmtModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	"github.com/stretchr/testify/assert"
)

func TestNewClustersService(t *testing.T) {
	service := NewClustersService(nil)
	assert.NotNil(t, service)
	assert.IsType(t, &ClustersService{}, service)
}

func TestClustersService_Get(t *testing.T) {
	service := NewClustersService(nil)
	ctx := context.Background()

	result, err := service.Get(ctx, "test-uuid")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "client is not initialized")
}

func TestClustersService_List(t *testing.T) {
	service := NewClustersService(nil)
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
			assert.Contains(t, err.Error(), "client is not initialized")
		})
	}
}

func TestClustersService_ListAll(t *testing.T) {
	service := NewClustersService(nil)
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
			assert.Contains(t, err.Error(), "client is not initialized")
		})
	}
}

func TestClustersService_NewIterator(t *testing.T) {
	service := NewClustersService(nil)

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

func TestClustersService_ContextHandling(t *testing.T) {
	service := NewClustersService(nil)

	t.Run("context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		result, err := service.Get(ctx, "test-uuid")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("context timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 0) // Immediate timeout
		defer cancel()

		result, err := service.Get(ctx, "test-uuid")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "client is not initialized")
	})
}

func TestClustersService_InterfaceCompliance(t *testing.T) {
	service := NewClustersService(nil)

	// Test that ClustersService implements the Getter interface
	var _ converged.Getter[clustermgmtModels.Cluster] = service

	// Test that ClustersService implements the Lister interface
	var _ converged.Lister[clustermgmtModels.Cluster] = service
}

func TestClustersService_ErrorConsistency(t *testing.T) {
	service := NewClustersService(nil)
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
	}

	for _, method := range methods {
		t.Run(method.name, func(t *testing.T) {
			err := method.fn()
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "client is not initialized")
		})
	}
}

func TestClustersService_MethodBehaviors(t *testing.T) {
	service := NewClustersService(nil)
	ctx := context.Background()

	t.Run("Get with empty UUID", func(t *testing.T) {
		result, err := service.Get(ctx, "")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("Get with valid UUID", func(t *testing.T) {
		result, err := service.Get(ctx, "test-uuid-123")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("List with nil options", func(t *testing.T) {
		result, err := service.List(ctx, nil)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("ListAll with nil options", func(t *testing.T) {
		result, err := service.ListAll(ctx, nil)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("NewIterator with nil options", func(t *testing.T) {
		iterator := service.NewIterator(nil)
		assert.Nil(t, iterator)
	})
}

func TestClustersService_NilService(t *testing.T) {
	var service *ClustersService
	ctx := context.Background()

	// Test that calling methods on nil service panics or handles gracefully
	t.Run("nil service Get", func(t *testing.T) {
		assert.Panics(t, func() {
			service.Get(ctx, "test")
		})
	})

	t.Run("nil service List", func(t *testing.T) {
		assert.Panics(t, func() {
			service.List(ctx)
		})
	})

	t.Run("nil service ListAll", func(t *testing.T) {
		assert.Panics(t, func() {
			service.ListAll(ctx)
		})
	})

	t.Run("nil service NewIterator", func(t *testing.T) {
		assert.Panics(t, func() {
			service.NewIterator()
		})
	})
}

func TestClustersService_ServiceInstantiation(t *testing.T) {
	t.Run("service with nil client", func(t *testing.T) {
		service := NewClustersService(nil)
		assert.NotNil(t, service)
		assert.Nil(t, service.client)
	})

	t.Run("service type assertion", func(t *testing.T) {
		service := NewClustersService(nil)
		assert.IsType(t, &ClustersService{}, service)

		// Verify the service has the expected structure
		assert.NotNil(t, service)
		assert.Nil(t, service.client)
	})
}
