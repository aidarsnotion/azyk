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

type CategoryHandler struct {
	usecase usecase.CategoryUsecase
}

func NewCategoryHandler(uc usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{usecase: uc}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.usecase.Create(&category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusCreated, category)
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	category, err := h.usecase.GetByID(int32(id))
	if err != nil {
		http.Error(w, "Категория не найдена", http.StatusNotFound)
		return
	}
	util.WriteJSON(w, http.StatusOK, category)
}

func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	categories, err := h.usecase.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, categories)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	category.ID = int32(id)
	if err := h.usecase.Update(&category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, category)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	if err := h.usecase.Delete(int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, map[string]string{"message": "Категория удалена"})
}
