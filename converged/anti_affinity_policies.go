package converged

type AntiAffinityPolicies[AntiAffinityPolicy any] interface {
	// Getter is the interface for Get operations.
	Getter[AntiAffinityPolicy]

	// Lister is the interface for List operations.
	Lister[AntiAffinityPolicy]

	// Creator is the interface for Create operations.
	Creator[AntiAffinityPolicy]

	// Updater is the interface for Update operations.
	Updater[AntiAffinityPolicy]

	// Deleter is the interface for Delete operations.
	Deleter[AntiAffinityPolicy]

	// AsyncCreator is the interface for Async Create operations.
	AsyncCreator[AntiAffinityPolicy]

	// AsyncUpdater is the interface for Async Update operations.
	AsyncUpdater[AntiAffinityPolicy]

	// AsyncDeleter is the interface for Async Delete operations.
	AsyncDeleter[AntiAffinityPolicy]
}
