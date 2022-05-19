package main

import (
	"log"
	"net/http"
	"os"

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

	http.HandleFunc("/webhook", withAuth(setupWebhook()))
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

func withAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Gitlab-Token")
		if token == os.Getenv("GITLAB_TOKEN") {
			next.ServeHTTP(w, r)
			return
		}
		response.ErrorResponse("Unauthorized", http.StatusUnauthorized, w, r)
	})
}
