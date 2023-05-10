package domain

import (
	"testing"

	"github.com/google/uuid"
)

func Test_NewProduct(t *testing.T) {
	t.Run("id should be a valid uuid", func(t *testing.T) {
		id := uuid.NewString()
		p, err := NewProduct(id, Price{}, "")
		if err != nil {
			t.Errorf("A valid uuid should be accepted but got this error: %v", err)
		}
		if p.ID != id {
			t.Errorf("Expected id %s was %v", id, p.ID)
		}
	})
	t.Run("invalid id should return error", func(t *testing.T) {
		invalidUUID := "invalid-id"
		if _, err := NewProduct(invalidUUID, Price{}, ""); err == nil {
			t.Errorf("Id %s should have been invalid but wasn't", invalidUUID)
		}

		invalidASCII := "475191éb-呐cd6-4549-bd0e-495😑525c5723"
		if _, err := NewProduct(invalidASCII, Price{}, ""); err == nil {
			t.Errorf("Id %s should have been invalid but wasn't", invalidUUID)
		}
	})

	t.Run("description should be valid", func(t *testing.T) {
		validDescription := "Valid description"
		p, err := NewProduct(uuid.NewString(), Price{}, validDescription)
		if err != nil {
			t.Errorf("Description: %s should have been valid, but wasn't", validDescription)
		}
		if p.Description != validDescription {
			t.Errorf("Product description should be: %s but instead is: %v", validDescription, p.Description)
		}
	})

	t.Run("should return error when invalid description", func(t *testing.T) {
		invalidDescription := "A Description of Exactly fifty one utf-8 characters"
		if _, err := NewProduct(uuid.NewString(), Price{}, invalidDescription); err == nil {
			t.Errorf("51 char description should have returned an error")
		}

	})

	t.Run("should add the Price", func(t *testing.T) {
		price := NewPriceInEuros(45)
		p, err := NewProduct(uuid.NewString(), price, "description")
		if err != nil {
			t.Errorf("Price: %v should have been valid, but wasn't", price)
		}

		if p.Price.Amount != price.Amount {
			t.Errorf("Expected price amount of %v, but got %v", price.Amount, p.Price.Amount)
		}

		if p.Price.Currency != price.Currency {
			t.Errorf("Expected price currency of %v, but got %v", price.Currency, p.Price.Currency)
		}
	})

	t.Run("should add a timestamp on Product creation", func(t *testing.T) {
		p, err := NewProduct(uuid.NewString(), Price{}, "description")
		if err != nil {
			t.Errorf("Creation: should not have had any kind of error, got %v", err)
		}

		if p.Creation.IsZero() {
			t.Errorf("should have set a creation Timestamp but got Zero value %v", p.Creation.String())
		}
	})

}
