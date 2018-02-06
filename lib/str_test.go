package lib

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	r := Fibonacci(10)
	if r != 55 {
		t.Errorf("Fibonacci(10) failed. Got %d, expected 55.", r)
	}
}