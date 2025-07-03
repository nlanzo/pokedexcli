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

	go func() {
		for {
			time.Sleep(interval)
		}
	}()

	return &cache
}