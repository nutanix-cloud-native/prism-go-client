package converged

// RoleMemberships defines the interface for IAM Role Membership API operations.
type RoleMemberships[RoleMembership, RoleMembershipProjection any] interface {
	// Getter is the interface for Get operations.
	Getter[RoleMembership]

	// Lister is the interface for List operations.
	Lister[RoleMembershipProjection]

	// Creator is the interface for Create operations.
	Creator[RoleMembership]

	// Deleter is the interface for Delete operations.
	Deleter[RoleMembership]
}
