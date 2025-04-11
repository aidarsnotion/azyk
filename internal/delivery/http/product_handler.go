package http

import (
	"azyk/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"azyk/internal/domain/models"
	"azyk/internal/usecase"
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
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.productService.CreateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusCreated, product)
}

// GetProductByID — GET /products/{id}
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	product, err := h.productService.GetProductByID(int32(id))
	if err != nil {
		http.Error(w, "Продукт не найден", http.StatusNotFound)
		return
	}
	util.WriteJSON(w, http.StatusOK, product)
}

// UpdateProduct — PUT /products/{id}
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product.ID = int32(id)
	if err := h.productService.UpdateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, product)
}

// DeleteProduct — DELETE /products/{id}
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	if err := h.productService.DeleteProduct(int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, map[string]string{"message": "Продукт удалён"})
}

// ListProducts — GET /products
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.ListProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, products)
}
