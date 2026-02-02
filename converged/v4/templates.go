package v4

import (
	"context"
	"errors"
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	vmmPrismModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	templateModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// TemplatesService provides implementation for all Templates interface methods.
type TemplatesService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewTemplatesService creates a new TemplatesService instance.
func NewTemplatesService(client *v4prismGoClient.Client) *TemplatesService {
	return &TemplatesService{client: client, entitiesName: "template"}
}

// Get returns the template for the given UUID.
func (s *TemplatesService) Get(ctx context.Context, uuid string) (*templateModels.Template, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*templateModels.GetTemplateApiResponse, templateModels.Template](
		func() (*templateModels.GetTemplateApiResponse, error) {
			return s.client.TemplatesApiInstance.GetTemplateById(&uuid)
		},
		s.entitiesName,
	)
}

// List returns a list of templates.
func (s *TemplatesService) List(ctx context.Context, opts ...converged.ODataOption) ([]templateModels.Template, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams.Expand != nil || myParams.Apply != nil {
		return nil, fmt.Errorf("expand and apply are not supported for listing templates")
	}

	return GenericListEntities[*templateModels.ListTemplatesApiResponse, templateModels.Template](
		func(reqParams *V4ODataParams) (*templateModels.ListTemplatesApiResponse, error) {
			return s.client.TemplatesApiInstance.ListTemplates(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing templates.
func (s *TemplatesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[templateModels.Template] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*templateModels.ListTemplatesApiResponse, templateModels.Template](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*templateModels.ListTemplatesApiResponse, error) {
			return s.client.TemplatesApiInstance.ListTemplates(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// Create creates a new template.
func (s *TemplatesService) Create(ctx context.Context, template *templateModels.Template) (*templateModels.Template, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	operation, err := s.CreateAsync(ctx, template)
	if err != nil {
		return nil, fmt.Errorf("failed to create template: %w", err)
	}

	result, err := operation.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create template: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("failed to create template: operation completed but no template returned")
	}

	return result[0], nil
}

// CreateAsync creates a new template asynchronously.
func (s *TemplatesService) CreateAsync(ctx context.Context, template *templateModels.Template) (converged.Operation[templateModels.Template], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	taskRef, err := CallAPI[*templateModels.CreateTemplateApiResponse, vmmPrismModels.TaskReference](
		s.client.TemplatesApiInstance.CreateTemplate(template),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create template: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for created template")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		s.Get,
	), nil
}
