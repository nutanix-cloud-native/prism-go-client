package v3

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-logr/logr"
	"github.com/keploy/go-sdk/integrations/khttpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
)

func TestNewV3ClientBasicAuth(t *testing.T) {
	// verifies positive client creation
	cred := prismgoclient.Credentials{
		URL:                "foo.com",
		Username:           "username",
		Password:           "password",
		Port:               "",
		Endpoint:           "0.0.0.0",
		Insecure:           true,
		FoundationEndpoint: "10.0.0.0",
		FoundationPort:     "8000",
		RequiredFields:     nil,
	}
	v3Client, err := NewV3Client(cred)
	assert.NoError(t, err)
	assert.NotNil(t, v3Client)

	// verify missing client scenario
	cred = prismgoclient.Credentials{
		URL:      "foo.com",
		Insecure: true,
		RequiredFields: map[string][]string{
			"prism_central": {"username", "password", "endpoint"},
		},
	}

	v3Client, err = NewV3Client(cred)
	assert.Nil(t, v3Client)
	assert.EqualError(t, err, "username, password and endpoint are required for basic auth")
}

func TestNewV3ClientAPIKey(t *testing.T) {
	// verifies positive client creation
	cred := prismgoclient.Credentials{
		URL:                "foo.com",
		Port:               "",
		Endpoint:           "0.0.0.0",
		Insecure:           true,
		FoundationEndpoint: "10.0.0.0",
		FoundationPort:     "8000",
		RequiredFields:     nil,
		APIKey:             "my-api-key",
	}
	v3Client, err := NewV3Client(cred)
	assert.NoError(t, err)
	assert.NotNil(t, v3Client)

	// verify missing client scenario
	cred = prismgoclient.Credentials{
		URL:      "foo.com",
		Insecure: true,
		RequiredFields: map[string][]string{
			"prism_central": {"username", "password", "endpoint"},
		},
		APIKey: "my-api-key",
	}

	v3Client, err = NewV3Client(cred)
	assert.Nil(t, v3Client)
	assert.EqualError(t, err, "endpoint is required for api key auth")
}

func TestNewV3ClientWithPEMEncodedCertBundle(t *testing.T) {
	// verifies positive client creation
	cred := prismgoclient.Credentials{
		URL:      "foo.com",
		Username: "username",
		Password: "password",
		Port:     "9440",
		Endpoint: "foo.com",
	}
	certBundle := `-----BEGIN CERTIFICATE-----
MIIDdTCCAl2gAwIBAgILBAAAAAABFUtaw5QwDQYJKoZIhvcNAQEFBQAwVzELMAkG
A1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv
b3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw05ODA5MDExMjAw
MDBaFw0yODAxMjgxMjAwMDBaMFcxCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i
YWxTaWduIG52LXNhMRAwDgYDVQQLEwdSb290IENBMRswGQYDVQQDExJHbG9iYWxT
aWduIFJvb3QgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDaDuaZ
jc6j40+Kfvvxi4Mla+pIH/EqsLmVEQS98GPR4mdmzxzdzxtIK+6NiY6arymAZavp
xy0Sy6scTHAHoT0KMM0VjU/43dSMUBUc71DuxC73/OlS8pF94G3VNTCOXkNz8kHp
1Wrjsok6Vjk4bwY8iGlbKk3Fp1S4bInMm/k8yuX9ifUSPJJ4ltbcdG6TRGHRjcdG
snUOhugZitVtbNV4FpWi6cgKOOvyJBNPc1STE4U6G7weNLWLBYy5d4ux2x8gkasJ
U26Qzns3dLlwR5EiUWMWea6xrkEmCMgZK9FGqkjWZCrXgzT/LCrBbBlDSgeF59N8
9iFo7+ryUp9/k5DPAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8E
BTADAQH/MB0GA1UdDgQWBBRge2YaRQ2XyolQL30EzTSo//z9SzANBgkqhkiG9w0B
AQUFAAOCAQEA1nPnfE920I2/7LqivjTFKDK1fPxsnCwrvQmeU79rXqoRSLblCKOz
yj1hTdNGCbM+w6DjY1Ub8rrvrTnhQ7k4o+YviiY776BQVvnGCv04zcQLcFGUl5gE
38NflNUVyRRBnMRddWQVDf9VMOyGj/8N7yy5Y0b2qvzfvGn9LhJIZJrglfCm7ymP
AbEVtQwdpf5pLGkkeB6zpxxxYu7KyJesF12KwvhHhm4qxFYxldBniYUr+WymXUad
DKqC5JlR3XC321Y9YeRq4VzW9v493kHMB65jUr9TU/Qr6cf9tveCX4XSQRjbgbME
HMUfpIBvFSDJ3gyICh3WZlXi/EjJKSZp4A==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIEYDCCA0igAwIBAgILBAAAAAABL07hRQwwDQYJKoZIhvcNAQEFBQAwVzELMAkG
A1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv
b3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw0xMTA0MTMxMDAw
MDBaFw0yMjA0MTMxMDAwMDBaMF0xCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i
YWxTaWduIG52LXNhMTMwMQYDVQQDEypHbG9iYWxTaWduIE9yZ2FuaXphdGlvbiBW
YWxpZGF0aW9uIENBIC0gRzIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB
AQDdNR3yIFQmGtDvpW+Bdllw3Of01AMkHyQOnSKf1Ccyeit87ovjYWI4F6+0S3qf
ZyEcLZVUunm6tsTyDSF0F2d04rFkCJlgePtnwkv3J41vNnbPMYzl8QbX3FcOW6zu
zi2rqqlwLwKGyLHQCAeV6irs0Z7kNlw7pja1Q4ur944+ABv/hVlrYgGNguhKujiz
4MP0bRmn6gXdhGfCZsckAnNate6kGdn8AM62pI3ffr1fsjqdhDFPyGMM5NgNUqN+
ARvUZ6UYKOsBp4I82Y4d5UcNuotZFKMfH0vq4idGhs6dOcRmQafiFSNrVkfB7cVT
5NSAH2v6gEaYsgmmD5W+ZoiTAgMBAAGjggElMIIBITAOBgNVHQ8BAf8EBAMCAQYw
EgYDVR0TAQH/BAgwBgEB/wIBADAdBgNVHQ4EFgQUXUayjcRLdBy77fVztjq3OI91
nn4wRwYDVR0gBEAwPjA8BgRVHSAAMDQwMgYIKwYBBQUHAgEWJmh0dHBzOi8vd3d3
Lmdsb2JhbHNpZ24uY29tL3JlcG9zaXRvcnkvMDMGA1UdHwQsMCowKKAmoCSGImh0
dHA6Ly9jcmwuZ2xvYmFsc2lnbi5uZXQvcm9vdC5jcmwwPQYIKwYBBQUHAQEEMTAv
MC0GCCsGAQUFBzABhiFodHRwOi8vb2NzcC5nbG9iYWxzaWduLmNvbS9yb290cjEw
HwYDVR0jBBgwFoAUYHtmGkUNl8qJUC99BM00qP/8/UswDQYJKoZIhvcNAQEFBQAD
ggEBABvgiADHBREc/6stSEJSzSBo53xBjcEnxSxZZ6CaNduzUKcbYumlO/q2IQen
fPMOK25+Lk2TnLryhj5jiBDYW2FQEtuHrhm70t8ylgCoXtwtI7yw07VKoI5lkS/Z
9oL2dLLffCbvGSuXL+Ch7rkXIkg/pfcNYNUNUUflWP63n41edTzGQfDPgVRJEcYX
pOBWYdw9P91nbHZF2krqrhqkYE/Ho9aqp9nNgSvBZnWygI/1h01fwlr1kMbawb30
hag8IyrhFHvBN91i0ZJsumB9iOQct+R2UTjEqUdOqCsukNK1OFHrwZyKarXMsh3o
wFZUTKiL8IkyhtyTMr5NGvo1dbU=
-----END CERTIFICATE-----`
	v3Client, err := NewV3Client(cred, WithPEMEncodedCertBundle([]byte(certBundle)))
	require.NoError(t, err)
	assert.NotNil(t, v3Client)
}

func TestNewV3ClientWithInsecureAndCustomTransport(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	// Changing insecure to true as that leads to modifying the transport underneath
	creds.Insecure = true

	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	v3Client, err := NewV3Client(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)
	assert.NotNil(t, v3Client)
}

func TestWithCertificate(t *testing.T) {
	t.Run("valid certificate", func(t *testing.T) {
		// Create a test certificate
		cert := &x509.Certificate{
			Raw: []byte("test-cert-data"),
		}

		client := &Client{}
		opt := WithCertificate(cert)

		err := opt(client)
		assert.NoError(t, err)
		assert.Len(t, client.clientOpts, 1)
	})

	t.Run("nil certificate", func(t *testing.T) {
		client := &Client{}
		opt := WithCertificate(nil)

		err := opt(client)
		assert.NoError(t, err) // WithCertificate doesn't validate nil, it passes to internal client
		assert.Len(t, client.clientOpts, 1)
	})
}

func TestWithPEMEncodedCertBundle(t *testing.T) {
	t.Run("valid PEM certificate bundle", func(t *testing.T) {
		// Use a valid test certificate from the existing test
		certBundle := `-----BEGIN CERTIFICATE-----
MIIDdTCCAl2gAwIBAgILBAAAAAABFUtaw5QwDQYJKoZIhvcNAQEFBQAwVzELMAkG
A1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv
b3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw05ODA5MDExMjAw
MDBaFw0yODAxMjgxMjAwMDBaMFcxCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i
YWxTaWduIG52LXNhMRAwDgYDVQQLEwdSb290IENBMRswGQYDVQQDExJHbG9iYWxT
aWduIFJvb3QgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDaDuaZ
jc6j40+Kfvvxi4Mla+pIH/EqsLmVEQS98GPR4mdmzxzdzxtIK+6NiY6arymAZavp
xy0Sy6scTHAHoT0KMM0VjU/43dSMUBUc71DuxC73/OlS8pF94G3VNTCOXkNz8kHp
1Wrjsok6Vjk4bwY8iGlbKk3Fp1S4bInMm/k8yuX9ifUSPJJ4ltbcdG6TRGHRjcdG
snUOhugZitVtbNV4FpWi6cgKOOvyJBNPc1STE4U6G7weNLWLBYy5d4ux2x8gkasJ
U26Qzns3dLlwR5EiUWMWea6xrkEmCMgZK9FGqkjWZCrXgzT/LCrBbBlDSgeF59N8
9iFo7+ryUp9/k5DPAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8E
BTADAQH/MB0GA1UdDgQWBBRge2YaRQ2XyolQL30EzTSo//z9SzANBgkqhkiG9w0B
AQUFAAOCAQEA1nPnfE920I2/7LqivjTFKDK1fPxsnCwrvQmeU79rXqoRSLblCKOz
yj1hTdNGCbM+w6DjY1Ub8rrvrTnhQ7k4o+YviiY776BQVvnGCv04zcQLcFGUl5gE
38NflNUVyRRBnMRddWQVDf9VMOyGj/8N7yy5Y0b2qvzfvGn9LhJIZJrglfCm7ymP
AbEVtQwdpf5pLGkkeB6zpxxxYu7KyJesF12KwvhHhm4qxFYxldBniYUr+WymXUad
DKqC5JlR3XC321Y9YeRq4VzW9v493kHMB65jUr9TU/Qr6cf9tveCX4XSQRjbgbME
HMUfpIBvFSDJ3gyICh3WZlXi/EjJKSZp4A==
-----END CERTIFICATE-----`

		client := &Client{}
		opt := WithPEMEncodedCertBundle([]byte(certBundle))

		err := opt(client)
		assert.NoError(t, err)
		assert.Len(t, client.clientOpts, 1)
	})

	t.Run("multiple certificates in bundle", func(t *testing.T) {
		// Use valid test certificates
		certBundle := `-----BEGIN CERTIFICATE-----
MIIDdTCCAl2gAwIBAgILBAAAAAABFUtaw5QwDQYJKoZIhvcNAQEFBQAwVzELMAkG
A1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv
b3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw05ODA5MDExMjAw
MDBaFw0yODAxMjgxMjAwMDBaMFcxCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i
YWxTaWduIG52LXNhMRAwDgYDVQQLEwdSb290IENBMRswGQYDVQQDExJHbG9iYWxT
aWduIFJvb3QgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDaDuaZ
jc6j40+Kfvvxi4Mla+pIH/EqsLmVEQS98GPR4mdmzxzdzxtIK+6NiY6arymAZavp
xy0Sy6scTHAHoT0KMM0VjU/43dSMUBUc71DuxC73/OlS8pF94G3VNTCOXkNz8kHp
1Wrjsok6Vjk4bwY8iGlbKk3Fp1S4bInMm/k8yuX9ifUSPJJ4ltbcdG6TRGHRjcdG
snUOhugZitVtbNV4FpWi6cgKOOvyJBNPc1STE4U6G7weNLWLBYy5d4ux2x8gkasJ
U26Qzns3dLlwR5EiUWMWea6xrkEmCMgZK9FGqkjWZCrXgzT/LCrBbBlDSgeF59N8
9iFo7+ryUp9/k5DPAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8E
BTADAQH/MB0GA1UdDgQWBBRge2YaRQ2XyolQL30EzTSo//z9SzANBgkqhkiG9w0B
AQUFAAOCAQEA1nPnfE920I2/7LqivjTFKDK1fPxsnCwrvQmeU79rXqoRSLblCKOz
yj1hTdNGCbM+w6DjY1Ub8rrvrTnhQ7k4o+YviiY776BQVvnGCv04zcQLcFGUl5gE
38NflNUVyRRBnMRddWQVDf9VMOyGj/8N7yy5Y0b2qvzfvGn9LhJIZJrglfCm7ymP
AbEVtQwdpf5pLGkkeB6zpxxxYu7KyJesF12KwvhHhm4qxFYxldBniYUr+WymXUad
DKqC5JlR3XC321Y9YeRq4VzW9v493kHMB65jUr9TU/Qr6cf9tveCX4XSQRjbgbME
HMUfpIBvFSDJ3gyICh3WZlXi/EjJKSZp4A==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIEYDCCA0igAwIBAgILBAAAAAABL07hRQwwDQYJKoZIhvcNAQEFBQAwVzELMAkG
A1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv
b3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw0xMTA0MTMxMDAw
MDBaFw0yMjA0MTMxMDAwMDBaMF0xCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i
YWxTaWduIG52LXNhMTMwMQYDVQQDEypHbG9iYWxTaWduIE9yZ2FuaXphdGlvbiBW
YWxpZGF0aW9uIENBIC0gRzIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB
AQDdNR3yIFQmGtDvpW+Bdllw3Of01AMkHyQOnSKf1Ccyeit87ovjYWI4F6+0S3qf
ZyEcLZVUunm6tsTyDSF0F2d04rFkCJlgePtnwkv3J41vNnbPMYzl8QbX3FcOW6zu
zi2rqqlwLwKGyLHQCAeV6irs0Z7kNlw7pja1Q4ur944+ABv/hVlrYgGNguhKujiz
4MP0bRmn6gXdhGfCZsckAnNate6kGdn8AM62pI3ffr1fsjqdhDFPyGMM5NgNUqN+
ARvUZ6UYKOsBp4I82Y4d5UcNuotZFKMfH0vq4idGhs6dOcRmQafiFSNrVkfB7cVT
5NSAH2v6gEaYsgmmD5W+ZoiTAgMBAAGjggElMIIBITAOBgNVHQ8BAf8EBAMCAQYw
EgYDVR0TAQH/BAgwBgEB/wIBADAdBgNVHQ4EFgQUXUayjcRLdBy77fVztjq3OI91
nn4wRwYDVR0gBEAwPjA8BgRVHSAAMDQwMgYIKwYBBQUHAgEWJmh0dHBzOi8vd3d3
Lmdsb2JhbHNpZ24uY29tL3JlcG9zaXRvcnkvMDMGA1UdHwQsMCowKKAmoCSGImh0
dHA6Ly9jcmwuZ2xvYmFsc2lnbi5uZXQvcm9vdC5jcmwwPQYIKwYBBQUHAQEEMTAv
MC0GCCsGAQUFBzABhiFodHRwOi8vb2NzcC5nbG9iYWxzaWduLmNvbS9yb290cjEw
HwYDVR0jBBgwFoAUYHtmGkUNl8qJUC99BM00qP/8/UswDQYJKoZIhvcNAQEFBQAD
ggEBABvgiADHBREc/6stSEJSzSBo53xBjcEnxSxZZ6CaNduzUKcbYumlO/q2IQen
fPMOK25+Lk2TnLryhj5jiBDYW2FQEtuHrhm70t8ylgCoXtwtI7yw07VKoI5lkS/Z
9oL2dLLffCbvGSuXL+Ch7rkXIkg/pfcNYNUNUUflWP63n41edTzGQfDPgVRJEcYX
pOBWYdw9P91nbHZF2krqrhqkYE/Ho9aqp9nNgSvBZnWygI/1h01fwlr1kMbawb30
hag8IyrhFHvBN91i0ZJsumB9iOQct+R2UTjEqUdOqCsukNK1OFHrwZyKarXMsh3o
wFZUTKiL8IkyhtyTMr5NGvo1dbU=
-----END CERTIFICATE-----`

		client := &Client{}
		opt := WithPEMEncodedCertBundle([]byte(certBundle))

		err := opt(client)
		assert.NoError(t, err)
		assert.Len(t, client.clientOpts, 2)
	})

	t.Run("invalid PEM block type", func(t *testing.T) {
		// Create invalid PEM block
		invalidPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "INVALID",
			Bytes: []byte("test-data"),
		})

		client := &Client{}
		opt := WithPEMEncodedCertBundle(invalidPEM)

		err := opt(client)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected PEM block type")
	})

	t.Run("empty certificate bundle", func(t *testing.T) {
		client := &Client{}
		opt := WithPEMEncodedCertBundle([]byte{})

		err := opt(client)
		assert.NoError(t, err)
		assert.Len(t, client.clientOpts, 0)
	})

	t.Run("invalid certificate data", func(t *testing.T) {
		// Create PEM block with invalid certificate data
		invalidPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: []byte("invalid-cert-data"),
		})

		client := &Client{}
		opt := WithPEMEncodedCertBundle(invalidPEM)

		err := opt(client)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "x509: malformed certificate")
	})
}

func TestWithRoundTripper(t *testing.T) {
	t.Run("valid round tripper", func(t *testing.T) {
		transport := &http.Transport{}
		client := &Client{}
		opt := WithRoundTripper(transport)

		err := opt(client)
		assert.NoError(t, err)
		assert.Len(t, client.clientOpts, 1)
	})

	t.Run("nil round tripper", func(t *testing.T) {
		client := &Client{}
		opt := WithRoundTripper(nil)

		err := opt(client)
		assert.NoError(t, err)
		assert.Len(t, client.clientOpts, 1)
	})
}

func TestWithLogger(t *testing.T) {
	t.Run("valid logger", func(t *testing.T) {
		logger := logr.Discard()
		client := &Client{}
		opt := WithLogger(&logger)

		err := opt(client)
		assert.NoError(t, err)
		assert.Len(t, client.clientOpts, 1)
	})

	t.Run("nil logger", func(t *testing.T) {
		client := &Client{}
		opt := WithLogger(nil)

		err := opt(client)
		assert.NoError(t, err)
		assert.Len(t, client.clientOpts, 1)
	})
}

func TestWithUserAgent(t *testing.T) {
	t.Run("valid user agent", func(t *testing.T) {
		userAgent := "test-user-agent"
		client := &Client{}
		opt := WithUserAgent(userAgent)

		err := opt(client)
		assert.NoError(t, err)
		assert.Len(t, client.clientOpts, 1)
	})

	t.Run("empty user agent", func(t *testing.T) {
		userAgent := ""
		client := &Client{}
		opt := WithUserAgent(userAgent)

		err := opt(client)
		assert.NoError(t, err)
		assert.Len(t, client.clientOpts, 1)
	})
}

func TestNewV3ClientWithOptions(t *testing.T) {
	t.Run("with certificate option", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			APIKey:   "test-api-key",
			Endpoint: "https://test.example.com",
		}

		cert := &x509.Certificate{
			Raw: []byte("test-cert-data"),
		}

		client, err := NewV3Client(credentials, WithCertificate(cert))
		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("with PEM certificate bundle option", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			APIKey:   "test-api-key",
			Endpoint: "https://test.example.com",
		}

		// Use a valid test certificate
		certBundle := `-----BEGIN CERTIFICATE-----
MIIDdTCCAl2gAwIBAgILBAAAAAABFUtaw5QwDQYJKoZIhvcNAQEFBQAwVzELMAkG
A1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv
b3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw05ODA5MDExMjAw
MDBaFw0yODAxMjgxMjAwMDBaMFcxCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i
YWxTaWduIG52LXNhMRAwDgYDVQQLEwdSb290IENBMRswGQYDVQQDExJHbG9iYWxT
aWduIFJvb3QgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDaDuaZ
jc6j40+Kfvvxi4Mla+pIH/EqsLmVEQS98GPR4mdmzxzdzxtIK+6NiY6arymAZavp
xy0Sy6scTHAHoT0KMM0VjU/43dSMUBUc71DuxC73/OlS8pF94G3VNTCOXkNz8kHp
1Wrjsok6Vjk4bwY8iGlbKk3Fp1S4bInMm/k8yuX9ifUSPJJ4ltbcdG6TRGHRjcdG
snUOhugZitVtbNV4FpWi6cgKOOvyJBNPc1STE4U6G7weNLWLBYy5d4ux2x8gkasJ
U26Qzns3dLlwR5EiUWMWea6xrkEmCMgZK9FGqkjWZCrXgzT/LCrBbBlDSgeF59N8
9iFo7+ryUp9/k5DPAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8E
BTADAQH/MB0GA1UdDgQWBBRge2YaRQ2XyolQL30EzTSo//z9SzANBgkqhkiG9w0B
AQUFAAOCAQEA1nPnfE920I2/7LqivjTFKDK1fPxsnCwrvQmeU79rXqoRSLblCKOz
yj1hTdNGCbM+w6DjY1Ub8rrvrTnhQ7k4o+YviiY776BQVvnGCv04zcQLcFGUl5gE
38NflNUVyRRBnMRddWQVDf9VMOyGj/8N7yy5Y0b2qvzfvGn9LhJIZJrglfCm7ymP
AbEVtQwdpf5pLGkkeB6zpxxxYu7KyJesF12KwvhHhm4qxFYxldBniYUr+WymXUad
DKqC5JlR3XC321Y9YeRq4VzW9v493kHMB65jUr9TU/Qr6cf9tveCX4XSQRjbgbME
HMUfpIBvFSDJ3gyICh3WZlXi/EjJKSZp4A==
-----END CERTIFICATE-----`

		client, err := NewV3Client(credentials, WithPEMEncodedCertBundle([]byte(certBundle)))
		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("with round tripper option", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			APIKey:   "test-api-key",
			Endpoint: "https://test.example.com",
		}

		transport := &http.Transport{}

		client, err := NewV3Client(credentials, WithRoundTripper(transport))
		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("with logger option", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			APIKey:   "test-api-key",
			Endpoint: "https://test.example.com",
		}

		logger := logr.Discard()

		client, err := NewV3Client(credentials, WithLogger(&logger))
		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("with user agent option", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			APIKey:   "test-api-key",
			Endpoint: "https://test.example.com",
		}

		userAgent := "custom-user-agent"

		client, err := NewV3Client(credentials, WithUserAgent(userAgent))
		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("with multiple options", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			APIKey:   "test-api-key",
			Endpoint: "https://test.example.com",
		}

		cert := &x509.Certificate{
			Raw: []byte("test-cert-data"),
		}

		logger := logr.Discard()
		userAgent := "custom-user-agent"

		client, err := NewV3Client(credentials,
			WithCertificate(cert),
			WithLogger(&logger),
			WithUserAgent(userAgent),
		)
		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("with invalid option", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			APIKey:   "test-api-key",
			Endpoint: "https://test.example.com",
		}

		// Create an option that returns an error
		invalidOpt := func(c *Client) error {
			return fmt.Errorf("invalid option")
		}

		client, err := NewV3Client(credentials, invalidOpt)
		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "invalid option")
	})
}

func TestConstants(t *testing.T) {
	t.Run("library version", func(t *testing.T) {
		assert.Equal(t, "v3", libraryVersion)
	})

	t.Run("absolute path", func(t *testing.T) {
		assert.Equal(t, "api/nutanix/v3", absolutePath)
	})

	t.Run("user agent", func(t *testing.T) {
		assert.Equal(t, "nutanix/v3", userAgent)
	})
}

func TestClientStructure(t *testing.T) {
	t.Run("client fields", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			APIKey:   "test-api-key",
			Endpoint: "https://test.example.com",
		}

		client, err := NewV3Client(credentials)
		require.NoError(t, err)

		assert.NotNil(t, client.client)
		assert.NotNil(t, client.V3)
		assert.NotNil(t, client.clientOpts)
	})

	t.Run("client options are preserved", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			APIKey:   "test-api-key",
			Endpoint: "https://test.example.com",
		}

		cert := &x509.Certificate{
			Raw: []byte("test-cert-data"),
		}

		client, err := NewV3Client(credentials, WithCertificate(cert))
		require.NoError(t, err)

		// The clientOpts should contain the certificate option
		assert.Len(t, client.clientOpts, 4) // credentials, userAgent, absolutePath, certificate
	})
}

func TestClientWithInvalidCredentials(t *testing.T) {
	t.Run("empty credentials", func(t *testing.T) {
		credentials := prismgoclient.Credentials{}

		client, err := NewV3Client(credentials)
		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "username, password and endpoint are required for basic auth")
	})

	t.Run("credentials with only username", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			Username: "testuser",
		}

		client, err := NewV3Client(credentials)
		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "username, password and endpoint are required for basic auth")
	})

	t.Run("credentials with only password", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			Password: "testpass",
		}

		client, err := NewV3Client(credentials)
		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "username, password and endpoint are required for basic auth")
	})

	t.Run("credentials with only endpoint", func(t *testing.T) {
		credentials := prismgoclient.Credentials{
			Endpoint: "https://test.example.com",
		}

		client, err := NewV3Client(credentials)
		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "username, password and endpoint are required for basic auth")
	})
}
