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

// AminoAcidCompositionHandler обрабатывает HTTP-запросы для аминокислотного состава.
type AminoAcidCompositionHandler struct {
	service usecase.AminoAcidCompositionService
}

// NewAminoAcidCompositionHandler создаёт новый обработчик.
func NewAminoAcidCompositionHandler(svc usecase.AminoAcidCompositionService) *AminoAcidCompositionHandler {
	return &AminoAcidCompositionHandler{service: svc}
}

// CreateComposition — POST /compositions/amino
func (h *AminoAcidCompositionHandler) CreateComposition(w http.ResponseWriter, r *http.Request) {
	var composition models.AminoAcidComposition
	if err := json.NewDecoder(r.Body).Decode(&composition); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(&composition); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	util.WriteJSON(w, http.StatusCreated, composition)
}

// GetCompositionByID — GET /compositions/amino/:id
func (h *AminoAcidCompositionHandler) GetCompositionByID(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	composition, err := h.service.GetByID(int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	util.WriteJSON(w, http.StatusOK, composition)
}

// UpdateComposition — PUT /compositions/amino/:id
func (h *AminoAcidCompositionHandler) UpdateComposition(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var composition models.AminoAcidComposition
	if err := json.NewDecoder(r.Body).Decode(&composition); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	composition.ID = int32(id)
	if err := h.service.Update(&composition); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, composition)
}

// DeleteComposition — DELETE /compositions/amino/:id
func (h *AminoAcidCompositionHandler) DeleteComposition(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Delete(int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// ListCompositionsByProduct — GET /compositions/amino?product_id=...
func (h *AminoAcidCompositionHandler) ListCompositionsByProduct(w http.ResponseWriter, r *http.Request) {
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
