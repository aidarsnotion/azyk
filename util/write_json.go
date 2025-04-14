package util

import (
	"azyk/internal/domain/responsebody"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func WriteJSON( /*r *http.Request.Id,*/ w http.ResponseWriter, status int, data interface{}) {
	var response responsebody.ResponseBody
	response.Id = "1"
	response.Code = strconv.Itoa(status)
	if status != 0 {
		response.Message = fmt.Sprintf("%v", data)
		response.Data = []interface{}{}
	} else {
		response.Data = data
		response.Message = "Success"
	}
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)
}
