package cart

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTakeMorePayLess(t *testing.T) {
	discounter := TakeMorePayLess{
		&itemOfferRepositoryStub{},
	}

	t.Run("calculate discount when cart's subtotal is invalid", func(t *testing.T) {
		cart := newMockCart(
			[]Item{
				{uuid.New(), uuid.New(), *newMoney(-10), 1},
			},
			Discount{"none", *newMoney(0)},
		)

		actual := discounter.calculate(cart)

		assert.Equal(t, "take-more-pay-less", actual.rule)
		assert.True(t, actual.IsZero())
	})

	t.Run("calculate discount when cart's item offer is not found", func(t *testing.T) {
		cart := newMockCartWithoutItemOffer()

		actual := discounter.calculate(cart)

		assert.Equal(t, "take-more-pay-less", actual.rule)
		assert.True(t, actual.IsZero())
	})

	t.Run("calcalute discount when expired cart's item offer", func(t *testing.T) {
		cart := newMockCartWithoutItemOffer()

		actual := discounter.calculate(cart)

		assert.Equal(t, "take-more-pay-less", actual.rule)
		assert.True(t, actual.IsZero())
	})

	t.Run("calcalute discount when the offer is not applicable", func(t *testing.T) {
		cart := newMockCartWithItemOffer(2)

		actual := discounter.calculate(cart)

		assert.Equal(t, "take-more-pay-less", actual.rule)
		assert.True(t, actual.IsZero())
	})

	t.Run("calculate discount when cart has items on offer", func(t *testing.T) {
		cart := newMockCartWithItemOffer(5)

		expect := newMoney(40)

		actual := discounter.calculate(cart)
		isEquals, _ := actual.value.Equals(expect)

		assert.Equal(t, "take-more-pay-less", actual.rule)
		assert.True(t, isEquals)
	})
}

type itemOfferRepositoryStub struct{}

func (r *itemOfferRepositoryStub) FindByItemId(id uuid.UUID) (ItemOffer, error) {

	switch id.String() {

	// expired offer
	case "e0cd6104-8b97-45ba-9ff5-98b0c44b6b3e":
		return ItemOffer{
			id:             uuid.New(),
			itemId:         id,
			targetQuantity: 3,
			chargeQuantity: 2,
			expireAt:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
		}, nil

	// active offer
	case "5e96a677-6aef-4ca6-876c-ef289e936440":
		return ItemOffer{
			id:             uuid.New(),
			itemId:         id,
			targetQuantity: 5,
			chargeQuantity: 1,
			expireAt:       time.Now().AddDate(1, 0, 0),
		}, nil

	// offer not found
	default:
		return ItemOffer{}, ErrProductOfferNotFound
	}
}

func newMockCartWithItemOffer(quantity int) Cart {
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

func newMockCartWithoutItemOffer() Cart {
	return newMockCart(
		[]Item{
			{
				id:          uuid.New(),
				cartegoryId: uuid.New(),
				price:       *newMoney(10),
				quantity:    3,
			},
		},
		Discount{"none", *newMoney(0)},
	)
}

func mockCartWithExpiredItemOffer() Cart {
	return newMockCart(
		[]Item{
			{
				id:          uuid.MustParse("e0cd6104-8b97-45ba-9ff5-98b0c44b6b3e"),
				cartegoryId: uuid.New(),
				price:       *newMoney(10),
				quantity:    3,
			},
		},
		Discount{"none", *newMoney(0)},
	)
}
