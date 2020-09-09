package models

import (
	"encoding/json"
	"log"

	"huobi_Golang/api/utils"
)

func init() {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		panic(err.Error())
	}
	defer engineScan.Close()
	err = engineScan.Sync2(new(Statis))
	if err != nil {
		panic(err.Error())
	}
	//log.Println("init")
}

type Statis struct {
	StracyId int     `xorm:"not null pk BIGINT(20)"`
	Context  string  `xorm:"CHAR(66)"`
	Totalfee float64 `xorm:"Numeric"`
	Lastime  uint64  `xorm:"BIGINT(20)"`
}
type Context struct {
	Price float64 //
	Base  float64 //btc
	Quoa  float64 //usdt
}

//计算卖单收益
func (c *Context) Fee(ctxts []*Context, Base, Quoa float64) ([]*Context, float64) {
	fee := float64(0)
	if len(ctxts) == 0 {
		return ctxts, Quoa / 100
	}
	cost := float64(0)
	for i := len(ctxts) - 1; i >= 0; i-- {
		if Base > ctxts[i].Base {
			cost += ctxts[i].Quoa
			Base -= ctxts[i].Base
			ctxts = ctxts[0:i]
		} else {
			cost += ctxts[i].Quoa * Base / ctxts[i].Base
			ctxts[i].Quoa = ctxts[i].Quoa * (ctxts[i].Base - Base) / ctxts[i].Base
			ctxts[i].Base -= Base
			Base = 0
			break
		}
	}

	if Base > 0.00000001 {
		fee = Base / 100
	}
	fee += Quoa - cost
	if fee < 0 {
		log.Println("你怕是个智障")
	}
	return ctxts, fee
}

func (c *Context) Add(ctxts []*Context, Base, Quoa float64) []*Context {
	tmpctxt := new(Context)
	tmpctxt.Price = Quoa / Base
	tmpctxt.Base = Base
	tmpctxt.Quoa = Quoa
	ctxts = append(ctxts, tmpctxt)

	return c.Desc(ctxts)
}

//Desc 价格降序排列
func (c *Context) Desc(ctxts []*Context) []*Context {
	for i := len(ctxts) - 1; i > 0; i-- {
		if ctxts[i-1].Price < ctxts[i].Price {
			ctxts[i-1].Price, ctxts[i].Price = ctxts[i].Price, ctxts[i-1].Price
		} else {
			break
		}
	}
	return ctxts
}
func (m *Statis) Get(StracyId int) (*Statis, error) {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return nil, err
	}
	defer engineScan.Close()
	//engineScan.ShowSQL(true)
	//defer engineScan.ShowSQL(false)
	db := new(Statis)
	db.StracyId = StracyId
	has, err := engineScan.Get(db)
	if err != nil {
		return nil, err
	}
	if !has {
		db.StracyId = StracyId
		db.Context = "[]"
		db.Lastime = 0
		if _, err = engineScan.Insert(db); err != nil {
			return nil, err
		}
		return db, nil
	}
	return db, nil
}

func (m *Statis) Compute(StracyId int) error {
	//
	if err := new(Orders).Flush(StracyId); err != nil {
		return err
	}
	statis, err := m.Get(StracyId)
	if err != nil {
		return err
	}

	orders, _ := new(Orders).ListCompute(StracyId, statis.Lastime)
	if orders == nil || len(orders) == 0 {
		return nil
	}
	log.Println(orders)
	ctxs := make([]*Context, 0, 10)
	log.Println(ctxs)
	err = json.Unmarshal([]byte(statis.Context), &ctxs)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	//准备好了开始处理
	for i := 0; i < len(orders); i++ {
		if statis.Lastime < orders[i].Finishedat {
			statis.Lastime = orders[i].Finishedat
		}
		if orders[i].Type == "buy" {
			ctxs = new(Context).Add(ctxs, orders[i].Fieldamount, orders[i].Fieldcashamount)
			continue
		}
		if orders[i].Type == "sell" {
			var fee float64
			ctxs, fee = new(Context).Fee(ctxs, orders[i].Fieldamount, orders[i].Fieldcashamount)
			orders[i].UpdateFee(orders[i].OrderId, fee)
			statis.Totalfee += fee
			continue
		}
	}
	ctxsJson, _ := json.Marshal(ctxs)
	statis.Context = string(ctxsJson)
	statis.Update()
	return nil
}
func (m *Statis) Update() error {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return err
	}
	defer engineScan.Close()
	_, err = engineScan.Exec("UPDATE "+m.TableName()+" SET `context` = ?, `lastime` = ?, `totalfee` = ? where `stracy_id` = ?", m.Context, m.Lastime, m.Totalfee, m.StracyId)
	return err
}
func (m *Statis) TableName() string {
	return "statis"
}
