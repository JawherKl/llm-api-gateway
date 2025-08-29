package ratelimit

import (
    "context"
    "time"
    "github.com/go-redis/redis/v8"
    "fmt"
)

type RateLimiter struct {
    client *redis.Client
    limit  int           // max requests
    window time.Duration // per window
}

func New(client *redis.Client, limit int, window time.Duration) *RateLimiter {
    return &RateLimiter{client: client, limit: limit, window: window}
}

func (r *RateLimiter) Allow(ctx context.Context, key string) (bool, error) {
    redisKey := fmt.Sprintf("ratelimit:%s", key)
    count, err := r.client.Incr(ctx, redisKey).Result()
    if err != nil {
        return false, err
    }
    if count == 1 {
        r.client.Expire(ctx, redisKey, r.window)
    }
    return int(count) <= r.limit, nil
}