package models

import (
	"time"

	"huobi_Golang/api/utils"
)

func init() {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		panic(err.Error())
	}
	defer engineScan.Close()
	err = engineScan.Sync2(new(Stracy))
	if err != nil {
		panic(err.Error())
	}
	//log.Println("init")
}

//Stracy 策略配置
type Stracy struct {
	//自增id
	Id int `xorm:"not null pk autoincr INT(11)"`
	//本金用于统计收益
	Capital float64 `xorm:"Numeric"`
	Name    string  `xorm:"CHAR(66)"`
	//拥有者
	Owner string `xorm:"CHAR(66)"`
	//API秘钥
	AccessKey string `xorm:"CHAR(66)"`
	SecretKey string `xorm:"CHAR(66)"`
	//火币子账户id
	AccountId string `xorm:"CHAR(66)"`
	//价格区间
	MinPrice float64 `xorm:"Numeric"`
	MaxPrice float64 `xorm:"Numeric"`
	//格子高度
	HeightPrice float64 `xorm:"Numeric"`
	//平仓差价
	SellPrice float64 `xorm:"Numeric"`
	//每格量(列usdt)
	Amount float64 `xorm:"Numeric"`
	//基础币 BTC
	BaseCurrency string `xorm:"CHAR(66)"`
	//报价币 USDT
	QuoteCurrency string `xorm:"CHAR(66)"`
	//模式 1,正常模式 2,自动买手动卖 3,自动卖手动买
	Model int64 `xorm:"BIGINT(20)"`

	Datas string `xorm:"CHAR(1024)"`
	//1 启动服务 0 不启动
	Run int64 `xorm:"BIGINT(20)"`

	//更新时间
	LastUpdate int64 `xorm:"BIGINT(20)"`
}

func (m *Stracy) Delete(Id int) (err error) {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return err
	}
	defer engineScan.Close()
	tm := new(Stracy)
	tm.Id = Id
	_, err = engineScan.Delete(tm)
	return err
}

func (m *Stracy) Get(Id int) (has bool, err error) {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return false, err
	}
	defer engineScan.Close()
	m.Id = Id
	return engineScan.Get(m)
}
func (m *Stracy) Update(p LHEditdParams) error {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return err
	}
	defer engineScan.Close()
	// engineScan.ShowSQL(true)
	// defer engineScan.ShowSQL(false)
	tm := new(Stracy)

	tm.Capital = p.Capital
	tm.Owner = p.Owner
	tm.BaseCurrency = p.BaseCurrency
	tm.QuoteCurrency = p.QuoteCurrency
	tm.AccessKey = p.AccessKey
	tm.SecretKey = p.SecretKey
	tm.AccountId = p.AccountId
	tm.Name = p.Name
	tm.Datas = p.Datas
	tm.MinPrice = p.MinPrice
	tm.MaxPrice = p.MaxPrice
	tm.HeightPrice = p.HeightPrice
	tm.SellPrice = p.SellPrice
	tm.Amount = p.Amount
	tm.Model = p.Model
	tm.Run = p.Run
	tm.LastUpdate = time.Now().Unix()
	_, err = engineScan.Id(p.Id).Update(tm)
	return err
}

//Add 添加一条记录
func (m *Stracy) Add(p LHAaddParams) error {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return err
	}
	defer engineScan.Close()

	tm := new(Stracy)
	tm.Capital = p.Capital
	tm.Owner = p.Owner
	tm.BaseCurrency = p.BaseCurrency
	tm.QuoteCurrency = p.QuoteCurrency
	tm.AccessKey = p.AccessKey
	tm.SecretKey = p.SecretKey
	tm.AccountId = p.AccountId
	tm.Name = p.Name
	tm.Datas = p.Datas
	tm.MinPrice = p.MinPrice
	tm.MaxPrice = p.MaxPrice
	tm.HeightPrice = p.HeightPrice
	tm.SellPrice = p.SellPrice
	tm.Amount = p.Amount
	tm.Model = p.Model
	tm.LastUpdate = time.Now().Unix()
	_, err = engineScan.Insert(tm)
	return err
}

//List 用户策略列表
func (m *Stracy) List(Owner string) ([]*Stracy, error) {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return nil, err
	}
	defer engineScan.Close()
	infos := make([]*Stracy, 0)
	err = engineScan.SQL("select * from "+m.TableName()+" where owner=?", Owner).Find(&infos)
	return infos, err
}
func (m *Stracy) ListRun() ([]*Stracy, error) {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return nil, err
	}
	defer engineScan.Close()
	// engineScan.ShowSQL(true)
	// defer engineScan.ShowSQL(false)
	infos := make([]*Stracy, 0)
	err = engineScan.SQL("select * from " + m.TableName() + " where run=1").Find(&infos)
	return infos, err
}

//保存服务运行数据
func (m *Stracy) Save(Id int, Datas string) error {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return err
	}
	defer engineScan.Close()
	_, err = engineScan.Exec("update `"+m.TableName()+"` set datas=? where id=?", Datas, Id)
	return err
}
func (m *Stracy) UpdateRun(Id int, Run int64) error {
	engineScan, err := utils.Engine_scan()
	if err != nil {
		return err
	}
	defer engineScan.Close()
	_, err = engineScan.Exec("update `"+m.TableName()+"` set `Run`=? where Id=?", Run, Id)
	return err
}

// TableName a
func (m *Stracy) TableName() string {
	return "stracy"
}
