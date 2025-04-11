package usecase

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
)

// ChemicalCompositionService определяет интерфейс бизнес-логики для сущности ChemicalComposition.
type ChemicalCompositionService interface {
	Create(comp *models.ChemicalComposition) error
	GetByID(id int32) (*models.ChemicalComposition, error)
	Update(comp *models.ChemicalComposition) error
	Delete(id int32) error
	ListByProduct(productID int32) ([]models.ChemicalComposition, error)
	List() ([]models.ChemicalComposition, error)
}

type chemicalCompositionService struct {
	repo repository.ChemicalCompositionRepository
}

// NewChemicalCompositionService создаёт новый экземпляр сервиса.
func NewChemicalCompositionService(repo repository.ChemicalCompositionRepository) ChemicalCompositionService {
	return &chemicalCompositionService{repo: repo}
}

func (s *chemicalCompositionService) Create(comp *models.ChemicalComposition) error {
	// Здесь можно добавить бизнес-валидацию.
	return s.repo.Create(comp)
}

func (s *chemicalCompositionService) GetByID(id int32) (*models.ChemicalComposition, error) {
	return s.repo.GetByID(id)
}

func (s *chemicalCompositionService) Update(comp *models.ChemicalComposition) error {
	return s.repo.Update(comp)
}

func (s *chemicalCompositionService) Delete(id int32) error {
	return s.repo.Delete(id)
}

func (s *chemicalCompositionService) ListByProduct(productID int32) ([]models.ChemicalComposition, error) {
	return s.repo.ListByProduct(productID)
}

func (s *chemicalCompositionService) List() ([]models.ChemicalComposition, error) {
	return s.repo.List()
}
