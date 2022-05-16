package handler

import (
	"github.com/lelikptz/gitlab-webhook-notifier/internal/response"
	"log"
	"net/http"
)

type WebHookHandler struct {
	Notifier
	Parser
}

func NewWebHookHandler(notifier Notifier, parser Parser) *WebHookHandler {
	return &WebHookHandler{notifier, parser}
}

func (h *WebHookHandler) Webhook(w http.ResponseWriter, r *http.Request) {
	log.Printf("method: %v, request uri: %v", r.Method, r.RequestURI)

	if r.Method != http.MethodPost {
		response.NotFoundErrorResponse(w, r)
		return
	}

	payload, err := h.Parser.GetPayload(r)
	if err != nil {
		response.ErrorResponse("bad request", http.StatusBadRequest, w, r)
		return
	}

	go h.Send(payload.GetMessage())

	response.SuccessHandler("OK", w, r)
}
