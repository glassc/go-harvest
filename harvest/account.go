package harvest

import (
	"encoding/json"
)

// Account related services from the harvest api
type AccountService service

type WhoAmI struct {
	Company Company `json:"company"`
}

type Company struct {
	BaseURI string `json:"base_uri"`
	FullDomain string `json:"full_domain"`
	Name string `json:"name"`
	Active bool `json:"active"`
	WeekStartDay string `json:"week_start_day"`
	TimeFormat string `json:"time_format"`
	Clock string `json:"clock"`
	DecimalSymbol string `json:"decimal_symbol"`
	ColorScheme string `json:"color_scheme"`
	Modules Modules `json:"modules"`
	ThousandsSeparator string `json:"thousands_separator"`
	PlanType string `json:"plan_type"`
}

type Modules struct {
	Expenses bool `json:"expenses"`
	Invoices bool `json:"invoices"`
	Approval bool `json:"approval"`
	Team bool `json:"team"`
}

// Gets information about the current user
// Harvest API Docs: http://help.getharvest.com/api-v1/introduction/overview/who-am-i/
func (s *AccountService) WhoAmI() (*WhoAmI, error) {
	resp, err := s.client.get("account/who_am_i")

	if err != nil {
		return nil, err
	}

	result := new(WhoAmI)
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, nil

}