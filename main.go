package main

import (
	"net/http"
	"log"
)


func main() {
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		log.Default().Print("Health check")
	})

	http.HandleFunc("/api/username", func(w http.ResponseWriter, r *http.Request) {
		log.Default().Print("Username check")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}