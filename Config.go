package password

import (
	"crypto/sha512"
	"hash"
)

//Config hashing configurations
type Config struct {
	HashBytes  int
	SaltBytes  int
	Iterations int
	Algo       func() hash.Hash
	Encoding   string
}

//DefaultConfig get default config for hashing
func DefaultConfig() Config {
	return CreateConfig(64, 16, 500000)
}

//CreateConfig create specialized config
func CreateConfig(hashBytes int, saltBytes int, iterations int) Config {
	return Config{
		HashBytes:  hashBytes,
		SaltBytes:  saltBytes,
		Iterations: iterations,
		Algo:       sha512.New,
		Encoding:   "base64",
	}
}
