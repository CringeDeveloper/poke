package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    *sync.RWMutex
	entry map[string]cacheEntry
}

func (c *Cache) Add(url string, data []byte) {
	go func() {
		entry := cacheEntry{
			createdAt: time.Now(),
			val:       data}

		defer c.mu.Unlock()
		c.mu.Lock()
		c.entry[url] = entry
	}()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if cache, ok := c.entry[key]; ok {
		return cache.val, true
	}
	return []byte{}, false
}

func (c *Cache) reepLoop(cacheInterval time.Duration) {
	tick := time.Tick(cacheInterval)
	for {
		select {
		case <-tick:
			c.mu.Lock()
			defer c.mu.Unlock()
			for k, v := range c.entry {
				if time.Since(v.createdAt) > cacheInterval {
					delete(c.entry, k)
				}
			}
		}
	}
}

func NewCache(cacheInterval time.Duration) Cache {
	cache := Cache{
		mu:    &sync.RWMutex{},
		entry: make(map[string]cacheEntry),
	}

	go cache.reepLoop(cacheInterval)
	return cache
}
