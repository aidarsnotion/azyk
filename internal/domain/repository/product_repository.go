package repository

import (
	"azyk/internal/domain/models"

	"gorm.io/gorm"
)

// ProductRepository задаёт интерфейс для CRUD‑операций с продуктами.
type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id int32) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id int32) error
	List() ([]models.Product, error)
}

// productRepository — реализация ProductRepository с использованием GORM.
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository возвращает новую реализацию репозитория для Product.
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// Create создаёт новый продукт в базе данных.
func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// GetByID находит продукт по его идентификатору.
func (r *productRepository) GetByID(id int32) (*models.Product, error) {
	var product models.Product
	// Для демонстрации загружаем связанные составы
	if err := r.db.
		Preload("AminoAcidCompositions").
		Preload("MineralCompositions").
		Preload("ChemicalCompositions").
		Preload("FattyAcidCompositions").
		Preload("VitaminCompositions").
		Preload("Categories").
		Preload("Region").
		First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// Update обновляет данные продукта в базе.
func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// Delete удаляет продукт по его идентификатору.
func (r *productRepository) Delete(id int32) error {
	return r.db.Delete(&models.Product{}, id).Error
}

// List возвращает список всех продуктов.
func (r *productRepository) List() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.
		Preload("AminoAcidCompositions").
		Preload("MineralCompositions").
		Preload("ChemicalCompositions").
		Preload("FattyAcidCompositions").
		Preload("VitaminCompositions").
		Preload("Categories").
		Preload("Region").
		Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
