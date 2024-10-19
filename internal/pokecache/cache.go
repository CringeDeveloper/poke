package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    sync.Mutex
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

func NewCache() *Cache {
	return &Cache{
		mu:    sync.Mutex{},
		entry: make(map[string]cacheEntry),
	}
}
