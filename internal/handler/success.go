package handler

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Message string `json:"message"`
}

func SuccessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := SuccessResponse{Message: "OK"}

	json.NewEncoder(w).Encode(response)
}
