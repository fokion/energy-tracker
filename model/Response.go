package model

type OctopusResponse struct {
	count    int64
	next     string
	previous string
	results  []OctopusDataPoint
}

type OctopusDataPoint struct {
	consumption float64
	start       string `json:"interval_start"`
	end         string `json:"interval_end"`
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
