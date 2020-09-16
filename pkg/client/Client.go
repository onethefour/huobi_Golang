package client

import "C"
import (
	"fmt"
	log "huobi_Golang/common/log"
	"huobi_Golang/pkg/client/contractclient"
	"huobi_Golang/pkg/model/c2c"
	"strconv"

	"github.com/shopspring/decimal"
)

type Client struct {
	Accounts map[string]string //spot：现货账户， margin：逐仓杠杆账户，otc：OTC 账户，point：点卡账户，super-margin：全仓杠杆账户, investment: C2C杠杆借出账户, borrow: C2C杠杆借入账户
	*C2client
	*AccountClient
	*CommonClient
	*CrossMarginClient
	*ETFClient
	*IsolatedMarginClient
	*MarketClient
	*OrderClient
	*SubUserClient
	*WalletClient
	*contractclient.ContractClient
	AccessKey string
	SecretKey string
	Host      string
}

func NewClient(AccessKey, SecretKey, Host string) *Client {
	if Host == "" {
		Host = "api.huobi.pro"
	}
	client := &Client{
		AccessKey: AccessKey,
		SecretKey: SecretKey,
		Host:      Host,
		Accounts:  make(map[string]string),
	}

	client.C2client = new(C2client).Init(AccessKey, SecretKey, Host)
	client.AccountClient = new(AccountClient).Init(AccessKey, SecretKey, Host)
	client.CommonClient = new(CommonClient).Init(Host)
	client.CrossMarginClient = new(CrossMarginClient).Init(AccessKey, SecretKey, Host)
	client.ETFClient = new(ETFClient).Init(AccessKey, SecretKey, Host)
	client.IsolatedMarginClient = new(IsolatedMarginClient).Init(AccessKey, SecretKey, Host)
	client.MarketClient = new(MarketClient).Init(Host)
	client.OrderClient = new(OrderClient).Init(AccessKey, SecretKey, Host)
	client.SubUserClient = new(SubUserClient).Init(AccessKey, SecretKey, Host)
	client.WalletClient = new(WalletClient).Init(AccessKey, SecretKey, Host)
	client.ContractClient = new(contractclient.ContractClient).Init(AccessKey, SecretKey, Host)
	resp, err := client.GetAccountInfo()
	if err != nil {
		panic(err.Error())
	}
	for _, account := range resp {
		if account.State != "working" {
			panic("error account State:" + account.State)

		}
		client.Accounts[account.Type] = strconv.FormatInt(account.Id, 10)
	}

	return client
}
func (c *Client) GetC2cBalance() (*c2c.AccountBalanceData, error) {
	return c.C2client.GetC2CBalance(c.Accounts["investment"])
}
func (c *Client) BalanceOf(coin string) (balance float64, err error) {
	amount := decimal.NewFromInt(0)
	for ct, accountId := range c.Accounts {
		if ct != "spot" {
			continue
		}
		if resp, err := c.GetAccountBalance(accountId); err != nil {
			return 0, err
		} else {
			for _, banance := range resp.List {
				if banance.Currency == coin {
					tmpmnt, err := decimal.NewFromString(banance.Balance)
					if err != nil {
						log.Info(err.Error())
						return 0, err
					}
					amount = tmpmnt.Add(amount)
					break
				}
			}
		}
	}
	balance, _ = amount.Float64()
	return balance, nil
}
func (c *Client) String(f float64, decimal int) string {
	if decimal == 0 {
		return fmt.Sprintf("%v", uint64(f))
	} else if decimal <= 10 {
		return fmt.Sprintf(fmt.Sprintf("%%.%vf", decimal), f)
	} else {
		return fmt.Sprintf("%.10f", f)
	}
}
func (c *Client) Price(symbol string) (buyPrice float64, sellPrice float64, err error) {
	if ticker, err := c.GetLast24hCandlestickAskBid(symbol); err != nil {
		return 0, 0, err
	} else {
		buyPrice, _ = ticker.Bid[0].Float64()
		sellPrice, _ = ticker.Ask[0].Float64()
		return buyPrice, sellPrice, err
	}
}
