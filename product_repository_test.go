package main

import (
	"testing"

	"github.com/google/uuid"
)

func Test_Save(t *testing.T) {
	t.Run("should save the product inMemory", func(t *testing.T) {
		description := "product description saved"
		p, _:= NewProduct(uuid.NewString(), NewPriceInEuros(22), description)
		sut := NewInMemoryProductRepository()

		if err := sut.Save(p); err != nil {
			t.Errorf("Should have saved the product correctly, got an error: %v", err)
		}

		retreivedP, err := sut.RetreiveById(p.ID)
		if err != nil {
			t.Errorf("Product not found on repository: %v", err)
		}
		if retreivedP.Description != description {
			t.Errorf("Mismatching saved description, expected %v but got %v", description, retreivedP.Description)
		}

	})
}

func Test_RetreiveById(t *testing.T) {
	t.Run("should return an error if product ID does not exist", func(t *testing.T) {
		sut := NewInMemoryProductRepository()
		if _, err := sut.RetreiveById("unexisting"); err == nil {
			t.Errorf("Should have returned an error, got nil")
		}
	})
}