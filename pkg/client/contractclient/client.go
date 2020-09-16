package contractclient

import (
	"encoding/json"
	"huobi_Golang/cmd/config"
	"huobi_Golang/common/log"
	"huobi_Golang/internal"
	"huobi_Golang/internal/requestbuilder"
	"huobi_Golang/pkg/client/websocketclientbase"
	"huobi_Golang/pkg/model"
	"huobi_Golang/pkg/model/contract"
)

// Responsible to operate isolated margin
type ContractClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
	websocketclientbase.WebSocketClientBase
	publicUrlBuilder *requestbuilder.PublicUrlBuilder
}

func (p *ContractClient) Init(accessKey string, secretKey string, host string) *ContractClient {
	host = config.MARKET_URL
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	p.publicUrlBuilder = new(requestbuilder.PublicUrlBuilder).Init(host)
	p.WebSocketClientBase.Init(host)
	return p
}
func (p *ContractClient) FutureContractInfo(symbol, contractType, contractCode string) (ret *contract.ContractInfoResponse, err error) {
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", symbol)
	request.AddParam("contract_type", contractType)
	request.AddParam("contractCode", contractCode)

	strRequest := "/api/v1/contract_contract_info"
	url := p.publicUrlBuilder.Build(strRequest, request)
	resp, err := internal.HttpGet(url)
	if err != nil {
		return
	}
	log.Info(resp)
	ret = new(contract.ContractInfoResponse)
	err = json.Unmarshal([]byte(resp), ret)
	return
}

//
func (p *ContractClient) FutureHistoryBasic(symbol, period, basis_price_type string, size int64) (ret *contract.ContractBasis, err error) {
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", symbol)
	request.AddParam("period", period)
	request.AddParam("basis_price_type", basis_price_type)
	request.Add("size", size)
	strRequest := "/index/market/history/basis"
	url := p.publicUrlBuilder.Build(strRequest, request)
	resp, err := internal.HttpGet(url)
	if err != nil {
		return
	}
	ret = new(contract.ContractBasis)
	err = json.Unmarshal([]byte(resp), ret)
	return ret, err
}
func (p *ContractClient) FuturePrice(symbol string) (buyPrice float64, sellPrice float64, err error) {
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", symbol)
	strRequest := "/market/detail/merged"
	url := p.publicUrlBuilder.Build(strRequest, request)
	resp, err := internal.HttpGet(url)
	if err != nil {
		return
	}
	ret := new(contract.DetailMerged)
	log.Info(resp)
	if err = json.Unmarshal([]byte(resp), ret); err != nil {
		return 0, 0, err
	}
	return ret.Tick.Ask[0], ret.Tick.Bid[0], nil
}
