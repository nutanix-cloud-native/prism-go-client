package facade

type Capability string

const (
	CapabilityAntiAffinityPolicy Capability = "anti_affinity_policy"
	CapabilityCategories         Capability = "categories"
)

type FacadeCapabilities interface {
	HasCapability(capability Capability) bool
	ListCapabilities() []Capability
}
