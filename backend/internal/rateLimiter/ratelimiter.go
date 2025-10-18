package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	visitors map[string]*Visitor
	mu       sync.Mutex
	rate     float64
	burst    int
	period   time.Duration
}

type Visitor struct {
	tokens      float64
	lastRequest time.Time
}
