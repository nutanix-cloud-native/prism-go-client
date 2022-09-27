package kubernetes

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"

	"github.com/nutanix-cloud-native/prism-go-client/environment/credentials"
	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
)

func runSecretInformer(ctx context.Context, clientset kubernetes.Interface) coreinformers.SecretInformer {
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)
	secretInformer := informerFactory.Core().V1().Secrets()
	informer := secretInformer.Informer()
	By("spawning secret informer", func() {
		go informer.Run(ctx.Done())
	})
	By("waiting for secret informer cache to be in sync", func() {
		Expect(
			cache.WaitForCacheSync(ctx.Done(), informer.HasSynced),
		).To(BeTrue())
	})
	return secretInformer
}

var _ = Describe("Kubernetes Environment Provider", Ordered, func() {
	var (
		secretName      = "nutanix-credentials"
		secretNamespace = "kube-system"
		secretInformer  coreinformers.SecretInformer
		prov            types.Provider
		ip              = rand.String(10)
		username        = rand.String(10)
		password        = rand.String(10)
		fakeClientset   = fake.NewSimpleClientset(&corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      secretName,
				Namespace: secretNamespace,
			},
			Data: map[string][]byte{
				"credentials": []byte(fmt.Sprintf(`
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
			`, username, password)),
			},
		})
		prismEndpoint = credentials.NutanixPrismEndpoint{
			Address:  ip,
			Port:     9440,
			Insecure: true,
			CredentialRef: &credentials.NutanixCredentialReference{
				Kind:      credentials.SecretKind,
				Name:      secretName,
				Namespace: secretNamespace,
			},
		}
	)
	BeforeAll(func() {
		secretInformer = runSecretInformer(context.TODO(), fakeClientset)
		prov = NewProvider(prismEndpoint, secretInformer)
	})
	It("must be able to look up secret", func() {
		_, err := fakeClientset.CoreV1().Secrets(secretNamespace).Get(context.TODO(),
			secretName, metav1.GetOptions{})
		Expect(err).To(Succeed())
		_, err = secretInformer.Lister().Secrets(secretNamespace).Get(secretName)
		Expect(err).To(Succeed())
	})
	It("must get management endpoint", func() {
		me, err := prov.GetManagementEndpoint(nil)
		Expect(err).To(Succeed())
		addr, _ := url.Parse(fmt.Sprintf("https://%s:9440", ip))
		Expect(me).To(BeEquivalentTo(&types.ManagementEndpoint{
			ApiCredentials: types.ApiCredentials{
				Username: username,
				Password: password,
			},
			Insecure: true,
			Address:  addr,
		}))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kubernetes Environment Provider Suite")
}
