package converged

import "context"

type Projects[Project any] interface {
	// Getter is the interface for Get operations.
	Getter[Project]

	// Lister is the interface for List operations.
	Lister[Project]

	// Get default project.
	GetDefaultProject(ctx context.Context) (*Project, error)

	// GetByName resolves a project by name.
	GetByName(ctx context.Context, name string) (*Project, error)
}
