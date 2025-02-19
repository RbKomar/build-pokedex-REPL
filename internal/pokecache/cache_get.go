package pokecache

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	entry, exists := cache.entries[key]
	if exists {
		return entry.val, true
	}
	return nil, false
}
