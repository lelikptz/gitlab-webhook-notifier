package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	response := ErrorResponse{Error: "Page not found", Code: http.StatusNotFound}
	json.NewEncoder(w).Encode(response)
}
