package certutils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"math/big"
	"time"
)

// GenerateCACertForTesting generates a CA Certificate and returns a PEM encoded version of the certificate
// and a base64 encoded version of the PEM encoded certificate.
func GenerateCACertForTesting() (string, string, error) {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2022),
		Subject: pkix.Name{
			Organization:  []string{"Nutanix, Inc."},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Jose"},
			StreetAddress: []string{"Technology Drive"},
			PostalCode:    []string{"95110"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	// create our private and public key
	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", "", err
	}

	// create the CA
	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return "", "", err
	}

	// pem encode
	caPEM := new(bytes.Buffer)
	if err := pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	}); err != nil {
		return "", "", err
	}

	b64Cert := base64.StdEncoding.EncodeToString(caPEM.Bytes())
	return caPEM.String(), b64Cert, nil
}
