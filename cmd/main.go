package main

import (
	"huobi_Golang/cmd/accountclientexample"
	"huobi_Golang/cmd/accountwebsocketclientexample"
	"huobi_Golang/cmd/commonclientexample"
	"huobi_Golang/cmd/crossmarginclientexample"
	"huobi_Golang/cmd/etfclientexample"
	"huobi_Golang/cmd/isolatedmarginclientexample"
	"huobi_Golang/cmd/marketclientexample"
	"huobi_Golang/cmd/marketwebsocketclientexample"
	"huobi_Golang/cmd/orderclientexample"
	"huobi_Golang/cmd/orderwebsocketclientexample"
	"huobi_Golang/cmd/subuserclientexample"
	"huobi_Golang/cmd/walletclientexample"
	"huobi_Golang/common/config"
	"huobi_Golang/common/log"
	"huobi_Golang/common/perflogger"
	"huobi_Golang/pkg/client"
	"huobi_Golang/pkg/model/margin"
)

func main() {

	huobi := client.NewClient(config.AccessKey, config.SecretKey, config.Host)
	log.Info("%v", huobi.Accounts)

	if ticker, err := huobi.GetLast24hCandlestickAskBid("btcusdt"); err != nil {
		log.Info("Get account error: %s", err)
	} else {
		price, _ := ticker.Ask[0].Float64()
		log.Info("account:%+v", price)
		return
	}
	if amount, err := huobi.BalanceOf("usdt"); err != nil {
		log.Info("Get amount error: %s", err)
	} else {
		log.Info("usdt:%+v", amount)
	}
	if amount, err := huobi.BalanceOf("btc"); err != nil {
		log.Info("Get amount error: %s", err)
	} else {
		log.Info("btc:%+v", amount)
	}
	return
	//investment borrow
	if resp, err := huobi.IsolatedMarginClient.GetMarginLoanInfo(margin.GetMarginLoanInfoOptionalRequest{}); err != nil {
		log.Info("Get account error: %s", err)
	} else {
		log.Info("account:%+v", resp)
	}
	return
	if resp, err := huobi.CrossMarginClient.GetMarginLoanInfo(); err != nil {
		log.Info("Get account error: %s", err)
	} else {
		log.Info("account:%+v", resp)
	}
	return
	if err := huobi.GetC2CRate(huobi.Accounts["investment"]); err != nil {
		log.Info("GetC2CRate error: %s", err)
		return
	}
	return
	if resp, err := huobi.GetC2cBalance(); err != nil {
		log.Info("Get account error: %s", err)
	} else {
		log.Info("account:%+v", resp)
	}
	if resp, err := huobi.C2client.GetC2CBalance(huobi.Accounts["borrow"]); err != nil {
		log.Info("Get account error: %s", err)
	} else {
		log.Info("account:%+v", resp)
	}
	if resp, err := huobi.GetAccountBalance(huobi.Accounts["borrow"]); err != nil {
		log.Info("Get account error: %s", err)
	} else {
		log.Info("account:%+v", resp)
	}

	if resp, err := huobi.GetAccountBalance(huobi.Accounts["investment"]); err != nil {
		log.Info("Get account error: %s", err)
	} else {
		log.Info("account:%+v", resp)
	}
	if amount, err := huobi.BalanceOf("usdt"); err != nil {
		log.Info("Get amount error: %s", err)
	} else {
		log.Info("usdt:%+v", amount)
	}
	if amount, err := huobi.BalanceOf("btc"); err != nil {
		log.Info("Get amount error: %s", err)
	} else {
		log.Info("btc:%+v", amount)
	}

	//runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	subuserclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	accountwebsocketclientexample.RunAllExamples()
	orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}
