package password

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/mclacore/passh/pkg/config"
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

	return string(password)
}

func MasterPasswordTimeout(input string) {
	timeVal, timeValErr := config.LoadConfigValue("auth", "timeout")
	if timeValErr != nil {
		log.Printf("Error loading timeout value: %v", timeValErr)
		os.Exit(2)
	}

	if timeVal == "" {
		config.SaveConfigValue("auth", "timeout", "900")
	}

	timeout, timeoutErr := strconv.Atoi(timeVal)
	if timeoutErr != nil {
		log.Printf("Error converting timeout string to int: %v", timeoutErr)
		os.Exit(2)
	}

	time.Sleep(time.Duration(timeout) * time.Second)
	config.SaveConfigValue("auth", "timeout", "")
}
