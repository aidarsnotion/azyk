package util

import "net/http"

type responseMapping struct {
	InternalCode int
	Message      string
}

var responseCodeMap = map[int]responseMapping{
	http.StatusOK: {
		InternalCode: 100,
		Message:      "Success",
	},
	http.StatusCreated: {
		InternalCode: 101,
		Message:      "Resource created",
	},
	http.StatusAccepted: {
		InternalCode: 102,
		Message:      "Request accepted and processing",
	},
	http.StatusBadRequest: {
		InternalCode: 200,
		Message:      "Bad request",
	},
	http.StatusUnauthorized: {
		InternalCode: 201,
		Message:      "Unauthorized",
	},
	http.StatusForbidden: {
		InternalCode: 202,
		Message:      "Forbidden",
	},
	http.StatusNotFound: {
		InternalCode: 203,
		Message:      "Resource not found",
	},
	http.StatusConflict: {
		InternalCode: 204,
		Message:      "Conflict: resource already exists or invalid state",
	},
	http.StatusInternalServerError: {
		InternalCode: 300,
		Message:      "Internal server error",
	},
	http.StatusServiceUnavailable: {
		InternalCode: 301,
		Message:      "Service temporarily unavailable",
	},
}
