package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type intClosedRangeSuite struct {
	suite.Suite
	icRange IntClosedRange
}

func (suite *intClosedRangeSuite) SetupSuite() {
	suite.icRange = IntClosedRange{3, 7}
}

func (suite *intClosedRangeSuite) Test_整数閉区間は下端点を持つ() {
	suite.Equal(3, suite.icRange.Lower())
}

func (suite *intClosedRangeSuite) Test_整数閉区間は上端点を持つ() {
	suite.Equal(7, suite.icRange.Upper())
}

func (suite *intClosedRangeSuite) Test_整数閉区間の文字列表記を返せる() {
	suite.Equal("[3,7]", suite.icRange.Notation())
}

func (suite *intClosedRangeSuite) Test_整数閉区間は指定した整数を含むか判定できる() {
	tests := map[string]struct {
		input    int
		includes bool
	}{
		"5は含まれる":    {5, true},
		"1は含まれない":   {1, false},
		"下端点3は含まれる": {3, true},
		"上端点7は含まれる": {7, true},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.Equal(test.includes, suite.icRange.Includes(test.input))
		})
	}
}

func Test_整数閉区間を表す(t *testing.T) {
	suite.Run(t, new(intClosedRangeSuite))
}
