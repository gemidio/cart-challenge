package cart

import "github.com/google/uuid"

type ProductOfferRepository interface {
	IsOffer(id uuid.UUID) bool
}
