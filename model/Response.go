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
type OctopusRatesResponse struct {
	Count    int64  `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		ValueWithoutVat   float64 `json:"value_exc_vat"`
		ValueIncludingVat float64 `json:"value_inc_vat"`
		Start             string  `json:"valid_from"`
		End               string  `json:"valid_to"`
	} `json:"results"`
}
type OctopusProductResponse struct {
	Count    int64  `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Code         string `json:"code"`
		Direction    string `json:"direction"`
		FullName     string `json:"full_name"`
		DisplayName  string `json:"display_name"`
		Description  string `json:"description"`
		IsVariable   bool   `json:"is_variable"`
		IsGreen      bool   `json:"is_green"`
		IsTracker    bool   `json:"is_tracker"`
		IsPrepay     bool   `json:"is_prepay"`
		IsBusiness   bool   `json:"is_business"`
		IsRestricted bool   `json:"is_restricted"`
		Term         int    `json:"term"`
		From         string `json:"available_from"`
		To           string `json:"available_to"`
		Links        []struct {
			Href string `json:"href"`
		}
		Brand string `json:"brand"`
	}
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
