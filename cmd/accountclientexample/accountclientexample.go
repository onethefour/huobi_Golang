package accountclientexample

import (
	"github.com/shopspring/decimal"
	"huobi_Golang/common/config"
	"huobi_Golang/common/log"
	"huobi_Golang/pkg/client"
	"huobi_Golang/pkg/model/account"
	"huobi_Golang/pkg/model/subuser"
)

func RunAllExamples() {
	getAccountInfo()
	getAccountBalance()
	transferAccount()
	getAccountHistory()
	getAccountLedger()
	transferFromFutureToSpot()
	transferFromSpotToFuture()
	subUserTransfer()
	getSubUserAggregateBalance()
	getSubUserAccount()
}

func getAccountInfo() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetAccountInfo()
	if err != nil {
		log.Error("Get account error: %s", err)
	} else {
		log.Info("Get account, count=%d", len(resp))
		for _, result := range resp {
			log.Info("account: %+v", result)
		}
	}
}

func transferAccount() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := account.TransferAccountRequest{
		FromUser:        125753978,
		FromAccountType: "spot",
		FromAccount:     11136102,
		ToUser:          128654510,
		ToAccountType:   "spot",
		ToAccount:       12825690,
		Currency:        "ht",
		Amount:          "0.18",
	}
	resp, err := client.TransferAccount(request)
	if err != nil {
		log.Error("Transfer account error: %s", err)
	} else {
		log.Info("Transfer account, %v", resp.Data)
	}
}

func getAccountBalance() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetAccountBalance(config.AccountId)
	if err != nil {
		log.Error("Get account balance error: %s", err)
	} else {
		log.Info("Get account balance: id=%d, type=%s, state=%s, count=%d",
			resp.Id, resp.Type, resp.State, len(resp.List))
		if resp.List != nil {
			for _, result := range resp.List {
				log.Info("Account balance: %+v", result)
			}
		}
	}
}

func getAccountHistory() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	getAccountHistoryOptionalRequest := account.GetAccountHistoryOptionalRequest{}
	resp, err := client.GetAccountHistory(config.AccountId, getAccountHistoryOptionalRequest)
	if err != nil {
		log.Error("Get account history error: %s", err)
	} else {
		log.Info("Get account history, count=%d", len(resp))
		for _, result := range resp {
			log.Info("Account history: %+v", result)
		}
	}
}

func getAccountLedger() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	getAccountLedgerOptionalRequest := account.GetAccountLedgerOptionalRequest{}
	resp, err := client.GetAccountLedger(config.AccountId, getAccountLedgerOptionalRequest)
	if err != nil {
		log.Error("Get account ledger error: %s", err)
	} else {
		log.Info("Get account ledger, count=%d", len(resp))
		for _, l := range resp {
			log.Info("Account legder: AccountId: %d, Currency: %s, Amount: %v, Transferer: %d, Transferee: %d", l.AccountId, l.Currency, l.TransactAmt, l.Transferer, l.Transferee)
		}
	}
}

func getSubUserAccount() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetSubUserAccount(config.SubUid)
	if err != nil {
		log.Error("Get sub user account error: %s", err)
	} else {
		log.Info("Get sub user account, count=%d", len(resp))
		for _, account := range resp {
			log.Info("account id: %d, type: %s, currency count=%d", account.Id, account.Type, len(account.List))
			if account.List != nil {
				for _, currency := range account.List {
					log.Info("currency: %+v", currency)
				}
			}
		}
	}
}

func transferFromFutureToSpot() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	futuresTransferRequest := account.FuturesTransferRequest{Currency: "btc", Amount: decimal.NewFromFloat(0.001), Type: "futures-to-pro"}
	resp, err := client.FuturesTransfer(futuresTransferRequest)
	if err != nil {
		log.Error("Transfer from future to spot error: %s", err)
	} else {
		log.Info("Transfer from future to spot success: id=%d", resp)
	}
}

func transferFromSpotToFuture() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	futuresTransferRequest := account.FuturesTransferRequest{Currency: "btc", Amount: decimal.NewFromFloat(0.001), Type: "pro-to-futures"}
	resp, err := client.FuturesTransfer(futuresTransferRequest)
	if err != nil {
		log.Error("Transfer from spot to future error: %s", err)
	} else {
		log.Info("Transfer from spot to future success: id=%d", resp)
	}
}

func getSubUserAggregateBalance() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetSubUserAggregateBalance()
	if err != nil {
		log.Error("Get sub user aggregated balance error: %s", err)
	} else {
		log.Info("Get sub user aggregated balance, count=%d", len(resp))
		for _, result := range resp {
			log.Info("balance: %+v", result)
		}
	}
}

func subUserTransfer() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "usdt"
	subUserTransferRequest := subuser.SubUserTransferRequest{
		SubUid:   config.SubUid,
		Currency: currency,
		Amount:   decimal.NewFromInt(1),
		Type:     "master-transfer-in",
	}
	resp, err := client.SubUserTransfer(subUserTransferRequest)
	if err != nil {
		log.Error("Transfer error: %s", err)
	} else {
		log.Info("Transfer successfully, id=%s", resp)

	}
}
