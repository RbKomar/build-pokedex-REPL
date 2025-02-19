package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{entries: make(map[string]cacheEntry, 10), mu: sync.RWMutex{}, interval: interval}
	go newCache.reapLoop()
	return newCache
}
