package routes

import (
	"todo-list/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/task", handlers.CreateTask)
	api.Get("/task", handlers.GetTasks)
	api.Get("/task/:id", handlers.GetTask)
	api.Put("/task/:id", handlers.UpdateTask)
	api.Delete("/task/:id", handlers.DeleteTask)
}