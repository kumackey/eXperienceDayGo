package main

import "strconv"

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
	return "[" + strconv.Itoa(r.min) + "," + strconv.Itoa(r.max) + "]"
}

func (r IntClosedRange) Includes(i int) bool {
	return r.min <= i && i <= r.max
}
