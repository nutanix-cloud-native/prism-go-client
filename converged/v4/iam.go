package v4

import (
	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	iamAuthnModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	// iamAuthzModels "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
)

// IAMService provides implementation for all IAM interface methods.
// NOTE: Currently only Users API is available in iam-go-client v4.0.1.
// Additional services (Roles, AccessPolicies, Operations, ServiceAccounts, APIKeys)
// are defined in the converged/iam.go interfaces and will be implemented here
// when they become available in future SDK versions.
type IAMService struct {
	client *v4prismGoClient.Client
	users  *UsersService
	// TODO: Uncomment when available in iam-go-client SDK
	// roles           *RolesService
	// accessPolicies  *AccessPoliciesService
	// operations      *OperationsService
	// serviceAccounts *ServiceAccountsService
	// apiKeys         *APIKeysService
}

// NewIAMService creates a new IAMService instance.
func NewIAMService(client *v4prismGoClient.Client) *IAMService {
	return &IAMService{
		client: client,
		users:  NewUsersService(client),
		// TODO: Uncomment when available in iam-go-client SDK
		// roles:           NewRolesService(client),
		// accessPolicies:  NewAccessPoliciesService(client),
		// operations:      NewOperationsService(client),
		// serviceAccounts: NewServiceAccountsService(client),
		// apiKeys:         NewAPIKeysService(client),
	}
}

// Users returns the Users service.
func (s *IAMService) Users() converged.Users[iamAuthnModels.User] {
	return s.users
}

// ========================================
// The following service implementations are commented out until the corresponding
// API instances become available in the iam-go-client SDK.
//
// To enable them:
// 1. Uncomment the import for iamAuthzModels above
// 2. Add the API instances to v4/v4.go Client struct
// 3. Initialize the API instances in v4/v4.go initIAMApiInstance function
// 4. Uncomment the service implementations below
// ========================================

/*
// Roles returns the Roles service.
func (s *IAMService) Roles() converged.Roles[iamAuthzModels.Role] {
	return s.roles
}

// AccessPolicies returns the AccessPolicies service.
func (s *IAMService) AccessPolicies() converged.AccessPolicies[iamAuthzModels.AccessPolicy] {
	return s.accessPolicies
}

// Operations returns the Operations service.
func (s *IAMService) Operations() converged.Operations[iamAuthzModels.Operation] {
	return s.operations
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

	return GenericCreateEntity[*iamAuthzModels.CreateRoleApiResponse, iamAuthzModels.Role](
		func() (*iamAuthzModels.CreateRoleApiResponse, error) {
			return s.client.RolesApiInstance.CreateRole(role, nil)
		},
		s.entitiesName,
	)
}

// Update updates an existing role.
func (s *RolesService) Update(ctx context.Context, uuid string, role *iamAuthzModels.Role) (*iamAuthzModels.Role, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericUpdateEntity[*iamAuthzModels.UpdateRoleApiResponse, iamAuthzModels.Role](
		func() (*iamAuthzModels.UpdateRoleApiResponse, error) {
			return s.client.RolesApiInstance.UpdateRoleById(&uuid, role, nil)
		},
		s.entitiesName,
	)
}

// Delete deletes an existing role.
func (s *RolesService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	return GenericDeleteEntity(
		func() (*iamAuthzModels.DeleteRoleApiResponse, error) {
			return s.client.RolesApiInstance.DeleteRoleById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// AccessPoliciesService provides implementation for IAM Access Policies API operations.
type AccessPoliciesService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewAccessPoliciesService creates a new AccessPoliciesService instance.
func NewAccessPoliciesService(client *v4prismGoClient.Client) *AccessPoliciesService {
	return &AccessPoliciesService{
		client:       client,
		entitiesName: "access policy",
	}
}

// Get returns the access policy for the given UUID.
func (s *AccessPoliciesService) Get(ctx context.Context, uuid string) (*iamAuthzModels.AccessPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthzModels.GetAccessPolicyApiResponse, iamAuthzModels.AccessPolicy](
		func() (*iamAuthzModels.GetAccessPolicyApiResponse, error) {
			return s.client.AccessPoliciesApiInstance.GetAccessPolicyById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of access policies.
func (s *AccessPoliciesService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthzModels.AccessPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing Access Policies")
	}

	return GenericListEntities[*iamAuthzModels.ListAccessPoliciesApiResponse, iamAuthzModels.AccessPolicy](
		func(reqParams *V4ODataParams) (*iamAuthzModels.ListAccessPoliciesApiResponse, error) {
			return s.client.AccessPoliciesApiInstance.ListAccessPolicies(
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

// NewIterator returns an iterator for listing access policies.
func (s *AccessPoliciesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthzModels.AccessPolicy] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthzModels.ListAccessPoliciesApiResponse, iamAuthzModels.AccessPolicy](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthzModels.ListAccessPoliciesApiResponse, error) {
			return s.client.AccessPoliciesApiInstance.ListAccessPolicies(
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

// Create creates a new access policy.
func (s *AccessPoliciesService) Create(ctx context.Context, policy *iamAuthzModels.AccessPolicy) (*iamAuthzModels.AccessPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericCreateEntity[*iamAuthzModels.CreateAccessPolicyApiResponse, iamAuthzModels.AccessPolicy](
		func() (*iamAuthzModels.CreateAccessPolicyApiResponse, error) {
			return s.client.AccessPoliciesApiInstance.CreateAccessPolicy(policy, nil)
		},
		s.entitiesName,
	)
}

// Update updates an existing access policy.
func (s *AccessPoliciesService) Update(ctx context.Context, uuid string, policy *iamAuthzModels.AccessPolicy) (*iamAuthzModels.AccessPolicy, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericUpdateEntity[*iamAuthzModels.UpdateAccessPolicyApiResponse, iamAuthzModels.AccessPolicy](
		func() (*iamAuthzModels.UpdateAccessPolicyApiResponse, error) {
			return s.client.AccessPoliciesApiInstance.UpdateAccessPolicyById(&uuid, policy, nil)
		},
		s.entitiesName,
	)
}

// Delete deletes an existing access policy.
func (s *AccessPoliciesService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	return GenericDeleteEntity(
		func() (*iamAuthzModels.DeleteAccessPolicyApiResponse, error) {
			return s.client.AccessPoliciesApiInstance.DeleteAccessPolicyById(&uuid, nil)
		},
		s.entitiesName,
	)
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

// ServiceAccountsService provides implementation for IAM Service Accounts API operations.
type ServiceAccountsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewServiceAccountsService creates a new ServiceAccountsService instance.
func NewServiceAccountsService(client *v4prismGoClient.Client) *ServiceAccountsService {
	return &ServiceAccountsService{
		client:       client,
		entitiesName: "service account",
	}
}

// Get returns the service account for the given UUID.
func (s *ServiceAccountsService) Get(ctx context.Context, uuid string) (*iamAuthnModels.ServiceAccount, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthnModels.GetServiceAccountApiResponse, iamAuthnModels.ServiceAccount](
		func() (*iamAuthnModels.GetServiceAccountApiResponse, error) {
			return s.client.ServiceAccountsApiInstance.GetServiceAccountById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of service accounts.
func (s *ServiceAccountsService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthnModels.ServiceAccount, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing Service Accounts")
	}

	return GenericListEntities[*iamAuthnModels.ListServiceAccountsApiResponse, iamAuthnModels.ServiceAccount](
		func(reqParams *V4ODataParams) (*iamAuthnModels.ListServiceAccountsApiResponse, error) {
			return s.client.ServiceAccountsApiInstance.ListServiceAccounts(
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

// NewIterator returns an iterator for listing service accounts.
func (s *ServiceAccountsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthnModels.ServiceAccount] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthnModels.ListServiceAccountsApiResponse, iamAuthnModels.ServiceAccount](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthnModels.ListServiceAccountsApiResponse, error) {
			return s.client.ServiceAccountsApiInstance.ListServiceAccounts(
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

// Create creates a new service account.
func (s *ServiceAccountsService) Create(ctx context.Context, serviceAccount *iamAuthnModels.ServiceAccount) (*iamAuthnModels.ServiceAccount, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericCreateEntity[*iamAuthnModels.CreateServiceAccountApiResponse, iamAuthnModels.ServiceAccount](
		func() (*iamAuthnModels.CreateServiceAccountApiResponse, error) {
			return s.client.ServiceAccountsApiInstance.CreateServiceAccount(serviceAccount, nil)
		},
		s.entitiesName,
	)
}

// Update updates an existing service account.
func (s *ServiceAccountsService) Update(ctx context.Context, uuid string, serviceAccount *iamAuthnModels.ServiceAccount) (*iamAuthnModels.ServiceAccount, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericUpdateEntity[*iamAuthnModels.UpdateServiceAccountApiResponse, iamAuthnModels.ServiceAccount](
		func() (*iamAuthnModels.UpdateServiceAccountApiResponse, error) {
			return s.client.ServiceAccountsApiInstance.UpdateServiceAccountById(&uuid, serviceAccount, nil)
		},
		s.entitiesName,
	)
}

// Delete deletes an existing service account.
func (s *ServiceAccountsService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	return GenericDeleteEntity(
		func() (*iamAuthnModels.DeleteServiceAccountApiResponse, error) {
			return s.client.ServiceAccountsApiInstance.DeleteServiceAccountById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// APIKeysService provides implementation for IAM API Keys operations.
type APIKeysService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewAPIKeysService creates a new APIKeysService instance.
func NewAPIKeysService(client *v4prismGoClient.Client) *APIKeysService {
	return &APIKeysService{
		client:       client,
		entitiesName: "api key",
	}
}

// Get returns the API key for the given UUID.
func (s *APIKeysService) Get(ctx context.Context, uuid string) (*iamAuthnModels.ApiKey, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*iamAuthnModels.GetApiKeyApiResponse, iamAuthnModels.ApiKey](
		func() (*iamAuthnModels.GetApiKeyApiResponse, error) {
			return s.client.ApiKeysApiInstance.GetApiKeyById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// List returns a list of API keys.
func (s *APIKeysService) List(ctx context.Context, opts ...converged.ODataOption) ([]iamAuthnModels.ApiKey, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams != nil && (myParams.Apply != nil || myParams.Expand != nil) {
		return nil, fmt.Errorf("apply and expand options are not supported for listing API Keys")
	}

	return GenericListEntities[*iamAuthnModels.ListApiKeysApiResponse, iamAuthnModels.ApiKey](
		func(reqParams *V4ODataParams) (*iamAuthnModels.ListApiKeysApiResponse, error) {
			return s.client.ApiKeysApiInstance.ListApiKeys(
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

// NewIterator returns an iterator for listing API keys.
func (s *APIKeysService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[iamAuthnModels.ApiKey] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*iamAuthnModels.ListApiKeysApiResponse, iamAuthnModels.ApiKey](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*iamAuthnModels.ListApiKeysApiResponse, error) {
			return s.client.ApiKeysApiInstance.ListApiKeys(
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

// Create creates a new API key.
func (s *APIKeysService) Create(ctx context.Context, apiKey *iamAuthnModels.ApiKey) (*iamAuthnModels.ApiKey, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericCreateEntity[*iamAuthnModels.CreateApiKeyApiResponse, iamAuthnModels.ApiKey](
		func() (*iamAuthnModels.CreateApiKeyApiResponse, error) {
			return s.client.ApiKeysApiInstance.CreateApiKey(apiKey, nil)
		},
		s.entitiesName,
	)
}

// Update updates an existing API key.
func (s *APIKeysService) Update(ctx context.Context, uuid string, apiKey *iamAuthnModels.ApiKey) (*iamAuthnModels.ApiKey, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericUpdateEntity[*iamAuthnModels.UpdateApiKeyApiResponse, iamAuthnModels.ApiKey](
		func() (*iamAuthnModels.UpdateApiKeyApiResponse, error) {
			return s.client.ApiKeysApiInstance.UpdateApiKeyById(&uuid, apiKey, nil)
		},
		s.entitiesName,
	)
}

// Delete deletes an existing API key.
func (s *APIKeysService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	return GenericDeleteEntity(
		func() (*iamAuthnModels.DeleteApiKeyApiResponse, error) {
			return s.client.ApiKeysApiInstance.DeleteApiKeyById(&uuid, nil)
		},
		s.entitiesName,
	)
}

// RevokeAPIKey revokes an API key.
func (s *APIKeysService) RevokeAPIKey(ctx context.Context, apiKeyExtId string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	_, err := s.client.ApiKeysApiInstance.RevokeApiKey(&apiKeyExtId, nil)
	if err != nil {
		return fmt.Errorf("failed to revoke %s: %w", s.entitiesName, err)
	}

	return nil
}
*/
