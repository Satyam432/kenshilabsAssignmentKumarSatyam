package middleware

import (
	"example.com/m/pkg/utils"
	"example.com/m/platform/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AuthenticateUser() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Get the token from the Authorization header
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or malformed JWT",
			})
		}

		// Extract the email from the token
		email, err := utils.ExtractEmailFromToken(tokenString)
		if err != nil {
			utils.ErrorHandler(err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired JWT",
			})
		}

		// Connect to MongoDB
		g := database.GetMongoCLient()
		// Check if the user exists in MongoDB
		collection := g.Database("UserTask").Collection("Users")
		var result bson.M
		err = collection.FindOne(c.Context(), bson.M{"email": email}).Decode(&result)
		if err != nil {
			utils.ErrorHandler(err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// If user exists, set user info in the context and proceed
		c.Locals("user", result)
		return c.Next()
	}
}
