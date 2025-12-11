package main

import (
	"log"

	"github.com/ayonqfl/go-fiber-gorm/database"
	"github.com/ayonqfl/go-fiber-gorm/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Welcome(C *fiber.Ctx) error {
	return C.SendString("Welcome to OMS API's")
}

func setupRoutes(app *fiber.App) {
	// Wellcome endpoint
	app.Get("/api", Welcome)

	// User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)
	log.Fatal(app.Listen(":9000"))
}
