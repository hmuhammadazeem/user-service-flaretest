package main

import (
	"net/http"
	"log"

	"github.com/hmuhammadazeem/user-service/app/routing"
)


func main() {
	http.HandleFunc("/api/health", routing.HealthHandler())
	
	log.Default().Print("Started listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}