package main

import (
	"fmt"
	"gin-exercise/m/v2/infrastructure/config"
	gormdb "gin-exercise/m/v2/infrastructure/db"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(p Product) error
	RetreiveById(id string) (Product, error)
}

type InMemoryProductRepository struct {
	products map[string]Product
}

type GormProductRepository struct {
	db *gorm.DB
}

func (g GormProductRepository) Save(p Product) error {
	//TODO implement me
	panic("implement me")
}

func (g GormProductRepository) RetreiveById(id string) (Product, error) {
	//TODO implement me
	panic("implement me")
}

func NewGormProductRepository() (ProductRepository, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	db, err := gormdb.NewDB(cfg)
	if err != nil {
		return nil, err
	}
	return GormProductRepository{db}, nil
}
func NewInMemoryProductRepository() ProductRepository {
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
