package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Holla Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("French", func(t *testing.T) {
		got := Hello("Solomon", "French")
		want := "Bonjur Solomon"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// this is important b/c it will tell our test suit that this is a
	// helper function so when it fails the line number reported
	// will be in our main function

	t.Helper()

	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}
