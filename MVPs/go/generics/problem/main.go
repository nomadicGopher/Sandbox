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

// TODO: Put inserts or updates the value for the given key in the cache.
// If the key already exists, its value is updated and it is marked as most recently used.
// If the cache is at capacity, the least recently used item is evicted.
func (c *Cache[K, V]) Put(key K, value V) {
	if elem, ok := c.store[key]; ok {
		// Update value and move to front
		elem.Value.(*entry[K, V]).value = value
		c.evict.MoveToFront(elem)
		return
	}

	// Evict if necessary
	if len(c.store) >= c.capacity {
		if backElem := c.evict.Back(); backElem != nil {
			evictEntry := backElem.Value.(*entry[K, V])
			delete(c.store, evictEntry.key)
			c.evict.Remove(backElem)
		}
	}

	// Insert new entry
	newEntry := &entry[K, V]{key: key, value: value}
	elem := c.evict.PushFront(newEntry)
	c.store[key] = elem
}

// TODO: Get retrieves the value associated with the given key from the cache.
// If the key exists, it is marked as most recently used and its value is returned.
// Returns the value and true if found, otherwise returns the zero value and false.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	var zero V
	if elem, ok := c.store[key]; ok {
		c.evict.MoveToFront(elem)
		return elem.Value.(*entry[K, V]).value, true
	}
	return zero, false
}

// TODO: Remove deletes the entry associated with the given key from the cache.
// Returns true if the key was present and removed, otherwise returns false.
func (c *Cache[K, V]) Remove(key K) bool {
	if elem, ok := c.store[key]; ok {
		c.evict.Remove(elem)
		delete(c.store, key)
		return true
	}
	return false
}
