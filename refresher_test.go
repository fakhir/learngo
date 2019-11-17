package main

import (
	"testing"
)

// TestSwap tests the swap() function.
func TestSwap(t *testing.T) {
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
