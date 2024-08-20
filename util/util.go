package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"path"
	"strings"

	"golang.org/x/crypto/argon2"
)

var saltLength uint32 = 16

func RelativeToAbsolutePath(relativePath string) (string, error) {
	var baseDir string
	if envDir, present := os.LookupEnv("BASE_DIR"); present {
		baseDir = envDir
	} else {
		execPath, err := os.Executable()
		if err != nil {
			return "", err
		}
		baseDir = path.Dir(execPath)
	}

	return path.Join(baseDir, relativePath), nil
}

func HashPassword(password string, salt []byte) string {
	var (
		memory      uint32 = 64 * 1024
		iterations  uint32 = 3
		parallelism uint8  = 2
		keyLength   uint32 = 32
	)

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		iterations,
		memory,
		parallelism,
		keyLength,
	)

	hashBase64 := base64.RawStdEncoding.EncodeToString(hash)
	saltBase64 := base64.RawStdEncoding.EncodeToString(salt)

	fullHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		memory,
		iterations,
		parallelism,
		hashBase64,
		saltBase64,
	)

	return fullHash
}

// Generates a random 16 byte salt
func GenerateSalt() ([]byte, error) {
  salt := make([]byte, saltLength)
  if _, err := rand.Read(salt); err != nil {
    return salt, err
  }

  return salt, nil
}

// Extracts the salt from Argon2id hash
func ExtractSalt(passhash string) ([]byte, error) {
  splits := strings.Split(passhash, "$")
  saltBase64 := splits[len(splits)-1]
  salt := make([]byte, saltLength)
  _, err := base64.RawStdEncoding.Decode(salt, []byte(saltBase64)) 
  return salt, err
}

func ValidatePassword(password, passhash string) (bool, error) {
  salt, err := ExtractSalt(passhash)
  if err != nil {
    return false, err
  }
  passwordHash := HashPassword(password, salt)
  return passwordHash == passhash, nil
}
