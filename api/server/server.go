package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/stephen-pp/instajot/database"
)

func CreateServer() {
	fmt.Println("Creating SQLite database...")
	db := database.CreateDatabase()

	fmt.Println("Creating server...")
	app := fiber.New()

	fmt.Println("Registering routes...")
	RegisterRoutes(app)

	// Determine port to use
	port := ""
	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	} else {
		port = ":3000"
	}

	fmt.Printf("Listening on port %s...\n", port)
	app.Listen(port)

	// Close the database connection when the server shuts down
	app.Hooks().OnShutdown(func() error {
		db.Close()
		return nil
	})
}
