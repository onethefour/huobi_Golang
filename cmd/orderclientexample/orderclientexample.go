package orderclientexample

import (
	"huobi_Golang/common/config"
	"huobi_Golang/common/log"
	"huobi_Golang/pkg/client"
	"huobi_Golang/pkg/model"
	"huobi_Golang/pkg/model/order"
)

func RunAllExamples() {
	placeOrder()
	placeOrders()
	cancelOrderById()
	cancelOrderByClient()
	getOpenOrders()
	cancelOrdersByCriteria()
	cancelOrdersByIds()
	getOrderById()
	getOrderByCriteria()
	getMatchResultById()
	getHistoryOrders()
	getLast48hOrders()
	getMatchResultByCriteria()
	getTransactFeeRate()
}

func placeOrder() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := order.PlaceOrderRequest{
		AccountId: config.AccountId,
		Type:      "buy-limit",
		Source:    "spot-api",
		Symbol:    "btcusdt",
		Price:     "1.1",
		Amount:    "1",
	}
	resp, err := client.PlaceOrder(&request)
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			log.Info("Place order successfully, order id: %s", resp.Data)
		case "error":
			log.Error("Place order error: %s", resp.ErrorMessage)
		}
	}
}

func placeOrders() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := order.PlaceOrderRequest{
		AccountId: config.AccountId,
		Type:      "buy-limit",
		Source:    "spot-api",
		Symbol:    "btcusdt",
		Price:     "1.1",
		Amount:    "1",
	}

	var requests []order.PlaceOrderRequest
	requests = append(requests, request)
	requests = append(requests, request)
	resp, err := client.PlaceOrders(requests)
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, r := range resp.Data {
					if r.OrderId != 0 {
						log.Info("Place order successfully: order id %d", r.OrderId)
					} else {
						log.Info("Place order error: %s", r.ErrorMessage)
					}
				}
			}
		case "error":
			log.Error("Place order error: %s", resp.ErrorMessage)
		}
	}
}

func cancelOrderById() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelOrderById("1")
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			log.Info("Cancel order successfully, order id: %s", resp.Data)
		case "error":
			log.Info("Cancel order error: %s", resp.ErrorMessage)
		}
	}
}

func cancelOrderByClient() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelOrderByClientOrderId("1")
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			log.Info("Cancel order successfully, order id: %d", resp.Data)
		case "error":
			log.Info("Cancel order error: %s", resp.ErrorMessage)
		}
	}
}

func getOpenOrders() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("account-id", config.AccountId)
	request.AddParam("symbol", "btcusdt")
	resp, err := client.GetOpenOrders(request)
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, o := range resp.Data {
					log.Info("Open orders, symbol: %s, price: %s, amount: %s", o.Symbol, o.Price, o.Amount)
				}
				log.Info("There are total %d open orders", len(resp.Data))
			}
		case "error":
			log.Error("Get open order error: %s", resp.ErrorMessage)
		}
	}
}

func cancelOrdersByCriteria() {
	request := order.CancelOrdersByCriteriaRequest{
		AccountId: config.AccountId,
		Symbol:    "btcusdt",
	}

	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelOrdersByCriteria(&request)
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				d := resp.Data
				log.Info("Cancel orders successfully, success count: %d, failed count: %d, next id: %d", d.SuccessCount, d.FailedCount, d.NextId)
			}
		case "error":
			log.Error("Cancel orders error: %s", resp.ErrorMessage)
		}
	}
}

func cancelOrdersByIds() {
	request := order.CancelOrdersByIdsRequest{
		OrderIds: []string{"1", "2"},
	}

	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelOrdersByIds(&request)
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				if resp.Data.Success != nil {
					for _, id := range resp.Data.Success {
						log.Info("Cancel orders successfully, id: %s", id)
					}
				}
				if resp.Data.Failed != nil {
					for _, f := range resp.Data.Failed {
						id := f.OrderId
						if id == "" {
							id = f.ClientOrderId
						}
						log.Error("Cancel orders failed, id: %s, error: %s", id, f.ErrorMessage)
					}
				}
			}
		case "error":
			log.Error("Cancel orders error: %s", resp.ErrorMessage)
		}
	}
}

func getOrderById() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetOrderById("1")
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				o := resp.Data
				log.Info("Get order, symbol: %s, price: %s, amount: %s, filled amount: %s, filled cash amount: %s, filled fees: %s",
					o.Symbol, o.Price, o.Amount, o.FilledAmount, o.FilledCashAmount, o.FilledFees)
			}
		case "error":
			log.Error("Get order by id error: %s", resp.ErrorMessage)
		}
	}
}

func getOrderByCriteria() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("clientOrderId", "cid12345")
	resp, err := client.GetOrderByCriteria(request)
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				o := resp.Data
				log.Info("Get order, symbol: %s, price: %s, amount: %s, filled amount: %s, filled cash amount: %s, filled fees: %s",
					o.Symbol, o.Price, o.Amount, o.FilledAmount, o.FilledCashAmount, o.FilledFees)
			}
		case "error":
			log.Error("Get order by criteria error: %s", resp.ErrorMessage)
		}
	}
}

func getMatchResultById() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetMatchResultsById("63403286375")
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, r := range resp.Data {
					log.Info("Match result, symbol: %s, filled amount: %s, filled fees: %s", r.Symbol, r.FilledAmount, r.FilledFees)
				}
				log.Info("There are total %d match results", len(resp.Data))
			}
		case "error":
			log.Error("Get match results error: %s", resp.ErrorMessage)
		}
	}
}

func getHistoryOrders() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", "btcusdt")
	request.AddParam("states", "canceled")
	resp, err := client.GetHistoryOrders(request)
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, o := range resp.Data {
					log.Info("Order history, symbol: %s, price: %s, amount: %s, state: %s", o.Symbol, o.Price, o.Amount, o.State)
				}
				log.Info("There are total %d orders", len(resp.Data))
			}
		case "error":
			log.Error("Get history order error: %s", resp.ErrorMessage)
		}
	}
}

func getLast48hOrders() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", "btcusdt")
	resp, err := client.GetLast48hOrders(request)
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, o := range resp.Data {
					log.Info("Order history, symbol: %s, price: %s, amount: %s, state: %s", o.Symbol, o.Price, o.Amount, o.State)
				}
				log.Info("There are total %d orders", len(resp.Data))
			}
		case "error":
			log.Error("Get history order error: %s", resp.ErrorMessage)
		}
	}
}

func getMatchResultByCriteria() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", "btcusdt")
	resp, err := client.GetMatchResultsByCriteria(request)
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, r := range resp.Data {
					log.Info("Match result, symbol: %s, filled amount: %s, filled fees: %s", r.Symbol, r.FilledAmount, r.FilledFees)
				}
				log.Info("There are total %d match results", len(resp.Data))
			}
		case "error":
			log.Error("Get match results error: %s", resp.ErrorMessage)
		}
	}
}

func getTransactFeeRate() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("symbols", "btcusdt,eosht")
	resp, err := client.GetTransactFeeRate(request)
	if err != nil {
		log.Error(err.Error())
	} else {
		switch resp.Code {
		case 200:
			if resp.Data != nil {
				for _, f := range resp.Data {
					log.Info("Fee rate , symbol: %s, maker-taker fee: %s-%s, actual maker-taker fee: %s-%s",
						f.Symbol, f.MakerFeeRate, f.TakerFeeRate, f.ActualMakerRate, f.ActualTakerRate)
				}
				log.Info("There are total %d fee rate result", len(resp.Data))
			}
		default:
			log.Error("Get fee error: %s", resp.Message)
		}
	}
}
