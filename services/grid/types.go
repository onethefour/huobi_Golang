package grid

type Ctxt struct {
	BuyMount float64 //买的量 usdt 计量
	BuyPrice float64 //买的价格
	BuyOrder string  //买的订单

	//SellMount float64 //卖的量 btc 计量
	SellPrice float64 //卖的价格
	SellOrder string  //卖的订单

	TotalBaseCurrency float64 // 到此需要买的币(btc) 用于初始化
}
