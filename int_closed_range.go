package main

import "strconv"

type IntClosedRange struct {
	lower int
	upper int
}

func (r IntClosedRange) Lower() int {
	return r.lower
}

func (r IntClosedRange) Upper() int {
	return r.upper
}

func (r IntClosedRange) Notation() string {
	// 文字列と数字の結合はgoでは地道な方法しかなかった
	return "[" + strconv.Itoa(r.lower) + "," + strconv.Itoa(r.upper) + "]"
}

func (r IntClosedRange) Includes(i int) bool {
	return r.lower <= i && i <= r.upper
}
