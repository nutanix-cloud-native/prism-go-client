package converged

// AuthorizationPolicies defines the interface for IAM Authorization Policies API operations.
type AuthorizationPolicies[AuthorizationPolicy, AuthorizationPolicyProjection any] interface {
	// Getter is the interface for Get operations.
	Getter[AuthorizationPolicy]

	// Lister is the interface for List operations.
	Lister[AuthorizationPolicyProjection]

	// Creator is the interface for Create operations.
	Creator[AuthorizationPolicy]

	// Deleter is the interface for Delete operations.
	Deleter[AuthorizationPolicy]
}
