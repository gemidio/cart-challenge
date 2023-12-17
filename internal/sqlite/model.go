package sqlite

type ProductOffer struct {
	ID      string
	IsOffer bool `gorm:"column:is_offer"`
}
