package infrastructure

import (
	"fmt"
	"gin-exercise/m/v2/domain"
)

type InMemoryProductRepository struct {
	products map[string]domain.Product
}

func NewInMemoryProductRepository() domain.ProductRepository {
	return InMemoryProductRepository{make(map[string]domain.Product)}
}

func (r InMemoryProductRepository) Save(p domain.Product) error {
	r.products[p.ID] = p
	return nil
}

func (r InMemoryProductRepository) RetrieveById(id string) (domain.Product, error) {
	p, ok := r.products[id]
	if !ok {
		return domain.Product{}, fmt.Errorf("unexisting product with ID %v", id)
	}

	return p, nil
}
