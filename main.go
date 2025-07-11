package main

import (
	"context"
	"log"
	"net/http"

	gen "github.com/RicardoTlatelpa/uniqueidgen"
	"github.com/redis/go-redis/v9"
)
var (
	redisClient *redis.Client
	ctx = context.Background()

	idGen *gen.Gen	
	baseURL = "http://localhost:8080/"
)

func main() {
	var err error

	redisClient = redis.NewClient(&redis.Options{
		Addr:		"localhost:6379",
		Password:	"",
		DB:		0,
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
	idGen, err = gen.NewGen(1)
	if err != nil {
		log.Fatalf("failed to create generator: %v", err)
	}

	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/", handleRedirect)

	log.Println("server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
