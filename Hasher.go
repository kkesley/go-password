package password

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

//Hasher hash functions given config
func Hasher(password string, config Config) (string, error) {
	salt := make([]byte, config.SaltBytes)
	if _, err := rand.Read(salt); err != nil {
		return "", errors.New("Cannot create hash")
	}
	derivationKey := pbkdf2.Key([]byte(password), salt, config.Iterations, config.HashBytes, config.Algo)
	saltBytes := make([]byte, 4)
	iterationBytes := make([]byte, 4)
	combinedBytes := make([]byte, 0)

	writeUint32BE(saltBytes, uint32(len(salt)))
	writeUint32BE(iterationBytes, uint32(config.Iterations))
	combinedBytes = append(combinedBytes, saltBytes...)
	combinedBytes = append(combinedBytes, iterationBytes...)
	combinedBytes = append(combinedBytes, salt...)
	combinedBytes = append(combinedBytes, derivationKey...)
	hash := base64.StdEncoding.EncodeToString(combinedBytes)
	return hash, nil
}
