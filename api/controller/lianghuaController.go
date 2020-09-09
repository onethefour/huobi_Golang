package controller

import (
	"encoding/json"
	"huobi_Golang/common/config"
	"huobi_Golang/pkg/client"
	"huobi_Golang/services/grid"
	"log"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"huobi_Golang/api/models"
	"huobi_Golang/api/robot"
	"huobi_Golang/api/utils"
	//"huobi_Golang/models"
)

//运行服务
func init() {
	list, err := new(models.Stracy).ListRun()
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(list); i++ {
		go robot.NewGridBuy(list[i].Id, list[i].Model, list[i].AccountId, list[i].BaseCurrency, list[i].QuoteCurrency, list[i].Datas, list[i].Owner, list[i].AccessKey, list[i].SecretKey).Start()
	}
}

type LianghuaController struct {
}

func (ctl *LianghuaController) Router(r *gin.Engine) {
	group := r.Group("/lianghua")
	{
		group.POST("/get_accounts", ctl.get_accounts)
		group.POST("/get_symbols", ctl.get_symbols)
		group.POST("/get_balance", ctl.get_balance)
		group.POST("/get_price", ctl.get_price)
		group.POST("/get_order", ctl.get_order)
		group.POST("/add", ctl.add)
		group.POST("/edit", ctl.edit)
		group.POST("/list", ctl.list)
		group.POST("/start", ctl.start)
		group.POST("/stop", ctl.stop)
		group.POST("/action", ctl.action)
		group.POST("/delete", ctl.delete)
		group.POST("/get", ctl.get)
		//group.POST("/GetTradeDetail", ctl.GetTradeDetail)
		group.POST("/order_list", ctl.order_list)
		group.POST("/profit", ctl.profit)
	}
}
func (ctl *LianghuaController) profit(ctx *gin.Context) {
	var params models.LHStartParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	new(models.Statis).Compute(params.Id)

	h24 := time.Now().UnixNano()/1e6 - 24*60*60*1000
	d7 := time.Now().UnixNano()/1e6 - 7*24*60*60*1000
	profith24, _ := new(models.Orders).SumProfit(params.Id, h24)
	profitd7, _ := new(models.Orders).SumProfit(params.Id, d7)
	totalProfit := float64(0)
	if statis, _ := new(models.Statis).Get(params.Id); statis != nil {
		totalProfit = statis.Totalfee
	}
	ctx.JSON(200, gin.H{"code": 0, "data": gin.H{"h24": profith24, "d7": profitd7, "all": totalProfit}})
}
func (ctl *LianghuaController) order_list(ctx *gin.Context) {

	//new(models.Orders).Flush(19)
	err := new(models.Statis).Compute(19)
	if err != nil {
		log.Println(err.Error())
	}
	infos, _ := new(models.Orders).List(19)
	ctx.JSON(200, gin.H{"code": 1, "data": infos})
}
func (ctl *LianghuaController) get_order(ctx *gin.Context) {
	var params models.LHgetOrderParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	stracy := new(models.Stracy)
	if has, _ := stracy.Get(params.Id); !has {
		ctx.JSON(200, gin.H{"code": 1, "message": "data not exist"})
		return
	}
	//huobi := models.NewHuoBiEx(stracy.AccessKey, stracy.SecretKey)
	//order := huobi.GetOrder(params.Orderid)
	huobi := client.NewClient(stracy.AccessKey, stracy.SecretKey, "")
	order, err := huobi.GetOrderById(params.Orderid)
	ctx.JSON(200, gin.H{"code": 1, "data": order, "error": err})
}

func (ctl *LianghuaController) get(ctx *gin.Context) {
	var params models.LHStartParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	info := new(models.Stracy)
	has, err := info.Get(params.Id)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	if !has {
		ctx.JSON(200, gin.H{"code": 1, "message": "data not exist"})
		return
	}
	ctx.JSON(200, gin.H{"code": 0, "data": info})
}
func (ctl *LianghuaController) start(ctx *gin.Context) {
	var params models.LHStartParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	robot.Stop(params.Id)
	info := new(models.Stracy)
	has, err := info.Get(params.Id)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	if !has {
		ctx.JSON(200, gin.H{"code": 1, "message": "data not exist"})
		return
	}
	new(models.Stracy).UpdateRun(params.Id, 1)
	go robot.NewGridBuy(info.Id, info.Model, info.AccountId, info.BaseCurrency, info.QuoteCurrency, info.Datas, info.Owner, info.AccessKey, info.SecretKey).Start()
	ctx.JSON(200, gin.H{"code": 0, "message": "success"})
}
func (ctl *LianghuaController) stop(ctx *gin.Context) {
	var params models.LHStartParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	robot.Stop(params.Id)
	new(models.Stracy).UpdateRun(params.Id, 0)
	ctx.JSON(200, gin.H{"code": 0, "message": "success"})
}
func (ctl *LianghuaController) action(ctx *gin.Context) {
	var params models.LHStartParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	robot.Action(params.Id)
	ctx.JSON(200, gin.H{"code": 0, "message": "success"})
}
func (ctl *LianghuaController) delete(ctx *gin.Context) {
	var params models.LHStartParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	robot.Stop(params.Id)
	if err := new(models.Stracy).Delete(params.Id); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"code": 0, "message": "success"})
}

func (ctl *LianghuaController) list(ctx *gin.Context) {
	list, err := new(models.Stracy).List("")
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 0, "data": list})
}
func (ctl *LianghuaController) edit(ctx *gin.Context) {
	var params models.LHEditdParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	//停止服务
	robot.Stop(params.Id)
	oldStracy := new(models.Stracy)
	has, err := oldStracy.Get(params.Id)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	if !has {
		ctx.JSON(200, gin.H{"code": 1, "message": "data not exist"})
		return
	}
	params.Run = oldStracy.Run

	oldCtxts := make([]*robot.Ctxt, 0)
	if err := json.Unmarshal([]byte(oldStracy.Datas), &oldCtxts); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	//log.Println(oldStracy.Datas)
	//取消order
	oldHuobi := client.NewClient(oldStracy.AccessKey, oldStracy.SecretKey, "")
	newHuobi := client.NewClient(params.AccessKey, params.SecretKey, "")
	for i := 0; i < len(oldCtxts); i++ {
		//log.Println(oldCtxts[i])
		if oldCtxts[i].BuyOrder != "" {
			oldHuobi.CancelOrderById(oldCtxts[i].BuyOrder)
			newHuobi.CancelOrderById(oldCtxts[i].BuyOrder)
		}
		if oldCtxts[i].SellOrder != "" {
			oldHuobi.CancelOrderById(oldCtxts[i].SellOrder)
			newHuobi.CancelOrderById(oldCtxts[i].SellOrder)
		}
	}
	Price := params.MaxPrice
	newCtxts := make([]*robot.Ctxt, 0)
	TotalBaseCurrency := float64(0)
	for i := 0; i < 100 && Price > params.MinPrice; i++ {
		tctxt := new(robot.Ctxt)
		tctxt.BuyMount = params.Amount
		tctxt.BuyPrice = utils.Digits(params.MaxPrice*math.Pow((100-params.HeightPrice)/100, float64(i)), 4)
		tctxt.SellPrice = utils.Digits(tctxt.BuyPrice*(100+params.SellPrice)/100, 4)
		tctxt.TotalBaseCurrency = TotalBaseCurrency + utils.Digits(params.Amount/tctxt.BuyPrice, 4)
		TotalBaseCurrency = tctxt.TotalBaseCurrency
		Price = tctxt.BuyPrice
		newCtxts = append(newCtxts, tctxt)
	}
	datas, _ := json.Marshal(newCtxts)
	params.Datas = string(datas)
	if err := new(models.Stracy).Update(params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"code": 0, "message": "success"})
	//启动服务
	if params.Run == 1 {
		go grid.NewGrid(params.Id, params.Model, params.AccountId, params.BaseCurrency, params.QuoteCurrency, params.Datas, params.Owner, params.AccessKey, params.SecretKey).Start()
	}
	//newCtxts,覆盖 oldCtxts
}
func (ctl *LianghuaController) add(ctx *gin.Context) {
	var params models.LHAaddParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	Price := params.MaxPrice
	ctxts := make([]*grid.Ctxt, 0)
	TotalBaseCurrency := float64(0)
	for i := 0; i < 100 && Price > params.MinPrice; i++ {
		tctxt := new(grid.Ctxt)
		tctxt.BuyMount = params.Amount
		tctxt.BuyPrice = utils.Digits(params.MaxPrice*math.Pow((100-params.HeightPrice)/100, float64(i)), 4)
		tctxt.SellPrice = utils.Digits(tctxt.BuyPrice*(100+params.SellPrice)/100, 4)
		tctxt.TotalBaseCurrency = TotalBaseCurrency + utils.Digits(params.Amount/tctxt.BuyPrice, 4)
		TotalBaseCurrency = tctxt.TotalBaseCurrency
		Price = tctxt.BuyPrice
		ctxts = append(ctxts, tctxt)
	}
	datas, _ := json.Marshal(ctxts)
	params.Datas = string(datas)
	if err := new(models.Stracy).Add(params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"code": 0, "message": "success"})
}

func (ctl *LianghuaController) get_symbols(ctx *gin.Context) {
	client := new(client.CommonClient).Init(config.Host)
	ret, err := client.GetSymbols()
	ctx.JSON(200, gin.H{"status": "ok", "data": ret, "err-msg": err})
}

func (ctl *LianghuaController) GetTradeDetail(ctx *gin.Context) {
	//huobi := models.NewHuoBiEx("", "")
	huobi := new(client.MarketClient).Init(config.Host)
	ret, _ := huobi.GetLatestTrade("ethusdt")
	ctx.JSON(200, ret)
}

func (ctl *LianghuaController) get_accounts(ctx *gin.Context) {
	var params models.GetAccountsParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	huobi := new(client.AccountClient).Init(params.AccessKey, params.SecretKey, "")
	accounts, err := huobi.GetAccountInfo()
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 0, "data": accounts})
	return
}
func (ctl *LianghuaController) get_balance(ctx *gin.Context) {
	var params models.GetBalanceParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	grid := robot.NewGridBuy(1, 1, params.AccountId, "btc", "usdt", "[]", "", params.AccessKey, params.SecretKey)
	balance, err := grid.BalanceOf(params.Currency)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"code": 0, "data": balance})
	return
}
func (ctl *LianghuaController) get_price(ctx *gin.Context) {
	var params models.GetPriceParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	grid := robot.NewGridBuy(1, 1, params.AccountId, params.BaseCurrency, params.QuoteCurrency, "[]", "", params.AccessKey, params.SecretKey)
	price, err := grid.BuyPrice()
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"code": 0, "data": price})
	return
}
