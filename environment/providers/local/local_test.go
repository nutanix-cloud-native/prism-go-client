package local

import (
	"fmt"
	"os"
	"testing"

	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestLocal(t *testing.T) {
	g := NewWithT(t)
	ip := rand.String(10)
	username := rand.String(10)
	password := rand.String(10)
	prov := &provider{}
	val, err := prov.Get(nil, "Foo")
	g.Expect(val).To(BeNil())
	g.Expect(err).To(BeIdenticalTo(types.ErrNotFound))

	os.Setenv(userEnv, username)
	os.Setenv(passwordEnv, password)
	os.Setenv(categoriesEnv, "k8s/foo,project/bar")

	endpoint := fmt.Sprintf("https://%s:9440", ip)
	os.Setenv(endpointEnv, endpoint)
	me, err := prov.GetManagementEndpoint(nil)
	g.Expect(err).To(Succeed())
	g.Expect(me.Address.String()).To(BeIdenticalTo(endpoint))

	os.Setenv(endpointEnv, ":not-valid")
	_, err = prov.GetManagementEndpoint(nil)
	g.Expect(err).NotTo(Succeed())

	os.Setenv(endpointEnv, endpoint)
	me, err = prov.GetManagementEndpoint(nil)
	g.Expect(err).To(Succeed())
	g.Expect(me.Address.String()).To(BeIdenticalTo(endpoint))
	g.Expect(me.ApiCredentials).To(BeIdenticalTo(types.ApiCredentials{
		Username: username,
		Password: password,
	}))
	val, err = prov.Get(nil, types.CategoriesKey)
	g.Expect(err).To(BeNil())
	g.Expect(val).To(BeEquivalentTo([]string{"k8s/foo", "project/bar"}))
}
