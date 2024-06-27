package routes

import (
	"example.com/m/app/controller"
	"example.com/m/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(a *fiber.App) {
	route := a.Group("/")

	route.Post("/signup", controller.UserSignUp)
	route.Post("/signin", controller.UserSignIn)
	route.Post("/signout", controller.UserSignOut)

	// Task Management Endpoints
	route.Post("/tasks", middleware.AuthenticateUser(), controller.PostTask)
	route.Get("/tasks", controller.GetTask)
	route.Get("/tasks/:id", controller.GetTaskById)
	route.Put("/tasks/:id", controller.UpdateTask)
	route.Delete("/tasks/:id", controller.DeleteTask)
}
