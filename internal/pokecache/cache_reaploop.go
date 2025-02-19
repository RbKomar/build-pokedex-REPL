package pokecache

import "time"

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()
	for range ticker.C {
		cache.mu.Lock()
		for key, val := range cache.entries {
			if time.Since(val.createdAt) > cache.interval {
				delete(cache.entries, key)
			}
		}
		cache.mu.Unlock()
	}
}
