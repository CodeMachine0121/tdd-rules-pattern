package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FeeServiceSuite struct {
	FeeService
	suite.Suite
}

func (s *FeeServiceSuite) SetupTest() {
	s.FeeService = NewFeeService()
}

func (s *FeeServiceSuite) TestFee_Local_0to5km_0to1k_Should_Be_60() {

	fee := s.FeeService.Calculate(CalculateFeeParams{"Local", 3.0, 0.8})
	s.Equal(60.0, fee)
}

func (s *FeeServiceSuite) TestFee_Local_0to5km_1to3k_Should_Be_90() {

	fee := s.FeeService.Calculate(CalculateFeeParams{"Local", 2.0, 2.2})
	s.Equal(90.0, fee)
}

func (s *FeeServiceSuite) TestFee_Local_5to20km_0to1k_Should_Be_80() {

	fee := s.FeeService.Calculate(CalculateFeeParams{"Local", 6.0, 1.0})
	s.Equal(80.0, fee)
}

func TestFeeService(t *testing.T) {
	suite.Run(t, new(FeeServiceSuite))
}
