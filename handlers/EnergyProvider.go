package handlers

import (
	"energy-tracker/model"
	"strings"
)

type ElectricityProvider interface {
	FetchElectricity() *model.Consumption
}

type GasProvider interface {
	FetchGas() *model.Consumption
}

func GetEndpoint(url string) func(energyType model.EnergyType) string {
	return func(energyType model.EnergyType) string {
		url = strings.ReplaceAll(url, "{account_number}", energyType.AccountNumber)
		url = strings.ReplaceAll(url, "{serial_number}", energyType.MeterNumber)
		return url
	}
}
