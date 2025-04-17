package usecase

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
)

type RegionUsecase interface {
	Create(region *models.Region) error
	GetByID(id int32) (*models.Region, error)
	List() ([]models.Region, error)
	Update(region *models.Region) (*models.Region, error) // !!! здесь надо возвращать *Region
	Delete(id int32) (*models.Region, error)
}

type regionUsecase struct {
	repo repository.RegionRepository
}

func NewRegionUsecase(r repository.RegionRepository) RegionUsecase {
	return &regionUsecase{repo: r}
}

func (u *regionUsecase) Create(region *models.Region) error {
	return u.repo.Create(region)
}

func (u *regionUsecase) GetByID(id int32) (*models.Region, error) {
	return u.repo.GetByID(id)
}

func (u *regionUsecase) List() ([]models.Region, error) {
	return u.repo.List()
}

func (u *regionUsecase) Update(region *models.Region) (*models.Region, error) {
	return u.repo.Update(region)
}

func (u *regionUsecase) Delete(id int32) (*models.Region, error) {
	return u.repo.Delete(id)
}
