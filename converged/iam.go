package converged

// IAM defines the interface for Identity and Access Management operations.
// It provides access to all IAM-related services including Users, Roles, AuthorizationPolicies, and Operations.
type IAM[User, Role, AuthorizationPolicy, Operation, UserGroup, DirectoryService, SAMLIdentityProvider any] interface {
	// Users provides access to IAM User operations.
	Users() Users[User]

	// Roles provides access to IAM Role operations.
	Roles() Roles[Role]

	// AuthorizationPolicies provides access to IAM Authorization Policy operations.
	AuthorizationPolicies() AuthorizationPolicies[AuthorizationPolicy]

	// Operations provides access to IAM Operation operations.
	Operations() Operations[Operation]

	// UserGroups provides access to IAM User Group operations.
	UserGroups() UserGroups[UserGroup]

	// DirectoryServices provides access to IAM Directory Service operations.
	DirectoryServices() DirectoryServices[DirectoryService]

	// SAMLIdentityProviders provides access to SAML Identity Provider operations.
	SAMLIdentityProviders() SAML[SAMLIdentityProvider]
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

// AuthorizationPolicies defines the interface for IAM Authorization Policies API operations.
type AuthorizationPolicies[AuthorizationPolicy any] interface {
	// Getter is the interface for Get operations.
	Getter[AuthorizationPolicy]

	// Lister is the interface for List operations.
	Lister[AuthorizationPolicy]

	// Creator is the interface for Create operations.
	Creator[AuthorizationPolicy]

	// Updater is the interface for Update operations.
	Updater[AuthorizationPolicy]

	// Deleter is the interface for Delete operations.
	Deleter[AuthorizationPolicy]
}

// Operations defines the interface for IAM Operations API operations.
// Operations represent the available actions/permissions in the system.
type Operations[Operation any] interface {
	// Getter is the interface for Get operations.
	Getter[Operation]

	// Lister is the interface for List operations.
	Lister[Operation]
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
