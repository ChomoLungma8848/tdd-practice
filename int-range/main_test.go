package main

import (
	"testing"
)

func TestUpperEndpoint(t *testing.T) {
	integerClosedRange, _ := NewIntegerClosedRange(1, 2)
	got := integerClosedRange.Upper()
	expected := 2
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}

	integerClosedRange, _ = NewIntegerClosedRange(1, 5)
	got = integerClosedRange.Upper()
	expected = 5
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}

	integerClosedRange, _ = NewIntegerClosedRange(1, 100)
	got = integerClosedRange.Upper()
	expected = 100
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}

func TestLowerEndpoint(t *testing.T) {
	integerClosedRange, _ := NewIntegerClosedRange(1, 5)
	got := integerClosedRange.Lower()
	expected := 1
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}

	integerClosedRange, _ = NewIntegerClosedRange(3, 5)
	got = integerClosedRange.Lower()
	expected = 3
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}

	integerClosedRange, _ = NewIntegerClosedRange(4, 5)
	got = integerClosedRange.Lower()
	expected = 4
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}

func TestValidRange(t *testing.T) {
	_, err := NewIntegerClosedRange(1, 5)
	if err != nil {
		t.Errorf("Error must be nil")
	}

	_, err = NewIntegerClosedRange(5, 1)
	if err == nil {
		t.Errorf("Error must not be nil")
	}
}
