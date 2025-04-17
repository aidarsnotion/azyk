package util

import (
	"azyk/internal/domain/responsebody"
	"encoding/json"
	"fmt"
	"net/http"
)

type WriteJSONParams[T any] struct {
	W           http.ResponseWriter
	Status      int
	Data        T
	RequestID   string
	Cmd         string
	Total       int
	TotalPages  int
	CurrentPage int
	Message     string // если хочешь переопределить текст
}

func WriteJSON[T any](w http.ResponseWriter, status int, data T, requestId string) {
	var response responsebody.Response[any]
	response.RequestId = requestId
	response.Code = status

	var message string
	var result any = data

	switch status {
	case http.StatusOK:
		message = "Success"

	case http.StatusCreated:
		message = "Resource created"

	case http.StatusAccepted:
		message = "Request accepted and processing"

	case http.StatusNoContent:
		w.WriteHeader(http.StatusNoContent)
		return

	case http.StatusBadRequest:
		message = fmt.Sprintf("Bad request: %v", data)
		result = []any{}

	case http.StatusUnauthorized:
		message = "Unauthorized"
		result = []any{}

	case http.StatusForbidden:
		message = "Forbidden"
		result = []any{}

	case http.StatusNotFound:
		message = "Resource not found"
		result = []any{}

	case http.StatusConflict:
		message = "Conflict: resource already exists or state invalid"
		result = []any{}

	case http.StatusUnprocessableEntity:
		message = "Unprocessable entity"
		result = []any{}

	case http.StatusTooManyRequests:
		message = "Too many requests"
		result = []any{}

	case http.StatusInternalServerError:
		message = fmt.Sprintf("Internal error: %v", data)
		result = []any{}

	case http.StatusServiceUnavailable:
		message = "Service temporarily unavailable"
		result = []any{}

	default:
		message = fmt.Sprintf("Unhandled status code: %d", status)
		result = []any{}
	}

	response.Message = message
	response.Data = result

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)
}

func WriteJSONresponse[T any](params WriteJSONParams[T]) {
	w := params.W

	var response responsebody.Response[any] // <<--- ВСЕГДА any для правильного пустого [] при ошибке
	response.RequestId = params.RequestID
	response.Cmd = params.Cmd

	mapping, ok := responseCodeMap[params.Status]
	if !ok {
		mapping = responseMapping{
			InternalCode: 999,
			Message:      fmt.Sprintf("Unhandled error: %d", params.Status),
		}
	}

	var resultData any
	if params.Status >= 400 {
		resultData = []any{}
		response.Message = params.Message
	} else {
		resultData = params.Data
		response.Message = mapping.Message
	}

	response.Code = mapping.InternalCode
	response.Data = resultData

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(params.Status)
	_ = json.NewEncoder(w).Encode(response)
}
