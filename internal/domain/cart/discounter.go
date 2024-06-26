package cart

type Discounter interface {
	calculate(c Cart) Discount
}

type TakeThreePayTwoDiscounter struct {
	repository ProductOfferRepository
}

func (tt *TakeThreePayTwoDiscounter) calculate(c Cart) Discount {
	_, err := c.Subtotal()

	if err != nil {
		discount, _ := NewDiscount("take-3-pay-2", 0)

		return discount
	}

	total := newMoney(0)

	for _, item := range c.items {

		if !tt.repository.IsOffer(item.id) {
			continue
		}

		if item.quantity%3 != 0 {
			continue
		}

		multiplier := int64(item.quantity / 3)
		price := item.price.Multiply(multiplier)
		total, _ = total.Add(price)
	}

	return Discount{"take-3-pay-2", *total}
}
