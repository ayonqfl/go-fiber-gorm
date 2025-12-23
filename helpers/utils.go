package helpers

import (
	"errors"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetCurrentUser(c *fiber.Ctx) (*TokenData, error) {
	user := c.Locals("user")
	if user == nil {
		return nil, errors.New("user not found in context")
	}

	currentUser, ok := user.(*TokenData)
	if !ok {
		return nil, errors.New("invalid user type")
	}

	return currentUser, nil
}

// Helper functions to get individual fields
func GetUserID(c *fiber.Ctx) (string, error) {
	userID := c.Locals("user_id")
	if userID == nil {
		return "", errors.New("user_id not found")
	}

	id, ok := userID.(string)
	if !ok {
		return "", errors.New("invalid user_id type")
	}

	return id, nil
}

func GetUsername(c *fiber.Ctx) (string, error) {
	username := c.Locals("username")
	if username == nil {
		return "", errors.New("username not found")
	}

	name, ok := username.(string)
	if !ok {
		return "", errors.New("invalid username type")
	}

	return name, nil
}

// GetEnvBool retrieves boolean value from environment variable
func GetEnvBool(key string, defaultValue bool) bool {
	value := strings.ToLower(strings.TrimSpace(os.Getenv(key)))
	if value == "" {
		return defaultValue
	}
	return value == "true" || value == "1" || value == "yes"
}
