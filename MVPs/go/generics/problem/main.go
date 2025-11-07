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

// Put adds or updates an item in the cache.
func (c *Cache[K, V]) Put(key K, value V) {
	newEntry := &entry[K, V]{key: key, value: value}
	elem := c.evict.PushFront(newEntry)
	c.store[key] = elem
}

// Get retrieves an item from the cache by key.
// Returns the value and true if found, otherwise returns the zero value and false.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	if elem, ok := c.store[key]; ok {
		return elem.Value.(*entry[K, V]).value, true
	}
	var zero V
	return zero, false
}

// Remove deletes an item from the cache by key.
// Returns true if the item was present and removed, otherwise returns false.
func (c *Cache[K, V]) Remove(key K) bool {
	_, ok := c.store[key]
	if ok {
		delete(c.store, key)
	}
	return ok
}
