package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rbd := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	limiter := NewTokenBucket(TokenBucketConfig{
		Client:         rbd,
		Capacity:       10,
		RefillRate:     1,
		RefillInterval: 1,
	})

	allowed, remaining, err := limiter.Allow(ctx, "user:123")
	if err != nil {
		log.Fatalf("rate limiter error: %v", err)
	}

	if allowed {
		fmt.Printf("Request allowed. %.0f tokens remaining.\n", remaining)
	} else {
		fmt.Println("Request denied. Rate limit exceeded.")
	}
}
