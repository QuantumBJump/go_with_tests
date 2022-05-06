package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got'%s', wanted '%s'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", english)
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when no name is supplied", func(t *testing.T) {
		got := Hello("", english)
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("en Espanol", func(t *testing.T) {
		got := Hello("Juan", spanish)
		want := "Hola, Juan"

		assertCorrectMessage(t, got, want)
	})
	t.Run("en Francais", func(t *testing.T) {
		got := Hello("Elodie", french)
		want := "Bonjour, Elodie"

		assertCorrectMessage(t, got, want)
	})
}
