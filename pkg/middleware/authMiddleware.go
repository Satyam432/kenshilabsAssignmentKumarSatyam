package middleware

// import (
// 	"github.com/gofiber/fiber/v2"
// )

// func AuthenticateUser() func(*fiber.Ctx) error {

// }

// func AssignAuthentication() func(*fiber.Ctx) error {
// 	return func(c *fiber.Ctx) error {
// 		// Define a struct to hold the request body data
// 		userData := struct {
// 			Password string `json:"password"`
// 			Email    string `json:"email"`
// 		}{}

// 		// Parse the request body
// 		if err := c.BodyParser(&userData); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
// 		}

// 		// Set the email in the context locals
// 		c.Locals("email", userData.Email)

// 		// Continue with the next handler
// 		return c.Next()
// 	}
// }
