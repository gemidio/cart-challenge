package cart

import (
	"errors"
	"slices"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

const (
	percentageCoupon = "percentage"
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

func NoDiscount(rule string) Discount {
	discount, _ := NewDiscount(rule, 0)

	return discount
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
	coupon   string
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

func (c *Cart) AddPromotionalCoupon(label string) {
	// c.promotionalCoupon = label
}

func NewCoupon(label string, ctype string, minimumAmount float64, value float64, expireAt time.Time) (Coupon, error) {

	if slices.Contains([]string{"percentage", "money"}, ctype) {
		return Coupon{}, errors.New("coupon type is invalid")
	}

	return Coupon{label, ctype, minimumAmount, value, expireAt}, nil
}

type Coupon struct {
	label         string
	ctype         string
	minimumAmount float64
	value         float64
	expireAt      time.Time
}

func (c Coupon) isApplicable(amount *money.Money) bool {
	return c.minimumAmount <= float64(amount.Amount())
}

func (c Coupon) isPercentage() bool {
	return c.ctype == percentageCoupon
}

func (c Coupon) Percentage() float64 {

	if !c.isPercentage() {
		return float64(0)
	}

	return c.value / 100
}

func (c Coupon) isExpired() bool {
	return c.expireAt.Before(time.Now())
}

func (c Coupon) Label() string {
	return c.label
}
