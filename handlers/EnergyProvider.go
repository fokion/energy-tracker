package handlers

import (
	"energy-tracker/model"
)

type ElectricityProvider interface {
	FetchElectricity() *model.Consumption
}

type GasProvider interface {
	FetchGas() *model.Consumption
}
