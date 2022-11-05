package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Graham")
	want := "Hello, Graham"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
