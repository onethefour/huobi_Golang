package marketwebsocketclientexample

import (
	"fmt"
	"huobi_Golang/common/config"
	"huobi_Golang/common/log"
	"huobi_Golang/pkg/client/marketwebsocketclient"
	"huobi_Golang/pkg/model/market"
)

func RunAllExamples() {
	reqAndSubscribeCandlestick()
	reqAndSubscribeDepth()
	reqAndSubscribeMarketByPrice()
	subscribeFullMarketByPrice()
	subscribeBBO()
	reqAndSubscribeTrade()
	reqAndSubscribeLast24hCandlestick()
}

func reqAndSubscribeCandlestick() {

	client := new(marketwebsocketclient.CandlestickWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			client.Request("btcusdt", "1min", 1569361140, 1569366420, "2305")

			client.Subscribe("btcusdt", "1min", "2118")
		},
		func(response interface{}) {
			resp, ok := response.(market.SubscribeCandlestickResponse)
			if ok {
				if &resp != nil {
					if resp.Tick != nil {
						t := resp.Tick
						log.Info("Candlestick update, id: %d, count: %d, vol: %v [%v-%v-%v-%v]",
							t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
					}

					if resp.Data != nil {
						log.Info("WebSocket returned data, count=%d", len(resp.Data))
						for _, t := range resp.Data {
							log.Info("Candlestick data, id: %d, count: %d, vol: %v [%v-%v-%v-%v]",
								t.Id, t.Count, t.Vol, t.Open, t.Count, t.Low, t.High)
						}
					}
				}
			} else {
				log.Warn("Unknown response: %v", resp)
			}

		})

	client.Connect(true)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.UnSubscribe("btcusdt", "1min", "2118")

	client.Close()
	log.Info("Client closed")
}

func reqAndSubscribeDepth() {

	client := new(marketwebsocketclient.DepthWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			client.Request("btcusdt", market.STEP4, "1153")

			client.Subscribe("btcusdt", market.STEP4, "1153")
		},
		func(resp interface{}) {
			depthResponse, ok := resp.(market.SubscribeDepthResponse)
			if ok {
				if &depthResponse != nil {
					if depthResponse.Tick != nil {
						log.Info("WebSocket received depth update")
						if depthResponse.Tick.Asks != nil {
							a := depthResponse.Tick.Asks
							log.Info("Ask, count=%d", len(a))
							for i := len(a) - 1; i >= 0; i-- {
								log.Info("%v - %v", a[i][0], a[i][1])
							}
						}
						if depthResponse.Tick.Bids != nil {
							b := depthResponse.Tick.Bids
							log.Info("Bid, count=%d", len(b))
							for i := 0; i < len(b); i++ {
								log.Info("%v - %v", b[i][0], b[i][1])
							}
						}
					}

					if depthResponse.Data != nil {
						log.Info("WebSocket received depth data")
						if depthResponse.Data.Asks != nil {
							a := depthResponse.Data.Asks
							log.Info("Ask, count=%d", len(a))
							for i := len(a) - 1; i >= 0; i-- {
								log.Info("%v - %v", a[i][0], a[i][1])
							}
						}
						if depthResponse.Data.Bids != nil {
							b := depthResponse.Data.Bids
							log.Info("Bid, count=%d", len(b))
							for i := 0; i < len(b); i++ {
								log.Info("%v - %v", b[i][0], b[i][1])
							}
						}
					}
				}
			} else {
				log.Warn("Unknown response: %v", resp)
			}

		})

	client.Connect(true)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.UnSubscribe("btcusdt", "1min", "2118")

	client.Close()
	log.Info("Client closed")
}

func reqAndSubscribeMarketByPrice() {

	client := new(marketwebsocketclient.MarketByPriceWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			client.Request("btcusdt", "1437")

			client.Subscribe("btcusdt", "1437")
		},
		func(resp interface{}) {
			depthResponse, ok := resp.(market.SubscribeMarketByPriceResponse)
			if ok {
				if &depthResponse != nil {
					if depthResponse.Tick != nil {
						t := depthResponse.Tick
						log.Info("WebSocket received MBP update: prevSeqNum=%d, seqNum=%d", t.PrevSeqNum, t.SeqNum)
						if t.Asks != nil {
							log.Info("Ask, count=%d", len(t.Asks))
							for i := len(t.Asks) - 1; i >= 0; i-- {
								log.Info("%v - %v"+
									"", t.Asks[i][0], t.Asks[i][1])
							}
						}
						if t.Bids != nil {
							log.Info("Bid, count=%d", len(t.Bids))
							for i := 0; i < len(t.Bids); i++ {
								log.Info("%v - %v", t.Bids[i][0], t.Bids[i][1])
							}
						}
					}

					if depthResponse.Data != nil {
						d := depthResponse.Data
						log.Info("WebSocket received MBP data, seqNum=%d", d.SeqNum)
						if d.Asks != nil {
							a := d.Asks
							log.Info("Ask, count=%d", len(a))
							for i := len(a) - 1; i >= 0; i-- {
								log.Info("%v - %v", a[i][0], a[i][1])
							}
						}
						if d.Bids != nil {
							b := depthResponse.Data.Bids
							log.Info("Bid, count=%d", len(b))
							for i := 0; i < len(b); i++ {
								log.Info("%v - %v", b[i][0], b[i][1])
							}
						}
					}
				}
			} else {
				log.Warn("Unknown response: %v", resp)
			}

		})

	client.Connect(true)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.UnSubscribe("btcusdt", "1437")

	client.Close()
	log.Info("Client closed")
}

func subscribeFullMarketByPrice() {

	client := new(marketwebsocketclient.MarketByPriceWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			client.SubscribeFull("btcusdt", 20, "1437")
		},
		func(resp interface{}) {
			depthResponse, ok := resp.(market.SubscribeMarketByPriceResponse)
			if ok {
				if &depthResponse != nil {
					if depthResponse.Tick != nil {
						t := depthResponse.Tick
						log.Info("WebSocket received Full MBP update: seqNum=%d", t.SeqNum)
						if t.Asks != nil {
							log.Info("Ask, count=%d", len(t.Asks))
							for i := len(t.Asks) - 1; i >= 0; i-- {
								log.Info("%v - %v", t.Asks[i][0], t.Asks[i][1])
							}
						}
						if t.Bids != nil {
							log.Info("Bid, count=%d", len(t.Bids))
							for i := 0; i < len(t.Bids); i++ {
								log.Info("%v - %v", t.Bids[i][0], t.Bids[i][1])
							}
						}
					}
				}
			} else {
				log.Warn("Unknown response: %v", resp)
			}

		})

	client.Connect(true)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.UnSubscribeFull("btcusdt", 20, "1437")

	client.Close()
	log.Info("Client closed")
}

func subscribeBBO() {

	client := new(marketwebsocketclient.BestBidOfferWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			client.Subscribe("btcusdt", "2118")
		},
		func(resp interface{}) {
			bboResponse, ok := resp.(market.SubscribeBestBidOfferResponse)
			if ok {
				if bboResponse.Tick != nil {
					t := bboResponse.Tick
					log.Info("WebSocket received update, symbol: %s, ask: [%v, %v], bid: [%v, %v]", t.Symbol, t.Ask, t.AskSize, t.Bid, t.BidSize)
				}
			}

		})

	client.Connect(true)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.UnSubscribe("btcusdt", "2118")

	client.Close()
	log.Info("Connection closed")
}

func reqAndSubscribeTrade() {

	client := new(marketwebsocketclient.TradeWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			client.Request("btcusdt", "1608")

			client.Subscribe("btcusdt", "1608")
		},
		func(resp interface{}) {
			depthResponse, ok := resp.(market.SubscribeTradeResponse)
			if ok {
				if &depthResponse != nil {
					if depthResponse.Tick != nil && depthResponse.Tick.Data != nil {
						log.Info("WebSocket received trade update: count=%d", len(depthResponse.Tick.Data))
						for _, t := range depthResponse.Tick.Data {
							log.Info("Trade update, id: %d, price: %v, amount: %v", t.TradeId, t.Price, t.Amount)
						}
					}

					if depthResponse.Data != nil {
						log.Info("WebSocket received trade data: count=%d", len(depthResponse.Data))
						for _, t := range depthResponse.Data {
							log.Info("Trade data, id: %d, price: %v, amount: %v", t.TradeId, t.Price, t.Amount)
						}
					}
				}
			} else {
				log.Warn("Unknown response: %v", resp)
			}

		})

	client.Connect(true)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.UnSubscribe("btcusdt", "1608")

	client.Close()
	log.Info("Client closed")
}

func reqAndSubscribeLast24hCandlestick() {
	// Initialize a new instance
	client := new(marketwebsocketclient.Last24hCandlestickWebSocketClient).Init(config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Connected handler
		func() {
			client.Request("btcusdt", "1608")

			client.Subscribe("btcusdt", "1608")
		},
		// Response handler
		func(resp interface{}) {
			candlestickResponse, ok := resp.(market.SubscribeLast24hCandlestickResponse)
			if ok {
				if &candlestickResponse != nil {
					if candlestickResponse.Tick != nil {
						t := candlestickResponse.Tick
						log.Info("WebSocket received candlestick update, id: %d, count: %v, volume: %v [%v-%v-%v-%v]",
							t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
					}

					if candlestickResponse.Data != nil {
						t := candlestickResponse.Data
						log.Info("WebSocket received candlestick data, id: %d, count: %v, volume: %v [%v-%v-%v-%v]",
							t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
					}
				}
			} else {
				log.Warn("Unknown response: %v", resp)
			}
		})

	// Connect to the server and wait for the handler to handle the response
	client.Connect(true)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.UnSubscribe("btcusdt", "1608")

	client.Close()
	log.Info("Client closed")
}