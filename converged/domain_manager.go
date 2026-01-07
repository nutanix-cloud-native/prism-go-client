package converged

import "context"

// DomainManager defines the interface for Domain Manager operations.
type DomainManager[DomainManagerEntity any] interface {
	// Getter is the interface for Get operations.
	Getter[DomainManagerEntity]

	// Lister is the interface for List operations.
	Lister[DomainManagerEntity]

	// GetPrismCentralVersion gets the Prism Central version from Domain Manager API.
	// This is a lightweight alternative to V3 GetPrismCentral() API.
	GetPrismCentralVersion(ctx context.Context) (string, error)
}
