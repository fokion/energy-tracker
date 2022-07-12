package model

import (
	"encoding/json"
	"energy-tracker/handlers"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type ConnectionProvider interface {
	Call(method string, url string, body io.Reader) (*http.Response, error)
}

type BasicAuth struct {
	username string
	password string
}

func (provider *BasicAuth) Call(method string, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(provider.username, provider.password)
	return client.Do(req)
}

func GetBasicAuthProvider(username string, password string) *BasicAuth {
	return &BasicAuth{username: username, password: password}
}

type APIHandler interface {
	GET(body io.Reader) (*Consumption, error)
}

type PagedAPI struct {
	Provider ConnectionProvider
	Handler  ConsumptionHandler
	Url      string
}
type ConsumptionHandler interface {
	Convert(response *http.Response, consumption *Consumption) (*Consumption, *string, error)
}

type OctopusHandler struct {
	Start                 int64
	End                   int64
	GasCalculator         handlers.EnergyCalculator
	ElectricityCalculator handlers.EnergyCalculator
}

func (handler *OctopusHandler) Convert(response *http.Response, consumption *Consumption) (*Consumption, *string, error) {
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, nil, errors.New("could not read the body response")
	}
	var resp OctopusResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, nil, errors.New("could not parse the response")
	}
	points := []DataPoint{}
	var start int64
	var end int64

	for _, octopusDataPoint := range resp.Results {
		timestamp, err := time.Parse(time.RFC3339, octopusDataPoint.End)
		if err != nil {
			fmt.Println(fmt.Sprintf("could not parse %s", octopusDataPoint.End))
		} else {
			datapoint := DataPoint{Consumption: octopusDataPoint.Consumption, Timestamp: timestamp.UnixMilli()}
			points = append(points, datapoint)
		}
	}
	if consumption == nil {
		return &Consumption{Points: points, Start: start, End: end}, &resp.Next, nil
	}
	consumption.End = end
	consumption.Points = append(consumption.Points, points...)
	return consumption, &resp.Next, nil
}
func (apiHandler *PagedAPI) GET(body io.Reader) (*Consumption, error) {
	provider := apiHandler.Provider
	url := apiHandler.Url

	var consumption *Consumption
	for url != "" {
		resp, err := provider.Call("GET", url, nil)
		if err != nil {
			return nil, err
		}
		updatedConsumption, nextUrl, err := apiHandler.Handler.Convert(resp, consumption)
		if err != nil {
			return nil, err
		}
		url = ""
		if nextUrl != nil {
			url = *nextUrl
			fmt.Println(url)
		}
		consumption = updatedConsumption
	}
	return consumption, nil
}
func GetOctopusApiHandler(provider ConnectionProvider, url string, handler *OctopusHandler) *PagedAPI {
	return &PagedAPI{Provider: provider, Url: url, Handler: handler}
}
