package handler

import (
	"github.com/lelikptz/gitlab-webhook-notifier/internal/response"
	"net/http"
)

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	response.NotFoundErrorResponse(w, r)
}
