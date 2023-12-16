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
			{uuid.New(), uuid.New(), *money.NewFromFloat(5, money.BRL), 1},
			{uuid.New(), uuid.New(), *money.NewFromFloat(5, money.BRL), 1},
		},
		Discount{
			"some-discount",
			*money.NewFromFloat(5, money.BRL),
		},
	)

	t.Run("get subtotal", func(t *testing.T) {
		subtotal, err := cart.Subtotal()
		expect := money.NewFromFloat(10, money.BRL)

		isEquals, _ := subtotal.Equals(expect)

		assert.NoError(t, err)
		assert.True(t, isEquals)
	})

	t.Run("get total", func(t *testing.T) {
		total, err := cart.Total()
		expect := money.NewFromFloat(5, money.BRL)

		isEquals, _ := total.Equals(expect)

		assert.NoError(t, err)
		assert.True(t, isEquals)
	})
}

func TestCartErrNegativeSubtotal(t *testing.T) {

	cart := newMockCart(
		[]Item{
			{uuid.New(), uuid.New(), *money.NewFromFloat(-1, money.BRL), 1},
		},
		Discount{
			"some-discount",
			*money.NewFromFloat(5, money.BRL),
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

func newMockCart(itens []Item, discount Discount) Cart {
	return Cart{
		id:       uuid.New(),
		userId:   uuid.New(),
		discount: discount,
		itens:    itens,
	}
}
