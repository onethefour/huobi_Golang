package models

type LoginParams struct {
	Account string `form:"account" binding:"required"`
	Pwd     string `form:"pwd" binding:"required"`
}

type LoinParams struct {
	//Islogin参数会被CheckLogin中间件自动加上,
	//有这个参数表示接口需要登录,否则无需登录
	Islogin int64 `form:"islogin" binding:"-"`
}
type GetAccountParams struct {
	//Islogin参数会被CheckLogin中间件自动加上,
	//有这个参数表示接口需要登录,否则无需登录
	Islogin int64 `form:"islogin" binding:"-"`
}

type SelectParams struct {
	Islogin int64 `form:"islogin" binding:"required"`
	Start   int64 `form:"start" binding:"-" validate:"min=1"`
	End     int64 `form:"end" binding:"-" validate:"min=1"`
}

type AddParams struct {
	Owner         string `form:"Owner"`
	BaseCurrency  string `form:"BaseCurrency"`
	QuoteCurrency string `form:"QuoteCurrency"`
	Model         uint64 `form:"Model"`
}

type GetAccountsParams struct {
	AccessKey string `form:"AccessKey" binding:"required"`
	SecretKey string `form:"SecretKey" binding:"required"`
}

type GetBalanceParams struct {
	AccessKey string `form:"AccessKey" binding:"required"`
	SecretKey string `form:"SecretKey" binding:"required"`
	Currency  string `form:"Currency" binding:"required"`
	AccountId string `form:"AccountId" binding:"required"`
}
type GetPriceParams struct {
	AccessKey     string `form:"AccessKey" binding:"required"`
	SecretKey     string `form:"SecretKey" binding:"required"`
	AccountId     string `form:"AccountId" binding:"required"`
	QuoteCurrency string `form:"QuoteCurrency" binding:"required"`
	BaseCurrency  string `form:"BaseCurrency" binding:"required"`
}
type LHAaddParams struct {
	Owner         string
	Capital       float64 `form:"Capital" binding:"required"`
	AccessKey     string  `form:"AccessKey" binding:"required"`
	SecretKey     string  `form:"SecretKey" binding:"required"`
	AccountId     string  `form:"AccountId" binding:"required"`
	QuoteCurrency string  `form:"QuoteCurrency" binding:"required"`
	BaseCurrency  string  `form:"BaseCurrency" binding:"required"`
	Name          string  `form:"Name" binding:"required"`

	Datas       string
	MinPrice    float64 `form:"MinPrice" binding:"required"`
	MaxPrice    float64 `form:"MaxPrice" binding:"required"`
	HeightPrice float64 `form:"HeightPrice" binding:"required"`
	Percent     float64 `form:"Percent" binding:"max=10"`
	SellPrice   float64 `form:"SellPrice" binding:"required"`
	Amount      float64 `form:"Amount" binding:"required"`
	Model       int64   `form:"Model" binding:"required"`
}
type LHEditdParams struct {
	Id            int `form:"Id" binding:"required"`
	Owner         string
	Capital       float64 `form:"Capital" binding:"required"`
	AccessKey     string  `form:"AccessKey" binding:"required"`
	SecretKey     string  `form:"SecretKey" binding:"required"`
	AccountId     string  `form:"AccountId" binding:"required"`
	QuoteCurrency string  `form:"QuoteCurrency" binding:"required"`
	BaseCurrency  string  `form:"BaseCurrency" binding:"required"`
	Name          string  `form:"Name" binding:"required"`
	Datas         string
	Run           int64
	MinPrice      float64 `form:"MinPrice" binding:"required"`
	MaxPrice      float64 `form:"MaxPrice" binding:"required"`
	HeightPrice   float64 `form:"HeightPrice" binding:"required"`
	SellPrice     float64 `form:"SellPrice" binding:"required"`
	Amount        float64 `form:"Amount" binding:"required"`
	Model         int64   `form:"Model" binding:"required"`
}
type LHStartParams struct {
	Id int `form:"Id" binding:"required"`
}
type LHgetOrderParams struct {
	Id      int    `form:"Id" binding:"required"`
	Orderid string `form:"orderid" binding:"required"`
}
