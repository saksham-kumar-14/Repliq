package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	visitors map[string]*Visitor
	mu       sync.Mutex
	rate     int
	burst    int
	period   time.Duration
}

type Visitor struct {
	tokens      int
	lastRequest time.Time
}
