package main

import (
	"crypto/md5"

	"encoding/hex"
	"fmt"
	"log"

	

	"github.com/redis/go-redis/v9"
)


func generateShortURL(longURL string) string {
    // Generate MD5 hash of the long URL
    hash := md5.Sum([]byte(longURL))
    hashStr := hex.EncodeToString(hash[:])

    // Use the first 8 characters of the hash as the short URL
    shortURL := hashStr[:8]

    return shortURL
}
// Function to store the mapping between a short URL and a long URL in Redis
func storeURLMapping(redisClient *redis.Client, shortURL, longURL string) error {
	// Use the SET command to store the mapping in Redis
	err := redisClient.Set(ctx, shortURL, longURL, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to store URL mapping in Redis: %v", err)
	}
	return nil
}

// Function to retrieve the long URL corresponding to a given short URL from Redis
func getLongURL(redisClient *redis.Client, shortURL string) (string, error) {
	// Use the GET command to retrieve the long URL from Redis
	longURL, err := redisClient.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("long URL not found for short URL: %s", shortURL)
	} else if err != nil {
		return "", fmt.Errorf("failed to retrieve long URL from Redis: %v", err)
	}

	// Log the retrieved long URL
	log.Printf("Retrieved long URL from Redis for short URL %s: %s", shortURL, longURL)

	return longURL, nil
}
