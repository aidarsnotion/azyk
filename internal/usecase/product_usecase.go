package usecase

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
)

type ProductUsecase interface {
	CreateProduct(product *models.Product) error
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
	// Другие бизнес-логики...
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(r repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo: r}
}

func (uc *productUsecase) CreateProduct(product *models.Product) error {
	// Здесь можно добавить бизнес-логику, валидации и расчёты
	return uc.repo.Create(product)
}

func (uc *productUsecase) GetProductByID(id uint) (*models.Product, error) {
	return uc.repo.GetByID(id)
}

func (uc *productUsecase) UpdateProduct(product *models.Product) error {
	return uc.repo.Update(product)
}

func (uc *productUsecase) DeleteProduct(id uint) error {
	return uc.repo.Delete(id)
}
