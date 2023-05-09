package main

import "fmt"

type ProductRepository interface {
	Save(p Product) error
	RetreiveById(id string) (Product, error)
}

type InMemoryProductRepository struct{
	products map[string]Product
}

func NewInMemoryProductRepository() ProductRepository{
	return InMemoryProductRepository{make(map[string]Product)}
}

func (r InMemoryProductRepository) Save(p Product) error {
	r.products[p.ID] = p
	return nil
}

func (r InMemoryProductRepository) RetreiveById(id string) (Product, error) {
	p, ok := r.products[id]
	if !ok {
		return Product{}, fmt.Errorf("unexisting product with ID %v", id)
	}

	return p, nil
}