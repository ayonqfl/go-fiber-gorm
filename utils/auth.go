package utils

import (
	"errors"
	"github.com/ayonqfl/go-fiber-gorm/helpers"
	"github.com/gofiber/fiber/v2"
)

func GetCurrentUser(c *fiber.Ctx) (*helpers.TokenData, error) {
	user := c.Locals("user")
	if user == nil {
		return nil, errors.New("user not found in context")
	}

	currentUser, ok := user.(*helpers.TokenData)
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