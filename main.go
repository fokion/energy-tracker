package main

import (
	"energy-tracker/model"
	"energy-tracker/utils"
	"flag"
	"fmt"
)

var (
	fileName string
)

func main() {
	flag.StringVar(&fileName, "f", "", "Specify configuration")
	flag.Parse()
	properties, err := utils.ReadProperties(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(properties)

	provider := model.NewOptopusProvider(properties)
	fmt.Println(provider.FetchElectricity())

}
