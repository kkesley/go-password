package password

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

//GenerateRefreshTokenHash generates refresh token given secret and message
func GenerateRefreshTokenHash(secret []byte, message []byte) (RefreshToken, error) {
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
