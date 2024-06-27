package routes

import (
	"example.com/m/app/controller"
	"example.com/m/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(a *fiber.App) {
	route := a.Group("/")

	// Task Management Endpoints
	route.Post("/tasks", middleware.AuthenticateUser(), controller.PostTask)
	route.Get("/tasks", middleware.AuthenticateUser(), controller.GetTask)
	route.Get("/tasks/:id", middleware.AuthenticateUser(), controller.GetTaskById)
	route.Put("/tasks/:id", middleware.AuthenticateUser(), controller.UpdateTask)
	route.Delete("/tasks/:id", middleware.AuthenticateUser(), controller.DeleteTask)
}
