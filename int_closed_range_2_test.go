package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Lowerメソッドは下端点を返せる(t *testing.T) {
	icRange := IntClosedRange{3, 7}
	assert.Equal(t, 3, icRange.Lower())
}

func Test_Upperメソッドは上端点を返せる(t *testing.T) {
	icRange := IntClosedRange{3, 7}
	assert.Equal(t, 7, icRange.Upper())
}

func Test_Notationメソッドは文字列表記を返せる(t *testing.T) {
	icRange := IntClosedRange{3, 7}
	assert.Equal(t, "[3,7]", icRange.Notation())
}

func Test_Includesメソッドは指定した整数を含むか判定できる(t *testing.T) {
	tests := map[string]struct {
		input    int
		includes bool
	}{
		"9は含まれない":   {9, false},
		"5は含まれる":    {5, true},
		"1は含まれない":   {1, false},
		"下端点3は含まれる": {3, true},
		"上端点7は含まれる": {7, true},
	}

	icRange := IntClosedRange{3, 7}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.includes, icRange.Includes(test.input))
		})
	}
}
