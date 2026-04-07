package converged

// ProtectionPolicies defines the interface for Data Policies Protection Policies API operations.
type ProtectionPolicies[ProtectionPolicy any] interface {
	Getter[ProtectionPolicy]
	Lister[ProtectionPolicy]
	Creator[ProtectionPolicy]
	Updater[ProtectionPolicy]
	Deleter[ProtectionPolicy]
	AsyncCreator[ProtectionPolicy]
	AsyncUpdater[ProtectionPolicy]
	AsyncDeleter[ProtectionPolicy]
}

// RecoveryPlans defines the interface for Data Policies Recovery Plans API operations.
type RecoveryPlans[RecoveryPlan any] interface {
	Getter[RecoveryPlan]
	Lister[RecoveryPlan]
	Creator[RecoveryPlan]
	Updater[RecoveryPlan]
	Deleter[RecoveryPlan]
	AsyncCreator[RecoveryPlan]
	AsyncUpdater[RecoveryPlan]
	AsyncDeleter[RecoveryPlan]
}

// DataPolicies groups the datapolicies namespace services (v4.2 API): protection policies and recovery plans.
type DataPolicies[ProtectionPolicy, RecoveryPlan any] struct {
	ProtectionPolicies ProtectionPolicies[ProtectionPolicy]
	RecoveryPlans      RecoveryPlans[RecoveryPlan]
}
