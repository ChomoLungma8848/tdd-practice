package main

import(
	"testing"
)

func TestUpperEndpoint(t *testing.T) {
	integerClosedRange := NewIntegerClosedRange(1, 2)
	if integerClosedRange.upper != 2 {
		t.Errorf( "expected 2, but got %v", integerClosedRange.upper)
	}
}
