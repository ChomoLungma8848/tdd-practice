package main

type integerClosedRange struct {
	lower int
	upper int
}

func (i integerClosedRange) Lower() int {
	return 1
}

func (i integerClosedRange) Upper() int {
	return i.upper
}

func main() {
}

func NewIntegerClosedRange(lower int, upper int) integerClosedRange {
	return integerClosedRange{lower: lower, upper: upper}
}
