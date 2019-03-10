package underscore

import (
	"testing"
)

func TestIsList(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got '%t' want '%t'", got, want)
		}
	}

	t.Run("Argument as a list", func(t *testing.T) {
		got := IsList([]int{})
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("Argument", func(t *testing.T) {
		got := IsList(map[string]string{})
		want := false
		assertCorrectMessage(t, got, want)
	})
}
