package converged

import "context"

// Templates defines the interface for Prism Central templates.
type Templates[Template, VM, TemplateDeployParams any] interface {
	// Getter is the interface for Get operations.
	Getter[Template]

	// Lister is the interface for List operations.
	Lister[Template]

	// Creator is the interface for Create operations.
	Creator[Template]

	// DeployVmWithTemplate deploys VMs from a template version.
	DeployVmWithTemplate(ctx context.Context, templateUUID string, params *TemplateDeployParams) (Operation[VM], error)
}
