package db

import "time"

type ProductDB struct {
	ID            string    `gorm:"primaryKey"`
	CreatedAt     time.Time `gorm:"not null"`
	Description   string    `gorm:"not null"`
	PriceAmount   float32   `gorm:"not null"`
	PriceCurrency string    `gorm:"not null"`
}
