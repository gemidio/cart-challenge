package main

import (
	"fmt"

	"github.com/gemidio/cart-challenge/internal/sqlite"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("começou")

	id := uuid.MustParse("123e4567-e89b-12d3-a456-426655440000")
	repo := sqlite.ProductOfferRepository{}
	result := repo.IsOffer(id)

	fmt.Println(result)
}
