package cache

import (
    "context"
    "time"
    "github.com/go-redis/redis/v8"
)

type Cache struct {
    client *redis.Client
}

func New(addr string) *Cache {
    rdb := redis.NewClient(&redis.Options{
        Addr: addr,
    })
    return &Cache{client: rdb}
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
    return c.client.Get(ctx, key).Result()
}

func (c *Cache) Set(ctx context.Context, key, value string, ttl time.Duration) error {
    return c.client.Set(ctx, key, value, ttl).Err()
}