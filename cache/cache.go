package cache

import (
	"sync"
	"time"
)

// -------- Core Data Structures --------

// Item represents a single cached value with expiration support.
// Store the value, its creation time, and TTL (time-to-live).
type Item struct {
	// YOUR CODE HERE
	value     []byte
	createdAt time.Time
	ttl       time.Duration
}

// Cache is an in-memory key-value store with TTL-based expiration.
// Use a map to store items, protect it with sync.RWMutex for concurrency safety.
type Cache struct {
	mu sync.RWMutex
	// YOUR CODE HERE
	kv map[string]Item
}

// -------- Constructor --------

// NewCache creates a new Cache and starts a background goroutine
// that periodically evicts expired items (garbage collection).
// Accept a cleanup interval (e.g. every 5s) and launch a goroutine with a ticker.
func NewCache(cleanupInterval time.Duration) *Cache {
	// YOUR CODE HERE
	c := Cache{
		kv: make(map[string]Item),
	}
	ticker := time.NewTicker(cleanupInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				c.evictExpired()
			}
		}
	}()
	return &c
}

// -------- Core Operations --------

// Set stores a key-value pair with a given TTL.
//
//	Lock the mutex, create an Item, store it in the map.
//
// If ttl == 0, the item never expires.
func (c *Cache) Set(key string, value []byte, ttl time.Duration) {
	// YOUR CODE HERE
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	i := Item{
		value:     value,
		createdAt: now,
		ttl:       ttl,
	}
	c.kv[key] = i
}

// Get retrieves a value by key.
// RLock the mutex, check if the key exists, check if it's expired.
// Return the value and true if found & valid, nil and false otherwise.
func (c *Cache) Get(key string) ([]byte, bool) {
	// YOUR CODE HERE
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.kv[key]
	if !ok {
		return nil, false
	}

	if item.isExpired() {
		return nil, false
	}

	return item.value, true
}

// Delete removes a key from the cache.
// Lock the mutex, delete the key from the map.
func (c *Cache) Delete(key string) {
	// YOUR CODE HERE
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.kv, key)
}

// Keys returns all non-expired keys currently in the cache.
// RLock, iterate the map, skip expired items.
func (c *Cache) Keys() []string {
	// YOUR CODE HERE
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := []string{}
	for k, v := range c.kv {
		if !v.isExpired() {
			keys = append(keys, k)
		}
	}
	return keys
}

// Count returns the number of non-expired items in the cache.
func (c *Cache) Count() int {
	// YOUR CODE HERE
	count := 0

	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, v := range c.kv {
		if !v.isExpired() {
			count++
		}
	}
	return count
}

// -------- Expiration --------

// isExpired checks whether a given Item has exceeded its TTL.
// Return false if TTL is 0 (never expires), otherwise compare time.Now() with creation + TTL.
func (item *Item) isExpired() bool {
	// YOUR CODE HERE
	if item.ttl == 0 {
		return false
	}
	return time.Now().After(item.createdAt.Add(item.ttl))
}

// evictExpired iterates all items and removes any that are expired.
// Lock the mutex, range over the map, delete expired entries.
func (c *Cache) evictExpired() {
	// YOUR CODE HERE
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.kv {
		if v.isExpired() {
			delete(c.kv, k)
		}
	}
}
