package cart

type Discounter interface {
	calculate(c Cart) Discount
}

type TakeMorePayLess struct {
	repository ItemOfferRepository
}

func (tt *TakeMorePayLess) calculate(c Cart) Discount {
	_, err := c.Subtotal()

	if err != nil {
		discount, _ := NewDiscount("take-more-pay-less", 0)

		return discount
	}

	total := newMoney(0)
	for _, item := range c.items {
		itemOffer, err := tt.repository.FindByItemId(item.id)

		if err != nil || itemOffer.isExpired() {
			continue
		}

		if !itemOffer.isApplicable(item) {
			continue
		}

		mul := itemOffer.discountApplyTo(item)
		price := item.price.Multiply(int64(mul))
		total, _ = total.Add(price)
	}

	return Discount{"take-more-pay-less", *total}
}
