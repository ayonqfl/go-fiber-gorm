package middleware

import (
	"github.com/ayonqfl/go-fiber-gorm/database"
	"github.com/ayonqfl/go-fiber-gorm/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// AuthMiddleware creates a new authentication middleware
func AuthMiddleware() fiber.Handler {
	// Paths that should skip authentication
	skipPaths := []string{
		"/auth/login",
		"/auth/register",
		"/public",
	}

	// Keywords in path that should skip authentication
	skipKeywords := []string{
		"auth",
		"public",
		"system",
	}

	// Get JWT secret from environment variable
	jwtSecret := helpers.GetJWTSecret()

	return func(c *fiber.Ctx) error {
		path := c.Path()

		// Check if path should skip authentication
		if helpers.ShouldSkipAuth(path, skipPaths, skipKeywords) {
			log.Info("Skipping auth for path: ", path)
			return c.Next()
		}

		// Extract token from Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Warn("Missing Authorization header")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "fail",
				"message": "Invalid token",
			})
		}

		// Extract Bearer token
		token := helpers.ExtractBearerToken(authHeader)
		if token == "" {
			log.Warn("Invalid Authorization header format")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "fail",
				"message": "Invalid token",
			})
		}

		// Validate token
		userData, err := helpers.ValidateToken(token, jwtSecret)
		if err != nil {
			log.Warnf("Token validation failed: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "fail",
				"message": "Invalid token",
			})
		}

		// Verify user exists in database and is active
		if err := helpers.VerifyUserExists(database.Database.Db, userData.Username); err != nil {
			log.Warnf("User verification failed: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "fail",
				"message": "User not found or inactive",
			})
		}

		// Store user data in context locals for access in handlers
		c.Locals("user", userData)
		c.Locals("user_id", userData.UserID)
		c.Locals("id", userData.ID)
		c.Locals("username", userData.Username)
		c.Locals("users_roles", userData.UsersRoles)
		c.Locals("client_code", userData.ClientCode)
		c.Locals("acc_type", userData.AccType)
		c.Locals("branch", userData.Branch)

		log.Infof("User authenticated: %s (ID: %d)", userData.Username, userData.ID)
		return c.Next()
	}
}
