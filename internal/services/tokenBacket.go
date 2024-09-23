package services

import (
	"math"
	"sync"
	"time"
)

type TokenBucketManager struct {
	tokenBuckets map[string]map[string]*TokenBucket
	mu           sync.Mutex
}

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

	now := time.Now()
	elapsed := time.Since(tb.lastReset)
	tb.tokens += int32(math.Floor(elapsed.Seconds() * tb.rate))
	tb.tokens = min(tb.tokens, tb.capacity)
	tb.lastReset = now.Add(0)

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}


func NewTokenBucketManager() *TokenBucketManager {
	return &TokenBucketManager{
		tokenBuckets: make(map[string]map[string]*TokenBucket),
		mu:           sync.Mutex{},
	}
}

func (tbManager *TokenBucketManager) AddBucketMemory(login, ip string, tb *TokenBucket) {
	tbManager.mu.Lock()
	defer tbManager.mu.Unlock()
	loginBuckets, ok := tbManager.tokenBuckets[login]
	if !ok {
		loginBuckets = make(map[string]*TokenBucket)
		tbManager.tokenBuckets[login] = loginBuckets
	}
	loginBuckets[ip] = tb
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
