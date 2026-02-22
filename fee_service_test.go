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
	fee := s.FeeService.Calculate(CalculateFeeParams{"Local", 0.0, 0.0})
	s.Equal(60.0, fee)
	fee = s.FeeService.Calculate(CalculateFeeParams{"Local", 4.9, 0.9})
	s.Equal(60.0, fee)
}

func (s *FeeServiceSuite) TestFee_Local_0to5km_1to3k_Should_Be_90() {
	fee := s.FeeService.Calculate(CalculateFeeParams{"Local", 2.0, 1.0})
	s.Equal(90.0, fee)
	fee = s.FeeService.Calculate(CalculateFeeParams{"Local", 2.0, 2.9})
	s.Equal(90.0, fee)
}

func (s *FeeServiceSuite) TestFee_Local_5to20km_0to1k_Should_Be_80() {
	fee := s.FeeService.Calculate(CalculateFeeParams{"Local", 5.0, 0.5})
	s.Equal(80.0, fee)
	fee = s.FeeService.Calculate(CalculateFeeParams{"Local", 19.9, 0.5})
	s.Equal(80.0, fee)
}

func (s *FeeServiceSuite) TestFee_Local_5to20km_1to3k_Should_Be_120() {
	fee := s.FeeService.Calculate(CalculateFeeParams{"Local", 10.0, 2.0})
	s.Equal(120.0, fee)
}

func (s *FeeServiceSuite) TestFee_Local_20to999km_0to1k_Should_Be_150() {
	fee := s.FeeService.Calculate(CalculateFeeParams{"Local", 25.0, 0.5})
	s.Equal(150.0, fee)
}

func (s *FeeServiceSuite) TestFee_Local_20to999km_1to3k_Should_Be_210() {
	fee := s.FeeService.Calculate(CalculateFeeParams{"Local", 30.0, 2.5})
	s.Equal(210.0, fee)
}

func (s *FeeServiceSuite) TestFee_Remote_0to5km_0to1k_Should_Be_100() {
	fee := s.FeeService.Calculate(CalculateFeeParams{"Remote", 2.0, 0.5})
	s.Equal(100.0, fee)
}

func (s *FeeServiceSuite) TestFee_Remote_0to5km_1to3k_Should_Be_140() {
	fee := s.FeeService.Calculate(CalculateFeeParams{"Remote", 1.0, 1.5})
	s.Equal(140.0, fee)
}

func (s *FeeServiceSuite) TestFee_Remote_5to20km_0to1k_Should_Be_130() {
	fee := s.FeeService.Calculate(CalculateFeeParams{"Remote", 10.0, 0.8})
	s.Equal(130.0, fee)
}

func (s *FeeServiceSuite) TestFee_Remote_5to20km_1to3k_Should_Be_180() {
	fee := s.FeeService.Calculate(CalculateFeeParams{"Remote", 15.0, 2.0})
	s.Equal(180.0, fee)
}

func (s *FeeServiceSuite) TestFee_Remote_20to999km_0to1k_Should_Be_220() {
	fee := s.FeeService.Calculate(CalculateFeeParams{"Remote", 50.0, 0.5})
	s.Equal(220.0, fee)
}

func (s *FeeServiceSuite) TestFee_Remote_20to999km_1to3k_Should_Be_300() {
	fee := s.FeeService.Calculate(CalculateFeeParams{"Remote", 100.0, 2.8})
	s.Equal(300.0, fee)
}

func TestFeeService(t *testing.T) {
	suite.Run(t, new(FeeServiceSuite))
}
