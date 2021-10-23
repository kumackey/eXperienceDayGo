package main

type IntClosedRange struct {
	min int
	max int
}

func (r IntClosedRange) Lower() int {
	return r.min
}

func (r IntClosedRange) Upper() int {
	return r.max
}

func (r IntClosedRange) Notation() string {
	return "[3,7]"
}
