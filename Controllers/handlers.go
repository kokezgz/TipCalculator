package Controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) handlerRates(w http.ResponseWriter, r *http.Request) {
	rates, _ := c.rateService.GetRates()
	response, _ := json.Marshal(rates)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (c *Controller) handlerTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, _ := c.transactionService.GetTransactions()
	response, _ := json.Marshal(transactions)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (c *Controller) handlerTips(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tips, _ := c.tipService.CallService(vars["sku"], vars["currency"])
	response, _ := json.Marshal(tips)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
