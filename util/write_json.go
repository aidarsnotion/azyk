package util

import (
	"azyk/internal/domain/responsebody"
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON[T any](w http.ResponseWriter, status int, data T, requestId string) {
	var response responsebody.Response[T]
	response.RequestId = requestId
	response.Code = status

	switch status {
	case http.StatusOK:
		response.Data = data
		response.Message = "Success"
	case http.StatusCreated:
		response.Data = data
		response.Message = "Resource created"
	case http.StatusBadRequest:
		response.Message = fmt.Sprintf("%v", data)
		response.Data = *new(T)
	case http.StatusInternalServerError:
		response.Message = fmt.Sprintf("%v", data)
		response.Data = *new(T)
	default: // По хорошему нужно пройтись по другим статусам тоже и сделать обработку
		response.Message = fmt.Sprintf("Unhandled status: %v", status)
		response.Data = *new(T)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)
}
