package repository

import (
	"azyk/internal/domain/models"

	"gorm.io/gorm"
)

// AminoAcidCompositionRepository задаёт интерфейс для CRUD‑операций с аминокислотным составом.
type AminoAcidCompositionRepository interface {
	Create(comp *models.AminoAcidComposition) error
	GetByID(id int32) (*models.AminoAcidComposition, error)
	Update(comp *models.AminoAcidComposition) error
	Delete(id int32) error
	List() ([]models.AminoAcidComposition, error)
	ListByProduct(productID int32) ([]models.AminoAcidComposition, error)
}

type aminoAcidCompositionRepository struct {
	db *gorm.DB
}

// NewAminoAcidCompositionRepository возвращает новый экземпляр репозитория.
func NewAminoAcidCompositionRepository(db *gorm.DB) AminoAcidCompositionRepository {
	return &aminoAcidCompositionRepository{db: db}
}

func (r *aminoAcidCompositionRepository) Create(comp *models.AminoAcidComposition) error {
	return r.db.Create(comp).Error
}

func (r *aminoAcidCompositionRepository) GetByID(id int32) (*models.AminoAcidComposition, error) {
	var comp models.AminoAcidComposition
	if err := r.db.First(&comp, id).Error; err != nil {
		return nil, err
	}
	return &comp, nil
}

func (r *aminoAcidCompositionRepository) Update(comp *models.AminoAcidComposition) error {
	return r.db.Save(comp).Error
}

func (r *aminoAcidCompositionRepository) Delete(id int32) error {
	return r.db.Delete(&models.AminoAcidComposition{}, id).Error
}

func (r *aminoAcidCompositionRepository) List() ([]models.AminoAcidComposition, error) {
	var comps []models.AminoAcidComposition
	if err := r.db.Find(&comps).Error; err != nil {
		return nil, err
	}
	return comps, nil
}

func (r *aminoAcidCompositionRepository) ListByProduct(productID int32) ([]models.AminoAcidComposition, error) {
	var comps []models.AminoAcidComposition
	if err := r.db.Where("product_id = ?", productID).Find(&comps).Error; err != nil {
		return nil, err
	}
	return comps, nil
}
