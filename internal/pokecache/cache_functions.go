package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries: make(map[string]CacheEntry),
	}

	go cache.reapLoop(interval)

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val: val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	entry, ok := c.entries[key]
	c.mu.RUnlock()
	return entry.Val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		c.mu.Lock()
		for k, v := range c.entries {
			if time.Since(v.CreatedAt) > interval {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}