package models

import (
	"fmt"
	"huobi_Golang/pkg/client"
	"log"
	"strconv"

	"huobi_Golang/api/utils"
)

func init() {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		panic(err.Error())
	}
	defer engineScan.Close()
	err = engineScan.Sync2(new(Orders))
	if err != nil {
		panic(err.Error())
	}
	//log.Println("init")
}

type Orders struct {
	OrderId  int    `xorm:"not null pk BIGINT(20)"`
	StracyId int    `xorm:"default 0 BIGINT(20)"`
	Type     string `xorm:"CHAR(66)"`
	//订单价格
	Price float64 `xorm:"Numeric"`
	//已成交数量
	Fieldamount float64 `xorm:"Numeric"`
	//已成交总金额
	Fieldcashamount float64 `xorm:"Numeric"`
	//收益
	Profit float64 `xorm:"Numeric"`
	//订单变为终结态的时间，不是成交时间，包含“已撤单”状态
	Finishedat uint64 `xorm:"BIGINT(20)"`
	//submitted 已提交, partial-filled 部分成交, partial-canceled 部分成交撤销, filled 完全成交, canceled 已撤销， created
	State string `xorm:"CHAR(66)"`
	//更新时间
	Updatetime uint64 `xorm:"BIGINT(20)"`
}

func (m *Orders) Add(orderid string, stracyid int) error {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return err
	}
	defer engineScan.Close()
	tm := new(Orders)
	orderid_int64, err := strconv.ParseInt(orderid, 10, 64)
	if err != nil {
		return err
	}
	tm.OrderId = int(orderid_int64)
	tm.StracyId = stracyid
	_, err = engineScan.Insert(tm)
	return err
}

//List
func (m *Orders) List(starcyid int) ([]*Orders, error) {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return nil, err
	}
	defer engineScan.Close()
	infos := make([]*Orders, 0)
	err = engineScan.SQL("select * from " + m.TableName() + " where 1").Find(&infos)
	return infos, err
}

//List
func (m *Orders) ListPending(starcyid int) ([]*Orders, error) {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return nil, err
	}
	defer engineScan.Close()
	infos := make([]*Orders, 0)
	err = engineScan.SQL("select * from "+m.TableName()+" where Finishedat=0 and stracy_id=?", starcyid).Find(&infos)
	return infos, err
}
func (m *Orders) ListCompute(starcyid int, Finishedat uint64) ([]*Orders, error) {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return nil, err
	}
	defer engineScan.Close()
	// engineScan.ShowSQL(true)
	// defer engineScan.ShowSQL(false)
	infos := make([]*Orders, 0)
	err = engineScan.SQL("select * from "+m.TableName()+" where finishedat>? and stracy_id=? order by finishedat asc", Finishedat, starcyid).Find(&infos)
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(infos)
	return infos, err
}

//更新order信息
func (m *Orders) Flush(stracyid int) error {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return err
	}
	defer engineScan.Close()

	// engineScan.ShowSQL(true)
	// defer engineScan.ShowSQL(false)
	infos, err := m.ListPending(stracyid)
	if err != nil || len(infos) == 0 {
		return nil
	}
	stracy := new(Stracy)
	if has, _ := stracy.Get(stracyid); !has {
		return nil
	}
	//huobi := NewHuoBiEx(stracy.AccessKey, stracy.SecretKey)
	huobi := new(client.OrderClient).Init(stracy.AccessKey, stracy.SecretKey, "")
	for i := 0; i < len(infos); i++ {
		order, err := huobi.GetOrderById(fmt.Sprint(infos[i].OrderId))
		if err != nil {
			continue
		}
		//if order.Data.Finishedat > 0 {
		//	engineScan.Exec("delete from "+m.TableName()+" where order_id=?", infos[i].OrderId)
		//	continue
		//}
		if order.Data.Finishedat > 0 {
			infos[i].Finishedat = order.Data.Finishedat
			infos[i].Fieldamount, _ = strconv.ParseFloat(order.Data.FilledAmount, 64)
			infos[i].Fieldcashamount, _ = strconv.ParseFloat(order.Data.FilledCashAmount, 64)
			fee, _ := strconv.ParseFloat(order.Data.FilledFees, 64)
			if order.Data.Type == "sell-limit" || order.Data.Type == "sell-market" {
				infos[i].Type = "sell"
				infos[i].Fieldcashamount -= fee

			} else {
				infos[i].Type = "buy"
				infos[i].Fieldamount -= fee
			}
			infos[i].State = order.Data.State
			err := infos[i].Update()
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
	//清除已经取消的交易
	_, err = engineScan.Exec("delete from " + m.TableName() + " where finishedat>0 and `fieldamount`<0.000000001 and `fieldcashamount`<0.000000001")
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}
func (m *Orders) Update() error {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return err
	}
	defer engineScan.Close()
	_, err = engineScan.Exec("UPDATE `orders` SET `fieldamount` = ?, `fieldcashamount` = ?, `finishedat` = ?, `state` = ?,`type` = ? where `order_id` = ?", m.Fieldamount, m.Fieldcashamount, m.Finishedat, m.State, m.Type, m.OrderId)
	return err
}
func (m *Orders) UpdateFee(OrderId int, Fee float64) error {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return err
	}
	defer engineScan.Close()
	_, err = engineScan.Exec("UPDATE `orders` SET `profit` = ? where `order_id` = ?", Fee, OrderId)
	return err
}
func (m *Orders) SumProfit(stracyid int, starts int64) (float64, error) {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return 0, err
	}
	defer engineScan.Close()
	//engineScan.ShowSQL(true)
	//defer engineScan.ShowSQL(false)
	return engineScan.Where("stracy_id=? and finishedat >?", stracyid, starts).Sum(m, "profit")
}

// TableName a
func (m *Orders) TableName() string {
	return "orders"
}
