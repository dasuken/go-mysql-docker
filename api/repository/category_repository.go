package repository

import (
	"github.com/noguchidaisuke/go-mysql-docker/api/models"
	"gorm.io/gorm"
	"time"
)

type CategoriesRepository interface {
	Save(category *models.Category) (*models.Category, error)
	FindAll() ([]*models.Category, error)
	Find(category_id uint64) (*models.Category, error)
	Update(category *models.Category) error
	Delete(category_id uint64) error
}

type categoriesRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) CategoriesRepository {
	return &categoriesRepositoryImpl{db: db}
}

func (r *categoriesRepositoryImpl) Save(category *models.Category) (*models.Category, error) {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Category{}).Create(category).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return category, tx.Commit().Error
}

func (r *categoriesRepositoryImpl) FindAll() ([]*models.Category, error) {
	categories := []*models.Category{}
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoriesRepositoryImpl) Find(category_id uint64) (*models.Category, error) {
	category := &models.Category{}
	err := r.db.Where("id = ?", category_id).Preload("Products").Find(category).Error
	return category, err
}

func (r *categoriesRepositoryImpl) Update(category *models.Category) error {
	tx := r.db.Begin()

	columns := map[string]interface{} {
		"description": category.Description,
		"updated_at": time.Now(),
	}

	err := tx.Debug().Model(&models.Category{}).Where("id = ?", category.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *categoriesRepositoryImpl) Delete(category_id uint64) error {
	tx := r.db.Begin()

	err := tx.Debug().Model(&models.Category{}).Where("id = ?", category_id).Delete(&models.Category{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
