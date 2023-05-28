package memorycache

import (
	"sync"
	"time"
)

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, expiration time.Duration) error
}

type MemoryCache struct {
	cache *sync.Map
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		cache: &sync.Map{},
	}
}

func (c *MemoryCache) Get(key string) (interface{}, error) {
	value, ok := c.cache.Load(key)
	if !ok {
		return nil, nil
	}
	return value, nil
}

func (c *MemoryCache) Set(key string, value interface{}, expiration time.Duration) error {
	c.cache.Store(key, value)
	time.AfterFunc(expiration, func() {
		c.cache.Delete(key)
	})
	return nil
}
