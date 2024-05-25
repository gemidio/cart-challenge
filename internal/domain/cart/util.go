package cart

import (
	"github.com/Rhymond/go-money"
)

func newMoney(amount float64) *money.Money {
	return money.NewFromFloat(amount, money.BRL)
}
