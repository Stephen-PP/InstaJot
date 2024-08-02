package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/stephen-pp/instajot/internal/api"
	"github.com/stephen-pp/instajot/internal/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Start database
	dbUrl := os.Getenv("API_DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "instajot.sqlite3"
	}
	err = database.InitDB(dbUrl)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	server := api.APIServer{
		DB: database.GetDB(),
	}

	serverPort := os.Getenv("API_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	server.Start(os.Getenv("API_IP"), serverPort)
}
