package db

import (
	domain "gin-exercise/m/v2/domain"
	"time"
)

type Product struct {
	ID            string    `gorm:"primaryKey"`
	CreatedAt     time.Time `gorm:"not null"`
	Description   string    `gorm:"not null"`
	PriceAmount   float32   `gorm:"not null"`
	PriceCurrency string    `gorm:"not null"`
}

func NewProductDB(p domain.Product) Product {

	return Product{
		ID:            p.ID,
		CreatedAt:     p.Creation,
		Description:   p.Description,
		PriceAmount:   p.Price.Amount,
		PriceCurrency: p.Price.Currency,
	}
}
func (product Product) ToProduct() domain.Product {

	return domain.Product{
		ID:          product.ID,
		Description: product.Description,
		Price: domain.Price{
			Amount:   product.PriceAmount,
			Currency: product.PriceCurrency,
		},
		Creation: product.CreatedAt,
	}
}
