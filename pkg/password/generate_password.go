package password

import (
	"math/rand"
)

func GeneratePassword(length int64, lowercase, uppercase, numbers, special bool) string {
	charset := ""

	if !lowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}

	if uppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if numbers {
		charset += "0123456789"
	}

	if special {
		charset += "!@#$%^&*()_+~><"
	}

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}
