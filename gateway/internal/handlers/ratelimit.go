package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/JawherKl/gateway/internal/ratelimit"
    "context"
)

func RateLimitMiddleware(limiter *ratelimit.RateLimiter) gin.HandlerFunc {
    return func(c *gin.Context) {
        apiKey := c.GetHeader("Authorization")
        allowed, err := limiter.Allow(context.Background(), apiKey)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Rate limiter error"})
            return
        }
        if !allowed {
            c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
            return
        }
        c.Next()
    }
}