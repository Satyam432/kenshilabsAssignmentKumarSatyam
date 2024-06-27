package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	println("Request received:", c.Path())
	return c.Next()
}

// func ErrorHandler(c *fiber.Ctx, err error) error {
// 	// Log the error
// 	fmt.Println("Error:", err)

// 	// Return a generic error response
// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 		"error": "Internal Server Error",
// 	})
// }

func ErrorHandler(err error) {
	// Log the error
	fmt.Println("Error:", err)
}
