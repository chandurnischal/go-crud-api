package main

import (
	"backend/database"
	"backend/handlers"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer database.DB.Close()

	app.Get("/", handlers.CheckDatabaseConnection)
	app.Post("/create", handlers.CreateUserHandler)
	app.Get("/retrieve", handlers.RetrieveByNameHandler)
	app.Delete("/delete", handlers.DeleteUserHandler)

	log.Fatal(app.Listen(":8000"))
}
