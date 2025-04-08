package http

import (
	"azyk/internal/domain/models"
	"azyk/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FattyAcidCompositionHandler обрабатывает HTTP‑запросы для FattyAcidComposition.
type FattyAcidCompositionHandler struct {
	svc usecase.FattyAcidCompositionService
}

// NewFattyAcidCompositionHandler создаёт новый обработчик.
func NewFattyAcidCompositionHandler(svc usecase.FattyAcidCompositionService) *FattyAcidCompositionHandler {
	return &FattyAcidCompositionHandler{svc: svc}
}

// CreateComposition — POST /compositions/fatty
func (h *FattyAcidCompositionHandler) CreateComposition(c *gin.Context) {
	var comp models.FattyAcidComposition
	if err := c.ShouldBindJSON(&comp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.Create(&comp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, comp)
}

// GetCompositionByID — GET /compositions/fatty/:id
func (h *FattyAcidCompositionHandler) GetCompositionByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	comp, err := h.svc.GetByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Запись не найдена"})
		return
	}
	c.JSON(http.StatusOK, comp)
}

// UpdateComposition — PUT /compositions/fatty/:id
func (h *FattyAcidCompositionHandler) UpdateComposition(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	var comp models.FattyAcidComposition
	if err := c.ShouldBindJSON(&comp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comp.ID = int32(id)
	if err := h.svc.Update(&comp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comp)
}

// DeleteComposition — DELETE /compositions/fatty/:id
func (h *FattyAcidCompositionHandler) DeleteComposition(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	if err := h.svc.Delete(int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Запись удалена"})
}

// ListCompositionsByProduct — GET /compositions/fatty?product_id=...
func (h *FattyAcidCompositionHandler) ListCompositionsByProduct(c *gin.Context) {
	productIDStr := c.Query("product_id")
	if productIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр product_id обязателен"})
		return
	}
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат product_id"})
		return
	}
	comps, err := h.svc.ListByProduct(int32(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comps)
}
