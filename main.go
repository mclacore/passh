package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	special   = "!@#$%^&*()_+"
)

func generatePassword(n int) string {
	pw := strings.Builder{}
	pw.Grow(n)
	for i := 0; i < n; i++ {
		chars := lowercase + uppercase + numbers + special
		pw.WriteByte(chars[rand.Intn(len(chars))])
	}

	return pw.String()
}

func main() {
	fmt.Println(generatePassword(20))
}
