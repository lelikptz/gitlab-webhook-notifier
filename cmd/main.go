package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/lelikptz/gitlab-webhook-notifier/internal/operation/webhook"
	"github.com/lelikptz/gitlab-webhook-notifier/internal/response"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	log.Println("Start web server")

	http.HandleFunc("/webhook", setupWebhook())
	http.HandleFunc("/", response.NotFoundErrorResponse)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func setupWebhook() func(writer http.ResponseWriter, request *http.Request) {
	return webhook.NewHandler(
		*webhook.NewNotifier(),
		*webhook.NewRequestParser(),
	).Handle
}
