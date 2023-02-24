package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func MakeRespon(w http.ResponseWriter, status int, msg string, res interface{}) {
	w.Header().Set("Content-type", "application/json")
	var response Response
	response.Status = http.StatusText(status)
	response.Message = msg
	response.Data = res
	userJson, _ := json.Marshal(response)
	w.WriteHeader(status)
	w.Write(userJson)
}
