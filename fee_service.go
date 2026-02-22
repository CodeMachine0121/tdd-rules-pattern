package main

type FeeService struct {
	rules []FeeRules
}

type CalculateFeeParams struct {
	Region   string
	Distance float64
	Weight   float64
}

func (service FeeService) Calculate(params CalculateFeeParams) float64 {

	for _, rule := range service.rules {
		if rule.Region == params.Region &&
			rule.DistanceRange.Contains(params.Distance) &&
			rule.WeightRange.Contains(params.Weight) {
			return rule.Fee
		}
	}

	return 60.0

}

func NewFeeService() FeeService {
	return FeeService{
		rules: []FeeRules{
			{Region: "Local", DistanceRange: DistanceRange{Min: 0, Max: 5}, WeightRange: WeightRange{Min: 0, Max: 1}, Fee: 60.0},
			{Region: "Local", DistanceRange: DistanceRange{Min: 0, Max: 5}, WeightRange: WeightRange{Min: 1, Max: 3}, Fee: 90.0},
			{Region: "Local", DistanceRange: DistanceRange{Min: 5, Max: 20}, WeightRange: WeightRange{Min: 0, Max: 1}, Fee: 80.0},
		},
	}
}

type DistanceRange struct {
	Min float64
	Max float64
}

func (d DistanceRange) Contains(value float64) bool {
	return value >= d.Min && value <= d.Max
}

type WeightRange struct {
	Min float64
	Max float64
}

func (w WeightRange) Contains(value float64) bool {
	return value >= w.Min && value <= w.Max
}

type FeeRules struct {
	Region        string
	DistanceRange DistanceRange
	WeightRange   WeightRange
	Fee           float64
}
