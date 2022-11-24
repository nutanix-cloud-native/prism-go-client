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
	krand "k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"

	"github.com/nutanix-cloud-native/prism-go-client/environment/credentials"
	"github.com/nutanix-cloud-native/prism-go-client/environment/types"
	"github.com/nutanix-cloud-native/prism-go-client/internal/certutils"
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

func runCMInformer(ctx context.Context, clientset kubernetes.Interface) coreinformers.ConfigMapInformer {
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)
	cmInformer := informerFactory.Core().V1().ConfigMaps()
	informer := cmInformer.Informer()
	By("spawning cm informer", func() {
		go informer.Run(ctx.Done())
	})
	By("waiting for cm informer cache to be in sync", func() {
		Expect(
			cache.WaitForCacheSync(ctx.Done(), informer.HasSynced),
		).To(BeTrue())
	})
	return cmInformer
}

var _ = Describe("Kubernetes Environment Provider", Ordered, func() {
	expectedCACert, expectedB64CACert, err := certutils.GenerateCACertForTesting()
	Expect(err).ToNot(HaveOccurred())

	var (
		secretName      = "nutanix-credentials"
		cmName          = "user-ca-bundle"
		secretNamespace = "kube-system"
		secretInformer  coreinformers.SecretInformer
		cmInformer      coreinformers.ConfigMapInformer
		prov            types.Provider
		ip              = krand.String(10)
		username        = krand.String(10)
		password        = krand.String(10)

		fakeSecret = &corev1.Secret{
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
		}
		fakeCM = &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      cmName,
				Namespace: secretNamespace,
			},
			BinaryData: map[string][]byte{
				certBundleKey: []byte(expectedB64CACert),
			},
		}

		fakeClientset = fake.NewSimpleClientset(fakeSecret, fakeCM)
		prismEndpoint = credentials.NutanixPrismEndpoint{
			Address:  ip,
			Port:     9440,
			Insecure: true,
			CredentialRef: &credentials.NutanixCredentialReference{
				Kind:      credentials.SecretKind,
				Name:      secretName,
				Namespace: secretNamespace,
			},
			AdditionalTrustBundle: &credentials.NutanixTrustBundleReference{
				Kind:      credentials.NutanixTrustBundleKindConfigMap,
				Name:      cmName,
				Namespace: secretNamespace,
			},
		}
	)
	BeforeAll(func() {
		secretInformer = runSecretInformer(context.TODO(), fakeClientset)
		cmInformer = runCMInformer(context.TODO(), fakeClientset)
		prov = NewProvider(prismEndpoint, secretInformer, cmInformer)
	})
	It("must be able to look up secret", func() {
		_, err := fakeClientset.CoreV1().Secrets(secretNamespace).Get(context.TODO(), secretName, metav1.GetOptions{})
		Expect(err).To(Succeed())

		_, err = secretInformer.Lister().Secrets(secretNamespace).Get(secretName)
		Expect(err).To(Succeed())

		_, err = fakeClientset.CoreV1().ConfigMaps(secretNamespace).Get(context.TODO(), cmName, metav1.GetOptions{})
		Expect(err).To(Succeed())

		_, err = cmInformer.Lister().ConfigMaps(secretNamespace).Get(cmName)
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
			Insecure:              true,
			Address:               addr,
			AdditionalTrustBundle: expectedCACert,
		}))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kubernetes Environment Provider Suite")
}
