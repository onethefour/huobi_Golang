package grid

import (
	"encoding/json"
	"fmt"
	"strconv"

	appmodels "huobi_Golang/api/models"
	"huobi_Golang/common/log"
	"huobi_Golang/pkg/model/order"
	"huobi_Golang/services"
	"time"
)

func (g *Grid) Stop() {
	g.Run = false
	for g.Status == 1 {
		time.Sleep(time.Second * 1)
	}
	return
}
func (g *Grid) Register() {
	if err := services.Register(g.Id, g); err != nil {
		log.Info(err.Error())
		services.Stop(g.Id)
		services.Register(g.Id, g)
	}
}
func (g *Grid) Start() {
	g.Register()
	defer services.Remove(g.Id)
	// 标记为running
	g.Status = 1
	defer func() {
		g.Status = 0
	}()

	if g.Model == 1 {
		g.work_classical()
	} else if g.Model == 2 {
		g.work_buy()
	} else if g.Model == 3 {
		g.work_sell()
	}
}
func (g *Grid) Action() error {
	return nil
}

func (g *Grid) BuyPrice() (price float64, err error) {
	if ticker, err := g.GetLast24hCandlestickAskBid(g.BaseCurrency + g.QuoteCurrency); err != nil {
		return 0, err
	} else {
		price, _ = ticker.Ask[0].Float64()
		return
	}
}
func (g *Grid) SellPrice() (price float64, err error) {
	if ticker, err := g.GetLast24hCandlestickAskBid(g.BaseCurrency + g.QuoteCurrency); err != nil {
		return 0, err
	} else {
		price, _ = ticker.Bid[0].Float64()
		return
	}
}
func (g *Grid) Price() (buyPrice float64, sellPrice float64, err error) {
	if ticker, err := g.GetLast24hCandlestickAskBid(g.BaseCurrency + g.QuoteCurrency); err != nil {
		return 0, 0, err
	} else {
		buyPrice, _ = ticker.Bid[0].Float64()
		sellPrice, _ = ticker.Ask[0].Float64()
		return
	}
}
func (m *Grid) BuyMarket(_usdt float64) error {

	PlaceOrderRequest := &order.PlaceOrderRequest{
		AccountId: m.Accounts["spot"],
		Symbol:    m.BaseCurrency + m.QuoteCurrency,
		Type:      "buy-market",
		Amount:    m.String(_usdt, m.Symbol.AmountPrecision),
		Source:    "api",
	}

	log.Info("%v %v", m.Symbol.AmountPrecision, PlaceOrderRequest)

	ret, err := m.PlaceOrder(PlaceOrderRequest)
	if err != nil {
		return err
	}
	return new(appmodels.Orders).Add(ret.Data, m.Id)
}
func (m *Grid) SellMarket(_btc float64) error {

	PlaceOrderRequest := &order.PlaceOrderRequest{
		AccountId: m.Accounts["spot"],
		Symbol:    m.BaseCurrency + m.QuoteCurrency,
		Type:      "sell-market",
		Amount:    m.String(_btc, m.Symbol.AmountPrecision),
		Source:    "api",
	}

	log.Info("%v %v", m.Symbol.AmountPrecision, PlaceOrderRequest)

	ret, err := m.PlaceOrder(PlaceOrderRequest)
	if err != nil {
		return err
	}
	return new(appmodels.Orders).Add(ret.Data, m.Id)
}
func (m *Grid) String(f float64, decimal int) string {
	if decimal == 0 {
		return fmt.Sprintf("%v", uint64(f))
	} else if decimal <= 10 {
		return fmt.Sprintf(fmt.Sprintf("%%.%vf", decimal), f)
	} else {
		return fmt.Sprintf("%.10f", f)
	}
}

func (m *Grid) GetOrder(orderid string) (order *order.GetOrderResponse, err error) {
	if order, err = m.GetOrderById(orderid); err != nil {
		return nil, err
	}
	if order.Data.Amount_f, err = strconv.ParseFloat(order.Data.Amount, 64); err != nil {
		return nil, err
	}

	if order.Data.Price_f, err = strconv.ParseFloat(order.Data.Price, 64); err != nil {
		return nil, err
	}
	return order, nil
}

//CancelOrder 取消订单
func (m *Grid) CancelOrder(orderid string) error {
	_, err := m.CancelOrderById(orderid)
	return err
}
func (m *Grid) Save() error {
	db := new(appmodels.Stracy)
	data, _ := json.Marshal(m.Datas)
	return db.Save(m.Id, string(data))
}
func (m *Grid) sleepOrStop(ts uint64) error {
	for i := 0; uint64(i) < ts; i++ {
		time.Sleep(time.Second)
		if m.Run == false {
			return fmt.Errorf("stop server:%v", m.Id)
		}
	}
	return nil
}

//SellOrder 限价卖币订单
func (m *Grid) SellOrder(_btc float64, _price float64) (string, error) {
	resq := &order.PlaceOrderRequest{
		AccountId: m.Accounts["splot"],
		Symbol:    m.BaseCurrency + m.QuoteCurrency,
		Type:      "sell-limit",
		Amount:    m.String(_btc, m.AmountPrecision),
		Price:     m.String(_price, m.PricePrecision),
		Source:    "api",
	}
	ret, err := m.PlaceOrder(resq)
	if err != nil {
		return "", err
	}
	new(appmodels.Orders).Add(ret.Data, m.Id)
	return ret.Data, nil
}

//BuyOrder 限价买币订单
func (m *Grid) BuyOrder(_btc float64, _price float64) (string, error) {
	resq := &order.PlaceOrderRequest{
		AccountId: m.Accounts["splot"],
		Symbol:    m.BaseCurrency + m.QuoteCurrency,
		Type:      "buy-limit",
		Amount:    m.String(_btc, m.AmountPrecision),
		Price:     m.String(_price, m.PricePrecision),
		Source:    "api",
	}
	ret, err := m.PlaceOrder(resq)
	if err != nil {
		return "", err
	}
	new(appmodels.Orders).Add(ret.Data, m.Id)
	return ret.Data, nil
}
