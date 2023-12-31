package api

import (
	"github.com/gofiber/fiber/v2"
)

func CreateServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}