package pokecache

import "time"

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry := cacheEntry{createdAt: time.Now(), val: val}
	cache.entries[key] = entry
}
