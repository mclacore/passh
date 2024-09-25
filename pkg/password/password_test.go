package password_test

import (
	"testing"
	"regexp"

	"github.com/mclacore/passh/pkg/password"
)

func TestSimplePassword(t *testing.T) {
	want := "abcdefghijkl"
	got := password.GeneratePassword(12, false, false, false, false)

	if len(got) != len(want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPasswordWithNumbers(t *testing.T) {
	password := "abcdefgh1234"
	want := regexp.MustCompile(`\d`).MatchString(password)
	got := password.GeneratePassword(12, false, false, true, false)

	if !want {
		t.Errorf("got %q, want %q", got, want)
	}
}
