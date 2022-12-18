package memorycache

import (
	"context"
	"time"

	"github.com/patrickmn/go-cache"
)

type Client struct {
	cache *cache.Cache
}

func New(defaultExpiration time.Duration, cleanupInterval time.Duration) Client {
	return Client{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (c Client) GetInt(ctx context.Context, key string) int {
	value, found := c.cache.Get(key)

	if !found {
		return 0
	}

	return value.(int)
}

func (c Client) SetInt(ctx context.Context, key string, value, ttl int) {
	c.cache.Set(key, value, time.Duration(ttl)*time.Minute)
}
