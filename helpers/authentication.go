package helpers

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	// "github.com/ayonqfl/go-fiber-gorm/database"
	"github.com/ayonqfl/go-fiber-gorm/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// TokenData represents the JWT payload structure from your FastAPI login response
type TokenData struct {
	ID            int    `json:"id"`
	UserID        string `json:"user_id"`
	Username      string `json:"username"`
	Name          string `json:"name"`
	UsersRoles    string `json:"users_roles"`
	AccType       string `json:"acc_type"`
	DealerGroupID string `json:"dealer_group_id"`
	MarginAllowed bool   `json:"margin_allowed"`
	Branch        string `json:"branch"`
	ClientCode    string `json:"client_code,omitempty"`
	DeviceOS      string `json:"device_os"`
	jwt.RegisteredClaims
}

// ShouldSkipAuth checks if the path should skip authentication
func ShouldSkipAuth(path string, skipPaths []string, skipKeywords []string) bool {
	// Check if path is in skip paths
	for _, skipPath := range skipPaths {
		if strings.HasPrefix(path, skipPath) {
			return true
		}
	}

	// Check if path contains skip keywords
	for _, keyword := range skipKeywords {
		if strings.Contains(path, keyword) {
			return true
		}
	}

	return false
}

// ExtractBearerToken extracts the token from Bearer format
func ExtractBearerToken(authHeader string) string {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}
	return parts[1]
}

// ValidateToken validates the JWT token and returns the token data
func ValidateToken(tokenString string, jwtSecret string) (*TokenData, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &TokenData{}, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Check if token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extract claims
	claims, ok := token.Claims.(*TokenData)
	if !ok {
		return nil, errors.New("failed to extract token claims")
	}

	// Verify token has not expired
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token has expired")
	}

	// Verify required fields
	if claims.Username == "" {
		return nil, errors.New("invalid token: missing username")
	}

	return claims, nil
}

// VerifyUserExists checks if the user exists in the database and is active
func VerifyUserExists(db *gorm.DB, username string) error {
	var user models.User
	result := db.Where("LOWER(username) = LOWER(?)", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return result.Error
	}

	// Check if user account is active
	if user.AccountStatus != "active" {
		return errors.New("user account is not active")
	}

	return nil
}

// GetJWTSecret retrieves JWT secret from environment variable
func GetJWTSecret() string {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET_KEY not found in environment variables")
	}
	return jwtSecret
}

// GetUserFromContext retrieves user data from context
func GetUserFromContext(c *fiber.Ctx) *TokenData {
	user := c.Locals("user")
	if user == nil {
		return nil
	}
	userData, ok := user.(*TokenData)
	if !ok {
		return nil
	}
	return userData
}

// GetUserIDFromContext retrieves user_id from context as string
func GetUserIDFromContext(c *fiber.Ctx) string {
	userID := c.Locals("user_id")
	if userID == nil {
		return ""
	}
	userIDStr, ok := userID.(string)
	if !ok {
		return ""
	}
	return userIDStr
}

// GetUsernameFromContext retrieves username from context
func GetUsernameFromContext(c *fiber.Ctx) string {
	username := c.Locals("username")
	if username == nil {
		return ""
	}
	usernameStr, ok := username.(string)
	if !ok {
		return ""
	}
	return usernameStr
}

// GetUserRoleFromContext retrieves user role from context
func GetUserRoleFromContext(c *fiber.Ctx) string {
	role := c.Locals("users_roles")
	if role == nil {
		return ""
	}
	roleStr, ok := role.(string)
	if !ok {
		return ""
	}
	return roleStr
}

// GetClientCodeFromContext retrieves client code from context
func GetClientCodeFromContext(c *fiber.Ctx) string {
	clientCode := c.Locals("client_code")
	if clientCode == nil {
		return ""
	}
	clientCodeStr, ok := clientCode.(string)
	if !ok {
		return ""
	}
	return clientCodeStr
}

// GetUserIDIntFromContext retrieves user id (int) from context
func GetUserIDIntFromContext(c *fiber.Ctx) int {
	id := c.Locals("id")
	if id == nil {
		return 0
	}
	idInt, ok := id.(int)
	if !ok {
		return 0
	}
	return idInt
}

