package cart

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductOfferNotFound error = errors.New("product offer not found")
)

type ItemOfferRepository interface {
	FindByItemId(id uuid.UUID) (ItemOffer, error)
}
