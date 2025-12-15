package middleware

import (
	"github.com/ayonqfl/go-fiber-gorm/database"
	"github.com/ayonqfl/go-fiber-gorm/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func AuthMiddleware() fiber.Handler {
	// Get JWT secret from environment variable
	jwtSecret := helpers.GetJWTSecret()

	return func(c *fiber.Ctx) error {
		// Extract token from Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Warn("Missing Authorization header for path: ", c.Path())
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  401,
				"message": "Invalid token",
				"errors":  nil,
				"data":    []string{},
			})
		}

		// Extract Bearer token
		token := helpers.ExtractBearerToken(authHeader)
		if token == "" {
			log.Warn("Invalid Authorization header format for path: ", c.Path())
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  401,
				"message": "Invalid token",
				"errors":  nil,
				"data":    []string{},
			})
		}

		// Validate token
		userData, err := helpers.ValidateToken(token, jwtSecret)
		if err != nil {
			log.Warnf("Token validation failed: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  401,
				"message": "Invalid token",
				"errors":  nil,
				"data":    []string{},
			})
		}

		// Verify user exists in database and is active
		if err := helpers.VerifyUserExists(database.GetQtraderDB(), userData.Username); err != nil {
			log.Warnf("User verification failed: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  401,
				"message": "User not found or inactive",
				"errors":  nil,
				"data":    []string{},
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
