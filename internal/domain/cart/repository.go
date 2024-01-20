package cart

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCouponNotFound error = errors.New("coupon not found")
)

type ProductOfferRepository interface {
	IsOffer(id uuid.UUID) bool
}

type PromotionCouponRepository interface {
	Find(label string) (Coupon, error)
}
