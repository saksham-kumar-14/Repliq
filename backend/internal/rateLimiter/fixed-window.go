package ratelimiter

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func New(rate int, burst int, period time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*Visitor),
		rate:     float64(rate),
		burst:    burst,
		period:   period,
	}

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			rl.cleanup()
		}
	}()

	return rl
}

func (rl *RateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	for ip, v := range rl.visitors {
		if now.Sub(v.lastRequest) > rl.period*2 {
			delete(rl.visitors, ip)
		}
	}
}

func (rl *RateLimiter) Limit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()

		rl.mu.Lock()
		v, exists := rl.visitors[ip]
		if !exists {
			v = &Visitor{
				tokens:      float64(rl.burst),
				lastRequest: time.Now(),
			}
			rl.visitors[ip] = v
		}

		now := time.Now()
		elapsed := now.Sub(v.lastRequest)
		v.lastRequest = now

		refillTokens := (elapsed.Seconds() / rl.period.Seconds()) * rl.rate
		v.tokens += refillTokens
		if v.tokens > float64(rl.burst) {
			v.tokens = float64(rl.burst)
		}

		if v.tokens >= 1 {
			v.tokens--
			rl.mu.Unlock()
			return next(c)
		}

		rl.mu.Unlock()
		return c.JSON(http.StatusTooManyRequests, map[string]string{
			"error": "rate limit exceeded",
		})
	}
}
