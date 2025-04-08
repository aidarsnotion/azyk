package http

import (
	"azyk/internal/domain/models"
	"azyk/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// MineralCompositionHandler обрабатывает HTTP-запросы для MineralComposition.
type MineralCompositionHandler struct {
	service usecase.MineralCompositionService
}

// NewMineralCompositionHandler создаёт новый обработчик.
func NewMineralCompositionHandler(svc usecase.MineralCompositionService) *MineralCompositionHandler {
	return &MineralCompositionHandler{service: svc}
}

// CreateComposition — POST /compositions/mineral
func (h *MineralCompositionHandler) CreateComposition(c *gin.Context) {
	var comp models.MineralComposition
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

// GetCompositionByID — GET /compositions/mineral/:id
func (h *MineralCompositionHandler) GetCompositionByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	comp, err := h.service.GetByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Запись не найдена"})
		return
	}
	c.JSON(http.StatusOK, comp)
}

// UpdateComposition — PUT /compositions/mineral/:id
func (h *MineralCompositionHandler) UpdateComposition(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	var comp models.MineralComposition
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

// DeleteComposition — DELETE /compositions/mineral/:id
func (h *MineralCompositionHandler) DeleteComposition(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{"message": "Запись удалена"})
}

// ListCompositionsByProduct — GET /compositions/mineral?product_id=...
func (h *MineralCompositionHandler) ListCompositionsByProduct(c *gin.Context) {
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
