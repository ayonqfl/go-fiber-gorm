package helpers

import (
	"errors"
	"fmt"
	"os"

	"strconv"
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
	token, err := jwt.ParseWithClaims(
		tokenString,
		&TokenData{},
		func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(jwtSecret), nil
		})

	if err != nil || !token.Valid {
		log.Warnf("Token parse error: %v", err)
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*TokenData)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	if claims.Username == "" {
		return nil, errors.New("missing username in token")
	}

	if claims.Username == "" {
		log.Warn("Token missing username")
		return nil, errors.New("invalid token: missing username")
	}

	// Check expiration
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		log.Warnf("Token expired for user: %s", claims.Username)
		return nil, errors.New("token has expired")
	}

	// Redis validation
	alias := os.Getenv("ACCESS_TOKEN_ALIAS")
	if alias == "" {
		log.Warn("ACCESS_TOKEN_ALIAS not set in environment")
		return nil, errors.New("ACCESS_TOKEN_ALIAS not set")
	}

	username := strings.TrimSpace(claims.Username)
	cacheKey := alias + username

	log.Infof("Checking Redis for key: %s", cacheKey)

	redisToken, err := RedisGet(cacheKey)
	if err != nil {
		log.Warnf("Redis error for user %s: %v", username, err)
		return nil, fmt.Errorf("redis error: %w", err)
	}

	if redisToken == "" {
		log.Warnf("No token found in Redis for user: %s", username)
		return nil, errors.New("token not found in cache")
	}

	if redisToken != tokenString {
		log.Warnf("Token mismatch for user %s. Redis token length: %d, Input token length: %d",
			username, len(redisToken), len(tokenString))
		return nil, errors.New("invalid token (token mismatch)")
	}

	// Refresh TTL
	ttlStr := os.Getenv("ACCESS_TOKEN_EXPIRE_MINUTES")
	ttlMinutes, err := strconv.Atoi(ttlStr)
	if err != nil || ttlMinutes <= 0 {
		ttlMinutes = 300 // Default 5 hours
	}

	if err := RedisSetTTL(cacheKey, tokenString, ttlMinutes); err != nil {
		log.Warnf("Failed to refresh token TTL for user %s: %v", username, err)
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
