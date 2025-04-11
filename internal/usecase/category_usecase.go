// usecase/category_usecase.go
package usecase

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
)

type CategoryUsecase interface {
	Create(category *models.Category) error
	GetByID(id int32) (*models.Category, error)
	List() ([]models.Category, error)
	Update(category *models.Category) error
	Delete(id int32) error
}

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(r repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{repo: r}
}

func (u *categoryUsecase) Create(category *models.Category) error {
	return u.repo.Create(category)
}

func (u *categoryUsecase) GetByID(id int32) (*models.Category, error) {
	return u.repo.GetByID(id)
}

func (u *categoryUsecase) List() ([]models.Category, error) {
	return u.repo.List()
}

func (u *categoryUsecase) Update(category *models.Category) error {
	return u.repo.Update(category)
}

func (u *categoryUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
