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
	var (
		secretName      = "nutanix-credentials"
		cmName          = "user-ca-bundle"
		secretNamespace = "kube-system"
		secretInformer  coreinformers.SecretInformer
		cmInformer      coreinformers.ConfigMapInformer
		prov            types.Provider
		ip              = rand.String(10)
		username        = rand.String(10)
		password        = rand.String(10)

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
				certBundleKey: []byte("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNyRENDQVpRQ0NRQ05seHR1azgvekp6QU5CZ2txaGtpRzl3MEJBUXNGQURBWE1SVXdFd1lEVlFRRERBeGgKWkcxcGMzTnBiMjVmWTJFd0lCY05Nakl4TVRFNE1UVXdNekV4V2hnUE1qSTVOakE1TURJeE5UQXpNVEZhTUJjeApGVEFUQmdOVkJBTU1ER0ZrYldsemMybHZibDlqWVRDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDCkFRb0NnZ0VCQUtyTTZ4VndlbGppN29hNThoc0w3UERZS0E4cVJ4THFYc2JBQ3pIdXZHcUVUQktmU0gxbFVRTlkKYVNRUlQzUTlWOXhsYzBuS3c1TGN3L1JXZ3ZMN0JPaEZtY0gyQWMrT2ZLQUFoRTVLTnBVeTgwV3pzaXh5ZFI2WApoVEVTbVVETHBsamFQREdLZWxCY0ZvT1BNYzRDeFNJUExMZ0p3dTNIT1FiYit6VS9pcUs4VVpKRDdwMDdycFN6CmdKMkRvTkVWSnBEU1h2N3ZwMERpZ1duRDN5dmw5RmxaVzJLOFZjZXp4ME5LRi9pbkplNFJ2cmQxcy9PSEhYQzQKczBmM2NWOVljZmYzL1BoV21RQ1NocmVwQm5WeUVMcXRpWTg5TkRqYVhlVVpsR1JwTTQ3Z21RQktaSCtnMXlENgorSlJTOGp6ZDlZS3VDZFF3NGR3REtPMjlyZ29OM1VrQ0F3RUFBVEFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBCk9CSXJkT0M2VWNCMFhOQ0U1VGx3NlJzM21EL1pGdmJmdk1YWmNCaGJGTlF3NHN0bmNyTGk2WEpKdVBVNVZHMDcKV05wQklUMzEyV1hLbHkxbGxmeWh4dGZRNGJ4SnI3SGUrLzRMeGU1bEVFTCtjK1BSQzVxMlhUby91VDExdHJjVQpiMnl1aDAvRzNZSnFqOUoyTDM4WEZ1Z0tDNGwrQUtoeTJsa1hLVUZaeVp0eGJMWWdLazFkb1RKeE1RVGl4dWxuCjNVOS9MdTlEMDRpQW9LS1RYMDgrZG51MXpTd0Rhd251UHoyN0lWOGdLNW5GN1VNak1iNFB2cEZvaTRDYmNJOTYKd2dIMy9KL0hDTWJRVG5rMVcxMEQwMHJFcU1UNUVVTCtjZlVwbG5vT1RYV2YvK3NaaXZvSU1vTmhSdFB0czB0SAptazdaUGs0bzd1RGZGNURXSkpUZkx3PT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="),
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
			AdditionalTrustBundle: "-----BEGIN CERTIFICATE-----\nMIICrDCCAZQCCQCNlxtuk8/zJzANBgkqhkiG9w0BAQsFADAXMRUwEwYDVQQDDAxh\nZG1pc3Npb25fY2EwIBcNMjIxMTE4MTUwMzExWhgPMjI5NjA5MDIxNTAzMTFaMBcx\nFTATBgNVBAMMDGFkbWlzc2lvbl9jYTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC\nAQoCggEBAKrM6xVwelji7oa58hsL7PDYKA8qRxLqXsbACzHuvGqETBKfSH1lUQNY\naSQRT3Q9V9xlc0nKw5Lcw/RWgvL7BOhFmcH2Ac+OfKAAhE5KNpUy80WzsixydR6X\nhTESmUDLpljaPDGKelBcFoOPMc4CxSIPLLgJwu3HOQbb+zU/iqK8UZJD7p07rpSz\ngJ2DoNEVJpDSXv7vp0DigWnD3yvl9FlZW2K8Vcezx0NKF/inJe4Rvrd1s/OHHXC4\ns0f3cV9Ycff3/PhWmQCShrepBnVyELqtiY89NDjaXeUZlGRpM47gmQBKZH+g1yD6\n+JRS8jzd9YKuCdQw4dwDKO29rgoN3UkCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEA\nOBIrdOC6UcB0XNCE5Tlw6Rs3mD/ZFvbfvMXZcBhbFNQw4stncrLi6XJJuPU5VG07\nWNpBIT312WXKly1llfyhxtfQ4bxJr7He+/4Lxe5lEEL+c+PRC5q2XTo/uT11trcU\nb2yuh0/G3YJqj9J2L38XFugKC4l+AKhy2lkXKUFZyZtxbLYgKk1doTJxMQTixuln\n3U9/Lu9D04iAoKKTX08+dnu1zSwDawnuPz27IV8gK5nF7UMjMb4PvpFoi4CbcI96\nwgH3/J/HCMbQTnk1W10D00rEqMT5EUL+cfUplnoOTXWf/+sZivoIMoNhRtPts0tH\nmk7ZPk4o7uDfF5DWJJTfLw==\n-----END CERTIFICATE-----\n",
		}))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kubernetes Environment Provider Suite")
}
