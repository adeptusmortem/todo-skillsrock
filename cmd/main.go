package main

import (
	"log"
	"todo-list/database"
	"todo-list/routes"

	"github.com/gofiber/fiber/v2"
)

// const (
// 	sConnection string = "path_to_db"
// )

func main() {

	// Connect to DB
	if err := database.Connect(sConnection); err != nil {
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