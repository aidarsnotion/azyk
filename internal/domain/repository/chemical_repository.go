package repository

import (
	"azyk/internal/domain/models"
	"gorm.io/gorm"
)

// ChemicalCompositionRepository задаёт интерфейс CRUD‑операций для ChemicalComposition.
type ChemicalCompositionRepository interface {
	Create(comp *models.ChemicalComposition) error
	GetByID(id int32) (*models.ChemicalComposition, error)
	Update(comp *models.ChemicalComposition) error
	Delete(id int32) error
	List() ([]models.ChemicalComposition, error)
	ListByProduct(productID int32) ([]models.ChemicalComposition, error)
}

type chemicalCompositionRepository struct {
	db *gorm.DB
}

// NewChemicalCompositionRepository возвращает новую реализацию репозитория для ChemicalComposition.
func NewChemicalCompositionRepository(db *gorm.DB) ChemicalCompositionRepository {
	return &chemicalCompositionRepository{db: db}
}

func (r *chemicalCompositionRepository) Create(comp *models.ChemicalComposition) error {
	return r.db.Create(comp).Error
}

func (r *chemicalCompositionRepository) GetByID(id int32) (*models.ChemicalComposition, error) {
	var comp models.ChemicalComposition
	if err := r.db.First(&comp, id).Error; err != nil {
		return nil, err
	}
	return &comp, nil
}

func (r *chemicalCompositionRepository) Update(comp *models.ChemicalComposition) error {
	return r.db.Save(comp).Error
}

func (r *chemicalCompositionRepository) Delete(id int32) error {
	return r.db.Delete(&models.ChemicalComposition{}, id).Error
}

func (r *chemicalCompositionRepository) List() ([]models.ChemicalComposition, error) {
	var comps []models.ChemicalComposition
	if err := r.db.Find(&comps).Error; err != nil {
		return nil, err
	}
	return comps, nil
}

func (r *chemicalCompositionRepository) ListByProduct(productID int32) ([]models.ChemicalComposition, error) {
	var comps []models.ChemicalComposition
	if err := r.db.Where("product_id = ?", productID).Find(&comps).Error; err != nil {
		return nil, err
	}
	return comps, nil
}
