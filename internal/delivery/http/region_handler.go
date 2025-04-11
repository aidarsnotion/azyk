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

type RegionHandler struct {
	usecase usecase.RegionUsecase
}

func NewRegionHandler(uc usecase.RegionUsecase) *RegionHandler {
	return &RegionHandler{usecase: uc}
}

func (h *RegionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var region models.Region
	if err := json.NewDecoder(r.Body).Decode(&region); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.usecase.Create(&region); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusCreated, region)
}

func (h *RegionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	region, err := h.usecase.GetByID(int32(id))
	if err != nil {
		http.Error(w, "Регион не найден", http.StatusNotFound)
		return
	}
	util.WriteJSON(w, http.StatusOK, region)
}

func (h *RegionHandler) List(w http.ResponseWriter, r *http.Request) {
	regions, err := h.usecase.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, regions)
}

func (h *RegionHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат id", http.StatusBadRequest)
		return
	}
	var region models.Region
	if err := json.NewDecoder(r.Body).Decode(&region); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	region.ID = int32(id)
	if err := h.usecase.Update(&region); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, region)
}

func (h *RegionHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	util.WriteJSON(w, http.StatusOK, map[string]string{"message": "Регион удалён"})
}
