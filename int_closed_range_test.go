package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_整数閉区間は下端点と上端点を持つ(t *testing.T) {
	icRange := IntClosedRange{3, 7}
	assert.Equal(t, 3, icRange.Lower())
	assert.Equal(t, 7, icRange.Upper())
}

func Test_整数閉区間の文字列表記を返せる(t *testing.T) {
	icRange := IntClosedRange{3, 7}
	assert.Equal(t, "[3,7]", icRange.Notation())

	icRange = IntClosedRange{2, 8}
	assert.Equal(t, "[2,8]", icRange.Notation())
}

func Test_整数閉区間は指定した整数を含むか判定できる(t *testing.T) {
	icRange := IntClosedRange{3, 7}
	assert.True(t, icRange.Includes(5))
	assert.False(t, icRange.Includes(1))
}

func Test_整数閉区間は下端点3を含む(t *testing.T) {
	icRange := IntClosedRange{3, 7}
	assert.True(t, icRange.Includes(3))
}

func Test_整数閉区間は上端点7を含む(t *testing.T) {
	icRange := IntClosedRange{3, 7}
	assert.True(t, icRange.Includes(7))
}
