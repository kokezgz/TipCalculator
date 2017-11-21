package Rates

import (
	"encoding/json"
)

func parseRates(value []byte) []Rate {
	var rates []Rate
	err := json.Unmarshal(value, &rates)

	if err != nil {
		return nil
	}

	return rates
}
