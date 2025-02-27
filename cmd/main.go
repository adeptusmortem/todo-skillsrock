package main

import (
	"log"
	"todo-list/database"
	"todo-list/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Connect to DB
	if err := database.Connect("path_to_db"); err != nil {
		log.Fatalf("Error on conecting to DB: %v", err)
	}

	// Automigrate
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("Error on DB Migration: %v", err)
	}

	// Make fiber app
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	// Register correct routes
	routes.RegisterRoutes(app)

    log.Fatal(app.Listen(":3000"))
}