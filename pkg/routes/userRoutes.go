package routes

import (
	"example.com/m/app/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/signup", controller.UserSignUp)
	route.Post("/signin", controller.UserSignIn)
	route.Post("/signout", controller.UserSignOut)

}
