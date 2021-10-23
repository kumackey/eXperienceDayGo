package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntClosedRange_LowerAndUpper(t *testing.T) {
	icRange := IntClosedRange{3, 7}
	assert.Equal(t, 3, icRange.Lower())
	assert.Equal(t, 7, icRange.Upper())
}

func TestIntClosedRange_Notaion(t *testing.T) {
	icRange := IntClosedRange{3, 7}
	assert.Equal(t, "[3,7]", icRange.Notation())

	icRange = IntClosedRange{2, 8}
	assert.Equal(t, "[2,8]", icRange.Notation())
}
