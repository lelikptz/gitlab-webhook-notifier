package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func ErrorResponse(text string, code int, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := Error{Error: text, Code: code}
	json.NewEncoder(w).Encode(response)
}

func NotFoundErrorResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse("Page not found", http.StatusNotFound, w, r)
}
