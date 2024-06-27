package main

import (
	"example.com/m/pkg/routes"
	"example.com/m/platform/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitializeMongoDB()

	app := fiber.New()

	routes.UserRoutes(app)
	routes.TaskRoutes(app)

	app.Listen(":3000")
	// app.Listen(":3000")
}
