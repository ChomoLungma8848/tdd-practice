package main

type integerClosedRange struct {
	lower int
	upper int
}

func main() {
}

func NewIntegerClosedRange( lower int, upper int) integerClosedRange {
	return integerClosedRange{ lower: lower, upper: upper }
}
