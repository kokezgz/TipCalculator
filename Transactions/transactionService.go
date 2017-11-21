package Transactions

import (
	"../Services"
	"../Utils"
)

type ITransactionService interface {
	GetTransactions() ([]Transaction, error)
	GetTransactionsBySku(sku string) ([]Transaction, error)
	injectTrans()
}

type TransactionService struct {
	client Services.IclientService
	backup Utils.IBackup
	config Utils.Config
}

func (t *TransactionService) GetTransactions() ([]Transaction, error) {
	t.injectTrans()
	t.backup.SetFileSection(t.config.Backup.Dir, t.config.Backup.File)

	response, err := t.client.CallService(t.config.Service.Route, t.config.Service.EndPoint, t.config.Service.ContentType)

	if err != nil {
		response, err = t.backup.ReadBackup()

		if err != nil {
			return nil, err
		}
	} else {
		err = t.backup.WriteBackup(response)
		if err != nil {
			println(err.Error())
		}
	}

	transactions := parseTransaction(response)
	return transactions, err
}

func (t *TransactionService) GetTransactionsBySku(sku string) ([]Transaction, error) {
	t.injectTrans()
	t.backup.SetFileSection(t.config.Backup.Dir, t.config.Backup.File)

	responseF, err := t.client.CallService(t.config.Service.Route, t.config.Service.EndPoint, t.config.Service.ContentType)

	if err != nil {
		responseF, err = t.backup.ReadBackup()

		if err != nil {
			return nil, err
		}
	}

	transactions := parseTransaction(responseF)
	transactionsBySku := findBySku(sku, transactions)

	if len(transactionsBySku) == 0 {
		response, err := t.backup.ReadBackup()

		if err != nil {
			return nil, err
		}

		transactions = parseTransaction(response)
		transactionsBySku = findBySku(sku, transactions)
	}

	err = t.backup.WriteBackup(responseF)
	if err != nil {
		println(err.Error())
	}

	return transactionsBySku, err
}

func (t *TransactionService) injectTrans() {
	//Injects
	var inj Services.ClientService
	var inj2 Utils.Backup
	t.client = &inj
	t.backup = &inj2

	//New Config
	inj3 := Utils.NewSettings("Transactions")
	t.config = inj3
}
