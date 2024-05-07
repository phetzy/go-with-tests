package integers

import (
	"fmt"
	"testing"
)

func TestIntegers(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

/*
	--'%d' is used to format integers instead of %q for strings
	--Testable functions begin with Example
		(much like test functions begin with Test),
		and reside in a package's _test.go files



*/
