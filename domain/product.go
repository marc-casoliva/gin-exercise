package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Price struct {
	Amount   float32
	Currency string
}

func NewPriceInEuros(amount float32) Price {
	return Price{amount, "EUR"}
}

type Product struct {
	ID string
	Price
	Creation    time.Time
	Description string
}

func NewProduct(id string, price Price, description string) (Product, error) {
	if _, err := uuid.Parse(id); err != nil {
		return Product{}, err
	}
	if len(description) > 50 {
		return Product{}, fmt.Errorf("description cannot be longer than 50 utf characters")
	}
	return Product{ID: id, Price: price, Creation: time.Now(), Description: description}, nil
}
