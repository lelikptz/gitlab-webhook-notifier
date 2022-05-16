package main

import (
	"github.com/joho/godotenv"
	"github.com/lelikptz/gitlab-webhook-notifier/internal/handler"
	"github.com/lelikptz/gitlab-webhook-notifier/internal/notifier"
	"github.com/lelikptz/gitlab-webhook-notifier/internal/webhook"
	"log"
	"net/http"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	log.Println("Start web server")

	http.HandleFunc("/webhook", setup())
	http.HandleFunc("/", handler.HandleNotFound)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func setup() func(writer http.ResponseWriter, request *http.Request) {
	return handler.NewWebHookHandler(
		notifier.NewTelegramNotifier(),
		webhook.NewRequestParser(),
	).Webhook
}
