package main

import (
	"fmt"
	"gin-exercise/m/v2/domain"
	"gin-exercise/m/v2/infrastructure/config"
	gormdb "gin-exercise/m/v2/infrastructure/db"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(p domain.Product) error
	RetreiveById(id string) (domain.Product, error)
}

type InMemoryProductRepository struct {
	products map[string]domain.Product
}

type GormProductRepository struct {
	db *gorm.DB
}

func (g GormProductRepository) Save(p domain.Product) error {
	productDB := gormdb.NewProductDB(p)
	tx := g.db.Create(productDB)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (g GormProductRepository) RetreiveById(id string) (domain.Product, error) {
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
	return InMemoryProductRepository{make(map[string]domain.Product)}
}

func (r InMemoryProductRepository) Save(p domain.Product) error {
	r.products[p.ID] = p
	return nil
}

func (r InMemoryProductRepository) RetreiveById(id string) (domain.Product, error) {
	p, ok := r.products[id]
	if !ok {
		return domain.Product{}, fmt.Errorf("unexisting product with ID %v", id)
	}

	return p, nil
}
