package password_test

import (
	"testing"

	"github.com/mclacore/passh/pkg/password"
)

func TestSimplePassword(t *testing.T) {
	want := "abcdefghijkl"
	got := password.GeneratePassword(12, false, false, false, false)

	if len(got) != len(want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
