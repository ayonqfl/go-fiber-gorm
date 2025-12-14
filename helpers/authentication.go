package helpers

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	// "github.com/ayonqfl/go-fiber-gorm/database"
	"github.com/ayonqfl/go-fiber-gorm/models"
	// "github.com/gofiber/fiber/v2"
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

// ExtractBearerToken extracts the token from Bearer format
func ExtractBearerToken(authHeader string) string {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}
	return parts[1]
}

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
		log.Warnf("Token parse error: %v", err)
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Check if token is valid
	if !token.Valid {
		log.Warn("Token is not valid")
		return nil, errors.New("invalid token")
	}

	// Extract claims
	claims, ok := token.Claims.(*TokenData)
	if !ok {
		log.Warn("Failed to extract claims")
		return nil, errors.New("failed to extract token claims")
	}

	// // Debug expiration start
	// log.Infof("=== Token Expiration Debug ===")
	// log.Infof("Token ExpiresAt: %v", claims.ExpiresAt)
	// if claims.ExpiresAt != nil {
	// 	log.Infof("Expiration Time: %v", claims.ExpiresAt.Time)
	// 	log.Infof("Current Time: %v", time.Now())
	// 	log.Infof("Time Until Expiry: %v", time.Until(claims.ExpiresAt.Time))
	// 	log.Infof("Is Expired: %v", claims.ExpiresAt.Before(time.Now()))
	// } else {
	// 	log.Warn("WARNING: Token has no expiration time!")
	// }
	// log.Infof("=============================")
	// // Debug expiration end

	// Verify token has not expired
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		log.Warn("Token has expired!")
		return nil, errors.New("token has expired")
	}

	// Verify required fields
	if claims.Username == "" {
		log.Warn("Token missing username")
		return nil, errors.New("invalid token: missing username")
	}

	log.Infof("Token validation successful for user: %s", claims.Username)
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


