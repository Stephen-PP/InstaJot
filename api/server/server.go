package server

import (
	"fmt"

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

	app.Listen(":3000")

	// Close the database connection when the server shuts down
	app.Hooks().OnShutdown(func() error {
		db.Close()
		return nil
	})
}
