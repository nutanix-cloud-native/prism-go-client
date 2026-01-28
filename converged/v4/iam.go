package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	iamAuthnModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	iamAuthzModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
)

// IAMService provides implementation for all IAM interface methods.
type IAMService struct {
	client                *v4prismGoClient.Client
	users                 *UsersService
	roles                 *RolesService
	authorizationPolicies *AuthorizationPoliciesService
	operations            *OperationsService
	userGroups            *UserGroupsService
	directoryServices     *DirectoryServicesService
	samlIdentityProviders *SAMLIdentityProvidersService
}

// NewIAMService creates a new IAMService instance.
func NewIAMService(client *v4prismGoClient.Client) *IAMService {
	return &IAMService{
		client:                client,
		users:                 NewUsersService(client),
		roles:                 NewRolesService(client),
		authorizationPolicies: NewAuthorizationPoliciesService(client),
		operations:            NewOperationsService(client),
		userGroups:            NewUserGroupsService(client),
		directoryServices:     NewDirectoryServicesService(client),
		samlIdentityProviders: NewSAMLIdentityProvidersService(client),
	}
}

// Users returns the Users service.
func (s *IAMService) Users() converged.Users[iamAuthnModels.User] {
	return s.users
}

// Roles returns the Roles service.
func (s *IAMService) Roles() converged.Roles[iamAuthzModels.Role] {
	return s.roles
}

// AuthorizationPolicies returns the AuthorizationPolicies service.
func (s *IAMService) AuthorizationPolicies() converged.AuthorizationPolicies[iamAuthzModels.AuthorizationPolicy] {
	return s.authorizationPolicies
}

// Operations returns the Operations service.
func (s *IAMService) Operations() converged.Operations[iamAuthzModels.Operation] {
	return s.operations
}

// UserGroups returns the UserGroups service.
func (s *IAMService) UserGroups() converged.UserGroups[iamAuthnModels.UserGroup] {
	return s.userGroups
}

// DirectoryServices returns the DirectoryServices service.
func (s *IAMService) DirectoryServices() converged.DirectoryServices[iamAuthnModels.DirectoryService] {
	return s.directoryServices
}

// SAMLIdentityProviders returns the SAML Identity Providers service.
func (s *IAMService) SAMLIdentityProviders() converged.SAML[iamAuthnModels.SamlIdentityProvider] {
	return s.samlIdentityProviders
}

// RolesService provides implementation for IAM Roles API operations.
type RolesService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewRolesService creates a new RolesService instance.
func NewRolesService(client *v4prismGoClient.Client) *RolesService {
	return &RolesService{
		client:       client,
		entitiesName: "role",
	}
}

// Get returns the role for the given UUID.
func (s *RolesService) Get(ctx context.Context, uuid string) (*iamAuthzModels.Role, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthzModels.GetRoleApiResponse, iamAuthzModels.Role](
		func() (*iamAuthzModels.GetRoleApiResponse, error) {
			return s.client.RolesApiInstance.GetRoleById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of roles.
func (s *RolesService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthzModels.Role, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing Roles")
	}

	return GenericListEntities[*iamAuthzModels.ListRolesApiResponse, iamAuthzModels.Role](
		func(reqParams *V4ODataParams) (*iamAuthzModels.ListRolesApiResponse, error) {
			return s.client.RolesApiInstance.ListRoles(
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

// NewIterator returns an iterator for listing roles.
func (s *RolesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthzModels.Role] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthzModels.ListRolesApiResponse, iamAuthzModels.Role](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthzModels.ListRolesApiResponse, error) {
			return s.client.RolesApiInstance.ListRoles(
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

// Create creates a new role.
func (s *RolesService) Create(ctx context.Context, role *iamAuthzModels.Role) (*iamAuthzModels.Role, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*iamAuthzModels.CreateRoleApiResponse, iamAuthzModels.Role](
		s.client.RolesApiInstance.CreateRole(role, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	return &result, nil
}

// Update updates an existing role.
func (s *RolesService) Update(ctx context.Context, uuid string, role *iamAuthzModels.Role) (*iamAuthzModels.Role, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*iamAuthzModels.UpdateRoleApiResponse, iamAuthzModels.Role](
		s.client.RolesApiInstance.UpdateRoleById(&uuid, role, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	return &result, nil
}

// Delete deletes an existing role.
func (s *RolesService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	_, err := s.client.RolesApiInstance.DeleteRoleById(&uuid, nil)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %w", s.entitiesName, err)
	}
	return nil
}

// AuthorizationPoliciesService provides implementation for IAM Authorization Policies API operations.
type AuthorizationPoliciesService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewAuthorizationPoliciesService creates a new AuthorizationPoliciesService instance.
func NewAuthorizationPoliciesService(client *v4prismGoClient.Client) *AuthorizationPoliciesService {
	return &AuthorizationPoliciesService{
		client:       client,
		entitiesName: "authorization policy",
	}
}

// Get returns the authorization policy for the given UUID.
func (s *AuthorizationPoliciesService) Get(ctx context.Context, uuid string) (*iamAuthzModels.AuthorizationPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthzModels.GetAuthorizationPolicyApiResponse, iamAuthzModels.AuthorizationPolicy](
		func() (*iamAuthzModels.GetAuthorizationPolicyApiResponse, error) {
			return s.client.AuthorizationPoliciesApiInstance.GetAuthorizationPolicyById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of authorization policies.
func (s *AuthorizationPoliciesService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthzModels.AuthorizationPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing Authorization Policies")
	}

	return GenericListEntities[*iamAuthzModels.ListAuthorizationPoliciesApiResponse, iamAuthzModels.AuthorizationPolicy](
		func(reqParams *V4ODataParams) (*iamAuthzModels.ListAuthorizationPoliciesApiResponse, error) {
			return s.client.AuthorizationPoliciesApiInstance.ListAuthorizationPolicies(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing authorization policies.
func (s *AuthorizationPoliciesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthzModels.AuthorizationPolicy] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthzModels.ListAuthorizationPoliciesApiResponse, iamAuthzModels.AuthorizationPolicy](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthzModels.ListAuthorizationPoliciesApiResponse, error) {
			return s.client.AuthorizationPoliciesApiInstance.ListAuthorizationPolicies(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// Create creates a new authorization policy.
func (s *AuthorizationPoliciesService) Create(ctx context.Context, policy *iamAuthzModels.AuthorizationPolicy) (*iamAuthzModels.AuthorizationPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*iamAuthzModels.CreateAuthorizationPolicyApiResponse, iamAuthzModels.AuthorizationPolicy](
		s.client.AuthorizationPoliciesApiInstance.CreateAuthorizationPolicy(policy, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	return &result, nil
}

// Update updates an existing authorization policy.
func (s *AuthorizationPoliciesService) Update(ctx context.Context, uuid string, policy *iamAuthzModels.AuthorizationPolicy) (*iamAuthzModels.AuthorizationPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*iamAuthzModels.UpdateAuthorizationPolicyApiResponse, iamAuthzModels.AuthorizationPolicy](
		s.client.AuthorizationPoliciesApiInstance.UpdateAuthorizationPolicyById(&uuid, policy, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	return &result, nil
}

// Delete deletes an existing authorization policy.
func (s *AuthorizationPoliciesService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	_, err := s.client.AuthorizationPoliciesApiInstance.DeleteAuthorizationPolicyById(&uuid, nil)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %w", s.entitiesName, err)
	}
	return nil
}

// OperationsService provides implementation for IAM Operations API operations.
type OperationsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewOperationsService creates a new OperationsService instance.
func NewOperationsService(client *v4prismGoClient.Client) *OperationsService {
	return &OperationsService{
		client:       client,
		entitiesName: "operation",
	}
}

// Get returns the operation for the given UUID.
func (s *OperationsService) Get(ctx context.Context, uuid string) (*iamAuthzModels.Operation, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthzModels.GetOperationApiResponse, iamAuthzModels.Operation](
		func() (*iamAuthzModels.GetOperationApiResponse, error) {
			return s.client.OperationsApiInstance.GetOperationById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of operations.
func (s *OperationsService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthzModels.Operation, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing Operations")
	}

	return GenericListEntities[*iamAuthzModels.ListOperationsApiResponse, iamAuthzModels.Operation](
		func(reqParams *V4ODataParams) (*iamAuthzModels.ListOperationsApiResponse, error) {
			return s.client.OperationsApiInstance.ListOperations(
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

// NewIterator returns an iterator for listing operations.
func (s *OperationsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthzModels.Operation] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthzModels.ListOperationsApiResponse, iamAuthzModels.Operation](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthzModels.ListOperationsApiResponse, error) {
			return s.client.OperationsApiInstance.ListOperations(
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

// UserGroupsService provides implementation for IAM User Groups API operations.
type UserGroupsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewUserGroupsService creates a new UserGroupsService instance.
func NewUserGroupsService(client *v4prismGoClient.Client) *UserGroupsService {
	return &UserGroupsService{
		client:       client,
		entitiesName: "user group",
	}
}

// Get returns the user group for the given UUID.
func (s *UserGroupsService) Get(ctx context.Context, uuid string) (*iamAuthnModels.UserGroup, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthnModels.GetUserGroupApiResponse, iamAuthnModels.UserGroup](
		func() (*iamAuthnModels.GetUserGroupApiResponse, error) {
			return s.client.UserGroupsApiInstance.GetUserGroupById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of user groups.
func (s *UserGroupsService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthnModels.UserGroup, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing User Groups")
	}

	return GenericListEntities[*iamAuthnModels.ListUserGroupsApiResponse, iamAuthnModels.UserGroup](
		func(reqParams *V4ODataParams) (*iamAuthnModels.ListUserGroupsApiResponse, error) {
			return s.client.UserGroupsApiInstance.ListUserGroups(
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

// NewIterator returns an iterator for listing user groups.
func (s *UserGroupsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthnModels.UserGroup] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthnModels.ListUserGroupsApiResponse, iamAuthnModels.UserGroup](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthnModels.ListUserGroupsApiResponse, error) {
			return s.client.UserGroupsApiInstance.ListUserGroups(
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

// Create creates a new user group.
func (s *UserGroupsService) Create(ctx context.Context, userGroup *iamAuthnModels.UserGroup) (*iamAuthnModels.UserGroup, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*iamAuthnModels.CreateUserGroupApiResponse, iamAuthnModels.UserGroup](
		s.client.UserGroupsApiInstance.CreateUserGroup(userGroup, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	return &result, nil
}

// Delete deletes an existing user group.
func (s *UserGroupsService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	_, err := s.client.UserGroupsApiInstance.DeleteUserGroupById(&uuid, nil)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %w", s.entitiesName, err)
	}
	return nil
}

// DirectoryServicesService provides implementation for IAM Directory Services API operations.
type DirectoryServicesService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewDirectoryServicesService creates a new DirectoryServicesService instance.
func NewDirectoryServicesService(client *v4prismGoClient.Client) *DirectoryServicesService {
	return &DirectoryServicesService{
		client:       client,
		entitiesName: "directory service",
	}
}

// Get returns the directory service for the given UUID.
func (s *DirectoryServicesService) Get(ctx context.Context, uuid string) (*iamAuthnModels.DirectoryService, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthnModels.GetDirectoryServiceApiResponse, iamAuthnModels.DirectoryService](
		func() (*iamAuthnModels.GetDirectoryServiceApiResponse, error) {
			return s.client.DirectoryServicesApiInstance.GetDirectoryServiceById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of directory services.
func (s *DirectoryServicesService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthnModels.DirectoryService, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing Directory Services")
	}

	return GenericListEntities[*iamAuthnModels.ListDirectoryServicesApiResponse, iamAuthnModels.DirectoryService](
		func(reqParams *V4ODataParams) (*iamAuthnModels.ListDirectoryServicesApiResponse, error) {
			return s.client.DirectoryServicesApiInstance.ListDirectoryServices(
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

// NewIterator returns an iterator for listing directory services.
func (s *DirectoryServicesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthnModels.DirectoryService] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthnModels.ListDirectoryServicesApiResponse, iamAuthnModels.DirectoryService](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthnModels.ListDirectoryServicesApiResponse, error) {
			return s.client.DirectoryServicesApiInstance.ListDirectoryServices(
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

// Create creates a new directory service.
func (s *DirectoryServicesService) Create(ctx context.Context, directoryService *iamAuthnModels.DirectoryService) (*iamAuthnModels.DirectoryService, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*iamAuthnModels.CreateDirectoryServiceApiResponse, iamAuthnModels.DirectoryService](
		s.client.DirectoryServicesApiInstance.CreateDirectoryService(directoryService, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	return &result, nil
}

// Update updates an existing directory service.
func (s *DirectoryServicesService) Update(ctx context.Context, uuid string, directoryService *iamAuthnModels.DirectoryService) (*iamAuthnModels.DirectoryService, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*iamAuthnModels.UpdateDirectoryServiceApiResponse, iamAuthnModels.DirectoryService](
		s.client.DirectoryServicesApiInstance.UpdateDirectoryServiceById(&uuid, directoryService, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	return &result, nil
}

// Delete deletes an existing directory service.
func (s *DirectoryServicesService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	_, err := s.client.DirectoryServicesApiInstance.DeleteDirectoryServiceById(&uuid, nil)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %w", s.entitiesName, err)
	}
	return nil
}

// SAMLIdentityProvidersService provides implementation for SAML Identity Providers API operations.
type SAMLIdentityProvidersService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewSAMLIdentityProvidersService creates a new SAMLIdentityProvidersService instance.
func NewSAMLIdentityProvidersService(client *v4prismGoClient.Client) *SAMLIdentityProvidersService {
	return &SAMLIdentityProvidersService{
		client:       client,
		entitiesName: "SAML identity provider",
	}
}

// Get returns the SAML identity provider for the given UUID.
func (s *SAMLIdentityProvidersService) Get(ctx context.Context, uuid string) (*iamAuthnModels.SamlIdentityProvider, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthnModels.GetSamlIdentityProviderApiResponse, iamAuthnModels.SamlIdentityProvider](
		func() (*iamAuthnModels.GetSamlIdentityProviderApiResponse, error) {
			return s.client.SAMLIdentityProvidersApiInstance.GetSamlIdentityProviderById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of SAML identity providers.
func (s *SAMLIdentityProvidersService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthnModels.SamlIdentityProvider, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing SAML Identity Providers")
	}

	return GenericListEntities[*iamAuthnModels.ListSamlIdentityProvidersApiResponse, iamAuthnModels.SamlIdentityProvider](
		func(reqParams *V4ODataParams) (*iamAuthnModels.ListSamlIdentityProvidersApiResponse, error) {
			return s.client.SAMLIdentityProvidersApiInstance.ListSamlIdentityProviders(
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

// NewIterator returns an iterator for listing SAML identity providers.
func (s *SAMLIdentityProvidersService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthnModels.SamlIdentityProvider] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthnModels.ListSamlIdentityProvidersApiResponse, iamAuthnModels.SamlIdentityProvider](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthnModels.ListSamlIdentityProvidersApiResponse, error) {
			return s.client.SAMLIdentityProvidersApiInstance.ListSamlIdentityProviders(
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

// Create creates a new SAML identity provider.
func (s *SAMLIdentityProvidersService) Create(ctx context.Context, samlIdp *iamAuthnModels.SamlIdentityProvider) (*iamAuthnModels.SamlIdentityProvider, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*iamAuthnModels.CreateSamlIdentityProviderApiResponse, iamAuthnModels.SamlIdentityProvider](
		s.client.SAMLIdentityProvidersApiInstance.CreateSamlIdentityProvider(samlIdp, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	return &result, nil
}

// Update updates an existing SAML identity provider.
func (s *SAMLIdentityProvidersService) Update(ctx context.Context, uuid string, samlIdp *iamAuthnModels.SamlIdentityProvider) (*iamAuthnModels.SamlIdentityProvider, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*iamAuthnModels.UpdateSamlIdentityProviderApiResponse, iamAuthnModels.SamlIdentityProvider](
		s.client.SAMLIdentityProvidersApiInstance.UpdateSamlIdentityProviderById(&uuid, samlIdp, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	return &result, nil
}

// Delete deletes an existing SAML identity provider.
func (s *SAMLIdentityProvidersService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	_, err := s.client.SAMLIdentityProvidersApiInstance.DeleteSamlIdentityProviderById(&uuid, nil)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %w", s.entitiesName, err)
	}
	return nil
}
