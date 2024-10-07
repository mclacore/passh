package password

import (
	"fmt"
	"math/rand"
)

func GeneratePassword(length int, lowercase, uppercase, numbers, special bool) string {
	charset := []byte("")
	lower := []byte("abcdefghijklmnopqrstuvwxyz")
	upper := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digits := []byte("0123456789")
	specials := []byte("!@#$%^&*()_+~><")

	if !lowercase {
		charset = append(charset, lower...)
	}

	if uppercase {
		charset = append(charset, upper...)
	}

	if numbers {
		charset = append(charset, digits...)
	}

	if special {
		charset = append(charset, specials...)
	}

	password := make([]byte, length)
	for i := range password {
		// something is panicing here when only enabling uppercase in password
		// also when running passh pass new -x true, it should error out but instead
		// is setting the charset to lowercase + uppercase
		password[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Println(string(charset))

	return string(password)
}
