package http

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/requestbody"
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
	var req requestbody.RequestWithPayload[models.Region]

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), req.RequestID, req.Cmd)
		return
	}

	if err := util.ValidateStruct(req); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), req.RequestID, req.Cmd)
		return
	}

	if err := h.usecase.Create(&req.Data); err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, err.Error(), req.RequestID, req.Cmd)
		return
	}

	// Успешный ответ
	util.SuccessResponse(w, http.StatusCreated, req.Data, req.RequestID, req.Cmd)
}

func (h *RegionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, err.Error(), idStr)
		return
	}
	region, err := h.usecase.GetByID(int32(id))
	if err != nil {
		util.WriteJSON(w, http.StatusNotFound, err.Error(), idStr)
		return
	}
	util.WriteJSON(w, http.StatusOK, region, idStr)
}

func (h *RegionHandler) List(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	regions, err := h.usecase.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(w, http.StatusOK, regions, idStr)
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
	util.WriteJSON(w, http.StatusOK, region, idStr)
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
	util.WriteJSON(w, http.StatusOK, map[string]string{"message": "Регион удалён"}, idStr)
}
