package secretdir

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/rand"
)

// Write regular and symlink file for given key-val pair.
func writeParam(t *testing.T, path, key, val string) error {
	filename := rand.String(10)
	keyPath := filepath.Join(path, key)
	dataPath := filepath.Join(path, filename)

	if err := os.WriteFile(dataPath, []byte(val), 0o600); err != nil {
		return err
	}
	if err := os.Symlink(filename, keyPath); err != nil {
		return err
	}
	return nil
}

func TestBasicAuth(t *testing.T) {
	g := NewWithT(t)
	path, err := os.MkdirTemp(".", "basic-auth")
	g.Expect(err).To(Succeed())
	defer func() {
		if !t.Failed() {
			os.RemoveAll(path)
		}
	}()
	t.Logf("Temporary directory %s", path)
	os.Setenv(envSecretDir, path)
	ip := rand.String(10)
	user := rand.String(10)
	passwd := rand.String(10)
	key := fmt.Sprintf("%s:9440:%s:%s", ip, user, passwd)
	g.Expect(writeParam(t, path, secretKeyName, key)).To(Succeed())
	t.Log("Key", key)

	prov := &provider{}
	me, err := prov.GetManagementEndpoint(nil)
	g.Expect(err).To(Succeed())
	g.Expect(me.Address.String()).To(BeIdenticalTo(fmt.Sprintf("https://%s:9440", ip)))
	g.Expect(me.ApiCredentials.Username).To(BeIdenticalTo(user))
	g.Expect(me.ApiCredentials.Password).To(BeIdenticalTo(passwd))
	g.Expect(me.ApiCredentials.KeyPair).To(BeEmpty())
}

func TestTLSAuth(t *testing.T) {
	g := NewWithT(t)
	path, err := os.MkdirTemp(".", "tls-auth")
	g.Expect(err).To(Succeed())
	defer func() {
		if !t.Failed() {
			os.RemoveAll(path)
		}
	}()
	ip := rand.String(10)
	cert := rand.String(512)
	t.Logf("Temporary directory %s", path)
	os.Setenv(envSecretDir, path)
	g.Expect(writeParam(t, path, secretKeyCertName, cert)).To(Succeed())
	g.Expect(writeParam(t, path, secretKeyEndpoint,
		fmt.Sprintf("%s:9440", ip))).To(Succeed())

	prov := &provider{}
	me, err := prov.GetManagementEndpoint(nil)
	g.Expect(err).To(Succeed())
	g.Expect(me.Address.String()).To(BeIdenticalTo(fmt.Sprintf("https://%s:9440", ip)))
	g.Expect(me.ApiCredentials.Username).To(BeEmpty())
	g.Expect(me.ApiCredentials.Password).To(BeEmpty())
	g.Expect(me.ApiCredentials.KeyPair).To(BeIdenticalTo(cert))
}

func TestNutanixAuth(t *testing.T) {
	g := NewWithT(t)
	path, err := os.MkdirTemp(".", "nutanix-auth")
	g.Expect(err).To(Succeed())
	defer func() {
		if !t.Failed() {
			os.RemoveAll(path)
		}
	}()
	ip := rand.String(10)
	user := rand.String(10)
	passwd := rand.String(10)
	creds := fmt.Sprintf(`
		[
			{
			  "type": "basic_auth", 
			  "data": { 
				"prismCentral":{
				  "username": "%s", 
				  "password": "%s"
				},
				"prismElements": null
			  }
			}
		]
	`, user, passwd)
	endpoint := fmt.Sprintf("%s:9440", ip)
	t.Logf("Temporary directory %s", path)
	os.Setenv(envSecretDir, path)

	// Write three keys
	g.Expect(writeParam(t, path, secretKeyCredentials, creds)).To(Succeed())
	g.Expect(writeParam(t, path, secretKeyInsecure, "true")).To(Succeed())
	g.Expect(writeParam(t, path, secretKeyEndpoint, endpoint)).To(Succeed())

	prov := &provider{}
	me, err := prov.GetManagementEndpoint(nil)
	g.Expect(err).To(Succeed())
	addr, err := url.Parse(fmt.Sprintf("https://%s:9440", ip))
	g.Expect(err).To(Succeed())
	g.Expect(me).To(BeEquivalentTo(&types.ManagementEndpoint{
		ApiCredentials: types.ApiCredentials{
			Username: user,
			Password: passwd,
		},
		Insecure: true,
		Address:  addr,
	}))
}
