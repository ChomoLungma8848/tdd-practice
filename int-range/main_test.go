package main

import (
	"testing"
)

func TestUpperEndpoint(t *testing.T) {
	integerClosedRange := NewIntegerClosedRange(1, 2)
	got := integerClosedRange.Upper()
	expected := 2
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}

	integerClosedRange = NewIntegerClosedRange(1, 5)
	got = integerClosedRange.Upper()
	expected = 5
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}

	integerClosedRange = NewIntegerClosedRange(1, 100)
	got = integerClosedRange.Upper()
	expected = 100
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}

func TestLowerEndpoint(t *testing.T) {
	integerClosedRange := NewIntegerClosedRange(1, 5)
	got := integerClosedRange.Lower()
	expected := 1
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}

}
