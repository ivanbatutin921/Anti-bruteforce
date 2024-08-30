package services

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity  int32
	rate      float64 //скорость добавления
	tokens    int32
	lastReset time.Time
	mu        sync.Mutex
}

func NewTokenbucket(capacity int32, rate float64) *TokenBucket {
	return &TokenBucket{
		capacity:  capacity,
		rate:      rate,
		tokens:    capacity,
		lastReset: time.Now(),
		mu:        sync.Mutex{},
	}
}

func (tb *TokenBucket) Take(ip string, tokens int32) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now
	elapsed := time.Since(tb.lastReset)
	tb.tokens += int32(elapsed.Seconds() * float64(tb.rate))
	tb.tokens = min(tb.tokens, tb.capacity)
	tb.lastReset = now()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
