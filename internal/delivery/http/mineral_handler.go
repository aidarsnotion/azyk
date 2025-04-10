package http

import (
	"azyk/internal/domain/models"
	"azyk/internal/usecase"
	"azyk/util"
	"encoding/json"
	"github.com/gorilla/mux"
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
func (h *MineralCompositionHandler) CreateComposition(w http.ResponseWriter, r *http.Request) {
	var comp models.MineralComposition
	if err := json.NewDecoder(r.Body).Decode(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusCreated, comp)
}

// GetCompositionByID — GET /compositions/mineral/{id}
func (h *MineralCompositionHandler) GetCompositionByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	comp, err := h.service.GetByID(int32(id))
	if err != nil {
		http.Error(w, "Запись не найдена", http.StatusNotFound)
		return
	}
	util.WriteJSON(w, http.StatusOK, comp)
}

// UpdateComposition — PUT /compositions/mineral/{id}
func (h *MineralCompositionHandler) UpdateComposition(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	var comp models.MineralComposition
	if err := json.NewDecoder(r.Body).Decode(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comp.ID = int32(id)
	if err := h.service.Update(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, comp)
}

// DeleteComposition — DELETE /compositions/mineral/{id}
func (h *MineralCompositionHandler) DeleteComposition(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	if err := h.service.Delete(int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, map[string]string{"message": "Запись удалена"})
}

// ListCompositionsByProduct — GET /compositions/mineral?product_id=...
func (h *MineralCompositionHandler) ListCompositionsByProduct(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("product_id")
	if productIDStr == "" {
		http.Error(w, "Параметр product_id обязателен", http.StatusBadRequest)
		return
	}
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Неверный формат product_id", http.StatusBadRequest)
		return
	}
	comps, err := h.service.ListByProduct(int32(productID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, comps)
}
