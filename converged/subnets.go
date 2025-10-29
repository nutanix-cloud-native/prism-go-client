package converged

type Subnets[Subnet, SubnetListParams any] interface {
	// Getter is the interface for Get operations.
	Getter[Subnet]

	// Lister is the interface for List operations.
	Lister[Subnet, SubnetListParams]
}
