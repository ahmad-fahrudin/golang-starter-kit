package middleware

import (
	"net/http"
	"sync"
	"time"

	"golang-starter-kit/internal/models"

	"github.com/gin-gonic/gin"
)

// RateLimitData holds the rate limiting information for an IP
type RateLimitData struct {
	Count     int
	LastReset time.Time
}

// RateLimiter manages rate limiting for login attempts
type RateLimiter struct {
	mu          sync.Mutex
	limits      map[string]*RateLimitData
	maxAttempts int
	window      time.Duration
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(maxAttempts int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limits:      make(map[string]*RateLimitData),
		maxAttempts: maxAttempts,
		window:      window,
	}
}

// IsAllowed checks if the IP is allowed to make a request
func (rl *RateLimiter) IsAllowed(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	data, exists := rl.limits[ip]

	if !exists || now.Sub(data.LastReset) > rl.window {
		// Reset or initialize
		rl.limits[ip] = &RateLimitData{
			Count:     1,
			LastReset: now,
		}
		return true
	}

	if data.Count >= rl.maxAttempts {
		return false
	}

	data.Count++
	return true
}

// RateLimitMiddleware creates a rate limiting middleware for login
func RateLimitMiddleware(maxAttempts int, window time.Duration) gin.HandlerFunc {
	limiter := NewRateLimiter(maxAttempts, window)

	return func(c *gin.Context) {
		ip := c.ClientIP()

		if !limiter.IsAllowed(ip) {
			c.JSON(http.StatusTooManyRequests, models.ErrorResponse{
				Error:   "rate_limit_exceeded",
				Message: "Too many login attempts. Please try again later.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
