package certificate

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"time"
)

// config.we can use to generate x509 certificates with
type CertificateConfig struct {
	Subject CertificateSubject
	Expires time.Time
	KeyType x509.PublicKeyAlgorithm
}

type CertificateSubject struct {
	Organization  string
	Country       string
	Province      string
	Locality      string
	StreetAddress string
	PostalCode    string
	CommonName    string
}

// Creates a certificate authority private key and certificate, returns the
// private key and certificate in as arrays of bytes in PEM encoding.
func CreateCA(config *CertificateConfig) ([]byte, []byte, error) {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization:  []string{config.Subject.Organization},
			Country:       []string{config.Subject.Country},
			Province:      []string{config.Subject.Province},
			Locality:      []string{config.Subject.Locality},
			StreetAddress: []string{config.Subject.StreetAddress},
			PostalCode:    []string{config.Subject.PostalCode},
			CommonName:    config.Subject.CommonName,
		},
		NotBefore: time.Now(),
		NotAfter:  config.Expires,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth,
		},
		KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		IsCA:     true,
	}

	caPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return []byte{}, []byte{}, err
	}

	caBytes, err := x509.CreateCertificate(
		rand.Reader,
		ca,
		ca,
		&caPrivateKey.PublicKey,
		caPrivateKey,
	)
	if err != nil {
		return []byte{}, []byte{}, err
	}

	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	caPrivateKeyBytes := x509.MarshalPKCS1PrivateKey(caPrivateKey)
	caPrivateKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivateKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: caPrivateKeyBytes,
	})

	return caPEM.Bytes(), caPrivateKeyPEM.Bytes(), nil
}

func CreateCertificate(config *CertificateConfig) ([]byte, []byte, error) {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization:  []string{config.Subject.Organization},
			Country:       []string{config.Subject.Country},
			Province:      []string{config.Subject.Province},
			Locality:      []string{config.Subject.Locality},
			StreetAddress: []string{config.Subject.StreetAddress},
			PostalCode:    []string{config.Subject.PostalCode},
			CommonName:    config.Subject.CommonName,
		},
		NotBefore: time.Now(),
		NotAfter:  config.Expires,
	}
	log.Println(cert)
	return []byte{}, []byte{}, nil
}
