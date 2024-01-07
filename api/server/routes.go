package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stephen-pp/instajot/validators"
)

func RegisterRoutes(app *fiber.App) {
	// Registration route
	app.Post("/register", func(c *fiber.Ctx) error {
		user := new(validators.UserRegistration)

		// Parse the request body into the user registration struct
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(validators.FailureResponse{
				Success: false,
				Data:    validators.EmptyStruct{},
				Error:   err.Error(),
			})
		}

		// Validate the user registration struct
		err := validators.Validate(user)
		if len(err) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(validators.FailureResponse{
				Success: false,
				Data:    validators.EmptyStruct{},
				Error:   validators.ParseValidationErrors(err),
			})
		}

		// Return a response
		return c.Status(fiber.StatusOK).JSON(validators.SuccessResponse{
			Success: true,
			Data:    "Hello, World!",
			Error:   validators.EmptyStruct{},
		})
	})

	// Login route
	app.Post("/login", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Logout route
	app.Get("/logout", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
