package model

import (
	"energy-tracker/utils"
)

type OctopusProvider struct {
	Gas         APIHandler
	Electricity APIHandler
}

func NewOptopusProvider(config utils.Config) *OctopusProvider {

	return &OctopusProvider{
		Gas: GetOctopusApiHandler(
			GetBasicAuthProvider(config["api_key"], ""),
			"https://api.octopus.energy/v1/gas-meter-points/{account_number}/meters/{serial_number}/consumption/",
			EnergyType{
				AccountNumber: config["gas_account_number"],
				MeterNumber:   config["gas_serial_number"],
			},
		),
		Electricity: GetOctopusApiHandler(
			GetBasicAuthProvider(config["api_key"], ""),
			"https://api.octopus.energy/v1/electricity-meter-points/{account_number}/meters/{serial_number}/consumption/",
			EnergyType{
				AccountNumber: config["electricity_account_number"],
				MeterNumber:   config["electricity_serial_number"],
			},
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
