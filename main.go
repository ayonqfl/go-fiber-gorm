package main

import (
	"log"

	"github.com/ayonqfl/go-fiber-gorm/database"
	"github.com/ayonqfl/go-fiber-gorm/middleware"
	"github.com/ayonqfl/go-fiber-gorm/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Welcome(C *fiber.Ctx) error {
	return C.SendString("Welcome to OMS API's")
}

func setupRoutes(app *fiber.App) {
	// Define the Wellcome routes
	app.Get("/api", Welcome)

	// Apply authentication middleware for protected routes
	protected := app.Use(middleware.AuthMiddleware())

	// Define the User routes
	routes.UserHandlers(protected.Group("/api/users"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDb()
	app := fiber.New(fiber.Config{
		AppName: "qTrader OMS API",
	})

	setupRoutes(app)
	log.Fatal(app.Listen(":9000"))
}
