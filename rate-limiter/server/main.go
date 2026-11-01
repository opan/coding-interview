package main

import (
	"log"
	"os"
	"time"

	ratelimiter "github.com/coding-interview/rate-limiter/server/rate-limiter"
)

func main() {
	rateLimiterType := os.Getenv("RATE_LIMITER_TYPE")
	var limiter ratelimiter.RateLimiter

	switch rateLimiterType {
	case "fixed":
		limiter = ratelimiter.NewFixedRateLimiter(5, 1*time.Minute)
		// Use the limiter as needed
	default:
		panic("Unsupported rate limiter type")
	}

	allowed := limiter.Allow("user123")

	if allowed {
		log.Println("Request allowed")
	} else {
		log.Println("Request denied")
	}

}
