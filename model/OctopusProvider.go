package model

import (
	"energy-tracker/utils"
	"sort"
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

	sort.Slice(datapoints, func(i int, j int) bool {
		return datapoints[i].Timestamp < datapoints[j].Timestamp
	})
	//drop decimals
	var daysCharged = int((datapoints[len(datapoints)-1].Timestamp - datapoints[0].Timestamp) / (24 * 3600 * 1000))
	var sum = 0.0
	for _, datapoint := range datapoints {
		sum += datapoint.Consumption
	}
	//the prices are in pences so we need to divide by 100
	return sum*(calculator.UnitPrice/100) + (calculator.StandingCharge/100)*float64(daysCharged)
}
