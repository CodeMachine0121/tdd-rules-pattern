package main

type DistanceRange struct {
	Min float64
	Max float64
}

func (d DistanceRange) Contains(value float64) bool {
	return value >= d.Min && value < d.Max
}

type WeightRange struct {
	Min float64
	Max float64
}

func (w WeightRange) Contains(value float64) bool {
	return value >= w.Min && value < w.Max
}

type FeeRules struct {
	Region        string
	DistanceRange DistanceRange
	WeightRange   WeightRange
	Fee           float64
}
