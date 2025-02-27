package handlers

import (
	"log"
	"strconv"
	"time"
	"todo-list/database"
	"todo-list/models"

	"github.com/gofiber/fiber/v2"
)

// Create new task
func CreateTask(c *fiber.Ctx) error {
	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		log.Printf("error 400: %v", err)
		return c.Status(400).JSON(fiber.Map{
            "error": "Bad Request",
        })
	}

	task.Created_at = time.Now()
	task.Updated_at = time.Now()

	result := database.DB.Create(&task)
	if result.Error != nil {
		log.Printf("error 500: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{
            "error": "Internal Server Error",
        })
	}

	return c.Status(201).JSON(task)
}

// Get specific task by id
func GetTask(c *fiber.Ctx) error {
	task := new(models.Task)
	result := database.DB.Where("ID = ?", c.Params("id")).First(&task)

	if result.Error != nil {
		log.Printf("err 404: %v", result.Error)
		return c.Status(404).JSON(fiber.Map{
            "error": "Not Found",
        })
	}

	return c.JSON(task)
}

// Get all tasks
func GetTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	result := database.DB.Find(&tasks)

	if result.Error != nil {
		log.Printf("err 404: %v", result.Error)
		return c.Status(404).JSON(fiber.Map{
            "error": "Not Found",
        })
	}

	return c.JSON(tasks)
}

// Update specific task
func UpdateTask(c *fiber.Ctx) error {
	var task models.Task
	
	if err := c.BodyParser(&task); err != nil {
		log.Printf("err 400: %v", err)
		return c.Status(400).JSON(fiber.Map{
            "error": "Bad Request",
        })
	}

	task.ID, _ = strconv.Atoi(c.Params("id"))
	result := database.DB.Save(&task)
	if result.Error != nil {
		log.Printf("err 500: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{
            "error": "Internal Server Error",
        })
	}

	return c.Status(200).JSON(task)
}

// Delete specific task
func DeleteTask(c *fiber.Ctx) error {
	var task models.Task
	result := database.DB.Delete(&task, c.Params("id"))

	if result.Error != nil {
		log.Printf("err 404: %v", result.Error)
		return c.Status(404).JSON(fiber.Map{
            "error": "Not Found",
        })
	}
	
	return c.Status(200).JSON(fiber.Map{
            "message": "Succesfully deleted!",
        })
}
