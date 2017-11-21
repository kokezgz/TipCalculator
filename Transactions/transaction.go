package Transactions

type Transaction struct {
	Sku      string
	Currency string
	Amount   float64 `json:"amount,string"`
}
