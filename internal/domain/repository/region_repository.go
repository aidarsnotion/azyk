package repository

import (
	"azyk/internal/domain/models"
	"gorm.io/gorm"
)

type RegionRepository interface {
	Create(region *models.Region) error
	GetByID(id int32) (*models.Region, error)
	List() ([]models.Region, error)
	Update(region *models.Region) error
	Delete(id int32) error
}

type regionRepository struct {
	db *gorm.DB
}

func NewRegionRepository(db *gorm.DB) RegionRepository {
	return &regionRepository{db: db}
}

func (r *regionRepository) Create(region *models.Region) error {
	return r.db.Create(region).Error
}

func (r *regionRepository) GetByID(id int32) (*models.Region, error) {
	var region models.Region
	if err := r.db.Preload("Parent").Preload("Children").Preload("Categories").First(&region, id).Error; err != nil {
		return nil, err
	}
	return &region, nil
}

func (r *regionRepository) List() ([]models.Region, error) {
	var regions []models.Region
	if err := r.db.Preload("Parent").Preload("Children").Preload("Categories").Find(&regions).Error; err != nil {
		return nil, err
	}
	return regions, nil
}

func (r *regionRepository) Update(region *models.Region) error {
	return r.db.Save(region).Error
}

func (r *regionRepository) Delete(id int32) error {
	return r.db.Delete(&models.Region{}, id).Error
}
