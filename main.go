package main

import (
	"energy-tracker/model"
	"energy-tracker/utils"
	"flag"
	"fmt"
	"strconv"
	"time"
)

var (
	fileName string
	daysStr  string
)

func main() {
	flag.StringVar(&fileName, "f", "", "Specify configuration")
	flag.StringVar(&daysStr, "days", "10", "Get the last x number of days of data ( default 10 )")
	flag.Parse()
	properties, err := utils.ReadProperties(fileName)
	if err != nil {
		panic(err)
	}

	var days int64
	days, err = strconv.ParseInt(daysStr, 10, 64)

	if err != nil {
		days = 10
	}

	gasUnitPrice, _ := strconv.ParseFloat(properties["price_gas"], 64)
	gasStandingPrice, _ := strconv.ParseFloat(properties["standing_gas"], 64)
	electricityUnitPrice, _ := strconv.ParseFloat(properties["price_electricity"], 64)
	electricityStandingPrice, _ := strconv.ParseFloat(properties["standing_electricity"], 64)

	handler := &model.OctopusHandler{
		Start:                 time.Now().Add(time.Duration(-days*24) * time.Hour),
		End:                   time.Now(),
		GasCalculator:         model.EnergyCalculator{UnitPrice: gasUnitPrice, StandingCharge: gasStandingPrice},
		ElectricityCalculator: model.EnergyCalculator{UnitPrice: electricityUnitPrice, StandingCharge: electricityStandingPrice},
	}
	provider := model.NewOctopusProvider(properties, handler)
	electricalConsumption := provider.FetchElectricity()

	//	fmt.Println(electricalConsumption.Points)
	fmt.Println(handler.ElectricityCalculator.GetCost(electricalConsumption.Points))
}
