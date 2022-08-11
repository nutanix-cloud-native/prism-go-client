package secretdir

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/rand"
)

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
	g.Expect(os.WriteFile(filepath.Join(path, secretKeyName),
		[]byte(key),
		0o600)).To(Succeed())
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
	g.Expect(os.WriteFile(filepath.Join(path, secretKeyCertName),
		[]byte(cert),
		0o600)).To(Succeed())
	g.Expect(os.WriteFile(filepath.Join(path, secretKeyCertEndpoint),
		[]byte(fmt.Sprintf("%s:9440", ip)),
		0o600)).To(Succeed())

	prov := &provider{}
	me, err := prov.GetManagementEndpoint(nil)
	g.Expect(err).To(Succeed())
	g.Expect(me.Address.String()).To(BeIdenticalTo(fmt.Sprintf("https://%s:9440", ip)))
	g.Expect(me.ApiCredentials.Username).To(BeEmpty())
	g.Expect(me.ApiCredentials.Password).To(BeEmpty())
	g.Expect(me.ApiCredentials.KeyPair).To(BeIdenticalTo(cert))
}
