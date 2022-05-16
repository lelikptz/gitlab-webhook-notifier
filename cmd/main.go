package main

import (
	"github.com/lelikptz/gitlab-webhook-notifier/internal/handler"
	"log"
	"net/http"
)

func main() {
	log.Println("Start web server")

	http.HandleFunc("/", handler.ErrorHandler)
	http.HandleFunc("/webhook", handler.SuccessHandler)
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	}
}
