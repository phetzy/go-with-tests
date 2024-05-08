package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("David", "English")
		want := "Hello, David"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Pierre", "French")
		want := "Bonjour, Pierre"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
NOTE:
   -Tests need to be in a file appending with "_test.go"
   -Test functions must start with word "Test"
   -Test function take one argument only "t *testing.T"
   -Testing needs to be imported similar to "fmt"
   -T in testing.T is the "hook"
   -can use T.Fail() when you want the end result as a fail
   t.Errorf with print a message on failed test
		"f" allows you to format your error message
		"%q" used as placeholder for variables
		for test %q is helpful as it wraps your values in double quotes ""
   testing.TB satisfies both testing t and testing b
		tells test code to fail if needed
   t.Helper() tells the test suite that this method is a test helper
   function name start with a lowercase letter
   public functions start with a capital letter
   constants can be grouped in a block instead of individually called
*/
