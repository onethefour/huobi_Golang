package websocketclientbase

import (
	"fmt"
	"github.com/gorilla/websocket"
	"huobi_Golang/common/log"
	"huobi_Golang/internal/gzip"
	"huobi_Golang/internal/model"
	"huobi_Golang/internal/requestbuilder"
	"huobi_Golang/pkg/model/auth"
	"huobi_Golang/pkg/model/base"
	"sync"
	"time"
)

const (
	websocketV2Path = "/ws/v2"
)

// It will be invoked after websocket v2 authentication response received
type AuthenticationV2ResponseHandler func(resp *auth.WebSocketV2AuthenticationResponse)

// The base class that responsible to get data from websocket authentication v2
type WebSocketV2ClientBase struct {
	host string
	conn *websocket.Conn

	authenticationResponseHandler AuthenticationV2ResponseHandler
	messageHandler                MessageHandler
	responseHandler               ResponseHandler

	stopReadChannel   chan int
	stopTickerChannel chan int
	ticker            *time.Ticker
	lastReceivedTime  time.Time
	sendMutex         *sync.Mutex

	requestBuilder *requestbuilder.WebSocketV2RequestBuilder
}

// Initializer
func (p *WebSocketV2ClientBase) Init(accessKey string, secretKey string, host string) *WebSocketV2ClientBase {
	p.host = host
	p.stopReadChannel = make(chan int, 1)
	p.stopTickerChannel = make(chan int, 1)
	p.requestBuilder = new(requestbuilder.WebSocketV2RequestBuilder).Init(accessKey, secretKey, host, websocketV2Path)
	p.sendMutex = &sync.Mutex{}
	return p
}

// Set callback handler
func (p *WebSocketV2ClientBase) SetHandler(authHandler AuthenticationV2ResponseHandler, msgHandler MessageHandler, repHandler ResponseHandler) {
	p.authenticationResponseHandler = authHandler
	p.messageHandler = msgHandler
	p.responseHandler = repHandler
}

// Connect to websocket server
// if autoConnect is true, then the connection can be re-connect if no data received after the pre-defined timeout
func (p *WebSocketV2ClientBase) Connect(autoConnect bool) {
	p.connectWebSocket()

	if autoConnect {
		p.startTicker()
	}
}

// Send data to websocket server
func (p *WebSocketV2ClientBase) Send(data string) {
	if p.conn == nil {
		log.Error("WebSocket sent error: no connection available")
		return
	}

	p.sendMutex.Lock()
	err := p.conn.WriteMessage(websocket.TextMessage, []byte(data))
	p.sendMutex.Unlock()

	if err != nil {
		log.Error("WebSocket sent error: data=%s, error=%s", data, err)
	}
}

// Close the connection to server
func (p *WebSocketV2ClientBase) Close() {
	p.stopTicker()
	p.disconnectWebSocket()
}

// connect to server
func (p *WebSocketV2ClientBase) connectWebSocket() {
	var err error
	url := fmt.Sprintf("wss://%s%s", p.host, websocketV2Path)
	log.Debug("WebSocket connecting...")
	p.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Error("WebSocket connected error: %s", err)
		return
	}
	log.Info("WebSocket connected")

	auth, err := p.requestBuilder.Build()
	if err != nil {
		log.Error("Signature generated error: %s", err)
		return
	}

	p.Send(auth)

	p.startReadLoop()
}

// disconnect with server
func (p *WebSocketV2ClientBase) disconnectWebSocket() {
	if p.conn == nil {
		return
	}

	// start a new goroutine to send a signal
	go p.stopReadLoop()

	log.Debug("WebSocket disconnecting...")
	err := p.conn.Close()
	if err != nil {
		log.Error("WebSocket disconnect error: %s", err)
		return
	}

	log.Info("WebSocket disconnected")
}

// initialize a ticker and start a goroutine tickerLoop()
func (p *WebSocketV2ClientBase) startTicker() {
	p.ticker = time.NewTicker(TimerIntervalSecond * time.Second)
	p.lastReceivedTime = time.Now()

	go p.tickerLoop()
}

// stop ticker and stop the goroutine
func (p *WebSocketV2ClientBase) stopTicker() {
	p.ticker.Stop()
	p.stopTickerChannel <- 1
}

// defines a for loop that will run based on ticker's frequency
// It checks the last data that received from server, if it is longer than the threshold,
// it will force disconnect server and connect again.
func (p *WebSocketV2ClientBase) tickerLoop() {
	log.Debug("tickerLoop started")
	for {
		select {
		// start a goroutine readLoop()
		case <-p.stopTickerChannel:
			log.Debug("tickerLoop stopped")
			return

		// Receive tick from tickChannel
		case <-p.ticker.C:
			elapsedSecond := time.Now().Sub(p.lastReceivedTime).Seconds()
			log.Debug("WebSocket received data %f sec ago", elapsedSecond)

			if elapsedSecond > ReconnectWaitSecond {
				log.Info("WebSocket reconnect...")
				p.disconnectWebSocket()
				p.connectWebSocket()
			}
		}
	}
}

// start a goroutine readLoop()
func (p *WebSocketV2ClientBase) startReadLoop() {
	go p.readLoop()
}

// stop the goroutine readLoop()
func (p *WebSocketV2ClientBase) stopReadLoop() {
	p.stopReadChannel <- 1
}

// defines a for loop to read data from server
// it will stop once it receives the signal from stopReadChannel
func (p *WebSocketV2ClientBase) readLoop() {
	log.Debug("readLoop started")
	for {
		select {
		// Receive data from stopChannel
		case <-p.stopReadChannel:
			log.Debug("readLoop stopped")
			return

		default:
			if p.conn == nil {
				log.Error("Read error: no connection available")
				time.Sleep(TimerIntervalSecond * time.Second)
				continue
			}

			msgType, buf, err := p.conn.ReadMessage()
			if err != nil {
				log.Error("Read error: %s", err)
				time.Sleep(TimerIntervalSecond * time.Second)
				continue
			}

			p.lastReceivedTime = time.Now()

			// decompress gzip data if it is binary message
			var message string
			if msgType == websocket.BinaryMessage {
				message, err = gzip.GZipDecompress(buf)
				if err != nil {
					log.Error("UnGZip data error: %s", err)
					return
				}
			} else if msgType == websocket.TextMessage {
				message = string(buf)
			}

			// Try to pass as PingV2Message
			// If it is Ping then respond Pong
			pingV2Msg := model.ParsePingV2Message(message)
			if pingV2Msg.IsPing() {
				log.Debug("Received Ping: %d", pingV2Msg.Data.Timestamp)
				pongMsg := fmt.Sprintf("{\"action\": \"pong\", \"data\": { \"ts\": %d } }", pingV2Msg.Data.Timestamp)
				p.Send(pongMsg)
				log.Debug("Respond  Pong: %d", pingV2Msg.Data.Timestamp)
			} else {
				// Try to pass as websocket v2 authentication response
				// If it is then invoke authentication handler
				wsV2Resp := base.ParseWSV2Resp(message)
				if wsV2Resp != nil {
					switch wsV2Resp.Action {
					case "req":
						authResp := auth.ParseWSV2AuthResp(message)
						if authResp != nil && p.authenticationResponseHandler != nil {
							p.authenticationResponseHandler(authResp)
						}

					case "sub", "push":
						{
							result, err := p.messageHandler(message)
							if err != nil {
								log.Error("Handle message error: %s", err)
								continue
							}
							if p.responseHandler != nil {
								p.responseHandler(result)
							}
						}
					}
				}
			}
		}
	}
}
