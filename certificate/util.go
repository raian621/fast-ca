package certificate

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

var ErrIncorrectKeyLength = errors.New("key length must be 32 bytes long for AES-256 encryption")

// Encrypts `pemBytes` using AES-256-GCM encryption with `key` as a passphrase.
// `key` must be 32 bytes for AES-256 encryption
func Encrypt(pemBytes, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return []byte{}, ErrIncorrectKeyLength
	}
	ciph, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	gcm, err := cipher.NewGCM(ciph)
	if err != nil {
		return []byte{}, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, err
	}

	encrypted := gcm.Seal(nonce, nonce, pemBytes, nil)

	return encrypted, nil
}

// Decrypts `encryptedBytes` using AES-256-GCM decryption with `key` as a
// passphrase. `key` must be 32 bytes for AES-256 decryption
func Decrypt(encryptedBytes, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return []byte{}, ErrIncorrectKeyLength
	}
	ciph, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	gcm, err := cipher.NewGCM(ciph)
	if err != nil {
		return []byte{}, err
	}

	nonce, cipherText := encryptedBytes[:gcm.NonceSize()], encryptedBytes[gcm.NonceSize():]

	return gcm.Open(nil, nonce, cipherText, nil)
}
