package main

import (
	"log"

	"github.com/ayonqfl/go-fiber-gorm/database"
	"github.com/ayonqfl/go-fiber-gorm/helpers"
	"github.com/ayonqfl/go-fiber-gorm/middleware"
	"github.com/ayonqfl/go-fiber-gorm/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to OMS API's")
}

func setupRoutes(app *fiber.App) {
	// Public routes (no authentication)
	app.Get("/api", Welcome)

	// Protected routes group with authentication middleware
	api := app.Group("/api")
	api.Use(middleware.AuthMiddleware())

	// Define the User routes under protected group
	routes.UserHandlers(api.Group("/users"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDb()
	helpers.InitRedis()

	app := fiber.New(fiber.Config{
		AppName: "qTrader OMS API",
	})

	setupRoutes(app)
	log.Fatal(app.Listen(":9000"))
}
