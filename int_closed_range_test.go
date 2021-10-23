package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntClosedRange(t *testing.T) {
	var icRange IntClosedRange
	assert.Equal(t, 3, icRange.lower())
}
