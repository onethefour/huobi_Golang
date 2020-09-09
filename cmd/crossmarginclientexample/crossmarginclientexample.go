package crossmarginclientexample

import (
	"huobi_Golang/common/config"
	"huobi_Golang/common/log"
	"huobi_Golang/pkg/client"
	"huobi_Golang/pkg/model/margin"
)

func RunAllExamples() {
	transferIn()
	transferOut()
	getMarginLoanInfo()
	marginOrders()
	marginOrdersRepay()
	marginLoanOrders()
	marginAccountsBalance()
}

//  Transfer specific asset from spot trading account to cross margin account.
func transferIn() {
	request := margin.CrossMarginTransferRequest{
		Currency: "usdt",
		Amount:   "1.0"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.TransferIn(request)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info("Data: %+v", resp)
	}
}

//  Transfer specific asset from cross margin account to spot trading account.
func transferOut() {
	request := margin.CrossMarginTransferRequest{
		Currency: "usdt",
		Amount:   "1.0"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.TransferOut(request)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info("Data: %+v", resp)
	}
}

//  Get the loan interest rates and loan quota applied on the user.
func getMarginLoanInfo() {
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetMarginLoanInfo()
	if err != nil {
		log.Error(err.Error())
	} else {
		for _, info := range resp {
			log.Info("Info: %+v", info)
		}
	}
}

//  Place an order to apply a margin loan.
func marginOrders() {
	request := margin.CrossMarginOrdersRequest{Currency: "usdt",
		Amount: "1.0"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.ApplyLoan(request)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info("Data: %+v", resp)
	}
}

//  Repays margin loan with you asset in your margin account.
func marginOrdersRepay() {
	orderId := "12345"
	request := margin.MarginOrdersRepayRequest{Amount: "1.0"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.Repay(orderId, request)
	if err != nil {
		log.Error("Repay error: %s", err)
	} else {
		log.Info("Repay successfully, id=%d", resp)
	}
}

//  Get the margin orders based on a specific searching criteria.
func marginLoanOrders() {
	optionalRequest := margin.CrossMarginLoanOrdersOptionalRequest{}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.MarginLoanOrders(optionalRequest)
	if err != nil {
		log.Error(err.Error())
	} else {
		for _, order := range resp {
			log.Info("Order: %+v", order)
		}
	}
}

//  Get the balance of the margin loan account.
func marginAccountsBalance() {
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.MarginAccountsBalance("")
	if err != nil {
		log.Error(err.Error())
	} else {
		for _, account := range resp.List {
			log.Info("Account: %+v", account)
		}
	}
}
