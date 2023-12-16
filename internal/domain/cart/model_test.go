package cart

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCart(t *testing.T) {
	cart := newMockCart(
		[]Item{
			{uuid.New(), uuid.New(), *newMoney(5), 1},
			{uuid.New(), uuid.New(), *newMoney(5), 1},
		},
		Discount{
			"some-discount",
			*newMoney(5),
		},
	)

	t.Run("get subtotal", func(t *testing.T) {
		subtotal, err := cart.Subtotal()
		expect := newMoney(10)

		isEquals, _ := subtotal.Equals(expect)

		assert.NoError(t, err)
		assert.True(t, isEquals)
	})

	t.Run("get total", func(t *testing.T) {
		total, err := cart.Total()
		expect := newMoney(5)

		isEquals, _ := total.Equals(expect)

		assert.NoError(t, err)
		assert.True(t, isEquals)
	})
}

func TestCartErrNegativeSubtotal(t *testing.T) {

	cart := newMockCart(
		[]Item{
			{uuid.New(), uuid.New(), *newMoney(-1), 1},
		},
		Discount{
			"some-discount",
			*newMoney(5),
		},
	)

	t.Run("get negative subtotal", func(t *testing.T) {
		_, err := cart.Subtotal()

		assert.ErrorIs(t, err, ErrNegativeSubtotal)
	})

	t.Run("get negative total", func(t *testing.T) {
		_, err := cart.Total()

		assert.ErrorIs(t, err, ErrNegativeSubtotal)
	})
}

func TestCartErrNegativeTotal(t *testing.T) {
	cart := newMockCart(
		[]Item{
			{uuid.New(), uuid.New(), *newMoney(3), 1},
		},
		Discount{
			"some-discount",
			*newMoney(5),
		},
	)

	_, err := cart.Total()

	assert.Error(t, err, ErrNegativeTotal)
}

func TestNewDiscount(t *testing.T) {

	t.Run("get valid discount", func(t *testing.T) {
		discount, _ := NewDiscount("some-discount", 10)
		expectTotal := money.NewFromFloat(10, money.BRL)

		isEquals, _ := discount.value.Equals(expectTotal)

		assert.Equal(t, "some-discount", discount.rule)
		assert.True(t, isEquals)
	})

	t.Run("get invalid discount error", func(t *testing.T) {
		_, err := NewDiscount("some-discount", -10)

		assert.Error(t, err, ErrInvalidDiscount)
	})

	t.Run("get discount when give negative value", func(t *testing.T) {
		discount, _ := NewDiscount("some-discount", -1)

		assert.True(t, discount.IsZero())
	})
}

func newMockCart(items []Item, discount Discount) Cart {
	return Cart{
		id:       uuid.New(),
		userId:   uuid.New(),
		discount: discount,
		items:    items,
	}
}
