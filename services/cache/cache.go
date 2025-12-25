package cache

import (
	"context"
	"log"
	"sync"

	"github.com/radovskyb/not-a-rule-engine/services"
)

// Cache is just a generic in-mem cache for now.
type Cache interface {
	Get(ctx context.Context, key string) (any, bool)
	Set(ctx context.Context, key string, value any)
	Delete(ctx context.Context, key string)
}

type cache struct {
	mu    sync.RWMutex
	items map[string]any
}

func New() *cache {
	return &cache{
		items: map[string]any{},
		mu:    sync.RWMutex{},
	}
}

func (c *cache) Call(ctx context.Context, params any) (any, error) {
	log.Println("calling cache service")

	return nil, nil
}

func (c *cache) Funcs() (int, map[string]services.FncParam) {
	log.Println("retrieving allowed funcs + params for cache service")
	return services.CacheServiceID, nil
}

func (c *cache) Get(ctx context.Context, key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	return item, found
}

func (c *cache) Set(ctx context.Context, key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = value
}

func (c *cache) Delete(ctx context.Context, key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}
