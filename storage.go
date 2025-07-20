package main

import (
	"fmt"
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

func incrementClick(shortID string) {
	key := fmt.Sprintf("clicks:%s", shortID)
	redisClient.Incr(ctx, key)
}

func getClicks(shortID string) int64 {
	key := fmt.Sprintf("clicks:%s", shortID)
	val, err := redisClient.Get(ctx, key).Int64()

	if err != nil {
		return 0
	}
	return val
}
