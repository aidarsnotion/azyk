package http

import (
	"net/http"
	"strconv"

	"azyk/internal/domain/models"
	"azyk/internal/usecase"
	"github.com/gin-gonic/gin"
)

// ProductHandler обрабатывает HTTP‑запросы для работы с продуктами.
type ProductHandler struct {
	productService usecase.ProductService
}

// NewProductHandler создаёт новый обработчик для продуктов.
func NewProductHandler(productService usecase.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// CreateProduct — POST /products
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.productService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// GetProductByID — GET /products/:id
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	product, err := h.productService.GetProductByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// UpdateProduct — PUT /products/:id
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = int32(id)
	if err := h.productService.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// DeleteProduct — DELETE /products/:id
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	if err := h.productService.DeleteProduct(int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Продукт удалён"})
}

// ListProducts — GET /products
func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.productService.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
