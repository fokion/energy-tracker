package model

import (
	"energy-tracker/handlers"
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
			handlers.GetEndpoint(config["gas_endpoint"])(
				EnergyType{
					AccountNumber: config["gas_account_number"],
					MeterNumber:   config["gas_serial_number"],
				}),
			handler,
		),
		Electricity: GetOctopusApiHandler(
			GetBasicAuthProvider(config["api_key"], ""),
			handlers.GetEndpoint(config["electricity_endpoint"])(
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
