package sqlite

import (
	"github.com/google/uuid"
)

type ProductOfferRepository struct{}

func (po ProductOfferRepository) IsOffer(id uuid.UUID) bool {
	var productOffer ProductOffer
	database.Unscoped().Find(&productOffer, "id = ?", id.String())

	return productOffer.IsOffer
}
