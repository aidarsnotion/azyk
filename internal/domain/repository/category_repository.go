// repository/category_repository.go
package repository

import (
	"azyk/internal/domain/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetByID(id int32) (*models.Category, error)
	List() ([]models.Category, error)
	Update(category *models.Category) error
	Delete(id int32) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetByID(id int32) (*models.Category, error) {
	var category models.Category
	if err := r.db.Preload("Region").First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) List() ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Preload("Region").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id int32) error {
	return r.db.Delete(&models.Category{}, id).Error
}
