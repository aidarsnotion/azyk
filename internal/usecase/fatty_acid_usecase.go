package usecase

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
)

// FattyAcidCompositionService определяет интерфейс бизнес‑логики для FattyAcidComposition.
type FattyAcidCompositionService interface {
	Create(comp *models.FattyAcidComposition) error
	GetByID(id int32) (*models.FattyAcidComposition, error)
	Update(comp *models.FattyAcidComposition) error
	Delete(id int32) error
	List() ([]models.FattyAcidComposition, error)
	ListByProduct(productID int32) ([]models.FattyAcidComposition, error)
}

type fattyAcidCompositionService struct {
	repo repository.FattyAcidCompositionRepository
}

// NewFattyAcidCompositionService создаёт новый экземпляр сервиса.
func NewFattyAcidCompositionService(repo repository.FattyAcidCompositionRepository) FattyAcidCompositionService {
	return &fattyAcidCompositionService{repo: repo}
}

func (s *fattyAcidCompositionService) Create(comp *models.FattyAcidComposition) error {
	// Здесь можно добавить бизнес-валидацию, если требуется
	return s.repo.Create(comp)
}

func (s *fattyAcidCompositionService) GetByID(id int32) (*models.FattyAcidComposition, error) {
	return s.repo.GetByID(id)
}

func (s *fattyAcidCompositionService) Update(comp *models.FattyAcidComposition) error {
	return s.repo.Update(comp)
}

func (s *fattyAcidCompositionService) Delete(id int32) error {
	return s.repo.Delete(id)
}

func (s *fattyAcidCompositionService) List() ([]models.FattyAcidComposition, error) {
	return s.repo.List()
}

func (s *fattyAcidCompositionService) ListByProduct(productID int32) ([]models.FattyAcidComposition, error) {
	return s.repo.ListByProduct(productID)
}
