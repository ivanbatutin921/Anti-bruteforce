package services

import (
	"errors"
	"math"
	"sync"
	"time"

	pb "github.com/ivanbatutin921/Anti-bruteforce/internal/protobuf"
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

func (tbManager *TokenBucketManager) AddBucketMemory(req *pb.AuthRequest, tb *TokenBucket) error {
	if tbManager == nil {
		return errors.New("TokenBucketManager is nil")
	}
	tbManager.mu.Lock()
	defer tbManager.mu.Unlock()
	loginBuckets, ok := tbManager.tokenBuckets[req.Login]
	if !ok {
		loginBuckets = make(map[string]*TokenBucket)
		tbManager.tokenBuckets[req.Login] = loginBuckets
	}
	if _, ok := loginBuckets[req.Ip]; ok {
		return errors.New("TokenBucket with same login and ip already exists")
	}
	loginBuckets[req.Ip] = tb
	return nil
}

func (tbManager *TokenBucketManager) ResetBucket(req *pb.BucketRequest) error {
	tbManager.mu.Lock()
	defer tbManager.mu.Unlock()
	if loginBuckets, ok := tbManager.tokenBuckets[req.Login]; ok {
		if bucket, ok := loginBuckets[req.Ip]; ok {
			bucket.mu.Lock()
			bucket.tokens = bucket.capacity
			bucket.lastReset = time.Now()
			bucket.mu.Unlock()
			delete(loginBuckets, req.Ip)
			if len(loginBuckets) == 0 {
				delete(tbManager.tokenBuckets, req.Login)
			}
			return nil
		}
	}
	return errors.New("TokenBucket with login and ip not found")
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
