package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/stephen-pp/instajot/database"
	"github.com/stephen-pp/instajot/validators"
)

func CreateServer() {
	fmt.Println("Creating SQLite database...")
	db := database.CreateDatabase()

	fmt.Println("Creating server...")
	app := fiber.New(fiber.Config{
		// Set custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(validators.FailureResponse{
				Success: false,
				Data:    validators.EmptyStruct{},
				Error:   err.Error(),
			})
		},
	})

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
