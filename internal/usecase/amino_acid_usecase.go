package usecase

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
)

// AminoAcidCompositionService задаёт интерфейс бизнес-логики для аминокислотного состава.
type AminoAcidCompositionService interface {
	Create(comp *models.AminoAcidComposition) error
	GetByID(id int32) (*models.AminoAcidComposition, error)
	Update(comp *models.AminoAcidComposition) error
	Delete(id int32) error
	ListByProduct(productID int32) ([]models.AminoAcidComposition, error)
}

type aminoAcidCompositionService struct {
	repo repository.AminoAcidCompositionRepository
}

// NewAminoAcidCompositionService создаёт новый экземпляр сервиса.
func NewAminoAcidCompositionService(repo repository.AminoAcidCompositionRepository) AminoAcidCompositionService {
	return &aminoAcidCompositionService{repo: repo}
}

func (s *aminoAcidCompositionService) Create(comp *models.AminoAcidComposition) error {
	// Здесь можно добавить бизнес-валидацию, если требуется.
	return s.repo.Create(comp)
}

func (s *aminoAcidCompositionService) GetByID(id int32) (*models.AminoAcidComposition, error) {
	return s.repo.GetByID(id)
}

func (s *aminoAcidCompositionService) Update(comp *models.AminoAcidComposition) error {
	return s.repo.Update(comp)
}

func (s *aminoAcidCompositionService) Delete(id int32) error {
	return s.repo.Delete(id)
}

func (s *aminoAcidCompositionService) ListByProduct(productID int32) ([]models.AminoAcidComposition, error) {
	return s.repo.ListByProduct(productID)
}
