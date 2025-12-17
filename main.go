package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		// Check QtraderDB
		qtraderSQL, err := database.GetQtraderDB().DB()
		qtraderHealthy := err == nil && qtraderSQL.Ping() == nil

		// Check TradeDB
		tradeSQL, err := database.GetTradeDB().DB()
		tradeHealthy := err == nil && tradeSQL.Ping() == nil

		// Check TradeDB
		marketSQL, err := database.GetMarketDB().DB()
		marketHealthy := err == nil && marketSQL.Ping() == nil

		status := "healthy"
		httpStatus := fiber.StatusOK

		if !qtraderHealthy || !tradeHealthy || !marketHealthy {
			status = "unhealthy"
			httpStatus = fiber.StatusServiceUnavailable
		}

		return c.Status(httpStatus).JSON(fiber.Map{
			"status": status,
			"databases": fiber.Map{
				"qtraderdb": qtraderHealthy,
				"tradedb":   tradeHealthy,
				"marketdb":  marketHealthy,
			},
		})
	})

	// Protected routes group with authentication middleware
	api := app.Group("/api")
	api.Use(middleware.AuthMiddleware())

	// Define the User routes under protected group
	routes.UserHandlers(api.Group("/users"))
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to all databases
	database.ConnectDatabases()

	// Initialize Redis
	helpers.InitRedis()

	// Defer closing database connections
	defer func() {
		if err := database.CloseDatabases(); err != nil {
			log.Printf("Error closing databases: %v", err)
		}
	}()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "qTrader OMS API",
	})

	// Setup routes
	setupRoutes(app)

	// Channel to listen for interrupt signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := app.Listen(":9000"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	log.Println("Server started on :9000")

	// Wait for interrupt signal
	<-c
	log.Println("Shutting down gracefully...")

	// Shutdown the server
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error during shutdown: %v", err)
	}

	log.Println("Server shutdown complete")
}
