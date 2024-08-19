package certificate

import (
	"reflect"
	"testing"
)

var exampleKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`)

func TestCertificateEncryption(t *testing.T) {
	key := []byte("01234567890123456789012345678901")
	cipherText, err := Encrypt(exampleKey, key)
	if err != nil {
		t.Fatal(err)
	}

	decryptedKey, err := Decrypt(cipherText, key)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(exampleKey, decryptedKey) {
		t.Fatalf(
			"Expected `%s`, got `%s`",
			string(exampleKey),
			string(decryptedKey),
		)
	}
}

func TestCertificateEncryptionIncorrectKeyLength(t *testing.T) {
	key := []byte("1234")
	_, err := Encrypt(exampleKey, key)
	if err != ErrIncorrectKeyLength {
		t.Error("expected `ErrIncorrectKeyLength`")
	}
	_, err = Decrypt(exampleKey, key)
	if err != ErrIncorrectKeyLength {
		t.Error("expected `ErrIncorrectKeyLength`")
	}

	key = []byte("123456789012345678901234567890123")
	_, err = Encrypt(exampleKey, key)
	if err != ErrIncorrectKeyLength {
		t.Error("expected `ErrIncorrectKeyLength`")
	}
	_, err = Decrypt(exampleKey, key)
	if err != ErrIncorrectKeyLength {
		t.Error("expected `ErrIncorrectKeyLength`")
	}
}
