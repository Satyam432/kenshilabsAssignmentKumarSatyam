package controller

import (
	"example.com/m/app/models"
	"example.com/m/pkg/utils"
	"example.com/m/platform/database"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid"
)

func PostTask(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	email, err := utils.ExtractEmailFromToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}
	taskRequestData := new(models.CreateTask)
	if err := c.BodyParser(taskRequestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	taskId, _ := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 20)

	task := models.Task{
		TaskId:     taskId,
		TaskName:   taskRequestData.TaskName,
		TaskDetail: taskRequestData.TaskDetail,
		CreatedBY:  email,
	}

	// Save the Task to the database
	g := database.GetMongoCLient()
	_, errInserting := g.Database("UserTask").Collection("Tasks").InsertOne(c.Context(), task)
	if errInserting != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errInserting.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(task)
}

func GetTask(c *fiber.Ctx) error {
	return nil
}

func GetTaskById(c *fiber.Ctx) error {
	return nil
}

func UpdateTask(c *fiber.Ctx) error {
	return nil
}

func DeleteTask(c *fiber.Ctx) error {
	return nil
}
