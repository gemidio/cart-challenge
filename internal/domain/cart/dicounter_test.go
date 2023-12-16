package cart

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTakeThreePayTwo(t *testing.T) {
	discounter := TakeThreePayTwoDiscounter{
		&mockProductOfferRepository{},
	}

	t.Run("get discount when cart's subtotal is invalid", func(t *testing.T) {
		cart := newMockCart(
			[]Item{
				{uuid.New(), uuid.New(), *newMoney(-10), 1},
			},
			Discount{"none", *newMoney(0)},
		)

		discount := discounter.calculate(cart)

		assert.Equal(t, "take-3-pay-2", discount.rule)
		assert.True(t, discount.IsZero())
	})

	t.Run("get discount when item is not offer", func(t *testing.T) {
		cart := newMockCart(
			[]Item{
				{uuid.New(), uuid.New(), *newMoney(10), 3},
			},
			Discount{"none", *newMoney(0)},
		)

		discount := discounter.calculate(cart)

		assert.Equal(t, "take-3-pay-2", discount.rule)
		assert.True(t, discount.IsZero())
	})

	t.Run("get discount when item is offer", func(t *testing.T) {
		cart := mockCartWithItemOffer(3)

		expect := newMoney(10)

		discount := discounter.calculate(cart)
		isEquals, _ := discount.value.Equals(expect)

		assert.Equal(t, "take-3-pay-2", discount.rule)
		assert.True(t, isEquals)
	})

	t.Run("get discount when item is offer but quantity is insufficient", func(t *testing.T) {
		cart := mockCartWithItemOffer(2)

		discount := discounter.calculate(cart)

		assert.Equal(t, "take-3-pay-2", discount.rule)
		assert.True(t, discount.IsZero())
	})
}

type mockProductOfferRepository struct{}

func (r *mockProductOfferRepository) IsOffer(id uuid.UUID) bool {
	return id.String() == "5e96a677-6aef-4ca6-876c-ef289e936440"
}

func mockCartWithItemOffer(quantity int) Cart {
	return newMockCart(
		[]Item{
			{
				// ID on offer
				id:          uuid.MustParse("5e96a677-6aef-4ca6-876c-ef289e936440"),
				cartegoryId: uuid.New(),
				price:       *newMoney(10),
				quantity:    quantity,
			},
		},
		Discount{"none", *newMoney(0)},
	)
}
