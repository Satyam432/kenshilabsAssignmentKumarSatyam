package main

import (
	"example.com/m/pkg/routes"
	"example.com/m/pkg/utils"
	"example.com/m/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	database.InitializeMongoDB()

	app := fiber.New()
	app.Use(recover.New())
	app.Use(utils.Logger)
	routes.UserRoutes(app)
	routes.TaskRoutes(app)

	app.Listen(":3000")
	// app.Listen(":3000")
}
