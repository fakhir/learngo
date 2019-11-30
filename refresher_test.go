package main

import (
	"fmt"
	"testing"
)

// This test can be executed by running this command (add -v for verbose output):
// $ go test

// TestSwap tests the swap() function. All test function names must begin
// with "Test".
func TestSwap(t *testing.T) {
	// The following functions can be called within the test:
	// - Fail(): Fails the test but continues next test in the file.
	// - FailNow(): Fails the test and skips all other tests in this file. The next
	//   test case is then executed.
	// - Log(args ...interface{}): Logs an informational/error message.
	// - Fatal(args ...interface{}): Equivalent to Log() followed by FailNow().

	a, b := swap("hello", "world")
	if a == "world" && b == "hello" {
		t.Log("swap is successful but we will fail it!")
		//t.Fail()
	}
}

// TestNothing does nothing.
func TestNothing(t *testing.T) {
	t.Log("Testing nothing")
	return
}

// ExampleSwap is both an example function used for auto-doc generation and
// a test function. Such functions must begin with "Example". The last line
// must be a comment describing the function's displayed output. Also note
// that Example functions do not take the usual test function parameter.
func ExampleSwap() {
	a, b := swap("one", "two")
	fmt.Println(a, b)
	// Output:
	// two one
}
