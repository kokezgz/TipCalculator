package Tips

import (
	"../Transactions"
)

type Tip struct {
	Transaction Transactions.Transaction
	Tip         float64 `json:"rate,string"`
}
