package taoli

import (
	"errors"
	"huobi_Golang/pkg/client"
	"huobi_Golang/pkg/model/common"
)

type Taoli struct {
	Id int
	//模式 0,正常模式 1,自动买手动卖 2,自动卖手动买
	Model int64
	//基础币 例如 btc
	BaseCurrency string
	//计量币 例如 usdt
	QuoteCurrency string
	*common.Symbol
	//货币交易所api
	*client.Client
	//执行上下文
	Datas []*Ctxt
	//停止运行标记
	Run bool
	//0 stop 1 running
	Status uint64
}
type Ctxt struct {
	BuyMount float64 //买的量 usdt 计量
	BuyPrice float64 //买的价格
	BuyOrder string  //买的订单

	//SellMount float64 //卖的量 btc 计量
	SellPrice float64 //卖的价格
	SellOrder string  //卖的订单

	TotalBaseCurrency float64 // 到此需要买的币(btc) 用于初始化
}

func (g *Taoli) init() error {
	resp, err := g.GetSymbols()
	if err != nil {
		return err
	}
	for k, symbol := range resp {
		if symbol.BaseCurrency == g.BaseCurrency && symbol.QuoteCurrency == g.QuoteCurrency {
			g.Symbol = &resp[k]
			return nil
		}
	}
	return errors.New("不存在的交易对:" + g.BaseCurrency + "/" + g.QuoteCurrency)
}
