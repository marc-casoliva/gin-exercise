package infrastructure

import (
	"fmt"
	"gin-exercise/m/v2/domain"
	"gin-exercise/m/v2/infrastructure/config"
	gormdb "gin-exercise/m/v2/infrastructure/db"
	"gorm.io/gorm"
)

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

func (g GormProductRepository) RetrieveById(id string) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func NewGormProductRepository() (domain.ProductRepository, error) {
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
