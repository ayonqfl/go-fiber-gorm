package helpers

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx         = context.Background()
	RedisClient *redis.Client
)

func InitRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	pass := os.Getenv("REDIS_PASS")
	db   := mustInt(os.Getenv("REDIS_DB"))

	if host == "" || port == "" {
		log.Fatal("REDIS_HOST or REDIS_PORT not set")
	}

	addr := host + ":" + port

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	// Test connection
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Redis connection failed (%s): %v", addr, err)
	}

	log.Printf("Redis connected successfully (%s, db=%d)", addr, db)
}

func mustInt(v string) int {
	i, _ := strconv.Atoi(v)
	return i
}

// RedisGet retrieves a value and handles JSON decoding (matching FastAPI behavior)
func RedisGet(key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	// Try to unmarshal as JSON string (FastAPI stores as JSON)
	var jsonStr string
	if err := json.Unmarshal([]byte(val), &jsonStr); err == nil {
		// Successfully decoded as JSON string
		return jsonStr, nil
	}

	// If not JSON, return as is
	return val, nil
}

// RedisSetTTL stores a value with TTL, encoding as JSON (matching FastAPI behavior)
func RedisSetTTL(key string, value string, ttlMinutes int) error {
	// Encode as JSON to match FastAPI behavior
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return RedisClient.Set(
		ctx,
		key,
		jsonData,
		time.Duration(ttlMinutes)*time.Minute,
	).Err()
}

func RedisDelete(key string) error {
	return RedisClient.Del(ctx, key).Err()
}