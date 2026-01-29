package converged

// Images defines the interface for Prism Central image operations.
// Images are disk or ISO artifacts used for VM provisioning.
type Images[Image any] interface {
	Getter[Image]
	Lister[Image]
	Creator[Image]
	Updater[Image]
	Deleter[Image]
	AsyncCreator[Image]
	AsyncUpdater[Image]
	AsyncDeleter[Image]
}
