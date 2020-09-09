package orderwebsocketclientexample

import (
	"fmt"
	"huobi_Golang/common/config"
	"huobi_Golang/common/log"
	"huobi_Golang/pkg/client/orderwebsocketclient"
	"huobi_Golang/pkg/model/auth"
	"huobi_Golang/pkg/model/order"
)

func RunAllExamples() {
	reqOrderV1()
	reqOrdersV1()
	subOrderUpdateV1()
	subOrderUpdateV2()
	subTradeClear()
}

func reqOrderV1() {
	client := new(orderwebsocketclient.RequestOrderWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)
	client.SetHandler(
		func(resp *auth.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				err := client.Request("1", "1601")
				if err != nil {
					log.Error("Request error: %s", err)
				} else {
					log.Debug("Sent request")
				}
			} else {
				log.Error("Authentication error: %d", resp.ErrorCode)
			}

		},
		func(resp interface{}) {
			reqResponse, ok := resp.(order.RequestOrderV1Response)
			if ok {
				if &reqResponse.Data != nil {
					o := reqResponse.Data
					log.Info("Request order, id: %d, state: %s, symbol: %s, price: %s, filled amount: %s", o.Id, o.State, o.Symbol, o.Price, o.FilledAmount)
				} else {
					log.Error("Request order error: %s", reqResponse.ErrorCode)
				}
			} else {
				log.Warn("Received unknown response: %v", resp)
			}
		})

	err := client.Connect(true)
	if err != nil {
		log.Error("Client Connect error: %s", err)
		return
	}

	fmt.Println("Press ENTER to stop...")
	fmt.Scanln()

	client.Close()
	log.Info("Client closed")
}

func reqOrdersV1() {
	client := new(orderwebsocketclient.RequestOrdersWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)
	client.SetHandler(
		func(resp *auth.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				req := order.RequestOrdersRequest{
					AccountId: 11136102,
					Symbol:    "btcusdt",
					States:    "submitted, created, canceled",
				}
				err := client.Request(req)
				if err != nil {
					log.Error("Request error: %s", err)
				} else {
					log.Debug("Sent request")
				}
			} else {
				log.Error("Authentication error: %d", resp.ErrorCode)
			}

		},
		func(resp interface{}) {
			reqResponse, ok := resp.(order.RequestOrdersV1Response)
			if ok {
				if &reqResponse.Data != nil {
					for _, o := range reqResponse.Data {
						log.Info("Request order, id: %d, state: %s, symbol: %s, price: %s, filled amount: %s", o.Id, o.State, o.Symbol, o.Price, o.FilledAmount)
					}
				} else {
					log.Error("Request order error: %s", reqResponse.ErrorCode)
				}
			} else {
				log.Warn("Received unknown response: %+v", resp)
			}
		})

	err := client.Connect(true)
	if err != nil {
		log.Error("Client Connect error: %s", err)
		return
	}

	fmt.Println("Press ENTER to stop...")
	fmt.Scanln()

	client.Close()
	log.Info("Client closed")
}

func subOrderUpdateV1() {
	// Initialize a new instance
	client := new(orderwebsocketclient.SubscribeOrderWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Connected handler
		func(resp *auth.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				err := client.Subscribe("btcusdt", "1601")
				if err != nil {
					log.Error("Subscribe error: %s", err)
				} else {
					log.Debug("Sent subscription")
				}
			} else {
				log.Error("Authentication error: %d", resp.ErrorCode)
			}

		},
		// Response handler
		func(resp interface{}) {
			subResponse, ok := resp.(order.SubscribeOrderV1Response)
			if ok {
				if &subResponse.Data != nil {
					o := subResponse.Data
					log.Info("Order update, id: %d, state: %s, symbol: %s, price: %s, filled amount: %s", o.OrderId, o.OrderState, o.Symbol, o.Price, o.FilledAmount)
				}
			} else {
				log.Warn("Received unknown response: %+v", resp)
			}
		})

	// Connect to the server and wait for the handler to handle the response
	err := client.Connect(true)
	if err != nil {
		log.Error("Client Connect error: %s", err)
		return
	}

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	err = client.UnSubscribe("1", "1250")
	if err != nil {
		log.Error("UnSubscribe error: %s", err)
	}

	client.Close()
	log.Info("Client closed")
}

func subOrderUpdateV2() {
	// Initialize a new instance
	client := new(orderwebsocketclient.SubscribeOrderWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Connected handler
		func(resp *auth.WebSocketV2AuthenticationResponse) {
			if resp.IsSuccess() {
				// Subscribe if authentication passed
				client.Subscribe("btcusdt", "1149")
			} else {
				log.Error("Authentication error, code: %d, message:%s", resp.Code, resp.Message)
			}
		},
		// Response handler
		func(resp interface{}) {
			subResponse, ok := resp.(order.SubscribeOrderV2Response)
			if ok {
				if subResponse.Action == "sub" {
					if subResponse.IsSuccess() {
						log.Info("Subscription topic %s successfully", subResponse.Ch)
					} else {
						log.Error("Subscription topic %s error, code: %d, message: %s", subResponse.Ch, subResponse.Code, subResponse.Message)
					}
				} else if subResponse.Action == "push" {
					if subResponse.Data != nil {
						o := subResponse.Data
						log.Info("Order update, event: %s, symbol: %s, type: %s, status: %s",
							o.EventType, o.Symbol, o.Type, o.OrderStatus)
					}
				}
			} else {
				log.Warn("Received unknown response: %v", resp)
			}
		})

	// Connect to the server and wait for the handler to handle the response
	client.Connect(true)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.UnSubscribe("1", "1250")

	client.Close()
	log.Info("Client closed")
}

func subTradeClear() {
	// Initialize a new instance
	client := new(orderwebsocketclient.SubscribeTradeClearWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Connected handler
		func(resp *auth.WebSocketV2AuthenticationResponse) {
			if resp.IsSuccess() {
				// Subscribe if authentication passed
				client.Subscribe("btcusdt", "1149")
			} else {
				log.Error("Authentication error, code: %d, message:%s", resp.Code, resp.Message)
			}
		},
		// Response handler
		func(resp interface{}) {
			subResponse, ok := resp.(order.SubscribeTradeClearResponse)
			if ok {
				if subResponse.Action == "sub" {
					if subResponse.IsSuccess() {
						log.Info("Subscription topic %s successfully", subResponse.Ch)
					} else {
						log.Error("Subscription topic %s error, code: %d, message: %s", subResponse.Ch, subResponse.Code, subResponse.Message)
					}
				} else if subResponse.Action == "push" {
					if subResponse.Data != nil {
						o := subResponse.Data
						log.Info("Order update, symbol: %s, order id: %d, price: %s, volume: %s",
							o.Symbol, o.OrderId, o.TradePrice, o.TradeVolume)
					}
				}
			} else {
				log.Warn("Received unknown response: %v", resp)
			}
		})

	// Connect to the server and wait for the handler to handle the response
	client.Connect(true)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.UnSubscribe("btcusdt", "1250")

	client.Close()
	log.Info("Client closed")
}
