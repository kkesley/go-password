package password

import (
	"github.com/kkesley/random"
)

//GenerateRefreshToken generates refresh token
func GenerateRefreshToken() (RefreshToken, error) {
	secret := []byte(random.MakeID(16))
	message := []byte(random.MakeID(32))
	return GenerateRefreshTokenHash(secret, message)
}
