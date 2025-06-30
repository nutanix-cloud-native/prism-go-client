package v4

import (
	"slices"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
)

func (f *FacadeV4Client) ListCapabilities() []facade.Capability {
	return []facade.Capability{
		facade.CapabilityAntiAffinityPolicy,
		facade.CapabilityCategories,
	}
}

func (f *FacadeV4Client) HasCapability(capability facade.Capability) bool {
	capabilities := f.ListCapabilities()
	return slices.Contains(capabilities, capability)
}
