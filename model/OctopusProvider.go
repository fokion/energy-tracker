package model

import (
	"energy-tracker/utils"
)

type OctopusProvider struct {
	Gas         APIHandler
	Electricity APIHandler
}

func NewOctopusProvider(config utils.Config, handler *OctopusHandler) *OctopusProvider {

	return &OctopusProvider{
		Gas: GetOctopusApiHandler(
			GetBasicAuthProvider(config["api_key"], ""),
			GetEndpoint(config["gas_endpoint"])(
				EnergyType{
					AccountNumber: config["gas_account_number"],
					MeterNumber:   config["gas_serial_number"],
				}),
			handler,
		),
		Electricity: GetOctopusApiHandler(
			GetBasicAuthProvider(config["api_key"], ""),
			GetEndpoint(config["electricity_endpoint"])(
				EnergyType{
					AccountNumber: config["electricity_account_number"],
					MeterNumber:   config["electricity_serial_number"],
				}),
			handler,
		),
	}
}

func (provider *OctopusProvider) FetchElectricity() *Consumption {
	if provider.Electricity != nil {
		consumption, err := provider.Electricity.GET(nil)
		if err != nil {
			return nil
		}
		return consumption
	}
	return nil
}

func (provider *OctopusProvider) FetchGas() *Consumption {
	if provider.Gas != nil {
		consumption, err := provider.Gas.GET(nil)
		if err != nil {
			return nil
		}
		return consumption
	}
	return nil
}

type EnergyCalculator struct {
	UnitPrice      float64
	StandingCharge float64
}

func (calculator *EnergyCalculator) Calculate(from float64, to float64, price float64) float64 {
	return (to - from) * price
}

type DatapointCalculator interface {
	GetCost(datapoints []DataPoint) float64
}

func (calculator *EnergyCalculator) GetCost(datapoints []DataPoint) float64 {
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
