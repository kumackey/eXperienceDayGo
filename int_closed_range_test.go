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

func (suite *intClosedRangeSuite) Test_整数閉区間は下端点と上端点を持つ() {
	suite.Equal(3, suite.icRange.Lower())
	suite.Equal(7, suite.icRange.Upper())
}

func (suite *intClosedRangeSuite) Test_整数閉区間の文字列表記を返せる() {
	suite.Equal("[3,7]", suite.icRange.Notation())
}

func (suite *intClosedRangeSuite) Test_整数閉区間は指定した整数を含むか判定できる() {
	suite.True(suite.icRange.Includes(5))
	suite.False(suite.icRange.Includes(1))
}

func (suite *intClosedRangeSuite) Test_整数閉区間は下端点3を含む() {
	suite.True(suite.icRange.Includes(3))
}

func (suite *intClosedRangeSuite) Test_整数閉区間は上端点7を含む() {
	suite.True(suite.icRange.Includes(7))
}

func TestIntClosedRangeSuite(t *testing.T) {
	suite.Run(t, new(intClosedRangeSuite))
}
