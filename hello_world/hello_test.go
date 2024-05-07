package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("David")
	want := "Hello, David"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/* Tests need to be in a file appending with "_test.go"
   Test functions must start with word "Test"
   Test function take one argument only "t *testing.T"
   Testing needs to be imported similar to "fmt"
   T in testing.T is the "hook"
   can use T.Fail() when you want the end result as a fail
   t.Errorf with print a message on failed test
		f allows you to format your error message
		%q used as placeholder for variables
		for test %q is helpful as it wraps your values in double quotes ""
*/
