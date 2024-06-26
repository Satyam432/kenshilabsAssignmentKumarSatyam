package main

import (
	"example.com/m/platform/database"
	"github.com/gofiber/fiber/v3"
)

func main() {

	database.InitializeMongoDB()

	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
	// app.Listen(":3000")
}
