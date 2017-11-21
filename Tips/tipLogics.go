package Tips

import (
	"math"
	"strings"

	"../Rates"
	"../Transactions"
)

func calculate(transactions []Transactions.Transaction) []Tip {

	var tips []Tip
	for _, s := range transactions {
		n := s.Amount * 0.05
		tip := Tip{Transaction: s, Tip: n}
		tips = append(tips, tip)
	}

	return tips
}

func currencyConvert(currency string, tips []Tip, rates []Rates.Rate) Tip {

	tipConverted := Tip{}
	for _, s := range tips {
		if strings.EqualFold(s.Transaction.Currency, currency) {
			tipConverted.Transaction.Amount += s.Transaction.Amount
			tipConverted.Tip += s.Tip
		} else {
			tip := findDirectConvert(currency, s, rates)
			if tip.Tip == 0 {
				tip = findIndirectConvert(currency, s, rates)
			}

			tipConverted.Transaction.Amount += tip.Transaction.Amount
			tipConverted.Tip += tip.Tip
		}
	}

	tipConverted.Transaction.Sku = "TOTAL"
	tipConverted.Transaction.Currency = currency
	return tipConverted
}

func findDirectConvert(currency string, tip Tip, rates []Rates.Rate) Tip {

	for _, s := range rates {
		if s.From == tip.Transaction.Currency && s.To == currency {
			trans := Transactions.Transaction{Amount: tip.Transaction.Amount * s.Rate}
			tip := Tip{Tip: tip.Tip * s.Rate, Transaction: trans}
			return tip
		}
	}

	return Tip{}
}

func findIndirectConvert(currency string, tip Tip, rates []Rates.Rate) Tip {

	var ratesFrom []Rates.Rate
	for _, s := range rates {
		if s.From == tip.Transaction.Currency {
			ratesFrom = append(ratesFrom, s)
		}
	}

	for _, x := range ratesFrom {
		for _, z := range rates {
			if z.From == x.To && z.To == currency {
				trans := Transactions.Transaction{Amount: tip.Transaction.Amount * x.Rate}
				trans.Amount = trans.Amount * z.Rate
				tip := Tip{Tip: tip.Tip * x.Rate, Transaction: trans}
				tip.Tip = tip.Tip * z.Rate
				return tip
			}
		}
	}

	return Tip{}
}

func roundAllTips(tips []Tip) {
	for i := 0; i < len(tips); i++ {
		tips[i].Tip = round(tips[i].Tip, 0.5, 2)
		tips[i].Transaction.Amount = round(tips[i].Transaction.Amount, 0.5, 2)
	}
}

func round(val float64, roundOn float64, places int) (newVal float64) {

	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
