package model

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestGetEndpoint_with_both(t *testing.T) {
	url := "https://example.com/{account_number}/{serial_number}"
	energyType := EnergyType{
		AccountNumber: "1234",
		MeterNumber:   "5678",
	}
	assert.Equal(t, GetEndpoint(url)(energyType), "https://example.com/1234/5678")
}
func TestGetEndpoint_with_multiple(t *testing.T) {
	url := "https://example.com/{account_number}/{serial_number}/test/{account_number}"
	energyType := EnergyType{
		AccountNumber: "1234",
		MeterNumber:   "5678",
	}
	assert.Equal(t, GetEndpoint(url)(energyType), "https://example.com/1234/5678/test/1234")
}

func TestOctopusHandler_Convert_and_get_cost(t *testing.T) {

	startTime := time.Date(2022, 05, 24, 0, 0, 0, 0, time.Local)
	endTime := time.Date(2022, 06, 24, 0, 0, 0, 0, time.Local)
	dummyBody, _ := json.Marshal(OctopusResponse{Count: 1, Results: []struct {
		Consumption float64 `json:"Consumption"`
		Start       string  `json:"interval_start"`
		End         string  `json:"interval_end"`
	}{{
		Consumption: 286.7,
		Start:       startTime.Format(time.RFC3339),
		End:         endTime.Format(time.RFC3339),
	}}})

	handler := OctopusHandler{
		Start:                 startTime,
		End:                   endTime,
		ElectricityCalculator: EnergyCalculator{UnitPrice: 21.32, StandingCharge: 23.68},
	}

	response := http.Response{Body: io.NopCloser(strings.NewReader(string(dummyBody))), StatusCode: 200}

	var consumption *Consumption
	consumption, _, err := handler.Convert(&response, consumption)
	if err != nil {
		assert.Fail(t, "Had an issue during conversion")
	}
	assert.Equal(t, 1, len(consumption.Points))

}

func TestEnergyCalculator_GetCost(t *testing.T) {
	startTime := time.Date(2022, 05, 24, 0, 0, 0, 0, time.Local)
	endTime := time.Date(2022, 06, 24, 0, 0, 0, 0, time.Local)

	points := []DataPoint{
		{Timestamp: startTime.UnixMilli()},
		{Timestamp: endTime.UnixMilli(), Consumption: 286.7},
	}

	handler := OctopusHandler{
		Start:                 startTime,
		End:                   endTime,
		ElectricityCalculator: EnergyCalculator{UnitPrice: 21.32, StandingCharge: 23.68},
	}
	assert.Equal(t, fmt.Sprintf("%.1f", 68.46), fmt.Sprintf("%.1f", handler.ElectricityCalculator.GetCost(points)))
}
