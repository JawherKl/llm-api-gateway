package main

import (
    "github.com/gin-gonic/gin"
    "github.com/JawherKl/gateway/internal/config"
    "github.com/JawherKl/gateway/internal/handlers"
    "github.com/JawherKl/gateway/internal/cache"
    "github.com/go-redis/redis/v8"
    "github.com/JawherKl/gateway/internal/ratelimit"
    "time"
)

func main() {
    cfg := config.Load()
    r := gin.Default()

    // Redis client for cache and rate limiter
    redisClient := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
    // Initialize Redis cache
    redisCache := cache.New(cfg.RedisAddr)
    limiter := ratelimit.New(redisClient, 60, time.Minute) // 60 req/min per key

    // Register the Auth middleware globally or for specific routes
    r.Use(handlers.AuthMiddleware())
    r.Use(handlers.RateLimitMiddleware(limiter))

    // Routes
    r.POST("/gateway/query", handlers.QueryHandler(cfg, redisCache))

    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Run(":" + cfg.ServerPort)
}
