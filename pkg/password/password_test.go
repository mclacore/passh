package password_test

import (
	"strings"
	"testing"

	"github.com/mclacore/passh/pkg/password"
)

var (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	special   = "!@#$%^&*()_+~><"
)

func TestSimplePassword(t *testing.T) {
	length := 12
	fail := uppercase + numbers + special
	got := password.GeneratePassword(length, false, false, false, false)

	if len(got) != length {
		t.Errorf("password length mismatch, got: %v, want: %v", len(got), length)
	}

	if !strings.ContainsAny(got, lowercase) {
		t.Errorf("password does not contain lowercase, got: %v", got)
	}

	if strings.ContainsAny(got, fail) {
		t.Errorf("password contains unwanted chars, got: %v", got)
	}
}

func TestLowerUpperPassword(t *testing.T) {
	length := 8
	got := password.GeneratePassword(length, false, true, false, false)

	if len(got) != length {
		t.Errorf("password lenght mismatch, got: %v, want: %v", len(got), length)
	}

	if !strings.ContainsAny(got, lowercase) {
		t.Errorf("password does not contain lowercase, got: %v", got)
	}

	if !strings.ContainsAny(got, uppercase) {
		t.Errorf("password does not contain uppercase, got: %v", got)
	}

	if strings.ContainsAny(got, numbers) {
		t.Errorf("password contains unwanted numbers, got: %v", got)
	}

	if strings.ContainsAny(got, special) {
		t.Errorf("password contains unwanted specials, got: %v", got)
	}
}

func TestLowerNumberPassword(t *testing.T) {
	length := 15
	got := password.GeneratePassword(length, false, false, true, false)

	if len(got) != length {
		t.Errorf("password length mismatch, got: %v, want: %v", len(got), length)
	}

	if !strings.ContainsAny(got, lowercase) {
		t.Errorf("password does not contain lowercase, got: %v", got)
	}

	if !strings.ContainsAny(got, numbers) {
		t.Errorf("password does not contain numbers, got: %v", got)
	}

	if strings.ContainsAny(got, uppercase) {
		t.Errorf("password contains unwanted uppercase, got: %v", got)
	}
	
	if strings.ContainsAny(got, special) {
		t.Errorf("password contains unwanted specials, got: %v", got)
	}
}

func TestLowerSpecialPassword(t *testing.T) {
	length := 15
	got := password.GeneratePassword(length, false, false, false, true)

	if len(got) != length {
		t.Errorf("password length mismatch, got: %v, want: %v", len(got), length)
	}

	if !strings.ContainsAny(got, lowercase) {
		t.Errorf("password does not contain lowercase, got: %v", got)
	}

	if !strings.ContainsAny(got, special) {
		t.Errorf("password does not contain specials, got: %v", got)
	}

	if strings.ContainsAny(got, uppercase) {
		t.Errorf("password contains unwanted uppercase, got: %v", got)
	}

	if strings.ContainsAny(got, numbers) {
		t.Errorf("password contains unwanted numbers, got: %v", got)
	}
}

func TestUpperNumberSpecial(t *testing.T) {
	length := 20
	got := password.GeneratePassword(length, true, true, true, true)

	if len(got) != length {
		t.Errorf("password length mismatch, got: %v, want: %v", len(got), length)
	}

	if !strings.ContainsAny(got, uppercase) {
		t.Errorf("password does not contain uppercase, got: %v", got)
	}

	if !strings.ContainsAny(got, numbers) {
		t.Errorf("password does not contain numbers, got: %v", got)
	}

	if !strings.ContainsAny(got, special) {
		t.Errorf("password does not contain specials, got: %v", got)
	}

	if strings.ContainsAny(got, lowercase) {
		t.Errorf("password contains unwanted lowercase, got: %v", got)
	}
}

func TestLowerUpperNumberSpecial(t *testing.T) {
	length := 25
	got := password.GeneratePassword(length, false, true, true, true)

	if len(got) != length {
		t.Errorf("password length mismatch, got: %v, want: %v", len(got), length)
	}

	if !strings.ContainsAny(got, lowercase) {
		t.Errorf("password does not contain lowercase, got: %v", got)
	}

	if !strings.ContainsAny(got, uppercase) {
		t.Errorf("password does not contain uppercase, got: %v", got)
	}

	if !strings.ContainsAny(got, numbers) {
		t.Errorf("password does not contain numbers, got: %v", got)
	}

	if !strings.ContainsAny(got, special) {
		t.Errorf("password does not contain specials, got: %v", got)
	}
}
