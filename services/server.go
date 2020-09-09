package services

import (
	"errors"
	"sync"
)

var registor map[int]Server
var lock sync.Mutex

func init() {
	registor = make(map[int]Server)
}
func Register(Id int, rbt Server) error {
	lock.Lock()
	defer lock.Unlock()

	_, ok := registor[Id]
	if ok {
		return errors.New("robot exist")
	}
	registor[Id] = rbt
	return nil
}
func Stop(Id int) {
	lock.Lock()
	defer lock.Unlock()
	rbt, ok := registor[Id]
	if !ok {
		return
	}
	rbt.Stop()
	delete(registor, Id)
	return
}
func Remove(Id int) {
	lock.Lock()
	defer lock.Unlock()
	delete(registor, Id)
}
func Action(Id int) {
	lock.Lock()
	defer lock.Unlock()
	rbt, ok := registor[Id]
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
