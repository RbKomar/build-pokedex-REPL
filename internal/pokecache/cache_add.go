package pokecache

import (
	"log"
	"time"
)

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry := cacheEntry{createdAt: time.Now(), val: val}
	cache.entries[key] = entry
	log.Println("Added to cache for key: ", key)
}
