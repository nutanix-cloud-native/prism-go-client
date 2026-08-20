package converged

import "context"

type VMProfiles[VMProfile, VM, DeployParams any] interface {
	// Getter is the interface for Get operations.
	Getter[VMProfile]

	// Lister is the interface for List operations.
	Lister[VMProfile]

	// DeployVmWithVmProfile deploys a VM from a VM Profile.
	DeployVmWithVmProfile(ctx context.Context, vmProfileUUID string, params *DeployParams) (Operation[VM], error)
}
