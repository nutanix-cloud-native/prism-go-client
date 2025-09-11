package test

import (
	"context"
	"testing"

	models "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	v4ConvergedClient "github.com/nutanix-cloud-native/prism-go-client/converged_client/v4"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

var client *v4ConvergedClient.Client

func initializeConvergedClient(t *testing.T) {
	if client != nil {
		return
	}

	var err error
	creds := testhelpers.CredentialsFromEnvironment(t)

	client, err = v4ConvergedClient.NewClient(creds)
	assert.Nil(t, err)
}

func TestCategoriesService_Get(t *testing.T) {
	initializeConvergedClient(t)
	ctx := context.Background()

	categories, err := client.Categories.List(ctx, models.WithPage(0))
	assert.Nil(t, err)
	assert.True(t, len(categories) > 0)

	_, err = client.Categories.Get(ctx, *categories[0].ExtId)
	assert.Nil(t, err)
}

func TestCategoriesService_List(t *testing.T) {
	initializeConvergedClient(t)
	ctx := context.Background()

	testCases := []struct {
		name string
		opts []models.ODataOption
	}{
		{
			name: "no options",
			opts: []models.ODataOption{},
		},
		{
			name: "with page option",
			opts: []models.ODataOption{models.WithPage(1)},
		},
		{
			name: "with limit option",
			opts: []models.ODataOption{models.WithLimit(10)},
		},
		{
			name: "with filter option",
			opts: []models.ODataOption{models.WithFilter("name eq 'test'")},
		},
		{
			name: "with order by option",
			opts: []models.ODataOption{models.WithOrderBy("name")},
		},
		{
			name: "with expand option",
			opts: []models.ODataOption{models.WithExpand("metadata")},
		},
		{
			name: "with select option",
			opts: []models.ODataOption{models.WithSelect("name")},
		},
		{
			name: "with apply option",
			opts: []models.ODataOption{models.WithApply("groupby((name))")},
		},
		{
			name: "with multiple options",
			opts: []models.ODataOption{
				models.WithPage(1),
				models.WithLimit(10),
				models.WithFilter("name eq 'test'"),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.Categories.List(ctx, tt.opts...)
			assert.Nil(t, err)
		})
	}
}

func TestCategoriesService_NewIterator(t *testing.T) {
	initializeConvergedClient(t)
	ctx := context.Background()

	testCases := []struct {
		name string
		opts []models.ODataOption
	}{
		{
			name: "no options",
			opts: []models.ODataOption{},
		},
		{
			name: "with filter option",
			opts: []models.ODataOption{models.WithFilter("name eq 'test'")},
		},
		{
			name: "with order by option",
			opts: []models.ODataOption{models.WithOrderBy("name")},
		},
		{
			name: "with expand option",
			opts: []models.ODataOption{models.WithExpand("metadata")},
		},
		{
			name: "with select option",
			opts: []models.ODataOption{models.WithSelect("name")},
		},
		{
			name: "with multiple options",
			opts: []models.ODataOption{
				models.WithFilter("name eq 'test'"),
				models.WithOrderBy("name"),
				models.WithExpand("metadata"),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			iterator := client.Categories.NewIterator(ctx, tt.opts...)
			assert.NotNil(t, iterator)
			for _, err := range iterator {
				assert.Nil(t, err)
			}
		})
	}
}
