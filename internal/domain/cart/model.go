package cart

import (
	"errors"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

var (
	// ErrInvalidSubtotal happens when cart's subtotal is negataive
	ErrNegativeSubtotal = errors.New("subtotal is negative")

	// ErrInvalidTotal happens when cart's total is negataive
	ErrNegativeTotal = errors.New("total is negative")

	// ErrInvalidDiscount happens when total is less than zero
	ErrInvalidDiscount = errors.New("discount is invalid")
)

func NewDiscount(rule string, value float64) (Discount, error) {
	valueMoney := *newMoney(value)

	if valueMoney.IsNegative() {
		return Discount{
			rule,
			*newMoney(0),
		}, ErrInvalidDiscount
	}

	return Discount{rule, valueMoney}, nil
}

type Discount struct {
	rule  string
	value money.Money
}

func (d Discount) IsZero() bool {
	return d.value.IsZero()
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

	total, _ := subtotal.Subtract(&c.discount.value)

	if total.IsNegative() {
		return &money.Money{}, ErrNegativeTotal
	}

	return total, nil
}

type ItemOffer struct {
	id             uuid.UUID
	itemId         uuid.UUID
	targetQuantity int
	chargeQuantity int
	expireAt       time.Time
}

func (io ItemOffer) isApplicable(item Item) bool {
	return item.id == io.itemId && item.quantity >= int(io.targetQuantity)
}

func (io ItemOffer) isExpired() bool {
	return time.Now().After(io.expireAt)
}

func (io ItemOffer) discountApplyTo(item Item) int {

	rest := io.targetQuantity - io.chargeQuantity

	return (item.quantity / io.targetQuantity) * rest
}
