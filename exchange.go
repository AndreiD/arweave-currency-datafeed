package main

import (
	"arweave-datafeed/utils"
	"fmt"
	"github.com/buger/jsonparser"
)

const baseURL = "https://openexchangerates.org/api/latest.json"

func getExchangeRates() ([]byte, error) {

	resp, statusCode, err := utils.GetRequest(baseURL + "?app_id=" + configuration.Get("appID") + "&show_alternative=1&prettyprint=0")
	if err != nil {
		return nil, err
	}

	if statusCode != 200 {
		return nil, fmt.Errorf("openexchangerates.org did not return 200")
	}

	// extract just the rates
	rates, _, _, err := jsonparser.Get(resp, "rates")
	if err != nil {
		return nil, err
	}

	return rates, nil
}
