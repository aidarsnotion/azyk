package http

import (
	"azyk/internal/domain/models"
	"azyk/internal/domain/requestbody"
	"azyk/internal/usecase"
	"azyk/util"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"net/http"
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

	if req.Cmd != "CreateRegion" {
		util.ErrorResponse(w, http.StatusBadRequest, "cmd should be CreateRegion", req.RequestID, req.Cmd)
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
	var req requestbody.BaseRequest // здесь только RequestID, Cmd и пагинация

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), req.RequestID, req.Cmd)
	}

	if req.Cmd != "GetRegionByID" {
		util.ErrorResponse(w, http.StatusBadRequest, "cmd should be GetRegionByID", req.RequestID, req.Cmd)
	}

	region, err := h.usecase.GetByID(int32(req.ID))
	if err != nil {
		util.WriteJSON(w, http.StatusNotFound, err.Error(), req.RequestID)
		return
	}

	util.SuccessResponse(w, http.StatusCreated, region, req.RequestID, req.Cmd)
}

func (h *RegionHandler) List(w http.ResponseWriter, r *http.Request) {
	var req requestbody.BaseRequest // здесь только RequestID, Cmd и пагинация

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), req.RequestID, req.Cmd)
		return
	}

	if req.Cmd != "GetRegions" {
		util.ErrorResponse(w, http.StatusBadRequest, "cmd should be CreateRegion", req.RequestID, req.Cmd)
		return
	}

	if err := util.ValidateStruct(req); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), req.RequestID, req.Cmd)
		return
	}

	// Получаем список регионов
	regions, err := h.usecase.List()
	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, err.Error(), req.RequestID, req.Cmd)
		return
	}

	// Успешный ответ
	util.SuccessResponse(w, http.StatusOK, regions, req.RequestID, req.Cmd)
}

func (h *RegionHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req requestbody.RequestWithPayload[models.Region]

	// Декодинг запроса
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), req.RequestID, req.Cmd)
		return
	}

	if req.Cmd != "UpdateRegion" {
		util.ErrorResponse(w, http.StatusBadRequest, "cmd should be UpdateRegion", req.RequestID, req.Cmd)
		return
	}

	// Валидация структуры
	if err := util.ValidateStruct(req); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), req.RequestID, req.Cmd)
		return
	}

	// Обработка в usecase
	updatedRegion, err := h.usecase.Update(&req.Data)
	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, err.Error(), req.RequestID, req.Cmd)
		return
	}

	// Успешный ответ
	util.SuccessResponse(w, http.StatusOK, updatedRegion, req.RequestID, req.Cmd)
}

func (h *RegionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var req requestbody.RequestWithPayload[struct {
		ID int32 `json:"id"`
	}]

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), req.RequestID, req.Cmd)
		return
	}

	if req.Cmd != "DeleteRegion" {
		util.ErrorResponse(w, http.StatusBadRequest, "cmd should be DeleteRegion", req.RequestID, req.Cmd)
		return
	}

	if req.Data.ID == 0 {
		util.ErrorResponse(w, http.StatusBadRequest, "ID is required", req.RequestID, req.Cmd)
		return
	}

	deletedRegion, err := h.usecase.Delete(req.Data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.ErrorResponse(w, http.StatusNotFound, "Region not found", req.RequestID, req.Cmd)
			return
		}
		util.ErrorResponse(w, http.StatusInternalServerError, err.Error(), req.RequestID, req.Cmd)
		return
	}

	util.SuccessResponse(w, http.StatusOK, deletedRegion, req.RequestID, req.Cmd)
}
