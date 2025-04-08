package usecase

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/repository"
)

// ProductService задаёт интерфейс бизнес‑логики для работы с продуктами.
type ProductService interface {
	CreateProduct(product *models.Product) error
	GetProductByID(id int32) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id int32) error
	ListProducts() ([]models.Product, error)
}

// productService — реализация ProductService.
type productService struct {
	repo repository.ProductRepository
}

// NewProductService создаёт новый экземпляр ProductService.
func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

// CreateProduct создаёт продукт через репозиторий.
func (s *productService) CreateProduct(product *models.Product) error {
	// Здесь можно добавить дополнительную бизнес-валидацию
	return s.repo.Create(product)
}

// GetProductByID возвращает продукт по id.
func (s *productService) GetProductByID(id int32) (*models.Product, error) {
	return s.repo.GetByID(id)
}

// UpdateProduct обновляет данные продукта.
func (s *productService) UpdateProduct(product *models.Product) error {
	return s.repo.Update(product)
}

// DeleteProduct удаляет продукт по id.
func (s *productService) DeleteProduct(id int32) error {
	return s.repo.Delete(id)
}

// ListProducts возвращает список всех продуктов.
func (s *productService) ListProducts() ([]models.Product, error) {
	return s.repo.List()
}
