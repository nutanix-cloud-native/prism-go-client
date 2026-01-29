package converged

// Users defines the interface for IAM user operations.
// Users represent identities in Prism Central for authentication and authorization.
type Users[User any] interface {
	Getter[User]
	Lister[User]
}
