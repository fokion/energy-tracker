package model

type OctopusResponse struct {
	Count    int64  `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Consumption float64 `json:"consumption"`
		Start       string  `json:"interval_start"`
		End         string  `json:"interval_end"`
	} `json:"results"`
}

type Consumption struct {
	points []DataPoint
	start  int64
	end    int64
}

type DataPoint struct {
	timestamp   int64
	consumption float64
}
