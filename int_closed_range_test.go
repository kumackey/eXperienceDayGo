package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntClosedRange(t *testing.T) {
	icRange := IntClosedRange{3, 7}
	assert.Equal(t, 3, icRange.lower())
	assert.Equal(t, 7, icRange.upper())
}
