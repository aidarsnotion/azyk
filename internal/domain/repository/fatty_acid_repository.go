package repository

import (
	"azyk/internal/domain/models"
	"gorm.io/gorm"
)

// FattyAcidCompositionRepository задаёт интерфейс CRUD‑операций для FattyAcidComposition.
type FattyAcidCompositionRepository interface {
	Create(comp *models.FattyAcidComposition) error
	GetByID(id int32) (*models.FattyAcidComposition, error)
	Update(comp *models.FattyAcidComposition) error
	Delete(id int32) error
	List() ([]models.FattyAcidComposition, error)
	ListByProduct(productID int32) ([]models.FattyAcidComposition, error)
}

type fattyAcidCompositionRepository struct {
	db *gorm.DB
}

// NewFattyAcidCompositionRepository возвращает новую реализацию репозитория.
func NewFattyAcidCompositionRepository(db *gorm.DB) FattyAcidCompositionRepository {
	return &fattyAcidCompositionRepository{db: db}
}

func (r *fattyAcidCompositionRepository) Create(comp *models.FattyAcidComposition) error {
	return r.db.Create(comp).Error
}

func (r *fattyAcidCompositionRepository) GetByID(id int32) (*models.FattyAcidComposition, error) {
	var comp models.FattyAcidComposition
	if err := r.db.First(&comp, id).Error; err != nil {
		return nil, err
	}
	return &comp, nil
}

func (r *fattyAcidCompositionRepository) Update(comp *models.FattyAcidComposition) error {
	return r.db.Save(comp).Error
}

func (r *fattyAcidCompositionRepository) Delete(id int32) error {
	return r.db.Delete(&models.FattyAcidComposition{}, id).Error
}

func (r *fattyAcidCompositionRepository) List() ([]models.FattyAcidComposition, error) {
	var comps []models.FattyAcidComposition
	if err := r.db.Find(&comps).Error; err != nil {
		return nil, err
	}
	return comps, nil
}

func (r *fattyAcidCompositionRepository) ListByProduct(productID int32) ([]models.FattyAcidComposition, error) {
	var comps []models.FattyAcidComposition
	if err := r.db.Where("product_id = ?", productID).Find(&comps).Error; err != nil {
		return nil, err
	}
	return comps, nil
}
