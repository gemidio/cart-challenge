package sqlite3

import (
	"time"

	"github.com/google/uuid"
)

type ProductOfferRepository struct {
	time time.Time
}

func (po ProductOfferRepository) IsOffer(uuid uuid.UUID) bool {
	stmt, _ := database.Prepare("SELECT expire_at FROM products_offers WHERE product_id = ?;")
	defer stmt.Close()

	var expireAtColumn string
	err := stmt.QueryRow(uuid.String()).Scan(&expireAtColumn)

	if err != nil {
		return false
	}

	expireAt, err := time.Parse("2006-01-02 15:04:05", expireAtColumn)

	if err != nil {
		return false
	}

	return expireAt.After(po.time)
}
