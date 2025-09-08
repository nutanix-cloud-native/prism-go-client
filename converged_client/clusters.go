package convergedclient

type Clusters[Cluster any] interface {
	// Getter is the interface for Get operations.
	Getter[Cluster]

	// Lister is the interface for List operations.
	Lister[Cluster]
}
