package converged

// Categories defines the interface for category operations.
// Categories are used to tag and organize entities (VMs, hosts, etc.) in Prism Central.
type Categories[Category any] interface {
	Getter[Category]
	Lister[Category]
	Creator[Category]
	Updater[Category]
	Deleter[Category]
}
