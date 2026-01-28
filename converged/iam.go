package converged

import "context"

// IAM defines the interface for Identity and Access Management operations.
// It provides access to all IAM-related services including Users, Roles, AccessPolicies, and Operations.
type IAM[User, Role, AccessPolicy, Operation any] interface {
	// Users provides access to IAM User operations.
	Users() Users[User]

	// Roles provides access to IAM Role operations.
	Roles() Roles[Role]

	// AccessPolicies provides access to IAM Access Policy operations.
	AccessPolicies() AccessPolicies[AccessPolicy]

	// Operations provides access to IAM Operation operations.
	Operations() Operations[Operation]
}

// Roles defines the interface for IAM Roles API operations.
type Roles[Role any] interface {
	// Getter is the interface for Get operations.
	Getter[Role]

	// Lister is the interface for List operations.
	Lister[Role]

	// Creator is the interface for Create operations.
	Creator[Role]

	// Updater is the interface for Update operations.
	Updater[Role]

	// Deleter is the interface for Delete operations.
	Deleter[Role]
}

// AccessPolicies defines the interface for IAM Access Policies API operations.
type AccessPolicies[AccessPolicy any] interface {
	// Getter is the interface for Get operations.
	Getter[AccessPolicy]

	// Lister is the interface for List operations.
	Lister[AccessPolicy]

	// Creator is the interface for Create operations.
	Creator[AccessPolicy]

	// Updater is the interface for Update operations.
	Updater[AccessPolicy]

	// Deleter is the interface for Delete operations.
	Deleter[AccessPolicy]
}

// Operations defines the interface for IAM Operations API operations.
// Operations represent the available actions/permissions in the system.
type Operations[Operation any] interface {
	// Getter is the interface for Get operations.
	Getter[Operation]

	// Lister is the interface for List operations.
	Lister[Operation]
}

// ServiceAccounts defines the interface for IAM Service Accounts API operations.
type ServiceAccounts[ServiceAccount any] interface {
	// Getter is the interface for Get operations.
	Getter[ServiceAccount]

	// Lister is the interface for List operations.
	Lister[ServiceAccount]

	// Creator is the interface for Create operations.
	Creator[ServiceAccount]

	// Updater is the interface for Update operations.
	Updater[ServiceAccount]

	// Deleter is the interface for Delete operations.
	Deleter[ServiceAccount]
}

// APIKeys defines the interface for IAM API Keys operations.
type APIKeys[APIKey any] interface {
	// Getter is the interface for Get operations.
	Getter[APIKey]

	// Lister is the interface for List operations.
	Lister[APIKey]

	// Creator is the interface for Create operations.
	Creator[APIKey]

	// Updater is the interface for Update operations.
	Updater[APIKey]

	// Deleter is the interface for Delete operations.
	Deleter[APIKey]

	// RevokeAPIKey revokes an API key.
	RevokeAPIKey(ctx context.Context, apiKeyExtId string) error
}

// AuthenticationPolicies defines the interface for IAM Authentication Policies API operations.
type AuthenticationPolicies[AuthenticationPolicy any] interface {
	// Getter is the interface for Get operations.
	Getter[AuthenticationPolicy]

	// Updater is the interface for Update operations.
	Updater[AuthenticationPolicy]
}

// DirectoryServices defines the interface for IAM Directory Services API operations.
type DirectoryServices[DirectoryService any] interface {
	// Getter is the interface for Get operations.
	Getter[DirectoryService]

	// Lister is the interface for List operations.
	Lister[DirectoryService]

	// Creator is the interface for Create operations.
	Creator[DirectoryService]

	// Updater is the interface for Update operations.
	Updater[DirectoryService]

	// Deleter is the interface for Delete operations.
	Deleter[DirectoryService]
}

// UserGroups defines the interface for IAM User Groups API operations.
type UserGroups[UserGroup any] interface {
	// Getter is the interface for Get operations.
	Getter[UserGroup]

	// Lister is the interface for List operations.
	Lister[UserGroup]

	// Creator is the interface for Create operations.
	Creator[UserGroup]

	// Updater is the interface for Update operations.
	Updater[UserGroup]

	// Deleter is the interface for Delete operations.
	Deleter[UserGroup]
}

// SAML defines the interface for IAM SAML Identity Provider operations.
type SAML[SAMLIdentityProvider any] interface {
	// Getter is the interface for Get operations.
	Getter[SAMLIdentityProvider]

	// Lister is the interface for List operations.
	Lister[SAMLIdentityProvider]

	// Creator is the interface for Create operations.
	Creator[SAMLIdentityProvider]

	// Updater is the interface for Update operations.
	Updater[SAMLIdentityProvider]

	// Deleter is the interface for Delete operations.
	Deleter[SAMLIdentityProvider]
}

// Certificates defines the interface for IAM Certificate operations.
type Certificates[Certificate any] interface {
	// Getter is the interface for Get operations.
	Getter[Certificate]

	// Lister is the interface for List operations.
	Lister[Certificate]

	// Creator is the interface for Create operations.
	Creator[Certificate]

	// Updater is the interface for Update operations.
	Updater[Certificate]

	// Deleter is the interface for Delete operations.
	Deleter[Certificate]
}
