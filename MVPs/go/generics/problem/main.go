package main

import (
	"container/list"
)

// Cache structure using Go Generics
type Cache[K comparable, V any] struct {
	capacity int                 // Maximum number of items the cache can hold
	store    map[K]*list.Element // Maps keys to their corresponding list element for fast lookup
	evict    *list.List          // Doubly-linked list to track eviction order (LRU: least recently used at the back)
}

// Entry structure to hold key-value pairs in the cache
type entry[K comparable, V any] struct {
	key   K
	value V
}

// NewCache initializes a new cache with the given capacity
func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{
		capacity: capacity,
		store:    make(map[K]*list.Element),
		evict:    list.New(),
	}
}

// -----------------------------------------------------------------------------

// Put adds an item to the cache. If the cache is at capacity, removes the oldest item first.
func (c *Cache[K, V]) Put(key K, value V) {
	// If key exists, update its value and move to front
	if element, exists := c.store[key]; exists {
		element.Value.(*entry[K, V]).value = value
		c.evict.MoveToFront(element)
		return
	}
	// If at capacity, evict least recently used (at back)
	if len(c.store) >= c.capacity {
		oldest := c.evict.Back()
		if oldest != nil {
			entry := oldest.Value.(*entry[K, V])
			delete(c.store, entry.key)
			c.evict.Remove(oldest)
		}
	}
	// Insert new entry at front
	newEntry := &entry[K, V]{key: key, value: value}
	element := c.evict.PushFront(newEntry)
	c.store[key] = element
}

// Get retrieves an item from the cache by key.
// Returns the value and true if found, otherwise returns the zero value and false.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	if element, exists := c.store[key]; exists {
		return element.Value.(*entry[K, V]).value, true
	}
	var empty V
	return empty, false
}

// Remove deletes an item from the cache by key.
// Returns true if the item was present and removed, otherwise returns false.
func (c *Cache[K, V]) Remove(key K) bool {
	if _, exists := c.store[key]; exists {
		delete(c.store, key)
		return true
	}
	return false
}
