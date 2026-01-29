package converged

import "context"

// DomainManager defines the interface for Domain Manager operations.
// It provides access to domain manager entities and Prism Central version.
type DomainManager[DomainManagerEntity any] interface {
	Getter[DomainManagerEntity]
	Lister[DomainManagerEntity]

	// GetPrismCentralVersion returns the Prism Central version from the Domain Manager API.
	// It is a lightweight alternative to the V3 GetPrismCentral() API.
	GetPrismCentralVersion(ctx context.Context) (string, error)
}
