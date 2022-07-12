package handlers

import (
	"energy-tracker/model"
)

type Calculator interface {
	Calculate(from float64, to float64, price float64) float64
}
type DatapointCalculator interface {
	GetCost(datapoints []model.DataPoint) float64
}
type EnergyCalculator struct {
	UnitPrice      float64
	StandingCharge float64
}

func (calculator *EnergyCalculator) Calculate(from float64, to float64, price float64) float64 {
	return (to - from) * price
}

func (calculator *EnergyCalculator) GetCost(datapoints []model.DataPoint) float64 {
	var sum = 0.0
	//sort the datapoints
	var previous = datapoints[0].Consumption
	//drop decimals
	var daysCharged = int((datapoints[len(datapoints)-1].Timestamp - datapoints[0].Timestamp) / (24 * 3600))
	for _, datapoint := range datapoints {
		sum += calculator.Calculate(previous, datapoint.Consumption, calculator.UnitPrice)
	}
	sum += calculator.StandingCharge * float64(daysCharged)

	return sum
}
