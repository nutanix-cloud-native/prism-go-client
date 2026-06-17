package converged

// Alerts defines the interface for managing alerts.
type Alerts[Alert any] interface {
	// Getter is the interface for Get operations.
	Getter[Alert]

	// Lister is the interface for List operations.
	Lister[Alert]
}
