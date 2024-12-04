package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	ttl  time.Duration
	mu   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	return Cache{map[string]cacheEntry{}, interval, &sync.Mutex{}}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.data[key]; ok {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop() {

}
