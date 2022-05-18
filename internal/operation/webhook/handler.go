package webhook

import (
	"log"
	"net/http"

	"github.com/lelikptz/gitlab-webhook-notifier/internal/response"
)

type Handler struct {
	Notifier
	RequestParser
}

func NewHandler(notifier Notifier, requestParser RequestParser) *Handler {
	return &Handler{notifier, requestParser}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("method: %v, request uri: %v", r.Method, r.RequestURI)

	if r.Method != http.MethodPost {
		response.NotFoundErrorResponse(w, r)
		return
	}

	payload, err := h.RequestParser.GetPayload(r)
	if err != nil {
		response.ErrorResponse("bad request", http.StatusBadRequest, w, r)
		return
	}

	go h.Notifier.Send(payload.GetMessage())

	response.SuccessHandler("OK", w, r)
}
