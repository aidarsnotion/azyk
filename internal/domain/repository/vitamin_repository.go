package repository

import (
	"azyk/internal/domain/models"
	"gorm.io/gorm"
)

// VitaminCompositionRepository задаёт интерфейс CRUD‑операций для VitaminComposition.
type VitaminCompositionRepository interface {
	Create(comp *models.VitaminComposition) error
	GetByID(id int32) (*models.VitaminComposition, error)
	Update(comp *models.VitaminComposition) error
	Delete(id int32) error
	List() ([]models.VitaminComposition, error)
	ListByProduct(productID int32) ([]models.VitaminComposition, error)
}

type vitaminCompositionRepository struct {
	db *gorm.DB
}

// NewVitaminCompositionRepository возвращает новую реализацию репозитория.
func NewVitaminCompositionRepository(db *gorm.DB) VitaminCompositionRepository {
	return &vitaminCompositionRepository{db: db}
}

func (r *vitaminCompositionRepository) Create(comp *models.VitaminComposition) error {
	return r.db.Create(comp).Error
}

func (r *vitaminCompositionRepository) GetByID(id int32) (*models.VitaminComposition, error) {
	var comp models.VitaminComposition
	if err := r.db.First(&comp, id).Error; err != nil {
		return nil, err
	}
	return &comp, nil
}

func (r *vitaminCompositionRepository) Update(comp *models.VitaminComposition) error {
	return r.db.Save(comp).Error
}

func (r *vitaminCompositionRepository) Delete(id int32) error {
	return r.db.Delete(&models.VitaminComposition{}, id).Error
}

func (r *vitaminCompositionRepository) List() ([]models.VitaminComposition, error) {
	var comps []models.VitaminComposition
	if err := r.db.Find(&comps).Error; err != nil {
		return nil, err
	}
	return comps, nil
}

func (r *vitaminCompositionRepository) ListByProduct(productID int32) ([]models.VitaminComposition, error) {
	var comps []models.VitaminComposition
	if err := r.db.Where("product_id = ?", productID).Find(&comps).Error; err != nil {
		return nil, err
	}
	return comps, nil
}
