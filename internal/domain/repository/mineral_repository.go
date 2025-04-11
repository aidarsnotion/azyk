package repository

import (
	"azyk/internal/domain/models"
	"gorm.io/gorm"
)

// MineralCompositionRepository задаёт интерфейс CRUD‑операций для MineralComposition.
type MineralCompositionRepository interface {
	Create(comp *models.MineralComposition) error
	GetByID(id int32) (*models.MineralComposition, error)
	Update(comp *models.MineralComposition) error
	Delete(id int32) error
	List() ([]models.MineralComposition, error)
	ListByProduct(productID int32) ([]models.MineralComposition, error)
}

type mineralCompositionRepository struct {
	db *gorm.DB
}

// NewMineralCompositionRepository возвращает новый экземпляр репозитория для MineralComposition.
func NewMineralCompositionRepository(db *gorm.DB) MineralCompositionRepository {
	return &mineralCompositionRepository{db: db}
}

func (r *mineralCompositionRepository) Create(comp *models.MineralComposition) error {
	return r.db.Create(comp).Error
}

func (r *mineralCompositionRepository) GetByID(id int32) (*models.MineralComposition, error) {
	var comp models.MineralComposition
	if err := r.db.First(&comp, id).Error; err != nil {
		return nil, err
	}
	return &comp, nil
}

func (r *mineralCompositionRepository) Update(comp *models.MineralComposition) error {
	return r.db.Save(comp).Error
}

func (r *mineralCompositionRepository) Delete(id int32) error {
	return r.db.Delete(&models.MineralComposition{}, id).Error
}

func (r *mineralCompositionRepository) List() ([]models.MineralComposition, error) {
	var comps []models.MineralComposition
	if err := r.db.Find(&comps).Error; err != nil {
		return nil, err
	}
	return comps, nil
}

func (r *mineralCompositionRepository) ListByProduct(productID int32) ([]models.MineralComposition, error) {
	var comps []models.MineralComposition
	if err := r.db.Where("product_id = ?", productID).Find(&comps).Error; err != nil {
		return nil, err
	}
	return comps, nil
}
