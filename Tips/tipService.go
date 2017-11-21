package Tips

import (
	"../Rates"
	"../Transactions"
)

type ITipService interface {
	CallService(sku string, currency string) ([]Tip, error)
}

type TipService struct {
	rateService        Rates.IRateService
	transactionService Transactions.ITransactionService
}

func (c *TipService) CallService(sku string, currency string) ([]Tip, error) {
	c.injects()
	rates, _ := c.rateService.GetRates()
	transactions, _ := c.transactionService.GetTransactionsBySku(sku)

	tips := calculate(transactions)
	tipConverted := currencyConvert(currency, tips, rates)

	tips = append([]Tip{tipConverted}, tips...)

	roundAllTips(tips)
	return tips, nil
}

func (c *TipService) injects() {
	var inj Rates.RateService
	var inj2 Transactions.TransactionService

	c.rateService = &inj
	c.transactionService = &inj2
}
