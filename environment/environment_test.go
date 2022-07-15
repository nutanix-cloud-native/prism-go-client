package environment

import (
	"net/url"
	"testing"

	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	. "github.com/onsi/gomega"
)

type testProvider struct {
	// single management endpoint
	mgmtEndpoint *types.ManagementEndpoint
	// variables as JSON doc
	vars map[string]interface{}
}

func (prov *testProvider) GetManagementEndpoint(
	topology types.Topology,
) (*types.ManagementEndpoint, error) {
	return prov.mgmtEndpoint, nil
}

func (prov *testProvider) Get(topology types.Topology, key string) (
	interface{}, error,
) {
	if val, ok := prov.vars[key]; ok {
		return val, nil
	}
	return nil, types.ErrNotFound
}

func TestEnvironment(t *testing.T) {
	g := NewWithT(t)
	me := &types.ManagementEndpoint{
		Address: &url.URL{Scheme: "https", Host: "pc"},
		ApiCredentials: types.ApiCredentials{
			Username: "k8s",
			Password: "k8s/4u",
		},
	}
	env := NewEnvironment(
		&testProvider{
			mgmtEndpoint: me,
			vars: map[string]interface{}{
				"foo": true,
			},
		},
		&testProvider{
			vars: map[string]interface{}{
				"bar": false,
			},
		},
	)

	val, err := env.Get(nil, "bar")
	g.Expect(err).To(BeNil())
	g.Expect(val).To(BeFalse())
	val, err = env.Get(nil, "foo")
	g.Expect(err).To(BeNil())
	g.Expect(val).To(BeTrue())
	_, err = env.Get(nil, "none")
	g.Expect(err).To(BeIdenticalTo(types.ErrNotFound))

	g.Expect(env.GetManagementEndpoint(nil)).To(BeIdenticalTo(me))
}
