package main

import "errors"

type integerClosedRange struct {
	lower int
	upper int
}

func (i integerClosedRange) Lower() int {
	return i.lower
}

func (i integerClosedRange) Upper() int {
	return i.upper
}

func main() {
}

func NewIntegerClosedRange(lower int, upper int) (integerClosedRange, error) {
	if lower >= upper {
		return integerClosedRange{}, errors.New("invalid arguments")
	}

	return integerClosedRange{lower: lower, upper: upper}, nil
}
