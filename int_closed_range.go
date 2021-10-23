package main

type IntClosedRange struct {
	min int
	max int
}

func (r IntClosedRange) lower() interface{} {
	return r.min
}

func (r IntClosedRange) upper() interface{} {
	return r.max
}
