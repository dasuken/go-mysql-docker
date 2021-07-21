package repository

import (
	"github.com/noguchidaisuke/go-mysql-docker/api/models"
	"gorm.io/gorm"
)

type ProductsRepository interface {
	Save(product *models.Product) (*models.Product, error)
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductsRepository {
	return &productRepositoryImpl{db:db}
}

func (r *productRepositoryImpl) Save(product *models.Product) (*models.Product, error) {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Product{}).Create(product).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return product, tx.Commit().Error
}