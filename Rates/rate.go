package Rates

type Rate struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate,string"`
}
