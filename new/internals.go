package zcache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type cacheEntry[V any] struct {
	value      V
	expireTime *time.Duration
}

type zCache[K comparable, V any] struct {
	entries map[K]*cacheEntry[V]
	mu      sync.Mutex
	// TODO: add others here. (LRU, TTL, other options, etc.)
}

/**
 *
 *
 */
func newZCache[K comparable, V any]() *zCache[K, V] {
	return &zCache[K, V]{
		entries: make(map[K]*cacheEntry[V]),
		mu:      sync.Mutex{},
	}
}

/**
 *
 *
 */
func (this *zCache[K, V]) Get(key K) (V, bool) {
	this.mu.Lock()
	defer this.mu.Unlock()
	log.Printf("Get -- Searching for entry.")

	var value V

	if len(this.entries) == 0 {
		log.Printf("Get -- Cache is empty.")
		return value, false
	}

	entry, found := this.entries[key]
	if !found {
		log.Printf("Get -- Entry not found.")
		return value, false
	}

	// TODO: Check expiration, move to front of cache is LRU, etc.

	return entry.value, true
}

/**
 *
 *
 */
func (this *zCache[K, V]) Set(key K, value V) error {
	this.mu.Lock()
	defer this.mu.Unlock()
	log.Printf("Set -- Attempting to add entry.")

	_, found := this.entries[key]
	if found {
		log.Printf("Set -- Key already exists.")
		return fmt.Errorf("Key already exists.")
	}

	// TODO: Cache size check, refresh TTL, move to front, etc.
	newEntry := &cacheEntry[V]{
		value:      value,
		expireTime: nil,
	}

	log.Print("DEBUG -- Set -- setting ... ", newEntry.value)
	this.entries[key] = newEntry

	log.Printf("Set -- Key & value pair stored")
	return nil
}
