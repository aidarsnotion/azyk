package util

import "net/http"

func SuccessResponse[T any](w http.ResponseWriter, status int, data T, requestId, cmd string) {
	WriteJSONresponse(WriteJSONParams[T]{
		W:         w,
		Status:    status,
		Data:      data,
		RequestID: requestId,
		Cmd:       cmd,
	})
}

func ErrorResponse(w http.ResponseWriter, status int, message, requestId, cmd string) {
	WriteJSONresponse(WriteJSONParams[any]{
		W:         w,
		Status:    status,
		Message:   message,
		RequestID: requestId,
		Cmd:       cmd,
	})
}
