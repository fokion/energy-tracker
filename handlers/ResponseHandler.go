package handlers

import (
	"energy-tracker/model"
	"energy-tracker/utils"
)

type EnergyProvider interface {
	FetchElectricity() *model.Consumption
	FetchGas() *model.Consumption
}

type PagedResponse interface {
	GetNext() *model.Consumption
}

type OctopusProvider struct {
	Gas         *model.EnergyType
	Electricity *model.EnergyType
}

func NewOptopusProvider(config utils.Config) *OctopusProvider {
	return &OctopusProvider{
		Gas: &model.EnergyType{
			AccountNumber: config["gas_account_number"],
			MeterNumber:   config["gas_serial_number"],
		},
		Electricity: &model.EnergyType{
			AccountNumber: config["electricity_account_number"],
			MeterNumber:   config["electricity_serial_number"],
		},
	}
}

func (provider *OctopusProvider) FetchElectricity() *model.Consumption {
	return nil
}
