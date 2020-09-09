package utils

import (
	"math"
	"testing"
)

func TestGridInit(t *testing.T) {
	HeightPrice := 2.1
	Price := 0.0250
	MaxPrice := 0.0250
	MinPrice := 0.016
	SellPrice := 0.01
	for i := 0; i < 100 && Price > MinPrice; i++ {
		//BuyMount := params.Amount
		BuyPrice := Digits(MaxPrice*math.Pow((100-HeightPrice)/100, float64(i)), 4)
		SellPrice = Digits(BuyPrice*(100+SellPrice)/100, 3)
		Price = BuyPrice
		t.Log(Price, MaxPrice*math.Pow((100-HeightPrice)/100, float64(i)))
	}
	t.Log(Digits(0.020401, 3))
	return
}
