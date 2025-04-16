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

// FattyAcidCompositionHandler обрабатывает HTTP‑запросы для FattyAcidComposition.
type FattyAcidCompositionHandler struct {
	svc usecase.FattyAcidCompositionService
}

// NewFattyAcidCompositionHandler создаёт новый обработчик.
func NewFattyAcidCompositionHandler(svc usecase.FattyAcidCompositionService) *FattyAcidCompositionHandler {
	return &FattyAcidCompositionHandler{svc: svc}
}

// CreateComposition — POST /compositions/fatty
func (h *FattyAcidCompositionHandler) CreateComposition(w http.ResponseWriter, r *http.Request) {
	var comp models.FattyAcidComposition
	if err := json.NewDecoder(r.Body).Decode(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.svc.Create(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusCreated, comp, r.URL.Query().Get("REQUEST_ID"))
}

// GetCompositionByID — GET /compositions/fatty/{id}
func (h *FattyAcidCompositionHandler) GetCompositionByID(w http.ResponseWriter, r *http.Request) {
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
	util.WriteJSON(w, http.StatusOK, comp, r.URL.Query().Get("REQUEST_ID"))
}

// UpdateComposition — PUT /compositions/fatty/{id}
func (h *FattyAcidCompositionHandler) UpdateComposition(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	var comp models.FattyAcidComposition
	if err := json.NewDecoder(r.Body).Decode(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comp.ID = int32(id)
	if err := h.svc.Update(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, comp, r.URL.Query().Get("REQUEST_ID"))
}

// DeleteComposition — DELETE /compositions/fatty/{id}
func (h *FattyAcidCompositionHandler) DeleteComposition(w http.ResponseWriter, r *http.Request) {
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
	util.WriteJSON(w, http.StatusOK, map[string]string{"message": "Запись удалена"}, r.URL.Query().Get("REQUEST_ID"))
}

// ListCompositionsByProduct — GET /compositions/fatty?product_id=...
func (h *FattyAcidCompositionHandler) ListCompositionsByProduct(w http.ResponseWriter, r *http.Request) {
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
	util.WriteJSON(w, http.StatusOK, comps, r.URL.Query().Get("REQUEST_ID"))
}
