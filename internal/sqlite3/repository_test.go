package sqlite3

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProductOfferRepository(t *testing.T) {
	setUp()
	Seed()
	defer tearDown()

	repository := ProductOfferRepository{
		time: time.Date(2023, 12, 1, 23, 59, 59, 0, time.UTC),
	}

	t.Run("check product without offer", func(t *testing.T) {
		id, _ := uuid.NewUUID()

		actual := repository.IsOffer(id)

		assert.False(t, actual)
	})

	t.Run("check product with active offer", func(t *testing.T) {
		// The expiration date of this product offer is 2024-01-01 12:00:00
		id, _ := uuid.Parse("6d5544f4-8dfe-4d11-9dca-4160ff265893")

		actual := repository.IsOffer(id)

		assert.True(t, actual)
	})

	t.Run("check product with inactive offer", func(t *testing.T) {
		repository := ProductOfferRepository{
			time: time.Date(2025, 12, 1, 0, 0, 0, 0, time.UTC),
		}

		// The expiration date of this product offer is 2024-01-01 12:00:00
		id, _ := uuid.Parse("e16fefdd-dd32-424a-a4dd-466dc86613d3")

		actual := repository.IsOffer(id)

		assert.False(t, actual)
	})

	t.Run("check offer register with invalid expiration date format", func(t *testing.T) {
		stmt, err := database.Prepare("INSERT INTO products_offers VALUES (999,'6fb99b91-b00f-4753-9a74-37a69956fcf3','');")

		log.Println(err)
		stmt.QueryRow(999, "6fb99b91-b00f-4753-9a74-37a69956fcf3", "2020-01-01")
		stmt.Exec()

		id, _ := uuid.Parse("6fb99b91-b00f-4753-9a74-37a69956fcf3")
		actual := repository.IsOffer(id)

		assert.False(t, actual)
	})
}
