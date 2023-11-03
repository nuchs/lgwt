package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to someone", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris!"
		assertEqual(t, got, want)
	})
	t.Run("Say 'Hello, World!' by default", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		assertEqual(t, got, want)
	})
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
