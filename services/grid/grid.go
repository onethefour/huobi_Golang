package grid

import (
	"encoding/json"
	"errors"
	"huobi_Golang/api/models"
	"huobi_Golang/common/config"
	"huobi_Golang/pkg/client"
	"huobi_Golang/pkg/model/common"
)

func NewGrid(stracy *models.Stracy) (*Grid, error) {
	datas := make([]*Ctxt, 0)
	if err := json.Unmarshal([]byte(stracy.Datas), &datas); err != nil {
		return nil, err
	}
	grid := &Grid{
		Id:            stracy.Id,
		Model:         stracy.Model,
		BaseCurrency:  stracy.BaseCurrency,
		QuoteCurrency: stracy.QuoteCurrency,
		Client:        client.NewClient(stracy.SecretKey, stracy.AccessKey, config.Host),
		Datas:         datas,
	}
	err := grid.init()
	return grid, err
}

type Grid struct {
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

func (g *Grid) init() error {
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
