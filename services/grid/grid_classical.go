package grid

import (
	"huobi_Golang/api/utils"
	"huobi_Golang/common/log"
)

func (g *Grid) work_classical() (err error) {
	if err = g.classical_init(); err != nil {
		return
	}
	var affacted bool
	for g.Run {
		if affacted {
			if err = g.Save(); err != nil {
				log.Info(err.Error())
			}
		}

		if err = g.sleepOrStop(10); err != nil {
			return
		}

		if affacted, err = g.classical_clear_order(); affacted || err != nil {
			if err != nil {
				log.Info(err.Error())
			}
			continue
		}
		if affacted, err = g.classica_new_order(); err != nil {
			log.Info(err.Error())
		}
	}
	return
}

//classical_clear_order 清除已经终结的订单
func (m *Grid) classical_clear_order() (affected bool, err error) {
	for i := 0; i < len(m.Datas); i++ {
		orderId := m.Datas[i].BuyOrder
		if orderId != "" {
			orderReturn, err := m.GetOrder(orderId)
			if err != nil {
				return affected, err
			}
			if orderReturn.Data.Finishedat > 0 {
				m.Datas[i].BuyOrder = ""
				affected = true
			}
		}
		orderId = m.Datas[i].SellOrder
		if orderId != "" {
			orderReturn, err := m.GetOrder(orderId)
			if err != nil {
				return affected, err
			}
			if orderReturn.Data.Finishedat > 0 {
				m.Datas[i].SellOrder = ""
				affected = true
			}
		}
	}
	return
}
func (m *Grid) classical_init() (err error) {
	//计算各个高度需要买的btc
	for i := 0; i < len(m.Datas); i++ {
		if i == 0 {
			m.Datas[i].TotalBaseCurrency = m.Datas[i].BuyMount / m.Datas[i].BuyPrice
		} else {
			m.Datas[i].TotalBaseCurrency = m.Datas[i-1].TotalBaseCurrency + m.Datas[i].BuyMount/m.Datas[i].BuyPrice
		}
	}
	var buyPrice, sellPrice, buyNeedBtc, sellNeedBtc, nowBtc float64
	if buyPrice, sellPrice, err = m.Price(); err != nil {
		log.Info(err.Error())
		return err
	}

	//当前价格所在的高度至少需要买多少btc
	for i := 0; i < len(m.Datas); i++ {
		if m.Datas[i].BuyPrice > buyPrice {
			buyNeedBtc = m.Datas[i].TotalBaseCurrency
		} else {
			break
		}
	}
	//当前价格所在的高度至多需要买多少btc
	for i := 0; i < len(m.Datas); i++ {
		if m.Datas[i].SellPrice > sellPrice {
			sellNeedBtc = m.Datas[i].TotalBaseCurrency
		} else {
			break
		}
	}
	//当前账户上btc
	if nowBtc, err = m.Client.BalanceOf(m.BaseCurrency); err != nil {
		log.Info(err.Error())
		return err
	}
	//log.Panicln("needBtc", needBtc, "nowBtc", nowBtc, m.Datas[0].TotalBaseCurrency)
	//补些btc,一般是首次启动初始化
	if nowBtc < buyNeedBtc-(0.5*m.Datas[0].TotalBaseCurrency) {
		if err = m.BuyMarket((buyNeedBtc - nowBtc) * buyPrice); err != nil {
			return err
		}

		if nowBtc, err = m.BalanceOf(m.BaseCurrency); err != nil {
			log.Info(err.Error())
			return err
		}
	}
	//卖掉多出的btc
	if nowBtc > sellNeedBtc+(0.5*m.Datas[0].TotalBaseCurrency) {
		if err = m.SellMarket(nowBtc - sellNeedBtc); err != nil {
			log.Info(err.Error())
			return err
		}
		if nowBtc, err = m.BalanceOf(m.BaseCurrency); err != nil {
			log.Info(err.Error())
			return err
		}
	}
	//策略改动后,之前的订单需要重新下单
	for i := 0; i < len(m.Datas); i++ {
		orderId := m.Datas[i].BuyOrder
		orderPrice := m.Datas[i].BuyPrice
		orderAmount := m.Datas[i].BuyMount / m.Datas[i].BuyPrice
		//m.Datas[i].BuyOrder = ""
		if orderId == "" {
			orderId = m.Datas[i].SellOrder
			orderPrice = m.Datas[i].SellPrice
			//m.Datas[i].SellOrder = ""
		}
		if orderId != "" {
			orderReturn, err := m.GetOrder(orderId)
			if err != nil {
				log.Info(err.Error())
				return err
			}
			if orderReturn.Data.Price_f > orderPrice*1.005 || orderReturn.Data.Price_f < orderPrice*0.996 {
				m.CancelOrder(orderId)
				continue
			}
			if orderReturn.Data.Amount_f > orderAmount*1.005 || orderReturn.Data.Amount_f < orderAmount*0.996 {
				m.CancelOrder(orderId)
				continue
			}
		}
	}
	return nil
}

//下买单和卖单
func (m *Grid) classica_new_order() (affected bool, err error) {
	var buyPrice, sellPrice, buyNeedBtc, sellNeedBtc, nowBtc float64
	//需要下卖单的的最低高度
	sellIndex := -1
	//需要下买单的最高高度
	buyIndex := -1
	//当前价格
	if buyPrice, err = m.BuyPrice(); err != nil {
		log.Info(err.Error())
		return affected, err
	}
	//当前价格
	if sellPrice, err = m.SellPrice(); err != nil {
		log.Info(err.Error())
		return affected, err
	}
	if nowBtc, err = m.BalanceOf(m.BaseCurrency); err != nil {
		log.Info(err.Error())
		return affected, err
	}

	//var needBtc float64
	for i := 0; i < len(m.Datas); i++ {
		if buyPrice < m.Datas[i].BuyPrice {
			buyIndex = i
			buyNeedBtc = m.Datas[buyIndex].TotalBaseCurrency
		} else {
			break
		}
	}
	//var needBtc float64
	for i := 0; i < len(m.Datas); i++ {
		if sellPrice < m.Datas[i].SellPrice {
			sellIndex = i
			sellNeedBtc = m.Datas[sellIndex].TotalBaseCurrency
		} else {
			break
		}
	}
	//补些btc
	if nowBtc < buyNeedBtc-(0.5*m.Datas[0].TotalBaseCurrency) && m.Datas[buyIndex].BuyOrder == "" {
		if err = m.BuyMarket((buyNeedBtc - nowBtc) * buyPrice); err != nil {
			log.Info(err.Error())
			return affected, err
		}
		if nowBtc, err = m.BalanceOf(m.BaseCurrency); err != nil {
			log.Info(err.Error())
			return affected, err
		}
	}
	//卖掉多出的btc
	if nowBtc > sellNeedBtc+(0.5*m.Datas[0].TotalBaseCurrency) && m.Datas[sellIndex].SellOrder == "" {
		if err = m.SellMarket(nowBtc - sellNeedBtc); err != nil {
			log.Info(err.Error())
			return affected, err
		}
		if nowBtc, err = m.BalanceOf(m.BaseCurrency); err != nil {
			log.Info(err.Error())
			return affected, err
		}
	}
	activeIndex := -1
	for i := 0; i < len(m.Datas); i++ {
		if nowBtc > m.Datas[i].TotalBaseCurrency-(0.5*m.Datas[0].TotalBaseCurrency) {
			activeIndex = i
		} else {
			break
		}
	}
	if activeIndex >= 0 {
		if m.Datas[activeIndex].SellOrder == "" {
			sellAmount := nowBtc
			if activeIndex > 0 { //四舍五入不精准,可能到时略有误差,最后一格检查是否有足够btc
				sellAmount = nowBtc - m.Datas[activeIndex-1].TotalBaseCurrency
			}
			orderid, err := m.SellOrder(utils.Digits(sellAmount, 6), m.Datas[activeIndex].SellPrice)
			if err != nil {
				log.Info("%v %v", activeIndex, sellAmount, err.Error())
				return affected, err
			}
			m.Datas[activeIndex].SellOrder = orderid
			affected = true

			for j := 0; j < len(m.Datas); j++ {
				if j != activeIndex && m.Datas[j].SellOrder != "" {
					m.CancelOrder(m.Datas[j].SellOrder)
					m.Datas[j].SellOrder = ""
				}
			}
		}
	}
	if activeIndex < len(m.Datas)-1 {
		if m.Datas[activeIndex+1].BuyOrder == "" {
			buyAmount := m.Datas[activeIndex+1].TotalBaseCurrency - nowBtc
			orderid, err := m.BuyOrder(utils.Digits(buyAmount, 6), m.Datas[activeIndex+1].BuyPrice)
			if err != nil {
				log.Info(err.Error())
				return affected, err
			}
			m.Datas[activeIndex+1].BuyOrder = orderid
			affected = true
			//删除其他订单
			for j := 0; j < len(m.Datas); j++ {
				if j != activeIndex+1 && m.Datas[j].BuyOrder != "" {
					m.CancelOrder(m.Datas[j].BuyOrder)
					m.Datas[j].BuyOrder = ""
				}
			}
		}
	}
	return affected, nil
}
