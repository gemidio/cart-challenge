package cart

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Discount struct {
	rule  string
	total money.Money
}

type Item struct {
	id          uuid.UUID
	cartegoryId uuid.UUID
	price       money.Money
	quantity    int
}

type Cart struct {
	id       uuid.UUID
	userId   uuid.UUID
	discount Discount
	itens    []Item
}

func (c *Cart) Subtotal() *money.Money {
	subtotal := money.NewFromFloat(0, money.BRL)

	for _, item := range c.itens {
		subtotal, _ = subtotal.Add(item.price.Multiply(int64(item.quantity)))
	}

	return subtotal
}

func (c *Cart) Total() *money.Money {
	total, _ := c.Subtotal().Subtract(&c.discount.total)

	return total
}
