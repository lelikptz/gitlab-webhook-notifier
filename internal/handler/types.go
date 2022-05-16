package handler

import "net/http"

type Notifier interface {
	Send(message string)
}

type Parser interface {
	GetPayload(r *http.Request) (Payload, error)
}

type Payload interface {
	GetMessage() string
}
