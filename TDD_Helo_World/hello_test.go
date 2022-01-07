package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, received, expected string) {
		t.Helper()
		if received != expected {
			t.Errorf("received %q expected %q", received, expected)
		}
	}

	t.Run("When passed a name function returns a string saying Hello <name>", func(t *testing.T) {
		received := Hello("Tim", "")
		expected := "Hello Tim"

		assertCorrectMessage(t, received, expected)
	})

	t.Run("When passed an empty string function returns a string saying Hello World", func(t *testing.T) {
		received := Hello("", "")
		expected := "Hello World"

		assertCorrectMessage(t, received, expected)
	})

	t.Run("When passed a second argument of language string function returns a string saying Hello World in that given language", func(t *testing.T) {
		received := Hello("Tim", "spanish")
		expected := "Hola Tim"

		assertCorrectMessage(t, received, expected)
	})

	t.Run("When passed a second argument of language string function returns a string saying Hello World in that given language", func(t *testing.T) {
		received := Hello("Tim", "french")
		expected := "Bonjour Tim"

		assertCorrectMessage(t, received, expected)
	})
}
