package usecase

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
)

// VitaminCompositionService определяет интерфейс бизнес-логики для VitaminComposition.
type VitaminCompositionService interface {
	Create(comp *models.VitaminComposition) error
	GetByID(id int32) (*models.VitaminComposition, error)
	Update(comp *models.VitaminComposition) error
	Delete(id int32) error
	List() ([]models.VitaminComposition, error)
	ListByProduct(productID int32) ([]models.VitaminComposition, error)
}

type vitaminCompositionService struct {
	repo repository.VitaminCompositionRepository
}

// NewVitaminCompositionService создаёт новый экземпляр сервиса.
func NewVitaminCompositionService(repo repository.VitaminCompositionRepository) VitaminCompositionService {
	return &vitaminCompositionService{repo: repo}
}

func (s *vitaminCompositionService) Create(comp *models.VitaminComposition) error {
	// Добавьте бизнес-валидацию, если необходимо.
	return s.repo.Create(comp)
}

func (s *vitaminCompositionService) GetByID(id int32) (*models.VitaminComposition, error) {
	return s.repo.GetByID(id)
}

func (s *vitaminCompositionService) Update(comp *models.VitaminComposition) error {
	return s.repo.Update(comp)
}

func (s *vitaminCompositionService) Delete(id int32) error {
	return s.repo.Delete(id)
}

func (s *vitaminCompositionService) List() ([]models.VitaminComposition, error) {
	return s.repo.List()
}

func (s *vitaminCompositionService) ListByProduct(productID int32) ([]models.VitaminComposition, error) {
	return s.repo.ListByProduct(productID)
}
