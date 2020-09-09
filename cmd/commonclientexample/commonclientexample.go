package commonclientexample

import (
	"huobi_Golang/common/config"
	"huobi_Golang/common/log"
	"huobi_Golang/pkg/client"
	"huobi_Golang/pkg/model/common"
)

func RunAllExamples() {
	getSystemStatus()
	getSymbols()
	getCurrencys()
	getV2ReferenceCurrencies()
	getTimestamp()
}

func getSystemStatus() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetSystemStatus()
	if err != nil {
		log.Error("Get system status error: %s", err)
	} else {
		log.Info("Get system status %s", resp)
	}
}

func getSymbols() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetSymbols()
	if err != nil {
		log.Error("Get symbols error: %s", err)
	} else {
		log.Info("Get symbols, count=%d", len(resp))
		for _, result := range resp {
			log.Info("symbol=%s, BaseCurrency=%s, QuoteCurrency=%s", result.Symbol, result.BaseCurrency, result.QuoteCurrency)
		}
	}
}

func getCurrencys() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetCurrencys()

	if err != nil {
		log.Error("Get currency error: %s", err)
	} else {
		log.Info("Get currency, count=%d", len(resp))
		for _, result := range resp {
			log.Info("currency: %+v", result)
		}
	}
}

func getV2ReferenceCurrencies() {
	optionalRequest := common.GetV2ReferenceCurrencies{Currency: "", AuthorizedUser: "true"}

	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetV2ReferenceCurrencies(optionalRequest)

	if err != nil {
		log.Error("Get reference currency error: %s", err)
	} else {
		log.Info("Get reference currency, count=%d", len(resp))
		for _, result := range resp {
			log.Info("currency:%s, ", result.Currency)

			for _, chain := range result.Chains {
				log.Info("Chain: %+v", chain)
			}
		}
	}
}

func getTimestamp() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetTimestamp()

	if err != nil {
		log.Error("Get timestamp error: %s", err)
	} else {
		log.Info("Get timestamp: %d", resp)
	}
}
