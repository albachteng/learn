package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello with a name", func(t *testing.T) {
		got := Hello("Graham")
		want := "Hello, Graham"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello without a name", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
