package pokecache

import (
	"sync"
	"time"
)


type Cache struct {
	entries map[string]CacheEntry
	mu sync.RWMutex
}

type CacheEntry struct {
	CreatedAt time.Time
	Val      []byte
}


func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries: make(map[string]CacheEntry),
	}

	go cache.reapLoop(interval)

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	entry, ok := c.entries[key]
	c.mu.RUnlock()
	return entry.Val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.entries {
			if time.Since(v.CreatedAt) > interval {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}