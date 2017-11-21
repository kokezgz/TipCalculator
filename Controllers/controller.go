package Controllers

import (
	"fmt"
	"net/http"
	"time"

	"../Rates"
	"../Tips"
	"../Transactions"
	"github.com/gorilla/mux"
)

type Controller struct {
	rateService        Rates.IRateService
	transactionService Transactions.ITransactionService
	tipService         Tips.ITipService
}

func (c *Controller) StartServer() {
	c.injects()

	r := mux.NewRouter()
	r.HandleFunc("/api/rates", c.handlerRates).Methods("GET")
	r.HandleFunc("/api/transactions", c.handlerTransactions).Methods("GET")
	r.HandleFunc("/api/{sku}/{currency}", c.handlerTips).Methods("GET")

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("The Server Start in port 8100")
	srv.ListenAndServe()
}

func (c *Controller) injects() {
	var inj Rates.RateService
	var inj2 Transactions.TransactionService
	var inj3 Tips.TipService

	c.rateService = &inj
	c.transactionService = &inj2
	c.tipService = &inj3
}
