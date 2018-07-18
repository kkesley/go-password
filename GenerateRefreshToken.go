package password

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/kkesley/random"
)

//GenerateRefreshToken generates refresh token
func GenerateRefreshToken() (RefreshToken, error) {
	secret := []byte(random.MakeID(16))
	message := []byte(random.MakeID(32))
	hash := hmac.New(sha256.New, secret)
	if _, err := hash.Write(message); err != nil {
		return RefreshToken{}, err
	}

	return RefreshToken{
		Secret: string(secret),
		Token:  string(message),
		Hash:   hex.EncodeToString(hash.Sum(nil)),
	}, nil
}

//RefreshToken holds the refresh token
type RefreshToken struct {
	Secret string `json:"-"`
	Token  string `json:"-"`
	Hash   string `json:"-"`
}
