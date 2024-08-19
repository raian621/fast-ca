package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"path"

	"golang.org/x/crypto/argon2"
)

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

func HashPassword(password string) (string, error) {
	var (
		memory      uint32 = 64 * 1024
		iterations  uint32 = 3
		parallelism uint8  = 2
		saltLength  uint32 = 16
		keyLength   uint32 = 32
	)

	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

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

  fmt.Println(len(fullHash))

	return fullHash, nil
}
