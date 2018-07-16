package password

import (
	"bytes"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"
)

//Verifier verify password and a hash given a config
func Verifier(password string, hash string, config Config) bool {
	sDec, _ := base64.StdEncoding.DecodeString(hash)
	saltBytes := readUint32BE(sDec, 0)
	iterations := readUint32BE(sDec, 4)
	if uint32(len(sDec)) < saltBytes+8 {
		return false
	}
	salt := sDec[8 : saltBytes+8]
	hashBytes := sDec[8+saltBytes:]

	derivationKey := pbkdf2.Key([]byte(password), salt, int(iterations), len(hashBytes), config.Algo)

	equality := bytes.Equal(hashBytes, derivationKey)
	return equality
}
