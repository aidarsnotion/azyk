package usecase

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
)

// MineralCompositionService задаёт интерфейс бизнес-логики для MineralComposition.
type MineralCompositionService interface {
	Create(comp *models.MineralComposition) error
	GetByID(id int32) (*models.MineralComposition, error)
	Update(comp *models.MineralComposition) error
	Delete(id int32) error
	ListByProduct(productID int32) ([]models.MineralComposition, error)
	List() ([]models.MineralComposition, error)
}

type mineralCompositionService struct {
	repo repository.MineralCompositionRepository
}

// NewMineralCompositionService создаёт новый экземпляр сервиса.
func NewMineralCompositionService(repo repository.MineralCompositionRepository) MineralCompositionService {
	return &mineralCompositionService{repo: repo}
}

func (s *mineralCompositionService) Create(comp *models.MineralComposition) error {
	// Можно добавить бизнес-валидацию, если требуется.
	return s.repo.Create(comp)
}

func (s *mineralCompositionService) GetByID(id int32) (*models.MineralComposition, error) {
	return s.repo.GetByID(id)
}

func (s *mineralCompositionService) Update(comp *models.MineralComposition) error {
	return s.repo.Update(comp)
}

func (s *mineralCompositionService) Delete(id int32) error {
	return s.repo.Delete(id)
}

func (s *mineralCompositionService) ListByProduct(productID int32) ([]models.MineralComposition, error) {
	return s.repo.ListByProduct(productID)
}

func (s *mineralCompositionService) List() ([]models.MineralComposition, error) {
	return s.repo.List()
}
