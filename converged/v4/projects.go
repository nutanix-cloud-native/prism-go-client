package v4

import (
	"context"
	"errors"
	"fmt"
	"strings"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	multidomainModels "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/config"
	projectRequest "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/request/projects"
)

// ProjectsService provides implementation for Projects API operations.
type ProjectsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewProjectsService creates a new ProjectsService instance.
func NewProjectsService(client *v4prismGoClient.Client) *ProjectsService {
	return &ProjectsService{client: client, entitiesName: "project"}
}

// Get returns the project for the given UUID.
func (s *ProjectsService) Get(ctx context.Context, uuid string) (*multidomainModels.Project, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	return GenericGetEntity[*multidomainModels.GetProjectApiResponse, multidomainModels.Project](
		func() (*multidomainModels.GetProjectApiResponse, error) {
			return s.client.ProjectsApiInstance.ServiceClient.GetProjectById(ctx, &projectRequest.GetProjectByIdRequest{
				ExtId: &uuid,
			})
		},
		s.entitiesName,
	)
}

// List returns a list of projects.
func (s *ProjectsService) List(ctx context.Context, opts ...converged.ODataOption) ([]multidomainModels.Project, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing Projects")
	}

	return GenericListEntities[*multidomainModels.ListProjectsApiResponse, multidomainModels.Project](
		func(reqParams *V4ODataParams) (*multidomainModels.ListProjectsApiResponse, error) {
			return s.client.ProjectsApiInstance.ServiceClient.ListProjects(ctx, &projectRequest.ListProjectsRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Select_:  reqParams.Select,
			})
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing projects.
func (s *ProjectsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[multidomainModels.Project] {
	if s.client == nil {
		return nil
	}
	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return ErrorOnlyIterator[multidomainModels.Project](fmt.Errorf("failed to convert options to V4ODataParams: %w", err))
	}
	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return ErrorOnlyIterator[multidomainModels.Project](fmt.Errorf("apply and expand options are not supported for listing Projects"))
	}
	return GenericNewIterator[*multidomainModels.ListProjectsApiResponse, multidomainModels.Project](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*multidomainModels.ListProjectsApiResponse, error) {
			return s.client.ProjectsApiInstance.ServiceClient.ListProjects(ctx, &projectRequest.ListProjectsRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Select_:  reqParams.Select,
			})
		},
		opts,
		s.entitiesName,
	)
}

// GetDefaultProject returns the default project.
func (s *ProjectsService) GetDefaultProject(ctx context.Context) (*multidomainModels.Project, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	result, err := s.List(ctx, converged.WithFilter("isDefault eq true and isSystemDefined eq true"))
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		// The system default project always exists on PC. An empty filtered list
		// means the caller authenticated but is not allowed to see it — same
		// meaning as HTTP 403 (ErrUnauthorized), even though PC returns 200 + [].
		return nil, &converged.APIError{
			Kind:    converged.ErrUnauthorized,
			Message: "credentials are not authorized to access the default project",
		}
	}

	if len(result) > 1 {
		return nil, errors.New("multiple default projects found")
	}

	return &result[0], nil
}

// GetByName resolves a project by name.
func (s *ProjectsService) GetByName(ctx context.Context, name string) (*multidomainModels.Project, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	needle := strings.TrimSpace(name)
	if needle == "" {
		return nil, errors.New("project name cannot be empty")
	}

	result, err := s.List(ctx, converged.WithFilter(fmt.Sprintf("name eq '%s'", strings.ReplaceAll(needle, "'", "''"))))
	if err != nil {
		return nil, fmt.Errorf("failed to lookup project by name %q: %w", needle, err)
	}

	if len(result) == 0 {
		return nil, &converged.APIError{Kind: converged.ErrNotFound, Message: fmt.Sprintf("project %q not found", needle)}
	}
	if len(result) > 1 {
		return nil, fmt.Errorf("project %q is ambiguous (%d matches)", needle, len(result))
	}

	return &result[0], nil
}
