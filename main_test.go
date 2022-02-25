package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sanndy", "English")
		want := "Hello, Sanndy"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying empty words", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in German", func(t *testing.T) {
		got := Hello("Hans", "German")
		want := "Hallo, Hans"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in German without name", func(t *testing.T) {
		got := Hello("", "German")
		want := "Hallo, Welt"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish without name", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, Tierra"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Hindi", func(t *testing.T) {
		got := Hello("Abhe", "Hindi")
		want := "Hello, Abhe"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Hindi without name", func(t *testing.T) {
		got := Hello("", "Hindi")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func ExampleHello() {
	greeting := Hello("Abhe", "English")
	fmt.Print(greeting)
	// Output: Hello, Abhe
}
