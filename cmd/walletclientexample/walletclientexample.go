package walletclientexample

import (
	"huobi_Golang/common/config"
	"huobi_Golang/common/log"
	"huobi_Golang/pkg/client"
	"huobi_Golang/pkg/model/wallet"
)

func RunAllExamples() {
	getDepositAddress()
	getWithdrawQuota()
	createWithdraw()
	cancelWithdraw()
	queryDepositWithdraw()
}

func getDepositAddress() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "usdt"
	resp, err := client.GetDepositAddress(currency)
	if err != nil {
		log.Error("Get deposit address error: %s", err)
	} else {
		log.Info("Get deposit address, count=%d", len(resp))
		for _, result := range resp {
			log.Info("DepositAddress: %+v", result)
		}
	}
}

func getWithdrawQuota() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "usdt"
	resp, err := client.GetWithdrawQuota(currency)
	if err != nil {
		log.Error("Get withdraw quota error: %s", err)
	} else {
		log.Info("Currency: %s, Chain: %s, MaxWithdrawAmt: %s", resp.Currency, resp.Chains[0].Chain, resp.Chains[0].MaxWithdrawAmt)
	}
}

func createWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	createWithdrawRequest := wallet.CreateWithdrawRequest{
		Address:  "xxxx",
		Amount:   "1.0",
		Currency: "usdt",
		Fee:      "1.0"}
	resp, err := client.CreateWithdraw(createWithdrawRequest)
	if err != nil {
		log.Error("Create withdraw request error: %s", err)
	} else {
		log.Info("Create withdraw request successfully: id=%d", resp)
	}
}

func cancelWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelWithdraw(12345)
	if err != nil {
		log.Error("Cancel withdraw error: %s", err)
	} else {
		log.Info("Cancel withdraw successfully: id=%d", resp)
	}
}

func queryDepositWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	depositType := "deposit"
	queryDepositWithdrawOptionalRequest := wallet.QueryDepositWithdrawOptionalRequest{Currency: "usdt"}
	resp, err := client.QueryDepositWithdraw(depositType, queryDepositWithdrawOptionalRequest)
	if err != nil {
		log.Error("Query deposit-withdraw history error: %s", err)
	} else {
		log.Info("Query deposit-withdraw history, count=%d", len(resp))
		for _, result := range resp {
			log.Info("resp: %+v", result)
		}
	}
}
