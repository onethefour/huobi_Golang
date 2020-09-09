package client

import (
	"encoding/json"
	"errors"
	log "huobi_Golang/common/log"
	"huobi_Golang/internal"
	"huobi_Golang/internal/requestbuilder"
	"huobi_Golang/pkg/model"
	"huobi_Golang/pkg/model/c2c"
)

type C2client struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
	publicUrlBuilder  *requestbuilder.PublicUrlBuilder
}

// Initializer
func (p *C2client) Init(accessKey string, secretKey string, host string) *C2client {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	p.publicUrlBuilder = new(requestbuilder.PublicUrlBuilder).Init(host)
	return p
}

func (p *C2client) GetC2CBalance(accountId string) (*c2c.AccountBalanceData, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("accountId", accountId)
	request.AddParam("currency", "")
	url := p.privateUrlBuilder.Build("GET", "/v2/c2c/account", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := &c2c.GetAccountBalanceResponse{}
	//log.Info(getResp)
	jsonErr := json.Unmarshal([]byte(getResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}
func (p *C2client) GetC2CRate(accountId string) error {
	request := new(model.GetRequest).Init()
	request.AddParam("accountId", accountId)
	//request.AddParam("symbols", "all")
	url := p.privateUrlBuilder.Build("GET", "/v1/cross-margin/loan-info", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return getErr
	}
	rjson, _ := json.Marshal(getResp)
	log.Info(string(rjson))
	return nil
}
