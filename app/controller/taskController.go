package controller

import (
	"fmt"

	"example.com/m/app/models"
	"example.com/m/pkg/utils"
	"example.com/m/platform/database"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid"
	"go.mongodb.org/mongo-driver/bson"
)

func PostTask(c *fiber.Ctx) error {

	user := c.Locals("user").(bson.M)
	email := user["email"].(string)
	taskRequestData := new(models.CreateTask)
	if err := c.BodyParser(taskRequestData); err != nil {

		utils.ErrorHandler(err)
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
		utils.ErrorHandler(errInserting)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errInserting.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(task)
}

func GetTask(c *fiber.Ctx) error {

	user := c.Locals("user").(bson.M)
	email := user["email"].(string)
	// Get MongoDB client
	client := database.GetMongoCLient()

	// Define filter to find tasks by email
	filter := bson.M{"createdby": email}

	// Perform find operation to get all tasks for the user
	cursor, err := client.Database("UserTask").Collection("Tasks").Find(c.Context(), filter)
	if err != nil {
		utils.ErrorHandler(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer cursor.Close(c.Context())

	// Iterate through the cursor and collect tasks
	var tasks []bson.M
	for cursor.Next(c.Context()) {
		var task bson.M
		if err := cursor.Decode(&task); err != nil {
			utils.ErrorHandler(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		tasks = append(tasks, task)
	}

	// Check if any tasks were found
	if len(tasks) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No tasks found for the user",
		})
	}

	// Return tasks in JSON response
	return c.Status(fiber.StatusOK).JSON(tasks)
}

func GetTaskById(c *fiber.Ctx) error {
	taskId := c.Params("id")
	client := database.GetMongoCLient()
	filter := bson.M{"taskid": taskId}
	var task models.Task
	err := client.Database("UserTask").Collection("Tasks").FindOne(c.Context(), filter).Decode(&task)
	//If user does not exist, return error
	if err != nil {
		utils.ErrorHandler(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	taskId := c.Params("id")
	taskUpdate := new(models.CreateTask) // Adjust this based on your task update structure
	if err := c.BodyParser(taskUpdate); err != nil {
		utils.ErrorHandler(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	client := database.GetMongoCLient()
	filter := bson.M{"taskid": taskId}

	//Update
	updateFilter := bson.M{
		"$set": bson.M{
			"taskname":   taskUpdate.TaskName,
			"taskdetail": taskUpdate.TaskDetail,
		},
	}

	// Perform update operation
	result, err := client.Database("UserTask").Collection("Tasks").UpdateOne(c.Context(), filter, updateFilter)
	if err != nil {
		utils.ErrorHandler(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Check if the task was updated successfully
	if result.ModifiedCount == 0 {
		utils.ErrorHandler(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found or no changes applied",
		})
	}

	// Return success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task updated successfully",
	})

}

func DeleteTask(c *fiber.Ctx) error {
	taskId := c.Params("id")
	// Get MongoDB client
	client := database.GetMongoCLient()

	// Define filter to find the task by ID
	filter := bson.M{"taskid": taskId}

	// Perform delete operation
	result, err := client.Database("UserTask").Collection("Tasks").DeleteOne(c.Context(), filter)
	if err != nil {
		utils.ErrorHandler(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		utils.ErrorHandler(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	// Return success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Task with ID %s deleted successfully", taskId),
	})

}
