package password

import (
	"testing"
)

func TestVerifier(test *testing.T) {
	password := "apassword"
	hashed, err := Hasher(password, DefaultConfig())
	if err != nil {
		test.Fail()
	}
	if !Verifier(password, hashed, DefaultConfig()) {
		test.Fail()
	}
}

func TestWrongVerifier(test *testing.T) {
	password := "apassword"
	password2 := "bpassword"
	hashed, err := Hasher(password, DefaultConfig())
	if err != nil {
		test.Fail()
	}
	if Verifier(password2, hashed, DefaultConfig()) {
		test.Fail()
	}
}
