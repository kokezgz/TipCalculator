package Transactions

import (
	"encoding/json"
)

func parseTransaction(value []byte) []Transaction {

	var transactions []Transaction
	err := json.Unmarshal(value, &transactions)

	if err != nil {
		return nil
	}

	return transactions
}

func findBySku(sku string, value []Transaction) []Transaction {

	var find []Transaction
	for _, s := range value {
		if s.Sku == sku {
			find = append(find, s)
		}
	}

	return find
}
