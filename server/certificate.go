package server

import (
	"crypto/x509"
	"io"
	"log"
	"net/url"
	"time"

	"github.com/raian621/fast-ca/certificate"
	"github.com/raian621/fast-ca/models"
)

func newCA(db models.Queryable, formData url.Values, certFile, keyFile io.Writer) error {
  log.Println(formData)
  expires, err := time.Parse(time.RFC3339[:len(formData["expires"][0])], formData["expires"][0])
  if err != nil {
    log.Println(err)
    return err
  }

  // returns an empty string if the form data field `key` isn't present in the
  // `values` map
  unwrapFormValue := func(values url.Values, key string) string {
    if value, ok := values[key]; ok {
      return value[0]
    }
    return ""
  }

  config := certificate.CertificateConfig{
    Subject: certificate.CertificateSubject{
      Organization: unwrapFormValue(formData, "subject[organization]"),
      Country: unwrapFormValue(formData, "subject[country]"),
      Province: unwrapFormValue(formData, "subject[province]"),
      Locality: unwrapFormValue(formData, "subject[locality]"),
      StreetAddress: unwrapFormValue(formData, "subject[street_address]"),
      PostalCode: unwrapFormValue(formData, "subject[postal_code]"),
      CommonName: unwrapFormValue(formData, "subject[common_name]"),
    },
    Expires: expires,
    KeyType: x509.RSA,
  }

  caCertBytes, caPrivKeyBytes, err := certificate.CreateCA(&config)
  if err != nil {
    return err
  }

  encryptedCaCert, err := certificate.Encrypt(caCertBytes, []byte("12345678901234567890123456789012"))
  if err != nil {
    return err
  }
  encryptedCaKey, err := certificate.Encrypt(caPrivKeyBytes, []byte("12345678901234567890123456789012"))
  if err != nil {
    return err
  }

  if _, err := certFile.Write(encryptedCaCert); err != nil {
    return err
  }
  if _, err := keyFile.Write(encryptedCaKey); err != nil {
    return err
  }

  return nil
}
