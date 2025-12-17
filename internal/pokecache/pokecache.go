package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val 	  []byte
}

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}


func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string)([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if val, ok := c.entries[key]; !ok {
		return nil, false
	} else {
		return val.val, true
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.entries {
			cutoff := time.Now().Add(-interval)
			if v.createdAt.Before(cutoff) {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache {
		entries: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}
