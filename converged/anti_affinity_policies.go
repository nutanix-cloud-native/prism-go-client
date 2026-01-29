package converged

// AntiAffinityPolicies defines the interface for anti-affinity policy operations.
// It supports get, list, create, update, delete (sync and async) for policies
// that control VM placement across hosts.
type AntiAffinityPolicies[AntiAffinityPolicy any] interface {
	Getter[AntiAffinityPolicy]
	Lister[AntiAffinityPolicy]
	Creator[AntiAffinityPolicy]
	Updater[AntiAffinityPolicy]
	Deleter[AntiAffinityPolicy]
	AsyncCreator[AntiAffinityPolicy]
	AsyncUpdater[AntiAffinityPolicy]
	AsyncDeleter[AntiAffinityPolicy]
}
