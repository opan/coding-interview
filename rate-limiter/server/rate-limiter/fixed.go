package ratelimiter

import (
	"sync"
	"time"
)

type fixedRateLimiter struct {
	requests map[string]int
	limit    int
	window   time.Duration
	mutex    sync.Mutex
}

func NewFixedRateLimiter(limit int, window time.Duration) RateLimiter {
	return &fixedRateLimiter{
		requests: make(map[string]int),
		limit:    limit,
		window:   window,
		mutex:    sync.Mutex{},
	}
}

func (rl *fixedRateLimiter) Allow(key string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	count, exists := rl.requests[key]
	if count > rl.limit {
		return false
	}

	if count < rl.limit {
		rl.requests[key]++
	}

	if !exists {
		go rl.resetCounter(key)
	}

	return true
}

func (rl *fixedRateLimiter) resetCounter(key string) {
	time.Sleep(rl.window)
	rl.mutex.Lock()
	delete(rl.requests, key)
	rl.mutex.Unlock()
}
