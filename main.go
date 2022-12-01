package main

import (
	"net/http"
	"log"

	"github.com/hmuhammadazeem/user-service/app/routing"
)


func main() {
	// update later, after exploring graceful shutdown issues
	http.HandleFunc("/api/health", routing.HealthHandler())
	http.HandleFunc("/api/username", routing.CheckUsernameHandler())

	log.Default().Print("Started listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}