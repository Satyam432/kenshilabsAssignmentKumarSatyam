package controller

import (
	"example.com/m/app/models"
	"example.com/m/pkg/utils"
	"example.com/m/platform/database"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid"
	"go.mongodb.org/mongo-driver/bson"
)

func UserSignUp(c *fiber.Ctx) error {
	userRequestData := new(models.SigninRequest)
	if err := c.BodyParser(userRequestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// create alphanumeric nanois: userid
	userId, _ := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 20)

	// Create the User struct
	user := models.User{
		UserId:   userId,
		Email:    userRequestData.Email,
		Password: userRequestData.Password,
	}
	// Save the user to the database
	g := database.GetMongoCLient()
	_, err := g.Database("UserTask").Collection("Users").InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return a 201 status code with the created user
	return c.Status(fiber.StatusCreated).JSON(user)
}

func UserSignIn(c *fiber.Ctx) error {
	//Fetch userBody
	userRequestData := new(models.SigninRequest)
	if err := c.BodyParser(userRequestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//Extract email and password from userBody
	email := userRequestData.Email
	g := database.GetMongoCLient()
	var User models.User
	filter := bson.M{"email": email}
	//Check if User Exists by emailId
	err := g.Database("UserTask").Collection("Users").FindOne(c.Context(), filter).Decode(&User)
	//If user does not exist, return error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	//If user exists, check if password is correct
	if User.Password != userRequestData.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}
	//If user exists create a JWT token
	token, errToken := utils.GenerateToke(User)
	if errToken != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errToken.Error(),
		})
	}
	//Return a 200 status code with the JWT token
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// UserSignOut handles the user sign-out process and invalidates the token.
func UserSignOut(c *fiber.Ctx) error {
	// Retrieve the token from the request (e.g., from Authorization header)
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}

	// fmt.Println(authHeader)

	// // Typically the token is in the format "Bearer <token>"
	// parts := strings.Split(authHeader, " ")
	// if len(parts) != 2 || parts[0] != "Bearer" {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "Invalid authorization header format",
	// 	})
	// }
	// tokenString := parts[1]

	// Call invalidate token function
	newTokenString, err := utils.InvalidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return a success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Signed out successfully",
		"token":   newTokenString, // Optionally return the invalidated token
	})
}
