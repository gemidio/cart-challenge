package cart

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCart(t *testing.T) {
	cart := Cart{
		id:     uuid.New(),
		userId: uuid.New(),
		discount: Discount{
			"some-discount",
			*money.NewFromFloat(5, money.BRL),
		},
		items: []Item{
			{uuid.New(), uuid.New(), *money.NewFromFloat(5, money.BRL), 1},
			{uuid.New(), uuid.New(), *money.NewFromFloat(5, money.BRL), 1},
		},
	}

	t.Run("get subtotal", func(t *testing.T) {
		subtotal := cart.Subtotal()
		expect := money.NewFromFloat(10, money.BRL)

		isEquals, _ := subtotal.Equals(expect)

		assert.True(t, isEquals)
	})

	t.Run("get total", func(t *testing.T) {
		total := cart.Total()
		expect := money.NewFromFloat(5, money.BRL)

		isEquals, _ := total.Equals(expect)

		assert.True(t, isEquals)
	})
}
