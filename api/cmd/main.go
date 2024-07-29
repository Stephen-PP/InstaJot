package main

import (
	"log"
	"net/http"
)

func main() {
	// TODO: Initialize database
	// TODO: Setup routes
	// TODO: Configure middleware

	log.Println("Starting secure notes server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
