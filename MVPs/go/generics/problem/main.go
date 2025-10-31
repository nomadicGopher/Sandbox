package main

import (
	"container/list"
)

// Cache structure using Go Generics
type Cache[K comparable, V any] struct {
	capacity int
	store    map[K]*list.Element
	evict    *list.List
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

func (c *Cache[K, V]) Put(key K, value V) {
	// TODO
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	// TODO
}

func (c *Cache[K, V]) Remove(key K) bool {
	// TODO
}
