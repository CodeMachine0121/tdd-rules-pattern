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
			{Region: "Local", DistanceRange: DistanceRange{Min: 5, Max: 20}, WeightRange: WeightRange{Min: 1, Max: 3}, Fee: 120.0},
			{Region: "Local", DistanceRange: DistanceRange{Min: 20, Max: 999}, WeightRange: WeightRange{Min: 0, Max: 1}, Fee: 150.0},
			{Region: "Local", DistanceRange: DistanceRange{Min: 20, Max: 999}, WeightRange: WeightRange{Min: 1, Max: 3}, Fee: 210.0},
			{Region: "Remote", DistanceRange: DistanceRange{Min: 0, Max: 5}, WeightRange: WeightRange{Min: 0, Max: 1}, Fee: 100.0},
			{Region: "Remote", DistanceRange: DistanceRange{Min: 0, Max: 5}, WeightRange: WeightRange{Min: 1, Max: 3}, Fee: 140.0},
			{Region: "Remote", DistanceRange: DistanceRange{Min: 5, Max: 20}, WeightRange: WeightRange{Min: 0, Max: 1}, Fee: 130.0},
			{Region: "Remote", DistanceRange: DistanceRange{Min: 5, Max: 20}, WeightRange: WeightRange{Min: 1, Max: 3}, Fee: 180.0},
			{Region: "Remote", DistanceRange: DistanceRange{Min: 20, Max: 999}, WeightRange: WeightRange{Min: 0, Max: 1}, Fee: 220.0},
			{Region: "Remote", DistanceRange: DistanceRange{Min: 20, Max: 999}, WeightRange: WeightRange{Min: 1, Max: 3}, Fee: 300.0},
		},
	}
}
