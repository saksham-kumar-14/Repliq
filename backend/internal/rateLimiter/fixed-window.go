package ratelimiter

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func New(rate int, burst int, period time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*Visitor),
		rate:     rate,
		burst:    burst,
		period:   period,
	}

	go func() {
		for {
			time.Sleep(time.Minute)
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
		for now.Sub(v.lastRequest) > rl.period*2 {
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
			v = &Visitor{tokens: rl.burst, lastRequest: time.Now()}
			rl.visitors[ip] = v
		}

		now := time.Now()
		elapsed := now.Sub(v.lastRequest)
		// refill tokens
		refill := int(elapsed.Seconds() * float64(rl.rate) / rl.period.Seconds())
		if refill > 0 {
			v.tokens += refill
			if v.tokens > rl.burst {
				v.tokens = rl.burst
			}
		}
		v.lastRequest = now

		if v.tokens > 0 {
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
