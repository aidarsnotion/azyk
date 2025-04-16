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

// ChemicalCompositionHandler обрабатывает HTTP-запросы для ChemicalComposition.
type ChemicalCompositionHandler struct {
	service usecase.ChemicalCompositionService
}

// NewChemicalCompositionHandler создаёт новый обработчик.
func NewChemicalCompositionHandler(svc usecase.ChemicalCompositionService) *ChemicalCompositionHandler {
	return &ChemicalCompositionHandler{service: svc}
}

// CreateComposition — POST /compositions/chemical
func (h *ChemicalCompositionHandler) CreateComposition(w http.ResponseWriter, r *http.Request) {
	var comp models.ChemicalComposition
	if err := json.NewDecoder(r.Body).Decode(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusCreated, comp, r.URL.Query().Get("REQUEST_ID"))
}

// GetCompositionByID — GET /compositions/chemical/{id}
func (h *ChemicalCompositionHandler) GetCompositionByID(w http.ResponseWriter, r *http.Request) {
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
	util.WriteJSON(w, http.StatusOK, comp, r.URL.Query().Get("REQUEST_ID"))
}

// UpdateComposition — PUT /compositions/chemical/{id}
func (h *ChemicalCompositionHandler) UpdateComposition(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	var comp models.ChemicalComposition
	if err := json.NewDecoder(r.Body).Decode(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comp.ID = int32(id)
	if err := h.service.Update(&comp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, comp, r.URL.Query().Get("REQUEST_ID"))
}

// DeleteComposition — DELETE /compositions/chemical/{id}
func (h *ChemicalCompositionHandler) DeleteComposition(w http.ResponseWriter, r *http.Request) {
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
	util.WriteJSON(w, http.StatusOK, map[string]string{"message": "Запись удалена"}, r.URL.Query().Get("REQUEST_ID"))
}

// ListCompositionsByProduct — GET /compositions/chemical?product_id=...
func (h *ChemicalCompositionHandler) ListCompositionsByProduct(w http.ResponseWriter, r *http.Request) {
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
	util.WriteJSON(w, http.StatusOK, comps, r.URL.Query().Get("REQUEST_ID"))
}
