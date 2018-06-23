package password

import (
	"testing"
)

func TestHasher(test *testing.T) {
	hashed1, err := Hasher("apassword", DefaultConfig())
	hashed2, err := Hasher("apassword", DefaultConfig())
	if err != nil {
		test.Fail()
	}
	if hashed1 == hashed2 {
		test.Fail()
	}
}
