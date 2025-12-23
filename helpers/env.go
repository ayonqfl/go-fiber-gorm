package helpers

import (
	"os"
	"strings"
)

// GetEnvBool retrieves boolean value from environment variable
func GetEnvBool(key string, defaultValue bool) bool {
	value := strings.ToLower(strings.TrimSpace(os.Getenv(key)))
	if value == "" {
		return defaultValue
	}
	return value == "true" || value == "1" || value == "yes"
}
