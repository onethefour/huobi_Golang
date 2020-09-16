package start

import (
	"huobi_Golang/api/models"
	"huobi_Golang/services/grid"
	"huobi_Golang/services"
)

func NewServer(stracy *models.Stracy) (services.Server,error) {
	return  grid.NewGrid(stracy)
}
