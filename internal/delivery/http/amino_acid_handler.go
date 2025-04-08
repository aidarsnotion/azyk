package http

import (
	"azyk/internal/domain/models"
	"azyk/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AminoAcidCompositionHandler обрабатывает HTTP-запросы для аминокислотного состава.
type AminoAcidCompositionHandler struct {
	service usecase.AminoAcidCompositionService
}

// NewAminoAcidCompositionHandler создаёт новый обработчик.
func NewAminoAcidCompositionHandler(svc usecase.AminoAcidCompositionService) *AminoAcidCompositionHandler {
	return &AminoAcidCompositionHandler{service: svc}
}

// CreateComposition — POST /compositions/amino
func (h *AminoAcidCompositionHandler) CreateComposition(c *gin.Context) {
	var comp models.AminoAcidComposition
	if err := c.ShouldBindJSON(&comp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&comp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, comp)
}

// GetCompositionByID — GET /compositions/amino/:id
func (h *AminoAcidCompositionHandler) GetCompositionByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	comp, err := h.service.GetByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Состав не найден"})
		return
	}
	c.JSON(http.StatusOK, comp)
}

// UpdateComposition — PUT /compositions/amino/:id
func (h *AminoAcidCompositionHandler) UpdateComposition(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	var comp models.AminoAcidComposition
	if err := c.ShouldBindJSON(&comp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comp.ID = int32(id)
	if err := h.service.Update(&comp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comp)
}

// DeleteComposition — DELETE /compositions/amino/:id
func (h *AminoAcidCompositionHandler) DeleteComposition(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	if err := h.service.Delete(int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Состав удалён"})
}

// ListCompositionsByProduct — GET /compositions/amino?product_id=...
func (h *AminoAcidCompositionHandler) ListCompositionsByProduct(c *gin.Context) {
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
	comps, err := h.service.ListByProduct(int32(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comps)
}
