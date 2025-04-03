package env

import (
	"fmt"
	"os"
)

// PASSH_USER
func SetPasshUserEnv(input string) error {
	file, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintf(file, "%s=%s\n", "PASSH_USER", input)

	return nil
}

// PASSH_TIMEOUT
func SetPasshTimeoutEnv(input string) error {
	file, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintf(file, "%s=%s\n", "PASSH_TIMEOUT", input)
	return nil
}

// PASSH_PASS
func SetPasshTempPassEnv(input string) error {
	file, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintf(file, "%s=%s\n", "PASSH_PASS", input)
	return nil
}

// PASSH_PERSISTENT_PASS
func SetPasshPersistPassEnv(input string) error {
	file, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintf(file, "%s=%s\n", "PASSH_PERSISTENT_PASS", input)
	return nil
}
