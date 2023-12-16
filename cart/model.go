package cart

import (
	"errors"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

var (
	// ErrInvalidSubtotal happens when cart's subtotal is negataive
	ErrNegativeSubtotal = errors.New("subtotal is negative")

	// ErrInvalidTotal happens when cart's total is negataive
	ErrNegativeTotal = errors.New("total is negative")
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
	items    []Item
}

func (c *Cart) Subtotal() (*money.Money, error) {
	subtotal := newMoney(0)

	for _, item := range c.items {
		subtotal, _ = subtotal.Add(item.price.Multiply(int64(item.quantity)))
	}

	if subtotal.IsNegative() {
		return &money.Money{}, ErrNegativeSubtotal
	}

	return subtotal, nil
}

func (c *Cart) Total() (*money.Money, error) {
	subtotal, err := c.Subtotal()

	if err != nil {
		return &money.Money{}, err
	}

	total, _ := subtotal.Subtract(&c.discount.total)

	if total.IsNegative() {
		return &money.Money{}, ErrNegativeTotal
	}

	return total, nil
}
