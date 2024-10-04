package password

import (
	"fmt"
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
		// something is panicing here when only enabling uppercase in password
		// also when running passh pass new -x true, it should error out but instead
		// is setting the charset to lowercase + uppercase
		password[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Println(charset)

	return string(password)
}
