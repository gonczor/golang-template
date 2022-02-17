package main

import (
	a "app/algorithms"
	"testing"
)

func TestFibonacci(t *testing.T) {
	var result int = a.Fibonacci(10)
	var expected = 55
	if result != expected {
		t.Errorf("Fibonacci result wrong. Expected: %d, got: %d", expected, result)
	}
}
