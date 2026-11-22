package pokecache

import (
	"time"
	"sync"
)



type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	Data map[string]cacheEntry
	mu sync.Mutex
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.Data {
		diff := time.Now().Sub(value.createdAt)
		if diff > interval {
			delete(c.Data, key)
		}

	}

}

func (c *Cache) NewCache(interval time.Duration) {
	c.Data = make(map[string]cacheEntry)
	ticker := time.NewTicker(interval)
	go func() {
		defer ticker.Stop()
		for range ticker.C {
			c.reapLoop(interval)
		}
	}()}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.Data[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.Data[key]
	if ok {
		return  entry.val, ok
	}
	return nil, ok

}
