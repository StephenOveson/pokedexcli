package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c Cache) readLoop(t time.Duration) {
	ticker := time.NewTicker(t)

	for range ticker.C {
		c.reap(time.Now().UTC(), t)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}

func NewCache(t time.Duration) Cache {
	c := Cache{
		cache: map[string]cacheEntry{},
		mu:    &sync.Mutex{},
	}

	go c.readLoop(t)
	return c
}
