package traderexample

import (
	"fmt"
	"huobi_Golang/common/config"
	"huobi_Golang/common/log"
	"huobi_Golang/pkg/client/marketwebsocketclient"
	"huobi_Golang/pkg/model/market"
)

func RunAllExamples() {
	subMultipleBBO()
}

func subMultipleBBO() {
	client := new(marketwebsocketclient.BestBidOfferWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			go client.Subscribe("btcusdt", "")
			go client.Subscribe("etcusdt", "")
			go client.Subscribe("bchusdt", "")
			go client.Subscribe("bsvusdt", "")
			go client.Subscribe("dashusdt", "")
			go client.Subscribe("zecusdt", "")
		},
		func(resp interface{}) {
			bboResponse, ok := resp.(market.SubscribeBestBidOfferResponse)
			if ok {
				if bboResponse.Tick != nil {
					t := bboResponse.Tick
					log.Info("Received update, symbol: %s, ask: [%v, %v], bid: [%v, %v]", t.Symbol, t.Ask, t.AskSize, t.Bid, t.BidSize)
				}
			}

		})

	client.Connect(true)

	fmt.Println("Press ENTER to stop...")
	fmt.Scanln()

	client.Close()
	log.Info("Connection closed")
}
