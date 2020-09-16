package services

import (
	"errors"
	"sync"
)

var Servers map[int]Server
var lock sync.Mutex
type Ctxt struct {
	BuyMount float64 //买的量 usdt 计量
	BuyPrice float64 //买的价格
	BuyOrder string  //买的订单

	//SellMount float64 //卖的量 btc 计量
	SellPrice float64 //卖的价格
	SellOrder string  //卖的订单

	TotalBaseCurrency float64 // 到此需要买的币(btc) 用于初始化
}

func init() {
	Servers = make(map[int]Server)
}
func Register(Id int, rbt Server) error {
	lock.Lock()
	defer lock.Unlock()

	_, ok := Servers[Id]
	if ok {
		return errors.New("robot exist")
	}
	Servers[Id] = rbt
	return nil
}
func Stop(Id int) {
	lock.Lock()
	defer lock.Unlock()
	rbt, ok := Servers[Id]
	if !ok {
		return
	}
	rbt.Stop()
	delete(Servers, Id)
	return
}
func Remove(Id int) {
	lock.Lock()
	defer lock.Unlock()
	delete(Servers, Id)
}
func Action(Id int) {
	lock.Lock()
	defer lock.Unlock()
	rbt, ok := Servers[Id]
	if !ok {
		return
	}
	rbt.Action()
	return
}
type Server interface {
	Stop()
	Start()
	Action() error
}

