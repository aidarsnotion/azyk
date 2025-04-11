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

// VitaminCompositionHandler обрабатывает HTTP‑запросы для VitaminComposition.
type VitaminCompositionHandler struct {
	svc usecase.VitaminCompositionService
}

// NewVitaminCompositionHandler создаёт новый обработчик.
func NewVitaminCompositionHandler(svc usecase.VitaminCompositionService) *VitaminCompositionHandler {
	return &VitaminCompositionHandler{svc: svc}
}

// CreateComposition — POST /compositions/vitamin
func (h *VitaminCompositionHandler) CreateComposition(w http.ResponseWriter, r *http.Request) {
	var comp models.VitaminComposition
	if err := json.NewDecoder(r.Body).Decode(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.svc.Create(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusCreated, comp)
}

// GetCompositionByID — GET /compositions/vitamin/{id}
func (h *VitaminCompositionHandler) GetCompositionByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	comp, err := h.svc.GetByID(int32(id))
	if err != nil {
		http.Error(w, "Запись не найдена", http.StatusNotFound)
		return
	}
	util.WriteJSON(w, http.StatusOK, comp)
}

// UpdateComposition — PUT /compositions/vitamin/{id}
func (h *VitaminCompositionHandler) UpdateComposition(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	var comp models.VitaminComposition
	if err := json.NewDecoder(r.Body).Decode(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comp.ID = int32(id)
	if err := h.svc.Update(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, comp)
}

// DeleteComposition — DELETE /compositions/vitamin/{id}
func (h *VitaminCompositionHandler) DeleteComposition(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	if err := h.svc.Delete(int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, map[string]string{"message": "Запись удалена"})
}

// ListCompositionsByProduct — GET /compositions/vitamin?product_id=...
func (h *VitaminCompositionHandler) ListCompositionsByProduct(w http.ResponseWriter, r *http.Request) {
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
	comps, err := h.svc.ListByProduct(int32(productID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, comps)
}
