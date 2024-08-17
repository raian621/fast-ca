package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"strings"
	"testing"
	"time"
)

func TestCreateCertificateAuthority(t *testing.T) {
	t.Parallel()

	data := CertificateConfig{
		Subject: CertificateSubject{
			Organization:  "An Organization",
			Country:       "United States",
			Province:      "",
			Locality:      "Texas",
			StreetAddress: "1234 Street Ln",
			PostalCode:    "12345",
		},
		Expires: time.Now().UTC().AddDate(10, 0, 0),
	}
	certBytes, keyBytes, err := CreateCA(&data)

	if err != nil {
		t.Fatal(err)
	}

	certStr := string(certBytes)
	keyStr := string(keyBytes)

	if !strings.HasPrefix(certStr, "-----BEGIN CERTIFICATE-----") {
		t.Error("expected certificate to begin with `-----BEGIN CERTIFICATE-----`")
	}
	if !strings.HasPrefix(keyStr, "-----BEGIN RSA PRIVATE KEY-----") {
		t.Error("expected key to begin with `-----BEGIN RSA PRIVATE KEY-----`")
	}
	if !strings.HasSuffix(certStr, "-----END CERTIFICATE-----\n") {
		t.Error("expected certificate to begin with `-----END CERTIFICATE-----`")
	}
	if !strings.HasSuffix(keyStr, "-----END RSA PRIVATE KEY-----\n") {
		t.Error("expected key to end with `-----END RSA PRIVATE KEY-----`")
	}

	certBlock, _ := pem.Decode(certBytes)

	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		t.Fatal(err)
	}

	if cert.Subject.Organization[0] != data.Subject.Organization {
		t.Errorf(
			"expected `%s`, got `%s` for organization",
			data.Subject.Organization,
			cert.Subject.Organization,
		)
	}
	if cert.Subject.Country[0] != data.Subject.Country {
		t.Errorf(
			"expected `%s`, got `%s` for country",
			data.Subject.Country,
			cert.Subject.Country,
		)
	}
	if cert.Subject.Province[0] != data.Subject.Province {
		t.Errorf(
			"expected `%s`, got `%s` for province",
			data.Subject.Province,
			cert.Subject.Province,
		)
	}
	if cert.Subject.Locality[0] != data.Subject.Locality {
		t.Errorf(
			"expected `%s`, got `%s` for locality",
			data.Subject.Locality,
			cert.Subject.Locality,
		)
	}
	if cert.Subject.StreetAddress[0] != data.Subject.StreetAddress {
		t.Errorf(
			"expected `%s`, got `%s` for street address",
			data.Subject.StreetAddress,
			cert.Subject.StreetAddress,
		)
	}
	if cert.Subject.PostalCode[0] != data.Subject.PostalCode {
		t.Errorf(
			"expected `%s`, got `%s` for postal code",
			data.Subject.PostalCode,
			cert.Subject.PostalCode,
		)
	}
	if cert.NotAfter != data.Expires.Truncate(time.Second) {
		t.Errorf(
			"expected `%v`, got `%v` for expiration date",
			data.Expires.Truncate(time.Second),
			cert.NotAfter,
		)
	}
}
