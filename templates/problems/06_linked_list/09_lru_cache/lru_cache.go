package lru_cache

// LRUCache implements a Least Recently Used cache.
type LRUCache struct {
}

// Constructor creates a new LRUCache with the given capacity.
func Constructor(capacity int) LRUCache {
	return LRUCache{}
}

// Get returns the value of the key if it exists, otherwise -1.
func (c *LRUCache) Get(key int) int {
	return -1
}

// Put inserts or updates the key-value pair, evicting the LRU item if at capacity.
func (c *LRUCache) Put(key int, value int) {
}
