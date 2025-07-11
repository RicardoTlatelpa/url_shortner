package main

import (
	"github.com/redis/go-redis/v9"
)
func Save(shortID, longURL string) error {
	return redisClient.Set(ctx, shortID, longURL, 0).Err()
}

func Get(shortID string) (string, bool) {
	val, err := redisClient.Get(ctx, shortID).Result()
	if err == redis.Nil {
		return "", false
	} else if err != nil {
		return "", false
	}
	return val, true
} 
